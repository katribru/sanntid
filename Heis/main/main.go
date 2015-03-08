package main

import (
"../driver"
"../manager"
"time"
"fmt"
)

func main(){
	driver.Init()
	fmt.Println("Main kjoerer")

	i := driver.Get_floor_sensor_signal()
	//driver.Set_floor_indicator(i)
	fmt.Println("ETG =", i)
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
	go manager.Poll()
	var dir driver.Motor_direction
	for{
		select{
		case down := <- manager.Button_signal_down:
			dir = driver.DIRN_DOWN
			driver.Set_motor_direction(driver.DIRN_DOWN)
			driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, down, 1)

		case up := <- manager.Button_signal_up:
			//fmt.Println("Elevator order up at floor ", up)
			dir = driver.DIRN_UP
			driver.Set_button_lamp(driver.BUTTON_CALL_UP, up, 1)
			driver.Set_motor_direction(driver.DIRN_UP)

		case inside := <- manager.Button_signal_inside:
			//fmt.Println("Elevator order inside, floor ", inside)
			driver.Set_button_lamp(driver.BUTTON_COMMAND, inside, 1)

		case floor := <- manager.Floor_reached:
			fmt.Println("Reached floor ", floor)
			driver.Set_floor_indicator(floor)
			driver.Set_motor_direction(driver.DIRN_STOP)
			manager.My_orders[floor][driver.BUTTON_COMMAND] = 0
			driver.Set_button_lamp(driver.BUTTON_COMMAND, floor, 0)
			
			if (dir == driver.DIRN_DOWN){
				manager.My_orders[floor][driver.BUTTON_CALL_DOWN] = 0
				driver.Set_button_lamp(driver.BUTTON_CALL_DOWN, floor, 0)
			} else{
				manager.My_orders[floor][driver.BUTTON_CALL_UP] = 0
				driver.Set_button_lamp(driver.BUTTON_CALL_UP, floor, 0)
			}
	
		}
	}
	


	
}
