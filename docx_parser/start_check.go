package docx_parser

import (
	"regexp"
)

// 检查是否以阿拉伯数字开头
func startsWithArabicNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	return s[0] >= '0' && s[0] <= '9'
}

// 检查是否以 "第" + 阿拉伯数字开头
func startsWithDiAndArabicNumber(s string) bool {
	matched, _ := regexp.MatchString(`^第[0-9]+`, s)
	return matched
}

// 检查是否以汉字数字开头
func startsWithChineseNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	matched, _ := regexp.MatchString(`[一二三四五六七八九十]`, s)
	return matched
}

// 检查是否以 "第" + 汉字数字开头
func startsWithDiAndChineseNumber(s string) bool {
	matched, _ := regexp.MatchString(`^第[一二三四五六七八九十]`, s)
	return matched
}

// 综合检查函数
func CheckString(s string) bool {
	return startsWithArabicNumber(s) ||
		startsWithDiAndArabicNumber(s) ||
		startsWithChineseNumber(s) ||
		startsWithDiAndChineseNumber(s)
}
