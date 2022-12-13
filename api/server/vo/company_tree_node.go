package vo

type CompanyTreeNode struct {
	Id        uint              `json:"id"`
	Name      string            `json:"name"`
	Location  string   			`json:"location"`
	Owner	  uint				`json:"owner"`
	Ancestors string            `json:"ancestors"`
	Children  []CompanyTreeNode `json:"children"`
}
