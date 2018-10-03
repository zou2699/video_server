// error handling

package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorReponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFaild = ErrorReponse{HttpSC: 400, Error: Err{Error: "request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser = ErrorReponse{HttpSC:401,Error:Err{Error:"User authentication failed.",ErrorCode:"002"}}
)
