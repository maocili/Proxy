package test2

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func M() {
	m := map[int]int{}
	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	//n := 10
	ks := reflect.ValueOf(m).MapRange()
	var list []int64
	for ks.Next() {
		list = append(list, ks.Key().Int())
	}
	jsontext, _ := json.Marshal(list)
	fmt.Println(string(jsontext))

}
