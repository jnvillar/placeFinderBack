package main

import (
	"placeFinderBack/app"
	http2 "placeFinderBack/http"
	"placeFinderBack/places"
)

const port = "8080"

func main() {
	f := app.NewFactory()
	f.Places.Create(&places.Place{Name:"prueba", Latitude: "prueba", Longitude:"prueba"})

	http2.RegisterHttpHandlers(f)
	http2.StartServer(port)
}