package main

import (
"../driver"
"../manager"
"time"
"fmt"
)

func main(){
	driver.Init()
	fmt.Println("Main running...")

	i := driver.Get_floor_sensor_signal()
	//driver.Set_floor_indicator(i)
	fmt.Println("Floor =", i)
//	for{
//		j := driver.Get_floor_sensor_signal()
//		fmt.Println("ETG =", j)
//		if (j != i && j != -1){
//			driver.Set_floor_indicator(j)
//		}
//		i = j
//		time.Sleep(1000*time.Millisecond)
//	}
//	driver.Set_motor_direction(driver.DIRN_DOWN)
//	time.Sleep(1000*time.Millisecond)
//	driver.Set_motor_direction(driver.DIRN_STOP)
//	time.Sleep(1000*time.Millisecond)
//	driver.Set_motor_direction(driver.DIRN_UP)
//	time.Sleep(1000*time.Millisecond)
//	driver.Set_motor_direction(driver.DIRN_STOP)
//	for i := 0; i < driver.N_FLOORS; i++{
//		if (i != 0){
//			driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, i, 1)
//		}
//		if(i != driver.N_FLOORS - 1){
//			driver.Set_button_lamp(driver.BUTTON_CALL_UP, i, 1)
//		}
//
//		driver.Set_button_lamp(driver.BUTTON_COMMAND, i, 1)
//		time.Sleep(1000*time.Millisecond)
//	}
//	for{
//		if (driver.Get_stop_signal() != 0){
//			fmt.Println("STOP")
//			driver.Set_stop_lamp(1)
//		}
//		time.Sleep(1000*time.Millisecond)
//	}

	time.Sleep(1000*time.Millisecond)
	go driver.Poll()
	go network.spam_ImAlive()
	go network.receive_msg()

	var N_elevators_in_system int 
	


	for{
		select{
		case new_order := <- driver.New_order:
			fmt.Printf("%v", new_order)
		// case order_down := <- driver.Button_signal_down:
		// 	manager.My_orders[i][driver.BUTTON_CALL_DOWN] = 1
		// 	fmt.Printf("%v", manager.My_orders)

		// 	dir = driver.DIRN_DOWN
		// 	driver.Set_motor_direction(driver.DIRN_DOWN)
		// 	driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, order_down, 1)

		// case order_up := <- driver.Button_signal_up:
		// 	//fmt.Println("Elevator order up at floor ", up)
		// 	manager.My_orders[i][driver.BUTTON_CALL_UP] = 1
		// 	fmt.Printf("%v", manager.My_orders)
		// 	dir = driver.DIRN_UP
		// 	driver.Set_button_lamp(driver.BUTTON_CALL_UP, order_up, 1)
		// 	driver.Set_motor_direction(driver.DIRN_UP)

		// case order_inside := <- driver.Button_signal_inside:
		// 	//fmt.Println("Elevator order inside, floor ", inside)
		// 	manager.My_orders[i][driver.BUTTON_COMMAND] = 1
		// 	fmt.Printf("%v", manager.My_orders)
		// 	driver.Set_button_lamp(driver.BUTTON_COMMAND, order_inside, 1)

		case floor_reached := <- driver.Floor_reached:
			fmt.Println("Reached floor ", floor_reached)
			driver.Set_floor_indicator(floor_reached)
			driver.Set_motor_direction(driver.DIRN_STOP)
			manager.My_orders[floor][driver.BUTTON_COMMAND] = 0
			driver.Set_button_lamp(driver.BUTTON_COMMAND, floor_reached, 0)
			
			if (dir == driver.DIRN_DOWN){
				manager.My_orders[floor_reached][driver.BUTTON_CALL_DOWN] = 0
				driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, floor_reached, 0)
			} else{
				manager.My_orders[floor_reached][driver.BUTTON_CALL_UP] = 0
				driver.Set_button_lamp(driver.BUTTON_CALL_UP, floor_reached, 0)
			}

		case msg := <-incoming_message:
			elevators_in_system[ip] = time.Now()
			if msg.Message_type == "ping"{
				//manager.update_elevators_in_system(msg.Sender_IP)
			}else if msg.Message_type == "order_update"{
				//mager.update_orders(msg.Data)
			}else if msg.Message_type == "state_update"{
				//maager.update_all_states(msg.Data)
			}else {
				fmt.Println("Error: Invalid msg type")
			}
	
		}
	}
	


	
}
