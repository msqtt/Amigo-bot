package bot

import (
	"io/ioutil"

	"github.com/mosqu1t0/Amigo-bot/utils/logcat"
	"gopkg.in/yaml.v3"
)

type botYaml struct {
	Bot botConfig `yaml:"bot"`
}

type botConfig struct {
	Debug bool `yaml:"debug"`
	Ws    struct {
		Path string `yaml:"path"`
		Addr string `yaml:"addr"`
	} `yaml:"ws"`
	Root []int64 `yaml:"root"`
}

var (
	conf             botYaml
	DefaultBotConfig *botConfig
)

func init() {
	byte, err := ioutil.ReadFile("./config/application.yaml")
	if err != nil {
		logcat.ErrorEnd("Read bot conf err: ", err)
	}
	yaml.Unmarshal(byte, &conf)
	DefaultBotConfig = &conf.Bot
}
