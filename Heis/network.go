package main

import (
"fmt"
"net"
"time"
"os" //used in get_local_IPaddr()
"encoding/json"
"strings"
)


type Message struct {
		Sender_IP        string
		Message_type     string //ping, order_update elev_state_update
		Message_content  string //Annen type/container - eks bytes pakken?
	}

var incoming_message = make(chan Message)
	
func spam_ImAlive(){ //Kjøres som en go routine fra main
	for{
		broadcast_msg("ping", "ImAlive")
		time.Sleep(300*time.Millisecond)
	}		
}

func broadcast_msg(msg_type string, msg_content string){
	UDPAdr, err1 := net.ResolveUDPAddr("udp", "129.241.187.255:40101") 
	if err1 != nil{
		fmt.Println(err1)
	}
	UDPconn, err := net.DialUDP("udp", nil, UDPAdr)
	if err != nil{
		fmt.Println(err)
	}

	msg := Message{
		Sender_IP: get_local_IPadr(),
		Message_type: msg_type,
		Message_content: msg_content,
	}
	
	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("error:", err)
	}
	UDPconn.Write(b)
	
}
	
func test_broadcaster(){¨//Tas vekk når vi er ferdig med å teste modulen
	for i := 0; i < 10; i++ {
		msg := []string{"Broadcast msg"}
		broadcast_msg("ping", strings.Join(msg, " "))
		time.Sleep(100*time.Millisecond)	
	}
}


func receive_msg(){//Kjøres som en goroutine fra main 
	buffer := make([]byte,1024)
	UDPAdr, err1 := net.ResolveUDPAddr("udp", ":40101")
	if err1 != nil{
		fmt.Println(err1)
	}
	
	UDPConn, err2 := net.ListenUDP("udp", UDPAdr)
	if err2 != nil{
		fmt.Println(err2)
	}	
	for {
		n, Addr, err3 := UDPConn.ReadFromUDP(buffer)
		if err3 != nil{
			fmt.Println(err3)
		}
		var msg Message
		err4 := json.Unmarshal(buffer[0:n], &msg)
		if err4 != nil{
			fmt.Println(err4)
		}
		
		if msg.Sender_IP != get_local_IPadr(){
			fmt.Println(msg.Sender_IP)
			incoming_message <- msg //Sender hele structen
		
		}
		fmt.Println(Addr,n)
		//fmt.Printf("Rcv message, %d bytes: %s",n,buffer[0:n])
	}
}


func get_local_IPadr()string{
	addrs, err := net.InterfaceAddrs()
 	if err != nil {
		 fmt.Println(err)
		 os.Exit(1)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "Error: Could not get local IP" //Noe bedre å returnere her???
}

	

func main(){
	go receive_msg()
	time.Sleep(3000*time.Millisecond)	
	go test_broadcaster()
	go spam_ImAlive()
	time.Sleep(15000*time.Millisecond)
	fmt.Println("Done!")
}
