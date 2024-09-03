# PortalSessionExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the portal session. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_PORTAL_SESSION]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**AccountId** | **string** | Unique identifier for the account. | 
**CustomerId** | **string** | The ID of an existing customer. | 
**ReturnUrl** | **string** | The default URL to redirect customers to when they click on the portal&#39;s link to return to your website. | 
**Url** | **string** | The short-lived URL of the session that gives customers access to the customer portal. | 
**Token** | **string** | The short-lived token of the session that gives customers access to the customer portal. | 

## Methods

### NewPortalSessionExternal

`func NewPortalSessionExternal(id string, createdAt time.Time, updatedAt time.Time, accountId string, customerId string, returnUrl string, url string, token string, ) *PortalSessionExternal`

NewPortalSessionExternal instantiates a new PortalSessionExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalSessionExternalWithDefaults

`func NewPortalSessionExternalWithDefaults() *PortalSessionExternal`

NewPortalSessionExternalWithDefaults instantiates a new PortalSessionExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalSessionExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalSessionExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalSessionExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PortalSessionExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PortalSessionExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PortalSessionExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *PortalSessionExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortalSessionExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortalSessionExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortalSessionExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PortalSessionExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortalSessionExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortalSessionExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *PortalSessionExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PortalSessionExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PortalSessionExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PortalSessionExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetAccountId

`func (o *PortalSessionExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PortalSessionExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PortalSessionExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCustomerId

`func (o *PortalSessionExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PortalSessionExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PortalSessionExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetReturnUrl

`func (o *PortalSessionExternal) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PortalSessionExternal) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PortalSessionExternal) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetUrl

`func (o *PortalSessionExternal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PortalSessionExternal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PortalSessionExternal) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetToken

`func (o *PortalSessionExternal) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PortalSessionExternal) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PortalSessionExternal) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


