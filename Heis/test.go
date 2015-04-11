package main

import (
//"../driver"
//"../manager"
//"time"
"fmt"
"encoding/json"
)


func main(){

	distributed_orders := map[string]map[string]map[string]int{
		"local_IP": map[string]map[string]int{
			"UP": map[string]int{
				"1": 0,
				"2": 0,
				"3": 0,
				"4": 0,
			},
			"DOWN": map[string]int{
				"1": 0,
				"2": 0,
				"3": 0,
				"4": 0,
			},
			"INTERNAL": map[string]int{
				"1": 0,
				"2": 0,
				"3": 0,
				"4": 0,
			},
		},
	}
	fmt.Println(distributed_orders)

	b, err1 := json.Marshal(distributed_orders)
	if err1 != nil{
			fmt.Println(err1)
		}
	fmt.Println(b)

	
	content := make(map[string]map[string]map[string]int) 

	err4 := json.Unmarshal(b, &content)
		if err4 != nil{
			fmt.Println(err4)
		}

	fmt.Println("content: ", content)
	
	type Message struct {
		Sender_IP        string
		Message_type     string //ping, internal_orders, shared_orders, elev_state, 
								//ack_shared_orders_recieved, request_all_info, all_info, 
								//Slå sammen internal_orders og elev_state? -> JA, de vil alltid sendes sammen uansett 
		Message_content  []byte //Annen type/container - eks bytes pakken?
	}

	msg := Message{"123.123", "order", b}

	fmt.Println("struct msg: ", msg)

/*

	distributed_orders := map[string]map[string]map[int]int{
		"local_IP": map[string]map[int]int{
			"UP": map[int]int{
				1: 0,
				2: 0,
				3: 0,
				4: 0,
			},
			"DOWN": map[int]int{
				1: 0,
				2: 0,
				3: 0,
				4: 0,
			},
			"INTERNAL": map[int]int{
				1: 0,
				2: 0,
				3: 0,
				4: 0,
			},
		},
	}

	
	for elevator_IP, order_map := range distributed_orders{
		for order_type, floors := range order_map{
			for floor, status := range floors{
				fmt.Println(elevator_IP, order_type, floor, status)
			}
		}	
	}


	shared_orders := map[string]map[int]int{
		"UP": map[int]int{
			1: 0,
			2: 0,
			3: 0,
			4: 0,
		},
		"DOWN": map[int]int{
			1: 0,
			2: 0,
			3: 0,
			4: 0,
		}
	}
	fmt.Println("shared_orders: ",shared_orders["123.123"]["DOWN"])
	

	elevator_IP := "123.222"
	distributed_orders[elevator_IP] = map[string]map[int]int{
		"UP": map[int]int{
			1: 0,
			2: 0,
			3: 0,
			4: 0,
		},
		"DOWN": map[int]int{
			1: 7,
			2: 0,
			3: 0,
			4: 0,
		},
		"INTERNAL": map[int]int{
			1: 0,
			2: 0,
			3: 0,
			4: 0,
		},
	}
	fmt.Println("all_orders: ", distributed_orders["123.222"]["DOWN"])

	//fmt.Println(all_orders["123.222"]["DOWN"])

/*
"Elevator"
var elevators_to _be_added []string

func add_elevator_to_system(string elevator_IP){ //Spawn en ny tråd for hver heis som oppdages. Terminer når heisen er lagt til
	//Spør gitt heis om status, interne ordre og felles ordre. 
	Broadcast_msg("request_all_info", elevator_IP)
	//Legger til IP i kø av nye heiser man venter svar fra
	elevators_to _be_added = []string
	//Vent til dette er mottatt
	while elevator_IP
	internalOrders_sharedOrders_status := <- network.response_all_info_request
	//Merge felles ordre med eksisterende felles ordre. Ikke slett noen ordre, bare legg til
	//Legg interne ordre inn i distributed_orders. Bare de interne ordrene er lagt til 
	distributed_orders[elevator_IP] = map[string]map[int]int{
		"UP": map[int]int{
			1: 0,
			2: 0,
			3: 0,
			4: 0,
		},
		"DOWN": map[int]int{
			1: 7,
			2: 0,
			3: 0,
			4: 0,
		},
		"INTERNAL": internal_orders //Der internal_orders er map[int]int som mottas på request
		},
	}
	//Legg til heis IP i elevators_in_system med timestamp

	elevators_in_system[elevator_IP] = time.Now().Format("20060102150405")
	Broadcast oppdatert tabell for å sikre synkronisering? -> Ja, etter en gitt 
}
*/

/*
"Main"
case msg := <-incoming_message:
	...
	}else if msg.Message_type == "all_info"{
		response_all_info_request <- msg.
func add_elevator_to_system(string elevator_IP){

}
*/

}

/* TO DO:
- Lag map og tabell of interne og felles ordre - tilpass meldingene som sendes til dette
- Lag funksjonalitet for å bekrefte at felles ordretabell er oppdatert - "Network"


*/ 