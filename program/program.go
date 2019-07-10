package program

import (
	"encoding/json"
	"log"

	"github.com/shiguanghuxian/go-commonly-template/internal/config"
)

// Program 程序实体
type Program struct {
	cfg *config.Config
}

// New 创建程序实例
func New() (*Program, error) {
	// 初始化配置文件
	cfgChan, err := config.NewConfig("")
	if err != nil {
		return nil, err
	}

	return &Program{
		cfg: <-cfgChan,
	}, nil
}

// Run 启动程序
func (p *Program) Run() {
	js, _ := json.Marshal(p.cfg)
	log.Println(string(js))
}

// Stop 程序结束要做的事
func (p *Program) Stop() {

}
