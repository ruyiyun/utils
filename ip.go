package utils

import (
	"net"
	"strconv"
	"strings"
)

func inet_ntoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

// Convert net.IP to int64
func inet_aton(ipnr net.IP) int64 {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

//Ip地址字符串转换为十进制
func IpString2Int64(ipstr string) int64 {

	ip := net.ParseIP(ipstr)
	if ip != nil {
		return inet_aton(ip)
	}

	return 0

}

//将十进制IP转换为IPv4地址样式
func IpInt2String(ipint int64) string {
	ip := inet_ntoa(ipint)
	return ip.String()
}
