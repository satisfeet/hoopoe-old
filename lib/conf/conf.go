package conf

import (
    "os"
    "io/ioutil"
    "path/filepath"
    "encoding/json"
)

var conf *Conf

const (
  DEFAULT = "/etc/default.json"
  DEVELOPMENT = "/etc/development.json"
)

type Conf struct {
    Store   map[string]string   `json:"store"`
    Httpd   map[string]string   `json:"httpd"`
}

func Init() error {
    conf = &Conf{}

    if err := load(DEFAULT); err != nil {
        return err
    }
    if err := load(DEVELOPMENT); err != nil {
        return err
    }

    return nil
}

func Get(key string) map[string]string {
    switch (key) {
        case "store":
            return conf.Store
        case "httpd":
            return conf.Httpd
    }

    return make(map[string]string)
}

func load(path string) error {
    wd, err := os.Getwd()

    if err != nil {
        return err
    }

    file, err := ioutil.ReadFile(filepath.Join(wd, path))

    if err != nil {
        return err
    }

    return json.Unmarshal(file, conf)
}
