package main

import (
"fmt"
"net"
"time"
"os" //used in get_local_IPaddr()
"encoding/json"
"strings"
//"strings"
)

//Solveig IP: 129.241.187.141
const N_Elevators = 3
var active_elevators_in_system[N_Elevators]string



type Message struct {
		Sender_IP        string
		Message_type     string //ping, order_table, state_update
		Data             string //Annen type/container - eks bytes pakken?
	}
	
func spam_ImAlive(){ //Kjøres som en go routine fra main
	for{
		broadcast_msg("ping", "ImAlive")
		time.Sleep(100*time.Millisecond)
	}		
}

func watch_elevators_in_system(){ //Kjøres når man mottar en ping. Legg denne funksjonalitet til Manager. 
	//Trenger å vite hvilke heiser som er i systemet
	var 
	for{
		
	}
}


func broadcast_msg(msg_type string, data string){
	UDPAdr, err1 := net.ResolveUDPAddr("udp", "129.241.187.255:40101") //Gir hvilken port vi skal broadcaste fra/til? 
	if err1 != nil{
		fmt.Println(err1)
	}
	UDPconn, err := net.DialUDP("udp", nil, UDPAdr)
	if err != nil{
		fmt.Println(err)
	}
	
	
	//Egen funksjon pack_msg(data)? ---------------
	msg := Message{
		Sender_IP: get_local_IPadr(),
		Data: data,
	}
	//--------------------------------------------
	//msg := pack_msg(data)	
	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("error:", err)
	}
	UDPconn.Write(b)
	
}
	
func test_broadcaster(){
	for i := 0; i < 10; i++ {
		msg := []string{"Broadcast msg"}
		broadcast_msg("ping", strings.Join(msg, " "))
		time.Sleep(100*time.Millisecond)	
	}
}


func receive_msg(){//Kjøres som en go routine fra main? 
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
			//Pass whole struct msg to Manager. Manager wil do different things depending on msg type
			//Manager.handle_incomming_message(msg)
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

	// check the address type and if it is not a loopback then display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "Could not get local IP" //Noe bedre å returnere her???
}



//Ha IPadr i en string. Brukes til identifikasjon. Eks: if !strings.Contains(Addr, ".150"){
	
	

func main(){
	go receive_msg()
	time.Sleep(3000*time.Millisecond)	
	go test_broadcaster()
	time.Sleep(15000*time.Millisecond)
	fmt.Println("Done!")
	
}
