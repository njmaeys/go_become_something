package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// DummyStringVal just testing some stuff out
type DummyStringVal struct {
	MyVal string
}

// dummyString returns a dummy value.
func dummyString(valToReturn string) DummyStringVal {
	var myVal DummyStringVal

	myVal.MyVal = valToReturn

	return myVal
}

// inbound takes in a simple request
func inbound(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
}

func main() {
	fmt.Println("Go become something!")

	dummy := dummyString("dummy string")
	fmt.Println(dummy)

	///////////////////// Actually start here /////////////////
	http.HandleFunc("/inbound", inbound)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
