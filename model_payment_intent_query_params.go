/*
OpenPay API

super charge your subscription management.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
)

// checks if the PaymentIntentQueryParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentIntentQueryParams{}

// PaymentIntentQueryParams struct for PaymentIntentQueryParams
type PaymentIntentQueryParams struct {
	AmountAtom NullableIntRangeFilter `json:"amount_atom,omitempty"`
	CreatedAt NullableDateTimeFilter `json:"created_at,omitempty"`
	CustomerId NullableString `json:"customer_id,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []string `json:"expand,omitempty"`
	InvoiceId NullableString `json:"invoice_id,omitempty"`
	// Page number
	PageNumber *int32 `json:"page_number,omitempty"`
	// Page size
	PageSize *int32 `json:"page_size,omitempty"`
	PaymentMethodId NullableString `json:"payment_method_id,omitempty"`
	Search NullableSearchFilter `json:"search,omitempty"`
	// Sort direction.
	SortDescending *bool `json:"sort_descending,omitempty"`
	// Key name based on which data is sorted.
	SortKey *string `json:"sort_key,omitempty"`
	Status NullablePaymentIntentStatus `json:"status,omitempty"`
	UpdatedAt NullableDateTimeFilter `json:"updated_at,omitempty"`
}

// NewPaymentIntentQueryParams instantiates a new PaymentIntentQueryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentIntentQueryParams() *PaymentIntentQueryParams {
	this := PaymentIntentQueryParams{}
	var pageNumber int32 = 1
	this.PageNumber = &pageNumber
	var pageSize int32 = 100
	this.PageSize = &pageSize
	var sortDescending bool = false
	this.SortDescending = &sortDescending
	var sortKey string = "created_at"
	this.SortKey = &sortKey
	return &this
}

// NewPaymentIntentQueryParamsWithDefaults instantiates a new PaymentIntentQueryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentIntentQueryParamsWithDefaults() *PaymentIntentQueryParams {
	this := PaymentIntentQueryParams{}
	var pageNumber int32 = 1
	this.PageNumber = &pageNumber
	var pageSize int32 = 100
	this.PageSize = &pageSize
	var sortDescending bool = false
	this.SortDescending = &sortDescending
	var sortKey string = "created_at"
	this.SortKey = &sortKey
	return &this
}

// GetAmountAtom returns the AmountAtom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntentQueryParams) GetAmountAtom() IntRangeFilter {
	if o == nil || IsNil(o.AmountAtom.Get()) {
		var ret IntRangeFilter
		return ret
	}
	return *o.AmountAtom.Get()
}

// GetAmountAtomOk returns a tuple with the AmountAtom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntentQueryParams) GetAmountAtomOk() (*IntRangeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountAtom.Get(), o.AmountAtom.IsSet()
}

// HasAmountAtom returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasAmountAtom() bool {
	if o != nil && o.AmountAtom.IsSet() {
		return true
	}

	return false
}

// SetAmountAtom gets a reference to the given NullableIntRangeFilter and assigns it to the AmountAtom field.
func (o *PaymentIntentQueryParams) SetAmountAtom(v IntRangeFilter) {
	o.AmountAtom.Set(&v)
}
// SetAmountAtomNil sets the value for AmountAtom to be an explicit nil
func (o *PaymentIntentQueryParams) SetAmountAtomNil() {
	o.AmountAtom.Set(nil)
}

// UnsetAmountAtom ensures that no value is present for AmountAtom, not even an explicit nil
func (o *PaymentIntentQueryParams) UnsetAmountAtom() {
	o.AmountAtom.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntentQueryParams) GetCreatedAt() DateTimeFilter {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntentQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableDateTimeFilter and assigns it to the CreatedAt field.
func (o *PaymentIntentQueryParams) SetCreatedAt(v DateTimeFilter) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *PaymentIntentQueryParams) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *PaymentIntentQueryParams) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntentQueryParams) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntentQueryParams) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *PaymentIntentQueryParams) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *PaymentIntentQueryParams) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *PaymentIntentQueryParams) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetExpand returns the Expand field value if set, zero value otherwise.
func (o *PaymentIntentQueryParams) GetExpand() []string {
	if o == nil || IsNil(o.Expand) {
		var ret []string
		return ret
	}
	return o.Expand
}

// GetExpandOk returns a tuple with the Expand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntentQueryParams) GetExpandOk() ([]string, bool) {
	if o == nil || IsNil(o.Expand) {
		return nil, false
	}
	return o.Expand, true
}

// HasExpand returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasExpand() bool {
	if o != nil && !IsNil(o.Expand) {
		return true
	}

	return false
}

// SetExpand gets a reference to the given []string and assigns it to the Expand field.
func (o *PaymentIntentQueryParams) SetExpand(v []string) {
	o.Expand = v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntentQueryParams) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntentQueryParams) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasInvoiceId() bool {
	if o != nil && o.InvoiceId.IsSet() {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given NullableString and assigns it to the InvoiceId field.
func (o *PaymentIntentQueryParams) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}
// SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil
func (o *PaymentIntentQueryParams) SetInvoiceIdNil() {
	o.InvoiceId.Set(nil)
}

// UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
func (o *PaymentIntentQueryParams) UnsetInvoiceId() {
	o.InvoiceId.Unset()
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *PaymentIntentQueryParams) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntentQueryParams) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *PaymentIntentQueryParams) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *PaymentIntentQueryParams) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntentQueryParams) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *PaymentIntentQueryParams) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntentQueryParams) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId.Get()
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntentQueryParams) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodId.Get(), o.PaymentMethodId.IsSet()
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given NullableString and assigns it to the PaymentMethodId field.
func (o *PaymentIntentQueryParams) SetPaymentMethodId(v string) {
	o.PaymentMethodId.Set(&v)
}
// SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil
func (o *PaymentIntentQueryParams) SetPaymentMethodIdNil() {
	o.PaymentMethodId.Set(nil)
}

// UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
func (o *PaymentIntentQueryParams) UnsetPaymentMethodId() {
	o.PaymentMethodId.Unset()
}

// GetSearch returns the Search field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntentQueryParams) GetSearch() SearchFilter {
	if o == nil || IsNil(o.Search.Get()) {
		var ret SearchFilter
		return ret
	}
	return *o.Search.Get()
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntentQueryParams) GetSearchOk() (*SearchFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Search.Get(), o.Search.IsSet()
}

// HasSearch returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasSearch() bool {
	if o != nil && o.Search.IsSet() {
		return true
	}

	return false
}

// SetSearch gets a reference to the given NullableSearchFilter and assigns it to the Search field.
func (o *PaymentIntentQueryParams) SetSearch(v SearchFilter) {
	o.Search.Set(&v)
}
// SetSearchNil sets the value for Search to be an explicit nil
func (o *PaymentIntentQueryParams) SetSearchNil() {
	o.Search.Set(nil)
}

// UnsetSearch ensures that no value is present for Search, not even an explicit nil
func (o *PaymentIntentQueryParams) UnsetSearch() {
	o.Search.Unset()
}

// GetSortDescending returns the SortDescending field value if set, zero value otherwise.
func (o *PaymentIntentQueryParams) GetSortDescending() bool {
	if o == nil || IsNil(o.SortDescending) {
		var ret bool
		return ret
	}
	return *o.SortDescending
}

// GetSortDescendingOk returns a tuple with the SortDescending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntentQueryParams) GetSortDescendingOk() (*bool, bool) {
	if o == nil || IsNil(o.SortDescending) {
		return nil, false
	}
	return o.SortDescending, true
}

// HasSortDescending returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasSortDescending() bool {
	if o != nil && !IsNil(o.SortDescending) {
		return true
	}

	return false
}

// SetSortDescending gets a reference to the given bool and assigns it to the SortDescending field.
func (o *PaymentIntentQueryParams) SetSortDescending(v bool) {
	o.SortDescending = &v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise.
func (o *PaymentIntentQueryParams) GetSortKey() string {
	if o == nil || IsNil(o.SortKey) {
		var ret string
		return ret
	}
	return *o.SortKey
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntentQueryParams) GetSortKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SortKey) {
		return nil, false
	}
	return o.SortKey, true
}

// HasSortKey returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasSortKey() bool {
	if o != nil && !IsNil(o.SortKey) {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given string and assigns it to the SortKey field.
func (o *PaymentIntentQueryParams) SetSortKey(v string) {
	o.SortKey = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntentQueryParams) GetStatus() PaymentIntentStatus {
	if o == nil || IsNil(o.Status.Get()) {
		var ret PaymentIntentStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntentQueryParams) GetStatusOk() (*PaymentIntentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePaymentIntentStatus and assigns it to the Status field.
func (o *PaymentIntentQueryParams) SetStatus(v PaymentIntentStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *PaymentIntentQueryParams) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *PaymentIntentQueryParams) UnsetStatus() {
	o.Status.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntentQueryParams) GetUpdatedAt() DateTimeFilter {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret DateTimeFilter
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntentQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PaymentIntentQueryParams) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableDateTimeFilter and assigns it to the UpdatedAt field.
func (o *PaymentIntentQueryParams) SetUpdatedAt(v DateTimeFilter) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *PaymentIntentQueryParams) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *PaymentIntentQueryParams) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o PaymentIntentQueryParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentIntentQueryParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AmountAtom.IsSet() {
		toSerialize["amount_atom"] = o.AmountAtom.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
	}
	if !IsNil(o.Expand) {
		toSerialize["expand"] = o.Expand
	}
	if o.InvoiceId.IsSet() {
		toSerialize["invoice_id"] = o.InvoiceId.Get()
	}
	if !IsNil(o.PageNumber) {
		toSerialize["page_number"] = o.PageNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PaymentMethodId.IsSet() {
		toSerialize["payment_method_id"] = o.PaymentMethodId.Get()
	}
	if o.Search.IsSet() {
		toSerialize["search"] = o.Search.Get()
	}
	if !IsNil(o.SortDescending) {
		toSerialize["sort_descending"] = o.SortDescending
	}
	if !IsNil(o.SortKey) {
		toSerialize["sort_key"] = o.SortKey
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	return toSerialize, nil
}

type NullablePaymentIntentQueryParams struct {
	value *PaymentIntentQueryParams
	isSet bool
}

func (v NullablePaymentIntentQueryParams) Get() *PaymentIntentQueryParams {
	return v.value
}

func (v *NullablePaymentIntentQueryParams) Set(val *PaymentIntentQueryParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentIntentQueryParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentIntentQueryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentIntentQueryParams(val *PaymentIntentQueryParams) *NullablePaymentIntentQueryParams {
	return &NullablePaymentIntentQueryParams{value: val, isSet: true}
}

func (v NullablePaymentIntentQueryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentIntentQueryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


