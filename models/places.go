package models

import (
	"github.com/graphql-go/graphql"
	"placeFinderBack/places"
)

var placeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PlaceType",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(*places.Place).Name, nil
			},
		},
		"latitude": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(*places.Place).Latitude, nil
			},
		},
		"longitude": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(*places.Place).Longitude, nil
			},
		},
	},
})

var PlacesListQuery = &graphql.Field{
	Type:        graphql.NewNonNull(graphql.NewList(placeType)),
	Description: "List places",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		f := getFactory(p)
		return f.Places.List(), nil
	},
}

var PlacesCreateMutation = &graphql.Field{
	Type: graphql.ID,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"latitude": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"longitude": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		f := getFactory(p)
		name := p.Args["name"].(string)
		latitude := p.Args["latitude"].(string)
		longitude := p.Args["longitude"].(string)
		err := f.Places.Create(&places.Place{
			Name:      name,
			Latitude:  latitude,
			Longitude: longitude,
		})
		return success, err
	},
}