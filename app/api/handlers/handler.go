package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func ListMachines() {

}

func getNetworkData() {
	// Open the file.
	f, _ := os.Open("./go_data/parent.json")

	resp, err := http.Post("http://127.0.0.1:5000/v1/parent", "application/json", f)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	time.Sleep(2 * time.Second)
	defer resp.Body.Close()
}
