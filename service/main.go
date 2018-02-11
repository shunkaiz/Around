
package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json: "lon"`
}

type Post struct {
	User string `json : "user"`
	Message string `json: "message"`
	Location Location `json: "location"`
}

/**

{
	user: "jack",
	message: "hello",
	location:
	{
		lat: 100,
		lon: 10

	},
}


 */
func main() {
	fmt.Println("started-service")
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/search", handlerSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
const (
	DISTANCE = "200km"
)

func handlerSearch(w http.ResponseWriter, r* http.Request){
	fmt.Print("Received a search request")
	lat := r.URL.Query().Get("lat")
	lt,_ := strconv.ParseFloat(lat, 64)
	lon := r.URL.Query().Get("lon")
	ln,_ := strconv.ParseFloat(lon,64)
	ran := DISTANCE
	if val := r.URL.Query().Get("range"); val != "" {
		ran = val + "km"
	}
	fmt.Println("range is ", ran)
	p:= &Post{
		User: "1111",
		Message: "this is wahtever",
		Location:Location{
			Lat:lt,
			Lon:ln,
		},
	}
	js, err := json.Marshal(p)
	if err != nil {
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handlerPost (w http.ResponseWriter, r *http.Request){
	fmt.Print("Received a request for post")
	decoder := json.NewDecoder(r.Body)
	var p Post
	if err:= decoder.Decode(&p); err != nil{
		panic(err)
		return
	}
	fmt.Fprintf(w, "Post received: %s\n", p.Message)
}