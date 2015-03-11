package manager

import (
"../driver"
//"fmt"
)

var My_orders = [driver.N_FLOORS][driver.N_BUTTONS]int{
    [driver.N_BUTTONS]int{0, 0, 0},
    [driver.N_BUTTONS]int{0, 0, 0},
    [driver.N_BUTTONS]int{0, 0, 0},
    [driver.N_BUTTONS]int{0, 0, 0}}

var All_orders = [driver.N_FLOORS][driver.N_BUTTONS]int{
    [driver.N_BUTTONS]int{0, 0, 0},
    [driver.N_BUTTONS]int{0, 0, 0},
    [driver.N_BUTTONS]int{0, 0, 0},
    [driver.N_BUTTONS]int{0, 0, 0}}


func 