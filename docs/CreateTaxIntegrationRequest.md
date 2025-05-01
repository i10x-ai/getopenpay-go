# CreateTaxIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountingTz** | **string** | The timezone name for the accounting timezone (see pytz.all_timezones for a full list) | 
**ApiKeys** | **map[string]string** | These keys will allow to authenticate API requests to the tax processor. | 
**ApiName** | [**TaxIntegrationApiName**](TaxIntegrationApiName.md) | The name of the API used for the tax integration. | 

## Methods

### NewCreateTaxIntegrationRequest

`func NewCreateTaxIntegrationRequest(accountingTz string, apiKeys map[string]string, apiName TaxIntegrationApiName, ) *CreateTaxIntegrationRequest`

NewCreateTaxIntegrationRequest instantiates a new CreateTaxIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaxIntegrationRequestWithDefaults

`func NewCreateTaxIntegrationRequestWithDefaults() *CreateTaxIntegrationRequest`

NewCreateTaxIntegrationRequestWithDefaults instantiates a new CreateTaxIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountingTz

`func (o *CreateTaxIntegrationRequest) GetAccountingTz() string`

GetAccountingTz returns the AccountingTz field if non-nil, zero value otherwise.

### GetAccountingTzOk

`func (o *CreateTaxIntegrationRequest) GetAccountingTzOk() (*string, bool)`

GetAccountingTzOk returns a tuple with the AccountingTz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingTz

`func (o *CreateTaxIntegrationRequest) SetAccountingTz(v string)`

SetAccountingTz sets AccountingTz field to given value.


### GetApiKeys

`func (o *CreateTaxIntegrationRequest) GetApiKeys() map[string]string`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *CreateTaxIntegrationRequest) GetApiKeysOk() (*map[string]string, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *CreateTaxIntegrationRequest) SetApiKeys(v map[string]string)`

SetApiKeys sets ApiKeys field to given value.


### GetApiName

`func (o *CreateTaxIntegrationRequest) GetApiName() TaxIntegrationApiName`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *CreateTaxIntegrationRequest) GetApiNameOk() (*TaxIntegrationApiName, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *CreateTaxIntegrationRequest) SetApiName(v TaxIntegrationApiName)`

SetApiName sets ApiName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


