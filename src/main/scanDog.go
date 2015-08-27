package main
import (
    "net"
    "strconv"
    "fmt"
)

func getPortSlice() (portSlice []int) {
    for port := 0; port <= 65535; port++ {
        portSlice = append(portSlice, port)
    }
}

/**
    扫描指定IP和PORT，查看是否开启
 */
func doScan(ip string, port int) (ipAddr string, ok bool) {
    // 生成完整IP地址
    ipAddr = ip + ":" + strconv.Itoa(port)
    _, err := net.Dial("tcp", ipAddr)
    if err == nil {
        ok = true
        return
    }

    ok = false
    return
}

func main()  {
    // IP地址
    ip := "180.97.33.107"
    // PORT
    port := 80
    var successIp, ok = doScan(ip, port)
    if(ok) {
        fmt.Printf("%s is open \n", successIp)
    }else {
        fmt.Printf("%s is close \n", successIp)
    }
}