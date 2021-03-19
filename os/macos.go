package os

import "github.com/r1cebank/cbox/app"

type MacOS struct {
}

func (_ MacOS) GetName() string {
    return "macos"
}

func (_ MacOS) RunCommand(command string) string {
    return ""
}

func (_ MacOS) CreateBackup() string {
    return ""
}

func (_ MacOS) FinalizeBackup(backupId string) {
} 

func (_ MacOS) GetConfigurations(app app.App) [][]byte {
    configs := [][]byte{}
    return configs
}


