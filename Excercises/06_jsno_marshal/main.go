package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	First_Name string
	Last_Name  string
	Age        int
	Mobile_No  string
}

func main() {

	var unmarshal_data = []User{}
	users := []User{{
		First_Name: "Vikash", Last_Name: "Parashar", Age: 28, Mobile_No: "XXXXxxxxx",
	}, {
		First_Name: "Rampati", Last_Name: "Devi", Age: 48, Mobile_No: "XXXXxxxxx",
	}, {
		First_Name: "Niyati", Last_Name: "Parashar", Age: 3, Mobile_No: "XXXXxxxxx",
	}, {
		First_Name: "Ritika", Last_Name: "Parashar", Age: 4, Mobile_No: "XXXXxxxxx",
	}, {
		First_Name: "Khushboo", Last_Name: "Pandayr", Age: 24, Mobile_No: "XXXXxxxxx",
	}}
	// for i, v := range users {
	// 	fmt.Println(i)
	// 	fmt.Println(v)
	// }

	// 01. json marshal
	data, err := json.Marshal(users)
	if err != nil {
		log.Fatal("failed to covert into json")
	}
	fmt.Println(string(data)) // converting []byte data into json string

	// 02. json marshal indent
	data2, err := json.MarshalIndent(users, "", "")
	if err != nil {
		log.Fatal("failed to covert into json")
	}
	fmt.Println(string(data2)) // converting []byte data into json string

	/*
		// now lets convert json data into golang struct
		// this can be done using json unmarshal

		if err := json.Unmarshal(data2, &unmarshal_data); err != nil {
			log.Fatal("faile to convert json data into golang data type")
		}
		// fmt.Println(unmarshal_data)
		for _, v := range unmarshal_data {
			fmt.Println(v)
		}
	*/

}
