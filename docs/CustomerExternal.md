# CustomerExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the customer. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_CUSTOMER]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**AccountId** | **string** | Unique identifier for the account. | 
**Email** | **string** | Customer’s email address. | 
**FirstName** | **NullableString** |  | 
**LastName** | **NullableString** |  | 
**Address** | [**NullableCompleteAddress**](CompleteAddress.md) |  | 
**Subscriptions** | Pointer to [**[]SubscriptionExternal**](SubscriptionExternal.md) |  | [optional] 
**BalanceAtoms** | Pointer to [**[]CustomerBalanceExternal**](CustomerBalanceExternal.md) |  | [optional] 
**SubscribedToProducts** | Pointer to [**[]ProductExternal**](ProductExternal.md) |  | [optional] 
**LastSuccessfulPaymentIntent** | Pointer to [**NullablePaymentIntentExternal**](PaymentIntentExternal.md) |  | [optional] 
**Discount** | Pointer to [**NullableDiscountExternal**](DiscountExternal.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Notes** | **NullableString** |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCustomerExternal

`func NewCustomerExternal(id string, createdAt time.Time, updatedAt time.Time, accountId string, email string, firstName NullableString, lastName NullableString, address NullableCompleteAddress, notes NullableString, ) *CustomerExternal`

NewCustomerExternal instantiates a new CustomerExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerExternalWithDefaults

`func NewCustomerExternalWithDefaults() *CustomerExternal`

NewCustomerExternalWithDefaults instantiates a new CustomerExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CustomerExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *CustomerExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CustomerExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *CustomerExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CustomerExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CustomerExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CustomerExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetAccountId

`func (o *CustomerExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CustomerExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CustomerExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetEmail

`func (o *CustomerExternal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerExternal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerExternal) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *CustomerExternal) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerExternal) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerExternal) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *CustomerExternal) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *CustomerExternal) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *CustomerExternal) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerExternal) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerExternal) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *CustomerExternal) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *CustomerExternal) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAddress

`func (o *CustomerExternal) GetAddress() CompleteAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerExternal) GetAddressOk() (*CompleteAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerExternal) SetAddress(v CompleteAddress)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *CustomerExternal) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *CustomerExternal) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetSubscriptions

`func (o *CustomerExternal) GetSubscriptions() []SubscriptionExternal`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *CustomerExternal) GetSubscriptionsOk() (*[]SubscriptionExternal, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *CustomerExternal) SetSubscriptions(v []SubscriptionExternal)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *CustomerExternal) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetBalanceAtoms

`func (o *CustomerExternal) GetBalanceAtoms() []CustomerBalanceExternal`

GetBalanceAtoms returns the BalanceAtoms field if non-nil, zero value otherwise.

### GetBalanceAtomsOk

`func (o *CustomerExternal) GetBalanceAtomsOk() (*[]CustomerBalanceExternal, bool)`

GetBalanceAtomsOk returns a tuple with the BalanceAtoms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAtoms

`func (o *CustomerExternal) SetBalanceAtoms(v []CustomerBalanceExternal)`

SetBalanceAtoms sets BalanceAtoms field to given value.

### HasBalanceAtoms

`func (o *CustomerExternal) HasBalanceAtoms() bool`

HasBalanceAtoms returns a boolean if a field has been set.

### GetSubscribedToProducts

`func (o *CustomerExternal) GetSubscribedToProducts() []ProductExternal`

GetSubscribedToProducts returns the SubscribedToProducts field if non-nil, zero value otherwise.

### GetSubscribedToProductsOk

`func (o *CustomerExternal) GetSubscribedToProductsOk() (*[]ProductExternal, bool)`

GetSubscribedToProductsOk returns a tuple with the SubscribedToProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedToProducts

`func (o *CustomerExternal) SetSubscribedToProducts(v []ProductExternal)`

SetSubscribedToProducts sets SubscribedToProducts field to given value.

### HasSubscribedToProducts

`func (o *CustomerExternal) HasSubscribedToProducts() bool`

HasSubscribedToProducts returns a boolean if a field has been set.

### GetLastSuccessfulPaymentIntent

`func (o *CustomerExternal) GetLastSuccessfulPaymentIntent() PaymentIntentExternal`

GetLastSuccessfulPaymentIntent returns the LastSuccessfulPaymentIntent field if non-nil, zero value otherwise.

### GetLastSuccessfulPaymentIntentOk

`func (o *CustomerExternal) GetLastSuccessfulPaymentIntentOk() (*PaymentIntentExternal, bool)`

GetLastSuccessfulPaymentIntentOk returns a tuple with the LastSuccessfulPaymentIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulPaymentIntent

`func (o *CustomerExternal) SetLastSuccessfulPaymentIntent(v PaymentIntentExternal)`

SetLastSuccessfulPaymentIntent sets LastSuccessfulPaymentIntent field to given value.

### HasLastSuccessfulPaymentIntent

`func (o *CustomerExternal) HasLastSuccessfulPaymentIntent() bool`

HasLastSuccessfulPaymentIntent returns a boolean if a field has been set.

### SetLastSuccessfulPaymentIntentNil

`func (o *CustomerExternal) SetLastSuccessfulPaymentIntentNil(b bool)`

 SetLastSuccessfulPaymentIntentNil sets the value for LastSuccessfulPaymentIntent to be an explicit nil

### UnsetLastSuccessfulPaymentIntent
`func (o *CustomerExternal) UnsetLastSuccessfulPaymentIntent()`

UnsetLastSuccessfulPaymentIntent ensures that no value is present for LastSuccessfulPaymentIntent, not even an explicit nil
### GetDiscount

`func (o *CustomerExternal) GetDiscount() DiscountExternal`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *CustomerExternal) GetDiscountOk() (*DiscountExternal, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *CustomerExternal) SetDiscount(v DiscountExternal)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *CustomerExternal) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *CustomerExternal) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *CustomerExternal) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetMetadata

`func (o *CustomerExternal) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerExternal) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerExternal) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerExternal) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CustomerExternal) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CustomerExternal) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetNotes

`func (o *CustomerExternal) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CustomerExternal) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CustomerExternal) SetNotes(v string)`

SetNotes sets Notes field to given value.


### SetNotesNil

`func (o *CustomerExternal) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *CustomerExternal) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetCustomFields

`func (o *CustomerExternal) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CustomerExternal) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CustomerExternal) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CustomerExternal) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CustomerExternal) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CustomerExternal) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


