# InvoiceItemExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique identifier of the account. | 
**AmountAtom** | **int32** | Total amount of invoice_item in atomic units (in USD this is cents). | 
**AmountAtomConsideringDiscountApplied** | **int32** | Total amount of invoice_item in atomic units considering discounts. Contains both invoice-level and invoice item-level discounts | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) | Three-letter ISO currency code, in lowercase. | 
**CustomerId** | **string** | Unique identifier of the customer. | 
**Description** | **NullableString** |  | 
**DiscountAmountAtoms** | [**[]InvoiceItemDiscountAmountsExternal**](InvoiceItemDiscountAmountsExternal.md) | The amount_atom of all discount. | 
**Discounts** | **[]string** | The discounts only directly applied to the InvoiceItem. | 
**Id** | **string** | Unique identifier of the invoice_item. | 
**InvoiceId** | **NullableString** |  | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Name** | **string** | Name of the invoice_item. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**PeriodEnd** | **time.Time** | End of the usage period of the invoice_item. It is in &#39;ISO 8601&#39; format. | 
**PeriodStart** | **time.Time** | Start of the usage period of the invoice_item. It is in &#39;ISO 8601&#39; format. | 
**PriceId** | **string** | Unique identifier of the price. | 
**PriceType** | [**PriceTypeEnum**](PriceTypeEnum.md) | Type of the price. | 
**ProductId** | **string** | Unique identifier of the product. | 
**Proration** | **bool** | Whether the invoice item was created automatically as a proration adjustment when the customer switched plans. | 
**ProrationDetailsDebitInvoiceItem** | **NullableString** |  | 
**Quantity** | **int32** | Quantity of the invoice_item. | 
**SubscriptionCancelledAt** | Pointer to **NullableTime** |  | [optional] 
**SubscriptionId** | **NullableString** |  | 
**SubscriptionItemDescription** | Pointer to **NullableString** |  | [optional] 
**SubscriptionItemId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewInvoiceItemExternal

`func NewInvoiceItemExternal(accountId string, amountAtom int32, amountAtomConsideringDiscountApplied int32, createdAt time.Time, currency CurrencyEnum, customerId string, description NullableString, discountAmountAtoms []InvoiceItemDiscountAmountsExternal, discounts []string, id string, invoiceId NullableString, name string, periodEnd time.Time, periodStart time.Time, priceId string, priceType PriceTypeEnum, productId string, proration bool, prorationDetailsDebitInvoiceItem NullableString, quantity int32, subscriptionId NullableString, updatedAt time.Time, ) *InvoiceItemExternal`

NewInvoiceItemExternal instantiates a new InvoiceItemExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceItemExternalWithDefaults

`func NewInvoiceItemExternalWithDefaults() *InvoiceItemExternal`

NewInvoiceItemExternalWithDefaults instantiates a new InvoiceItemExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *InvoiceItemExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InvoiceItemExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InvoiceItemExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmountAtom

