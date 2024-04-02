package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type product struct {
	Name     string `json:"productname"`
	Price    int
	Password string   `json:"-"`              //<-- '-' means in json, pasword not show
	Tags     []string `json:"tags,omitempty"` //<-- 'omitempty' means nil/null not show
}

func main() {
	fmt.Println("creating Json Data Structure")
	// EncodeJson()
	// DecodeJson()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter value:")
	// comma ok || err err
	input, _ := reader.ReadString('\n')
	fmt.Println("value is ", input)
	fmt.Printf("type is %T", input)

}

func EncodeJson() {

	jsonProduct := []product{
		{"Laptop", 50000, "abc123", []string{"developer", "js"}},
		{"Keyboard", 500, "abc234", []string{"laptop use", "circle"}},
		{"Mouse", 1000, "abc345", nil},
	}

	//package this data as json data
	finalJson, err := json.MarshalIndent(jsonProduct, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromweb := []byte(`
		{
			"productname": "Laptop",
			"Price": 50000,
			"tags": ["developer","js"]
	}
	`)

	var product product
	checkValid := json.Valid(jsonDataFromweb)

	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(jsonDataFromweb, &product)
		fmt.Printf("%#v\n", product)
	} else {
		fmt.Println("Json was not Valid")
	}

	// some cases where you just want to add to key value pair data

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromweb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
