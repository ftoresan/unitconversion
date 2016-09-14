package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ftoresan/unitconversion/distance"
)

func milesToKilometers(w http.ResponseWriter, r *http.Request) {
	var sMiles = r.URL.Query().Get("miles")
	var miles, e = strconv.ParseFloat(sMiles, 32)
	var km float32
	if e == nil {
		km = distance.MilesToKilometers(float32(miles))
	} else {
		km = distance.MilesToKilometers(0)
	}
	fmt.Fprint(w, km)
}

func handleConversions() {
	http.HandleFunc("/milesToKilometers", milesToKilometers)

	http.ListenAndServe(":8080", nil)
}

func main() {
	handleConversions()
}
