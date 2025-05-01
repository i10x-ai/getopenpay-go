# CreateCreditNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditAmountAtom** | Pointer to **int32** | The integer amount representing the amount to credit the customer’s balance, which will be automatically applied to their next invoice. | [optional] [default to 0]
**Currency** | Pointer to [**CurrencyEnum**](CurrencyEnum.md) |  | [optional] 
**InvoiceId** | **string** | ID of the invoice | 
**Lines** | [**[]CreateCreditNoteLine**](CreateCreditNoteLine.md) | Line items that make up the credit note. | 
**Reason** | Pointer to [**NullableCreditNoteReason**](CreditNoteReason.md) |  | [optional] 
**RefundAmountAtom** | Pointer to **int32** | The integer amount representing the amount to refund. If set, a refund will be created for the charge associated with the invoice. | [optional] [default to 0]
**TotalAmountAtom** | **int32** | The int amount representing the total amount of the credit note. | 

## Methods

### NewCreateCreditNoteRequest

`func NewCreateCreditNoteRequest(invoiceId string, lines []CreateCreditNoteLine, totalAmountAtom int32, ) *CreateCreditNoteRequest`

NewCreateCreditNoteRequest instantiates a new CreateCreditNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCreditNoteRequestWithDefaults

`func NewCreateCreditNoteRequestWithDefaults() *CreateCreditNoteRequest`

NewCreateCreditNoteRequestWithDefaults instantiates a new CreateCreditNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditAmountAtom

`func (o *CreateCreditNoteRequest) GetCreditAmountAtom() int32`

GetCreditAmountAtom returns the CreditAmountAtom field if non-nil, zero value otherwise.

### GetCreditAmountAtomOk

`func (o *CreateCreditNoteRequest) GetCreditAmountAtomOk() (*int32, bool)`

GetCreditAmountAtomOk returns a tuple with the CreditAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountAtom

`func (o *CreateCreditNoteRequest) SetCreditAmountAtom(v int32)`

SetCreditAmountAtom sets CreditAmountAtom field to given value.

### HasCreditAmountAtom

`func (o *CreateCreditNoteRequest) HasCreditAmountAtom() bool`

HasCreditAmountAtom returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateCreditNoteRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCreditNoteRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCreditNoteRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateCreditNoteRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInvoiceId

`func (o *CreateCreditNoteRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreateCreditNoteRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreateCreditNoteRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetLines

`func (o *CreateCreditNoteRequest) GetLines() []CreateCreditNoteLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *CreateCreditNoteRequest) GetLinesOk() (*[]CreateCreditNoteLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *CreateCreditNoteRequest) SetLines(v []CreateCreditNoteLine)`

SetLines sets Lines field to given value.


### GetReason

`func (o *CreateCreditNoteRequest) GetReason() CreditNoteReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateCreditNoteRequest) GetReasonOk() (*CreditNoteReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateCreditNoteRequest) SetReason(v CreditNoteReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateCreditNoteRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *CreateCreditNoteRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *CreateCreditNoteRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRefundAmountAtom

`func (o *CreateCreditNoteRequest) GetRefundAmountAtom() int32`

GetRefundAmountAtom returns the RefundAmountAtom field if non-nil, zero value otherwise.

### GetRefundAmountAtomOk

`func (o *CreateCreditNoteRequest) GetRefundAmountAtomOk() (*int32, bool)`

GetRefundAmountAtomOk returns a tuple with the RefundAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmountAtom

`func (o *CreateCreditNoteRequest) SetRefundAmountAtom(v int32)`

SetRefundAmountAtom sets RefundAmountAtom field to given value.

### HasRefundAmountAtom

`func (o *CreateCreditNoteRequest) HasRefundAmountAtom() bool`

HasRefundAmountAtom returns a boolean if a field has been set.

### GetTotalAmountAtom

`func (o *CreateCreditNoteRequest) GetTotalAmountAtom() int32`

GetTotalAmountAtom returns the TotalAmountAtom field if non-nil, zero value otherwise.

### GetTotalAmountAtomOk

`func (o *CreateCreditNoteRequest) GetTotalAmountAtomOk() (*int32, bool)`

GetTotalAmountAtomOk returns a tuple with the TotalAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountAtom

`func (o *CreateCreditNoteRequest) SetTotalAmountAtom(v int32)`

SetTotalAmountAtom sets TotalAmountAtom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


