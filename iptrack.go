package main

import (
	"fmt"
	"net/http"
	"io"
)

func main(){
	resp, err:= http.Get("https://api.ipify.org?format=json")

	if err != nil {
		fmt.Println("Cant get data from API: %w", err.Error())
        
	}
	data, err:= io.ReadAll(resp.Body)
	if err!= nil {
        fmt.Println(err)
    }
	fmt.Println(string(data))

}

