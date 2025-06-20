/*
OpenPay API

super charge your subscription management.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the CheckoutSessionLineItemExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutSessionLineItemExternal{}

// CheckoutSessionLineItemExternal struct for CheckoutSessionLineItemExternal
type CheckoutSessionLineItemExternal struct {
	AmountSubtotalAtom int32 `json:"amount_subtotal_atom"`
	AmountTotalAtom int32 `json:"amount_total_atom"`
	BillingInterval NullableCalendarIntervalEnum `json:"billing_interval"`
	BillingIntervalCount NullableInt32 `json:"billing_interval_count"`
	CheckoutSessionId NullableString `json:"checkout_session_id"`
	// DateTime at which the object was created, in 'ISO 8601' format.
	CreatedAt time.Time `json:"created_at"`
	Currency CurrencyEnum `json:"currency"`
	CustomFields map[string]interface{} `json:"custom_fields"`
	Description NullableString `json:"description"`
	DescriptionDetailed NullableString `json:"description_detailed"`
	Id string `json:"id"`
	// If true, indicates that this object has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	Object *ObjectName `json:"object,omitempty"`
	PriceId string `json:"price_id"`
	ProductId string `json:"product_id"`
	Quantity int32 `json:"quantity"`
	// DateTime at which the object was updated, in 'ISO 8601' format.
	UpdatedAt time.Time `json:"updated_at"`
}

type _CheckoutSessionLineItemExternal CheckoutSessionLineItemExternal

// NewCheckoutSessionLineItemExternal instantiates a new CheckoutSessionLineItemExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSessionLineItemExternal(amountSubtotalAtom int32, amountTotalAtom int32, billingInterval NullableCalendarIntervalEnum, billingIntervalCount NullableInt32, checkoutSessionId NullableString, createdAt time.Time, currency CurrencyEnum, customFields map[string]interface{}, description NullableString, descriptionDetailed NullableString, id string, priceId string, productId string, quantity int32, updatedAt time.Time) *CheckoutSessionLineItemExternal {
	this := CheckoutSessionLineItemExternal{}
	this.AmountSubtotalAtom = amountSubtotalAtom
	this.AmountTotalAtom = amountTotalAtom
	this.BillingInterval = billingInterval
	this.BillingIntervalCount = billingIntervalCount
	this.CheckoutSessionId = checkoutSessionId
	this.CreatedAt = createdAt
	this.Currency = currency
	this.CustomFields = customFields
	this.Description = description
	this.DescriptionDetailed = descriptionDetailed
	this.Id = id
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	this.PriceId = priceId
	this.ProductId = productId
	this.Quantity = quantity
	this.UpdatedAt = updatedAt
	return &this
}

// NewCheckoutSessionLineItemExternalWithDefaults instantiates a new CheckoutSessionLineItemExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionLineItemExternalWithDefaults() *CheckoutSessionLineItemExternal {
	this := CheckoutSessionLineItemExternal{}
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// GetAmountSubtotalAtom returns the AmountSubtotalAtom field value
func (o *CheckoutSessionLineItemExternal) GetAmountSubtotalAtom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountSubtotalAtom
}

// GetAmountSubtotalAtomOk returns a tuple with the AmountSubtotalAtom field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionLineItemExternal) GetAmountSubtotalAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountSubtotalAtom, true
}

// SetAmountSubtotalAtom sets field value
func (o *CheckoutSessionLineItemExternal) SetAmountSubtotalAtom(v int32) {
	o.AmountSubtotalAtom = v
}

// GetAmountTotalAtom returns the AmountTotalAtom field value
func (o *CheckoutSessionLineItemExternal) GetAmountTotalAtom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountTotalAtom
}

// GetAmountTotalAtomOk returns a tuple with the AmountTotalAtom field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionLineItemExternal) GetAmountTotalAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountTotalAtom, true
}

// SetAmountTotalAtom sets field value
func (o *CheckoutSessionLineItemExternal) SetAmountTotalAtom(v int32) {
	o.AmountTotalAtom = v
}

// GetBillingInterval returns the BillingInterval field value
// If the value is explicit nil, the zero value for CalendarIntervalEnum will be returned
func (o *CheckoutSessionLineItemExternal) GetBillingInterval() CalendarIntervalEnum {
	if o == nil || o.BillingInterval.Get() == nil {
		var ret CalendarIntervalEnum
		return ret
	}

	return *o.BillingInterval.Get()
}

// GetBillingIntervalOk returns a tuple with the BillingInterval field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionLineItemExternal) GetBillingIntervalOk() (*CalendarIntervalEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingInterval.Get(), o.BillingInterval.IsSet()
}

// SetBillingInterval sets field value
func (o *CheckoutSessionLineItemExternal) SetBillingInterval(v CalendarIntervalEnum) {
	o.BillingInterval.Set(&v)
}

// GetBillingIntervalCount returns the BillingIntervalCount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *CheckoutSessionLineItemExternal) GetBillingIntervalCount() int32 {
	if o == nil || o.BillingIntervalCount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.BillingIntervalCount.Get()
}

// GetBillingIntervalCountOk returns a tuple with the BillingIntervalCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionLineItemExternal) GetBillingIntervalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingIntervalCount.Get(), o.BillingIntervalCount.IsSet()
}

// SetBillingIntervalCount sets field value
func (o *CheckoutSessionLineItemExternal) SetBillingIntervalCount(v int32) {
	o.BillingIntervalCount.Set(&v)
}

// GetCheckoutSessionId returns the CheckoutSessionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSessionLineItemExternal) GetCheckoutSessionId() string {
	if o == nil || o.CheckoutSessionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CheckoutSessionId.Get()
}

// GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionLineItemExternal) GetCheckoutSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckoutSessionId.Get(), o.CheckoutSessionId.IsSet()
}

// SetCheckoutSessionId sets field value
func (o *CheckoutSessionLineItemExternal) SetCheckoutSessionId(v string) {
	o.CheckoutSessionId.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *CheckoutSessionLineItemExternal) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionLineItemExternal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CheckoutSessionLineItemExternal) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCurrency returns the Currency field value
func (o *CheckoutSessionLineItemExternal) GetCurrency() CurrencyEnum {
	if o == nil {
		var ret CurrencyEnum
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionLineItemExternal) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CheckoutSessionLineItemExternal) SetCurrency(v CurrencyEnum) {
	o.Currency = v
}

// GetCustomFields returns the CustomFields field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *CheckoutSessionLineItemExternal) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionLineItemExternal) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// SetCustomFields sets field value
func (o *CheckoutSessionLineItemExternal) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSessionLineItemExternal) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionLineItemExternal) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *CheckoutSessionLineItemExternal) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetDescriptionDetailed returns the DescriptionDetailed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSessionLineItemExternal) GetDescriptionDetailed() string {
	if o == nil || o.DescriptionDetailed.Get() == nil {
		var ret string
		return ret
	}

	return *o.DescriptionDetailed.Get()
}

// GetDescriptionDetailedOk returns a tuple with the DescriptionDetailed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionLineItemExternal) GetDescriptionDetailedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DescriptionDetailed.Get(), o.DescriptionDetailed.IsSet()
}

// SetDescriptionDetailed sets field value
func (o *CheckoutSessionLineItemExternal) SetDescriptionDetailed(v string) {
	o.DescriptionDetailed.Set(&v)
}

// GetId returns the Id field value
func (o *CheckoutSessionLineItemExternal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionLineItemExternal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CheckoutSessionLineItemExternal) SetId(v string) {
	o.Id = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *CheckoutSessionLineItemExternal) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionLineItemExternal) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *CheckoutSessionLineItemExternal) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *CheckoutSessionLineItemExternal) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CheckoutSessionLineItemExternal) GetObject() ObjectName {
	if o == nil || IsNil(o.Object) {
		var ret ObjectName
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionLineItemExternal) GetObjectOk() (*ObjectName, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CheckoutSessionLineItemExternal) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given ObjectName and assigns it to the Object field.
func (o *CheckoutSessionLineItemExternal) SetObject(v ObjectName) {
	o.Object = &v
}

// GetPriceId returns the PriceId field value
func (o *CheckoutSessionLineItemExternal) GetPriceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PriceId
}

// GetPriceIdOk returns a tuple with the PriceId field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionLineItemExternal) GetPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceId, true
}

// SetPriceId sets field value
func (o *CheckoutSessionLineItemExternal) SetPriceId(v string) {
	o.PriceId = v
}

// GetProductId returns the ProductId field value
func (o *CheckoutSessionLineItemExternal) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionLineItemExternal) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *CheckoutSessionLineItemExternal) SetProductId(v string) {
	o.ProductId = v
}

// GetQuantity returns the Quantity field value
func (o *CheckoutSessionLineItemExternal) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionLineItemExternal) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CheckoutSessionLineItemExternal) SetQuantity(v int32) {
	o.Quantity = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CheckoutSessionLineItemExternal) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionLineItemExternal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CheckoutSessionLineItemExternal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CheckoutSessionLineItemExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutSessionLineItemExternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount_subtotal_atom"] = o.AmountSubtotalAtom
	toSerialize["amount_total_atom"] = o.AmountTotalAtom
	toSerialize["billing_interval"] = o.BillingInterval.Get()
	toSerialize["billing_interval_count"] = o.BillingIntervalCount.Get()
	toSerialize["checkout_session_id"] = o.CheckoutSessionId.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["currency"] = o.Currency
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["description"] = o.Description.Get()
	toSerialize["description_detailed"] = o.DescriptionDetailed.Get()
	toSerialize["id"] = o.Id
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["price_id"] = o.PriceId
	toSerialize["product_id"] = o.ProductId
	toSerialize["quantity"] = o.Quantity
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *CheckoutSessionLineItemExternal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount_subtotal_atom",
		"amount_total_atom",
		"billing_interval",
		"billing_interval_count",
		"checkout_session_id",
		"created_at",
		"currency",
		"custom_fields",
		"description",
		"description_detailed",
		"id",
		"price_id",
		"product_id",
		"quantity",
		"updated_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCheckoutSessionLineItemExternal := _CheckoutSessionLineItemExternal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCheckoutSessionLineItemExternal)

	if err != nil {
		return err
	}

	*o = CheckoutSessionLineItemExternal(varCheckoutSessionLineItemExternal)

	return err
}

type NullableCheckoutSessionLineItemExternal struct {
	value *CheckoutSessionLineItemExternal
	isSet bool
}

func (v NullableCheckoutSessionLineItemExternal) Get() *CheckoutSessionLineItemExternal {
	return v.value
}

func (v *NullableCheckoutSessionLineItemExternal) Set(val *CheckoutSessionLineItemExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionLineItemExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionLineItemExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionLineItemExternal(val *CheckoutSessionLineItemExternal) *NullableCheckoutSessionLineItemExternal {
	return &NullableCheckoutSessionLineItemExternal{value: val, isSet: true}
}

func (v NullableCheckoutSessionLineItemExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionLineItemExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


