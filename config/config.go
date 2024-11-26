package config

type Config struct {
	App   App
	Log   Log
	Mysql Mysql
}

type App struct {
	Port string
	Mode string
}

type Log struct {
	Path  string
	Name  string
	Level string
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Charset  string
}
