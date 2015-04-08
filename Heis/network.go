package main

import (
"fmt"
"net"
"time"
"os" //used in get_local_IPaddr()
"encoding/json"
"strings"
)

/*
//Flytt til manager!!

type Elevator struct{
	IP                string
	time_last_contact int 
}

elevators_in_system := make(map[string]int)

func check_connected_elevators()//Kjør fra main. 
	for{
		time_now := time.Now().Format("20060102150405") //Format YYYYMMDDhhmmss

		for elevator_ip, timestamp := range elevators_in_system {
			if (timestamp + 5) < time_now{ //Kan ikke vente mer enn 5 sekunder
				delete(elevators_in_system, elevator_ip) // delete_elevator_from_system(elevator_ip)//Funsksjon i elevator. Den skal oppdatere map? Noe mer?

			}
		}
	}
}

func set_timestamp(string elevator_IP){ //Kjøres i case ny msg
	if val, ok := elevators_in_system[elevator_IP] {
		add_elevator_to_system(elevator_IP)
	}
	else {
		elevators_in_system[elevator_IP] = time.Now().Format("20060102150405")
	}
}

func add_elevator_to_system(string elevator_IP){}
	elevators_in_system[elevator_IP] = time.Now().Format("20060102150405")
	//Broadcast oppdatert tabell slik at den nye blir oppdatert??
}

func delete_elevator_from_system(elevator_ip){
	delete(elevators_in_system, elevator_ip)
}

*/

type Message struct {
		Sender_IP        string
		Message_type     string //ping, order_update elev_state_update
		Message_content  string //Annen type/container - eks bytes pakken?
	}

var Incoming_message = make(chan Message)
	
func Spam_ImAlive(){ //Kjøres som en go routine fra main
	for{
		Broadcast_msg("ping", "ImAlive")
		time.Sleep(300*time.Millisecond)
	}		
}

func Broadcast_msg(msg_type string, msg_content string){
	UDPAdr, _ := net.ResolveUDPAddr("udp", "129.241.187.255:40101") 
	
	UDPconn, _ := net.DialUDP("udp", nil, UDPAdr)

	msg := Message{
		Sender_IP: get_local_IPadr(),
		Message_type: msg_type,
		Message_content: msg_content,
	}
	
	b, _ := json.Marshal(msg)

	UDPconn.Write(b)
	
}
	
func test_broadcaster(){//Tas vekk når vi er ferdig med å teste modulen
	for i := 0; i < 10; i++ {
		msg := []string{"Broadcast msg"}
		Broadcast_msg("ping", strings.Join(msg, " "))
		time.Sleep(100*time.Millisecond)	
	}
}

func test_listener(){//Tas vekk når vi er ferdig med å teste modulen
	for{
		select{
		case new_msg := <- Incoming_message:
			fmt.Println("Case, new_msg: ", new_msg)
		}
	}
}


func Receive_msg(){//Kjøres som en goroutine fra main 
	buffer := make([]byte,1024)
	UDPAdr, _ := net.ResolveUDPAddr("udp", ":40101")
	
	
	UDPConn, _ := net.ListenUDP("udp", UDPAdr)
	
	for {
		n, _, _ := UDPConn.ReadFromUDP(buffer)
		
		var msg Message
		err4 := json.Unmarshal(buffer[0:n], &msg)
		if err4 != nil{
			fmt.Println(err4)
		}
		
		if msg.Sender_IP != get_local_IPadr(){
			fmt.Println("Mottar fra annen IP: ",msg.Sender_IP)
			Incoming_message <- msg //Sender hele structen
		}

		fmt.Println("n: ", n)
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
	
	go Receive_msg()
	time.Sleep(3000*time.Millisecond)
	go test_listener()	
	go Spam_ImAlive()
	go test_broadcaster()
	time.Sleep(15000*time.Millisecond)
	fmt.Println("Done!")
}
