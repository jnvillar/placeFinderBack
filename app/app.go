package app

import "placeFinderBack/places"

type Factory struct {
	Places *places.Factory
}

func NewFactory() *Factory {
	f := &Factory{}
	f.Places = places.NewFactory()
	return f
}
