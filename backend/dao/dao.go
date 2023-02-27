package dao

import (
	"github.com/li-zeyuan/CSEIdiom/backend/config"
)

var D *Dao

type Dao struct {
	cfg  *config.Config
	User *User
}

func New(cfg *config.Config) {
	D = &Dao{
		cfg:  cfg,
		User: &User{},
	}
}

func Close() {

}
