/*
OpenPay API

Testing PaymentLinksAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package getopenpay

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func Test_getopenpay_PaymentLinksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PaymentLinksAPIService CreatePaymentLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PaymentLinksAPI.CreatePaymentLink(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentLinksAPIService GetPaymentLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var plinkId string

		resp, httpRes, err := apiClient.PaymentLinksAPI.GetPaymentLink(context.Background(), plinkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentLinksAPIService ListPaymentLinks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PaymentLinksAPI.ListPaymentLinks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentLinksAPIService OpenPaymentLinkPagePublic", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var secureToken string

		resp, httpRes, err := apiClient.PaymentLinksAPI.OpenPaymentLinkPagePublic(context.Background(), secureToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
