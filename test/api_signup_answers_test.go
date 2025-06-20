/*
OpenPay API

Testing SignupAnswersAPIService

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

func Test_getopenpay_SignupAnswersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SignupAnswersAPIService SignupQuestionnaireStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SignupAnswersAPI.SignupQuestionnaireStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
