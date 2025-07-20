package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/mail"
	"os"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

func main() {
	file, err := os.Open("results/race-result.eml")

	if err != nil {
		log.Panic(err.Error())
	}

	message, err := mail.ReadMessage(file)

	if err != nil {
		log.Panic(err.Error())
	}

	buf := new(strings.Builder)

	_, err = io.Copy(buf, message.Body)

	if err != nil {
		log.Panic(err.Error())
	}

	body := buf.String()

	var html string = getHtml(body)
	var event Event = parseEvent(html, message.Header.Get("Subject"))
	fmt.Println(event)
}

func getHtml(data string) string {
	scanner := bufio.NewScanner(strings.NewReader(data))

	var isHtml bool = false
	var htmlStrings []string
	for scanner.Scan() {
		var line string = scanner.Text()

		if strings.TrimRight(line, "\n") == "<html>" {
			isHtml = true
		}

		if isHtml {
			htmlStrings = append(htmlStrings, line)
			if strings.TrimRight(line, "\n") == "</html>" {
				isHtml = false
			}
		}
	}
	return strings.Join(htmlStrings, "")
}

func parseEvent(rawHtml string, subject string) Event {
	rootNode, _ := html.Parse(strings.NewReader(rawHtml))
	var tables []*html.Node = searchHtml(rootNode, "table", []*html.Node{})
	var driverInfoHtml []*html.Node = searchHtml(tables[0], "tr", []*html.Node{})

	var driverInfo DriverInfo = DriverInfo{
		Name: extractTextIter(driverInfoHtml[3])[1],
	}

	var raceInfo Event = Event{
		Location: getLocationFromSubject(subject),
		Position: stripPosition(extractTextIter(driverInfoHtml[4])[2]),
		RaceType: extractTextIter(driverInfoHtml[6])[1],
	}

	var raceData []DriverTime
	var firstRow = true
	for _, row := range searchHtml(tables[2], "tr", []*html.Node{}) {
		if firstRow == true {
			firstRow = false
			continue
		}
		if extractTextIter(row)[0] == raceInfo.Position {
			data := DriverTime{
				Pos:    extractTextIter(row)[0],
				Kart:   extractTextIter(row)[1],
				Racer:  driverInfo.Name,
				Best:   extractTextIter(row)[2],
				NoLaps: extractTextIter(row)[3],
				Avg:    extractTextIter(row)[4],
				Gap:    extractTextIter(row)[5],
			}
			raceData = append(raceData, data)
		} else {
			data := DriverTime{
				Pos:    extractTextIter(row)[0],
				Kart:   extractTextIter(row)[1],
				Racer:  extractTextIter(row)[2],
				Best:   extractTextIter(row)[3],
				NoLaps: extractTextIter(row)[4],
				Avg:    extractTextIter(row)[5],
				Gap:    extractTextIter(row)[6],
			}
			raceData = append(raceData, data)
		}
	}
	raceInfo.DriverTimes = raceData
	raceInfo.DriverInfo = driverInfo
	return raceInfo
} 

func getLocationFromSubject(subject string) string {
	if strings.Contains(subject, "Milton Keynes") {
		return "Milton Keynes"
	}
	return "Unrecognised"
}

func searchHtml(n *html.Node, term string, result []*html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == term {
		result = append(result, n)
	}

	// Recursively visit children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = searchHtml(c, term, result)
	}

	return result
}

func extractTextIter(n *html.Node) []string {
	var data []string
	// start with just the root
	stack := []*html.Node{n}

	for len(stack) > 0 {
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Type == html.TextNode {
			txt := strings.TrimSpace(node.Data)
			if txt != "" {
				data = append(data, txt)
			}
		}

		// push children in reverse so we process them in order
		for c := node.LastChild; c != nil; c = c.PrevSibling {
			stack = append(stack, c)
		}
	}
	return data
}

func stripPosition(position string) string {
	for index, char := range position {
		if !unicode.IsDigit(char) {
			return position[:index]
		}
	}
	return position
}

type DriverTime struct {
	Pos    string
	Kart   string
	Racer  string
	Best   string
	NoLaps string
	Avg    string
	Gap    string
}

type DriverInfo struct {
	Name     string
	Pos      string
	RaceType string
}

type Event struct {
	Location    string
	RaceType    string
	Position    string
	DriverInfo  DriverInfo
	DriverTimes []DriverTime
}
