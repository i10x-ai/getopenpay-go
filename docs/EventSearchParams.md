# EventSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**Expand** | Pointer to **[]string** | Specifies which fields in the response should be expanded. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of documents to retrieve. | [optional] [default to 50]
**ObjectId** | **string** | Unique ID of the object. | 
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**SortDescending** | Pointer to **bool** | Sort direction. | [optional] [default to false]
**SortKey** | Pointer to **string** | Key name based on which data is sorted. | [optional] [default to "created_at"]
**Types** | Pointer to [**[]EventType**](EventType.md) | An array of up to 20 strings containing specific event names. The list will be filtered to include only events with a matching event property. | [optional] 
**UpdatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 

## Methods

### NewEventSearchParams

`func NewEventSearchParams(objectId string, ) *EventSearchParams`

NewEventSearchParams instantiates a new EventSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSearchParamsWithDefaults

`func NewEventSearchParamsWithDefaults() *EventSearchParams`

NewEventSearchParamsWithDefaults instantiates a new EventSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EventSearchParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventSearchParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventSearchParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventSearchParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *EventSearchParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *EventSearchParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpand

`func (o *EventSearchParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *EventSearchParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *EventSearchParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *EventSearchParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetLimit

`func (o *EventSearchParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *EventSearchParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *EventSearchParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *EventSearchParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetObjectId

`func (o *EventSearchParams) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *EventSearchParams) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *EventSearchParams) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetPageNumber

`func (o *EventSearchParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *EventSearchParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *EventSearchParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *EventSearchParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *EventSearchParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EventSearchParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EventSearchParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *EventSearchParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortDescending

`func (o *EventSearchParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *EventSearchParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *EventSearchParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *EventSearchParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetSortKey

`func (o *EventSearchParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *EventSearchParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *EventSearchParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *EventSearchParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetTypes

`func (o *EventSearchParams) GetTypes() []EventType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *EventSearchParams) GetTypesOk() (*[]EventType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *EventSearchParams) SetTypes(v []EventType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *EventSearchParams) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EventSearchParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventSearchParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventSearchParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EventSearchParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *EventSearchParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *EventSearchParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


