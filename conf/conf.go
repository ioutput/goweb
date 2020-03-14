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
	Port     int64
	Name     string
}
type Setting struct {
	Mysql *MysqlConf
}

var Settings *Setting

func Init() (err error) {
	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//var cp *Setting
	if _, err = toml.DecodeFile("./config.toml", &Settings); err != nil {
		log.Fatal(err)
	}

	return
}
