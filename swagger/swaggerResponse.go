package swagResponse

type SuccessResponse200 struct {
	Code 	int		`json:"code" example:"200"`
	Data	interface{}		`json:"data" `
	Msg		string		`json:"msg" example:"操作成功"`
}

type FailureResponse400 struct {
	Code 	int		`json:"code" example:"400"`
	Data	string		`json:"data" example:"null"`
	Msg		string		`json:"msg" example:"错误请求"`
}

type FailureResponse401 struct {
	Code 	int		`json:"code" example:"401"`
	Data	string		`json:"data" example:"null"`
	Msg		string		`json:"msg" example:"权限不足"`
}

type FailureResponse422 struct {
	Code 	int		`json:"code" example:"422"`
	Data	string		`json:"data" example:"null"`
	Msg		string		`json:"msg" example:"语义错误"`
}

type FailureResponse500 struct {
	Code 	int		`json:"code" example:"500"`
	Data	string		`json:"data" example:"null"`
	Msg		string		`json:"msg" example:"服务器内部错误"`
}

