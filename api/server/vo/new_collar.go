package vo

type NewCollar struct {
	Area 	string 		`json:"area"`
	Iccid	string  	`json:"iccid"`      
	Police 	int			`json:"police"`     
	AllStep	int			`json:"all_step"`
	LastTime	string	`json:"last_time"`
    Temperature		float32		`json:"temperature"`
    Station 	string	`json:"station"`
    IsOnline	int 		`json:"isOnline"`
    SignalStrength	int 	`json:"signal_strength"`
    Type 	string		`json:"type"`
    Voltage 	float32		`json:"voltage"`
}