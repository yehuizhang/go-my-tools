package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

/* Documentation:
slice https://blog.golang.org/slices-intro
*/
func main() {

	filepath := "data/AccessLog_210801.csv"
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln("Could not open file: ", filepath)
	}

	reader := csv.NewReader(file)

	header, err := reader.Read()

	if err != nil {
		log.Fatalln("Unable to read the header")
	}

	log.Println(strings.Join(header, ","))

	// s := make([][]string, 1)
	var s [][]string // just declare

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Reading file error")
		}
		s = append(s, line)
	}
	log.Println("Lenth: ", len(s))
	getIpForLogin(s)
}

func getIpForLogin(list [][]string) {
	m := make(map[string]int)

	log.Println("getIpForLogin output")

	for _, row := range list {
		if row[0] == "Information" {

			ip := row[4]
			if ip[0:7] == "192.168" {
				ip = "192.168.*"
			}
			m[ip] = m[ip] + 1
		}
	}

	for k, v := range m {
		log.Println(k, v)
	}
}
