# ChargeExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountAtom** | **int32** | Charge amount without any fees, in smallest currency unit. | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) | Currency code, e.g., USD. | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomerId** | **string** | ID of the customer associated with this charge. | 
**Disputed** | **bool** | Indicates whether the charge is disputed. | 
**FailureCode** | **NullableString** |  | 
**FailureMessage** | **NullableString** |  | 
**Id** | **string** | Unique Identifier of the charge. | 
**InvoiceId** | **NullableString** |  | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**PaymentIntentId** | **string** | ID of the payment intent associated with this charge. | 
**PaymentIntentMappingId** | **NullableString** |  | 
**PaymentMethod** | Pointer to [**NullablePaymentMethodExternal**](PaymentMethodExternal.md) |  | [optional] 
**PaymentMethodId** | **string** | ID of the payment method used for this charge. | 
**PaymentProcessorId** | **NullableString** |  | 
**PaymentProcessorName** | **NullableString** |  | 
**ProcessorMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ProviderTypeFeeAmountAtom** | **int32** | Fee amount charged due to the payment provider type, in smallest currency unit. | 
**Refunded** | **bool** | Indicates whether the charge has been refunded. | 
**RefundedAmountAtom** | **int32** | Amount refunded, in smallest currency unit. | 
**Status** | [**ChargeStatusEnum**](ChargeStatusEnum.md) | Status of the charge. | 
**SubscriptionIds** | Pointer to **[]string** | List of subscription ids for which the charge is created. | [optional] 
**Subscriptions** | Pointer to [**[]SubscriptionExternal**](SubscriptionExternal.md) | List of subscriptions for which the charge is created. | [optional] 
**TotalChargeAmountAtom** | **int32** | Total charge amount including fees, in smallest currency unit. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewChargeExternal

`func NewChargeExternal(amountAtom int32, createdAt time.Time, currency CurrencyEnum, customerId string, disputed bool, failureCode NullableString, failureMessage NullableString, id string, invoiceId NullableString, paymentIntentId string, paymentIntentMappingId NullableString, paymentMethodId string, paymentProcessorId NullableString, paymentProcessorName NullableString, providerTypeFeeAmountAtom int32, refunded bool, refundedAmountAtom int32, status ChargeStatusEnum, totalChargeAmountAtom int32, updatedAt time.Time, ) *ChargeExternal`

NewChargeExternal instantiates a new ChargeExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeExternalWithDefaults

`func NewChargeExternalWithDefaults() *ChargeExternal`

NewChargeExternalWithDefaults instantiates a new ChargeExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetPaymentIntentMappingId

`func (o *ChargeExternal) GetPaymentIntentMappingId() string`

GetPaymentIntentMappingId returns the PaymentIntentMappingId field if non-nil, zero value otherwise.

### GetPaymentIntentMappingIdOk

`func (o *ChargeExternal) GetPaymentIntentMappingIdOk() (*string, bool)`

GetPaymentIntentMappingIdOk returns a tuple with the PaymentIntentMappingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentMappingId

`func (o *ChargeExternal) SetPaymentIntentMappingId(v string)`

SetPaymentIntentMappingId sets PaymentIntentMappingId field to given value.


### SetPaymentIntentMappingIdNil

`func (o *ChargeExternal) SetPaymentIntentMappingIdNil(b bool)`

 SetPaymentIntentMappingIdNil sets the value for PaymentIntentMappingId to be an explicit nil

### UnsetPaymentIntentMappingId
`func (o *ChargeExternal) UnsetPaymentIntentMappingId()`

UnsetPaymentIntentMappingId ensures that no value is present for PaymentIntentMappingId, not even an explicit nil
### GetPaymentMethod

`func (o *ChargeExternal) GetPaymentMethod() PaymentMethodExternal`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ChargeExternal) GetPaymentMethodOk() (*PaymentMethodExternal, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ChargeExternal) SetPaymentMethod(v PaymentMethodExternal)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *ChargeExternal) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *ChargeExternal) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *ChargeExternal) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
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


### GetPaymentProcessorId

`func (o *ChargeExternal) GetPaymentProcessorId() string`

GetPaymentProcessorId returns the PaymentProcessorId field if non-nil, zero value otherwise.

### GetPaymentProcessorIdOk

`func (o *ChargeExternal) GetPaymentProcessorIdOk() (*string, bool)`

GetPaymentProcessorIdOk returns a tuple with the PaymentProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProcessorId

`func (o *ChargeExternal) SetPaymentProcessorId(v string)`

SetPaymentProcessorId sets PaymentProcessorId field to given value.


### SetPaymentProcessorIdNil

