package src

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func createQcowFromBase(base, out string) (string, error) {
	//cmdArgs := fmt.Sprintf("qemu-img create -f qcow2 -o backing_file=%s %s", base, out)
	cmdArgs := fmt.Sprintf("cp %s %s", base, out)
	fmt.Println(cmdArgs)
	_ = runCmd(cmdArgs)
	return filepath.Abs(out)
}

func runCmd(cmd string) error {
	c := exec.Command("/bin/bash", "-c", cmd)
	err := c.Run()
	return err
}

// 检查path路径是否存在以及是否是文件, nil 为是文件且存在
func checkPath(path string) error {
	isExist, isFile := checkFileAndDirExist(path)
	if !isExist {
		return errors.New("file not exist")
	}
	if !isFile {
		return errors.New("path is a dir")
	}
	return nil
}

//检查路径存在以及是否是文件
func checkFileAndDirExist(path string) (isExist, isFile bool) {
	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, false
		}
		return true, false
	}
	return true, !f.IsDir()
}
