package main

import (
	"encoding/json"
	"fmt"
    "time"
)

func main() {
	type myObject struct {
		ArrayValue []int `json:"arrayValue"`
	}

	type myJSON struct {
		IntValue        int       `json:"intValue"`
		BoolValue       bool      `json:"boolValue"`
		StringValue     string    `json:"stringValue"`
		DateValue       time.Time `json:"dateValue"`
		ObjectValue     *myObject `json:"objectValue"`
		NullStringValue *string   `json:"nullStringValue"`
		NullIntValue    *int      `json:"nullIntValue"`
	}

	otherInt := 4321
	data := &myJSON{
		IntValue:    1234,
		BoolValue:   true,
		StringValue: "hello!",
		DateValue:   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		ObjectValue: &myObject{
			ArrayValue: []int{1, 2, 3, 4},
		},
		NullStringValue: nil,
		NullIntValue:    &otherInt,
	}

	jsonData, _ := json.Marshal(data)
	fmt.Printf("json data: %s\n", jsonData)
    otherInt = 0
	jsonData, _ = json.Marshal(data)
	fmt.Printf("json data: %s\n", jsonData)

}
