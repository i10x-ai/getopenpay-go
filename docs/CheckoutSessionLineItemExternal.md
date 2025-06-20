# CheckoutSessionLineItemExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountSubtotalAtom** | **int32** |  | 
**AmountTotalAtom** | **int32** |  | 
**BillingInterval** | [**NullableCalendarIntervalEnum**](CalendarIntervalEnum.md) |  | 
**BillingIntervalCount** | **NullableInt32** |  | 
**CheckoutSessionId** | **NullableString** |  | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) |  | 
**CustomFields** | **map[string]interface{}** |  | 
**Description** | **NullableString** |  | 
**DescriptionDetailed** | **NullableString** |  | 
**Id** | **string** |  | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**PriceId** | **string** |  | 
**ProductId** | **string** |  | 
**Quantity** | **int32** |  | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewCheckoutSessionLineItemExternal

`func NewCheckoutSessionLineItemExternal(amountSubtotalAtom int32, amountTotalAtom int32, billingInterval NullableCalendarIntervalEnum, billingIntervalCount NullableInt32, checkoutSessionId NullableString, createdAt time.Time, currency CurrencyEnum, customFields map[string]interface{}, description NullableString, descriptionDetailed NullableString, id string, priceId string, productId string, quantity int32, updatedAt time.Time, ) *CheckoutSessionLineItemExternal`

NewCheckoutSessionLineItemExternal instantiates a new CheckoutSessionLineItemExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionLineItemExternalWithDefaults

`func NewCheckoutSessionLineItemExternalWithDefaults() *CheckoutSessionLineItemExternal`

NewCheckoutSessionLineItemExternalWithDefaults instantiates a new CheckoutSessionLineItemExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountSubtotalAtom

`func (o *CheckoutSessionLineItemExternal) GetAmountSubtotalAtom() int32`

GetAmountSubtotalAtom returns the AmountSubtotalAtom field if non-nil, zero value otherwise.

### GetAmountSubtotalAtomOk

`func (o *CheckoutSessionLineItemExternal) GetAmountSubtotalAtomOk() (*int32, bool)`

GetAmountSubtotalAtomOk returns a tuple with the AmountSubtotalAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSubtotalAtom

`func (o *CheckoutSessionLineItemExternal) SetAmountSubtotalAtom(v int32)`

SetAmountSubtotalAtom sets AmountSubtotalAtom field to given value.


### GetAmountTotalAtom

`func (o *CheckoutSessionLineItemExternal) GetAmountTotalAtom() int32`

GetAmountTotalAtom returns the AmountTotalAtom field if non-nil, zero value otherwise.

### GetAmountTotalAtomOk

`func (o *CheckoutSessionLineItemExternal) GetAmountTotalAtomOk() (*int32, bool)`

GetAmountTotalAtomOk returns a tuple with the AmountTotalAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTotalAtom

`func (o *CheckoutSessionLineItemExternal) SetAmountTotalAtom(v int32)`

SetAmountTotalAtom sets AmountTotalAtom field to given value.


### GetBillingInterval

`func (o *CheckoutSessionLineItemExternal) GetBillingInterval() CalendarIntervalEnum`

GetBillingInterval returns the BillingInterval field if non-nil, zero value otherwise.

### GetBillingIntervalOk

`func (o *CheckoutSessionLineItemExternal) GetBillingIntervalOk() (*CalendarIntervalEnum, bool)`

GetBillingIntervalOk returns a tuple with the BillingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInterval

`func (o *CheckoutSessionLineItemExternal) SetBillingInterval(v CalendarIntervalEnum)`

SetBillingInterval sets BillingInterval field to given value.


### SetBillingIntervalNil

`func (o *CheckoutSessionLineItemExternal) SetBillingIntervalNil(b bool)`

 SetBillingIntervalNil sets the value for BillingInterval to be an explicit nil

### UnsetBillingInterval
`func (o *CheckoutSessionLineItemExternal) UnsetBillingInterval()`

UnsetBillingInterval ensures that no value is present for BillingInterval, not even an explicit nil
### GetBillingIntervalCount

`func (o *CheckoutSessionLineItemExternal) GetBillingIntervalCount() int32`

GetBillingIntervalCount returns the BillingIntervalCount field if non-nil, zero value otherwise.

### GetBillingIntervalCountOk

`func (o *CheckoutSessionLineItemExternal) GetBillingIntervalCountOk() (*int32, bool)`

GetBillingIntervalCountOk returns a tuple with the BillingIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingIntervalCount

