// run this code as follows:
//    OUTPUT_PATH=res.txt ./camelcase
// and see the result in file res.txt
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// A solution using regexp
func camelcase2(s string) int32 {
	re := regexp.MustCompile(`[A-Z]`)
	res := re.FindAll([]byte(s), -1)
	fmt.Printf("%q\n", res)
	return int32(1 + len(res))
}

// Another solution using characters
func camelcase(s string) int32 {
	var answer int32
	answer = 1
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			answer++
		}
	}
	return answer
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := camelcase(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
