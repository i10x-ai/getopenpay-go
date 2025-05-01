# RefundExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountAtom** | **int32** | amount_atom that you want to refund. | 
**AttemptErrorMessage** | Pointer to **NullableString** |  | [optional] 
**ChargeId** | **string** | Unique Identifier of the charge. | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Currency** | Pointer to [**CurrencyEnum**](CurrencyEnum.md) | Three-letter ISO currency code, in lowercase. | [optional] 
**Id** | **string** | Unique Identifier of the refund. | 
**InvoiceId** | **string** | Invoice id to which the refund is attached. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**PaymentIntentId** | **string** | Unique Identifier of the payment_intent. | 
**Reason** | [**RefundReasonEnum**](RefundReasonEnum.md) | Reason of the refund. | 
**Status** | [**RefundStatusEnum**](RefundStatusEnum.md) | Status of the refund. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewRefundExternal

`func NewRefundExternal(amountAtom int32, chargeId string, createdAt time.Time, id string, invoiceId string, paymentIntentId string, reason RefundReasonEnum, status RefundStatusEnum, updatedAt time.Time, ) *RefundExternal`

NewRefundExternal instantiates a new RefundExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundExternalWithDefaults

`func NewRefundExternalWithDefaults() *RefundExternal`

NewRefundExternalWithDefaults instantiates a new RefundExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountAtom

`func (o *RefundExternal) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *RefundExternal) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *RefundExternal) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetAttemptErrorMessage

`func (o *RefundExternal) GetAttemptErrorMessage() string`

GetAttemptErrorMessage returns the AttemptErrorMessage field if non-nil, zero value otherwise.

### GetAttemptErrorMessageOk

`func (o *RefundExternal) GetAttemptErrorMessageOk() (*string, bool)`

GetAttemptErrorMessageOk returns a tuple with the AttemptErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptErrorMessage

`func (o *RefundExternal) SetAttemptErrorMessage(v string)`

SetAttemptErrorMessage sets AttemptErrorMessage field to given value.

### HasAttemptErrorMessage

`func (o *RefundExternal) HasAttemptErrorMessage() bool`

HasAttemptErrorMessage returns a boolean if a field has been set.

### SetAttemptErrorMessageNil

`func (o *RefundExternal) SetAttemptErrorMessageNil(b bool)`

 SetAttemptErrorMessageNil sets the value for AttemptErrorMessage to be an explicit nil

### UnsetAttemptErrorMessage
`func (o *RefundExternal) UnsetAttemptErrorMessage()`

UnsetAttemptErrorMessage ensures that no value is present for AttemptErrorMessage, not even an explicit nil
### GetChargeId

`func (o *RefundExternal) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *RefundExternal) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *RefundExternal) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.


### GetCreatedAt

`func (o *RefundExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RefundExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RefundExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *RefundExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RefundExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RefundExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *RefundExternal) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *RefundExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RefundExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RefundExternal) SetId(v string)`

SetId sets Id field to given value.


### GetInvoiceId

`func (o *RefundExternal) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *RefundExternal) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *RefundExternal) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetIsDeleted

`func (o *RefundExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *RefundExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *RefundExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *RefundExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetObject

`func (o *RefundExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RefundExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RefundExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *RefundExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPaymentIntentId

`func (o *RefundExternal) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *RefundExternal) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *RefundExternal) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.


### GetReason

`func (o *RefundExternal) GetReason() RefundReasonEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RefundExternal) GetReasonOk() (*RefundReasonEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RefundExternal) SetReason(v RefundReasonEnum)`

SetReason sets Reason field to given value.


### GetStatus

`func (o *RefundExternal) GetStatus() RefundStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RefundExternal) GetStatusOk() (*RefundStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RefundExternal) SetStatus(v RefundStatusEnum)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *RefundExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RefundExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RefundExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


