# InvoiceItemDiscountAmountsPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountAtom** | **int32** | The amount_atom of the discount. | 
**CouponDescription** | Pointer to **NullableString** |  | [optional] 
**CouponName** | **string** | Name of the coupon that was applied to get this discount. | 

## Methods

### NewInvoiceItemDiscountAmountsPublic

`func NewInvoiceItemDiscountAmountsPublic(amountAtom int32, couponName string, ) *InvoiceItemDiscountAmountsPublic`

NewInvoiceItemDiscountAmountsPublic instantiates a new InvoiceItemDiscountAmountsPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceItemDiscountAmountsPublicWithDefaults

`func NewInvoiceItemDiscountAmountsPublicWithDefaults() *InvoiceItemDiscountAmountsPublic`

NewInvoiceItemDiscountAmountsPublicWithDefaults instantiates a new InvoiceItemDiscountAmountsPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountAtom

`func (o *InvoiceItemDiscountAmountsPublic) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *InvoiceItemDiscountAmountsPublic) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *InvoiceItemDiscountAmountsPublic) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetCouponDescription

`func (o *InvoiceItemDiscountAmountsPublic) GetCouponDescription() string`

GetCouponDescription returns the CouponDescription field if non-nil, zero value otherwise.

### GetCouponDescriptionOk

`func (o *InvoiceItemDiscountAmountsPublic) GetCouponDescriptionOk() (*string, bool)`

GetCouponDescriptionOk returns a tuple with the CouponDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponDescription

`func (o *InvoiceItemDiscountAmountsPublic) SetCouponDescription(v string)`

SetCouponDescription sets CouponDescription field to given value.

### HasCouponDescription

`func (o *InvoiceItemDiscountAmountsPublic) HasCouponDescription() bool`

HasCouponDescription returns a boolean if a field has been set.

### SetCouponDescriptionNil

`func (o *InvoiceItemDiscountAmountsPublic) SetCouponDescriptionNil(b bool)`

 SetCouponDescriptionNil sets the value for CouponDescription to be an explicit nil

### UnsetCouponDescription
`func (o *InvoiceItemDiscountAmountsPublic) UnsetCouponDescription()`

UnsetCouponDescription ensures that no value is present for CouponDescription, not even an explicit nil
### GetCouponName

`func (o *InvoiceItemDiscountAmountsPublic) GetCouponName() string`

GetCouponName returns the CouponName field if non-nil, zero value otherwise.

### GetCouponNameOk

`func (o *InvoiceItemDiscountAmountsPublic) GetCouponNameOk() (*string, bool)`

GetCouponNameOk returns a tuple with the CouponName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponName

`func (o *InvoiceItemDiscountAmountsPublic) SetCouponName(v string)`

SetCouponName sets CouponName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


