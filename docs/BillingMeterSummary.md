# BillingMeterSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregatedValue** | **int32** | Aggregated value of all the events within start_time and end_time. The aggregation strategy is defined on meter via default_aggregation. | 
**DocumentCount** | **int32** | Count of events found within start_time and end_time. | 
**EndDatetime** | **time.Time** | End datetime for this event summary | 
**MeterId** | **string** | The if of billing meter associated with this event summary. | 
**StartDatetime** | **time.Time** | Start datetime for this event summary | 

## Methods

### NewBillingMeterSummary

`func NewBillingMeterSummary(aggregatedValue int32, documentCount int32, endDatetime time.Time, meterId string, startDatetime time.Time, ) *BillingMeterSummary`

NewBillingMeterSummary instantiates a new BillingMeterSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingMeterSummaryWithDefaults

`func NewBillingMeterSummaryWithDefaults() *BillingMeterSummary`

NewBillingMeterSummaryWithDefaults instantiates a new BillingMeterSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatedValue

`func (o *BillingMeterSummary) GetAggregatedValue() int32`

GetAggregatedValue returns the AggregatedValue field if non-nil, zero value otherwise.

### GetAggregatedValueOk

`func (o *BillingMeterSummary) GetAggregatedValueOk() (*int32, bool)`

GetAggregatedValueOk returns a tuple with the AggregatedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedValue

`func (o *BillingMeterSummary) SetAggregatedValue(v int32)`

SetAggregatedValue sets AggregatedValue field to given value.


### GetDocumentCount

`func (o *BillingMeterSummary) GetDocumentCount() int32`

GetDocumentCount returns the DocumentCount field if non-nil, zero value otherwise.

### GetDocumentCountOk

`func (o *BillingMeterSummary) GetDocumentCountOk() (*int32, bool)`

GetDocumentCountOk returns a tuple with the DocumentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCount

`func (o *BillingMeterSummary) SetDocumentCount(v int32)`

SetDocumentCount sets DocumentCount field to given value.


### GetEndDatetime

`func (o *BillingMeterSummary) GetEndDatetime() time.Time`

GetEndDatetime returns the EndDatetime field if non-nil, zero value otherwise.

### GetEndDatetimeOk

`func (o *BillingMeterSummary) GetEndDatetimeOk() (*time.Time, bool)`

GetEndDatetimeOk returns a tuple with the EndDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDatetime

`func (o *BillingMeterSummary) SetEndDatetime(v time.Time)`

SetEndDatetime sets EndDatetime field to given value.


### GetMeterId

`func (o *BillingMeterSummary) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *BillingMeterSummary) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *BillingMeterSummary) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.


### GetStartDatetime

`func (o *BillingMeterSummary) GetStartDatetime() time.Time`

GetStartDatetime returns the StartDatetime field if non-nil, zero value otherwise.

### GetStartDatetimeOk

`func (o *BillingMeterSummary) GetStartDatetimeOk() (*time.Time, bool)`

GetStartDatetimeOk returns a tuple with the StartDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDatetime

`func (o *BillingMeterSummary) SetStartDatetime(v time.Time)`

SetStartDatetime sets StartDatetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


