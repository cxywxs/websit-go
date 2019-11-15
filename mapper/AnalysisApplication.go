package mapper

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Databases struct {
	Server   string `toml:"server"`
	Port     string `toml:"port"`
	Root     string `toml:"root"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

type Application struct {
	Databases *Databases
}

var application *Application = new(Application)

func AnalysisApplication() *Application {
	_, err := toml.DecodeFile("config/application.toml", application)
	if err != nil {
		fmt.Println(err)
	}
	return application
}
