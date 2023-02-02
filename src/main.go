package src

import (
	"embed"
	"fmt"
	"github.com/eust-w/dockerfileparser"
	"github.com/eust-w/govirsh"
	"os"
	"path/filepath"
	"strings"
)

const (
	C74String = "c74"
)

//go:embed defaultVm.xml
var defaultVm embed.FS

func Qcow2file(qcow, currentQcow, file string, pause bool) error {
	vmUuid, _ := generateUuid()
	vmName := vmUuid
	mac := ""
	emulato := "/usr/libexec/qemu-kvm"
	baseXml, _ := defaultVm.ReadFile("defaultVm.xml")
	err := os.WriteFile("./"+vmUuid+".xml", baseXml, 0666)
	defer func() {
		runCmd("rm -rf " + "./" + vmUuid + ".xml")
	}()
	cpu := uint(4)
	memory := uint(4)

	parseFile, err := dockerfileparser.ParseFile(file)
	if err != nil {
		return err
	}
	if qcow == "" {
		qcow = from(parseFile)
		if qcow == "" {
			qcow = C74String
		}
	}
	_ = os.Remove(currentQcow)
	currentQcow, err = createQcowFromBase(qcow, currentQcow)
	vm, err := govirsh.NewVm(vmName, vmUuid, mac, qcow, emulato, "./"+vmUuid+".xml", currentQcow, cpu, memory)
	if err != nil {
		return err
	}
	_, err = vm.CheckIp()
	if err != nil {
		return err
	}

	workDir := "./"
	for _, v := range parseFile.All {
		switch v[0] {
		case dockerfileparser.RunString:
			// remote ssh run
			cmd, out, err := run(workDir, v[1], vm.GetIp())
			fmt.Printf("%s:%s out:%s err:%v\n", dockerfileparser.RunString, cmd, out, err)
		case dockerfileparser.WorkdirString:
			workDir = v[1]
			fmt.Println(dockerfileparser.WorkdirString, ":", workDir)
			// 在虚拟机中创建此目录
		case dockerfileparser.CopyString:
			localfile, remotefile := strings.Split(v[1], " ")[0], strings.Split(v[1], " ")[1]
			out, err := copy(workDir, localfile, remotefile, vm.GetIp())
			fmt.Println(dockerfileparser.CopyString, ":", out, err)

		}
	}
	// 落盘
	run(workDir, "sync", vm.GetIp())
	//可选压缩镜像
	//可选暂停删除
	if pause {
		fmt.Println("pause vm, enter any key to destroy")
		fmt.Scanln()
	}
	return vm.DestroyVm()
}

//获取qcow2
func from(d dockerfileparser.DockerFileContent) string {
	if len(d.From) <= 0 {
		return ""
	}
	return d.From[0]
}

func run(workDir, cmd, ip string) (string, string, error) {
	realCmd := "cd " + workDir + ";" + cmd
	out, err := connectAndRunCmd(ip, 100, 100, []string{realCmd})
	return realCmd, out[0], err
}

func copy(workDir, localfile, remotefile, ip string) (string, error) {
	realRemote := filepath.Join(workDir, remotefile)
	return fmt.Sprintf("CP %s %s", localfile, realRemote), connectAndUpload(localfile, realRemote, ip, 100, 100, []string{})
}
