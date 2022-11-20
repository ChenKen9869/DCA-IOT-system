package vo

import "time"

type BiologyInfo struct {
	Id   uint   `json:"biology_id"`
	Name string `json:"biology_name"`
	Type string `json:"biology_type"`
	Gender string `json:"gender"`
	Birthday time.Time `json:"birthday"`
	CreateTime time.Time `json:"create_date"`
	FarmhouseId uint `json:"farmhouse_id"`
}

type EpidemicPreventionRecord struct {
	VaccineDescription      string		`json:"vaccine_description"`
	InoculationTime 		time.Time	`json:"inoculation_time"`
}

type MedicalRecord struct {
	Disease 		string		`json:"disease"`
	IllnessTime 	time.Time	`json:"illness_time"`
	TreatmentPlan	string		`json:"treatment_plan"`
}

type OperationRecord struct {
	Doctor				string		`json:"doctor"`
	OperationTime		time.Time	`json:"operation_time"`
	ProcessDescription	string		`json:"process_description"`
	Result				string		`json:"result"`
}

type AuthBology struct {
	BiologyId	uint	`json:"biology_id"`
	BiologyName	string	`json:"biology_name"`
	BiologyType string	`json:"biology_type"`
	Gender	string	`json:"gender"`
	FarmhouseId	uint	`json:"farmhouse_id"`
	Birthday	time.Time	`json:"birthday"`
	CreateDate	time.Time	`json:"create_date"`
}

