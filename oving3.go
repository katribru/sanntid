package main

import (
"fmt"
"net"
)

func main(){
	buffer := make([]byte,1024)
	
	UDPAdr, err1 := net.ResolveUDPAddr("udp", ":30000")
	if err1 != nil{
		fmt.Println(err1)
	}
	fmt.Println(*UDPAdr)
	
	UDPConn, err2 := net.ListenUDP("udp", UDPAdr)
	if err2 != nil{
		fmt.Println(err2)
	}
	fmt.Println(UDPConn)

	for true{
		n, Addr, err3 := UDPConn.ReadFromUDP(buffer)
		if err3 != nil{
			fmt.Println(err3)
		}
		fmt.Println(Addr)
		fmt.Printf("Rcv %d bytes: %s",n,buffer[0:n])
	}
}
