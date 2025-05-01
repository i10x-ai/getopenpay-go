# CreateBillingMeterEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventDatetime** | Pointer to **time.Time** | The time of the event. | [optional] 
**EventName** | **string** | The name of the meter event. Corresponds with the event_name field on a meter. | 
**Identifier** | Pointer to **string** | A unique identifier for the event. If not provided, one will be generated. We recommend using a globally unique identifier for this. | [optional] 
**Payload** | **map[string]interface{}** | The payload of the event. This must contain the fields corresponding to a meter’s event_payload_customer_mapping_key (default is customer_id) and event_payload_value_key (default is value).  | 

## Methods

### NewCreateBillingMeterEventRequest

`func NewCreateBillingMeterEventRequest(eventName string, payload map[string]interface{}, ) *CreateBillingMeterEventRequest`

NewCreateBillingMeterEventRequest instantiates a new CreateBillingMeterEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBillingMeterEventRequestWithDefaults

`func NewCreateBillingMeterEventRequestWithDefaults() *CreateBillingMeterEventRequest`

NewCreateBillingMeterEventRequestWithDefaults instantiates a new CreateBillingMeterEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventDatetime

`func (o *CreateBillingMeterEventRequest) GetEventDatetime() time.Time`

GetEventDatetime returns the EventDatetime field if non-nil, zero value otherwise.

### GetEventDatetimeOk

`func (o *CreateBillingMeterEventRequest) GetEventDatetimeOk() (*time.Time, bool)`

GetEventDatetimeOk returns a tuple with the EventDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDatetime

`func (o *CreateBillingMeterEventRequest) SetEventDatetime(v time.Time)`

SetEventDatetime sets EventDatetime field to given value.

### HasEventDatetime

`func (o *CreateBillingMeterEventRequest) HasEventDatetime() bool`

HasEventDatetime returns a boolean if a field has been set.

### GetEventName

`func (o *CreateBillingMeterEventRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *CreateBillingMeterEventRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *CreateBillingMeterEventRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetIdentifier

`func (o *CreateBillingMeterEventRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateBillingMeterEventRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateBillingMeterEventRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CreateBillingMeterEventRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetPayload

`func (o *CreateBillingMeterEventRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CreateBillingMeterEventRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CreateBillingMeterEventRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


