package main
import (
    "net"
    "strconv"
    "fmt"
    "strings"
    "flag"
    "os"
)

const(
    IP_PART_NUM = 4
    IP_END = "255.255.255.255"
    PORT_BEGIN = 0
    PORT_END = 65535
)

/**
    生成IP扫描范围参数
 */
func getIpParms(ipRange string) (beginIp, endIp string) {
    ipSlice := strings.Split(ipRange, "-")

    // 生成起始IP地址
    beginIpEntity := net.ParseIP(ipSlice[0])
    if beginIpEntity == nil {
        return "", ""
    }else {
        beginIp = beginIpEntity.String()
    }
    // 生成结束IP地址
    ipLenNum := len(ipSlice)  // 传入的IP值个数

    if ipLenNum == 1 {
        endIp = beginIp
    }else if ipLenNum == 2 {
        if ipSlice[1] == "" {  // 当结束IP为空时
            endIp = IP_END
        }else {
            endIpEntity := net.ParseIP(ipSlice[1])
            if endIpEntity == nil {
                return "", ""
            }else {
                endIp = endIpEntity.String()
            }
        }
    }else {
        return "", ""
    }

    return
}

func getPortParms(portRange string) (beginPort, endPort int, ok bool) {
    portSlice := strings.Split(portRange, "-")

    // 生成起始PORT
    var err error
    beginPort, err = strconv.Atoi(portSlice[0])
    if err != nil {
        return 0, 0, false
    }else {
        // 格式化起始PORT取值，即为[0, 65535]
        switch {
        case beginPort < PORT_BEGIN:
            beginPort = PORT_BEGIN
        case beginPort > PORT_END:
            beginPort = PORT_END
        }
    }
    // 生成结束PORT
    portLenNum := len(portSlice)  // 传入的IP值个数
    if portLenNum == 1 {
        endPort = beginPort
    }else if portLenNum == 2 {
        if portSlice[1] == "" {
            endPort = PORT_END
        }else {
            endPort, err = strconv.Atoi(portSlice[1])
        }
    }else {
        return 0, 0, false
    }

    ok = true
    return
}

/**
    根据当前传入IP获取下一个IP地址
 */
func getNextIp(ip string) (nextIp string) {
    ipPartSlice := strings.Split(ip, ".")
    addFlag := false  // 进位标记
    var partNum int  // 段位数字
    nextIpSlice := make([]string, IP_PART_NUM)
    for i := (IP_PART_NUM - 1); i >= 0; i-- {
        partNum, _ = strconv.Atoi(ipPartSlice[i])
        // 第4段位需要增加1 或者 前一段位产生进位的时候
        if addFlag || i == (IP_PART_NUM - 1) {
            if partNum == 255 {
                // 当第一段位为255时为最后一个IP，即：255.255.255.255
                if i == 0 {
                    return ""
                }else {
                    partNum = 0
                    addFlag = true
                }
            }else {
                partNum++
                addFlag = false
            }
        }
        nextIpSlice[i] = strconv.Itoa(partNum)
    }

    nextIp = strings.Join(nextIpSlice, ".")
    fmt.Println("nextIp is:", nextIp)  // TODO

    return
}

func getIpSlice(beginIp, endIp string) (ipSlice []string) {
    nextIp := beginIp
    for nextIp != "" && nextIp != endIp {
        ipSlice = append(ipSlice, nextIp)

        nextIp = getNextIp(nextIp)
    }
    ipSlice = append(ipSlice, endIp)

    return
}

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
    // 获取相关运行参数
    var ipRange *string = flag.String("ip", "127.0.0.1", "Target IP address's range like 0.0.0.1[-255.255.255.255]")
    var portRange *string = flag.String("port", "0", "Target PORT's range like 0[-65535]")
    // 开始解析命令参数
    flag.Parse()

    // IP地址范围
    fmt.Println("ipRange is:", *ipRange)  // TODO
    beginIp, endIp := getIpParms(*ipRange)
    if beginIp == "" {
        fmt.Println("请输入符合规范的IP地址范围")
        os.Exit(1)
    }
    fmt.Println("beginIp is:", beginIp)  // TODO
    fmt.Println("endIp is:", endIp)  // TODO
    ipSlice := getIpSlice(beginIp, endIp)

    // 端口范围
    fmt.Println("portRange is:", *portRange)  // TODO
    beginPort, endPort, ok := getPortParms(*portRange)
    if !ok {
        fmt.Println("请输入符合规范的PORT：0-65535")
        os.Exit(2)
    }
    fmt.Println("beginPort is:", beginPort)  // TODO
    fmt.Println("endPort is:", endPort)  // TODO
    portSlice := getPortSlice(beginPort, endPort)

    fmt.Println("ipSlice is:", ipSlice)  // TODO
//    fmt.Println("ipSlice count is:", len(ipSlice))  // TODO
    fmt.Println("portSlice is:", portSlice)  // TODO
    fmt.Println("portSlice count is:", len(portSlice))  // TODO
    index := 1  // 测试次数
    for _, ip := range ipSlice {
        for _, port := range portSlice {
            // 返回测试过的IP地址与测试结果
            testedIp, ok := doScan(ip, port)
            if(ok) {
                fmt.Printf("第%d次测试，%s is open \n", index, testedIp)
            }else {
                fmt.Printf("第%d次测试，%s is close \n", index, testedIp)
            }
            index++
        }
    }
}