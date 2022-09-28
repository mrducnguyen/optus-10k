package main

import (
	"bufio"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	dataMap := loadData()

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		query := strings.ToLower(html.EscapeString(r.URL.Query().Get("q")))
		fmt.Fprintf(w, strconv.FormatBool(dataMap[query]))
	})

	log.Println("Server listening on :8081 ...")
	log.Fatal(http.ListenAndServe(":8081", nil))
	log.Println("Server stopped!")
}

func loadData() map[string]bool {
	m := make(map[string]bool)
	file, err := os.Open("./phone_email.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		m[strings.ToLower(scanner.Text())] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return m
}
