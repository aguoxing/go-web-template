package configs

type DBConfig struct {
	Platform string
	Host     string
	Port     int16
	Dbname   string
	Username string
	Password string
	Arg      string
}
