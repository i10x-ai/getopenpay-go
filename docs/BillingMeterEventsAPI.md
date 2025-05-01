# \BillingMeterEventsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingMeterEvent**](BillingMeterEventsAPI.md#CreateBillingMeterEvent) | **Post** /billing/meter-events/ | Create Billing Meter Event
[**ListBillingMeterEventSummaries**](BillingMeterEventsAPI.md#ListBillingMeterEventSummaries) | **Post** /billing/meter-events/list | List Billing Meter Event Summaries



## CreateBillingMeterEvent

> BillingMeterEventExternal CreateBillingMeterEvent(ctx).CreateBillingMeterEventRequest(createBillingMeterEventRequest).Execute()

Create Billing Meter Event



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
	createBillingMeterEventRequest := *openapiclient.NewCreateBillingMeterEventRequest("EventName_example", map[string]interface{}(123)) // CreateBillingMeterEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingMeterEventsAPI.CreateBillingMeterEvent(context.Background()).CreateBillingMeterEventRequest(createBillingMeterEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingMeterEventsAPI.CreateBillingMeterEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillingMeterEvent`: BillingMeterEventExternal
	fmt.Fprintf(os.Stdout, "Response from `BillingMeterEventsAPI.CreateBillingMeterEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillingMeterEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingMeterEventRequest** | [**CreateBillingMeterEventRequest**](CreateBillingMeterEventRequest.md) |  | 

### Return type

[**BillingMeterEventExternal**](BillingMeterEventExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBillingMeterEventSummaries

> ListResponseBillingMeterSummary ListBillingMeterEventSummaries(ctx).ListBillingMeterEventSummariesRequest(listBillingMeterEventSummariesRequest).Execute()

List Billing Meter Event Summaries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func main() {
	listBillingMeterEventSummariesRequest := *openapiclient.NewListBillingMeterEventSummariesRequest("CustomerId_example", time.Now(), "MeterId_example", time.Now()) // ListBillingMeterEventSummariesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingMeterEventsAPI.ListBillingMeterEventSummaries(context.Background()).ListBillingMeterEventSummariesRequest(listBillingMeterEventSummariesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingMeterEventsAPI.ListBillingMeterEventSummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBillingMeterEventSummaries`: ListResponseBillingMeterSummary
	fmt.Fprintf(os.Stdout, "Response from `BillingMeterEventsAPI.ListBillingMeterEventSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingMeterEventSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listBillingMeterEventSummariesRequest** | [**ListBillingMeterEventSummariesRequest**](ListBillingMeterEventSummariesRequest.md) |  | 

### Return type

[**ListResponseBillingMeterSummary**](ListResponseBillingMeterSummary.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

