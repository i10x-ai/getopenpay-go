# MerchantTaxIdSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | **bool** | Whether this tax ID is the default for the country | 
**TaxId** | **string** | The tax ID value | 
**TaxIdType** | **string** | The type code for this tax ID | 

## Methods

### NewMerchantTaxIdSetting

`func NewMerchantTaxIdSetting(isDefault bool, taxId string, taxIdType string, ) *MerchantTaxIdSetting`

NewMerchantTaxIdSetting instantiates a new MerchantTaxIdSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantTaxIdSettingWithDefaults

`func NewMerchantTaxIdSettingWithDefaults() *MerchantTaxIdSetting`

NewMerchantTaxIdSettingWithDefaults instantiates a new MerchantTaxIdSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefault

`func (o *MerchantTaxIdSetting) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MerchantTaxIdSetting) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MerchantTaxIdSetting) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetTaxId

`func (o *MerchantTaxIdSetting) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *MerchantTaxIdSetting) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *MerchantTaxIdSetting) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.


### GetTaxIdType

`func (o *MerchantTaxIdSetting) GetTaxIdType() string`

GetTaxIdType returns the TaxIdType field if non-nil, zero value otherwise.

### GetTaxIdTypeOk

`func (o *MerchantTaxIdSetting) GetTaxIdTypeOk() (*string, bool)`

GetTaxIdTypeOk returns a tuple with the TaxIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdType

`func (o *MerchantTaxIdSetting) SetTaxIdType(v string)`

SetTaxIdType sets TaxIdType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


