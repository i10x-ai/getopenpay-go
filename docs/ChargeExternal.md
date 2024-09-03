# ChargeExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Identifier of the charge. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_CHARGE]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**AmountAtom** | **int32** | Amount charged, in smallest currency unit. | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) |  | 
**CustomerId** | **string** | Unique Identifier of the customer associated with this charge. | 
**PaymentMethodId** | **string** | Unique Identifier of the payment method used for this charge. | 
**PaymentIntentId** | **string** | Unique Identifier of the payment intent associated with this charge. | 
**InvoiceId** | **NullableString** |  | 
**Status** | [**ChargeStatusEnum**](ChargeStatusEnum.md) |  | 
**Refunded** | **bool** | Indicates whether the charge has been refunded. | 
**RefundedAmountAtom** | **int32** | Amount refunded, in smallest currency unit. | 
**Disputed** | **bool** | Indicates whether the charge is disputed. | 
**FailureCode** | **NullableString** |  | 
**FailureMessage** | **NullableString** |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewChargeExternal

`func NewChargeExternal(id string, createdAt time.Time, updatedAt time.Time, amountAtom int32, currency CurrencyEnum, customerId string, paymentMethodId string, paymentIntentId string, invoiceId NullableString, status ChargeStatusEnum, refunded bool, refundedAmountAtom int32, disputed bool, failureCode NullableString, failureMessage NullableString, ) *ChargeExternal`

NewChargeExternal instantiates a new ChargeExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeExternalWithDefaults

`func NewChargeExternalWithDefaults() *ChargeExternal`

NewChargeExternalWithDefaults instantiates a new ChargeExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChargeExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ChargeExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChargeExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChargeExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *ChargeExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChargeExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChargeExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChargeExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ChargeExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChargeExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChargeExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *ChargeExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ChargeExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ChargeExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ChargeExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetAmountAtom

`func (o *ChargeExternal) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *ChargeExternal) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *ChargeExternal) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetCurrency

`func (o *ChargeExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ChargeExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ChargeExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *ChargeExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ChargeExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ChargeExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPaymentMethodId

`func (o *ChargeExternal) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *ChargeExternal) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *ChargeExternal) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### GetPaymentIntentId

`func (o *ChargeExternal) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *ChargeExternal) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *ChargeExternal) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.


### GetInvoiceId

`func (o *ChargeExternal) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ChargeExternal) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ChargeExternal) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *ChargeExternal) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *ChargeExternal) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetStatus

`func (o *ChargeExternal) GetStatus() ChargeStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargeExternal) GetStatusOk() (*ChargeStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargeExternal) SetStatus(v ChargeStatusEnum)`

SetStatus sets Status field to given value.


### GetRefunded

`func (o *ChargeExternal) GetRefunded() bool`

GetRefunded returns the Refunded field if non-nil, zero value otherwise.

### GetRefundedOk

`func (o *ChargeExternal) GetRefundedOk() (*bool, bool)`

GetRefundedOk returns a tuple with the Refunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunded

`func (o *ChargeExternal) SetRefunded(v bool)`

SetRefunded sets Refunded field to given value.


### GetRefundedAmountAtom

`func (o *ChargeExternal) GetRefundedAmountAtom() int32`

GetRefundedAmountAtom returns the RefundedAmountAtom field if non-nil, zero value otherwise.

### GetRefundedAmountAtomOk

`func (o *ChargeExternal) GetRefundedAmountAtomOk() (*int32, bool)`

GetRefundedAmountAtomOk returns a tuple with the RefundedAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmountAtom

`func (o *ChargeExternal) SetRefundedAmountAtom(v int32)`

SetRefundedAmountAtom sets RefundedAmountAtom field to given value.


### GetDisputed

`func (o *ChargeExternal) GetDisputed() bool`

GetDisputed returns the Disputed field if non-nil, zero value otherwise.

### GetDisputedOk

`func (o *ChargeExternal) GetDisputedOk() (*bool, bool)`

GetDisputedOk returns a tuple with the Disputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputed

`func (o *ChargeExternal) SetDisputed(v bool)`

SetDisputed sets Disputed field to given value.


### GetFailureCode

`func (o *ChargeExternal) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *ChargeExternal) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *ChargeExternal) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.


### SetFailureCodeNil

`func (o *ChargeExternal) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *ChargeExternal) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetFailureMessage

`func (o *ChargeExternal) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *ChargeExternal) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *ChargeExternal) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.


### SetFailureMessageNil

`func (o *ChargeExternal) SetFailureMessageNil(b bool)`

 SetFailureMessageNil sets the value for FailureMessage to be an explicit nil

### UnsetFailureMessage
`func (o *ChargeExternal) UnsetFailureMessage()`

UnsetFailureMessage ensures that no value is present for FailureMessage, not even an explicit nil
### GetCustomFields

`func (o *ChargeExternal) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ChargeExternal) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ChargeExternal) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ChargeExternal) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ChargeExternal) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ChargeExternal) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


