# BillingMeterEventExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique identifier for the account. | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**EventDatetime** | Pointer to **time.Time** | The time of the event. | [optional] 
**EventName** | **string** | The name of the meter event. Corresponds with the event_name field on a meter. | 
**Id** | **string** | Unique identifier of meter event. | 
**Identifier** | Pointer to **string** | A unique identifier for the event. If not provided, one will be generated. We recommend using a globally unique identifier for this. | [optional] 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**Payload** | **map[string]interface{}** | The payload of the event. This must contain the fields corresponding to a meter’s event_payload_customer_mapping_key (default is customer_id) and event_payload_value_key (default is value).  | 
**Status** | [**BillingMeterEventStatus**](BillingMeterEventStatus.md) | Status of the billing meter event. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewBillingMeterEventExternal

`func NewBillingMeterEventExternal(accountId string, createdAt time.Time, eventName string, id string, payload map[string]interface{}, status BillingMeterEventStatus, updatedAt time.Time, ) *BillingMeterEventExternal`

NewBillingMeterEventExternal instantiates a new BillingMeterEventExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingMeterEventExternalWithDefaults

`func NewBillingMeterEventExternalWithDefaults() *BillingMeterEventExternal`

NewBillingMeterEventExternalWithDefaults instantiates a new BillingMeterEventExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BillingMeterEventExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BillingMeterEventExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BillingMeterEventExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *BillingMeterEventExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingMeterEventExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingMeterEventExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEventDatetime

`func (o *BillingMeterEventExternal) GetEventDatetime() time.Time`

GetEventDatetime returns the EventDatetime field if non-nil, zero value otherwise.

### GetEventDatetimeOk

`func (o *BillingMeterEventExternal) GetEventDatetimeOk() (*time.Time, bool)`

GetEventDatetimeOk returns a tuple with the EventDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDatetime

`func (o *BillingMeterEventExternal) SetEventDatetime(v time.Time)`

SetEventDatetime sets EventDatetime field to given value.

### HasEventDatetime

`func (o *BillingMeterEventExternal) HasEventDatetime() bool`

HasEventDatetime returns a boolean if a field has been set.

### GetEventName

`func (o *BillingMeterEventExternal) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *BillingMeterEventExternal) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *BillingMeterEventExternal) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetId

`func (o *BillingMeterEventExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingMeterEventExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingMeterEventExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *BillingMeterEventExternal) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BillingMeterEventExternal) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BillingMeterEventExternal) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BillingMeterEventExternal) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIsDeleted

`func (o *BillingMeterEventExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *BillingMeterEventExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *BillingMeterEventExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *BillingMeterEventExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetObject

`func (o *BillingMeterEventExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BillingMeterEventExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BillingMeterEventExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *BillingMeterEventExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPayload

`func (o *BillingMeterEventExternal) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *BillingMeterEventExternal) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *BillingMeterEventExternal) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### GetStatus

`func (o *BillingMeterEventExternal) GetStatus() BillingMeterEventStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingMeterEventExternal) GetStatusOk() (*BillingMeterEventStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingMeterEventExternal) SetStatus(v BillingMeterEventStatus)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *BillingMeterEventExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingMeterEventExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingMeterEventExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


