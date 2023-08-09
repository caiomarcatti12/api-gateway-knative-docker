package config

type Route struct {
	Path       string `yaml:"path"`
	Service    string `yaml:"service"`
	TTL        int    `yaml:"ttl"`
	Port       string `yaml:"port"`
	HealthPath string `yaml:"healthPath"` // Caminho para healthcheck
	Retry      int    `yaml:"retry"`      // Número de tentativas para healthcheck
	RetryDelay int    `yaml:"retryDelay"` // Intervalo entre tentativas em segundos
}
