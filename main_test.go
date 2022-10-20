package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/aws/aws-sdk-go/service/organizations/organizationsiface"
	"github.com/stretchr/testify/assert"
)

type MockedOrganizations struct {
    organizationsiface.OrganizationsAPI
}

func (m *MockedOrganizations) ListAccounts(in *organizations.ListAccountsInput) (*organizations.ListAccountsOutput, error) {
    return &organizations.ListAccountsOutput{
        Accounts: []*organizations.Account{
            {
                Arn:   aws.String(""),
                Email: aws.String("test1@example.com"),
                Id:    aws.String("234567890"),
                Name:  aws.String("test-1"),
            },
            {
                Arn:   aws.String(""),
                Email: aws.String("test2@example.com"),
                Id:    aws.String("123456789"),
                Name:  aws.String("test-2"),
            },
        },
    }, nil
}

func TestListAccounts(t *testing.T) {
	// This is main file's Organization Structure's OBJECT.
	// Type MockedOrganizations{} is also implementing the organizationsiface.OrganizationsAPI interface.
	// So MockOrganizations is also of type organizationsiface.OrganizationsAPI.
	// And client is of Type organizationsiface.OrganizationsAPI. SO set Client = &MockedOrganizations{}
	// Now main_test.go's ListAccounts method is defined on type MockedOrganizations.
	// so when we call ListAccount by Organizations below object. 
	// Below Organizations object Which has type MockedOrganizations will call its own main_test.go's ListAccounts method. So we get above ListAccounts method data by over-riding
    test := Organizations{
        Client: &MockedOrganizations{},
    }
    resp, err := WhatAreMyAccounts(&test)
    assert.Equal(t, len(resp.Accounts), 2)
    assert.NoError(t, err)
}