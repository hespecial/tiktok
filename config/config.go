package config

type App struct {
	Port string
	Mode string
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Charset  string
}
