package system

import "golang.org/x/sys/windows/registry"

// RegisterTable 用于windows注册表
func RegisterTable() error {
	key, _, err := registry.CreateKey(
		registry.LOCAL_MACHINE,
		`System\CurrentControlSet\Services\NPF`,
		registry.ALL_ACCESS)
	if err != nil {
		return err
	}

	_ = key.SetDWordValue(`Type`, uint32(1))
	_ = key.SetExpandStringValue(`ImagePath`, `system32\drivers\NPF.sys`)
	_ = key.SetStringValue(`DisplayName`, `WinPcap Packet Driver (NPF)`)
	_ = key.Close()
	return nil
}
