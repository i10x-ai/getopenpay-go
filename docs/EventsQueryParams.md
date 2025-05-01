# EventsQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**DeliverySuccess** | Pointer to **NullableBool** |  | [optional] 
**Expand** | Pointer to **[]string** | Specifies which fields in the response should be expanded. | [optional] 
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**SortDescending** | Pointer to **bool** | Sort direction. | [optional] [default to false]
**SortKey** | Pointer to **string** | Key name based on which data is sorted. | [optional] [default to "created_at"]
**Types** | Pointer to [**[]EventType**](EventType.md) | An array of up to 20 strings containing specific event names. The list will be filtered to include only events with a matching event property. | [optional] 
**UpdatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 

## Methods

### NewEventsQueryParams

`func NewEventsQueryParams() *EventsQueryParams`

NewEventsQueryParams instantiates a new EventsQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsQueryParamsWithDefaults

`func NewEventsQueryParamsWithDefaults() *EventsQueryParams`

NewEventsQueryParamsWithDefaults instantiates a new EventsQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EventsQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventsQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventsQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventsQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *EventsQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *EventsQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeliverySuccess

`func (o *EventsQueryParams) GetDeliverySuccess() bool`

GetDeliverySuccess returns the DeliverySuccess field if non-nil, zero value otherwise.

### GetDeliverySuccessOk

`func (o *EventsQueryParams) GetDeliverySuccessOk() (*bool, bool)`

GetDeliverySuccessOk returns a tuple with the DeliverySuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverySuccess

`func (o *EventsQueryParams) SetDeliverySuccess(v bool)`

SetDeliverySuccess sets DeliverySuccess field to given value.

### HasDeliverySuccess

`func (o *EventsQueryParams) HasDeliverySuccess() bool`

HasDeliverySuccess returns a boolean if a field has been set.

### SetDeliverySuccessNil

`func (o *EventsQueryParams) SetDeliverySuccessNil(b bool)`

 SetDeliverySuccessNil sets the value for DeliverySuccess to be an explicit nil

### UnsetDeliverySuccess
`func (o *EventsQueryParams) UnsetDeliverySuccess()`

UnsetDeliverySuccess ensures that no value is present for DeliverySuccess, not even an explicit nil
### GetExpand

`func (o *EventsQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *EventsQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *EventsQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *EventsQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetPageNumber

`func (o *EventsQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *EventsQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *EventsQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *EventsQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *EventsQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EventsQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EventsQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *EventsQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortDescending

`func (o *EventsQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *EventsQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *EventsQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *EventsQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetSortKey

`func (o *EventsQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *EventsQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *EventsQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *EventsQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetTypes

`func (o *EventsQueryParams) GetTypes() []EventType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *EventsQueryParams) GetTypesOk() (*[]EventType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *EventsQueryParams) SetTypes(v []EventType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *EventsQueryParams) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EventsQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventsQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventsQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EventsQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *EventsQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *EventsQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


