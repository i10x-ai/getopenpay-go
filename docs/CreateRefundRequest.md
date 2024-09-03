# CreateRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeId** | Pointer to **NullableString** |  | [optional] 
**PaymentIntentId** | Pointer to **NullableString** |  | [optional] 
**AmountAtom** | Pointer to **NullableInt32** |  | [optional] 
**Currency** | Pointer to [**CurrencyEnum**](CurrencyEnum.md) |  | [optional] [default to CURRENCYENUM_USD]
**Reason** | Pointer to [**RefundReasonEnum**](RefundReasonEnum.md) |  | [optional] [default to REFUNDREASONENUM_MANUAL]
**Description** | Pointer to **string** | Description of the refund. | [optional] [default to "Manually triggered refund by admin."]

## Methods

### NewCreateRefundRequest

`func NewCreateRefundRequest() *CreateRefundRequest`

NewCreateRefundRequest instantiates a new CreateRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRefundRequestWithDefaults

`func NewCreateRefundRequestWithDefaults() *CreateRefundRequest`

NewCreateRefundRequestWithDefaults instantiates a new CreateRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeId

`func (o *CreateRefundRequest) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *CreateRefundRequest) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *CreateRefundRequest) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.

### HasChargeId

`func (o *CreateRefundRequest) HasChargeId() bool`

HasChargeId returns a boolean if a field has been set.

### SetChargeIdNil

`func (o *CreateRefundRequest) SetChargeIdNil(b bool)`

 SetChargeIdNil sets the value for ChargeId to be an explicit nil

### UnsetChargeId
`func (o *CreateRefundRequest) UnsetChargeId()`

UnsetChargeId ensures that no value is present for ChargeId, not even an explicit nil
### GetPaymentIntentId

`func (o *CreateRefundRequest) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *CreateRefundRequest) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *CreateRefundRequest) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.

### HasPaymentIntentId

`func (o *CreateRefundRequest) HasPaymentIntentId() bool`

HasPaymentIntentId returns a boolean if a field has been set.

### SetPaymentIntentIdNil

`func (o *CreateRefundRequest) SetPaymentIntentIdNil(b bool)`

 SetPaymentIntentIdNil sets the value for PaymentIntentId to be an explicit nil

### UnsetPaymentIntentId
`func (o *CreateRefundRequest) UnsetPaymentIntentId()`

UnsetPaymentIntentId ensures that no value is present for PaymentIntentId, not even an explicit nil
### GetAmountAtom

`func (o *CreateRefundRequest) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *CreateRefundRequest) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *CreateRefundRequest) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.

### HasAmountAtom

`func (o *CreateRefundRequest) HasAmountAtom() bool`

HasAmountAtom returns a boolean if a field has been set.

### SetAmountAtomNil

`func (o *CreateRefundRequest) SetAmountAtomNil(b bool)`

 SetAmountAtomNil sets the value for AmountAtom to be an explicit nil

### UnsetAmountAtom
`func (o *CreateRefundRequest) UnsetAmountAtom()`

UnsetAmountAtom ensures that no value is present for AmountAtom, not even an explicit nil
### GetCurrency

`func (o *CreateRefundRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateRefundRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateRefundRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateRefundRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetReason

`func (o *CreateRefundRequest) GetReason() RefundReasonEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateRefundRequest) GetReasonOk() (*RefundReasonEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateRefundRequest) SetReason(v RefundReasonEnum)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateRefundRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDescription

`func (o *CreateRefundRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRefundRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRefundRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRefundRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


