# AuthorizePaymentMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizeAmountAtom** | Pointer to **int32** | Amount to authorize in atoms (smallest currency unit). | [optional] [default to 2000]
**Currency** | Pointer to [**CurrencyEnum**](CurrencyEnum.md) | Currency for the authorization amount. | [optional] 
**PaymentMethodId** | **string** | Unique identifier of the payment method. | 

## Methods

### NewAuthorizePaymentMethodRequest

`func NewAuthorizePaymentMethodRequest(paymentMethodId string, ) *AuthorizePaymentMethodRequest`

NewAuthorizePaymentMethodRequest instantiates a new AuthorizePaymentMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizePaymentMethodRequestWithDefaults

`func NewAuthorizePaymentMethodRequestWithDefaults() *AuthorizePaymentMethodRequest`

NewAuthorizePaymentMethodRequestWithDefaults instantiates a new AuthorizePaymentMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizeAmountAtom

`func (o *AuthorizePaymentMethodRequest) GetAuthorizeAmountAtom() int32`

GetAuthorizeAmountAtom returns the AuthorizeAmountAtom field if non-nil, zero value otherwise.

### GetAuthorizeAmountAtomOk

`func (o *AuthorizePaymentMethodRequest) GetAuthorizeAmountAtomOk() (*int32, bool)`

GetAuthorizeAmountAtomOk returns a tuple with the AuthorizeAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeAmountAtom

`func (o *AuthorizePaymentMethodRequest) SetAuthorizeAmountAtom(v int32)`

SetAuthorizeAmountAtom sets AuthorizeAmountAtom field to given value.

### HasAuthorizeAmountAtom

`func (o *AuthorizePaymentMethodRequest) HasAuthorizeAmountAtom() bool`

HasAuthorizeAmountAtom returns a boolean if a field has been set.

### GetCurrency

`func (o *AuthorizePaymentMethodRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AuthorizePaymentMethodRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AuthorizePaymentMethodRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AuthorizePaymentMethodRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *AuthorizePaymentMethodRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *AuthorizePaymentMethodRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *AuthorizePaymentMethodRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


