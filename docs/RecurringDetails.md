# RecurringDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregateUsage** | Pointer to [**NullableUsageAggMethodEnum**](UsageAggMethodEnum.md) |  | [optional] 
**TrialPeriodDays** | Pointer to **int32** |  | [optional] [default to 0]
**UsageType** | [**UsageTypeEnum**](UsageTypeEnum.md) |  | 

## Methods

### NewRecurringDetails

`func NewRecurringDetails(usageType UsageTypeEnum, ) *RecurringDetails`

NewRecurringDetails instantiates a new RecurringDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringDetailsWithDefaults

`func NewRecurringDetailsWithDefaults() *RecurringDetails`

NewRecurringDetailsWithDefaults instantiates a new RecurringDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregateUsage

`func (o *RecurringDetails) GetAggregateUsage() UsageAggMethodEnum`

GetAggregateUsage returns the AggregateUsage field if non-nil, zero value otherwise.

### GetAggregateUsageOk

`func (o *RecurringDetails) GetAggregateUsageOk() (*UsageAggMethodEnum, bool)`

GetAggregateUsageOk returns a tuple with the AggregateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateUsage

`func (o *RecurringDetails) SetAggregateUsage(v UsageAggMethodEnum)`

SetAggregateUsage sets AggregateUsage field to given value.

### HasAggregateUsage

`func (o *RecurringDetails) HasAggregateUsage() bool`

HasAggregateUsage returns a boolean if a field has been set.

### SetAggregateUsageNil

`func (o *RecurringDetails) SetAggregateUsageNil(b bool)`

 SetAggregateUsageNil sets the value for AggregateUsage to be an explicit nil

### UnsetAggregateUsage
`func (o *RecurringDetails) UnsetAggregateUsage()`

UnsetAggregateUsage ensures that no value is present for AggregateUsage, not even an explicit nil
### GetTrialPeriodDays

`func (o *RecurringDetails) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *RecurringDetails) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *RecurringDetails) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *RecurringDetails) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### GetUsageType

`func (o *RecurringDetails) GetUsageType() UsageTypeEnum`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *RecurringDetails) GetUsageTypeOk() (*UsageTypeEnum, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *RecurringDetails) SetUsageType(v UsageTypeEnum)`

SetUsageType sets UsageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


