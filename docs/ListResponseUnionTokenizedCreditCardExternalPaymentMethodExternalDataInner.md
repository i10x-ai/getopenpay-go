# ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to [**CompleteAddress**](CompleteAddress.md) |  | [optional] 
**CardBrand** | Pointer to **string** |  | [optional] 
**CardCountry** | Pointer to **string** |  | [optional] 
**CardFingerprint** | Pointer to **string** |  | [optional] 
**CardIin** | Pointer to **string** |  | [optional] 
**CardIssuer** | Pointer to **string** |  | [optional] 
**CardType** | Pointer to **string** |  | [optional] 
**CdeExternalId** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**CustomerId** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** | Display name for the payment method to show on the UI. | 
**ExpiryDate** | **string** |  | 
**Id** | **string** | Unique Identifier of the payment_method. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**IsFullDetailsKnown** | **bool** | Indicates if full details of the card are known | 
**LastFour** | Pointer to **string** |  | [optional] 
**Mappings** | Pointer to [**[]PaymentMethodMappingExternal**](PaymentMethodMappingExternal.md) | List of payment method mappings associated with this payment method. Include \&quot;mappings\&quot; in the expand parameter to retrieve this data. | [optional] [default to []]
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**Provider** | [**PaymentProviderType**](PaymentProviderType.md) | The provider type for this payment method | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner

`func NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner(createdAt time.Time, displayName string, expiryDate string, id string, isFullDetailsKnown bool, provider PaymentProviderType, updatedAt time.Time, ) *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner`

NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner instantiates a new ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInnerWithDefaults

`func NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInnerWithDefaults() *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner`

NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInnerWithDefaults instantiates a new ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetBillingAddress() CompleteAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetBillingAddressOk() (*CompleteAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetBillingAddress(v CompleteAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCardBrand

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### GetCardCountry

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### GetCardFingerprint

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardFingerprint() string`

GetCardFingerprint returns the CardFingerprint field if non-nil, zero value otherwise.

### GetCardFingerprintOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardFingerprintOk() (*string, bool)`

GetCardFingerprintOk returns a tuple with the CardFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFingerprint

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardFingerprint(v string)`

SetCardFingerprint sets CardFingerprint field to given value.

### HasCardFingerprint

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardFingerprint() bool`

HasCardFingerprint returns a boolean if a field has been set.

### GetCardIin

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardIin() string`

GetCardIin returns the CardIin field if non-nil, zero value otherwise.

### GetCardIinOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardIinOk() (*string, bool)`

GetCardIinOk returns a tuple with the CardIin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIin

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardIin(v string)`

SetCardIin sets CardIin field to given value.

### HasCardIin

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardIin() bool`

HasCardIin returns a boolean if a field has been set.

### GetCardIssuer

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardIssuer() string`

GetCardIssuer returns the CardIssuer field if non-nil, zero value otherwise.

### GetCardIssuerOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardIssuerOk() (*string, bool)`

GetCardIssuerOk returns a tuple with the CardIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIssuer

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardIssuer(v string)`

SetCardIssuer sets CardIssuer field to given value.

### HasCardIssuer

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardIssuer() bool`

HasCardIssuer returns a boolean if a field has been set.

### GetCardType

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCdeExternalId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCdeExternalId() string`

GetCdeExternalId returns the CdeExternalId field if non-nil, zero value otherwise.

### GetCdeExternalIdOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCdeExternalIdOk() (*string, bool)`

GetCdeExternalIdOk returns a tuple with the CdeExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdeExternalId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCdeExternalId(v string)`

SetCdeExternalId sets CdeExternalId field to given value.

### HasCdeExternalId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCdeExternalId() bool`

HasCdeExternalId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCustomerId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetExpiryDate

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsFullDetailsKnown

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetIsFullDetailsKnown() bool`

GetIsFullDetailsKnown returns the IsFullDetailsKnown field if non-nil, zero value otherwise.

### GetIsFullDetailsKnownOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetIsFullDetailsKnownOk() (*bool, bool)`

GetIsFullDetailsKnownOk returns a tuple with the IsFullDetailsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullDetailsKnown

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetIsFullDetailsKnown(v bool)`

SetIsFullDetailsKnown sets IsFullDetailsKnown field to given value.


### GetLastFour

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### GetMappings

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetMappings() []PaymentMethodMappingExternal`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetMappingsOk() (*[]PaymentMethodMappingExternal, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetMappings(v []PaymentMethodMappingExternal)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetMetadata

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetObject

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetProvider

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetProvider() PaymentProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetProviderOk() (*PaymentProviderType, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetProvider(v PaymentProviderType)`

SetProvider sets Provider field to given value.


### GetUpdatedAt

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


