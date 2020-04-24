package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

var sshdFile string

// 读取文件
func readConf() (string, error) {
	f, e := ioutil.ReadFile(sshdFile)
	if e != nil {
		return "", errors.New("read file error")
	}
	str := string(f)
	return str, nil
}

// 写入系统剪贴板
func writeClipboard() bool {
	var run = runtime.GOOS
	if run == "darwin" {
		cat := exec.Command("cat", sshdFile)
		pbocpy := exec.Command("pbcopy")
		pbocpy.Stdout = os.Stdout
		in, _ := pbocpy.StdinPipe()
		cat.Stdout = in
		pbocpy.Start()
		cat.Run()
		in.Close()
		pbocpy.Wait()
	} else if run == "windows" {
		// TODO
	}
	return false
}

func init() {
	f, err := user.Current()
	if err != nil {
		fmt.Println("read user home dir error")
		os.Exit(3)
	}
	sshdFile = filepath.Join(f.HomeDir, "./.ssh/id_rsa.pub")
}
