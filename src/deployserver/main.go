package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func reLaunch(){
	cmd := exec.Command("sh","./deploy.sh")
	err := cmd.Start()
	if err != nil{
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"<h1> deploy server !")
	reLaunch()
}

func main(){
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":8001",nil)
	fmt.Println("listen at 8001")
}