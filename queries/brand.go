package queries

import (
	"fmt"
	"go_graphql/graphqlclient"
	"go_graphql/models"
)

func GetAllBrands(client *graphqlclient.Client) ([]models.Brand, error) {
	query := `
    query getAllBrands {
        getAllBrands {
            response {
                status
                message
                code
            }
            data {
                brandUuid
                brandName
            }
        }
    }`
	var response struct {
		GetAllBrands struct {
			Response struct {
				Status  bool   `json:"status"`
				Message string `json:"message"`
				Code    int    `json:"code"`
			} `json:"response"`
			Data []models.Brand `json:"data"`
		} `json:"getAllBrands"`
	}

	err := client.DoQuery(query, &response)
	if err != nil {
		return nil, err
	}
	if !response.GetAllBrands.Response.Status {
		return nil, fmt.Errorf("failed to fetch brands: %s", response.GetAllBrands.Response.Message)
	}

	return response.GetAllBrands.Data, nil
}
