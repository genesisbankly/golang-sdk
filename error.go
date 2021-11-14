package genesis

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Module  string `json:"module"`
}
