package elevator

import (
"../driver"
"time"
"fmt"

)

type State int
const( 
	IDLE State = iota
	RUNNING
	DOOR_OPEN
	)

const N_ELEVATORS = 3
var dir driver.Motor_direction
var current_floor int
var current_state State
var elevators_in_system [N_ELEVATORS]string
var Door = make(chan int)

var my_orders = [driver.N_FLOORS][driver.N_BUTTONS]int{
[driver.N_BUTTONS]int{0, 0, 0},
[driver.N_BUTTONS]int{0, 0, 0},
[driver.N_BUTTONS]int{0, 0, 0},
[driver.N_BUTTONS]int{0, 0, 0}}

var all_orders = [driver.N_FLOORS][driver.N_BUTTONS]int{
[driver.N_BUTTONS]int{0, 0, 0},
[driver.N_BUTTONS]int{0, 0, 0},
[driver.N_BUTTONS]int{0, 0, 0},
[driver.N_BUTTONS]int{0, 0, 0}}



//func Redistribute_orders(){
//	for i := 0; i < driver.N_FLOORS; i++{
//		for j := 0; j < driver.N_BUTTONS; j++{
//			if All_orders[i][j] == 1{
//				var best_elevator int
//				cost := 1000
//				for k := 0, k < elevators_in_system; k++{
//					if Calculate_cost(k) < cost{
//						best_elevator := k
//					}
//				}
//			}
//		}
//	}
//}

func Init() {
	if (driver.Get_floor_sensor_signal() == -1){
		driver.Set_motor_direction(driver.DIRN_DOWN)
	}
	dir = driver.DIRN_DOWN
	for (driver.Get_floor_sensor_signal() == -1){
		//wait until floor reached
	}
	driver.Set_motor_direction(driver.DIRN_STOP)
	current_floor := driver.Get_floor_sensor_signal()
	driver.Set_floor_indicator(current_floor)
	fmt.Println("Floor ", current_floor)

	current_state = IDLE

}

func Event_new_order(order driver.Order){
	if empty_queue() == 1{
		fmt.Println("Empty queue")
		my_orders[order.Floor][order.Button] = 1
		fmt.Println("My orders:")
		for i := 0; i < 4; i++{
	 		fmt.Printf("%v\n", my_orders[i])
	 	}

		if(order.Button == 0){ //UP
	 		driver.Set_button_lamp(driver.BUTTON_CALL_UP, order.Floor, 1)
	 	} else if (order.Button == 1){ //DOWN
	 		driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, order.Floor, 1)
	 	} else if (order.Button == 2){ //INSIDE
	 		driver.Set_button_lamp(driver.BUTTON_COMMAND, order.Floor, 1)
	 	}

	 	dir = decide_direction()
		fmt.Println("Dir",dir)
		driver.Set_motor_direction(dir)
		current_state = RUNNING

		if order.Floor == current_floor{
	 		driver.Set_motor_direction(driver.DIRN_STOP)
			go open_door()
		}

	} else{
		my_orders[order.Floor][order.Button] = 1
		fmt.Println("My orders:")
		for i := 0; i < 4; i++{
	 		fmt.Printf("%v\n", my_orders[i])
	 	}
	 	
	 	if(order.Button == 0){ //UP
	 		fmt.Println("Elevator order up at floor ", order.Floor)
	 		driver.Set_button_lamp(driver.BUTTON_CALL_UP, order.Floor, 1)
	 		cost := calcualte_cost(order)
	 		fmt.Println("Cost: ", cost)
	 	} else if (order.Button == 1){ //DOWN
	 		fmt.Println("Elevator order down at floor ", order.Floor)
	 		driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, order.Floor, 1)
	 		cost := calcualte_cost(order)
	 		fmt.Println("Cost: ", cost)
	 	} else if (order.Button == 2){ //INSIDE
	 		fmt.Println("Order inside elevator at floor ", order.Floor + 1)
	 		driver.Set_button_lamp(driver.BUTTON_COMMAND, order.Floor, 1)
	 		cost := calcualte_cost(order)
	 		fmt.Println("Cost: ", cost)
	 	}
	 	if order.Floor == current_floor{
	 		driver.Set_motor_direction(driver.DIRN_STOP)
			go open_door()
		}
	}

 }

