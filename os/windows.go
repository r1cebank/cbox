package os

import "github.com/r1cebank/cbox/app"

type Windows struct {
}

func (_ Windows) GetName() string {
    return "windows"
}

func (_ Windows) RunCommand(command string) string {
    return ""
}

func (_ Windows) CreateBackup() string {
    return ""
}

func (_ Windows) FinalizeBackup(backupId string) {
} 

func (_ Windows) GetConfigurations(app app.App) [][]byte {
    configs := [][]byte{}
    return configs
}
