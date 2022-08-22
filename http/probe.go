package http

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "net/http/httptrace"
    "strings"
    "time"
)

func HttpProbe(context *gin.Context) {
    // 解析url参数，拼凑出需要探测的目标url
    host := context.Query("host")
    isHttps := context.Query("is_https")
    if host == "" {
        context.String(http.StatusBadRequest, "empty host")
    }
    schema := "http"
    if isHttps == "1" {
        schema = "https"
    }
    url := fmt.Sprintf("%v://%v", schema, host)

    // 创建一个request实例并给其配置上下文
    var (
        start, t1 time.Time
        dnsRes, targetIp string
    )
    req, _ := http.NewRequest("GET", url, nil)
    trace := &httptrace.ClientTrace{
        DNSStart: func(info httptrace.DNSStartInfo) {
            start = time.Now()
        },
        DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
            ips := make([]string, 0)
            for _, addr := range dnsInfo.Addrs {
                ips = append(ips, addr.IP.String())
            }
            dnsRes = strings.Join(ips, ", ")
        },
        ConnectStart: func(network, addr string) {
            t1 = time.Now()
            targetIp = addr
        },
    }
    req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

    // 创建一个client实例并传递连接超时时长等参数
    client := http.Client{
        Timeout: time.Duration(HttpProbeTimeout) * time.Second,
    }
    resp, err := client.Do(req)
    if start.IsZero() {
       start = t1
    }
    var msg string
    if err != nil {
        msg = fmt.Sprintf("http探测出错\n错误详情：%v\n探测目标：%v\n总耗时：%vms", err, url, timeConsume)
        context.String(http.StatusExpectationFailed, msg)
        return
    }
    defer resp.Body.Close()
    msg = fmt.Sprintf("探测目标：%v\nDNS解析结果：%v\n连接的IP：%v\n总耗时：%vms", url, dnsRes, targetIp, timeConsume)
    context.String(resp.StatusCode, msg)
}
