package driver

import (
//"fmt"
)

const N_BUTTONS = 3
const N_FLOORS = 4

type Button_type int
const( 
	BUTTON_CALL_UP Button_type = iota
	BUTTON_CALL_DOWN
	BUTTON_COMMAND
	)

//var button_types = [...]string{ 
//	"BUTTON_CALL_UP", 
//	"BUTTON_CALL_DOWN", 
//	"BUTTON_COMMAND",
//	}

type Motor_direction int
const(
	DIRN_DOWN = -1 << iota
	DIRN_STOP = 0 << iota
	DIRN_UP = 1<< iota
	)

var lamp_channel_matrix = [N_FLOORS][N_BUTTONS]int{
    [N_BUTTONS]int{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
    [N_BUTTONS]int{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
    [N_BUTTONS]int{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
    [N_BUTTONS]int{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4}}

var button_channel_matrix = [N_FLOORS][N_BUTTONS]int{
    [N_BUTTONS]int{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
    [N_BUTTONS]int{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
    [N_BUTTONS]int{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
    [N_BUTTONS]int{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4}}


func Get_floor_sensor_signal() int{
	if (io_read_bit(SENSOR_FLOOR1) != 0){
		return 0
	}else if (io_read_bit(SENSOR_FLOOR2)!= 0){
		return 1
	}else if (io_read_bit(SENSOR_FLOOR3) != 0){
		return 2
	}else if (io_read_bit(SENSOR_FLOOR4) != 0){
		return 3
	}else{
		return -1
	}
}

func Set_floor_indicator(floor int) {
	//assert(floor >= 0)
	//assert(floor < N_FLOORS)
	
	//Binary encoding. One light must always be on.
	if (floor&0x02 == 0x02){
		io_set_bit(LIGHT_FLOOR_IND1)
	}else{
		io_clear_bit(LIGHT_FLOOR_IND1)
	}
	if (floor&0x01 == 0x01){
		io_set_bit(LIGHT_FLOOR_IND2)
	}else{
		io_clear_bit(LIGHT_FLOOR_IND2)
	}
}

func Get_button_signal(button Button_type, floor int) int {
	//assert(floor >= 0);
	//assert(floor < N_FLOORS);
	//assert(!(button == BUTTON_CALL_UP && floor == N_FLOORS - 1));
	//assert(!(button == BUTTON_CALL_DOWN && floor == 0));
	//assert(button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND);
	if (io_read_bit(button_channel_matrix[floor][button]) != 0){
		return 1
	} else{
		return 0
	}
}

func Set_button_lamp(button Button_type, floor int, value int) {
	//assert(floor >= 0);
	//assert(floor < N_FLOORS);
	//assert(!(button == BUTTON_CALL_UP && floor == N_FLOORS - 1));
	//assert(!(button == BUTTON_CALL_DOWN && floor == 0));
	//assert(button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND);
	if (value != 0){
		io_set_bit(lamp_channel_matrix[floor][button])
	}else{
		io_clear_bit(lamp_channel_matrix[floor][button])
	}
}

func Set_motor_direction(dirn Motor_direction) {
	if (dirn == 0){
		io_write_analog(MOTOR, 0)
	} else if (dirn > 0) {
		io_clear_bit(MOTORDIR)
		io_write_analog(MOTOR, 2800)
	} else if (dirn < 0) {
		io_set_bit(MOTORDIR)
		io_write_analog(MOTOR, 2800)
	}
}

func Set_door_open_lamp(value int){
	if (value != 0){
		io_set_bit(LIGHT_DOOR_OPEN)
	} else{
		io_clear_bit(LIGHT_DOOR_OPEN)
	}
}

func Get_obstruction_signal() int {
	return io_read_bit(OBSTRUCTION)
}

func Get_stop_signal() int {
	return io_read_bit(STOP)
}

func Set_stop_lamp(value int){
	if (value != 0){
		io_set_bit(LIGHT_STOP)
	} else{
		io_clear_bit(LIGHT_STOP)
	}
}

func Init() int {
	if(io_init() == 0){
		return 0
	}
	for i := 0; i < N_FLOORS; i++{
		if (i != 0){
			Set_button_lamp(BUTTON_CALL_DOWN, i, 0)
		}
		if(i != N_FLOORS - 1){
			Set_button_lamp(BUTTON_CALL_UP, i, 0)
		}

		Set_button_lamp(BUTTON_COMMAND, i, 0)
	}
	Set_stop_lamp(0);
    	Set_door_open_lamp(0);
	//Vil vi at heisen skal kjoere til naermeste etg hvis den er mellom to etg naar vi starter? 
	//Skal vi default sette lyset i 1.etg?
    	Set_floor_indicator(0);
	//Return success
	return 1
}

