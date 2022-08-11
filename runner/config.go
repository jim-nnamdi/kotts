package runner

type Config struct {
	Dbdriver   string `mapstructure:"DB_DRIVER"`
	Dbhost     string `mapstructure:"DB_HOST"`
	ServerAddr string `mapstructure:"SERVER_ADDR"`
	DbSource   string `mapstructure:"DB_SOURCE"`
}
