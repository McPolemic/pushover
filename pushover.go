package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	formValues := url.Values{}
	formValues = loadEnvIntoValues(formValues)
	formValues = parseFlagsIntoValues(formValues)
	makeRequest(formValues)
}

func loadEnvIntoValues(formValues url.Values) url.Values {
	token := os.Getenv("PUSHOVER_TOKEN")
	user := os.Getenv("PUSHOVER_USER")

	if token == "" {
		log.Fatal("Need environment variable PUSHOVER_TOKEN")
	} else {
		formValues.Set("token", token)
	}

	if user == "" {
		log.Fatal("Need environment variable PUSHOVER_USER")
	} else {
		formValues.Set("user", user)
	}

	return formValues
}

func parseFlagsIntoValues(formValues url.Values) url.Values {

	titlePtr := flag.String("title",
		"",
		"Notification title")
	messagePtr := flag.String("message",
		"",
		"Notification message body")
	urlPtr := flag.String("url",
		"",
		"URL for more information")

	flag.Parse()

	if *titlePtr != "" {
		formValues.Set("title", *titlePtr)
	}

	if *messagePtr != "" {
		formValues.Set("message", *messagePtr)
	} else {
		log.Fatal("Must provide a --message value")
	}

	if *urlPtr != "" {
		formValues.Set("url", *urlPtr)
	}

	return formValues
}

func makeRequest(formValues url.Values) {
	resp, err := http.PostForm("https://api.pushover.net/1/messages.json", formValues)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
