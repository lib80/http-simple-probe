package main

import (
    "flag"
    "fmt"
    "http-simple-probe/config"
    "http-simple-probe/http"
)

var configFile string

func main() {
    //解析命令行参数，解析配置文件并将配置信息传递给gin
    flag.StringVar(&configFile, "c", "config/config.yaml", "config file path")
    flag.Parse()
    cfg, err := config.LoadFile(configFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    http.StartGin(cfg)
}