`func (o *InvoiceItemExternal) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *InvoiceItemExternal) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *InvoiceItemExternal) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetAmountAtomConsideringDiscountApplied

`func (o *InvoiceItemExternal) GetAmountAtomConsideringDiscountApplied() int32`

GetAmountAtomConsideringDiscountApplied returns the AmountAtomConsideringDiscountApplied field if non-nil, zero value otherwise.

### GetAmountAtomConsideringDiscountAppliedOk

`func (o *InvoiceItemExternal) GetAmountAtomConsideringDiscountAppliedOk() (*int32, bool)`

GetAmountAtomConsideringDiscountAppliedOk returns a tuple with the AmountAtomConsideringDiscountApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtomConsideringDiscountApplied

`func (o *InvoiceItemExternal) SetAmountAtomConsideringDiscountApplied(v int32)`

SetAmountAtomConsideringDiscountApplied sets AmountAtomConsideringDiscountApplied field to given value.


### GetCreatedAt

`func (o *InvoiceItemExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceItemExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceItemExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *InvoiceItemExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceItemExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceItemExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *InvoiceItemExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InvoiceItemExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InvoiceItemExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDescription

`func (o *InvoiceItemExternal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceItemExternal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceItemExternal) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *InvoiceItemExternal) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceItemExternal) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDiscountAmountAtoms

`func (o *InvoiceItemExternal) GetDiscountAmountAtoms() []InvoiceItemDiscountAmountsExternal`

GetDiscountAmountAtoms returns the DiscountAmountAtoms field if non-nil, zero value otherwise.

### GetDiscountAmountAtomsOk

`func (o *InvoiceItemExternal) GetDiscountAmountAtomsOk() (*[]InvoiceItemDiscountAmountsExternal, bool)`

GetDiscountAmountAtomsOk returns a tuple with the DiscountAmountAtoms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmountAtoms

`func (o *InvoiceItemExternal) SetDiscountAmountAtoms(v []InvoiceItemDiscountAmountsExternal)`

SetDiscountAmountAtoms sets DiscountAmountAtoms field to given value.


### GetDiscounts

`func (o *InvoiceItemExternal) GetDiscounts() []string`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *InvoiceItemExternal) GetDiscountsOk() (*[]string, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *InvoiceItemExternal) SetDiscounts(v []string)`

SetDiscounts sets Discounts field to given value.


### GetId

`func (o *InvoiceItemExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceItemExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceItemExternal) SetId(v string)`

SetId sets Id field to given value.


### GetInvoiceId

`func (o *InvoiceItemExternal) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceItemExternal) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceItemExternal) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *InvoiceItemExternal) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *InvoiceItemExternal) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetIsDeleted

`func (o *InvoiceItemExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *InvoiceItemExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *InvoiceItemExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *InvoiceItemExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetName

`func (o *InvoiceItemExternal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceItemExternal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceItemExternal) SetName(v string)`

SetName sets Name field to given value.


### GetObject

`func (o *InvoiceItemExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *InvoiceItemExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *InvoiceItemExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *InvoiceItemExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *InvoiceItemExternal) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *InvoiceItemExternal) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *InvoiceItemExternal) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetPeriodStart

`func (o *InvoiceItemExternal) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *InvoiceItemExternal) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *InvoiceItemExternal) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.


### GetPriceId

`func (o *InvoiceItemExternal) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *InvoiceItemExternal) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *InvoiceItemExternal) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetPriceType

`func (o *InvoiceItemExternal) GetPriceType() PriceTypeEnum`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *InvoiceItemExternal) GetPriceTypeOk() (*PriceTypeEnum, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *InvoiceItemExternal) SetPriceType(v PriceTypeEnum)`

SetPriceType sets PriceType field to given value.


### GetProductId

`func (o *InvoiceItemExternal) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *InvoiceItemExternal) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *InvoiceItemExternal) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetProration

`func (o *InvoiceItemExternal) GetProration() bool`

GetProration returns the Proration field if non-nil, zero value otherwise.

### GetProrationOk

`func (o *InvoiceItemExternal) GetProrationOk() (*bool, bool)`

GetProrationOk returns a tuple with the Proration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProration

`func (o *InvoiceItemExternal) SetProration(v bool)`

SetProration sets Proration field to given value.


### GetProrationDetailsDebitInvoiceItem

`func (o *InvoiceItemExternal) GetProrationDetailsDebitInvoiceItem() string`

GetProrationDetailsDebitInvoiceItem returns the ProrationDetailsDebitInvoiceItem field if non-nil, zero value otherwise.

### GetProrationDetailsDebitInvoiceItemOk

`func (o *InvoiceItemExternal) GetProrationDetailsDebitInvoiceItemOk() (*string, bool)`

GetProrationDetailsDebitInvoiceItemOk returns a tuple with the ProrationDetailsDebitInvoiceItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDetailsDebitInvoiceItem

`func (o *InvoiceItemExternal) SetProrationDetailsDebitInvoiceItem(v string)`

SetProrationDetailsDebitInvoiceItem sets ProrationDetailsDebitInvoiceItem field to given value.


### SetProrationDetailsDebitInvoiceItemNil

`func (o *InvoiceItemExternal) SetProrationDetailsDebitInvoiceItemNil(b bool)`

 SetProrationDetailsDebitInvoiceItemNil sets the value for ProrationDetailsDebitInvoiceItem to be an explicit nil

### UnsetProrationDetailsDebitInvoiceItem
`func (o *InvoiceItemExternal) UnsetProrationDetailsDebitInvoiceItem()`

UnsetProrationDetailsDebitInvoiceItem ensures that no value is present for ProrationDetailsDebitInvoiceItem, not even an explicit nil
### GetQuantity

`func (o *InvoiceItemExternal) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceItemExternal) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceItemExternal) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetSubscriptionCancelledAt

