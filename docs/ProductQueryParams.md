# ProductQueryParams

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
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**NameContains** | Pointer to **NullableString** |  | [optional] 
**Search** | Pointer to [**NullableSearchFilter**](SearchFilter.md) |  | [optional] 

## Methods

### NewProductQueryParams

`func NewProductQueryParams() *ProductQueryParams`

NewProductQueryParams instantiates a new ProductQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductQueryParamsWithDefaults

`func NewProductQueryParamsWithDefaults() *ProductQueryParams`

NewProductQueryParamsWithDefaults instantiates a new ProductQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *ProductQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ProductQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ProductQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ProductQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ProductQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ProductQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ProductQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ProductQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortKey

`func (o *ProductQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *ProductQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *ProductQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *ProductQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetSortDescending

`func (o *ProductQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *ProductQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *ProductQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *ProductQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProductQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ProductQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ProductQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ProductQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProductQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProductQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProductQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ProductQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ProductQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetExpand

`func (o *ProductQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *ProductQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *ProductQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *ProductQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetIsActive

`func (o *ProductQueryParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProductQueryParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProductQueryParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProductQueryParams) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProductQueryParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProductQueryParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCreatedBy

`func (o *ProductQueryParams) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProductQueryParams) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProductQueryParams) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ProductQueryParams) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ProductQueryParams) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ProductQueryParams) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetNameContains

`func (o *ProductQueryParams) GetNameContains() string`

GetNameContains returns the NameContains field if non-nil, zero value otherwise.

### GetNameContainsOk

`func (o *ProductQueryParams) GetNameContainsOk() (*string, bool)`

GetNameContainsOk returns a tuple with the NameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameContains

`func (o *ProductQueryParams) SetNameContains(v string)`

SetNameContains sets NameContains field to given value.

### HasNameContains

`func (o *ProductQueryParams) HasNameContains() bool`

HasNameContains returns a boolean if a field has been set.

### SetNameContainsNil

`func (o *ProductQueryParams) SetNameContainsNil(b bool)`

 SetNameContainsNil sets the value for NameContains to be an explicit nil

### UnsetNameContains
`func (o *ProductQueryParams) UnsetNameContains()`

UnsetNameContains ensures that no value is present for NameContains, not even an explicit nil
### GetSearch

`func (o *ProductQueryParams) GetSearch() SearchFilter`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ProductQueryParams) GetSearchOk() (*SearchFilter, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ProductQueryParams) SetSearch(v SearchFilter)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ProductQueryParams) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *ProductQueryParams) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *ProductQueryParams) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