func Event_floor_reached(floor_reached int){
	fmt.Println("Reached floor ", floor_reached + 1,"\n")
	driver.Set_floor_indicator(floor_reached)
	current_floor = floor_reached
	if decide_direction() == driver.DIRN_STOP{
		driver.Set_motor_direction(driver.DIRN_STOP)
		if empty_queue() == 0{
			go open_door()
		}
		
	}
/*	if (my_orders[floor_reached][driver.BUTTON_COMMAND] == 1){
		driver.Set_motor_direction(driver.DIRN_STOP)
		go open_door()
		my_orders[floor_reached][driver.BUTTON_COMMAND] = 0
		driver.Set_button_lamp(driver.BUTTON_COMMAND, floor_reached, 0)
	}
	if (my_orders[floor_reached][driver.BUTTON_CALL_UP] == 1 && dir == driver.DIRN_UP){
		driver.Set_motor_direction(driver.DIRN_STOP)
		go open_door()
		my_orders[floor_reached][driver.BUTTON_CALL_UP] = 0
		driver.Set_button_lamp(driver.BUTTON_CALL_UP, floor_reached, 0)
	}
	if (my_orders[floor_reached][driver.BUTTON_CALL_DOWN] == 1 && dir == driver.DIRN_DOWN){
		driver.Set_motor_direction(driver.DIRN_STOP)
		go open_door()
		my_orders[floor_reached][driver.BUTTON_CALL_DOWN] = 0
		driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, floor_reached, 0)
	}*/
}

func Event_door_timed_out(){
	fmt.Println("Event_door_timed_out")
	if driver.Get_floor_sensor_signal() != -1{
		remove_orders()
	}
	if current_state == DOOR_OPEN{
		new_direction := decide_direction()
		driver.Set_motor_direction(new_direction)
		if new_direction != driver.DIRN_STOP{
			driver.Set_door_open_lamp(0)
			current_state = RUNNING
			dir = new_direction
		}
		if empty_queue() == 1{
			driver.Set_door_open_lamp(0)
			current_state = IDLE
		}
	}
}

func calcualte_cost(order driver.Order) int {
	var cost int
	if order.Floor == current_floor{
		cost = 0
	} else if order.Floor > current_floor{
		cost = order.Floor - current_floor
	} else if order.Floor < current_floor{
		cost = current_floor - order.Floor
	}
	return cost
}

func open_door(){
	driver.Set_door_open_lamp(1)
	current_state = DOOR_OPEN
	time.Sleep(3000*time.Millisecond)
	Door <- 1
}

func empty_queue() int{
	for i := 0; i < driver.N_FLOORS; i++{
		for j := 0; j < driver.N_BUTTONS; j++{
			if my_orders[i][j] == 1{
				return 0
			}
		}
	}
	return 1
}


