package model

type RedisConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Password   string `yaml:"password"`
	DefaultDB  int    `yaml:"defaultDb"`
	MaxRetries int    `yaml:"max_retries"`
}
