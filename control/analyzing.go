package control

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"interface1/public/singleSelectParam"
	"interface1/public/batchSelectParam"
	"interface1/public/deleteParam"
)

func ToSingleSelectParam(r *http.Request) singleSelectParam.Param {
	fmt.Println("method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return singleSelectParam.Param{}
	}
	println("json:", string(body))

	var a singleSelectParam.Param
	if err = json.Unmarshal(body, &a); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return singleSelectParam.Param{}
	}

	return a
}

func ToBatchSelectParam(r *http.Request) batchSelectParam.Param {
	fmt.Println("method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return batchSelectParam.Param{}
	}
	println("json:", string(body))

	var a batchSelectParam.Param
	if err = json.Unmarshal(body, &a); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return batchSelectParam.Param{}
	}

	return a
}

func ToDeleteParam(r *http.Request) deleteParam.Param {
	fmt.Println("method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return deleteParam.Param{}
	}
	println("json:", string(body))

	var a deleteParam.Param
	if err = json.Unmarshal(body, &a); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return deleteParam.Param{}
	}

	return a
}
