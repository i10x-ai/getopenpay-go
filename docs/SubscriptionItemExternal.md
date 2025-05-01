# SubscriptionItemExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddAtPeriodEnd** | **bool** | Whether or not this item will be added before next renewal | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) | Three-letter ISO currency code, in lowercase. | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**DeletedAt** | **NullableTime** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DropAtEnd** | **bool** | Whether or not this item will be dropped from subscription before next renewal | 
**Id** | **string** | Unique Identifier of the subscription_item. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**PendingAttachmentToSubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Price** | [**PriceExternal**](PriceExternal.md) | Price object associated with subscription_item | 
**PriceId** | **string** | Unique Identifier of the price. | 
**ProductId** | **string** | Unique Identifier of the product. | 
**Quantity** | **int32** | Quantity of the product selected for the subscription_item. | 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewSubscriptionItemExternal

`func NewSubscriptionItemExternal(addAtPeriodEnd bool, createdAt time.Time, currency CurrencyEnum, deletedAt NullableTime, dropAtEnd bool, id string, price PriceExternal, priceId string, productId string, quantity int32, updatedAt time.Time, ) *SubscriptionItemExternal`

NewSubscriptionItemExternal instantiates a new SubscriptionItemExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionItemExternalWithDefaults

`func NewSubscriptionItemExternalWithDefaults() *SubscriptionItemExternal`

NewSubscriptionItemExternalWithDefaults instantiates a new SubscriptionItemExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddAtPeriodEnd

`func (o *SubscriptionItemExternal) GetAddAtPeriodEnd() bool`

GetAddAtPeriodEnd returns the AddAtPeriodEnd field if non-nil, zero value otherwise.

### GetAddAtPeriodEndOk

`func (o *SubscriptionItemExternal) GetAddAtPeriodEndOk() (*bool, bool)`

GetAddAtPeriodEndOk returns a tuple with the AddAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAtPeriodEnd

`func (o *SubscriptionItemExternal) SetAddAtPeriodEnd(v bool)`

SetAddAtPeriodEnd sets AddAtPeriodEnd field to given value.


### GetCreatedAt

`func (o *SubscriptionItemExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionItemExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionItemExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *SubscriptionItemExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionItemExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionItemExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetCustomFields

`func (o *SubscriptionItemExternal) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SubscriptionItemExternal) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SubscriptionItemExternal) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SubscriptionItemExternal) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *SubscriptionItemExternal) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *SubscriptionItemExternal) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDeletedAt

`func (o *SubscriptionItemExternal) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SubscriptionItemExternal) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SubscriptionItemExternal) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *SubscriptionItemExternal) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *SubscriptionItemExternal) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetDescription

`func (o *SubscriptionItemExternal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionItemExternal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionItemExternal) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionItemExternal) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SubscriptionItemExternal) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SubscriptionItemExternal) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDropAtEnd

`func (o *SubscriptionItemExternal) GetDropAtEnd() bool`

GetDropAtEnd returns the DropAtEnd field if non-nil, zero value otherwise.

### GetDropAtEndOk

`func (o *SubscriptionItemExternal) GetDropAtEndOk() (*bool, bool)`

GetDropAtEndOk returns a tuple with the DropAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropAtEnd

`func (o *SubscriptionItemExternal) SetDropAtEnd(v bool)`

SetDropAtEnd sets DropAtEnd field to given value.


### GetId

`func (o *SubscriptionItemExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionItemExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionItemExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *SubscriptionItemExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SubscriptionItemExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SubscriptionItemExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *SubscriptionItemExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetObject

`func (o *SubscriptionItemExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionItemExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionItemExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *SubscriptionItemExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPendingAttachmentToSubscriptionId

`func (o *SubscriptionItemExternal) GetPendingAttachmentToSubscriptionId() string`

GetPendingAttachmentToSubscriptionId returns the PendingAttachmentToSubscriptionId field if non-nil, zero value otherwise.

### GetPendingAttachmentToSubscriptionIdOk

`func (o *SubscriptionItemExternal) GetPendingAttachmentToSubscriptionIdOk() (*string, bool)`

GetPendingAttachmentToSubscriptionIdOk returns a tuple with the PendingAttachmentToSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAttachmentToSubscriptionId

`func (o *SubscriptionItemExternal) SetPendingAttachmentToSubscriptionId(v string)`

SetPendingAttachmentToSubscriptionId sets PendingAttachmentToSubscriptionId field to given value.

### HasPendingAttachmentToSubscriptionId

`func (o *SubscriptionItemExternal) HasPendingAttachmentToSubscriptionId() bool`

HasPendingAttachmentToSubscriptionId returns a boolean if a field has been set.

### SetPendingAttachmentToSubscriptionIdNil

`func (o *SubscriptionItemExternal) SetPendingAttachmentToSubscriptionIdNil(b bool)`

 SetPendingAttachmentToSubscriptionIdNil sets the value for PendingAttachmentToSubscriptionId to be an explicit nil

### UnsetPendingAttachmentToSubscriptionId
`func (o *SubscriptionItemExternal) UnsetPendingAttachmentToSubscriptionId()`

UnsetPendingAttachmentToSubscriptionId ensures that no value is present for PendingAttachmentToSubscriptionId, not even an explicit nil
### GetPrice

`func (o *SubscriptionItemExternal) GetPrice() PriceExternal`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SubscriptionItemExternal) GetPriceOk() (*PriceExternal, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SubscriptionItemExternal) SetPrice(v PriceExternal)`

SetPrice sets Price field to given value.


### GetPriceId

`func (o *SubscriptionItemExternal) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *SubscriptionItemExternal) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *SubscriptionItemExternal) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetProductId

`func (o *SubscriptionItemExternal) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SubscriptionItemExternal) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SubscriptionItemExternal) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetQuantity

`func (o *SubscriptionItemExternal) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionItemExternal) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionItemExternal) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetSubscriptionId

`func (o *SubscriptionItemExternal) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionItemExternal) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionItemExternal) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SubscriptionItemExternal) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *SubscriptionItemExternal) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *SubscriptionItemExternal) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetUpdatedAt

`func (o *SubscriptionItemExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionItemExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionItemExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


