# EventRetriggerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique id for the event which you would like to re-retigger. | 
**SelectedWebhooks** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEventRetriggerRequest

`func NewEventRetriggerRequest(eventId string, ) *EventRetriggerRequest`

NewEventRetriggerRequest instantiates a new EventRetriggerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRetriggerRequestWithDefaults

`func NewEventRetriggerRequestWithDefaults() *EventRetriggerRequest`

NewEventRetriggerRequestWithDefaults instantiates a new EventRetriggerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *EventRetriggerRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventRetriggerRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventRetriggerRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetSelectedWebhooks

`func (o *EventRetriggerRequest) GetSelectedWebhooks() []string`

GetSelectedWebhooks returns the SelectedWebhooks field if non-nil, zero value otherwise.

### GetSelectedWebhooksOk

`func (o *EventRetriggerRequest) GetSelectedWebhooksOk() (*[]string, bool)`

GetSelectedWebhooksOk returns a tuple with the SelectedWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedWebhooks

`func (o *EventRetriggerRequest) SetSelectedWebhooks(v []string)`

SetSelectedWebhooks sets SelectedWebhooks field to given value.

### HasSelectedWebhooks

`func (o *EventRetriggerRequest) HasSelectedWebhooks() bool`

HasSelectedWebhooks returns a boolean if a field has been set.

### SetSelectedWebhooksNil

`func (o *EventRetriggerRequest) SetSelectedWebhooksNil(b bool)`

 SetSelectedWebhooksNil sets the value for SelectedWebhooks to be an explicit nil

### UnsetSelectedWebhooks
`func (o *EventRetriggerRequest) UnsetSelectedWebhooks()`

UnsetSelectedWebhooks ensures that no value is present for SelectedWebhooks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


