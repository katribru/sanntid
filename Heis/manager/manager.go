package manager

import (
"../driver"
"time"
//"fmt"

)

type State int
const( 
	IDLE State = iota
	RUNNING
	DOOR_OPEN
	)
const N_ELEVATORS = 3
var dir driver.Motor_direction
var previous_floor int
var current_floor int
var current_state State
var elevators_in_system [N_ELEVATORS]string 





func Calculate_cost(elevator int) int{

}

func Redistribute_orders(){
	for i := 0; i < driver.N_FLOORS; i++{
		for j := 0; j < driver.N_BUTTONS; j++{
			if All_orders[i][j] == 1{
				var best_elevator int
				cost := 1000
				for k := 0, k < elevators_in_system; k++{
					if Calculate_cost(k) < cost{
						best_elevator := k
					}
				}
			}
	}
	}
}

