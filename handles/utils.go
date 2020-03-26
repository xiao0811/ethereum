package handles

import (
	"log"
	"regexp"
)

// AddressVerify 验证地址
func AddressVerify(addr string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}

// ErrorVerify 错误验证
func ErrorVerify(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// 根据
