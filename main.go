package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 5

func main() {
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			monitoring()
		case 2:
			showLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Command not found")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - exit the program")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func monitoring() {
	fmt.Println("Monitoring...")

	sites := readFileSites()

	for i := 0; i < 5; i++ {
		for _, site := range sites {
			resp, err := http.Get(site)

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			if resp.StatusCode == 200 {
				fmt.Println("Ok", site)
				saveLog(site, true)
			} else {
				fmt.Println("Error", site)
				saveLog(site, false)
			}
		}
		time.Sleep(delay * time.Second)
	}

}

func readFileSites() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func saveLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(file))
}
