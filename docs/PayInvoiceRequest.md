# PayInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodId** | Pointer to **NullableString** |  | [optional] 
**PaidOutOfBand** | Pointer to **bool** | Boolean representing whether an invoice is paid outside of OpenPay. This will result in no charge being made. Defaults to false. | [optional] [default to false]
**Comment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPayInvoiceRequest

`func NewPayInvoiceRequest() *PayInvoiceRequest`

NewPayInvoiceRequest instantiates a new PayInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayInvoiceRequestWithDefaults

`func NewPayInvoiceRequestWithDefaults() *PayInvoiceRequest`

NewPayInvoiceRequestWithDefaults instantiates a new PayInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodId

`func (o *PayInvoiceRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *PayInvoiceRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *PayInvoiceRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *PayInvoiceRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### SetPaymentMethodIdNil

`func (o *PayInvoiceRequest) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *PayInvoiceRequest) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetPaidOutOfBand

`func (o *PayInvoiceRequest) GetPaidOutOfBand() bool`

GetPaidOutOfBand returns the PaidOutOfBand field if non-nil, zero value otherwise.

### GetPaidOutOfBandOk

`func (o *PayInvoiceRequest) GetPaidOutOfBandOk() (*bool, bool)`

GetPaidOutOfBandOk returns a tuple with the PaidOutOfBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidOutOfBand

`func (o *PayInvoiceRequest) SetPaidOutOfBand(v bool)`

SetPaidOutOfBand sets PaidOutOfBand field to given value.

### HasPaidOutOfBand

`func (o *PayInvoiceRequest) HasPaidOutOfBand() bool`

HasPaidOutOfBand returns a boolean if a field has been set.

### GetComment

`func (o *PayInvoiceRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PayInvoiceRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PayInvoiceRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PayInvoiceRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *PayInvoiceRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *PayInvoiceRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


