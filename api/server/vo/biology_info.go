package vo

import "time"

type BiologyInfo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Gender string `json:"gender"`
	Birthday time.Time `json:"birthday"`
	// InGroup bool `json:"in_group"`
	CreateTime time.Time `json:"create_time"`
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