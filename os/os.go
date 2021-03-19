package os

import (
    "fmt"
    "log"
    "runtime"

    "github.com/go-ini/ini"

    "github.com/r1cebank/cbox/app"
)

// The common interface for all distribution
// Backup function will utilize the common interfaces
// defined here
type OperatingSystem interface{
    GetName() string
    RunCommand(string) string
    CreateBackup() string
    FinalizeBackup(string)
    GetConfigurations(app.App) [][]byte
}

func ReadOSRelease(configfile string) map[string]string {
    cfg, err := ini.Load(configfile)
    if err != nil {
        log.Fatal("Fail to read file: ", err)
    }

    ConfigParams := make(map[string]string)
    ConfigParams["ID"] = cfg.Section("").Key("ID").String()

    return ConfigParams
}


func GetOperatingSystem() OperatingSystem {
    runtimeOS := runtime.GOOS

    fmt.Println("os: ", ReadOSRelease("/etc/os-release")["ID"])
    // The runtime operating system is later used to refine
    switch runtimeOS {
    case "windows":
        return Windows{} 
    case "darwin":
        return MacOS{}
    case "linux":
        osRelease := ReadOSRelease("/etc/os-release")["ID"]
        if osRelease  == "arch" {
            return ArchLinux{}
        } else if osRelease == "ubuntu" {
            return Ubuntu{}
        }
    default:
        return nil 
    }
    return nil
}

