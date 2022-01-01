package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	title()
	reader := bufio.NewReader(os.Stdin)

	message("$", "Please specify the URL: ")
	prompt()
	url, _ := reader.ReadString('\n')
	url = strings.TrimSuffix(url, "\n")

	cleanurl := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(url, "https://", ""), "http://", ""), "www.", "")
	cleanurl = strings.Split(cleanurl, "/")[0]
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:80", cleanurl), 1*time.Second)
	if err != nil {
		error("This URL is unreachable.")
		return
	}
	defer conn.Close()
}

func message(prefix string, m string) {
	fmt.Printf("\u001b[32;1m[%s] \u001b[0m%s\n", prefix, m)
}

func error(e string) {
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
