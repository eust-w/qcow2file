package src

import (
	"errors"
	"github.com/eust-w/xssh"
	"time"
)

type SshConnect struct {
	ip       string
	user     string
	name     string
	password string
	port     string
	client   *xssh.Client
}

func (s SshConnect) upload(ori, dst string, ignored []string) error {
	err := checkPath(ori)
	if err != nil {
		err = s.client.UploadDir(ori, dst, ignored)
	} else {
		err = s.client.UploadFile(ori, dst)
	}
	return err
}

func (s SshConnect) close() error {
	s.client.Close()
	return nil
}

func (s *SshConnect) init(name, user, ip, port, password string) error {
	s.name, s.ip, s.port, s.user, s.password = name, ip, port, user, password
	var err error
	s.client, err = xssh.NewClient(ip, port, user, password)
	return err
}

func (s SshConnect) runCmd(cmd string) (string, error) {
	if cmd == "" {
		return "", errors.New("no cmd run")
	}
	output, err := s.client.Output(cmd)
	if err != nil {
		return string(output), err
	}
	return string(output), nil
}

func connectAndRunCmd(ip string, times int, waitDuration int, cmd []string) ([]string, error) {
	var err error
	var connect *SshConnect
	outList := make([]string, 0, len(cmd))
	for i := 0; i <= times; i++ {
		connect, err = newConnect("", "root", ip, "", "password")
		if err == nil {
			break
		}
		time.Sleep(time.Second * time.Duration(waitDuration))
	}
	for _, d := range cmd {
		var out string
		out, err = connect.runCmd(d)
		outList = append(outList, out)
	}
	connect.close()
	return outList, nil
}

func connectAndUpload(ori, dst, ip string, times int, waitDuration int, ignore []string) error {
	var connect *SshConnect
	var err error
	for i := 0; i <= times; i++ {
		connect, err = newConnect("", "root", ip, "", "password")
		if err == nil {
			break
		}
		time.Sleep(time.Second * time.Duration(waitDuration))
	}
	err = connect.upload(ori, dst, ignore)
	connect.close()
	return err
}

func newConnect(name, user, ip, port, password string) (s *SshConnect, err error) {
	if ip == "" || user == "" {
		err = errors.New("ip or user is null")
		return
	}
	if name == "" {
		name = ip
	}
	if port == "" {
		port = "22"
	}
	s = &SshConnect{}
	err = s.init(name, user, ip, port, password)
	return
}
