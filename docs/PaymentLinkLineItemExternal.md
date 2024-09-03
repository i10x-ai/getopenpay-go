# PaymentLinkLineItemExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the payment link line item. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_PAYMENT_LINK_LINE_ITEM]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**PaymentLinkId** | **string** | The ID of the payment link associated with this item | 
**PriceId** | **string** | Unique identifier of the price. | 
**Quantity** | **int32** | Quantity of the item. | 

## Methods

### NewPaymentLinkLineItemExternal

`func NewPaymentLinkLineItemExternal(id string, createdAt time.Time, updatedAt time.Time, paymentLinkId string, priceId string, quantity int32, ) *PaymentLinkLineItemExternal`

NewPaymentLinkLineItemExternal instantiates a new PaymentLinkLineItemExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentLinkLineItemExternalWithDefaults

`func NewPaymentLinkLineItemExternalWithDefaults() *PaymentLinkLineItemExternal`

NewPaymentLinkLineItemExternalWithDefaults instantiates a new PaymentLinkLineItemExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentLinkLineItemExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentLinkLineItemExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentLinkLineItemExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PaymentLinkLineItemExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentLinkLineItemExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentLinkLineItemExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *PaymentLinkLineItemExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentLinkLineItemExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentLinkLineItemExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentLinkLineItemExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PaymentLinkLineItemExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentLinkLineItemExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentLinkLineItemExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *PaymentLinkLineItemExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PaymentLinkLineItemExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PaymentLinkLineItemExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PaymentLinkLineItemExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetPaymentLinkId

`func (o *PaymentLinkLineItemExternal) GetPaymentLinkId() string`

GetPaymentLinkId returns the PaymentLinkId field if non-nil, zero value otherwise.

### GetPaymentLinkIdOk

`func (o *PaymentLinkLineItemExternal) GetPaymentLinkIdOk() (*string, bool)`

GetPaymentLinkIdOk returns a tuple with the PaymentLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLinkId

`func (o *PaymentLinkLineItemExternal) SetPaymentLinkId(v string)`

SetPaymentLinkId sets PaymentLinkId field to given value.


### GetPriceId

`func (o *PaymentLinkLineItemExternal) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *PaymentLinkLineItemExternal) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *PaymentLinkLineItemExternal) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetQuantity

`func (o *PaymentLinkLineItemExternal) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PaymentLinkLineItemExternal) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PaymentLinkLineItemExternal) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


