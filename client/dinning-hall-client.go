package client
//
//import (
//	"bytes"
//	component "dinning-hall/components"
//	"encoding/json"
//	"fmt"
//	"io"
//	"log"
//	"net/http"
//	"os"
//)
//
//func SendOrderRequest(order component.Order) {
//	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request){
//		fmt.Fprintf(w, "Hi")
//	})
//
//	fmt.Printf("Starting server at port 8080\n")
//	if err := http.ListenAndServe(":8080", nil); err != nil {
//		log.Fatal(err)
//	}
//	payloadBuf := new(bytes.Buffer)
//	json.NewEncoder(payloadBuf).Encode(order)
//	req, _ := http.NewRequest("POST", "http://localhost:8080/order", payloadBuf)
//
//	client := &http.Client{}
//	res, e := client.Do(req)
//	if e != nil {
//		log.Fatal(e)
//	}
//
//	defer res.Body.Close()
//
//	fmt.Println("response Status:", res.Status)
//
//	// Print the body to the stdout
//	io.Copy(os.Stdout, res.Body)
//}




