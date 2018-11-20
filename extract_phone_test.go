package myutils

import (
	"fmt"
	"testing"
)

func TestGetAllPhoneNumber(t *testing.T) {
	strData:="1835564632132165489798132134165465465465987测试 测试 18382197831"
	numbers := GetAllPhoneNumber(strData, CellPhone)
	fmt.Println(numbers)
}