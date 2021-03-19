package os

import "github.com/r1cebank/cbox/app"

type ArchLinux struct {
}

func (_ ArchLinux) GetName() string {
    return "archlinux"
}

func (_ ArchLinux) RunCommand(command string) string {
    return ""
}

func (_ ArchLinux) CreateBackup() string {
    return ""
}

func (_ ArchLinux) FinalizeBackup(backupId string) {
} 

func (_ ArchLinux) GetConfigurations(app app.App) [][]byte {
    configs := [][]byte{}
    return configs
}


