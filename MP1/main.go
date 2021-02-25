package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"./tcp"
	"./app"
	"gopkg.in/yaml.v2"
)

// Structs for yaml
type Process struct {
	Id int `yaml:"id"`
	Ip string `yaml:"ip"`
	Port string `yaml:"port"`
}
type Config struct {
	MinDelay int `yaml:"minDelay"`
	MaxDelay int `yaml:"maxDelay"`
	Process []Process `yaml:"process"`
}

func main() {
	// Parse command-line flags
	idx, idxValid := ParseFlags()
	if !idxValid {
		return
	}

	// Read config
	var cfg Config
	cfg.GetConf("config.yml")

	// Construct source (sender process) struct
	source := app.Process{
		Id: cfg.Process[idx].Id,
		Ip: cfg.Process[idx].Ip,
		Port: cfg.Process[idx].Port,
		MinDelay: cfg.MinDelay,
		MaxDelay: cfg.MaxDelay,
	}

	// Launch a server
	// Server waits for messages from other processes
	srv := tcp.NewServer(source.Ip + ":" + source.Port)

	// Prompt user
	fmt.Println("\n3 options in this process:")
	fmt.Println("1. Type a command in the format to send message")
	fmt.Println("\tFormat: send [destination ID (0-4)] [message content]")
	fmt.Println("2. Wait for messages")
	fmt.Println("3. Type 'q' to quit")
	fmt.Print("\n")

	// Send delayed messages to other processes
	for {
		// Take user input
		inputCmd := takeUserInput()

		// Break loop if input is q
		if inputCmd == "q" {
			break
		}

		// Check if user input is valid
		// If not valid, try again (continue)
		if !inputIsValid(inputCmd) {
			fmt.Println("\nInput does not match format below. Try again")
			fmt.Println("Format: send [destination ID (0-4)] [message content]")
			fmt.Print("\n")
			continue
		}

		// Parse user input into an array of words
		inputs := strings.Fields(inputCmd)

		// Convert input at index 1 to int
		destIdx, err := strconv.Atoi(inputs[1])
		if err != nil {
			fmt.Println(err)
		}

		// Construct destination (receiver process) struct
		destination := app.Process{
			Id: cfg.Process[destIdx].Id,
			Ip: cfg.Process[destIdx].Ip,
			Port: cfg.Process[destIdx].Port,
		}

		newMsg := app.Message{
			Source: source,
			Destination: destination,
			Content: strings.Join(inputs[2:], " "),
		}

		// Send
		go tcp.UnicastSend(newMsg)
	}

	// Once the loop above breaks, shutdown server
	srv.Stop()
}

func ParseFlags() (int, bool) {
	// Parse command-line arguments
	var id int
	flag.IntVar(&id, "ID", 0, "Process ID for config")
	flag.Parse()
	// Check correctness of arg
	if id < 0 || id > 4 {
		fmt.Println("Please input an ID within the range of 0 to 4")
		return -999, false
	} else {
		fmt.Println("ID entered:", id)
		return id, true
	}
}

// GetConf reads the yaml configuration into Config struct
func (c *Config) GetConf(configPath string) *Config {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

// helper function to take user inputs
func takeUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// returns if s matches the pattern
func inputIsValid(s string) bool {
	pattern := `^send [0-4]{1} `
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		fmt.Println(err)
	}
	return matched
}
