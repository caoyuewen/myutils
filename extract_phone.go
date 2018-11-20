package myutils

/*
工具类
提取一段文本的所有电话号码
目前仅支持 手机号码 和 12位座机号码

*/

import (
	"regexp"
)

const (
	CellPhone = 11		//手机号码类型
	LandLine  = 12		//座机号码类型

	CellPhoneRegular = `^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$`
	LandLineRegular  = `\d{3}-\d{8}|\d{4}-\d{7}`
)

//获取一段字符串中所有电话号码(手机)
func GetAllPhoneNumber(data string, phoneType int) []string {
	if !(phoneType == CellPhone || phoneType == LandLine) {
		return nil
	}
	var result []string
	dataArr := []rune(data)

	var phoneRegexp *regexp.Regexp
	if phoneType == CellPhone {
		phoneRegexp = regexp.MustCompile(CellPhoneRegular)
	} else {
		phoneRegexp = regexp.MustCompile(LandLineRegular)
	}
	for {
		if len(dataArr) < phoneType {
			break
		}


		s := phoneRegexp.FindString(getSubstring(dataArr, phoneType))
		if s != "" {
			result = append(result, s)
		}
		dataArr = dataArr[1:]
	}
	return result
}

/*截取字符的头几位*/
func getSubstring(data []rune, num int) string {
	if num <= len(data) {
		content := string(data[:num])
		return content

	}
	return ""
}
