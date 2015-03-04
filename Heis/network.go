package main

import (
"fmt"
"net"
"time"
"strings"
)

func broadcast_msg(){
	UDPAdr, err1 := net.ResolveUDPAddr("udp", "127.0.0.1:92929") //Gir hvilken port vi skal broadcaste fra. 
	if err1 != nil{
		fmt.Println(err1)
	}

	connUDP, err := net.DialUDP("udp", nil, UDPAdr)
	if err != nil{
		fmt.Println(err)
	}
	for i := 0; i < 10; i++ {
		n, err1 := SrvUDP.Write([]byte("Broadcast message!\n\x00"))
		if err1 != nil{
			fmt.Println("Error",err1,n)
		}
		time.Sleep(100*time.Millisecond)	
	}
	
}



func receive_msg(){
	buffer := make([]byte,1024)
	
	UDPAdr, err1 := net.ResolveUDPAddr("udp", ":92929")
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
			if !strings.Contains(Addr, ".150"){
				fmt.Printf("Rcv %d bytes: %s",n,buffer[0:n])
				}
	}
}
	
	

func main(){
	go read()
	time.Sleep(3000*time.Millisecond)	
	go write()
	time.Sleep(15000*time.Millisecond)
	fmt.Println("Done!")
	
}
