package providers

type Config struct {
	DBUser     string
	DBPass     string
	DBName     string
	ListenPort string
}

func NewConfig() *Config {
	dbUser := "root"
	dbPass := "123"
	dbName := "go_chat"
	port := ":9876"

	return &Config{
		DBUser:     dbUser,
		DBPass:     dbPass,
		DBName:     dbName,
		ListenPort: port,
	}
}
