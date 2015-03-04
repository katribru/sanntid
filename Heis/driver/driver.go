package driver

const N_BUTTONS = 3

static const int lamp_channel_matrix[N_FLOORS][N_BUTTONS] = {
    {LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
    {LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
    {LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
    {LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
};


static const int button_channel_matrix[N_FLOORS][N_BUTTONS] = {
    {BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
    {BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
    {BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
    {BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
};

func Set_door_open_lamp(value bool){
	if value{
		io_set_bit(LIGHT_DOOR_OPEN)
	} else{
		io_clear_bit(LIGHT_DOOR_OPEN)
	}
}

func Set_stop_lamp(value bool){
	if value{
		io_set_bit(LIGHT_STOP)
	} else{
		io_clear_bit(LIGHT_STOP)
	}
}

func Init() int {
	if(!Io_init()){
		return 0
	}
	
	for i := 0, i < N_FLOORS; i++{
		if (i != 0){
			Set_button_lamp(BUTTON_CALL_DOWN, i, 0)
		}
		if(i != N_FLOORS - 1){
			Set_button_lamp(BUTTON_CALL_UP, i, 0)
		}

		Set_button_lamp(BUTTON_COMMAND, i, 0)

		Set_stop_lamp(False);
    		Set_door_open_lamp(False);
    		Set_floor_indicator(0);
	}
	//Return success
	return 1;
	
}
//Copied from this computer
//New change