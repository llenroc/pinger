package dummy

import (
	"io"
	"log"
	"net/http"
)

// SampleData - Handler for sample data return/data post
func SampleData(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		io.WriteString(w, "[\"http://127.0.0.1:3000/status\"]\n")
	} else {
		log.Println("Got a post")
	}
}

// StatusResponse - 'OK' handler
func StatusResponse(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK\n")
}

func main() {
	http.HandleFunc("/http_checks", SampleData)
	http.HandleFunc("/status", StatusResponse)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
