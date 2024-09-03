# CreateDisputeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentIntentId** | Pointer to **NullableString** |  | [optional] 
**TheirPaymentIntentId** | Pointer to **NullableString** |  | [optional] 
**TheirDisputeId** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateDisputeRequest

`func NewCreateDisputeRequest() *CreateDisputeRequest`

NewCreateDisputeRequest instantiates a new CreateDisputeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDisputeRequestWithDefaults

`func NewCreateDisputeRequestWithDefaults() *CreateDisputeRequest`

NewCreateDisputeRequestWithDefaults instantiates a new CreateDisputeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentIntentId

`func (o *CreateDisputeRequest) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *CreateDisputeRequest) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *CreateDisputeRequest) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.

### HasPaymentIntentId

`func (o *CreateDisputeRequest) HasPaymentIntentId() bool`

HasPaymentIntentId returns a boolean if a field has been set.

### SetPaymentIntentIdNil

`func (o *CreateDisputeRequest) SetPaymentIntentIdNil(b bool)`

 SetPaymentIntentIdNil sets the value for PaymentIntentId to be an explicit nil

### UnsetPaymentIntentId
`func (o *CreateDisputeRequest) UnsetPaymentIntentId()`

UnsetPaymentIntentId ensures that no value is present for PaymentIntentId, not even an explicit nil
### GetTheirPaymentIntentId

`func (o *CreateDisputeRequest) GetTheirPaymentIntentId() string`

GetTheirPaymentIntentId returns the TheirPaymentIntentId field if non-nil, zero value otherwise.

### GetTheirPaymentIntentIdOk

`func (o *CreateDisputeRequest) GetTheirPaymentIntentIdOk() (*string, bool)`

GetTheirPaymentIntentIdOk returns a tuple with the TheirPaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheirPaymentIntentId

`func (o *CreateDisputeRequest) SetTheirPaymentIntentId(v string)`

SetTheirPaymentIntentId sets TheirPaymentIntentId field to given value.

### HasTheirPaymentIntentId

`func (o *CreateDisputeRequest) HasTheirPaymentIntentId() bool`

HasTheirPaymentIntentId returns a boolean if a field has been set.

### SetTheirPaymentIntentIdNil

`func (o *CreateDisputeRequest) SetTheirPaymentIntentIdNil(b bool)`

 SetTheirPaymentIntentIdNil sets the value for TheirPaymentIntentId to be an explicit nil

### UnsetTheirPaymentIntentId
`func (o *CreateDisputeRequest) UnsetTheirPaymentIntentId()`

UnsetTheirPaymentIntentId ensures that no value is present for TheirPaymentIntentId, not even an explicit nil
### GetTheirDisputeId

`func (o *CreateDisputeRequest) GetTheirDisputeId() string`

GetTheirDisputeId returns the TheirDisputeId field if non-nil, zero value otherwise.

### GetTheirDisputeIdOk

`func (o *CreateDisputeRequest) GetTheirDisputeIdOk() (*string, bool)`

GetTheirDisputeIdOk returns a tuple with the TheirDisputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheirDisputeId

`func (o *CreateDisputeRequest) SetTheirDisputeId(v string)`

SetTheirDisputeId sets TheirDisputeId field to given value.

### HasTheirDisputeId

`func (o *CreateDisputeRequest) HasTheirDisputeId() bool`

HasTheirDisputeId returns a boolean if a field has been set.

### SetTheirDisputeIdNil

`func (o *CreateDisputeRequest) SetTheirDisputeIdNil(b bool)`

 SetTheirDisputeIdNil sets the value for TheirDisputeId to be an explicit nil

### UnsetTheirDisputeId
`func (o *CreateDisputeRequest) UnsetTheirDisputeId()`

UnsetTheirDisputeId ensures that no value is present for TheirDisputeId, not even an explicit nil
### GetReason

`func (o *CreateDisputeRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateDisputeRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateDisputeRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateDisputeRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *CreateDisputeRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *CreateDisputeRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetMeta

`func (o *CreateDisputeRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateDisputeRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateDisputeRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateDisputeRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *CreateDisputeRequest) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *CreateDisputeRequest) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


