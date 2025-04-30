package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/mail"
	"os"
	"strings"
)

func main() {
	fmt.Println("Karting Rules!")

	file, err := os.Open("race-result.eml")

	if err != nil {
		log.Panic(err.Error())
	}

	m, err := mail.ReadMessage(file)

	if err != nil {
		log.Panic(err.Error())
	}

	buf := new(strings.Builder)

	n, err := io.Copy(buf, m.Body)

	fmt.Printf("Copied: %d bytes \n\n", n)
	if err != nil {
		log.Panic(err.Error())
	}

	body := buf.String()

	scanner := bufio.NewScanner(strings.NewReader(body))

	var isHtml bool = false
	var html []string
	for scanner.Scan() {
		var line string = scanner.Text()

		if strings.TrimRight(line, "\n") == "<html>" {
			isHtml = true
		}

		if isHtml {
			html = append(html, line)
			if strings.TrimRight(line, "\n") == "</html>" {
				isHtml = false
			}
		}
	}
	fmt.Print(html)
}
