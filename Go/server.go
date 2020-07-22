/*Prime or not server*/

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func prime(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.Bytes()

	temp := make(map[string]uint64)
	err := json.Unmarshal(newStr, &temp)
	if err != nil {
		return
	}
	ans := MillerRabin(temp["value"])
	if ans {
		w.Write([]byte("True"))

	} else {
		w.Write([]byte("False"))
	}
}

func main() {
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/prime", prime)
	if err := http.ListenAndServe(":8001", nil); err != nil {
		panic(err)
	}

}
