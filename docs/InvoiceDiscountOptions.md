# InvoiceDiscountOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coupons** | Pointer to **[]string** |  | [optional] 
**Discount** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceDiscountOptions

`func NewInvoiceDiscountOptions() *InvoiceDiscountOptions`

NewInvoiceDiscountOptions instantiates a new InvoiceDiscountOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDiscountOptionsWithDefaults

`func NewInvoiceDiscountOptionsWithDefaults() *InvoiceDiscountOptions`

NewInvoiceDiscountOptionsWithDefaults instantiates a new InvoiceDiscountOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoupons

`func (o *InvoiceDiscountOptions) GetCoupons() []string`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *InvoiceDiscountOptions) GetCouponsOk() (*[]string, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *InvoiceDiscountOptions) SetCoupons(v []string)`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *InvoiceDiscountOptions) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### GetDiscount

`func (o *InvoiceDiscountOptions) GetDiscount() string`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *InvoiceDiscountOptions) GetDiscountOk() (*string, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *InvoiceDiscountOptions) SetDiscount(v string)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *InvoiceDiscountOptions) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *InvoiceDiscountOptions) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *InvoiceDiscountOptions) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


