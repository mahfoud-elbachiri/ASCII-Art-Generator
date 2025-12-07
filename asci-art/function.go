package function

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Checkascii(s string) string {
	rr := ""
	for i := 0; i < len(s); i++ {
		if (s[i] >= 32 && s[i] <= 126) || s[i] == 10 || s[i] == 13 {
			rr += string(s[i])
		}
	}
	return rr
}

func Output(s string, ss http.ResponseWriter) {
	os.Remove("file.txt")
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_RDONLY, 0o644)
	if err != nil {
		fmt.Fprint(ss, `<h1 style="text-align:center">Internal Server Error</h1>`)
		http.Error(ss, "", http.StatusInternalServerError)
		return
	}
	file.WriteString(s)
}

func Art(word string, banner string) string {
	banner = banner + ".txt"
	print := ""
 
	numb := 1

 
	file, err := os.ReadFile("banner/" + banner)
	if err != nil {
		panic(err)
	}
	data := strings.ReplaceAll(string(file), "\r\n", "\n")
	Letters := strings.Split(data[numb:], "\n\n")

	var matrix []string

	for i := 0; i < 8; i++ {

		for j := 0; j < len(word); j++ {

			lines := strings.Split(Letters[int(rune(word[j])-32)], "\n")
			matrix = append(matrix, lines[i])

		}

		matrix = append(matrix, "\n")

	}
	print = strings.Join(matrix, "")

	return print
}
