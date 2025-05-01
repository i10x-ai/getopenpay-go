# CheckoutSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**InvoiceUrls** | **[]string** |  | 
**SubscriptionIds** | **[]string** |  | 

## Methods

### NewCheckoutSuccessResponse

`func NewCheckoutSuccessResponse(customerId string, invoiceUrls []string, subscriptionIds []string, ) *CheckoutSuccessResponse`

NewCheckoutSuccessResponse instantiates a new CheckoutSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSuccessResponseWithDefaults

`func NewCheckoutSuccessResponseWithDefaults() *CheckoutSuccessResponse`

NewCheckoutSuccessResponseWithDefaults instantiates a new CheckoutSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CheckoutSuccessResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CheckoutSuccessResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CheckoutSuccessResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInvoiceUrls

`func (o *CheckoutSuccessResponse) GetInvoiceUrls() []string`

GetInvoiceUrls returns the InvoiceUrls field if non-nil, zero value otherwise.

### GetInvoiceUrlsOk

`func (o *CheckoutSuccessResponse) GetInvoiceUrlsOk() (*[]string, bool)`

GetInvoiceUrlsOk returns a tuple with the InvoiceUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceUrls

`func (o *CheckoutSuccessResponse) SetInvoiceUrls(v []string)`

SetInvoiceUrls sets InvoiceUrls field to given value.


### GetSubscriptionIds

`func (o *CheckoutSuccessResponse) GetSubscriptionIds() []string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *CheckoutSuccessResponse) GetSubscriptionIdsOk() (*[]string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *CheckoutSuccessResponse) SetSubscriptionIds(v []string)`

SetSubscriptionIds sets SubscriptionIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


