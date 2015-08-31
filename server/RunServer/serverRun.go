package serverInit

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

func WritePid(pidFile string) {
	ioutil.WriteFile(pidFile, []byte(strconv.Itoa(os.Getpid())), 0744)
}

func StopPid(pidFile string) {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		fmt.Printf("cannot read pid file: %s\n", pidFile)
		return
	}
	pid, _ := strconv.Atoi(string(data))
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("read pid process error")
		return
	}
	err = process.Signal(syscall.SIGUSR1)
	if err != nil {
		fmt.Println("send signal to process error")
		return
	}
	err = os.Remove(pidFile)
	if err == os.ErrNotExist {
		err = nil
	}
	if err != nil {
		fmt.Println("send signal to process successfully and remove pid error: ", err)
	}
	fmt.Println("send signal to process successfully")
}
