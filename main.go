package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	title()
	reader := bufio.NewReader(os.Stdin)

	message("$", "Please specify the URL: ")
	prompt()
	url, _ := reader.ReadString('\n')
	url = strings.TrimSuffix(url, "\n")

	res := getResponse(url)
	if !(strings.Contains(res, ".png") || strings.Contains(res, ".jpg") || strings.Contains(res, ".jpeg")) {
		errorMsg("This site does not contain images.")
		return
	}

	for i := 0; i < len(res)-4; i++ {
		if res[i] == 's' && res[i+1] == 'r' && res[i+2] == 'c' && res[i+3] == '=' && res[i+4] == '"' {
			img := ""
			for j := i + 5; j < len(res)-4; j++ {
				if res[j] == '"' {
					break
				} else {
					img += fmt.Sprint(string(res[j]))
				}
			}
			if strings.HasSuffix(img, "png") || strings.HasSuffix(img, "jpeg") || strings.HasSuffix(img, "jpg") || strings.HasSuffix(img, "svg") || strings.HasSuffix(img, "ico") {
				fmt.Println(img)
			}
			img = ""
		}
	}
}

func getResponse(link string) string {
	res, err := http.Get(link)
	if err != nil {
		errorMsg("This URL is unreachable.")
		return ""
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		errorMsg("This URL is unreachable.")
		return ""
	}
	return string(content)
}

func message(prefix string, m string) {
	fmt.Printf("\u001b[32;1m[%s] \u001b[0m%s\n", prefix, m)
}

func errorMsg(e string) {
	fmt.Printf("\u001b[31;1m[!] \u001b[0m%s\n", e)
}

func prompt() {
	fmt.Print("\u001b[36;1m\u001b[4mgid\u001b[0m > \u001b[33m")
}

func title() {
	fmt.Println()
	fmt.Println("\u001b[33;1m --==<[{ \u001b[37;1mGo Image Downloader \u001b[33;1m }]>==--")
	fmt.Println()
}
