package models

import (
	"github.com/graphql-go/graphql"
	"placeFinderBack/app"
)

type Request struct {
	Factory *app.Factory
}

const success = "ok"

var Queries = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Query",
		Fields: queryList,
	})

var Mutations = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: mutationsList,
	})

var queryList = graphql.Fields{
	"listPlaces": PlacesListQuery,
}

var mutationsList = graphql.Fields{
	"createPlace": PlacesCreateMutation,
}

func getFactory(params graphql.ResolveParams) *app.Factory{
	return params.Context.Value(0).(Request).Factory
}
