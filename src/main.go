package src

import (
	"embed"
	"errors"
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
	
	// Load base XML configuration for VM
	baseXml, err := defaultVm.ReadFile("defaultVm.xml")
	if err != nil {
		return fmt.Errorf("failed to read base XML: %w", err)
	}
	
	err = os.WriteFile("./"+vmUuid+".xml", baseXml, 0666)
	if err != nil {
		return fmt.Errorf("failed to write VM XML file: %w", err)
	}
	defer func() {
		if err := os.Remove("./" + vmUuid + ".xml"); err != nil {
			fmt.Printf("warning: failed to clean up XML file: %v\n", err)
		}
	}()
	
	cpu := uint(4)
	memory := uint(4)

	// Parse Dockerfile for VM configuration
	parseFile, err := dockerfileparser.ParseFile(file)
	if err != nil {
		return fmt.Errorf("failed to parse Dockerfile: %w", err)
	}
	if qcow == "" {
		qcow = from(parseFile)
		if qcow == "" {
			qcow = C74String
		}
	}

	// Attempt to remove currentQcow if it exists
	if err := os.Remove(currentQcow); err != nil && !errors.Is(err, os.ErrNotExist) {
		fmt.Printf("warning: failed to remove current qcow file: %v\n", err)
	}
	
	// Create qcow image from base
	currentQcow, err = createQcowFromBase(qcow, currentQcow)
	if err != nil {
		return fmt.Errorf("failed to create qcow image: %w", err)
	}

	// Initialize VM
	vm, err := govirsh.NewVm(vmName, vmUuid, mac, qcow, emulato, "./"+vmUuid+".xml", currentQcow, cpu, memory)
	if err != nil {
		return fmt.Errorf("failed to initialize VM: %w", err)
	}
	defer func() {
		if err := vm.DestroyVm(); err != nil {
			fmt.Printf("warning: failed to destroy VM: %v\n", err)
		}
	}()

	// Check if VM acquired IP
	_, err = vm.CheckIp()
	if err != nil {
		return fmt.Errorf("failed to acquire IP for VM: %w", err)
	}

	workDir := "./"
	for _, v := range parseFile.All {
		switch v[0] {
		case dockerfileparser.RunString:
			// Execute remote SSH command
			cmd, out, err := run(workDir, v[1], vm.GetIp())
			if err != nil {
				fmt.Printf("error running command %s: %v\n", cmd, err)
			} else {
				fmt.Printf("%s:%s out:%s\n", dockerfileparser.RunString, cmd, out)
			}
		case dockerfileparser.WorkdirString:
			workDir = v[1]
			fmt.Printf("%s: %s\n", dockerfileparser.WorkdirString, workDir)
			// Ensure directory exists in VM
			if _, _, err := run("/", "mkdir -p "+workDir, vm.GetIp()); err != nil {
				fmt.Printf("error creating workdir %s on VM: %v\n", workDir, err)
			}
		case dockerfileparser.CopyString:
			localfile, remotefile := strings.Split(v[1], " ")[0], strings.Split(v[1], " ")[1]
			out, err := copy(workDir, localfile, remotefile, vm.GetIp())
			if err != nil {
				fmt.Printf("error copying file %s to %s: %v\n", localfile, remotefile, err)
			} else {
				fmt.Printf("%s: %s\n", dockerfileparser.CopyString, out)
			}
		}
	}
	
	// Run sync command to ensure all writes are flushed
	if _, _, err := run(workDir, "sync", vm.GetIp()); err != nil {
		fmt.Printf("warning: sync command failed: %v\n", err)
	}

	// Optional pause before VM destruction
	if pause {
		fmt.Println("pause VM, press Enter to continue and destroy...")
		fmt.Scanln()
	}
	
	return nil
}

// Extract qcow2 path from Dockerfile
func from(d dockerfileparser.DockerFileContent) string {
	if len(d.From) == 0 {
		return ""
	}
	return d.From[0]
}

// Run command on VM and return output
func run(workDir, cmd, ip string) (string, string, error) {
	realCmd := "cd " + workDir + ";" + cmd
	out, err := connectAndRunCmd(ip, 100, 100, []string{realCmd})
	if err != nil {
		return realCmd, "", fmt.Errorf("command execution failed: %w", err)
	}
	return realCmd, out[0], nil
}

// Copy file to VM
func copy(workDir, localfile, remotefile, ip string) (string, error) {
	realRemote := filepath.Join(workDir, remotefile)
	err := connectAndUpload(localfile, realRemote, ip, 100, 100, []string{})
	if err != nil {
		return "", fmt.Errorf("file copy failed: %w", err)
	}
	return fmt.Sprintf("CP %s %s", localfile, realRemote), nil
}
