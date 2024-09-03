# PromoRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstTimeTransaction** | Pointer to **NullableBool** |  | [optional] 
**MinimumAmountAtom** | Pointer to **NullableInt32** |  | [optional] 
**MinimumAmountCurrency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) |  | [optional] 

## Methods

### NewPromoRestrictions

`func NewPromoRestrictions() *PromoRestrictions`

NewPromoRestrictions instantiates a new PromoRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoRestrictionsWithDefaults

`func NewPromoRestrictionsWithDefaults() *PromoRestrictions`

NewPromoRestrictionsWithDefaults instantiates a new PromoRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstTimeTransaction

`func (o *PromoRestrictions) GetFirstTimeTransaction() bool`

GetFirstTimeTransaction returns the FirstTimeTransaction field if non-nil, zero value otherwise.

### GetFirstTimeTransactionOk

`func (o *PromoRestrictions) GetFirstTimeTransactionOk() (*bool, bool)`

GetFirstTimeTransactionOk returns a tuple with the FirstTimeTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTimeTransaction

`func (o *PromoRestrictions) SetFirstTimeTransaction(v bool)`

SetFirstTimeTransaction sets FirstTimeTransaction field to given value.

### HasFirstTimeTransaction

`func (o *PromoRestrictions) HasFirstTimeTransaction() bool`

HasFirstTimeTransaction returns a boolean if a field has been set.

### SetFirstTimeTransactionNil

`func (o *PromoRestrictions) SetFirstTimeTransactionNil(b bool)`

 SetFirstTimeTransactionNil sets the value for FirstTimeTransaction to be an explicit nil

### UnsetFirstTimeTransaction
`func (o *PromoRestrictions) UnsetFirstTimeTransaction()`

UnsetFirstTimeTransaction ensures that no value is present for FirstTimeTransaction, not even an explicit nil
### GetMinimumAmountAtom

`func (o *PromoRestrictions) GetMinimumAmountAtom() int32`

GetMinimumAmountAtom returns the MinimumAmountAtom field if non-nil, zero value otherwise.

### GetMinimumAmountAtomOk

`func (o *PromoRestrictions) GetMinimumAmountAtomOk() (*int32, bool)`

GetMinimumAmountAtomOk returns a tuple with the MinimumAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmountAtom

`func (o *PromoRestrictions) SetMinimumAmountAtom(v int32)`

SetMinimumAmountAtom sets MinimumAmountAtom field to given value.

### HasMinimumAmountAtom

`func (o *PromoRestrictions) HasMinimumAmountAtom() bool`

HasMinimumAmountAtom returns a boolean if a field has been set.

### SetMinimumAmountAtomNil

`func (o *PromoRestrictions) SetMinimumAmountAtomNil(b bool)`

 SetMinimumAmountAtomNil sets the value for MinimumAmountAtom to be an explicit nil

### UnsetMinimumAmountAtom
`func (o *PromoRestrictions) UnsetMinimumAmountAtom()`

UnsetMinimumAmountAtom ensures that no value is present for MinimumAmountAtom, not even an explicit nil
### GetMinimumAmountCurrency

`func (o *PromoRestrictions) GetMinimumAmountCurrency() CurrencyEnum`

GetMinimumAmountCurrency returns the MinimumAmountCurrency field if non-nil, zero value otherwise.

### GetMinimumAmountCurrencyOk

`func (o *PromoRestrictions) GetMinimumAmountCurrencyOk() (*CurrencyEnum, bool)`

GetMinimumAmountCurrencyOk returns a tuple with the MinimumAmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmountCurrency

`func (o *PromoRestrictions) SetMinimumAmountCurrency(v CurrencyEnum)`

SetMinimumAmountCurrency sets MinimumAmountCurrency field to given value.

### HasMinimumAmountCurrency

`func (o *PromoRestrictions) HasMinimumAmountCurrency() bool`

HasMinimumAmountCurrency returns a boolean if a field has been set.

### SetMinimumAmountCurrencyNil

`func (o *PromoRestrictions) SetMinimumAmountCurrencyNil(b bool)`

 SetMinimumAmountCurrencyNil sets the value for MinimumAmountCurrency to be an explicit nil

### UnsetMinimumAmountCurrency
`func (o *PromoRestrictions) UnsetMinimumAmountCurrency()`

UnsetMinimumAmountCurrency ensures that no value is present for MinimumAmountCurrency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


