package config

type AppConfig interface {
    GetServersConfig() ServersConfig
    GetFirebaseConfig() FirebaseConfig
}

type ServersConfig interface {
    GetRestConfig() RestConfig
    GetGRPCConfig() GRPCConfig
}

type RestConfig interface {
    GetEnabled() bool
    GetHost() string
    GetPort() int
}

type GRPCConfig interface {
    GetEnabled() bool
    GetHost() string
    GetPort() int
}

type FirebaseConfig interface {
    GetServiceAccountFile() string
}

// -- impl

type appConfig struct {
    ServersCfg serversConfig `yaml:"servers"`
    FirebaseCfg firebaseConfig `yaml:"firebase"`
}

func (c appConfig) GetServersConfig() ServersConfig {
    return c.ServersCfg
}

func (c appConfig) GetFirebaseConfig() FirebaseConfig {
    return c.FirebaseCfg
}

// -- serversConfig impl

type serversConfig struct {
    RestCfg restConfig `yaml:"rest"`
    GRPCCfg gRPCConfig `yaml:"gRPC"`
}

func (c serversConfig) GetRestConfig() RestConfig {
    return c.RestCfg
}

func (c serversConfig) GetGRPCConfig() GRPCConfig {
    return c.GRPCCfg
}

// -- restConfig impl

type restConfig struct {
    Enabled bool `yaml:"enabled"`
    Host string `yaml:"host"`
    Port int `yaml:"port"`
}

func (c restConfig) GetEnabled() bool {
    return c.Enabled
}

func (c restConfig) GetHost() string {
    return c.Host
}

func (c restConfig) GetPort() int {
    return c.Port
}

// -- gRPC impl

type gRPCConfig struct {
    Enabled bool `yaml:"enabled"`
    Host string `yaml:"host"`
    Port int `yaml:"port"`
}

func (c gRPCConfig) GetEnabled() bool {
    return c.Enabled
}

func (c gRPCConfig) GetHost() string {
    return c.Host
}

func (c gRPCConfig) GetPort() int {
    return c.Port
}

// -- firebase impl

type firebaseConfig struct {
    ServiceAccountFile string `yaml:"serviceAccountFile"`
}

func (f firebaseConfig) GetServiceAccountFile() string {
    return f.ServiceAccountFile
}