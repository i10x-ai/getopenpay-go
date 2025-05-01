# InvoiceDiscountAmountsExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountAtom** | **int32** | Amount in units of atom (for usd that is cents, e.g. 1200 for $12) | 
**DiscountId** | **string** | Id of the discount that was applied to get this discount amount. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 

## Methods

### NewInvoiceDiscountAmountsExternal

`func NewInvoiceDiscountAmountsExternal(amountAtom int32, discountId string, ) *InvoiceDiscountAmountsExternal`

NewInvoiceDiscountAmountsExternal instantiates a new InvoiceDiscountAmountsExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDiscountAmountsExternalWithDefaults

`func NewInvoiceDiscountAmountsExternalWithDefaults() *InvoiceDiscountAmountsExternal`

NewInvoiceDiscountAmountsExternalWithDefaults instantiates a new InvoiceDiscountAmountsExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountAtom

`func (o *InvoiceDiscountAmountsExternal) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *InvoiceDiscountAmountsExternal) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *InvoiceDiscountAmountsExternal) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetDiscountId

`func (o *InvoiceDiscountAmountsExternal) GetDiscountId() string`

GetDiscountId returns the DiscountId field if non-nil, zero value otherwise.

### GetDiscountIdOk

`func (o *InvoiceDiscountAmountsExternal) GetDiscountIdOk() (*string, bool)`

GetDiscountIdOk returns a tuple with the DiscountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountId

`func (o *InvoiceDiscountAmountsExternal) SetDiscountId(v string)`

SetDiscountId sets DiscountId field to given value.


### GetObject

`func (o *InvoiceDiscountAmountsExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *InvoiceDiscountAmountsExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *InvoiceDiscountAmountsExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *InvoiceDiscountAmountsExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


