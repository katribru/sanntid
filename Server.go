package main

import (
"fmt"
"net"
"time"
//"strings"
"encoding/json"
)

type UDP_message struct {
	Addr string
	Msg string
	Length int
}


func write(){
	UDPAdr, err1 := net.ResolveUDPAddr("udp", "129.241.187.255:50002")
	if err1 != nil{
		fmt.Println(err1)
	}

	SrvUDP, err := net.DialUDP("udp", nil, UDPAdr)
	if err != nil{
		fmt.Println(err)
	}
	
	var s string = "Hei client, fra server!\x00"
	
	m := UDP_message{"129.241.187.149",s,len(s)}
	//fmt.Printf("%+v\n",m)
	
	b,err2 := json.Marshal(m)
	if err2 != nil{
		fmt.Println(err2)
	}
	fmt.Printf("Sending: %s)",b[0:50])
	
	for i := 0; i < 10; i++ {
		n, err1 := SrvUDP.Write(b)
		if err1 != nil{
			fmt.Println("Error",err1,n)
		}
		time.Sleep(100*time.Millisecond)
	}
	
	
}

func read(){
	buffer := make([]byte,1024)
	
	UDPAdr, err1 := net.ResolveUDPAddr("udp", ":50002")
	if err1 != nil{
		fmt.Println(err1)
	}
	
	UDPConn, err2 := net.ListenUDP("udp", UDPAdr)
	if err2 != nil{
		fmt.Println(err2)
	}
	
	var m UDP_message
	
	for true{
		n, Addr, err3 := UDPConn.ReadFromUDP(buffer)
		if err3 != nil{
			fmt.Println(err3)
		}
		err4 := json.Unmarshal(buffer,&m)
		if err4 != nil{
			fmt.Println(err4)
		}
		fmt.Printf("%+v\n",m)
		fmt.Println(Addr,n)
		fmt.Printf("Rcv %d bytes: %s",n,buffer[0:n])
	}
}
	
	

func main(){
	go read()
	time.Sleep(3000*time.Millisecond)	
	go write()
	time.Sleep(15000*time.Millisecond)
	fmt.Println("Done!")
	
}
