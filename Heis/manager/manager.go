package manager

import (
"../driver"
"fmt"
)

var Button_signal_down = make(chan int)
var Button_signal_up = make(chan int)
var Button_signal_inside = make(chan int)
var Floor_reached = make(chan int)
var Dir driver.Motor_direction

var My_orders = [driver.N_FLOORS][driver.N_BUTTONS]int{
    [driver.N_BUTTONS]int{0, 0, 0},
    [driver.N_BUTTONS]int{0, 0, 0},
    [driver.N_BUTTONS]int{0, 0, 0},
    [driver.N_BUTTONS]int{0, 0, 0}}

func Poll(){
	count:= 0
	for{
		for i := 0; i < driver.N_FLOORS; i++{
			if (i != 0 && driver.Get_button_signal(driver.BUTTON_CALL_DOWN, i) != 0){
				My_orders[i][driver.BUTTON_CALL_DOWN] = 1
				fmt.Printf("%v", My_orders)
				Button_signal_down <- i
			}

			if(i != driver.N_FLOORS - 1 && driver.Get_button_signal(driver.BUTTON_CALL_UP, i) != 0){
				My_orders[i][driver.BUTTON_CALL_UP] = 1
				fmt.Printf("%v", My_orders)
				Button_signal_up <- i
				
			}
			if(driver.Get_button_signal(driver.BUTTON_COMMAND, i) != 0){
				My_orders[i][driver.BUTTON_COMMAND] = 1
				fmt.Printf("%v", My_orders)
				Button_signal_inside <- i
				
			}

		}
		floor := driver.Get_floor_sensor_signal()
		if (floor != -1 && count == 0){
			count = count + 1
			Floor_reached <- floor
		} else if(floor == -1){
			count = 0
		} else{
			count = count + 1
		}
	}
}

func Cost(){
	
}
