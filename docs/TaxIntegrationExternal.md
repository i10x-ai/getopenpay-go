# TaxIntegrationExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Identifier of the tax integration. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_TAX_INTEGRATION]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**ApiName** | [**TaxIntegrationApiName**](TaxIntegrationApiName.md) |  | 
**ApiKeys** | **map[string]interface{}** | These keys will allow to authenticate API requests to the tax processor. | 
**AccountingTz** | **string** | The timezone name for the accounting timezone (see pytz.all_timezones for a full list) | 

## Methods

### NewTaxIntegrationExternal

`func NewTaxIntegrationExternal(id string, createdAt time.Time, updatedAt time.Time, apiName TaxIntegrationApiName, apiKeys map[string]interface{}, accountingTz string, ) *TaxIntegrationExternal`

NewTaxIntegrationExternal instantiates a new TaxIntegrationExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxIntegrationExternalWithDefaults

`func NewTaxIntegrationExternalWithDefaults() *TaxIntegrationExternal`

NewTaxIntegrationExternalWithDefaults instantiates a new TaxIntegrationExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxIntegrationExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxIntegrationExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxIntegrationExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *TaxIntegrationExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TaxIntegrationExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TaxIntegrationExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *TaxIntegrationExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaxIntegrationExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaxIntegrationExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaxIntegrationExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TaxIntegrationExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaxIntegrationExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaxIntegrationExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *TaxIntegrationExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TaxIntegrationExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TaxIntegrationExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TaxIntegrationExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetApiName

`func (o *TaxIntegrationExternal) GetApiName() TaxIntegrationApiName`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *TaxIntegrationExternal) GetApiNameOk() (*TaxIntegrationApiName, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *TaxIntegrationExternal) SetApiName(v TaxIntegrationApiName)`

SetApiName sets ApiName field to given value.


### GetApiKeys

`func (o *TaxIntegrationExternal) GetApiKeys() map[string]interface{}`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *TaxIntegrationExternal) GetApiKeysOk() (*map[string]interface{}, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *TaxIntegrationExternal) SetApiKeys(v map[string]interface{})`

SetApiKeys sets ApiKeys field to given value.


### GetAccountingTz

`func (o *TaxIntegrationExternal) GetAccountingTz() string`

GetAccountingTz returns the AccountingTz field if non-nil, zero value otherwise.

### GetAccountingTzOk

`func (o *TaxIntegrationExternal) GetAccountingTzOk() (*string, bool)`

GetAccountingTzOk returns a tuple with the AccountingTz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingTz

`func (o *TaxIntegrationExternal) SetAccountingTz(v string)`

SetAccountingTz sets AccountingTz field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


