// without data struct
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	// Let's first read the `config.json` file
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Let's print the unmarshalled data!
	log.Printf("origin: %s\n", payload["origin"])
	log.Printf("user: %s\n", payload["user"])
	log.Printf("status: %t\n", payload["active"])
}
