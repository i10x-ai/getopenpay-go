# \BillingMetersAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingMeter**](BillingMetersAPI.md#CreateBillingMeter) | **Post** /billing/meters/ | Create Billing Meter
[**DeactivateBillingMeter**](BillingMetersAPI.md#DeactivateBillingMeter) | **Post** /billing/meters/{meter_id}/deactivate | Deactivate Billing Meter
[**GetBillingMeter**](BillingMetersAPI.md#GetBillingMeter) | **Get** /billing/meters/{meter_id} | Get Billing Meter
[**ListBillingMeters**](BillingMetersAPI.md#ListBillingMeters) | **Post** /billing/meters/list | List Billing Meters
[**ReactivateBillingMeter**](BillingMetersAPI.md#ReactivateBillingMeter) | **Post** /billing/meters/{meter_id}/reactivate | Reactivate Billing Meter



## CreateBillingMeter

> BillingMeterExternal CreateBillingMeter(ctx).CreateBillingMeterRequest(createBillingMeterRequest).Execute()

Create Billing Meter



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
	createBillingMeterRequest := *openapiclient.NewCreateBillingMeterRequest("DisplayName_example", "EventName_example") // CreateBillingMeterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingMetersAPI.CreateBillingMeter(context.Background()).CreateBillingMeterRequest(createBillingMeterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingMetersAPI.CreateBillingMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillingMeter`: BillingMeterExternal
	fmt.Fprintf(os.Stdout, "Response from `BillingMetersAPI.CreateBillingMeter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillingMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingMeterRequest** | [**CreateBillingMeterRequest**](CreateBillingMeterRequest.md) |  | 

### Return type

[**BillingMeterExternal**](BillingMeterExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateBillingMeter

> BillingMeterExternal DeactivateBillingMeter(ctx, meterId).Execute()

Deactivate Billing Meter



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
	meterId := "meterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingMetersAPI.DeactivateBillingMeter(context.Background(), meterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingMetersAPI.DeactivateBillingMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateBillingMeter`: BillingMeterExternal
	fmt.Fprintf(os.Stdout, "Response from `BillingMetersAPI.DeactivateBillingMeter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateBillingMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingMeterExternal**](BillingMeterExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingMeter

> BillingMeterExternal GetBillingMeter(ctx, meterId).Execute()

Get Billing Meter



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
	meterId := "meterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingMetersAPI.GetBillingMeter(context.Background(), meterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingMetersAPI.GetBillingMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingMeter`: BillingMeterExternal
	fmt.Fprintf(os.Stdout, "Response from `BillingMetersAPI.GetBillingMeter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingMeterExternal**](BillingMeterExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBillingMeters

> ListResponseBillingMeterExternal ListBillingMeters(ctx).BillingMeterQueryParams(billingMeterQueryParams).Execute()

List Billing Meters



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
	billingMeterQueryParams := *openapiclient.NewBillingMeterQueryParams(false) // BillingMeterQueryParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingMetersAPI.ListBillingMeters(context.Background()).BillingMeterQueryParams(billingMeterQueryParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingMetersAPI.ListBillingMeters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBillingMeters`: ListResponseBillingMeterExternal
	fmt.Fprintf(os.Stdout, "Response from `BillingMetersAPI.ListBillingMeters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingMetersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingMeterQueryParams** | [**BillingMeterQueryParams**](BillingMeterQueryParams.md) |  | 

### Return type

[**ListResponseBillingMeterExternal**](ListResponseBillingMeterExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactivateBillingMeter

> BillingMeterExternal ReactivateBillingMeter(ctx, meterId).Execute()

Reactivate Billing Meter



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
	meterId := "meterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingMetersAPI.ReactivateBillingMeter(context.Background(), meterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingMetersAPI.ReactivateBillingMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReactivateBillingMeter`: BillingMeterExternal
	fmt.Fprintf(os.Stdout, "Response from `BillingMetersAPI.ReactivateBillingMeter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactivateBillingMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingMeterExternal**](BillingMeterExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

