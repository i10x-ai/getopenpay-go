# CreditNoteExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**CreditAmountAtom** | **int32** | The integer amount representing the amount to credit the customer’s balance, which will be automatically applied to their next invoice. It is in atomic units (in USD this is cents). | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) | Three-letter ISO currency code, in lowercase. | 
**Id** | **string** | Unique Identifier of the credit_note. | 
**InvoiceId** | **string** | Unique ID of the invoice. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Items** | [**[]CreditNoteItemExternal**](CreditNoteItemExternal.md) | List of credit_note_items. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**Reason** | [**NullableCreditNoteReason**](CreditNoteReason.md) |  | 
**RefundAmountAtom** | **int32** | The integer amount representing the amount to refund. If set, a refund will be created for the charge associated with the invoice. It is in atomic units (in USD this is cents). | 
**TotalAmountAtom** | **int32** | The integer amount representing the total amount of the credit note. It is in atomic units (in USD this is cents). | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewCreditNoteExternal

`func NewCreditNoteExternal(createdAt time.Time, creditAmountAtom int32, currency CurrencyEnum, id string, invoiceId string, items []CreditNoteItemExternal, reason NullableCreditNoteReason, refundAmountAtom int32, totalAmountAtom int32, updatedAt time.Time, ) *CreditNoteExternal`

NewCreditNoteExternal instantiates a new CreditNoteExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteExternalWithDefaults

`func NewCreditNoteExternalWithDefaults() *CreditNoteExternal`

NewCreditNoteExternalWithDefaults instantiates a new CreditNoteExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CreditNoteExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreditNoteExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreditNoteExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreditAmountAtom

`func (o *CreditNoteExternal) GetCreditAmountAtom() int32`

GetCreditAmountAtom returns the CreditAmountAtom field if non-nil, zero value otherwise.

### GetCreditAmountAtomOk

`func (o *CreditNoteExternal) GetCreditAmountAtomOk() (*int32, bool)`

GetCreditAmountAtomOk returns a tuple with the CreditAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountAtom

`func (o *CreditNoteExternal) SetCreditAmountAtom(v int32)`

SetCreditAmountAtom sets CreditAmountAtom field to given value.


### GetCurrency

`func (o *CreditNoteExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditNoteExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditNoteExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetId

`func (o *CreditNoteExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditNoteExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditNoteExternal) SetId(v string)`

SetId sets Id field to given value.


### GetInvoiceId

`func (o *CreditNoteExternal) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreditNoteExternal) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreditNoteExternal) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetIsDeleted

`func (o *CreditNoteExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CreditNoteExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CreditNoteExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CreditNoteExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetItems

`func (o *CreditNoteExternal) GetItems() []CreditNoteItemExternal`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreditNoteExternal) GetItemsOk() (*[]CreditNoteItemExternal, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreditNoteExternal) SetItems(v []CreditNoteItemExternal)`

SetItems sets Items field to given value.


### GetObject

`func (o *CreditNoteExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreditNoteExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreditNoteExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *CreditNoteExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetReason

`func (o *CreditNoteExternal) GetReason() CreditNoteReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreditNoteExternal) GetReasonOk() (*CreditNoteReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreditNoteExternal) SetReason(v CreditNoteReason)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *CreditNoteExternal) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *CreditNoteExternal) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRefundAmountAtom

`func (o *CreditNoteExternal) GetRefundAmountAtom() int32`

GetRefundAmountAtom returns the RefundAmountAtom field if non-nil, zero value otherwise.

### GetRefundAmountAtomOk

`func (o *CreditNoteExternal) GetRefundAmountAtomOk() (*int32, bool)`

GetRefundAmountAtomOk returns a tuple with the RefundAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountAtom

`func (o *CreditNoteExternal) SetRefundAmountAtom(v int32)`

SetRefundAmountAtom sets RefundAmountAtom field to given value.


### GetTotalAmountAtom

`func (o *CreditNoteExternal) GetTotalAmountAtom() int32`

GetTotalAmountAtom returns the TotalAmountAtom field if non-nil, zero value otherwise.

### GetTotalAmountAtomOk

`func (o *CreditNoteExternal) GetTotalAmountAtomOk() (*int32, bool)`

GetTotalAmountAtomOk returns a tuple with the TotalAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountAtom

`func (o *CreditNoteExternal) SetTotalAmountAtom(v int32)`

SetTotalAmountAtom sets TotalAmountAtom field to given value.


### GetUpdatedAt

`func (o *CreditNoteExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreditNoteExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreditNoteExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


