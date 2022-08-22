package http

import (
    "github.com/gin-gonic/gin"
    "http-simple-probe/config"
    "net/http"
)
var HttpProbeTimeout int

func StartGin(cfg *config.Config) {
    HttpProbeTimeout = cfg.HttpProbeTimeout
    r := gin.Default()
    r.GET("/hello", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{"message": "welcome to this web."})
    })
    r.GET("/probe/http", HttpProbe)
    r.Run(cfg.HttpListenAddr)
}