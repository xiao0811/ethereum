package handles

import (
	"regexp"
)

// AddressVerify 验证地址
func AddressVerify(addr string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}
