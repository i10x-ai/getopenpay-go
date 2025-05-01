# CreateCreditNoteLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountAtom** | **int32** | The integer amount representing the gross amount being credited for this line item. | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) |  | 
**InvoiceItemId** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] [default to 1]
**Type** | [**CreditNoteLineType**](CreditNoteLineType.md) | The type of the credit note line item, one of invoice_line_item or custom_line_item. When the type is invoice_line_item there is an additional invoice_line_item property on the resource the value of which is the id of the credited line item on the invoice | 
**UnitAmountAtom** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateCreditNoteLine

`func NewCreateCreditNoteLine(amountAtom int32, currency CurrencyEnum, type_ CreditNoteLineType, ) *CreateCreditNoteLine`

NewCreateCreditNoteLine instantiates a new CreateCreditNoteLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCreditNoteLineWithDefaults

`func NewCreateCreditNoteLineWithDefaults() *CreateCreditNoteLine`

NewCreateCreditNoteLineWithDefaults instantiates a new CreateCreditNoteLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountAtom

`func (o *CreateCreditNoteLine) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *CreateCreditNoteLine) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *CreateCreditNoteLine) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetCurrency

`func (o *CreateCreditNoteLine) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCreditNoteLine) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCreditNoteLine) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetInvoiceItemId

`func (o *CreateCreditNoteLine) GetInvoiceItemId() string`

GetInvoiceItemId returns the InvoiceItemId field if non-nil, zero value otherwise.

### GetInvoiceItemIdOk

`func (o *CreateCreditNoteLine) GetInvoiceItemIdOk() (*string, bool)`

GetInvoiceItemIdOk returns a tuple with the InvoiceItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemId

`func (o *CreateCreditNoteLine) SetInvoiceItemId(v string)`

SetInvoiceItemId sets InvoiceItemId field to given value.

### HasInvoiceItemId

`func (o *CreateCreditNoteLine) HasInvoiceItemId() bool`

HasInvoiceItemId returns a boolean if a field has been set.

### SetInvoiceItemIdNil

`func (o *CreateCreditNoteLine) SetInvoiceItemIdNil(b bool)`

 SetInvoiceItemIdNil sets the value for InvoiceItemId to be an explicit nil

### UnsetInvoiceItemId
`func (o *CreateCreditNoteLine) UnsetInvoiceItemId()`

UnsetInvoiceItemId ensures that no value is present for InvoiceItemId, not even an explicit nil
### GetQuantity

`func (o *CreateCreditNoteLine) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateCreditNoteLine) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateCreditNoteLine) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreateCreditNoteLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetType

`func (o *CreateCreditNoteLine) GetType() CreditNoteLineType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCreditNoteLine) GetTypeOk() (*CreditNoteLineType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCreditNoteLine) SetType(v CreditNoteLineType)`

SetType sets Type field to given value.


### GetUnitAmountAtom

`func (o *CreateCreditNoteLine) GetUnitAmountAtom() int32`

GetUnitAmountAtom returns the UnitAmountAtom field if non-nil, zero value otherwise.

### GetUnitAmountAtomOk

`func (o *CreateCreditNoteLine) GetUnitAmountAtomOk() (*int32, bool)`

GetUnitAmountAtomOk returns a tuple with the UnitAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountAtom

`func (o *CreateCreditNoteLine) SetUnitAmountAtom(v int32)`

SetUnitAmountAtom sets UnitAmountAtom field to given value.

### HasUnitAmountAtom

`func (o *CreateCreditNoteLine) HasUnitAmountAtom() bool`

HasUnitAmountAtom returns a boolean if a field has been set.

### SetUnitAmountAtomNil

`func (o *CreateCreditNoteLine) SetUnitAmountAtomNil(b bool)`

 SetUnitAmountAtomNil sets the value for UnitAmountAtom to be an explicit nil

### UnsetUnitAmountAtom
`func (o *CreateCreditNoteLine) UnsetUnitAmountAtom()`

UnsetUnitAmountAtom ensures that no value is present for UnitAmountAtom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


