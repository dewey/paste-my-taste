package config

// Config contains all the variables needed
type Config struct {
	APIKey      string `env:"API_KEY" envDefault:""`
	Environment string `env:"ENVIRONMENT" envDefault:"develop"`
	Port        int    `env:"PORT" envDefault:"8080"`
}
