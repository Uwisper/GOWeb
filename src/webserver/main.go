package main

import (
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"<h1>hello,this is my first page</h1>" +
		"<h2>we changed a little to test deployserver</h2>" +
		"sh,is deploy server work？" +
		"<div>God,is this work?</div>")
}

func main(){
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":8000",nil)

}
