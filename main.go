package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/aws/aws-sdk-go/service/organizations/organizationsiface"
)

type Organizations struct {
    Client organizationsiface.OrganizationsAPI
}

func (s *Organizations) ListAccounts(input *organizations.ListAccountsInput) (*organizations.ListAccountsOutput, error) {
	restult, err := s.Client.ListAccounts(input)
	return restult, err
}


// This method has parameter client of type *Organization. On type Organization we have defined method ListAccounts. 
// Then ListAccount method will call through struct Organization's --> Client -->  organizationsiface.OrganizationsAPI --> ListAccounts (s.Client.organizationsiface.OrganizationsAPI.ListAccounts || s.Client.ListAccounts {TYPE PROMOTION}) 
// Finally API get triggred at network.
func WhatAreMyAccounts(client *Organizations) (*organizations.ListAccountsOutput, error) {
	input := &organizations.ListAccountsInput{
		MaxResults: aws.Int64(5),
		NextToken:  nil,
	}
    return client.ListAccounts(input)
}

func main() {
	// you need to pass svc in Organization
	// svc := organizations.New(nil) 
	client := Organizations{
		Client: &organizations.Organizations{}, // This type organizations.Organizations{} is implementing all methods of organizationsiface.OrganizationsAPI.
												// So this organizations.Organizations{} is also of type organizationsiface.OrganizationsAPI.
	}
	WhatAreMyAccounts(&client)
}

// Follow the below link to under stand more:-
// https://dev.to/danquack/mocking-the-aws-sdk-with-go-44hi#:~:text=Mocking%20a%20client%20library%20is,you%20are%20trying%20to%20mock. 