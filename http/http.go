package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/mnmtanish/go-graphiql"
	"placeFinderBack/app"
	"placeFinderBack/models"
	"log"
	"net/http"
)

func RegisterHttpHandlers(f *app.Factory){
	graphqlHandler := http.HandlerFunc(graphqlHandlerFunc)
	http.HandleFunc("/", graphiql.ServeGraphiQL)
	http.Handle("/graphql",addContext(graphqlHandler,f))
}

func StartServer(port string){
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	log.Println(fmt.Sprintf("Now server is running on port %s", port) )
}

func graphqlHandlerFunc(w http.ResponseWriter, r *http.Request) {
	opts := handler.NewRequestOptions(r)
	schema, extErr := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    models.Queries,
			Mutation: models.Mutations,
		},
	)
	if extErr != nil {
		fmt.Println("error trying to create schema: %v", extErr)
	}
	// execute graphql query
	params := graphql.Params{
		Schema:         schema, // defined in another file
		RequestString:  opts.Query,
		VariableValues: opts.Variables,
		OperationName:  opts.OperationName,
		Context:        r.Context(),
	}
	result := graphql.Do(params)
	// output JSON
	var buff []byte
	w.WriteHeader(http.StatusOK)
	buff, _ = json.Marshal(result)
	w.Write(buff)
}

func addContext(next http.Handler, f *app.Factory) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := models.Request{Factory: f}
		ctx := context.WithValue(r.Context(), 0, request)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
