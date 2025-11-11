package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const colorRed = "\033[0;31m"

type quoteDTO struct {
	Quote     string
	Character string
}

type messageDTO struct {
	Author  string
	Message string
}

func readBody(res *http.Response) (*quoteDTO, error) {
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("unable to read body")
	}

	var data []quoteDTO
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data[0], nil
}

func getQuote() (*quoteDTO, error) {
	response, err := http.Get(os.Getenv("QUOTES_URL"))

	if err != nil {
		return nil, errors.New("request failed")
	}

	return readBody(response)
}

func postMessage(message *messageDTO) error {
	jsonData, err := json.Marshal(message)

	if err != nil {
		return err
	}

	_, err = http.Post(os.Getenv("MESSAGES_URL"), "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}
	return nil
}

func main() {
	godotenv.Load()
	for {
		quote, err := getQuote()

		if err != nil {
			fmt.Fprintf(os.Stdout, "Red: %s %s", colorRed, "couldn't get qoute, quiting :(")
			return
		}

		err = postMessage(&messageDTO{
			Author:  quote.Character,
			Message: quote.Quote,
		})

		if err != nil {
			fmt.Fprintf(os.Stdout, "Red: %s %s", colorRed, "couldn't post message, quiting :(")
			return
		}

		time.Sleep(5 * time.Second)
	}

}
