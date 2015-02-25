package main

import (
"../driver"
"time"
"fmt"
)

func main(){
	driver.Io_init()
	driver.Set_door_open_lamp(true)
	driver.Set_stop_lamp(true)
	fmt.Println("Main kjoerer")
	time.Sleep(3000*time.Millisecond)
	driver.Set_door_open_lamp(false)
	driver.Set_stop_lamp(false)
}
