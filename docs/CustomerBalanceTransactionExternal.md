# CustomerBalanceTransactionExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_CUSTOMER_BALANCE_TRANSACTION]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**AmountAtom** | **int32** |  | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) |  | 
**EndingBalanceAmountAtom** | **int32** |  | 
**Type** | [**CustomerBalanceTransactionType**](CustomerBalanceTransactionType.md) |  | 
**InvoiceId** | **NullableString** |  | 
**CreditNoteId** | **NullableString** |  | 
**Description** | **NullableString** |  | 

## Methods

### NewCustomerBalanceTransactionExternal

`func NewCustomerBalanceTransactionExternal(id string, createdAt time.Time, updatedAt time.Time, amountAtom int32, currency CurrencyEnum, endingBalanceAmountAtom int32, type_ CustomerBalanceTransactionType, invoiceId NullableString, creditNoteId NullableString, description NullableString, ) *CustomerBalanceTransactionExternal`

NewCustomerBalanceTransactionExternal instantiates a new CustomerBalanceTransactionExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerBalanceTransactionExternalWithDefaults

`func NewCustomerBalanceTransactionExternalWithDefaults() *CustomerBalanceTransactionExternal`

NewCustomerBalanceTransactionExternalWithDefaults instantiates a new CustomerBalanceTransactionExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerBalanceTransactionExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerBalanceTransactionExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerBalanceTransactionExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CustomerBalanceTransactionExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerBalanceTransactionExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerBalanceTransactionExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *CustomerBalanceTransactionExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerBalanceTransactionExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerBalanceTransactionExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerBalanceTransactionExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CustomerBalanceTransactionExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerBalanceTransactionExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerBalanceTransactionExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *CustomerBalanceTransactionExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CustomerBalanceTransactionExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CustomerBalanceTransactionExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CustomerBalanceTransactionExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetAmountAtom

`func (o *CustomerBalanceTransactionExternal) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *CustomerBalanceTransactionExternal) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *CustomerBalanceTransactionExternal) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetCurrency

`func (o *CustomerBalanceTransactionExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerBalanceTransactionExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerBalanceTransactionExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetEndingBalanceAmountAtom

`func (o *CustomerBalanceTransactionExternal) GetEndingBalanceAmountAtom() int32`

GetEndingBalanceAmountAtom returns the EndingBalanceAmountAtom field if non-nil, zero value otherwise.

### GetEndingBalanceAmountAtomOk

`func (o *CustomerBalanceTransactionExternal) GetEndingBalanceAmountAtomOk() (*int32, bool)`

GetEndingBalanceAmountAtomOk returns a tuple with the EndingBalanceAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingBalanceAmountAtom

`func (o *CustomerBalanceTransactionExternal) SetEndingBalanceAmountAtom(v int32)`

SetEndingBalanceAmountAtom sets EndingBalanceAmountAtom field to given value.


### GetType

`func (o *CustomerBalanceTransactionExternal) GetType() CustomerBalanceTransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerBalanceTransactionExternal) GetTypeOk() (*CustomerBalanceTransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerBalanceTransactionExternal) SetType(v CustomerBalanceTransactionType)`

SetType sets Type field to given value.


### GetInvoiceId

`func (o *CustomerBalanceTransactionExternal) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CustomerBalanceTransactionExternal) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CustomerBalanceTransactionExternal) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *CustomerBalanceTransactionExternal) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *CustomerBalanceTransactionExternal) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetCreditNoteId

`func (o *CustomerBalanceTransactionExternal) GetCreditNoteId() string`

GetCreditNoteId returns the CreditNoteId field if non-nil, zero value otherwise.

### GetCreditNoteIdOk

`func (o *CustomerBalanceTransactionExternal) GetCreditNoteIdOk() (*string, bool)`

GetCreditNoteIdOk returns a tuple with the CreditNoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteId

`func (o *CustomerBalanceTransactionExternal) SetCreditNoteId(v string)`

SetCreditNoteId sets CreditNoteId field to given value.


### SetCreditNoteIdNil

`func (o *CustomerBalanceTransactionExternal) SetCreditNoteIdNil(b bool)`

 SetCreditNoteIdNil sets the value for CreditNoteId to be an explicit nil

### UnsetCreditNoteId
`func (o *CustomerBalanceTransactionExternal) UnsetCreditNoteId()`

UnsetCreditNoteId ensures that no value is present for CreditNoteId, not even an explicit nil
### GetDescription

`func (o *CustomerBalanceTransactionExternal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerBalanceTransactionExternal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerBalanceTransactionExternal) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *CustomerBalanceTransactionExternal) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CustomerBalanceTransactionExternal) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


