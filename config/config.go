package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type Config struct {
    HttpListenAddr string `yaml:"http_listen_addr"`
    HttpProbeTimeout int `yaml:"http_probe_timeout"`
}

func LoadFile(filename string) (*Config, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    cfg := &Config{}
    err2 := yaml.Unmarshal(data, cfg)
    if err2 != nil {
        return nil, err2
    }
    return cfg, nil
}
