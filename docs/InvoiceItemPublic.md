# InvoiceItemPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountAtom** | **int32** | Total amount of invoice_item in atomic units (in USD this is cents). | 
**AmountAtomConsideringDiscountApplied** | **int32** | Total amount of invoice_item in atomic units considering discounts | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) | Three-letter ISO currency code, in lowercase. | 
**DiscountAmountAtoms** | [**[]InvoiceItemDiscountAmountsPublic**](InvoiceItemDiscountAmountsPublic.md) | The amount_atom of all discount. | 
**Discounts** | **[]string** | The discounts only directly applied to the InvoiceItem. | 
**InvoiceItemDescription** | **NullableString** |  | 
**Name** | **string** | The name of the line item as it appears in the invoice. | 
**PeriodEnd** | **time.Time** | End of the usage period of the invoice_item. It is in &#39;ISO 8601&#39; format. | 
**PriceTiers** | Pointer to [**[]PriceTierPublic**](PriceTierPublic.md) | The price tiers of the product, if applicable. | [optional] 
**Quantity** | **int32** | Quantity of the line item. | 
**SubscriptionItemDescription** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceItemPublic

`func NewInvoiceItemPublic(amountAtom int32, amountAtomConsideringDiscountApplied int32, currency CurrencyEnum, discountAmountAtoms []InvoiceItemDiscountAmountsPublic, discounts []string, invoiceItemDescription NullableString, name string, periodEnd time.Time, quantity int32, ) *InvoiceItemPublic`

NewInvoiceItemPublic instantiates a new InvoiceItemPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceItemPublicWithDefaults

`func NewInvoiceItemPublicWithDefaults() *InvoiceItemPublic`

NewInvoiceItemPublicWithDefaults instantiates a new InvoiceItemPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountAtom

`func (o *InvoiceItemPublic) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *InvoiceItemPublic) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *InvoiceItemPublic) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetAmountAtomConsideringDiscountApplied

`func (o *InvoiceItemPublic) GetAmountAtomConsideringDiscountApplied() int32`

GetAmountAtomConsideringDiscountApplied returns the AmountAtomConsideringDiscountApplied field if non-nil, zero value otherwise.

### GetAmountAtomConsideringDiscountAppliedOk

`func (o *InvoiceItemPublic) GetAmountAtomConsideringDiscountAppliedOk() (*int32, bool)`

GetAmountAtomConsideringDiscountAppliedOk returns a tuple with the AmountAtomConsideringDiscountApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtomConsideringDiscountApplied

`func (o *InvoiceItemPublic) SetAmountAtomConsideringDiscountApplied(v int32)`

SetAmountAtomConsideringDiscountApplied sets AmountAtomConsideringDiscountApplied field to given value.


### GetCurrency

`func (o *InvoiceItemPublic) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceItemPublic) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceItemPublic) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetDiscountAmountAtoms

`func (o *InvoiceItemPublic) GetDiscountAmountAtoms() []InvoiceItemDiscountAmountsPublic`

GetDiscountAmountAtoms returns the DiscountAmountAtoms field if non-nil, zero value otherwise.

### GetDiscountAmountAtomsOk

`func (o *InvoiceItemPublic) GetDiscountAmountAtomsOk() (*[]InvoiceItemDiscountAmountsPublic, bool)`

GetDiscountAmountAtomsOk returns a tuple with the DiscountAmountAtoms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmountAtoms

`func (o *InvoiceItemPublic) SetDiscountAmountAtoms(v []InvoiceItemDiscountAmountsPublic)`

SetDiscountAmountAtoms sets DiscountAmountAtoms field to given value.


### GetDiscounts

`func (o *InvoiceItemPublic) GetDiscounts() []string`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *InvoiceItemPublic) GetDiscountsOk() (*[]string, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *InvoiceItemPublic) SetDiscounts(v []string)`

SetDiscounts sets Discounts field to given value.


### GetInvoiceItemDescription

`func (o *InvoiceItemPublic) GetInvoiceItemDescription() string`

GetInvoiceItemDescription returns the InvoiceItemDescription field if non-nil, zero value otherwise.

### GetInvoiceItemDescriptionOk

`func (o *InvoiceItemPublic) GetInvoiceItemDescriptionOk() (*string, bool)`

GetInvoiceItemDescriptionOk returns a tuple with the InvoiceItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemDescription

`func (o *InvoiceItemPublic) SetInvoiceItemDescription(v string)`

SetInvoiceItemDescription sets InvoiceItemDescription field to given value.


### SetInvoiceItemDescriptionNil

`func (o *InvoiceItemPublic) SetInvoiceItemDescriptionNil(b bool)`

 SetInvoiceItemDescriptionNil sets the value for InvoiceItemDescription to be an explicit nil

### UnsetInvoiceItemDescription
`func (o *InvoiceItemPublic) UnsetInvoiceItemDescription()`

UnsetInvoiceItemDescription ensures that no value is present for InvoiceItemDescription, not even an explicit nil
### GetName

`func (o *InvoiceItemPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceItemPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceItemPublic) SetName(v string)`

SetName sets Name field to given value.


### GetPeriodEnd

`func (o *InvoiceItemPublic) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *InvoiceItemPublic) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *InvoiceItemPublic) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetPriceTiers

`func (o *InvoiceItemPublic) GetPriceTiers() []PriceTierPublic`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *InvoiceItemPublic) GetPriceTiersOk() (*[]PriceTierPublic, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *InvoiceItemPublic) SetPriceTiers(v []PriceTierPublic)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *InvoiceItemPublic) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceItemPublic) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceItemPublic) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceItemPublic) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetSubscriptionItemDescription

`func (o *InvoiceItemPublic) GetSubscriptionItemDescription() string`

GetSubscriptionItemDescription returns the SubscriptionItemDescription field if non-nil, zero value otherwise.

### GetSubscriptionItemDescriptionOk

`func (o *InvoiceItemPublic) GetSubscriptionItemDescriptionOk() (*string, bool)`

GetSubscriptionItemDescriptionOk returns a tuple with the SubscriptionItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionItemDescription

`func (o *InvoiceItemPublic) SetSubscriptionItemDescription(v string)`

SetSubscriptionItemDescription sets SubscriptionItemDescription field to given value.

### HasSubscriptionItemDescription

`func (o *InvoiceItemPublic) HasSubscriptionItemDescription() bool`

HasSubscriptionItemDescription returns a boolean if a field has been set.

### SetSubscriptionItemDescriptionNil

`func (o *InvoiceItemPublic) SetSubscriptionItemDescriptionNil(b bool)`

 SetSubscriptionItemDescriptionNil sets the value for SubscriptionItemDescription to be an explicit nil

### UnsetSubscriptionItemDescription
`func (o *InvoiceItemPublic) UnsetSubscriptionItemDescription()`

UnsetSubscriptionItemDescription ensures that no value is present for SubscriptionItemDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


