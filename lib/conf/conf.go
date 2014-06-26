package conf

import (
    "io/ioutil"
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
    file, err := ioutil.ReadFile(path)

    if err != nil {
        return err
    }

    return json.Unmarshal(file, c)
}
