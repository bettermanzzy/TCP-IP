package main

import (
	"net/http"
	"fmt"
)

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("handle login",r.URL)
	fmt.Fprintf(w,"----- login ")
}
func main(){
	//http.HandleFunc("/",hello)
	http.HandleFunc("/user/login",login)
	err :=http.ListenAndServe(":8800",nil)
	if err !=nil{
		fmt.Println("http listen failed")
	}
}
