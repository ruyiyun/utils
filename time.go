package utils

import (
	"strconv"
	"strings"
	"time"
)

const TIME_LAYOUT_OFTEN = "2006-01-02 15:04:05"

// DateFormat pattern rules.
var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06", //A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2", // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon", // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

// Parse Date use PHP time format.
func DateParse(dateString, format string) (time.Time, error) {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return time.ParseInLocation(format, dateString, time.Local)
}

// Date takes a PHP like date func to Go's time format.
func Date(t time.Time, format string) string {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return t.Format(format)
}
func DateFormat(t time.Time, layout string) (datestring string) {
	datestring = t.Format(layout)
	return
}

// 解析常用的日期时间格式：2014-01-11 16:18:00，东八区
func TimeParseOften(value string) (time.Time, error) {
	local, _ := time.LoadLocation("Local")
	return time.ParseInLocation(TIME_LAYOUT_OFTEN, value, local)
}

//返回当前时区的当前时间
func TimeLocal() time.Time {
	format := "2006-01-02 15:04:05 -07:00 "
	dateString := DateFormat(time.Now(), format)
	formatedDate, _ := DateParse(dateString, format)
	return formatedDate
}

//返回当前时区的当前时间
func DateLocal() time.Time {
	format := "2006-01-02"
	dateString := DateFormat(time.Now(), format)
	formatedDate, _ := DateParse(dateString, format)
	return formatedDate
}

//返回当前时区的当前时间
func TimeLocalString() string {
	dateString := DateFormat(TimeLocal(), TIME_LAYOUT_OFTEN)
	return dateString
}

//返回当前时区的当前月日时,120615,12月6号15时
func TimeLocalMDHString() string {
	dateString := DateFormat(TimeLocal(), "010215")
	return dateString
}

//将unix 时间戳转换为时间字符串
//1441070992=>2015-09-01 09:29:52
func Timestamp2String(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(TIME_LAYOUT_OFTEN)
}

//把指定的时间转换为字符串
func DateTime2String(dt time.Time) string {
	dateString := DateFormat(dt, TIME_LAYOUT_OFTEN)
	return dateString
}

//得到多少分钟前的时间
func TheTimeString(counts time.Duration) string {
	baseTime := time.Now()
	date := baseTime.Add(counts)

	dateString := DateFormat(date, TIME_LAYOUT_OFTEN)
	return dateString
}

//得到多少分钟前的时间
func TheTime(counts time.Duration) time.Time {
	baseTime := time.Now()
	date := baseTime.Add(counts)

	return date
}

//返回当前月份，yms str类型，ymi int 类型
func TheYearMonthString() (yms string, ymi int) {
	yms = DateFormat(TimeLocal(), "200601")
	ymi, _ = strconv.Atoi(yms)
	return yms, ymi
}

//以时间作为文件夹目录
func GetDateAsDirName() string {
	format := "2006/01/02"
	dateString := DateFormat(time.Now(), format)
	return dateString
}

//以时间作为文件夹目录
func GetDateYYYYMMDD() string {
	format := "20060102"
	dateString := DateFormat(time.Now(), format)
	return dateString
}
