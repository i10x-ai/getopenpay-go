# CustomerBalanceExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountAtom** | **int32** |  | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) |  | 

## Methods

### NewCustomerBalanceExternal

`func NewCustomerBalanceExternal(amountAtom int32, currency CurrencyEnum, ) *CustomerBalanceExternal`

NewCustomerBalanceExternal instantiates a new CustomerBalanceExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerBalanceExternalWithDefaults

`func NewCustomerBalanceExternalWithDefaults() *CustomerBalanceExternal`

NewCustomerBalanceExternalWithDefaults instantiates a new CustomerBalanceExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountAtom

`func (o *CustomerBalanceExternal) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *CustomerBalanceExternal) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *CustomerBalanceExternal) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetCurrency

`func (o *CustomerBalanceExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerBalanceExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerBalanceExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


