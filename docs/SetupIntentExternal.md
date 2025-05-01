# SetupIntentExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique Identifier of the account. | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**CustomerId** | **NullableString** |  | 
**Id** | **string** | Unique Identifier of the setup intent. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**PaymentMethodIds** | **[]string** | Payment methods created from this setup intent. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewSetupIntentExternal

`func NewSetupIntentExternal(accountId string, createdAt time.Time, customerId NullableString, id string, paymentMethodIds []string, updatedAt time.Time, ) *SetupIntentExternal`

NewSetupIntentExternal instantiates a new SetupIntentExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupIntentExternalWithDefaults

`func NewSetupIntentExternalWithDefaults() *SetupIntentExternal`

NewSetupIntentExternalWithDefaults instantiates a new SetupIntentExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SetupIntentExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SetupIntentExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SetupIntentExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *SetupIntentExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SetupIntentExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SetupIntentExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCustomerId

`func (o *SetupIntentExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SetupIntentExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SetupIntentExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *SetupIntentExternal) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *SetupIntentExternal) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetId

`func (o *SetupIntentExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetupIntentExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetupIntentExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *SetupIntentExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SetupIntentExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SetupIntentExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *SetupIntentExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetObject

`func (o *SetupIntentExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SetupIntentExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SetupIntentExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *SetupIntentExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPaymentMethodIds

`func (o *SetupIntentExternal) GetPaymentMethodIds() []string`

GetPaymentMethodIds returns the PaymentMethodIds field if non-nil, zero value otherwise.

### GetPaymentMethodIdsOk

`func (o *SetupIntentExternal) GetPaymentMethodIdsOk() (*[]string, bool)`

GetPaymentMethodIdsOk returns a tuple with the PaymentMethodIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodIds

`func (o *SetupIntentExternal) SetPaymentMethodIds(v []string)`

SetPaymentMethodIds sets PaymentMethodIds field to given value.


### GetUpdatedAt

`func (o *SetupIntentExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SetupIntentExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SetupIntentExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


