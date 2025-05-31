package config

type Config struct {
	ServerAddress string
	DatabasePath  string
	DatabaseUrl   string
}

func Load() *Config {
	return &Config{
		ServerAddress: ":8080",
		DatabasePath:  "./db/stoy-svaya.db",
		DatabaseUrl:   "sqlite://db/stoy-svaya.db",
	}
}
