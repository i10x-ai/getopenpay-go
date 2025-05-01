# BillingMeterExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique identifier for the account. | 
**AggregationFormula** | Pointer to [**BillingMeterAggregationFormula**](BillingMeterAggregationFormula.md) | Specifies how events are aggregated. Allowed values are count to count the number of events and sum to sum each event’s value. | [optional] 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**DisplayName** | **string** | The meter’s name. | 
**EventName** | **string** | The name of the meter event to record usage for. Corresponds with the event_name field on meter events. | 
**EventPayloadCustomerMappingKey** | Pointer to **string** | The key in the usage event payload to use for mapping the event to a customer. | [optional] [default to "customer_id"]
**EventPayloadValueKey** | Pointer to **string** | The key in the usage event payload to use as the value for this meter. For example, if the event payload contains usage on a bytes_used field, then set the event_payload_value_key to “bytes_used”. | [optional] [default to "value"]
**Id** | **string** | Unique Identifier of the Billing Meter. | 
**IsActive** | **bool** | The meter’s status. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewBillingMeterExternal

`func NewBillingMeterExternal(accountId string, createdAt time.Time, displayName string, eventName string, id string, isActive bool, updatedAt time.Time, ) *BillingMeterExternal`

NewBillingMeterExternal instantiates a new BillingMeterExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingMeterExternalWithDefaults

`func NewBillingMeterExternalWithDefaults() *BillingMeterExternal`

NewBillingMeterExternalWithDefaults instantiates a new BillingMeterExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BillingMeterExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BillingMeterExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BillingMeterExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAggregationFormula

`func (o *BillingMeterExternal) GetAggregationFormula() BillingMeterAggregationFormula`

GetAggregationFormula returns the AggregationFormula field if non-nil, zero value otherwise.

### GetAggregationFormulaOk

`func (o *BillingMeterExternal) GetAggregationFormulaOk() (*BillingMeterAggregationFormula, bool)`

GetAggregationFormulaOk returns a tuple with the AggregationFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFormula

`func (o *BillingMeterExternal) SetAggregationFormula(v BillingMeterAggregationFormula)`

SetAggregationFormula sets AggregationFormula field to given value.

### HasAggregationFormula

`func (o *BillingMeterExternal) HasAggregationFormula() bool`

HasAggregationFormula returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingMeterExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingMeterExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingMeterExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisplayName

`func (o *BillingMeterExternal) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BillingMeterExternal) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BillingMeterExternal) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEventName

`func (o *BillingMeterExternal) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *BillingMeterExternal) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *BillingMeterExternal) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventPayloadCustomerMappingKey

`func (o *BillingMeterExternal) GetEventPayloadCustomerMappingKey() string`

GetEventPayloadCustomerMappingKey returns the EventPayloadCustomerMappingKey field if non-nil, zero value otherwise.

### GetEventPayloadCustomerMappingKeyOk

`func (o *BillingMeterExternal) GetEventPayloadCustomerMappingKeyOk() (*string, bool)`

GetEventPayloadCustomerMappingKeyOk returns a tuple with the EventPayloadCustomerMappingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPayloadCustomerMappingKey

`func (o *BillingMeterExternal) SetEventPayloadCustomerMappingKey(v string)`

SetEventPayloadCustomerMappingKey sets EventPayloadCustomerMappingKey field to given value.

### HasEventPayloadCustomerMappingKey

`func (o *BillingMeterExternal) HasEventPayloadCustomerMappingKey() bool`

HasEventPayloadCustomerMappingKey returns a boolean if a field has been set.

### GetEventPayloadValueKey

`func (o *BillingMeterExternal) GetEventPayloadValueKey() string`

GetEventPayloadValueKey returns the EventPayloadValueKey field if non-nil, zero value otherwise.

### GetEventPayloadValueKeyOk

`func (o *BillingMeterExternal) GetEventPayloadValueKeyOk() (*string, bool)`

GetEventPayloadValueKeyOk returns a tuple with the EventPayloadValueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPayloadValueKey

`func (o *BillingMeterExternal) SetEventPayloadValueKey(v string)`

SetEventPayloadValueKey sets EventPayloadValueKey field to given value.

### HasEventPayloadValueKey

`func (o *BillingMeterExternal) HasEventPayloadValueKey() bool`

HasEventPayloadValueKey returns a boolean if a field has been set.

### GetId

`func (o *BillingMeterExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingMeterExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingMeterExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIsActive

`func (o *BillingMeterExternal) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BillingMeterExternal) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BillingMeterExternal) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsDeleted

`func (o *BillingMeterExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *BillingMeterExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *BillingMeterExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *BillingMeterExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetObject

`func (o *BillingMeterExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BillingMeterExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BillingMeterExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *BillingMeterExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BillingMeterExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingMeterExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingMeterExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


