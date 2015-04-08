package main

import (
"fmt"
"net"
"time"
"os" //used in get_local_IPaddr()
"encoding/json"
"strings"
)

//Elevators_in_system := make(map[string]int) ADD TO MANAGER evt ordremodulen







//Flytt til manager!!

type Elevator struct{
	IP                string
	time_last_contact int 
}

elevators_in_system = make(map[string]int)

func check_connected_elevators()//Kjør fra main. Ligge i manager el kømodul da det er i kømodulen man trenger å vite hvilke heiser som er i systemet
	//Legg en set_timestamp() funksjon i kømodul som kan kjøres fra main case ny msg
	//Bruke len(elevators_in_system) som variabel på antall heiser som er conectet
	//n := len(m) //returns the number of items in a map
	for{
		time_now = time.Now() //Sjekk om det stemmer med formatet til time.Now()

		for elevator_ip, timestamp := range elevators_in_system {
			if (timestamp + 10) < time_now{
				elevator_disconected(elevator_ip)//Funsksjon i elevator. Den skal oppdatere map? Noe mer?

			}
		}	
	}
}

func update_elevators_in_system(string elevator_IP){
	// if elevator_IP not in map{
	//	add_elevator_to_system(elevator_IP)
	//}
	//Gå gjennom liste med 

	//Elevators in system - map med IP som key og siste timestamp som entry


}

func add_elevator_to_system(string elevator_IP)
	//Broadcast oppdatert tabell slik at den nye blir oppdatert
}



type Message struct {
		Sender_IP        string
		Message_type     string //ping, order_update elev_state_update
		Message_content  string //Annen type/container - eks bytes pakken?
	}

var incoming_message = make(chan Message)
	
func Spam_ImAlive(){ //Kjøres som en go routine fra main
	for{
		Broadcast_msg("ping", "ImAlive")
		time.Sleep(300*time.Millisecond)
	}		
}

func Broadcast_msg(msg_type string, msg_content string){
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
	
func test_broadcaster(){//Tas vekk når vi er ferdig med å teste modulen
	for i := 0; i < 10; i++ {
		msg := []string{"Broadcast msg"}
		Broadcast_msg("ping", strings.Join(msg, " "))
		time.Sleep(100*time.Millisecond)	
	}
}


func Receive_msg(){//Kjøres som en goroutine fra main 
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
			//set_timestamp()
		
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
	go Receive_msg()
	time.Sleep(3000*time.Millisecond)	
	go test_broadcaster()
	go Spam_ImAlive()
	time.Sleep(15000*time.Millisecond)
	fmt.Println("Done!")
}
