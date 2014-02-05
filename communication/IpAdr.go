package main
import(
    "net"
    "fmt"
    "strings"
)

func GetMyIP() string{
    allIPs,err:=net.InterfaceAddrs()
    if err!=nil{
        fmt.Println("IP receiving errors!!!!!!!!\n")
        return ""
    }
    return strings.Split(allIPs[1].String(),"/")[0]
}

func GetBIP(MyIP string) string{
    IP:=strings.Split(MyIP,".")
    return IP[0]+"."+IP[1]+"."+IP[2]+".255"
}
