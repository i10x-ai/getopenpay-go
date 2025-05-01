# CreateCustomerBalanceTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountAtom** | **int32** | The amount of the transaction in atomic units (in USD this is cents). A negative value is a credit for the customer’s balance, and a positive value is a debit to the customer’s balance. | 
**Currency** | Pointer to [**CurrencyEnum**](CurrencyEnum.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateCustomerBalanceTransactionRequest

`func NewCreateCustomerBalanceTransactionRequest(amountAtom int32, ) *CreateCustomerBalanceTransactionRequest`

NewCreateCustomerBalanceTransactionRequest instantiates a new CreateCustomerBalanceTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomerBalanceTransactionRequestWithDefaults

`func NewCreateCustomerBalanceTransactionRequestWithDefaults() *CreateCustomerBalanceTransactionRequest`

NewCreateCustomerBalanceTransactionRequestWithDefaults instantiates a new CreateCustomerBalanceTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountAtom

`func (o *CreateCustomerBalanceTransactionRequest) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *CreateCustomerBalanceTransactionRequest) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *CreateCustomerBalanceTransactionRequest) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetCurrency

`func (o *CreateCustomerBalanceTransactionRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCustomerBalanceTransactionRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCustomerBalanceTransactionRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateCustomerBalanceTransactionRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *CreateCustomerBalanceTransactionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCustomerBalanceTransactionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCustomerBalanceTransactionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCustomerBalanceTransactionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateCustomerBalanceTransactionRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateCustomerBalanceTransactionRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


