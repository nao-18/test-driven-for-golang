package main

import (
	"log"
	"net/http"

	server "github.com/nao-18/test-driven-for-golang/http-server"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	// fmt.Println(Hello("world", ""))
	// sleeper := &mocking.DefaultSleeper{}
	// mocking.Countdown(os.Stdout, sleeper)
	// http.ListenAndServe(":9000", http.HandlerFunc(di.MyGreeterHandler))

	// http-server
	handler := http.HandlerFunc(server.PlayerServer)
	if err := http.ListenAndServe(":9000", handler); err != nil {
		log.Fatalf("could not listen on port 9000 %v", err)
	}
}
