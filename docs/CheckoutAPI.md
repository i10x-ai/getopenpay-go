# \CheckoutAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCheckoutSession**](CheckoutAPI.md#CreateCheckoutSession) | **Post** /checkout/sessions | Create Checkout Session
[**GetCheckoutSession**](CheckoutAPI.md#GetCheckoutSession) | **Get** /checkout/sessions/{session_id} | Get Checkout Session
[**ListCheckoutSessions**](CheckoutAPI.md#ListCheckoutSessions) | **Post** /checkout/list | List Checkout Sessions
[**ValidateCheckoutAttempt**](CheckoutAPI.md#ValidateCheckoutAttempt) | **Get** /checkout/checkout_attempts/validate/{checkout_attempt_id} | Validate Checkout Attempt
[**VerifyCheckoutSessionPaymentMethod**](CheckoutAPI.md#VerifyCheckoutSessionPaymentMethod) | **Get** /checkout/sessions/{checkout_secure_token}/verify_payment_method/{payment_method_id} | Verify Checkout Session Payment Method



## CreateCheckoutSession

> CheckoutSessionExternal CreateCheckoutSession(ctx).CreateCheckoutSessionRequest(createCheckoutSessionRequest).Execute()

Create Checkout Session

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func main() {
	createCheckoutSessionRequest := *openapiclient.NewCreateCheckoutSessionRequest(openapiclient.CheckoutMode("payment")) // CreateCheckoutSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.CreateCheckoutSession(context.Background()).CreateCheckoutSessionRequest(createCheckoutSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.CreateCheckoutSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCheckoutSession`: CheckoutSessionExternal
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.CreateCheckoutSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCheckoutSessionRequest** | [**CreateCheckoutSessionRequest**](CreateCheckoutSessionRequest.md) |  | 

### Return type

[**CheckoutSessionExternal**](CheckoutSessionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckoutSession

> CheckoutSessionExternal GetCheckoutSession(ctx, sessionId).Execute()

Get Checkout Session

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func main() {
	sessionId := "cs_dev_abc123" // string | Unique identifier of the checkout session.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.GetCheckoutSession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.GetCheckoutSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCheckoutSession`: CheckoutSessionExternal
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.GetCheckoutSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Unique identifier of the checkout session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutSessionExternal**](CheckoutSessionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCheckoutSessions

> ListResponseCheckoutSessionExternal ListCheckoutSessions(ctx).CheckoutSessionQueryParams(checkoutSessionQueryParams).Execute()

List Checkout Sessions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func main() {
	checkoutSessionQueryParams := *openapiclient.NewCheckoutSessionQueryParams() // CheckoutSessionQueryParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.ListCheckoutSessions(context.Background()).CheckoutSessionQueryParams(checkoutSessionQueryParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.ListCheckoutSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCheckoutSessions`: ListResponseCheckoutSessionExternal
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.ListCheckoutSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCheckoutSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkoutSessionQueryParams** | [**CheckoutSessionQueryParams**](CheckoutSessionQueryParams.md) |  | 

### Return type

[**ListResponseCheckoutSessionExternal**](ListResponseCheckoutSessionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCheckoutAttempt

> CheckoutSuccessResponse ValidateCheckoutAttempt(ctx, checkoutAttemptId).Execute()

Validate Checkout Attempt

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func main() {
	checkoutAttemptId := "catt_dev_abc123" // string | The checkout attempt ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.ValidateCheckoutAttempt(context.Background(), checkoutAttemptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.ValidateCheckoutAttempt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateCheckoutAttempt`: CheckoutSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.ValidateCheckoutAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutAttemptId** | **string** | The checkout attempt ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateCheckoutAttemptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutSuccessResponse**](CheckoutSuccessResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyCheckoutSessionPaymentMethod

> CheckoutSessionExternal VerifyCheckoutSessionPaymentMethod(ctx, checkoutSecureToken, paymentMethodId).Execute()

Verify Checkout Session Payment Method



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func main() {
	checkoutSecureToken := "00000000-aaaa-bbbb-cccc-dddddddddddd" // string | The checkout session's secure token
	paymentMethodId := "cc_dev_abc123" // string | The payment method ID to verify

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.VerifyCheckoutSessionPaymentMethod(context.Background(), checkoutSecureToken, paymentMethodId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.VerifyCheckoutSessionPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyCheckoutSessionPaymentMethod`: CheckoutSessionExternal
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.VerifyCheckoutSessionPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutSecureToken** | **string** | The checkout session&#39;s secure token | 
**paymentMethodId** | **string** | The payment method ID to verify | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyCheckoutSessionPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CheckoutSessionExternal**](CheckoutSessionExternal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

