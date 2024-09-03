# CheckoutSessionQueryParams

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
**PaymentLinkId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCheckoutSessionQueryParams

`func NewCheckoutSessionQueryParams() *CheckoutSessionQueryParams`

NewCheckoutSessionQueryParams instantiates a new CheckoutSessionQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionQueryParamsWithDefaults

`func NewCheckoutSessionQueryParamsWithDefaults() *CheckoutSessionQueryParams`

NewCheckoutSessionQueryParamsWithDefaults instantiates a new CheckoutSessionQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *CheckoutSessionQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *CheckoutSessionQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *CheckoutSessionQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *CheckoutSessionQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *CheckoutSessionQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CheckoutSessionQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CheckoutSessionQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CheckoutSessionQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortKey

`func (o *CheckoutSessionQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *CheckoutSessionQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *CheckoutSessionQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *CheckoutSessionQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetSortDescending

`func (o *CheckoutSessionQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *CheckoutSessionQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *CheckoutSessionQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *CheckoutSessionQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CheckoutSessionQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckoutSessionQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckoutSessionQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CheckoutSessionQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CheckoutSessionQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CheckoutSessionQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *CheckoutSessionQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CheckoutSessionQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CheckoutSessionQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CheckoutSessionQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CheckoutSessionQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CheckoutSessionQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetExpand

`func (o *CheckoutSessionQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *CheckoutSessionQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *CheckoutSessionQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *CheckoutSessionQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetPaymentLinkId

`func (o *CheckoutSessionQueryParams) GetPaymentLinkId() string`

GetPaymentLinkId returns the PaymentLinkId field if non-nil, zero value otherwise.

### GetPaymentLinkIdOk

`func (o *CheckoutSessionQueryParams) GetPaymentLinkIdOk() (*string, bool)`

GetPaymentLinkIdOk returns a tuple with the PaymentLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLinkId

`func (o *CheckoutSessionQueryParams) SetPaymentLinkId(v string)`

SetPaymentLinkId sets PaymentLinkId field to given value.

### HasPaymentLinkId

`func (o *CheckoutSessionQueryParams) HasPaymentLinkId() bool`

HasPaymentLinkId returns a boolean if a field has been set.

### SetPaymentLinkIdNil

`func (o *CheckoutSessionQueryParams) SetPaymentLinkIdNil(b bool)`

 SetPaymentLinkIdNil sets the value for PaymentLinkId to be an explicit nil

### UnsetPaymentLinkId
`func (o *CheckoutSessionQueryParams) UnsetPaymentLinkId()`

UnsetPaymentLinkId ensures that no value is present for PaymentLinkId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


