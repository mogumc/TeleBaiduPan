// 配置文件读取
// @author MoGuQAQ
// @version 1.0.0

package config

import (
	"TeleBaidu/global"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type config struct {
	Logger logger `yaml:"logger"`
	User   user   `yaml:"user"`
	TGBot  tginfo `yaml:"tginfo"`
}

type tginfo struct {
	Bot_token   string `yaml:"bottoken"`
	User_id     int    `yaml:"userid"`
	Privacymode bool   `yaml:"privacy"`
}

type logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Showline     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}

type user struct {
	Bduss      string `yaml:"bduss"`
	Is_VIP     int    `yaml:"is_vip"`
	AccLink    string `yaml:"acclink"`
	User_Agent string `yaml:"useragent"`
}

var Config *config

func init() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		global.Log.Errorf("解析文件失败: %v", err)
	}
	yaml.Unmarshal(yamlFile, &Config)
}

func UpdateYaml(config *config) {
	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		global.Log.Fatalf("发生错误!更新配置失败!")
	}
	err = ioutil.WriteFile("./config.yaml", yamlData, 0644)
	if err != nil {
		global.Log.Errorf("写入文件失败: %v", err)
	}
}
