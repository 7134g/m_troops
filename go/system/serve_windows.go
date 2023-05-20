package system

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func CreateWindowsServe() {
	cmd := exec.Command("cmd.exe", "/k")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	input, _ := cmd.StdinPipe()
	_ = cmd.Start()
	defer func() { _ = input.Close() }()

	_, _ = fmt.Fprintln(input, `sc create npf binPath= "system32\drivers\NPF.sys" type= kernel start= auto DisplayName= "WinPcap Packet Driver (NPF)"`)
	_, _ = fmt.Fprintln(input, `sc start npf`)
	_, _ = fmt.Fprintln(input, `exit`)
	_ = cmd.Wait()
}
