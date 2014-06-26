package conf

import (
    "os"
    "io/ioutil"
    "path/filepath"
    "encoding/json"
)

type Conf struct {
    Name    string              `json:"name"`
    Store   map[string]string   `json:"store"`
    Httpd   map[string]string   `json:"httpd"`
}

func New() *Conf {
    return &Conf{}
}

func (c *Conf) LoadJSON(path string) error {
    wd, err := os.Getwd()

    if err != nil {
        return err
    }

    file, err := ioutil.ReadFile(filepath.Join(wd, path))

    if err != nil {
        return err
    }

    return json.Unmarshal(file, c)
}
