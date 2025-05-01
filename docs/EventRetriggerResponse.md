# EventRetriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookDeliveryResults** | [**map[string]WebhookDeliveryResult**](WebhookDeliveryResult.md) | Mapping from webhook id to the delivery results | 

## Methods

### NewEventRetriggerResponse

`func NewEventRetriggerResponse(webhookDeliveryResults map[string]WebhookDeliveryResult, ) *EventRetriggerResponse`

NewEventRetriggerResponse instantiates a new EventRetriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRetriggerResponseWithDefaults

`func NewEventRetriggerResponseWithDefaults() *EventRetriggerResponse`

NewEventRetriggerResponseWithDefaults instantiates a new EventRetriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookDeliveryResults

`func (o *EventRetriggerResponse) GetWebhookDeliveryResults() map[string]WebhookDeliveryResult`

GetWebhookDeliveryResults returns the WebhookDeliveryResults field if non-nil, zero value otherwise.

### GetWebhookDeliveryResultsOk

`func (o *EventRetriggerResponse) GetWebhookDeliveryResultsOk() (*map[string]WebhookDeliveryResult, bool)`

GetWebhookDeliveryResultsOk returns a tuple with the WebhookDeliveryResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookDeliveryResults

`func (o *EventRetriggerResponse) SetWebhookDeliveryResults(v map[string]WebhookDeliveryResult)`

SetWebhookDeliveryResults sets WebhookDeliveryResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


