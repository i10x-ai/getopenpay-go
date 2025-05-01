# \BillingMeterEventAdjustmentsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingMeterEventAdjustments**](BillingMeterEventAdjustmentsAPI.md#CreateBillingMeterEventAdjustments) | **Post** /billing/meter-event-adjustments/ | Create Billing Meter Event Adjustments



## CreateBillingMeterEventAdjustments

> BillingMeterEventAdjustmentExternal CreateBillingMeterEventAdjustments(ctx).CreateBillingMeterEventAdjustmentRequest(createBillingMeterEventAdjustmentRequest).Execute()

Create Billing Meter Event Adjustments



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
	createBillingMeterEventAdjustmentRequest := *openapiclient.NewCreateBillingMeterEventAdjustmentRequest("CancelIdentifier_example", "EventName_example") // CreateBillingMeterEventAdjustmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingMeterEventAdjustmentsAPI.CreateBillingMeterEventAdjustments(context.Background()).CreateBillingMeterEventAdjustmentRequest(createBillingMeterEventAdjustmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingMeterEventAdjustmentsAPI.CreateBillingMeterEventAdjustments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillingMeterEventAdjustments`: BillingMeterEventAdjustmentExternal
	fmt.Fprintf(os.Stdout, "Response from `BillingMeterEventAdjustmentsAPI.CreateBillingMeterEventAdjustments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillingMeterEventAdjustmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingMeterEventAdjustmentRequest** | [**CreateBillingMeterEventAdjustmentRequest**](CreateBillingMeterEventAdjustmentRequest.md) |  | 

### Return type

[**BillingMeterEventAdjustmentExternal**](BillingMeterEventAdjustmentExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

