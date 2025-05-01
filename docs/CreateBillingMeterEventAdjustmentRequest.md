# CreateBillingMeterEventAdjustmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelIdentifier** | **string** | Unique identifier for the event. | 
**EventName** | **string** | The name of the meter event. Corresponds with the event_name field on a meter. | 
**Type** | Pointer to [**BillingMeterAdjustmentType**](BillingMeterAdjustmentType.md) | Specifies whether to cancel a single event. | [optional] 

## Methods

### NewCreateBillingMeterEventAdjustmentRequest

`func NewCreateBillingMeterEventAdjustmentRequest(cancelIdentifier string, eventName string, ) *CreateBillingMeterEventAdjustmentRequest`

NewCreateBillingMeterEventAdjustmentRequest instantiates a new CreateBillingMeterEventAdjustmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBillingMeterEventAdjustmentRequestWithDefaults

`func NewCreateBillingMeterEventAdjustmentRequestWithDefaults() *CreateBillingMeterEventAdjustmentRequest`

NewCreateBillingMeterEventAdjustmentRequestWithDefaults instantiates a new CreateBillingMeterEventAdjustmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelIdentifier

`func (o *CreateBillingMeterEventAdjustmentRequest) GetCancelIdentifier() string`

GetCancelIdentifier returns the CancelIdentifier field if non-nil, zero value otherwise.

### GetCancelIdentifierOk

`func (o *CreateBillingMeterEventAdjustmentRequest) GetCancelIdentifierOk() (*string, bool)`

GetCancelIdentifierOk returns a tuple with the CancelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelIdentifier

`func (o *CreateBillingMeterEventAdjustmentRequest) SetCancelIdentifier(v string)`

SetCancelIdentifier sets CancelIdentifier field to given value.


### GetEventName

`func (o *CreateBillingMeterEventAdjustmentRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *CreateBillingMeterEventAdjustmentRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *CreateBillingMeterEventAdjustmentRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetType

`func (o *CreateBillingMeterEventAdjustmentRequest) GetType() BillingMeterAdjustmentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateBillingMeterEventAdjustmentRequest) GetTypeOk() (*BillingMeterAdjustmentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateBillingMeterEventAdjustmentRequest) SetType(v BillingMeterAdjustmentType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateBillingMeterEventAdjustmentRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


