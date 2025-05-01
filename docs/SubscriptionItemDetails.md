# SubscriptionItemDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubscriptionItemDetails

`func NewSubscriptionItemDetails() *SubscriptionItemDetails`

NewSubscriptionItemDetails instantiates a new SubscriptionItemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionItemDetailsWithDefaults

`func NewSubscriptionItemDetailsWithDefaults() *SubscriptionItemDetails`

NewSubscriptionItemDetailsWithDefaults instantiates a new SubscriptionItemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFields

`func (o *SubscriptionItemDetails) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SubscriptionItemDetails) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SubscriptionItemDetails) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SubscriptionItemDetails) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *SubscriptionItemDetails) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *SubscriptionItemDetails) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDescription

`func (o *SubscriptionItemDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionItemDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionItemDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionItemDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SubscriptionItemDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SubscriptionItemDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


