# RefundQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountAtom** | Pointer to [**NullableIntRangeFilter**](IntRangeFilter.md) |  | [optional] 
**ChargeId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Expand** | Pointer to **[]string** | Specifies which fields in the response should be expanded. | [optional] 
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**PaymentIntentId** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to [**NullableRefundReasonEnum**](RefundReasonEnum.md) |  | [optional] 
**SortDescending** | Pointer to **bool** | Sort direction. | [optional] [default to false]
**SortKey** | Pointer to **string** | Key name based on which data is sorted. | [optional] [default to "created_at"]
**Status** | Pointer to [**NullableRefundStatusEnum**](RefundStatusEnum.md) |  | [optional] 
**UpdatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 

## Methods

### NewRefundQueryParams

`func NewRefundQueryParams() *RefundQueryParams`

NewRefundQueryParams instantiates a new RefundQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundQueryParamsWithDefaults

`func NewRefundQueryParamsWithDefaults() *RefundQueryParams`

NewRefundQueryParamsWithDefaults instantiates a new RefundQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountAtom

`func (o *RefundQueryParams) GetAmountAtom() IntRangeFilter`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *RefundQueryParams) GetAmountAtomOk() (*IntRangeFilter, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *RefundQueryParams) SetAmountAtom(v IntRangeFilter)`

SetAmountAtom sets AmountAtom field to given value.

### HasAmountAtom

`func (o *RefundQueryParams) HasAmountAtom() bool`

HasAmountAtom returns a boolean if a field has been set.

### SetAmountAtomNil

`func (o *RefundQueryParams) SetAmountAtomNil(b bool)`

 SetAmountAtomNil sets the value for AmountAtom to be an explicit nil

### UnsetAmountAtom
`func (o *RefundQueryParams) UnsetAmountAtom()`

UnsetAmountAtom ensures that no value is present for AmountAtom, not even an explicit nil
### GetChargeId

`func (o *RefundQueryParams) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *RefundQueryParams) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *RefundQueryParams) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.

### HasChargeId

`func (o *RefundQueryParams) HasChargeId() bool`

HasChargeId returns a boolean if a field has been set.

### SetChargeIdNil

`func (o *RefundQueryParams) SetChargeIdNil(b bool)`

 SetChargeIdNil sets the value for ChargeId to be an explicit nil

### UnsetChargeId
`func (o *RefundQueryParams) UnsetChargeId()`

UnsetChargeId ensures that no value is present for ChargeId, not even an explicit nil
### GetCreatedAt

`func (o *RefundQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RefundQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RefundQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RefundQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *RefundQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *RefundQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCustomerId

`func (o *RefundQueryParams) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *RefundQueryParams) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *RefundQueryParams) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *RefundQueryParams) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *RefundQueryParams) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *RefundQueryParams) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetExpand

`func (o *RefundQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *RefundQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *RefundQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *RefundQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetPageNumber

`func (o *RefundQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *RefundQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *RefundQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *RefundQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *RefundQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *RefundQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *RefundQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *RefundQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPaymentIntentId

`func (o *RefundQueryParams) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *RefundQueryParams) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *RefundQueryParams) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.

### HasPaymentIntentId

`func (o *RefundQueryParams) HasPaymentIntentId() bool`

HasPaymentIntentId returns a boolean if a field has been set.

### SetPaymentIntentIdNil

`func (o *RefundQueryParams) SetPaymentIntentIdNil(b bool)`

 SetPaymentIntentIdNil sets the value for PaymentIntentId to be an explicit nil

### UnsetPaymentIntentId
`func (o *RefundQueryParams) UnsetPaymentIntentId()`

UnsetPaymentIntentId ensures that no value is present for PaymentIntentId, not even an explicit nil
### GetReason

`func (o *RefundQueryParams) GetReason() RefundReasonEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RefundQueryParams) GetReasonOk() (*RefundReasonEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RefundQueryParams) SetReason(v RefundReasonEnum)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RefundQueryParams) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *RefundQueryParams) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *RefundQueryParams) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetSortDescending

`func (o *RefundQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *RefundQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *RefundQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *RefundQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetSortKey

`func (o *RefundQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *RefundQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *RefundQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *RefundQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetStatus

`func (o *RefundQueryParams) GetStatus() RefundStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RefundQueryParams) GetStatusOk() (*RefundStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RefundQueryParams) SetStatus(v RefundStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RefundQueryParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RefundQueryParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RefundQueryParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUpdatedAt

`func (o *RefundQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RefundQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RefundQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RefundQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *RefundQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *RefundQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


