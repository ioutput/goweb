package conf

import (
	"log"

	"github.com/BurntSushi/toml"
	//"path/filepath"
)

type MysqlConf struct {
	Username string
	Password string
	Host     string
	Port     uint32
	Name     string
}
type Setting struct {
	Mysql *MysqlConf
}

var Settings *Setting = &Setting{Mysql: &MysqlConf{Username: "root", Password: "wiu5201314", Host: "localhost", Port: 3306, Name: "test"}}

func Init() (err error) {
	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//Settings = new(Setting)
	if _, err = toml.DecodeFile("/etc/goweb/config.toml", &Settings); err != nil {
		log.Fatal(err)
	}

	return
}
