# InvoiceItemsQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Expand** | Pointer to **[]string** | Specifies which fields in the response should be expanded. | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**SortDescending** | Pointer to **bool** | Sort direction. | [optional] [default to false]
**SortKey** | Pointer to **string** | Key name based on which data is sorted. | [optional] [default to "created_at"]
**UpdatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 

## Methods

### NewInvoiceItemsQueryParams

`func NewInvoiceItemsQueryParams() *InvoiceItemsQueryParams`

NewInvoiceItemsQueryParams instantiates a new InvoiceItemsQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceItemsQueryParamsWithDefaults

`func NewInvoiceItemsQueryParamsWithDefaults() *InvoiceItemsQueryParams`

NewInvoiceItemsQueryParamsWithDefaults instantiates a new InvoiceItemsQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouponId

`func (o *InvoiceItemsQueryParams) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *InvoiceItemsQueryParams) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *InvoiceItemsQueryParams) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *InvoiceItemsQueryParams) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *InvoiceItemsQueryParams) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *InvoiceItemsQueryParams) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetCreatedAt

`func (o *InvoiceItemsQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceItemsQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceItemsQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InvoiceItemsQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *InvoiceItemsQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *InvoiceItemsQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCustomerId

`func (o *InvoiceItemsQueryParams) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InvoiceItemsQueryParams) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InvoiceItemsQueryParams) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InvoiceItemsQueryParams) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *InvoiceItemsQueryParams) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *InvoiceItemsQueryParams) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetExpand

`func (o *InvoiceItemsQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *InvoiceItemsQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *InvoiceItemsQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *InvoiceItemsQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetInvoiceId

`func (o *InvoiceItemsQueryParams) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceItemsQueryParams) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceItemsQueryParams) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *InvoiceItemsQueryParams) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *InvoiceItemsQueryParams) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *InvoiceItemsQueryParams) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetPageNumber

`func (o *InvoiceItemsQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *InvoiceItemsQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *InvoiceItemsQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *InvoiceItemsQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *InvoiceItemsQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *InvoiceItemsQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *InvoiceItemsQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *InvoiceItemsQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortDescending

`func (o *InvoiceItemsQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *InvoiceItemsQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *InvoiceItemsQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *InvoiceItemsQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetSortKey

`func (o *InvoiceItemsQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *InvoiceItemsQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *InvoiceItemsQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *InvoiceItemsQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InvoiceItemsQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceItemsQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceItemsQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InvoiceItemsQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *InvoiceItemsQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *InvoiceItemsQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


