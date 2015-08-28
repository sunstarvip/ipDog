package main
import (
    "net"
    "strconv"
    "fmt"
    "strings"
)

const IP_PART_NUM = 4

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
            }
        }
        nextIpSlice[i] = strconv.Itoa(partNum)
    }

    nextIp = strings.Join(nextIpSlice, ".")

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
    // IP地址
    beginIp := "180.97.33.107"
    endIp := "180.97.33.108"
    ipSlice := getIpSlice(beginIp, endIp)
    // PORT
    portSlice := getPortSlice(75,85)

    fmt.Println("ipSlice is:", ipSlice)  // TODO
    fmt.Println("portSlice is:", portSlice)  // TODO
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