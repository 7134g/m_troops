package system

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func QuickLink(absPath, workDir string) {
	cmd := exec.Command("cmd.exe", "/k")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	input, _ := cmd.StdinPipe()
	_ = cmd.Start()
	_, _ = fmt.Fprintln(input, "set Program="+absPath)
	_, _ = fmt.Fprintln(input, "set LnkName=TangGo")
	_, _ = fmt.Fprintln(input, "set WorkDir="+workDir)
	_, _ = fmt.Fprintln(input, "set Desc=TangGo")
	_, _ = fmt.Fprintln(input, "(echo Set WshShell=CreateObject(\"WScript.Shell\"^)")
	_, _ = fmt.Fprintln(input, "echo strDesKtop=WshShell.SpecialFolders(\"DesKtop\"^)")
	_, _ = fmt.Fprintln(input, "echo Set oShellLink=WshShell.CreateShortcut(strDesKtop^&\"\\%LnkName%.lnk\"^)")
	_, _ = fmt.Fprintln(input, "echo oShellLink.TargetPath=\"%Program%\"")
	_, _ = fmt.Fprintln(input, "echo oShellLink.WorkingDirectory=\"%WorkDir%\"")
	_, _ = fmt.Fprintln(input, "echo oShellLink.WindowStyle=1")
	_, _ = fmt.Fprintln(input, "echo oShellLink.Description=\"%Desc%\"")
	_, _ = fmt.Fprintln(input, "echo oShellLink.Save)>makelnk.vbs")
	_, _ = fmt.Fprintln(input, "makelnk.vbs")
	_, _ = fmt.Fprintln(input, "exit")
	_ = cmd.Wait()
}
