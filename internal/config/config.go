package config

import (
	"os"

	"github.com/naoina/toml"
	"github.com/shiguanghuxian/go-commonly-template/internal/common"
)

// Config 配置文件
type Config struct {
	Debug bool `toml:"debug"`
}

// NewConfig 初始化一个server配置文件对象
func NewConfig(path string) (cfgChan chan *Config, err error) {
	if path == "" {
		path = common.GetRootDir() + "config/cfg.toml"
	}
	cfgChan = make(chan *Config, 0)
	// 读取配置文件
	cfg, err := readConfFile(path)
	if err != nil {
		return
	}
	go watcher(cfgChan, path)
	go func() {
		cfgChan <- cfg
	}()
	return
}

// ReadConfFile 读取配置文件
func readConfFile(path string) (cfg *Config, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	cfg = new(Config)
	if err := toml.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}
	return
}
