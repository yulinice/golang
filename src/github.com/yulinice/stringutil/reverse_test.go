//包含_test.go结尾的,包含名为Text...
//且签名为 func (t *testing.T)函数的GO文件
//若该函数调用了像 t.Error 或 t.Fail 这样表示失败的函数，此测试即表示失败。
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in,want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) ==%q", c.in, got, c.want)
		}
	}
}
