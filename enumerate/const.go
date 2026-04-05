package main

import "fmt"

const (
	SuccessCode    = 1001
	ErrorCode      = 1002
	ServeErrorCode = 1003
)

func State(state int) (msg string) {
	if state == SuccessCode {
		return "成功"
	}
	if state == ErrorCode {
		return "失败"
	}
	if state == ServeErrorCode {
		return "服务失败"
	}
	return "错误"
}

func NetServe(num int) (state int, msg string) {
	if num == 1 {
		return SuccessCode, State(SuccessCode)
	}
	if num == 2 {
		return ErrorCode, State(ErrorCode)
	}
	if num == 3 {
		return ServeErrorCode, State(ServeErrorCode)

	}
	return -1, "未知错误"
}

func main() {
	code, msg := NetServe(-1)
	fmt.Println(code, msg)
}
