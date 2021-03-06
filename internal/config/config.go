package config

import (
    "os"
    "log"
    "sync"
    "bytes"
    "runtime"

    "io/ioutil"
    "text/template"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

var (
    once sync.Once
	cfg appConfig
)

func InitDefaultConfig() {
    once.Do(func() {
        // first unmarshall default config, then let config file override it
        initDefaultConfig()
	})
}

func InitConfig(filepath string) {
    once.Do(func() {
        // first unmarshall default config, then let config file override it
        initDefaultConfig()

        log.Printf("Initializing config")
        configAsByte, err := ioutil.ReadFile(filepath)
        if err != nil {
            log.Fatalf("Reading config file failed: [%v]", err)
        }

        err = yaml.Unmarshal(configAsByte, &cfg)
        if err != nil {
            log.Fatalf("Parsing config file failed: [%v]", err)
        }
	})
}

func GetAppConfig() AppConfig {
	return cfg
}

// --

func initDefaultConfig() {
    err := yaml.Unmarshal([]byte(parseConfigTemplate()), &cfg)
    if err != nil {
        log.Fatalf("Parsing default config failed: [%v]", err)
    }

    log.Println(cfg)
}

func parseConfigTemplate() string {
	params := make(map[string]string)
	params["serviceAccountFile"] = GetDefaultServiceAccountFilePath()

	temp, err := template.New("config-template").Parse(configTemplate)
    if err != nil {
        log.Fatal("Error occurred while creating config template")
    }

    var tpl bytes.Buffer
    err = temp.Execute(&tpl, params)
    if err != nil {
        log.Fatal("Error occurred while parsing config template")
    }
    return tpl.String()
}

func GetDefaultConfigBaseDirPath() string {
    // windows
    if runtime.GOOS == "windows" {
        appData := os.Getenv("APPDATA")
        return filepath.Join(appData, "pns")
    } else { // unix
        return "/etc/pns.d"
    }
}

func GetDefaultConfigPath() string {
    return filepath.Join(GetDefaultConfigBaseDirPath(), "config.yaml")
}

func GetDefaultServiceAccountFilePath() string {
    return filepath.Join(GetDefaultConfigBaseDirPath(), "serviceAccount.json")
}

// --

var configTemplate =
`
servers:
    rest:
        enabled: true
        host: ""
        port: 8080
    gRPC:
        enabled: true
        host: ""
        port: 18080
firebase:
    serviceAccountFile: "{{.serviceAccountFile}}"
`