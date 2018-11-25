package config

// Config contains all the variables needed
type Config struct {
	CacheExpiration   int    `env:"CACHE_EXPIRATION" envDefault:"30"`
	CacheExpiredPurge int    `env:"CACHE_EXPIRED_PURGE" envDefault:"60"`
	StorageBackend    string `env:"STORAGE_BACKEND" envDefault:"memory"`
	StoragePath       string `env:"STORAGE_PATH" envDefault:"/pmt-data"`
	APIKey            string `env:"API_KEY" envDefault:""`
	Environment       string `env:"ENVIRONMENT" envDefault:"develop"`
	Port              int    `env:"PORT" envDefault:"8080"`
}
