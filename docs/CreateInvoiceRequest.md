# CreateInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | Unique identifier of the subscription. The created invoice will only include pending invoice items for that subscription. The subscription’s billing cycle and regular subscription events won’t be affected. | 
**PaymentMethodId** | Pointer to **NullableString** |  | [optional] 
**CollectionMethod** | Pointer to [**NullableCollectionMethodEnum**](CollectionMethodEnum.md) |  | [optional] 
**Description** | Pointer to **string** | Description for newly created invoice | [optional] [default to "Manual creation of invoice"]
**Discounts** | Pointer to [**NullableDiscounts**](Discounts.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateInvoiceRequest

`func NewCreateInvoiceRequest(subscriptionId string, ) *CreateInvoiceRequest`

NewCreateInvoiceRequest instantiates a new CreateInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvoiceRequestWithDefaults

`func NewCreateInvoiceRequestWithDefaults() *CreateInvoiceRequest`

NewCreateInvoiceRequestWithDefaults instantiates a new CreateInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *CreateInvoiceRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateInvoiceRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateInvoiceRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetPaymentMethodId

`func (o *CreateInvoiceRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *CreateInvoiceRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *CreateInvoiceRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *CreateInvoiceRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### SetPaymentMethodIdNil

`func (o *CreateInvoiceRequest) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *CreateInvoiceRequest) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetCollectionMethod

`func (o *CreateInvoiceRequest) GetCollectionMethod() CollectionMethodEnum`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *CreateInvoiceRequest) GetCollectionMethodOk() (*CollectionMethodEnum, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *CreateInvoiceRequest) SetCollectionMethod(v CollectionMethodEnum)`

SetCollectionMethod sets CollectionMethod field to given value.

### HasCollectionMethod

`func (o *CreateInvoiceRequest) HasCollectionMethod() bool`

HasCollectionMethod returns a boolean if a field has been set.

### SetCollectionMethodNil

`func (o *CreateInvoiceRequest) SetCollectionMethodNil(b bool)`

 SetCollectionMethodNil sets the value for CollectionMethod to be an explicit nil

### UnsetCollectionMethod
`func (o *CreateInvoiceRequest) UnsetCollectionMethod()`

UnsetCollectionMethod ensures that no value is present for CollectionMethod, not even an explicit nil
### GetDescription

`func (o *CreateInvoiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInvoiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInvoiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateInvoiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscounts

`func (o *CreateInvoiceRequest) GetDiscounts() Discounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *CreateInvoiceRequest) GetDiscountsOk() (*Discounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *CreateInvoiceRequest) SetDiscounts(v Discounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *CreateInvoiceRequest) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### SetDiscountsNil

`func (o *CreateInvoiceRequest) SetDiscountsNil(b bool)`

 SetDiscountsNil sets the value for Discounts to be an explicit nil

### UnsetDiscounts
`func (o *CreateInvoiceRequest) UnsetDiscounts()`

UnsetDiscounts ensures that no value is present for Discounts, not even an explicit nil
### GetCustomFields

`func (o *CreateInvoiceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateInvoiceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateInvoiceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateInvoiceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CreateInvoiceRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CreateInvoiceRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


