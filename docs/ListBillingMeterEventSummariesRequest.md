# ListBillingMeterEventSummariesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The customer for which to fetch event summaries. | 
**EndTime** | **time.Time** | The timestamp from when to stop aggregating meter events (exclusive). Must be aligned with minute boundaries. | 
**MeterId** | **string** | Unique identifier for the meter object. | 
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**StartTime** | **time.Time** | The timestamp from when to start aggregating meter events (inclusive). Must be aligned with minute boundaries. | 
**ValueGroupingWindow** | Pointer to [**NullableMeterEventValueGroupingWindow**](MeterEventValueGroupingWindow.md) |  | [optional] 

## Methods

### NewListBillingMeterEventSummariesRequest

`func NewListBillingMeterEventSummariesRequest(customerId string, endTime time.Time, meterId string, startTime time.Time, ) *ListBillingMeterEventSummariesRequest`

NewListBillingMeterEventSummariesRequest instantiates a new ListBillingMeterEventSummariesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingMeterEventSummariesRequestWithDefaults

`func NewListBillingMeterEventSummariesRequestWithDefaults() *ListBillingMeterEventSummariesRequest`

NewListBillingMeterEventSummariesRequestWithDefaults instantiates a new ListBillingMeterEventSummariesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *ListBillingMeterEventSummariesRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ListBillingMeterEventSummariesRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ListBillingMeterEventSummariesRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetEndTime

`func (o *ListBillingMeterEventSummariesRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ListBillingMeterEventSummariesRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ListBillingMeterEventSummariesRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetMeterId

`func (o *ListBillingMeterEventSummariesRequest) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *ListBillingMeterEventSummariesRequest) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *ListBillingMeterEventSummariesRequest) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.


### GetPageNumber

`func (o *ListBillingMeterEventSummariesRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ListBillingMeterEventSummariesRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ListBillingMeterEventSummariesRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ListBillingMeterEventSummariesRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ListBillingMeterEventSummariesRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListBillingMeterEventSummariesRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListBillingMeterEventSummariesRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListBillingMeterEventSummariesRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStartTime

`func (o *ListBillingMeterEventSummariesRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ListBillingMeterEventSummariesRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ListBillingMeterEventSummariesRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetValueGroupingWindow

`func (o *ListBillingMeterEventSummariesRequest) GetValueGroupingWindow() MeterEventValueGroupingWindow`

GetValueGroupingWindow returns the ValueGroupingWindow field if non-nil, zero value otherwise.

### GetValueGroupingWindowOk

`func (o *ListBillingMeterEventSummariesRequest) GetValueGroupingWindowOk() (*MeterEventValueGroupingWindow, bool)`

GetValueGroupingWindowOk returns a tuple with the ValueGroupingWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueGroupingWindow

`func (o *ListBillingMeterEventSummariesRequest) SetValueGroupingWindow(v MeterEventValueGroupingWindow)`

SetValueGroupingWindow sets ValueGroupingWindow field to given value.

### HasValueGroupingWindow

`func (o *ListBillingMeterEventSummariesRequest) HasValueGroupingWindow() bool`

HasValueGroupingWindow returns a boolean if a field has been set.

### SetValueGroupingWindowNil

`func (o *ListBillingMeterEventSummariesRequest) SetValueGroupingWindowNil(b bool)`

 SetValueGroupingWindowNil sets the value for ValueGroupingWindow to be an explicit nil

### UnsetValueGroupingWindow
`func (o *ListBillingMeterEventSummariesRequest) UnsetValueGroupingWindow()`

UnsetValueGroupingWindow ensures that no value is present for ValueGroupingWindow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


