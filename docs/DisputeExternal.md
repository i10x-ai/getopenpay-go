# DisputeExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the dispute. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_DISPUTE]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**TheirDisputeId** | **string** | There dispute id. | 
**AccountId** | **string** | Unique identifier for the account. | 
**PaymentIntentId** | **string** | The payment intent id. | 
**TheirPaymentIntentId** | **string** | Their payment intent id | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDisputeExternal

`func NewDisputeExternal(id string, createdAt time.Time, updatedAt time.Time, theirDisputeId string, accountId string, paymentIntentId string, theirPaymentIntentId string, ) *DisputeExternal`

NewDisputeExternal instantiates a new DisputeExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeExternalWithDefaults

`func NewDisputeExternalWithDefaults() *DisputeExternal`

NewDisputeExternalWithDefaults instantiates a new DisputeExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DisputeExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisputeExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisputeExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *DisputeExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DisputeExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DisputeExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *DisputeExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DisputeExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DisputeExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DisputeExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DisputeExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DisputeExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DisputeExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *DisputeExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *DisputeExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *DisputeExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *DisputeExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetTheirDisputeId

`func (o *DisputeExternal) GetTheirDisputeId() string`

GetTheirDisputeId returns the TheirDisputeId field if non-nil, zero value otherwise.

### GetTheirDisputeIdOk

`func (o *DisputeExternal) GetTheirDisputeIdOk() (*string, bool)`

GetTheirDisputeIdOk returns a tuple with the TheirDisputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheirDisputeId

`func (o *DisputeExternal) SetTheirDisputeId(v string)`

SetTheirDisputeId sets TheirDisputeId field to given value.


### GetAccountId

`func (o *DisputeExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DisputeExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DisputeExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPaymentIntentId

`func (o *DisputeExternal) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *DisputeExternal) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *DisputeExternal) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.


### GetTheirPaymentIntentId

`func (o *DisputeExternal) GetTheirPaymentIntentId() string`

GetTheirPaymentIntentId returns the TheirPaymentIntentId field if non-nil, zero value otherwise.

### GetTheirPaymentIntentIdOk

`func (o *DisputeExternal) GetTheirPaymentIntentIdOk() (*string, bool)`

GetTheirPaymentIntentIdOk returns a tuple with the TheirPaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheirPaymentIntentId

`func (o *DisputeExternal) SetTheirPaymentIntentId(v string)`

SetTheirPaymentIntentId sets TheirPaymentIntentId field to given value.


### GetReason

`func (o *DisputeExternal) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DisputeExternal) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DisputeExternal) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DisputeExternal) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *DisputeExternal) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *DisputeExternal) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetMetadata

`func (o *DisputeExternal) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DisputeExternal) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DisputeExternal) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DisputeExternal) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DisputeExternal) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DisputeExternal) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


