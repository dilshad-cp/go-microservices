package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dilshad-cp/go-microservices/details"
	"github.com/gorilla/mux"
)

/*
import (

	"fmt"
	"unsafe"

	geo "github.com/dilshad-cp/go-microservices/geometry"

	"rsc.io/quote"

)

	func rectProps(length, width float64) (area, perimeter float64) {
		area = length * width
		perimeter = 2 * (length + width)
		return
	}

	func main() {
		xa := 31
		name := "Dilshad"
		fmt.Println(quote.Go())
		fmt.Println(xa, name)
		fmt.Printf("Hi, %s Your are %d", name, xa)
		fmt.Printf("Type %T Length %d\n", name, unsafe.Sizeof(name))

		a, p := rectProps(1, 2)

		fmt.Printf("Area : %f, Perimeter: %f\n", a, p)

		// daysOfTheMonth := make(map[string]int)
		// daysOfTheMonth["jan"] = 31

		var daysOfTheMonth = map[string]int{"jan": 31, "feb": 28}

		fmt.Println(daysOfTheMonth)

		geo.Area(1, 2)
	}
*/
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Application running homepage!")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Application running")
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "hellow you are requested %s, %s\n", title, page)

}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")
	response := map[string]string{
		"status": "OK",
		"time":   time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching details")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}

	ip := details.GetIp()

	fmt.Printf("Hostname:%s, Ip: %s", hostname, ip)

	response := map[string]string{
		"hostnaem": hostname,
		"ip":       ip.String(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	/*port := ":80"
	http.HandleFunc("/", rootHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Printf("Webserver started on port %s", port)
	http.ListenAndServe(port, nil)*/
	port := ":80"
	r := mux.NewRouter()
	r.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "hellow you are requested %s\n", r.URL.Path)

	})
	r.HandleFunc("/books/{title}/page/{page}", booksHandler)
	r.HandleFunc("/healthcheck", healthHandler)
	r.HandleFunc("/details", detailsHandler)
	r.HandleFunc("/", rootHandler)
	// http.Handle("/", r)
	log.Printf("Webserver started on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
