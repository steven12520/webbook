package common

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func Password(len int, pwdO string) (pwd string, salt string) {

	defaultPwd := "wangsy518"

	pwd = Md5([]byte(defaultPwd))
	return pwd, salt
}

func creates(ma map[string]string, token string) {

	list := make([]string, 0)

	for k, v := range ma {
		list = append(list, k+"="+v)
	}

	sort.Sort(sort.StringSlice(list))
	fmt.Println(strings.Join(list, ""))
}
