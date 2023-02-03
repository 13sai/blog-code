package pkg

import "github.com/13sai/gohelper/str"

func MD(s string) string {
	return str.Md5(s)
}
