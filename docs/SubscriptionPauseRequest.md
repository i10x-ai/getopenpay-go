# SubscriptionPauseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPreview** | Pointer to **bool** | Whether the request is in preview mode (Subscription won&#39;t actually pause) | [optional] [default to false]
**NumberOfBillingCyclesToSkip** | Pointer to **NullableInt32** |  | [optional] 
**ResumptionDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewSubscriptionPauseRequest

`func NewSubscriptionPauseRequest() *SubscriptionPauseRequest`

NewSubscriptionPauseRequest instantiates a new SubscriptionPauseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPauseRequestWithDefaults

`func NewSubscriptionPauseRequestWithDefaults() *SubscriptionPauseRequest`

NewSubscriptionPauseRequestWithDefaults instantiates a new SubscriptionPauseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPreview

`func (o *SubscriptionPauseRequest) GetIsPreview() bool`

GetIsPreview returns the IsPreview field if non-nil, zero value otherwise.

### GetIsPreviewOk

`func (o *SubscriptionPauseRequest) GetIsPreviewOk() (*bool, bool)`

GetIsPreviewOk returns a tuple with the IsPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreview

`func (o *SubscriptionPauseRequest) SetIsPreview(v bool)`

SetIsPreview sets IsPreview field to given value.

### HasIsPreview

`func (o *SubscriptionPauseRequest) HasIsPreview() bool`

HasIsPreview returns a boolean if a field has been set.

### GetNumberOfBillingCyclesToSkip

`func (o *SubscriptionPauseRequest) GetNumberOfBillingCyclesToSkip() int32`

GetNumberOfBillingCyclesToSkip returns the NumberOfBillingCyclesToSkip field if non-nil, zero value otherwise.

### GetNumberOfBillingCyclesToSkipOk

`func (o *SubscriptionPauseRequest) GetNumberOfBillingCyclesToSkipOk() (*int32, bool)`

GetNumberOfBillingCyclesToSkipOk returns a tuple with the NumberOfBillingCyclesToSkip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBillingCyclesToSkip

`func (o *SubscriptionPauseRequest) SetNumberOfBillingCyclesToSkip(v int32)`

SetNumberOfBillingCyclesToSkip sets NumberOfBillingCyclesToSkip field to given value.

### HasNumberOfBillingCyclesToSkip

`func (o *SubscriptionPauseRequest) HasNumberOfBillingCyclesToSkip() bool`

HasNumberOfBillingCyclesToSkip returns a boolean if a field has been set.

### SetNumberOfBillingCyclesToSkipNil

`func (o *SubscriptionPauseRequest) SetNumberOfBillingCyclesToSkipNil(b bool)`

 SetNumberOfBillingCyclesToSkipNil sets the value for NumberOfBillingCyclesToSkip to be an explicit nil

### UnsetNumberOfBillingCyclesToSkip
`func (o *SubscriptionPauseRequest) UnsetNumberOfBillingCyclesToSkip()`

UnsetNumberOfBillingCyclesToSkip ensures that no value is present for NumberOfBillingCyclesToSkip, not even an explicit nil
### GetResumptionDate

`func (o *SubscriptionPauseRequest) GetResumptionDate() time.Time`

GetResumptionDate returns the ResumptionDate field if non-nil, zero value otherwise.

### GetResumptionDateOk

`func (o *SubscriptionPauseRequest) GetResumptionDateOk() (*time.Time, bool)`

GetResumptionDateOk returns a tuple with the ResumptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumptionDate

`func (o *SubscriptionPauseRequest) SetResumptionDate(v time.Time)`

SetResumptionDate sets ResumptionDate field to given value.

### HasResumptionDate

`func (o *SubscriptionPauseRequest) HasResumptionDate() bool`

HasResumptionDate returns a boolean if a field has been set.

### SetResumptionDateNil

`func (o *SubscriptionPauseRequest) SetResumptionDateNil(b bool)`

 SetResumptionDateNil sets the value for ResumptionDate to be an explicit nil

### UnsetResumptionDate
`func (o *SubscriptionPauseRequest) UnsetResumptionDate()`

UnsetResumptionDate ensures that no value is present for ResumptionDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


