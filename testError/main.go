package main

import (
	"fmt"
)

var StErrMap = map[int]string{
	1001: "Invalid StProtocolType",
	1002: "StProtocol encode data type error",
}

type StError struct {
	ErrCode         int
	ErrMsg          string
	ErrTrackBackMsg string
}

func (e *StError) Error() string {
	return fmt.Sprintf("StError[%v] %v%v", e.ErrCode, e.ErrMsg, e.ErrTrackBackMsg)
}

func NewStError(code int, tbErr ...error) *StError {
	tbMsg := ""
	for _, e := range tbErr {
		if e != nil {
			tbMsg += "\n" + e.Error()
		}
	}

	if errMsg := StErrMap[code]; errMsg == "" {
		return &StError{
			ErrCode: 0,
			ErrMsg: "unknown error",
			ErrTrackBackMsg: tbMsg,
		}
	}
	return &StError{
		ErrCode: code,
		ErrMsg: StErrMap[code],
		ErrTrackBackMsg: tbMsg,
	}
}

func test1() error {
	return NewStError(1001)
}

func main() {
	fmt.Println(NewStError(1001))

}
