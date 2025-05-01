# CustomerInvoiceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerTaxIds** | Pointer to [**[]TaxIdSetting**](TaxIdSetting.md) | The tax IDs of the customer. | [optional] 
**DefaultNetD** | Pointer to **NullableInt32** |  | [optional] 
**EmailReceiptOnPaid** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCustomerInvoiceSettings

`func NewCustomerInvoiceSettings() *CustomerInvoiceSettings`

NewCustomerInvoiceSettings instantiates a new CustomerInvoiceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerInvoiceSettingsWithDefaults

`func NewCustomerInvoiceSettingsWithDefaults() *CustomerInvoiceSettings`

NewCustomerInvoiceSettingsWithDefaults instantiates a new CustomerInvoiceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerTaxIds

`func (o *CustomerInvoiceSettings) GetCustomerTaxIds() []TaxIdSetting`

GetCustomerTaxIds returns the CustomerTaxIds field if non-nil, zero value otherwise.

### GetCustomerTaxIdsOk

`func (o *CustomerInvoiceSettings) GetCustomerTaxIdsOk() (*[]TaxIdSetting, bool)`

GetCustomerTaxIdsOk returns a tuple with the CustomerTaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTaxIds

`func (o *CustomerInvoiceSettings) SetCustomerTaxIds(v []TaxIdSetting)`

SetCustomerTaxIds sets CustomerTaxIds field to given value.

### HasCustomerTaxIds

`func (o *CustomerInvoiceSettings) HasCustomerTaxIds() bool`

HasCustomerTaxIds returns a boolean if a field has been set.

### GetDefaultNetD

`func (o *CustomerInvoiceSettings) GetDefaultNetD() int32`

GetDefaultNetD returns the DefaultNetD field if non-nil, zero value otherwise.

### GetDefaultNetDOk

`func (o *CustomerInvoiceSettings) GetDefaultNetDOk() (*int32, bool)`

GetDefaultNetDOk returns a tuple with the DefaultNetD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetD

`func (o *CustomerInvoiceSettings) SetDefaultNetD(v int32)`

SetDefaultNetD sets DefaultNetD field to given value.

### HasDefaultNetD

`func (o *CustomerInvoiceSettings) HasDefaultNetD() bool`

HasDefaultNetD returns a boolean if a field has been set.

### SetDefaultNetDNil

`func (o *CustomerInvoiceSettings) SetDefaultNetDNil(b bool)`

 SetDefaultNetDNil sets the value for DefaultNetD to be an explicit nil

### UnsetDefaultNetD
`func (o *CustomerInvoiceSettings) UnsetDefaultNetD()`

UnsetDefaultNetD ensures that no value is present for DefaultNetD, not even an explicit nil
### GetEmailReceiptOnPaid

`func (o *CustomerInvoiceSettings) GetEmailReceiptOnPaid() bool`

GetEmailReceiptOnPaid returns the EmailReceiptOnPaid field if non-nil, zero value otherwise.

### GetEmailReceiptOnPaidOk

`func (o *CustomerInvoiceSettings) GetEmailReceiptOnPaidOk() (*bool, bool)`

GetEmailReceiptOnPaidOk returns a tuple with the EmailReceiptOnPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailReceiptOnPaid

`func (o *CustomerInvoiceSettings) SetEmailReceiptOnPaid(v bool)`

SetEmailReceiptOnPaid sets EmailReceiptOnPaid field to given value.

### HasEmailReceiptOnPaid

`func (o *CustomerInvoiceSettings) HasEmailReceiptOnPaid() bool`

HasEmailReceiptOnPaid returns a boolean if a field has been set.

### SetEmailReceiptOnPaidNil

`func (o *CustomerInvoiceSettings) SetEmailReceiptOnPaidNil(b bool)`

 SetEmailReceiptOnPaidNil sets the value for EmailReceiptOnPaid to be an explicit nil

### UnsetEmailReceiptOnPaid
`func (o *CustomerInvoiceSettings) UnsetEmailReceiptOnPaid()`

UnsetEmailReceiptOnPaid ensures that no value is present for EmailReceiptOnPaid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