`func (o *InvoiceItemExternal) GetSubscriptionCancelledAt() time.Time`

GetSubscriptionCancelledAt returns the SubscriptionCancelledAt field if non-nil, zero value otherwise.

### GetSubscriptionCancelledAtOk

`func (o *InvoiceItemExternal) GetSubscriptionCancelledAtOk() (*time.Time, bool)`

GetSubscriptionCancelledAtOk returns a tuple with the SubscriptionCancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCancelledAt

`func (o *InvoiceItemExternal) SetSubscriptionCancelledAt(v time.Time)`

SetSubscriptionCancelledAt sets SubscriptionCancelledAt field to given value.

### HasSubscriptionCancelledAt

`func (o *InvoiceItemExternal) HasSubscriptionCancelledAt() bool`

HasSubscriptionCancelledAt returns a boolean if a field has been set.

### SetSubscriptionCancelledAtNil

`func (o *InvoiceItemExternal) SetSubscriptionCancelledAtNil(b bool)`

 SetSubscriptionCancelledAtNil sets the value for SubscriptionCancelledAt to be an explicit nil

### UnsetSubscriptionCancelledAt
`func (o *InvoiceItemExternal) UnsetSubscriptionCancelledAt()`

UnsetSubscriptionCancelledAt ensures that no value is present for SubscriptionCancelledAt, not even an explicit nil
### GetSubscriptionId

`func (o *InvoiceItemExternal) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *InvoiceItemExternal) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *InvoiceItemExternal) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### SetSubscriptionIdNil

`func (o *InvoiceItemExternal) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *InvoiceItemExternal) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetSubscriptionItemDescription

`func (o *InvoiceItemExternal) GetSubscriptionItemDescription() string`

GetSubscriptionItemDescription returns the SubscriptionItemDescription field if non-nil, zero value otherwise.

### GetSubscriptionItemDescriptionOk

`func (o *InvoiceItemExternal) GetSubscriptionItemDescriptionOk() (*string, bool)`

GetSubscriptionItemDescriptionOk returns a tuple with the SubscriptionItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionItemDescription

`func (o *InvoiceItemExternal) SetSubscriptionItemDescription(v string)`

SetSubscriptionItemDescription sets SubscriptionItemDescription field to given value.

### HasSubscriptionItemDescription

`func (o *InvoiceItemExternal) HasSubscriptionItemDescription() bool`

HasSubscriptionItemDescription returns a boolean if a field has been set.

### SetSubscriptionItemDescriptionNil

`func (o *InvoiceItemExternal) SetSubscriptionItemDescriptionNil(b bool)`

 SetSubscriptionItemDescriptionNil sets the value for SubscriptionItemDescription to be an explicit nil

### UnsetSubscriptionItemDescription
`func (o *InvoiceItemExternal) UnsetSubscriptionItemDescription()`

UnsetSubscriptionItemDescription ensures that no value is present for SubscriptionItemDescription, not even an explicit nil
### GetSubscriptionItemId

`func (o *InvoiceItemExternal) GetSubscriptionItemId() string`

GetSubscriptionItemId returns the SubscriptionItemId field if non-nil, zero value otherwise.

### GetSubscriptionItemIdOk

`func (o *InvoiceItemExternal) GetSubscriptionItemIdOk() (*string, bool)`

GetSubscriptionItemIdOk returns a tuple with the SubscriptionItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionItemId

`func (o *InvoiceItemExternal) SetSubscriptionItemId(v string)`

SetSubscriptionItemId sets SubscriptionItemId field to given value.

### HasSubscriptionItemId

`func (o *InvoiceItemExternal) HasSubscriptionItemId() bool`

HasSubscriptionItemId returns a boolean if a field has been set.

### SetSubscriptionItemIdNil

`func (o *InvoiceItemExternal) SetSubscriptionItemIdNil(b bool)`

 SetSubscriptionItemIdNil sets the value for SubscriptionItemId to be an explicit nil

### UnsetSubscriptionItemId
`func (o *InvoiceItemExternal) UnsetSubscriptionItemId()`

UnsetSubscriptionItemId ensures that no value is present for SubscriptionItemId, not even an explicit nil
### GetUpdatedAt

`func (o *InvoiceItemExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceItemExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceItemExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


