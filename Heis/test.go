package main

import (
//"../driver"
//"../manager"
//"time"
"fmt"
)

func main(){
	//orders := make(map[string]map[string]int)

/*	type order_at_flors struct{
		1 int
		2 int
		3 int
		4 int
	}
*/
	x := map[int]int{
				1: 1,
				2: 0,
				3: 17,
				4: 0,
			}

	fmt.Println(x[3])

	all_orders := map[string]map[string]map[int]int{
		"123.123": map[string]map[int]int{
			"UP": map[int]int{
				1: 1,
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
		"123.124": map[string]map[int]int{
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
	
	for elevator_IP, order_map := range all_orders{
		for order_type, floors := range order_map{
			for floor, status := range floors{
				fmt.Println(elevator_IP, order_type, floor, status)
			}
		}	
	}

	fmt.Println(all_orders["123.123"]["DOWN"])



}

/* TO DO:
- Lag map og tabell of interne og felles ordre - tilpass meldingene som sendes til dette
- Lag funksjonalitet for å bekrefte at felles ordretabell er oppdatert - "Network"


*/ 