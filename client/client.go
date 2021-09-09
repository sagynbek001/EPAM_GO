package client

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{}
}

func (s *Client) Run() error {
	s.httpClient = &http.Client{Timeout: time.Duration(1) * time.Second}

	for {
		// taking our message from the console
		fmt.Print("Enter the city name: ")
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		// message[:len(message) - 1] - removes '\n' for logging
		cityName := message[:len(message)-2]

		resp, err := s.httpClient.Get("http://localhost:8080/weather?city=" + cityName)
		if err != nil {
			fmt.Printf("Error %s", err)
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error %s", err)
			return err
		}
		fmt.Printf("Body : %s", body)
	}
}
