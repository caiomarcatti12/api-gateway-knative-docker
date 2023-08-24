package cors

type CORSConfig struct {
	AllowedOrigins   []string `yaml:"allowedOrigins"`
	AllowedMethods   []string `yaml:"allowedMethods"`
	AllowedHeaders   []string `yaml:"allowedHeaders"`
	AllowCredentials bool     `yaml:"allowCredentials"`
	ExposedHeaders   []string `yaml:"exposedHeaders"`
	MaxAge           int      `yaml:"maxAge"`
}
