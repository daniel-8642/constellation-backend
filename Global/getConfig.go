package Global

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Web   Web   `yaml:"Web"`
	Mysql Mysql `yaml:"Mysql"`
}

type Web struct {
	Port          string `yaml:"port"`
	StaticWeb     string `yaml:"static_web"`
	StaticBackend string `yaml:"static_backend"`
	WebUrl        string `yaml:"web_url"`
	BackendUrl    string `yaml:"backend_url"`
	Key           string `yaml:"key"`
}
type Mysql struct {
	DriverName string `yaml:"driverName"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Ip         string `yaml:"ip"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
}

func GetWeb() *Web {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	config := new(Config)
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &config.Web
}
func GetMysql() *Mysql {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	config := new(Config)
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &config.Mysql
}
