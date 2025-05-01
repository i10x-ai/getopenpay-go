# CreateSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**[]SubscriptionExternal**](SubscriptionExternal.md) | List of subscriptions created. | 
**Invoices** | [**[]InvoiceExternal**](InvoiceExternal.md) | List of invoices created | 
**ProcessorsUsed** | **[]string** | List of successful processor IDs used for creating the subscriptions | 

## Methods

### NewCreateSubscriptionResponse

`func NewCreateSubscriptionResponse(created []SubscriptionExternal, invoices []InvoiceExternal, processorsUsed []string, ) *CreateSubscriptionResponse`

NewCreateSubscriptionResponse instantiates a new CreateSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionResponseWithDefaults

`func NewCreateSubscriptionResponseWithDefaults() *CreateSubscriptionResponse`

NewCreateSubscriptionResponseWithDefaults instantiates a new CreateSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CreateSubscriptionResponse) GetCreated() []SubscriptionExternal`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateSubscriptionResponse) GetCreatedOk() (*[]SubscriptionExternal, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateSubscriptionResponse) SetCreated(v []SubscriptionExternal)`

SetCreated sets Created field to given value.


### GetInvoices

`func (o *CreateSubscriptionResponse) GetInvoices() []InvoiceExternal`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *CreateSubscriptionResponse) GetInvoicesOk() (*[]InvoiceExternal, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *CreateSubscriptionResponse) SetInvoices(v []InvoiceExternal)`

SetInvoices sets Invoices field to given value.


### GetProcessorsUsed

`func (o *CreateSubscriptionResponse) GetProcessorsUsed() []string`

GetProcessorsUsed returns the ProcessorsUsed field if non-nil, zero value otherwise.

### GetProcessorsUsedOk

`func (o *CreateSubscriptionResponse) GetProcessorsUsedOk() (*[]string, bool)`

GetProcessorsUsedOk returns a tuple with the ProcessorsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorsUsed

`func (o *CreateSubscriptionResponse) SetProcessorsUsed(v []string)`

SetProcessorsUsed sets ProcessorsUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


