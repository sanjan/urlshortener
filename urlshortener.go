package main

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "strings"
    "github.com/pilu/go-base62"
)

type postdata_struct struct {
	LONGURL string `json:"url"`
}

type post_response_struct struct{
	Short string
}

type getdata_struct struct {
	SHORTURL string `json:"short"`
}

type get_response_struct struct{
	Original string	
}

var urlId int = 10000
var urlStore = make(map[string]string)

func parsePOSTreq(rw http.ResponseWriter, req *http.Request) {
    
    decoder := json.NewDecoder(req.Body)

    var pd postdata_struct  
    
    err := decoder.Decode(&pd)
    
    if err != nil {
          log.Println(err.Error())
          http.Error(rw, err.Error(), http.StatusInternalServerError)
          return
    }
    
    log.Println("Received long URL : "+ pd.LONGURL + " for processing")

    encodedID := base62.Encode(urlId)


	urlStore[encodedID]=pd.LONGURL

	shortURL := "http://localhost/" + encodedID

    log.Println("Generated short URL : "+ shortURL)

     
    respcontent := post_response_struct{ Short: shortURL}
    
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(rw, "HTTP 200\n")
    json.NewEncoder(rw).Encode(&respcontent)


    defer req.Body.Close()
    
    urlId++
}

func parseGETreq(rw http.ResponseWriter, req *http.Request) {

    if req.Method == "GET" {

    decoder := json.NewDecoder(req.Body)
        
    var gd getdata_struct  
    err := decoder.Decode(&gd)
    if err != nil {
          log.Println(err.Error())
          http.Error(rw, err.Error(), http.StatusInternalServerError)
          return
    }
    
    log.Println("Received short URL : "+ gd.SHORTURL + " for processing")

    urlelements := strings.Split(gd.SHORTURL, "/")
    shorturlkey := urlelements[len(urlelements)-1]

    originalURL := urlStore[shorturlkey]

    log.Println("Found original URL : "+ originalURL)

    respcontent := get_response_struct{ Original: originalURL }
    
    rw.Header().Set("Content-Type", "application/json") 
    fmt.Fprintf(rw, "HTTP 200\n")
    json.NewEncoder(rw).Encode(&respcontent)

    defer req.Body.Close()
     
    }

}

func main() {
    http.HandleFunc("/shorten", parsePOSTreq)
    http.HandleFunc("/original", parseGETreq)
    log.Println ("Listening on port 8082 ...")
    log.Fatal(http.ListenAndServe(":8082", nil))
}
