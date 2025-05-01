# CreditNoteItemExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountAtom** | **int32** | The integer amount representing the gross amount being credited for this credit_note_item. It is in atomic units (in USD this is cents). | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) | Three-letter ISO currency code, in lowercase. | 
**Id** | **string** | Unique Identifier of the credit_note_item. | 
**InvoiceItemId** | Pointer to **NullableString** |  | [optional] 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**Quantity** | Pointer to **int32** | Quantity of the product being credited. | [optional] [default to 1]
**Type** | [**CreditNoteLineType**](CreditNoteLineType.md) | The type of the credit note line item, one of invoice_line_item or custom_line_item. When the type is invoice_line_item there is an additional invoice_line_item property on the resource the value of which is the id of the credited line item on the invoice. | 
**UnitAmountAtom** | Pointer to **NullableInt32** |  | [optional] 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewCreditNoteItemExternal

`func NewCreditNoteItemExternal(amountAtom int32, createdAt time.Time, currency CurrencyEnum, id string, type_ CreditNoteLineType, updatedAt time.Time, ) *CreditNoteItemExternal`

NewCreditNoteItemExternal instantiates a new CreditNoteItemExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteItemExternalWithDefaults

`func NewCreditNoteItemExternalWithDefaults() *CreditNoteItemExternal`

NewCreditNoteItemExternalWithDefaults instantiates a new CreditNoteItemExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountAtom

`func (o *CreditNoteItemExternal) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *CreditNoteItemExternal) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *CreditNoteItemExternal) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetCreatedAt

`func (o *CreditNoteItemExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreditNoteItemExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreditNoteItemExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *CreditNoteItemExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditNoteItemExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditNoteItemExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetId

`func (o *CreditNoteItemExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditNoteItemExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditNoteItemExternal) SetId(v string)`

SetId sets Id field to given value.


### GetInvoiceItemId

`func (o *CreditNoteItemExternal) GetInvoiceItemId() string`

GetInvoiceItemId returns the InvoiceItemId field if non-nil, zero value otherwise.

### GetInvoiceItemIdOk

`func (o *CreditNoteItemExternal) GetInvoiceItemIdOk() (*string, bool)`

GetInvoiceItemIdOk returns a tuple with the InvoiceItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemId

`func (o *CreditNoteItemExternal) SetInvoiceItemId(v string)`

SetInvoiceItemId sets InvoiceItemId field to given value.

### HasInvoiceItemId

`func (o *CreditNoteItemExternal) HasInvoiceItemId() bool`

HasInvoiceItemId returns a boolean if a field has been set.

### SetInvoiceItemIdNil

`func (o *CreditNoteItemExternal) SetInvoiceItemIdNil(b bool)`

 SetInvoiceItemIdNil sets the value for InvoiceItemId to be an explicit nil

### UnsetInvoiceItemId
`func (o *CreditNoteItemExternal) UnsetInvoiceItemId()`

UnsetInvoiceItemId ensures that no value is present for InvoiceItemId, not even an explicit nil
### GetIsDeleted

`func (o *CreditNoteItemExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CreditNoteItemExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CreditNoteItemExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CreditNoteItemExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetObject

`func (o *CreditNoteItemExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreditNoteItemExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreditNoteItemExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *CreditNoteItemExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetQuantity

`func (o *CreditNoteItemExternal) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreditNoteItemExternal) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreditNoteItemExternal) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreditNoteItemExternal) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetType

`func (o *CreditNoteItemExternal) GetType() CreditNoteLineType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditNoteItemExternal) GetTypeOk() (*CreditNoteLineType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditNoteItemExternal) SetType(v CreditNoteLineType)`

SetType sets Type field to given value.


### GetUnitAmountAtom

`func (o *CreditNoteItemExternal) GetUnitAmountAtom() int32`

GetUnitAmountAtom returns the UnitAmountAtom field if non-nil, zero value otherwise.

### GetUnitAmountAtomOk

`func (o *CreditNoteItemExternal) GetUnitAmountAtomOk() (*int32, bool)`

GetUnitAmountAtomOk returns a tuple with the UnitAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountAtom

`func (o *CreditNoteItemExternal) SetUnitAmountAtom(v int32)`

SetUnitAmountAtom sets UnitAmountAtom field to given value.

### HasUnitAmountAtom

`func (o *CreditNoteItemExternal) HasUnitAmountAtom() bool`

HasUnitAmountAtom returns a boolean if a field has been set.

### SetUnitAmountAtomNil

`func (o *CreditNoteItemExternal) SetUnitAmountAtomNil(b bool)`

 SetUnitAmountAtomNil sets the value for UnitAmountAtom to be an explicit nil

### UnsetUnitAmountAtom
`func (o *CreditNoteItemExternal) UnsetUnitAmountAtom()`

UnsetUnitAmountAtom ensures that no value is present for UnitAmountAtom, not even an explicit nil
### GetUpdatedAt

`func (o *CreditNoteItemExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreditNoteItemExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreditNoteItemExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


