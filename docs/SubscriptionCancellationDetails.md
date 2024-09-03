# SubscriptionCancellationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **NullableString** |  | [optional] 
**Feedback** | Pointer to [**NullableSubscriptionCancelFeedbackEnum**](SubscriptionCancelFeedbackEnum.md) |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubscriptionCancellationDetails

`func NewSubscriptionCancellationDetails() *SubscriptionCancellationDetails`

NewSubscriptionCancellationDetails instantiates a new SubscriptionCancellationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCancellationDetailsWithDefaults

`func NewSubscriptionCancellationDetailsWithDefaults() *SubscriptionCancellationDetails`

NewSubscriptionCancellationDetailsWithDefaults instantiates a new SubscriptionCancellationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SubscriptionCancellationDetails) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SubscriptionCancellationDetails) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SubscriptionCancellationDetails) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SubscriptionCancellationDetails) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *SubscriptionCancellationDetails) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *SubscriptionCancellationDetails) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetFeedback

`func (o *SubscriptionCancellationDetails) GetFeedback() SubscriptionCancelFeedbackEnum`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *SubscriptionCancellationDetails) GetFeedbackOk() (*SubscriptionCancelFeedbackEnum, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *SubscriptionCancellationDetails) SetFeedback(v SubscriptionCancelFeedbackEnum)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *SubscriptionCancellationDetails) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### SetFeedbackNil

`func (o *SubscriptionCancellationDetails) SetFeedbackNil(b bool)`

 SetFeedbackNil sets the value for Feedback to be an explicit nil

### UnsetFeedback
`func (o *SubscriptionCancellationDetails) UnsetFeedback()`

UnsetFeedback ensures that no value is present for Feedback, not even an explicit nil
### GetReason

`func (o *SubscriptionCancellationDetails) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SubscriptionCancellationDetails) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SubscriptionCancellationDetails) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SubscriptionCancellationDetails) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *SubscriptionCancellationDetails) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *SubscriptionCancellationDetails) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


