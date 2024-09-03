# InvoiceItemDiscountAmountsExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_INVOICE_ITEM_DISCOUNT]
**DiscountId** | **string** | Id of the discount that was applied to get this discount amount. | 
**AmountAtom** | **int32** | The amount_atom, of the discount. | 

## Methods

### NewInvoiceItemDiscountAmountsExternal

`func NewInvoiceItemDiscountAmountsExternal(discountId string, amountAtom int32, ) *InvoiceItemDiscountAmountsExternal`

NewInvoiceItemDiscountAmountsExternal instantiates a new InvoiceItemDiscountAmountsExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceItemDiscountAmountsExternalWithDefaults

`func NewInvoiceItemDiscountAmountsExternalWithDefaults() *InvoiceItemDiscountAmountsExternal`

NewInvoiceItemDiscountAmountsExternalWithDefaults instantiates a new InvoiceItemDiscountAmountsExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *InvoiceItemDiscountAmountsExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *InvoiceItemDiscountAmountsExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *InvoiceItemDiscountAmountsExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *InvoiceItemDiscountAmountsExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetDiscountId

`func (o *InvoiceItemDiscountAmountsExternal) GetDiscountId() string`

GetDiscountId returns the DiscountId field if non-nil, zero value otherwise.

### GetDiscountIdOk

`func (o *InvoiceItemDiscountAmountsExternal) GetDiscountIdOk() (*string, bool)`

GetDiscountIdOk returns a tuple with the DiscountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountId

`func (o *InvoiceItemDiscountAmountsExternal) SetDiscountId(v string)`

SetDiscountId sets DiscountId field to given value.


### GetAmountAtom

`func (o *InvoiceItemDiscountAmountsExternal) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *InvoiceItemDiscountAmountsExternal) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *InvoiceItemDiscountAmountsExternal) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


