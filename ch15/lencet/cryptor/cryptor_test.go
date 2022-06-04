package cryptor

import (
	"fmt"
	"github.com/duke-git/lancet/v2/cryptor"
	"testing"
)

func TestMd5String(t *testing.T) {
	s := cryptor.HmacMd5("123456", "zzx")
	s1 := cryptor.Md5String("hello word")
	fmt.Println(s)  //5f4c9faaff0a1ad3007d9ddc06abe36d
	fmt.Println(s1) //13574ef0d58b50fab38ec841efe39df4
}
