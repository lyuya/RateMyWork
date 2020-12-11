package src

import (
	"./model"
	"encoding/json"
)
import (
	"fmt"
	"log"
	"net/http"
)

var Files []model.File

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here is the welcome page for RateMyWork API!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/files", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Files)
}

func main() {
	Files = []model.File{
		model.File{
			Id:       "111",
			FileName: "ok",
			FileType: "json",
		},
		model.File{
			Id:       "123",
			FileName: "okq",
			FileType: "json",
		},
	}
	handleRequests()
}
