# NonPciIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account ID | 
**IntegrationId** | **string** | The integration ID | 
**IntegrationType** | [**NonPciIntegrationEnum**](NonPciIntegrationEnum.md) |  | 
**Configuration** | Pointer to [**[]NonPciIntegrationConfigurationInner**](NonPciIntegrationConfigurationInner.md) |  | [optional] 

## Methods

### NewNonPciIntegration

`func NewNonPciIntegration(accountId string, integrationId string, integrationType NonPciIntegrationEnum, ) *NonPciIntegration`

NewNonPciIntegration instantiates a new NonPciIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonPciIntegrationWithDefaults

`func NewNonPciIntegrationWithDefaults() *NonPciIntegration`

NewNonPciIntegrationWithDefaults instantiates a new NonPciIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NonPciIntegration) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NonPciIntegration) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NonPciIntegration) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetIntegrationId

`func (o *NonPciIntegration) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *NonPciIntegration) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *NonPciIntegration) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegrationType

`func (o *NonPciIntegration) GetIntegrationType() NonPciIntegrationEnum`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *NonPciIntegration) GetIntegrationTypeOk() (*NonPciIntegrationEnum, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *NonPciIntegration) SetIntegrationType(v NonPciIntegrationEnum)`

SetIntegrationType sets IntegrationType field to given value.


### GetConfiguration

`func (o *NonPciIntegration) GetConfiguration() []NonPciIntegrationConfigurationInner`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *NonPciIntegration) GetConfigurationOk() (*[]NonPciIntegrationConfigurationInner, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *NonPciIntegration) SetConfiguration(v []NonPciIntegrationConfigurationInner)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *NonPciIntegration) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


