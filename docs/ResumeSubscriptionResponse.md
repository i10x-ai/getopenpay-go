# ResumeSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | Pointer to [**NullableInvoiceExternal**](InvoiceExternal.md) |  | [optional] 
**Subscription** | [**SubscriptionExternal**](SubscriptionExternal.md) | The subscription being resumed. | 

## Methods

### NewResumeSubscriptionResponse

`func NewResumeSubscriptionResponse(subscription SubscriptionExternal, ) *ResumeSubscriptionResponse`

NewResumeSubscriptionResponse instantiates a new ResumeSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeSubscriptionResponseWithDefaults

`func NewResumeSubscriptionResponseWithDefaults() *ResumeSubscriptionResponse`

NewResumeSubscriptionResponseWithDefaults instantiates a new ResumeSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoice

`func (o *ResumeSubscriptionResponse) GetInvoice() InvoiceExternal`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *ResumeSubscriptionResponse) GetInvoiceOk() (*InvoiceExternal, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *ResumeSubscriptionResponse) SetInvoice(v InvoiceExternal)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *ResumeSubscriptionResponse) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### SetInvoiceNil

`func (o *ResumeSubscriptionResponse) SetInvoiceNil(b bool)`

 SetInvoiceNil sets the value for Invoice to be an explicit nil

### UnsetInvoice
`func (o *ResumeSubscriptionResponse) UnsetInvoice()`

UnsetInvoice ensures that no value is present for Invoice, not even an explicit nil
### GetSubscription

`func (o *ResumeSubscriptionResponse) GetSubscription() SubscriptionExternal`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ResumeSubscriptionResponse) GetSubscriptionOk() (*SubscriptionExternal, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ResumeSubscriptionResponse) SetSubscription(v SubscriptionExternal)`

SetSubscription sets Subscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


