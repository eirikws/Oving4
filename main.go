package main

import(
    "fmt"
    "net"
    "runtime"
    "communication"
)

const BCAST_IP=GetBIP(MY_IP)
const TARGET_PORT = "20011"
const LISTEN_PORT = "30022"

func main(){
    GOMAXPROCS(NumCPU())
    msg:=MakeMessage(BCAST_IP,"muhhahahaha","testtesttesttesttesttesttesttesttesttesttest")
    go communication.SendMsgTo(BCAST_IP,TARGET_PORT,msg)
    go communication.ListenerCon(BCAST_IP,LISTEN_PORT)
}

