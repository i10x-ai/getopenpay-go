# CreateBillingMeterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationFormula** | Pointer to [**BillingMeterAggregationFormula**](BillingMeterAggregationFormula.md) | Specifies how events are aggregated. Allowed values are count to count the number of events and sum to sum each event’s value. | [optional] 
**DisplayName** | **string** | The meter’s name. | 
**EventName** | **string** | The name of the meter event to record usage for. Corresponds with the event_name field on meter events. | 
**EventPayloadCustomerMappingKey** | Pointer to **string** | The key in the usage event payload to use for mapping the event to a customer. | [optional] [default to "customer_id"]
**EventPayloadValueKey** | Pointer to **string** | The key in the usage event payload to use as the value for this meter. For example, if the event payload contains usage on a bytes_used field, then set the event_payload_value_key to “bytes_used”. | [optional] [default to "value"]

## Methods

### NewCreateBillingMeterRequest

`func NewCreateBillingMeterRequest(displayName string, eventName string, ) *CreateBillingMeterRequest`

NewCreateBillingMeterRequest instantiates a new CreateBillingMeterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBillingMeterRequestWithDefaults

`func NewCreateBillingMeterRequestWithDefaults() *CreateBillingMeterRequest`

NewCreateBillingMeterRequestWithDefaults instantiates a new CreateBillingMeterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationFormula

`func (o *CreateBillingMeterRequest) GetAggregationFormula() BillingMeterAggregationFormula`

GetAggregationFormula returns the AggregationFormula field if non-nil, zero value otherwise.

### GetAggregationFormulaOk

`func (o *CreateBillingMeterRequest) GetAggregationFormulaOk() (*BillingMeterAggregationFormula, bool)`

GetAggregationFormulaOk returns a tuple with the AggregationFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFormula

`func (o *CreateBillingMeterRequest) SetAggregationFormula(v BillingMeterAggregationFormula)`

SetAggregationFormula sets AggregationFormula field to given value.

### HasAggregationFormula

`func (o *CreateBillingMeterRequest) HasAggregationFormula() bool`

HasAggregationFormula returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateBillingMeterRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateBillingMeterRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateBillingMeterRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEventName

`func (o *CreateBillingMeterRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *CreateBillingMeterRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *CreateBillingMeterRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventPayloadCustomerMappingKey

`func (o *CreateBillingMeterRequest) GetEventPayloadCustomerMappingKey() string`

GetEventPayloadCustomerMappingKey returns the EventPayloadCustomerMappingKey field if non-nil, zero value otherwise.

### GetEventPayloadCustomerMappingKeyOk

`func (o *CreateBillingMeterRequest) GetEventPayloadCustomerMappingKeyOk() (*string, bool)`

GetEventPayloadCustomerMappingKeyOk returns a tuple with the EventPayloadCustomerMappingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPayloadCustomerMappingKey

`func (o *CreateBillingMeterRequest) SetEventPayloadCustomerMappingKey(v string)`

SetEventPayloadCustomerMappingKey sets EventPayloadCustomerMappingKey field to given value.

### HasEventPayloadCustomerMappingKey

`func (o *CreateBillingMeterRequest) HasEventPayloadCustomerMappingKey() bool`

HasEventPayloadCustomerMappingKey returns a boolean if a field has been set.

### GetEventPayloadValueKey

`func (o *CreateBillingMeterRequest) GetEventPayloadValueKey() string`

GetEventPayloadValueKey returns the EventPayloadValueKey field if non-nil, zero value otherwise.

### GetEventPayloadValueKeyOk

`func (o *CreateBillingMeterRequest) GetEventPayloadValueKeyOk() (*string, bool)`

GetEventPayloadValueKeyOk returns a tuple with the EventPayloadValueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPayloadValueKey

`func (o *CreateBillingMeterRequest) SetEventPayloadValueKey(v string)`

SetEventPayloadValueKey sets EventPayloadValueKey field to given value.

### HasEventPayloadValueKey

`func (o *CreateBillingMeterRequest) HasEventPayloadValueKey() bool`

HasEventPayloadValueKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


