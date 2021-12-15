package main

import(
"fmt"
)

func main() {
	str := ""
	for {
		n, _ := fmt.Scan(&str)
		if n == 0 {
			break
		} else {
			mp := map[byte]bool{}
			b1 := []byte(str)
			for  _, v := range b1 {
				mp[v] = true
			}
			length := len(b1)
			for i := 0; i < length; i++ {
				v := b1[i]
				if mp[v] == true {
					mp[v] = false
					continue
				}else {

				}
			}

			fmt.Println(string(b1))
		}
	}
}
