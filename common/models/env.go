package models

type Configs struct {
	Application Application `yaml:"application"`
	Jwt         Jwt         `yaml:"jwt"`
	Database    Database    `yaml:"database"`
}

type Application struct {
	Mode string `yaml:"mode"`
	Host string `yaml:"host"`
	Name string `yaml:"name"`
	Port int64  `yaml:"port"`
}
type Jwt struct {
	Secret  string `yaml:"secret"`
	Timeout int64  `yaml:"timeout"`
}
type Database struct {
	Driver    string `yaml:"driver"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      int64  `yaml:"port"`
	DBName    string `yaml:"dbName"`
	ParseTime bool   `yaml:"ParseTime"`
	Timeout   string `yaml:"timeout"`
	Loc       string `yaml:"loc"`
}