func decide_direction() driver.Motor_direction {
	var new_dir driver.Motor_direction

	orders_above := 0
	for i := (current_floor + 1); i < driver.N_FLOORS; i++{
		if (my_orders[i][driver.BUTTON_CALL_UP] == 1 || my_orders[i][driver.BUTTON_CALL_DOWN] == 1 || my_orders[i][driver.BUTTON_COMMAND] == 1){
			orders_above = 1
		}
	}
	orders_below := 0
	for i := (current_floor - 1); i >= 0; i--{
		if (my_orders[i][driver.BUTTON_CALL_UP] == 1 || my_orders[i][driver.BUTTON_CALL_DOWN] == 1 || my_orders[i][driver.BUTTON_COMMAND] == 1){
			orders_below = 1
		}
	}

	if (empty_queue() == 1){
		new_dir = driver.DIRN_STOP
		return new_dir
	}

	if (driver.Get_floor_sensor_signal() != -1){

		if dir == driver.DIRN_UP{
			if (my_orders[current_floor][driver.BUTTON_CALL_UP] == 1 || my_orders[current_floor][driver.BUTTON_COMMAND] == 1 ){
				new_dir = driver.DIRN_STOP
			} else if orders_above == 1{
				new_dir = driver.DIRN_UP
			} else if (my_orders[current_floor][driver.BUTTON_CALL_DOWN] == 1){
				new_dir = driver.DIRN_STOP
			} else{
				new_dir = driver.DIRN_DOWN
			}

		} else if dir == driver.DIRN_DOWN{
			if (my_orders[current_floor][driver.BUTTON_CALL_DOWN] == 1 || my_orders[current_floor][driver.BUTTON_COMMAND] == 1 ){
				new_dir = driver.DIRN_STOP
			} else if orders_below == 1{
				new_dir = driver.DIRN_DOWN
			} else if (my_orders[current_floor][driver.BUTTON_CALL_UP] == 1){
				new_dir = driver.DIRN_STOP
			} else{
				new_dir = driver.DIRN_UP
			}

		} 
	} else{
		if dir == driver.DIRN_UP{
			if orders_above == 1{
				new_dir = driver.DIRN_UP
			} else{
				new_dir = driver.DIRN_DOWN
			}

		} else if dir == driver.DIRN_DOWN{
			if orders_below == 1{
				new_dir = driver.DIRN_DOWN
			} else{
				new_dir = driver.DIRN_UP
			}
		}
	}

	return new_dir
}

func remove_orders(){
	orders_above := 0
	for i := (current_floor + 1); i < driver.N_FLOORS; i++{ 
		if ((my_orders[i][driver.BUTTON_CALL_UP]) == 1 || (my_orders[i][driver.BUTTON_CALL_DOWN]) == 1 || (my_orders[i][driver.BUTTON_COMMAND]) == 1){			
			orders_above = 1
		}			
	}

	orders_below := 0
	for i := (current_floor - 1); i >= 0; i--{ 
		if ((my_orders[i][driver.BUTTON_CALL_UP]) == 1 || (my_orders[i][driver.BUTTON_CALL_DOWN]) == 1 || (my_orders[i][driver.BUTTON_COMMAND]) == 1){			
			orders_below = 1
		}
	}
	
	my_orders[current_floor][driver.BUTTON_COMMAND] = 0;
	driver.Set_button_lamp(driver.BUTTON_COMMAND, current_floor, 0);

	if dir == driver.DIRN_UP{
		if(current_floor != 3){
			
			if(orders_above == 1 && orders_below == 1){
				my_orders[current_floor][driver.BUTTON_CALL_UP] = 0;
				driver.Set_button_lamp(driver.BUTTON_CALL_UP, current_floor, 0);
			} else{
				my_orders[current_floor][driver.BUTTON_CALL_UP] = 0;
				driver.Set_button_lamp(driver.BUTTON_CALL_UP, current_floor, 0);
				my_orders[current_floor][driver.BUTTON_CALL_DOWN] = 0;
				driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, current_floor, 0);
			}
		} else{
			my_orders[current_floor][driver.BUTTON_CALL_DOWN] = 0;
			driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, current_floor, 0);
		}
	} else if dir == driver.DIRN_DOWN{
		if(current_floor != 0){
			if(orders_above == 1 && orders_below == 1){
				my_orders[current_floor][driver.BUTTON_CALL_DOWN] = 0;
				driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, current_floor, 0);
			} else{
				my_orders[current_floor][driver.BUTTON_CALL_UP] = 0;
				driver.Set_button_lamp(driver.BUTTON_CALL_UP, current_floor, 0);
				my_orders[current_floor][driver.BUTTON_CALL_DOWN] = 0;
				driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, current_floor, 0);	
			}
		} else{
			my_orders[current_floor][driver.BUTTON_CALL_UP] = 0;
			driver.Set_button_lamp(driver.BUTTON_CALL_UP, current_floor, 0);
		}
	}
	
}