package common

import (
	"crypto/md5"
	"fmt"
	"strings"
	"sort"
)

func Password(len int, pwdO string) (pwd string, salt string) {

	defaultPwd := "wangsy518"

	pwd = Md5([]byte(defaultPwd))
	return pwd, salt
}
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func GetSign(ma map[string]string, token string)string {
	keylist:=make([]string,0)
	for k,v := range ma {
		keylist = append(keylist,strings.ToLower(k+"="+v))
	}
	sort.Sort(sort.StringSlice(keylist))
	keylist = append(keylist,token)

	str:=strings.ToLower(strings.Join(keylist , ""))
	//logs.Debug(str)
	return strings.ToUpper(Md5([]byte(str)))
}
