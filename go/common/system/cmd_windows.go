package system

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

// SetCmdTitle 设置cmd的title
func SetCmdTitle(title string) {
	kernel32, _ := syscall.LoadLibrary(`kernel32.dll`)
	sct, _ := syscall.GetProcAddress(kernel32, `SetConsoleTitleW`)
	syscall.Syscall(sct, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	syscall.FreeLibrary(kernel32)
}

// ExecuteMoreCmd 执行多条cmd
func ExecuteMoreCmd() {
	cmd := exec.Command("/bin/bash", "-c")
	cmd.Stdout = os.Stdout
	input, _ := cmd.StdinPipe()
	defer func() { _ = input.Close() }()
	_ = cmd.Start()
	_, _ = fmt.Fprintln(input, `echo "hello world"`)
	_ = cmd.Wait()
}
