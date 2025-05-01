# CreateSubscriptionItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddAtPeriodEnd** | Pointer to **bool** | If the flag is set to True, item will be added when renewing the subscription at next billing cycle. | [optional] [default to false]
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PriceId** | **string** | The ID of the price. | 
**ProrationBehavior** | Pointer to [**ProrationEnum**](ProrationEnum.md) | Determines how to handle prorations when the billable items changes.In case of subscription is in trialing state, invoice items if any will be for amount_atom 0. | [optional] 
**Quantity** | Pointer to **int32** | The quantity you&#39;d like to apply to the subscription item you&#39;re creating. | [optional] [default to 1]
**SubscriptionId** | **string** | The identifier of the subscription to modify | 

## Methods

### NewCreateSubscriptionItemRequest

`func NewCreateSubscriptionItemRequest(priceId string, subscriptionId string, ) *CreateSubscriptionItemRequest`

NewCreateSubscriptionItemRequest instantiates a new CreateSubscriptionItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionItemRequestWithDefaults

`func NewCreateSubscriptionItemRequestWithDefaults() *CreateSubscriptionItemRequest`

NewCreateSubscriptionItemRequestWithDefaults instantiates a new CreateSubscriptionItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddAtPeriodEnd

`func (o *CreateSubscriptionItemRequest) GetAddAtPeriodEnd() bool`

GetAddAtPeriodEnd returns the AddAtPeriodEnd field if non-nil, zero value otherwise.

### GetAddAtPeriodEndOk

`func (o *CreateSubscriptionItemRequest) GetAddAtPeriodEndOk() (*bool, bool)`

GetAddAtPeriodEndOk returns a tuple with the AddAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAtPeriodEnd

`func (o *CreateSubscriptionItemRequest) SetAddAtPeriodEnd(v bool)`

SetAddAtPeriodEnd sets AddAtPeriodEnd field to given value.

### HasAddAtPeriodEnd

`func (o *CreateSubscriptionItemRequest) HasAddAtPeriodEnd() bool`

HasAddAtPeriodEnd returns a boolean if a field has been set.

### GetCustomFields

`func (o *CreateSubscriptionItemRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateSubscriptionItemRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateSubscriptionItemRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateSubscriptionItemRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CreateSubscriptionItemRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CreateSubscriptionItemRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDescription

`func (o *CreateSubscriptionItemRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSubscriptionItemRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSubscriptionItemRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSubscriptionItemRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSubscriptionItemRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSubscriptionItemRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriceId

`func (o *CreateSubscriptionItemRequest) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *CreateSubscriptionItemRequest) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *CreateSubscriptionItemRequest) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetProrationBehavior

`func (o *CreateSubscriptionItemRequest) GetProrationBehavior() ProrationEnum`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *CreateSubscriptionItemRequest) GetProrationBehaviorOk() (*ProrationEnum, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *CreateSubscriptionItemRequest) SetProrationBehavior(v ProrationEnum)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *CreateSubscriptionItemRequest) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.

### GetQuantity

`func (o *CreateSubscriptionItemRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateSubscriptionItemRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateSubscriptionItemRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreateSubscriptionItemRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CreateSubscriptionItemRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateSubscriptionItemRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateSubscriptionItemRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


