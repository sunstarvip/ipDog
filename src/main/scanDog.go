package main
import (
    "net"
    "strconv"
    "fmt"
)

/**
    获取PORT范围slice
 */
func getPortSlice(beginPort, endPort int) (portSlice []int) {
    for port := beginPort; port <= endPort; port++ {
        portSlice = append(portSlice, port)
    }

    return
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
    portSlice := getPortSlice(0, 65535)

    for index, port := range portSlice {
        // 返回测试过的IP地址与测试结果
        testedIp, ok := doScan(ip, port)
        if(ok) {
            fmt.Printf("第%d次测试，%s is open \n", index, testedIp)
        }else {
            fmt.Printf("第%d次测试，%s is close \n", index, testedIp)
        }
    }
}