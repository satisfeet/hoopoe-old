package conf

import (
    "os"
    "io/ioutil"
    "encoding/json"
    "path/filepath"
)

const (
    DEFAULT     = "/etc/default.json"
    DEVELOPMENT = "/etc/development.json"
)

type Config struct {
    Name    string          `json:"name"`
    Store   StoreConfig     `json:"store"`
    Httpd   HttpdConfig     `json:"httpd"`
}

type StoreConfig struct {
    Name    string          `json:"name"`
    Host    string          `json:"host"`
}

type HttpdConfig struct {
    Port    string          `json:"port"`
}

func New() (*Config, error) {
    c := &Config{}

    if err := parse(DEFAULT, c); err != nil {
        return c, err
    }

    switch (os.Getenv("GO_ENV")) {
    default:
        if err := parse(DEVELOPMENT, c); err != nil {
            return c, err
        }
    }

    return c, nil
}

func parse (p string, c *Config) (error) {
    d, err := os.Getwd()

    if err != nil {
        return err
    }

    b, err := ioutil.ReadFile(filepath.Join(d, p))

    if err != nil {
        return err
    }

    return json.Unmarshal(b, c)
}
