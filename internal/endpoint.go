package internal

import (
	"io/ioutil"
	"log"
	"net/http"
)

/*
Note to self that capital letter means exported function and can be used outside
of the actual file it is in.
*/

// Inbound takes in a simple request
func Inbound(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
}
