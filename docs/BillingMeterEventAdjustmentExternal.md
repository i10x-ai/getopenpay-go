# BillingMeterEventAdjustmentExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique identifier for the account. | 
**CancelIdentifier** | **string** | Unique identifier for the event. | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**EventName** | **string** | The name of the meter event. Corresponds with the event_name field on a meter. | 
**Id** | **string** | Unique identifier of meter event. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**Status** | [**BillingMeterEventAdjustmentStatus**](BillingMeterEventAdjustmentStatus.md) | Status of the billing meter adjustment event. | 
**Type** | Pointer to [**BillingMeterAdjustmentType**](BillingMeterAdjustmentType.md) | Specifies whether to cancel a single event. | [optional] 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewBillingMeterEventAdjustmentExternal

`func NewBillingMeterEventAdjustmentExternal(accountId string, cancelIdentifier string, createdAt time.Time, eventName string, id string, status BillingMeterEventAdjustmentStatus, updatedAt time.Time, ) *BillingMeterEventAdjustmentExternal`

NewBillingMeterEventAdjustmentExternal instantiates a new BillingMeterEventAdjustmentExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingMeterEventAdjustmentExternalWithDefaults

`func NewBillingMeterEventAdjustmentExternalWithDefaults() *BillingMeterEventAdjustmentExternal`

NewBillingMeterEventAdjustmentExternalWithDefaults instantiates a new BillingMeterEventAdjustmentExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BillingMeterEventAdjustmentExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BillingMeterEventAdjustmentExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BillingMeterEventAdjustmentExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCancelIdentifier

`func (o *BillingMeterEventAdjustmentExternal) GetCancelIdentifier() string`

GetCancelIdentifier returns the CancelIdentifier field if non-nil, zero value otherwise.

### GetCancelIdentifierOk

`func (o *BillingMeterEventAdjustmentExternal) GetCancelIdentifierOk() (*string, bool)`

GetCancelIdentifierOk returns a tuple with the CancelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelIdentifier

`func (o *BillingMeterEventAdjustmentExternal) SetCancelIdentifier(v string)`

SetCancelIdentifier sets CancelIdentifier field to given value.


### GetCreatedAt

`func (o *BillingMeterEventAdjustmentExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingMeterEventAdjustmentExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingMeterEventAdjustmentExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEventName

`func (o *BillingMeterEventAdjustmentExternal) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *BillingMeterEventAdjustmentExternal) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *BillingMeterEventAdjustmentExternal) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetId

`func (o *BillingMeterEventAdjustmentExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingMeterEventAdjustmentExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingMeterEventAdjustmentExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *BillingMeterEventAdjustmentExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *BillingMeterEventAdjustmentExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *BillingMeterEventAdjustmentExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *BillingMeterEventAdjustmentExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetObject

`func (o *BillingMeterEventAdjustmentExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BillingMeterEventAdjustmentExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BillingMeterEventAdjustmentExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *BillingMeterEventAdjustmentExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetStatus

`func (o *BillingMeterEventAdjustmentExternal) GetStatus() BillingMeterEventAdjustmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingMeterEventAdjustmentExternal) GetStatusOk() (*BillingMeterEventAdjustmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingMeterEventAdjustmentExternal) SetStatus(v BillingMeterEventAdjustmentStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *BillingMeterEventAdjustmentExternal) GetType() BillingMeterAdjustmentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingMeterEventAdjustmentExternal) GetTypeOk() (*BillingMeterAdjustmentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingMeterEventAdjustmentExternal) SetType(v BillingMeterAdjustmentType)`

SetType sets Type field to given value.

### HasType

`func (o *BillingMeterEventAdjustmentExternal) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BillingMeterEventAdjustmentExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingMeterEventAdjustmentExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingMeterEventAdjustmentExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


