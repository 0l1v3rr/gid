package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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
	if !(strings.Contains(res, ".png") || strings.Contains(res, ".jpg") || strings.Contains(res, ".jpeg") || strings.Contains(res, ".gif") || strings.Contains(res, ".svg") || strings.Contains(res, ".ico")) {
		errorMsg("\nThis site does not contain images.")
		return
	}

	os.Mkdir("img", 0755)

	fmt.Println()
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
				newurl := ""
				if strings.HasPrefix(img, "https://") || strings.HasPrefix(img, "http://") {
					newurl = img
				} else {
					newurl = "http://" + strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(url, "https://", ""), "http://", ""), "www.", ""), "/")[0] + "/" + img
				}
				derr := downloadImg(newurl, fmt.Sprintf("img/%v", strings.Split(img, "/")[len(strings.Split(img, "/"))-1]))
				if derr != nil {
					if derr.Error() == "an unknown error occurred" {
						errorMsg("This site is not reachable.")
						return
					} else if derr.Error() == "received a non 200 status code" {
						errorMsg(fmt.Sprintf("%s - unreachable endpoint.", newurl))
					}
				} else {
					fmt.Printf("\u001b[0m[$] \u001b[32;1m%s \u001b[0mdownloaded.\n", newurl)
				}
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

func downloadImg(link, fileName string) error {
	response, err := http.Get(link)
	if err != nil {
		return errors.New("an unknown error occurred")
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received a non 200 status code")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return errors.New("an unknown error occurred while creating the file")
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return errors.New("an unknown error occurred while copying the file")
	}

	return nil
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
