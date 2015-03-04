package main

import (
"fmt"
"net"
"time"
)


func write(){
	UDPAdr, err1 := net.ResolveUDPAddr("udp", "129.241.187.255:20002")
	if err1 != nil{
		fmt.Println(err1)
	}

	SrvUDP, err := net.DialUDP("udp", nil, UDPAdr)
	if err != nil{
		fmt.Println(err)
	}
	
	SrvUDP.Write([]byte("Hei, Server!\x00"))
}

func read(){
	buffer := make([]byte,1024)
	
	UDPAdr, err1 := net.ResolveUDPAddr("udp", ":20002")
	if err1 != nil{
		fmt.Println(err1)
	}
	
	UDPConn, err2 := net.ListenUDP("udp", UDPAdr)
	if err2 != nil{
		fmt.Println(err2)
	}

	for true{
		n, Addr, err3 := UDPConn.ReadFromUDP(buffer)
		if err3 != nil{
			fmt.Println(err3)
		}
		fmt.Println(Addr,n)
		fmt.Printf("Rcv %d bytes: %s",n,buffer[0:n])
	}
}
	
	

func main(){
	go read()
	go write()
	time.Sleep(10000*time.Millisecond)	
}
