package di

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writter io.Writer, name string) {
	fmt.Fprintf(writter, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
