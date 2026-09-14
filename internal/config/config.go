package config

type Config struct {
	Coin       string
	Path       string
	Prefix     string
	StorageDir string
}

func Default() Config {
	return Config{
		Coin:       "BTC",
		Path:       "m/84'/0'/0'",
		Prefix:     "ln",
		StorageDir: ".wallets",
	}
}