`func (o *CheckoutSessionLineItemExternal) SetBillingIntervalCount(v int32)`

SetBillingIntervalCount sets BillingIntervalCount field to given value.


### SetBillingIntervalCountNil

`func (o *CheckoutSessionLineItemExternal) SetBillingIntervalCountNil(b bool)`

 SetBillingIntervalCountNil sets the value for BillingIntervalCount to be an explicit nil

### UnsetBillingIntervalCount
`func (o *CheckoutSessionLineItemExternal) UnsetBillingIntervalCount()`

UnsetBillingIntervalCount ensures that no value is present for BillingIntervalCount, not even an explicit nil
### GetCheckoutSessionId

`func (o *CheckoutSessionLineItemExternal) GetCheckoutSessionId() string`

GetCheckoutSessionId returns the CheckoutSessionId field if non-nil, zero value otherwise.

### GetCheckoutSessionIdOk

`func (o *CheckoutSessionLineItemExternal) GetCheckoutSessionIdOk() (*string, bool)`

GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutSessionId

`func (o *CheckoutSessionLineItemExternal) SetCheckoutSessionId(v string)`

SetCheckoutSessionId sets CheckoutSessionId field to given value.


### SetCheckoutSessionIdNil

`func (o *CheckoutSessionLineItemExternal) SetCheckoutSessionIdNil(b bool)`

 SetCheckoutSessionIdNil sets the value for CheckoutSessionId to be an explicit nil

### UnsetCheckoutSessionId
`func (o *CheckoutSessionLineItemExternal) UnsetCheckoutSessionId()`

UnsetCheckoutSessionId ensures that no value is present for CheckoutSessionId, not even an explicit nil
### GetCreatedAt

`func (o *CheckoutSessionLineItemExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckoutSessionLineItemExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckoutSessionLineItemExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *CheckoutSessionLineItemExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CheckoutSessionLineItemExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CheckoutSessionLineItemExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetCustomFields

`func (o *CheckoutSessionLineItemExternal) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CheckoutSessionLineItemExternal) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CheckoutSessionLineItemExternal) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.


### SetCustomFieldsNil

`func (o *CheckoutSessionLineItemExternal) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CheckoutSessionLineItemExternal) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDescription

`func (o *CheckoutSessionLineItemExternal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckoutSessionLineItemExternal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckoutSessionLineItemExternal) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *CheckoutSessionLineItemExternal) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CheckoutSessionLineItemExternal) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDescriptionDetailed

`func (o *CheckoutSessionLineItemExternal) GetDescriptionDetailed() string`

GetDescriptionDetailed returns the DescriptionDetailed field if non-nil, zero value otherwise.

### GetDescriptionDetailedOk

`func (o *CheckoutSessionLineItemExternal) GetDescriptionDetailedOk() (*string, bool)`

GetDescriptionDetailedOk returns a tuple with the DescriptionDetailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionDetailed

`func (o *CheckoutSessionLineItemExternal) SetDescriptionDetailed(v string)`

SetDescriptionDetailed sets DescriptionDetailed field to given value.


### SetDescriptionDetailedNil

`func (o *CheckoutSessionLineItemExternal) SetDescriptionDetailedNil(b bool)`

 SetDescriptionDetailedNil sets the value for DescriptionDetailed to be an explicit nil

### UnsetDescriptionDetailed
`func (o *CheckoutSessionLineItemExternal) UnsetDescriptionDetailed()`

UnsetDescriptionDetailed ensures that no value is present for DescriptionDetailed, not even an explicit nil
### GetId

`func (o *CheckoutSessionLineItemExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutSessionLineItemExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutSessionLineItemExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *CheckoutSessionLineItemExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CheckoutSessionLineItemExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CheckoutSessionLineItemExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CheckoutSessionLineItemExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetObject

`func (o *CheckoutSessionLineItemExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CheckoutSessionLineItemExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CheckoutSessionLineItemExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *CheckoutSessionLineItemExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPriceId

`func (o *CheckoutSessionLineItemExternal) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *CheckoutSessionLineItemExternal) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *CheckoutSessionLineItemExternal) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetProductId

`func (o *CheckoutSessionLineItemExternal) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CheckoutSessionLineItemExternal) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CheckoutSessionLineItemExternal) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetQuantity

`func (o *CheckoutSessionLineItemExternal) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CheckoutSessionLineItemExternal) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CheckoutSessionLineItemExternal) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUpdatedAt

`func (o *CheckoutSessionLineItemExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CheckoutSessionLineItemExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CheckoutSessionLineItemExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


