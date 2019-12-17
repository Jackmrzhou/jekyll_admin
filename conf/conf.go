package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
)

type Config struct {
	Token string
	User struct{
		Username string
		Password string
	}
	JekyllRoot string `yaml:"jekyll_root"`
	Local bool
}

func InitConfig(dir string) (*Config, error) {
	conf := &Config{}
	var filepath = path.Join(dir, "_admin_config.yml")
	if data, err := ioutil.ReadFile(filepath); err == nil{
		if err = yaml.Unmarshal(data, conf); err == nil {
			return conf, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}