# CustomerPaymentMethodQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**Expand** | Pointer to **[]string** | Specifies which fields in the response should be expanded. | [optional] 
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**ProviderType** | Pointer to [**NullablePaymentProviderType**](PaymentProviderType.md) |  | [optional] 
**SortDescending** | Pointer to **bool** | Sort direction. | [optional] [default to false]
**SortKey** | Pointer to **string** | Key name based on which data is sorted. | [optional] [default to "created_at"]
**UpdatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 

## Methods

### NewCustomerPaymentMethodQueryParams

`func NewCustomerPaymentMethodQueryParams() *CustomerPaymentMethodQueryParams`

NewCustomerPaymentMethodQueryParams instantiates a new CustomerPaymentMethodQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPaymentMethodQueryParamsWithDefaults

`func NewCustomerPaymentMethodQueryParamsWithDefaults() *CustomerPaymentMethodQueryParams`

NewCustomerPaymentMethodQueryParamsWithDefaults instantiates a new CustomerPaymentMethodQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CustomerPaymentMethodQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerPaymentMethodQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerPaymentMethodQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerPaymentMethodQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CustomerPaymentMethodQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CustomerPaymentMethodQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpand

`func (o *CustomerPaymentMethodQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *CustomerPaymentMethodQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *CustomerPaymentMethodQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *CustomerPaymentMethodQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetPageNumber

`func (o *CustomerPaymentMethodQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *CustomerPaymentMethodQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *CustomerPaymentMethodQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *CustomerPaymentMethodQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *CustomerPaymentMethodQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CustomerPaymentMethodQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CustomerPaymentMethodQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CustomerPaymentMethodQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProviderType

`func (o *CustomerPaymentMethodQueryParams) GetProviderType() PaymentProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *CustomerPaymentMethodQueryParams) GetProviderTypeOk() (*PaymentProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *CustomerPaymentMethodQueryParams) SetProviderType(v PaymentProviderType)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *CustomerPaymentMethodQueryParams) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### SetProviderTypeNil

`func (o *CustomerPaymentMethodQueryParams) SetProviderTypeNil(b bool)`

 SetProviderTypeNil sets the value for ProviderType to be an explicit nil

### UnsetProviderType
`func (o *CustomerPaymentMethodQueryParams) UnsetProviderType()`

UnsetProviderType ensures that no value is present for ProviderType, not even an explicit nil
### GetSortDescending

`func (o *CustomerPaymentMethodQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *CustomerPaymentMethodQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *CustomerPaymentMethodQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *CustomerPaymentMethodQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetSortKey

`func (o *CustomerPaymentMethodQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *CustomerPaymentMethodQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *CustomerPaymentMethodQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *CustomerPaymentMethodQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomerPaymentMethodQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerPaymentMethodQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerPaymentMethodQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomerPaymentMethodQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CustomerPaymentMethodQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CustomerPaymentMethodQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


