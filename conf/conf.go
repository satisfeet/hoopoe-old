package conf

import (
    "flag"
)

type Conf struct {
    Httpd map[string]string
    Store map[string]string
}

func New() *Conf {
    return &Conf{
        map[string]string{},
        map[string]string{},
    }
}

func (c *Conf) FromFlags() error {
    c.Httpd["addr"] = *flag.String("port", ":3001",
        "Port to listen for incoming HTTP requests.")

    c.Store["mongo"] = *flag.String("mongo", "mongodb://localhost/satisfeet",
        "URL to connect to mongodb server.")

    return nil
}
