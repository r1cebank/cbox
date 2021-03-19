package os

import "github.com/r1cebank/cbox/app"

type Ubuntu struct {
}


func (_ Ubuntu) GetName() string {
    return "ubuntu"
}

func (_ Ubuntu) RunCommand(command string) string {
    return ""
}

func (_ Ubuntu) CreateBackup() string {
    return ""
}

func (_ Ubuntu) FinalizeBackup(backupId string) {
} 

func (_ Ubuntu) GetConfigurations(app app.App) [][]byte {
    configs := [][]byte{}
    return configs
}

