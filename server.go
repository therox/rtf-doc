package rtfdoc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func rtfHandler(w http.ResponseWriter, r *http.Request) {
	// Here we'll be processing reauest body
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("error reading body: %v", err)
		}
		// Parse body as json
		var document []map[string]interface{}
		err = json.Unmarshal(body, &document)
		if err != nil {
			log.Fatalf("can't unmarshal body: %v", err)
		}
		// for _, docItem := range document {

		// }

		fmt.Println(string(body))
	} else {
		w.Write([]byte("Method  is not applicable"))

	}
	return
}

// RunServer run listener server for doc generate
func RunServer(port int) {

	http.HandleFunc("/generate_rtf_doc", rtfHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
