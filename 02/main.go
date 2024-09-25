package main

import (
	"encoding/json"
	"fmt"
)
 
func Appender(slice []int, val int) {
    slice[0] = 0
    slice = append(slice, val)
    slice[1] = 0
}
type Alert struct {
	Message interface{} `json:"message"`
}

type Alerts struct {
	Alerts []Alert `json:"alerts"`
}
 
func main() {
    // slice := []int{1, 2, 3, 4, 5}
    // Appender(slice, 6)
    // fmt.Println(slice)


	// Input:= []int{1,2,3,4,1,2,2}

	// fmt.Printf("Output  = %v\n",findDuplicate(Input))

	jsondata := `{
		"alerts": [
		  {
			"message": "hello world"
		  },
		  {
			"message": 911
		  },
		  {
			"message": "hello india"
		  }
		]
	  }`


	  var alerts Alerts

	  err := json.Unmarshal([]byte(jsondata) , &alerts)
	  if err != nil {
		fmt.Printf("error while unmarshal , err  = %v\n",err)
		return
	  }

	  stringCount := 0
	  intCount := 0

	  for _ , alert := range alerts.Alerts {
		switch alert.Message.(type) {
		case string:
			stringCount++
		case float64:
			intCount++
		}
	  }

	  fmt.Printf("Totasl strings are  = %d , Total int messages are  = %v", stringCount  , intCount)
}


func findDuplicate(slice []int) []int{

	elsemntsCOunt := make(map[int]int)
	duplicates := []int{}



	for _ , num := range slice {
		elsemntsCOunt[num]++
	}


	for key , value := range elsemntsCOunt {
		if value > 1 {
			duplicates = append(duplicates, key)
		}
	}


return duplicates

}