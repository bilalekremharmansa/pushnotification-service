package config

import (
    "log"
    "sync"

    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type config struct {
    ServerConfig struct {
        Host string `yaml:"host"`
        Port string `yaml:"port"`
    } `yaml:"server"`

    FirebaseConfig struct {
        ServiceAccountFile string `yaml:"serviceAccountFile"`
    } `yaml:"firebase"`
}

var (
    once sync.Once
	cfg config
)

func InitConfig(filepath string) {
    once.Do(func() {
        // first unmarshall default config, then let config file override it
        err := yaml.Unmarshal(defaultConfig, &cfg)
        if err != nil {
            log.Fatalf("Parsing default config failed: [%v]", err)
        }

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

func GetConfig() config {
	return cfg
}

// --

var defaultConfig = []byte(
`
server:
   host: ""
   port: 8080
firebase:
   serviceAccountFile: "/tmp/serviceAccount.json"
`)