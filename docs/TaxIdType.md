# TaxIdType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Human readable description of this tax ID type | 
**Example** | **string** | Example of a valid tax ID of this type | 
**TaxIdType** | [**TaxIdTypeEnum**](TaxIdTypeEnum.md) | The type code for this tax ID | 

## Methods

### NewTaxIdType

`func NewTaxIdType(description string, example string, taxIdType TaxIdTypeEnum, ) *TaxIdType`

NewTaxIdType instantiates a new TaxIdType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxIdTypeWithDefaults

`func NewTaxIdTypeWithDefaults() *TaxIdType`

NewTaxIdTypeWithDefaults instantiates a new TaxIdType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TaxIdType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxIdType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxIdType) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExample

`func (o *TaxIdType) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *TaxIdType) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *TaxIdType) SetExample(v string)`

SetExample sets Example field to given value.


### GetTaxIdType

`func (o *TaxIdType) GetTaxIdType() TaxIdTypeEnum`

GetTaxIdType returns the TaxIdType field if non-nil, zero value otherwise.

### GetTaxIdTypeOk

`func (o *TaxIdType) GetTaxIdTypeOk() (*TaxIdTypeEnum, bool)`

GetTaxIdTypeOk returns a tuple with the TaxIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdType

`func (o *TaxIdType) SetTaxIdType(v TaxIdTypeEnum)`

SetTaxIdType sets TaxIdType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


