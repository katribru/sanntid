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

func broadcast_msg(data string){
	UDPAdr, err1 := net.ResolveUDPAddr("udp", "127.0.0.1:55001") //Gir hvilken port vi skal broadcaste fra. 
	if err1 != nil{
		fmt.Println(err1)
	}
	UDPconn, err := net.DialUDP("udp", nil, UDPAdr)
	if err != nil{
		fmt.Println(err)
	}
	
	
	//Egen funksjon pack_msg(data)? ---------------
	type Message struct {
		Sender_IP        string
		//Message_type   string
		Data             string //Annen type/container - eks bytes pakken?
	}
	
	//ip = get_local_IPadr()
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
		broadcast_msg(strings.Join(msg, " "))
		time.Sleep(100*time.Millisecond)	
	}

}



func receive_msg(){
	buffer := make([]byte,1024)
	
	UDPAdr, err1 := net.ResolveUDPAddr("udp", ":55001")
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
		//if !strings.Contains(Addr, ".150"){
		//	fmt.Printf("Rcv %d bytes: %s",n,buffer[0:n])
		//}
		fmt.Println(Addr,n)
		fmt.Printf("Rcv message, %d bytes: %s",n,buffer[0:n])
	}
}




func get_local_IPadr()string{
	addrs, err := net.InterfaceAddrs()
 	if err != nil {
		 fmt.Println(err)
		 os.Exit(1)
	}
	for _, address := range addrs {

	// check the address type and if it is not a loopback the display it
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
	time.Sleep(10000*time.Millisecond)
	fmt.Println("Done!")
	
}
