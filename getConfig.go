package Config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Web   Web   `yaml:"web"`
	Mysql Mysql `yaml:"mysql"`
}
type Web struct {
	Port string `yaml:"port"`
}
type Mysql struct {
	DriverName string `yaml:"driverName"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	ip         string `yaml:"ip"`
	port       string `yaml:"port"`
	database   string `yaml:"database"`
}

var C Conf

func (c *Conf) getConf() *Conf {
	yamlFile, err := ioutil.ReadFile("../static/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
