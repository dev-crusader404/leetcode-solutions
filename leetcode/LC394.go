package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stk := []byte{}
	for i := range s {
		if s[i] != ']' {
			stk = append(stk, s[i])
		} else {
			var sb strings.Builder
			var idx int
			size := len(stk)
			for len(stk) > 0 && stk[size-1-idx] != '[' {
				idx++
			}
			sb.Write(stk[size-idx:])
			stk = stk[:size-idx-1]
			temp := sb.String()
			var num []byte
			var numIdx int
			for len(stk)-numIdx > 0 && stk[len(stk)-1-numIdx] >= '0' && stk[len(stk)-1-numIdx] <= '9' {
				numIdx++
			}
			num = stk[len(stk)-numIdx:]
			stk = stk[:len(stk)-numIdx]
			n, _ := strconv.Atoi(string(num))
			for n > 0 {
				for k := range temp {
					stk = append(stk, temp[k])
				}
				n--
			}
		}
	}
	return string(stk)
}

func RunLC394() {
	fmt.Println(decodeString("3[a]2[bc]"))
}
