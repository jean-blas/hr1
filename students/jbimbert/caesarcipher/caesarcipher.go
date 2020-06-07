// run this code as follows:
//    OUTPUT_PATH=res.txt ./caesarcipher
// and see the result in file res.txt
package main

import (
	"fmt"
)

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	res := make([]rune, len(s))
	for _, c := range s {
		switch {
		case c >= 'a' && c <= 'z':
			res = append(res, 'a'+(c-'a'+k)%26)
		case c >= 'A' && c <= 'Z':
			res = append(res, 'A'+(c-'A'+k)%26)
		default:
			res = append(res, c)
		}
	}
	return string(res)
}

func main() {
	var k, l int32
	var s string

	fmt.Scanf("%d\n", &l)
	fmt.Scanf("%s\n", &s)
	fmt.Scanf("%d\n", &k)

	result := caesarCipher(s, k)

	fmt.Printf("%s\n", result)
}
