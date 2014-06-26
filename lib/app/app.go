package app

import (
    "os"
    "path/filepath"

    "github.com/satisfeet/hoopoe/lib/conf"
)

type App struct {
    Conf    *conf.Conf
}

func New() *App {
    c := conf.New()

    return &App{c}
}

func (a *App) Configure(filename string) error {
    wd, err := os.Getwd()

    if err != nil {
        return err
    }

    return a.Conf.LoadJSON(filepath.Join(wd, filename))
}
