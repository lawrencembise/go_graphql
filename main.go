package main

import (
	"fmt"
	"go_graphql/graphqlclient"
	"go_graphql/queries"
	"log"
)

func main() {
	client := graphqlclient.NewClient("http://localhost:8000/graphql")

	brands, err := queries.GetAllBrands(client)
	if err != nil {
		log.Fatal("Error fetching brands:", err)
	}

	for _, brand := range brands {
		fmt.Printf("Brand: %s, UUID: %s\n", brand.BrandName, brand.BrandUuid)
	}
}
