// mxTranslate project main.go
package main

import (
	"log"
	"net/http"
)

const API_KEY = "AIzaSyArbtRYU_FDUqC1SuquW-vLC10I-a0EO3U"

func init() {
	log.Println("mxTranslate svr start")
	switchPwd()
	loggerInit()
}

func main() {
	http.Handle("/", getHandler())
	http.HandleFunc("/translate", translateHandler)
	log.Fatalln(http.ListenAndServe(":16022", nil))
}
