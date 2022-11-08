package system

import (
	"bytes"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

// CheckPermission 检测是否为管理员权限目录下
func CheckPermission(p string) (bool, error) {
	// 检测是否为管理员权限目录下
	f, err := os.Stat(p)
	if err != nil {
		return false, err
	}

	// 判断是否拥有写权限
	modeFile := string(f.Mode().String()[2]) + string(f.Mode().String()[5]) + string(f.Mode().String()[8])
	if modeFile == "www" {
		return true, nil
	}
	return false, nil
}

// GetHardDiskSpace 获取该路径下磁盘容量
func GetHardDiskSpace(s string) float64 {
	var info *disk.UsageStat
	var err error
	if runtime.GOOS == "windows" {
		info, err = disk.Usage(s[:2]) // window
	} else {
		info, err = disk.Usage("/") // linux
	}
	if err != nil || info == nil {
		log.Println("get disk error, info is nil, ", err)
		// 磁盘空间获取失败直接默认剩余10g空间
		return 10240.0
	}

	free := float64(info.Free) / 1024 / 1024
	// mb
	return free
}

// CPUType 获取cpu类型
func CPUType() int {
	hInfo, _ := host.Info()
	if hInfo.KernelArch == "x86_64" {
		return 64
	} else {
		return 32
	}
}

func QuickLink() {

	// 修改成ansi编码方式写vbs文件
	//	body := `Set WshShell=CreateObject("WScript.Shell")
	//strDesKtop=WshShell.SpecialFolders("Desktop")
	//Set oShellLink=WshShell.CreateShortcut(strDesKtop&"\TangGo.lnk")
	//oShellLink.TargetPath="%s"
	//oShellLink.WorkingDirectory="%s"
	//oShellLink.WindowStyle=1
	//oShellLink.Description="TangGo"
	//oShellLink.IconLocation="%s"
	//oShellLink.Save`

	content := bytes.NewBuffer(nil)
	content.WriteString("Set WshShell=CreateObject(\"WScript.Shell\")\n")
	content.WriteString("strDesKtop=WshShell.SpecialFolders(\"Desktop\")\n")
	content.WriteString(fmt.Sprintf("Set oShellLink=WshShell.CreateShortcut(strDesKtop&\"\\%s.lnk\")\n", "MyApp"))
	content.WriteString(fmt.Sprintf("oShellLink.TargetPath=\"%s\"\n", "c://app/run.exe"))
	content.WriteString(fmt.Sprintf("oShellLink.WorkingDirectory=\"%s\"\n", "c://app"))
	content.WriteString(fmt.Sprintf("oShellLink.Description=\"%s\"\n", "MyApp"))
	content.WriteString(fmt.Sprintf("oShellLink.IconLocation=\"%s\"\n", "c://app/run.icon"))
	content.WriteString("oShellLink.Save")

	//body = fmt.Sprintf(body, TargetPath, WorkingDirectory, icon)
	body := content.String()
	path := "./temp_makeLnk.vbs"
	f, err := os.Create(path)
	if err != nil {
		log.Println("temp_makeLnk.vbs error:", err)
		return
	}

	ansi, err := simplifiedchinese.GBK.NewEncoder().String(body)
	if err != nil {
		log.Println("temp_makeLnk.vbs error:", err)
		return
	}
	_, _ = f.WriteString(ansi)

	err = f.Close()
	if err != nil {
		log.Println("temp_makeLnk.vbs error:", err)
		return
	}
	cmd := exec.Command("wscript.exe", "//e:vbscript", path)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	result, err := cmd.Output()
	if err != nil {
		log.Println("cmd error:", err, string(result))
		return
	}
	log.Println("生成快捷方式：", body, string(result))
	log.Println("成功构建快捷方式")
}
