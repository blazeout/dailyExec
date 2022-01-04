package main

func modifyString(s string) string {
	// ? 替换为任何字符, 但是此字符不能和前一个字符或者后一个字符相等
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] == '?' {
			// 两种情况
			// 1. 不能和b[i-1]相等
			// 2. 不能和b[i+1]相等
			var notEqualChar1 byte
			var notEqualChar2 byte
			if i-1 >= 0 {
				notEqualChar1 = b[i-1]
			}
			if i+1 < len(b) {
				notEqualChar2 = b[i+1]
			}
			for char := byte('a'); char <= 'c'; char++ {
				if char != notEqualChar1 && char != notEqualChar2 {
					b[i] = char
					break
				}
			}
		}
	}
	return string(b)
}

func main() {
	// 字符串不适合直接修改
	println(modifyString("??yw?ipkj?"))
}
