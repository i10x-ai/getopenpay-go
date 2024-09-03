# ChargeQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**SortKey** | Pointer to **string** | Key name based on which data is sorted. | [optional] [default to "created_at"]
**SortDescending** | Pointer to **bool** | Sort direction. | [optional] [default to false]
**CreatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**UpdatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**Expand** | Pointer to **[]string** |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**PaymentMethodId** | Pointer to **NullableString** |  | [optional] 
**PaymentIntentId** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullableChargeStatusEnum**](ChargeStatusEnum.md) |  | [optional] 
**Refunded** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewChargeQueryParams

`func NewChargeQueryParams() *ChargeQueryParams`

NewChargeQueryParams instantiates a new ChargeQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeQueryParamsWithDefaults

`func NewChargeQueryParamsWithDefaults() *ChargeQueryParams`

NewChargeQueryParamsWithDefaults instantiates a new ChargeQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *ChargeQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ChargeQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ChargeQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ChargeQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ChargeQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ChargeQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ChargeQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ChargeQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortKey

`func (o *ChargeQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *ChargeQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *ChargeQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *ChargeQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetSortDescending

`func (o *ChargeQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *ChargeQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *ChargeQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *ChargeQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChargeQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChargeQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChargeQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChargeQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ChargeQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ChargeQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ChargeQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChargeQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChargeQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ChargeQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ChargeQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ChargeQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetExpand

`func (o *ChargeQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *ChargeQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *ChargeQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *ChargeQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetCustomerId

`func (o *ChargeQueryParams) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ChargeQueryParams) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ChargeQueryParams) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ChargeQueryParams) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *ChargeQueryParams) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *ChargeQueryParams) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetPaymentMethodId

`func (o *ChargeQueryParams) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *ChargeQueryParams) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *ChargeQueryParams) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *ChargeQueryParams) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### SetPaymentMethodIdNil

`func (o *ChargeQueryParams) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *ChargeQueryParams) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetPaymentIntentId

`func (o *ChargeQueryParams) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *ChargeQueryParams) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *ChargeQueryParams) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.

### HasPaymentIntentId

`func (o *ChargeQueryParams) HasPaymentIntentId() bool`

HasPaymentIntentId returns a boolean if a field has been set.

### SetPaymentIntentIdNil

`func (o *ChargeQueryParams) SetPaymentIntentIdNil(b bool)`

 SetPaymentIntentIdNil sets the value for PaymentIntentId to be an explicit nil

### UnsetPaymentIntentId
`func (o *ChargeQueryParams) UnsetPaymentIntentId()`

UnsetPaymentIntentId ensures that no value is present for PaymentIntentId, not even an explicit nil
### GetInvoiceId

`func (o *ChargeQueryParams) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ChargeQueryParams) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ChargeQueryParams) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *ChargeQueryParams) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *ChargeQueryParams) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *ChargeQueryParams) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetStatus

`func (o *ChargeQueryParams) GetStatus() ChargeStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargeQueryParams) GetStatusOk() (*ChargeStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargeQueryParams) SetStatus(v ChargeStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChargeQueryParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ChargeQueryParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ChargeQueryParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRefunded

`func (o *ChargeQueryParams) GetRefunded() bool`

GetRefunded returns the Refunded field if non-nil, zero value otherwise.

### GetRefundedOk

`func (o *ChargeQueryParams) GetRefundedOk() (*bool, bool)`

GetRefundedOk returns a tuple with the Refunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunded

`func (o *ChargeQueryParams) SetRefunded(v bool)`

SetRefunded sets Refunded field to given value.

### HasRefunded

`func (o *ChargeQueryParams) HasRefunded() bool`

HasRefunded returns a boolean if a field has been set.

### SetRefundedNil

`func (o *ChargeQueryParams) SetRefundedNil(b bool)`

 SetRefundedNil sets the value for Refunded to be an explicit nil

### UnsetRefunded
`func (o *ChargeQueryParams) UnsetRefunded()`

UnsetRefunded ensures that no value is present for Refunded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


