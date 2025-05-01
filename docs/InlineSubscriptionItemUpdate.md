# InlineSubscriptionItemUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddAtPeriodEnd** | Pointer to **NullableBool** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Delete** | Pointer to **bool** | A flag that, if set to true, will delete the specified item immediately from subscription. drop_at_end flag cannot be used while using this behaviour | [optional] [default to false]
**Description** | Pointer to **NullableString** |  | [optional] 
**DropAtEnd** | Pointer to **NullableBool** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**PriceId** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **NullableInt32** |  | [optional] 
**StartsAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewInlineSubscriptionItemUpdate

`func NewInlineSubscriptionItemUpdate() *InlineSubscriptionItemUpdate`

NewInlineSubscriptionItemUpdate instantiates a new InlineSubscriptionItemUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineSubscriptionItemUpdateWithDefaults

`func NewInlineSubscriptionItemUpdateWithDefaults() *InlineSubscriptionItemUpdate`

NewInlineSubscriptionItemUpdateWithDefaults instantiates a new InlineSubscriptionItemUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddAtPeriodEnd

`func (o *InlineSubscriptionItemUpdate) GetAddAtPeriodEnd() bool`

GetAddAtPeriodEnd returns the AddAtPeriodEnd field if non-nil, zero value otherwise.

### GetAddAtPeriodEndOk

`func (o *InlineSubscriptionItemUpdate) GetAddAtPeriodEndOk() (*bool, bool)`

GetAddAtPeriodEndOk returns a tuple with the AddAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAtPeriodEnd

`func (o *InlineSubscriptionItemUpdate) SetAddAtPeriodEnd(v bool)`

SetAddAtPeriodEnd sets AddAtPeriodEnd field to given value.

### HasAddAtPeriodEnd

`func (o *InlineSubscriptionItemUpdate) HasAddAtPeriodEnd() bool`

HasAddAtPeriodEnd returns a boolean if a field has been set.

### SetAddAtPeriodEndNil

`func (o *InlineSubscriptionItemUpdate) SetAddAtPeriodEndNil(b bool)`

 SetAddAtPeriodEndNil sets the value for AddAtPeriodEnd to be an explicit nil

### UnsetAddAtPeriodEnd
`func (o *InlineSubscriptionItemUpdate) UnsetAddAtPeriodEnd()`

UnsetAddAtPeriodEnd ensures that no value is present for AddAtPeriodEnd, not even an explicit nil
### GetCustomFields

`func (o *InlineSubscriptionItemUpdate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InlineSubscriptionItemUpdate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InlineSubscriptionItemUpdate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InlineSubscriptionItemUpdate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *InlineSubscriptionItemUpdate) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *InlineSubscriptionItemUpdate) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDelete

`func (o *InlineSubscriptionItemUpdate) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *InlineSubscriptionItemUpdate) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *InlineSubscriptionItemUpdate) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *InlineSubscriptionItemUpdate) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDescription

`func (o *InlineSubscriptionItemUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineSubscriptionItemUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineSubscriptionItemUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineSubscriptionItemUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineSubscriptionItemUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineSubscriptionItemUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDropAtEnd

`func (o *InlineSubscriptionItemUpdate) GetDropAtEnd() bool`

GetDropAtEnd returns the DropAtEnd field if non-nil, zero value otherwise.

### GetDropAtEndOk

`func (o *InlineSubscriptionItemUpdate) GetDropAtEndOk() (*bool, bool)`

GetDropAtEndOk returns a tuple with the DropAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropAtEnd

`func (o *InlineSubscriptionItemUpdate) SetDropAtEnd(v bool)`

SetDropAtEnd sets DropAtEnd field to given value.

### HasDropAtEnd

`func (o *InlineSubscriptionItemUpdate) HasDropAtEnd() bool`

HasDropAtEnd returns a boolean if a field has been set.

### SetDropAtEndNil

`func (o *InlineSubscriptionItemUpdate) SetDropAtEndNil(b bool)`

 SetDropAtEndNil sets the value for DropAtEnd to be an explicit nil

### UnsetDropAtEnd
`func (o *InlineSubscriptionItemUpdate) UnsetDropAtEnd()`

UnsetDropAtEnd ensures that no value is present for DropAtEnd, not even an explicit nil
### GetId

`func (o *InlineSubscriptionItemUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineSubscriptionItemUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineSubscriptionItemUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineSubscriptionItemUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InlineSubscriptionItemUpdate) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InlineSubscriptionItemUpdate) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPriceId

`func (o *InlineSubscriptionItemUpdate) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *InlineSubscriptionItemUpdate) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *InlineSubscriptionItemUpdate) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *InlineSubscriptionItemUpdate) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### SetPriceIdNil

`func (o *InlineSubscriptionItemUpdate) SetPriceIdNil(b bool)`

 SetPriceIdNil sets the value for PriceId to be an explicit nil

### UnsetPriceId
`func (o *InlineSubscriptionItemUpdate) UnsetPriceId()`

UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
### GetQuantity

`func (o *InlineSubscriptionItemUpdate) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InlineSubscriptionItemUpdate) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InlineSubscriptionItemUpdate) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InlineSubscriptionItemUpdate) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *InlineSubscriptionItemUpdate) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *InlineSubscriptionItemUpdate) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetStartsAt

`func (o *InlineSubscriptionItemUpdate) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *InlineSubscriptionItemUpdate) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *InlineSubscriptionItemUpdate) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *InlineSubscriptionItemUpdate) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *InlineSubscriptionItemUpdate) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *InlineSubscriptionItemUpdate) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


