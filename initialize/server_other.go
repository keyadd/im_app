package initialize

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"im_app/global"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type server interface {
	ListenAndServe() error
}

func InitServer(address string, router *gin.Engine) server {
	if global.GVA_CONFIG.System.Env != "debug" {
		initProcess()
	}
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func initProcess() {
	if syscall.Getppid() == 1 {
		if err := os.Chdir("./"); err != nil {
			panic(err)
		}
		syscall.Umask(0)
		return
	}
	fmt.Println("go daemon!!!")
	fp, err := os.OpenFile("daemon.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = fp.Close()
	}()
	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	cmd.Stdout = fp
	cmd.Stderr = fp
	cmd.Stdin = nil
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	_, _ = fp.WriteString(fmt.Sprintf(
		"[PID] %d Start At %s\n", cmd.Process.Pid, time.Now().Format("2006-01-02 15:04:05")))
	os.Exit(0)
}
