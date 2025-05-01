# WebhookDeliveryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptTime** | **time.Time** | Webhook delivery attempt time in &#39;ISO 8601&#39; format. | 
**EventId** | **string** | Unique id for the event which was sent to the webhook url. | 
**ResponseCode** | **int32** | HTTP response code received upon attempted delivery to webhook endpoint | 
**ResponseText** | Pointer to **NullableString** |  | [optional] 
**WebhookId** | **string** | Unique id for the webhook url to which the event was sent. | 

## Methods

### NewWebhookDeliveryResult

`func NewWebhookDeliveryResult(attemptTime time.Time, eventId string, responseCode int32, webhookId string, ) *WebhookDeliveryResult`

NewWebhookDeliveryResult instantiates a new WebhookDeliveryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryResultWithDefaults

`func NewWebhookDeliveryResultWithDefaults() *WebhookDeliveryResult`

NewWebhookDeliveryResultWithDefaults instantiates a new WebhookDeliveryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptTime

`func (o *WebhookDeliveryResult) GetAttemptTime() time.Time`

GetAttemptTime returns the AttemptTime field if non-nil, zero value otherwise.

### GetAttemptTimeOk

`func (o *WebhookDeliveryResult) GetAttemptTimeOk() (*time.Time, bool)`

GetAttemptTimeOk returns a tuple with the AttemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptTime

`func (o *WebhookDeliveryResult) SetAttemptTime(v time.Time)`

SetAttemptTime sets AttemptTime field to given value.


### GetEventId

`func (o *WebhookDeliveryResult) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhookDeliveryResult) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhookDeliveryResult) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetResponseCode

`func (o *WebhookDeliveryResult) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *WebhookDeliveryResult) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *WebhookDeliveryResult) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseText

`func (o *WebhookDeliveryResult) GetResponseText() string`

GetResponseText returns the ResponseText field if non-nil, zero value otherwise.

### GetResponseTextOk

`func (o *WebhookDeliveryResult) GetResponseTextOk() (*string, bool)`

GetResponseTextOk returns a tuple with the ResponseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseText

`func (o *WebhookDeliveryResult) SetResponseText(v string)`

SetResponseText sets ResponseText field to given value.

### HasResponseText

`func (o *WebhookDeliveryResult) HasResponseText() bool`

HasResponseText returns a boolean if a field has been set.

### SetResponseTextNil

`func (o *WebhookDeliveryResult) SetResponseTextNil(b bool)`

 SetResponseTextNil sets the value for ResponseText to be an explicit nil

### UnsetResponseText
`func (o *WebhookDeliveryResult) UnsetResponseText()`

UnsetResponseText ensures that no value is present for ResponseText, not even an explicit nil
### GetWebhookId

`func (o *WebhookDeliveryResult) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookDeliveryResult) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookDeliveryResult) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


