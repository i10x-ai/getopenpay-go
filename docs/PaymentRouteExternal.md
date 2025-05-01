# PaymentRouteExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique Identifier of the account. | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Id** | **string** | Unique Identifier of the payment route. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Name** | **string** | The name of the payment route. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**RouteConfig** | **map[string]interface{}** | The configuration object for the payment route. | 
**RouteType** | [**PaymentRouteType**](PaymentRouteType.md) | The type of payment route. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewPaymentRouteExternal

`func NewPaymentRouteExternal(accountId string, createdAt time.Time, id string, name string, routeConfig map[string]interface{}, routeType PaymentRouteType, updatedAt time.Time, ) *PaymentRouteExternal`

NewPaymentRouteExternal instantiates a new PaymentRouteExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRouteExternalWithDefaults

`func NewPaymentRouteExternalWithDefaults() *PaymentRouteExternal`

NewPaymentRouteExternalWithDefaults instantiates a new PaymentRouteExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PaymentRouteExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PaymentRouteExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PaymentRouteExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *PaymentRouteExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentRouteExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentRouteExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *PaymentRouteExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRouteExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRouteExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *PaymentRouteExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PaymentRouteExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PaymentRouteExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PaymentRouteExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetName

`func (o *PaymentRouteExternal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentRouteExternal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentRouteExternal) SetName(v string)`

SetName sets Name field to given value.


### GetObject

`func (o *PaymentRouteExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentRouteExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentRouteExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *PaymentRouteExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRouteConfig

`func (o *PaymentRouteExternal) GetRouteConfig() map[string]interface{}`

GetRouteConfig returns the RouteConfig field if non-nil, zero value otherwise.

### GetRouteConfigOk

`func (o *PaymentRouteExternal) GetRouteConfigOk() (*map[string]interface{}, bool)`

GetRouteConfigOk returns a tuple with the RouteConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteConfig

`func (o *PaymentRouteExternal) SetRouteConfig(v map[string]interface{})`

SetRouteConfig sets RouteConfig field to given value.


### GetRouteType

`func (o *PaymentRouteExternal) GetRouteType() PaymentRouteType`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *PaymentRouteExternal) GetRouteTypeOk() (*PaymentRouteType, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *PaymentRouteExternal) SetRouteType(v PaymentRouteType)`

SetRouteType sets RouteType field to given value.


### GetUpdatedAt

`func (o *PaymentRouteExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentRouteExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentRouteExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


