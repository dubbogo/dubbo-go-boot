package model

type ApplicationConfig map[string]interface{}

//type ApplicationConfig struct {
//	Server   *ServerConfig `yaml:"server"`
//	Redis    *RedisConfig  `yaml:"redis"`
//	Database *Database     `yaml:"database"`
//}

//type ServerConfig struct {
//	Host string `yaml:"host"`
//	Port string `yaml:"port"`
//}

//type RedisConfig struct {
//	Host       string `yaml:"host"`
//	Port       string `yaml:"port"`
//	Password   string `yaml:"password"`
//	DefaultDB  int    `yaml:"defaultDb"`
//	MaxRetries int    `yaml:"max_retries"`
//}

//type Database struct {
//	Dialect  string `yaml:"dialect"`
//	Host     string `yaml:"host"`
//	Port     string `yaml:"port"`
//	Database string `yaml:"database"`
//	Username string `yaml:"username"`
//	Password string `yaml:"password"`
//}