`func (o *ChargeExternal) SetPaymentProcessorIdNil(b bool)`

 SetPaymentProcessorIdNil sets the value for PaymentProcessorId to be an explicit nil

### UnsetPaymentProcessorId
`func (o *ChargeExternal) UnsetPaymentProcessorId()`

UnsetPaymentProcessorId ensures that no value is present for PaymentProcessorId, not even an explicit nil
### GetPaymentProcessorName

`func (o *ChargeExternal) GetPaymentProcessorName() string`

GetPaymentProcessorName returns the PaymentProcessorName field if non-nil, zero value otherwise.

### GetPaymentProcessorNameOk

`func (o *ChargeExternal) GetPaymentProcessorNameOk() (*string, bool)`

GetPaymentProcessorNameOk returns a tuple with the PaymentProcessorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProcessorName

`func (o *ChargeExternal) SetPaymentProcessorName(v string)`

SetPaymentProcessorName sets PaymentProcessorName field to given value.


### SetPaymentProcessorNameNil

`func (o *ChargeExternal) SetPaymentProcessorNameNil(b bool)`

 SetPaymentProcessorNameNil sets the value for PaymentProcessorName to be an explicit nil

### UnsetPaymentProcessorName
`func (o *ChargeExternal) UnsetPaymentProcessorName()`

UnsetPaymentProcessorName ensures that no value is present for PaymentProcessorName, not even an explicit nil
### GetProcessorMetadata

`func (o *ChargeExternal) GetProcessorMetadata() map[string]interface{}`

GetProcessorMetadata returns the ProcessorMetadata field if non-nil, zero value otherwise.

### GetProcessorMetadataOk

`func (o *ChargeExternal) GetProcessorMetadataOk() (*map[string]interface{}, bool)`

GetProcessorMetadataOk returns a tuple with the ProcessorMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorMetadata

`func (o *ChargeExternal) SetProcessorMetadata(v map[string]interface{})`

SetProcessorMetadata sets ProcessorMetadata field to given value.

### HasProcessorMetadata

`func (o *ChargeExternal) HasProcessorMetadata() bool`

HasProcessorMetadata returns a boolean if a field has been set.

### SetProcessorMetadataNil

`func (o *ChargeExternal) SetProcessorMetadataNil(b bool)`

 SetProcessorMetadataNil sets the value for ProcessorMetadata to be an explicit nil

### UnsetProcessorMetadata
`func (o *ChargeExternal) UnsetProcessorMetadata()`

UnsetProcessorMetadata ensures that no value is present for ProcessorMetadata, not even an explicit nil
### GetProviderTypeFeeAmountAtom

`func (o *ChargeExternal) GetProviderTypeFeeAmountAtom() int32`

GetProviderTypeFeeAmountAtom returns the ProviderTypeFeeAmountAtom field if non-nil, zero value otherwise.

### GetProviderTypeFeeAmountAtomOk

`func (o *ChargeExternal) GetProviderTypeFeeAmountAtomOk() (*int32, bool)`

GetProviderTypeFeeAmountAtomOk returns a tuple with the ProviderTypeFeeAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeFeeAmountAtom

`func (o *ChargeExternal) SetProviderTypeFeeAmountAtom(v int32)`

SetProviderTypeFeeAmountAtom sets ProviderTypeFeeAmountAtom field to given value.


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


### GetSubscriptionIds

`func (o *ChargeExternal) GetSubscriptionIds() []string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *ChargeExternal) GetSubscriptionIdsOk() (*[]string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *ChargeExternal) SetSubscriptionIds(v []string)`

SetSubscriptionIds sets SubscriptionIds field to given value.

### HasSubscriptionIds

`func (o *ChargeExternal) HasSubscriptionIds() bool`

HasSubscriptionIds returns a boolean if a field has been set.

### GetSubscriptions

`func (o *ChargeExternal) GetSubscriptions() []SubscriptionExternal`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ChargeExternal) GetSubscriptionsOk() (*[]SubscriptionExternal, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ChargeExternal) SetSubscriptions(v []SubscriptionExternal)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *ChargeExternal) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetTotalChargeAmountAtom

`func (o *ChargeExternal) GetTotalChargeAmountAtom() int32`

GetTotalChargeAmountAtom returns the TotalChargeAmountAtom field if non-nil, zero value otherwise.

### GetTotalChargeAmountAtomOk

`func (o *ChargeExternal) GetTotalChargeAmountAtomOk() (*int32, bool)`

GetTotalChargeAmountAtomOk returns a tuple with the TotalChargeAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChargeAmountAtom

`func (o *ChargeExternal) SetTotalChargeAmountAtom(v int32)`

SetTotalChargeAmountAtom sets TotalChargeAmountAtom field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


