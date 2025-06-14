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

// checks if the RefundExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefundExternal{}

// RefundExternal struct for RefundExternal
type RefundExternal struct {
	// amount_atom that you want to refund.
	AmountAtom int32 `json:"amount_atom"`
	AttemptErrorMessage NullableString `json:"attempt_error_message,omitempty"`
	// Unique Identifier of the charge.
	ChargeId string `json:"charge_id"`
	// DateTime at which the object was created, in 'ISO 8601' format.
	CreatedAt time.Time `json:"created_at"`
	// Three-letter ISO currency code, in lowercase.
	Currency *CurrencyEnum `json:"currency,omitempty"`
	// Unique Identifier of the refund.
	Id string `json:"id"`
	// Invoice id to which the refund is attached.
	InvoiceId string `json:"invoice_id"`
	// If true, indicates that this object has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	Object *ObjectName `json:"object,omitempty"`
	// Unique Identifier of the payment_intent.
	PaymentIntentId string `json:"payment_intent_id"`
	// Reason of the refund.
	Reason RefundReasonEnum `json:"reason"`
	// Status of the refund.
	Status RefundStatusEnum `json:"status"`
	// DateTime at which the object was updated, in 'ISO 8601' format.
	UpdatedAt time.Time `json:"updated_at"`
}

type _RefundExternal RefundExternal

// NewRefundExternal instantiates a new RefundExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundExternal(amountAtom int32, chargeId string, createdAt time.Time, id string, invoiceId string, paymentIntentId string, reason RefundReasonEnum, status RefundStatusEnum, updatedAt time.Time) *RefundExternal {
	this := RefundExternal{}
	this.AmountAtom = amountAtom
	this.ChargeId = chargeId
	this.CreatedAt = createdAt
	this.Id = id
	this.InvoiceId = invoiceId
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	this.PaymentIntentId = paymentIntentId
	this.Reason = reason
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewRefundExternalWithDefaults instantiates a new RefundExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundExternalWithDefaults() *RefundExternal {
	this := RefundExternal{}
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// GetAmountAtom returns the AmountAtom field value
func (o *RefundExternal) GetAmountAtom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountAtom
}

// GetAmountAtomOk returns a tuple with the AmountAtom field value
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetAmountAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountAtom, true
}

// SetAmountAtom sets field value
func (o *RefundExternal) SetAmountAtom(v int32) {
	o.AmountAtom = v
}

// GetAttemptErrorMessage returns the AttemptErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefundExternal) GetAttemptErrorMessage() string {
	if o == nil || IsNil(o.AttemptErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.AttemptErrorMessage.Get()
}

// GetAttemptErrorMessageOk returns a tuple with the AttemptErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefundExternal) GetAttemptErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttemptErrorMessage.Get(), o.AttemptErrorMessage.IsSet()
}

// HasAttemptErrorMessage returns a boolean if a field has been set.
func (o *RefundExternal) HasAttemptErrorMessage() bool {
	if o != nil && o.AttemptErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetAttemptErrorMessage gets a reference to the given NullableString and assigns it to the AttemptErrorMessage field.
func (o *RefundExternal) SetAttemptErrorMessage(v string) {
	o.AttemptErrorMessage.Set(&v)
}
// SetAttemptErrorMessageNil sets the value for AttemptErrorMessage to be an explicit nil
func (o *RefundExternal) SetAttemptErrorMessageNil() {
	o.AttemptErrorMessage.Set(nil)
}

// UnsetAttemptErrorMessage ensures that no value is present for AttemptErrorMessage, not even an explicit nil
func (o *RefundExternal) UnsetAttemptErrorMessage() {
	o.AttemptErrorMessage.Unset()
}

// GetChargeId returns the ChargeId field value
func (o *RefundExternal) GetChargeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChargeId
}

// GetChargeIdOk returns a tuple with the ChargeId field value
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetChargeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChargeId, true
}

// SetChargeId sets field value
func (o *RefundExternal) SetChargeId(v string) {
	o.ChargeId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RefundExternal) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RefundExternal) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *RefundExternal) GetCurrency() CurrencyEnum {
	if o == nil || IsNil(o.Currency) {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *RefundExternal) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CurrencyEnum and assigns it to the Currency field.
func (o *RefundExternal) SetCurrency(v CurrencyEnum) {
	o.Currency = &v
}

// GetId returns the Id field value
func (o *RefundExternal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RefundExternal) SetId(v string) {
	o.Id = v
}

// GetInvoiceId returns the InvoiceId field value
func (o *RefundExternal) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *RefundExternal) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *RefundExternal) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *RefundExternal) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *RefundExternal) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *RefundExternal) GetObject() ObjectName {
	if o == nil || IsNil(o.Object) {
		var ret ObjectName
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetObjectOk() (*ObjectName, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *RefundExternal) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given ObjectName and assigns it to the Object field.
func (o *RefundExternal) SetObject(v ObjectName) {
	o.Object = &v
}

// GetPaymentIntentId returns the PaymentIntentId field value
func (o *RefundExternal) GetPaymentIntentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentIntentId
}

// GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field value
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetPaymentIntentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentIntentId, true
}

// SetPaymentIntentId sets field value
func (o *RefundExternal) SetPaymentIntentId(v string) {
	o.PaymentIntentId = v
}

// GetReason returns the Reason field value
func (o *RefundExternal) GetReason() RefundReasonEnum {
	if o == nil {
		var ret RefundReasonEnum
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetReasonOk() (*RefundReasonEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *RefundExternal) SetReason(v RefundReasonEnum) {
	o.Reason = v
}

// GetStatus returns the Status field value
func (o *RefundExternal) GetStatus() RefundStatusEnum {
	if o == nil {
		var ret RefundStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetStatusOk() (*RefundStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RefundExternal) SetStatus(v RefundStatusEnum) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RefundExternal) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RefundExternal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RefundExternal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o RefundExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundExternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount_atom"] = o.AmountAtom
	if o.AttemptErrorMessage.IsSet() {
		toSerialize["attempt_error_message"] = o.AttemptErrorMessage.Get()
	}
	toSerialize["charge_id"] = o.ChargeId
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["id"] = o.Id
	toSerialize["invoice_id"] = o.InvoiceId
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["payment_intent_id"] = o.PaymentIntentId
	toSerialize["reason"] = o.Reason
	toSerialize["status"] = o.Status
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *RefundExternal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount_atom",
		"charge_id",
		"created_at",
		"id",
		"invoice_id",
		"payment_intent_id",
		"reason",
		"status",
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

	varRefundExternal := _RefundExternal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRefundExternal)

	if err != nil {
		return err
	}

	*o = RefundExternal(varRefundExternal)

	return err
}

type NullableRefundExternal struct {
	value *RefundExternal
	isSet bool
}

func (v NullableRefundExternal) Get() *RefundExternal {
	return v.value
}

func (v *NullableRefundExternal) Set(val *RefundExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundExternal(val *RefundExternal) *NullableRefundExternal {
	return &NullableRefundExternal{value: val, isSet: true}
}

func (v NullableRefundExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


