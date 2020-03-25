package defs

type Err struct {
	Error   string `json:"error"`
	ErrCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpSC: 400,
		Error: Err{
			Error:   "Request body is not correct",
			ErrCode: "001",
		},
	}

	ErrorNotAuthUser = ErrorResponse{
		HttpSC: 401,
		Error: Err{
			Error:   "User authentication failed.",
			ErrCode: "002",
		},
	}
)
