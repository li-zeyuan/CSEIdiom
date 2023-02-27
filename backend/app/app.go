package app

import (
	"github.com/li-zeyuan/CSEIdiom/backend/config"
)

var (
	_app *App
)

type App struct {
	Conf *config.Config
}

func New(conf *config.Config) (*App, error) {
	_app = &App{
		Conf: conf,
	}

	return _app, nil
}

func GetApp() *App {
	return _app
}

func GetCfg() *config.Config {
	return _app.Conf
}

func (a *App) Close() {

}
