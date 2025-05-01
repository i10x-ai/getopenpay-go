# PaymentMethodMappingExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Id** | **string** | Unique identifier of the payment method mapping. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**PaymentMethodId** | **string** | External id for payment method | 
**ProcessorId** | **string** | External id for processor | 
**ProcessorName** | [**PaymentProcessorName**](PaymentProcessorName.md) | Name of the payment processor. | 
**TheirPaymentMethodId** | **string** | The payment method identifier from the processor&#39;s perspective. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewPaymentMethodMappingExternal

`func NewPaymentMethodMappingExternal(createdAt time.Time, id string, paymentMethodId string, processorId string, processorName PaymentProcessorName, theirPaymentMethodId string, updatedAt time.Time, ) *PaymentMethodMappingExternal`

NewPaymentMethodMappingExternal instantiates a new PaymentMethodMappingExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodMappingExternalWithDefaults

`func NewPaymentMethodMappingExternalWithDefaults() *PaymentMethodMappingExternal`

NewPaymentMethodMappingExternalWithDefaults instantiates a new PaymentMethodMappingExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PaymentMethodMappingExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethodMappingExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethodMappingExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *PaymentMethodMappingExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodMappingExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodMappingExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *PaymentMethodMappingExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PaymentMethodMappingExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PaymentMethodMappingExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PaymentMethodMappingExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetObject

`func (o *PaymentMethodMappingExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodMappingExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodMappingExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *PaymentMethodMappingExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *PaymentMethodMappingExternal) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *PaymentMethodMappingExternal) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *PaymentMethodMappingExternal) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### GetProcessorId

`func (o *PaymentMethodMappingExternal) GetProcessorId() string`

GetProcessorId returns the ProcessorId field if non-nil, zero value otherwise.

### GetProcessorIdOk

`func (o *PaymentMethodMappingExternal) GetProcessorIdOk() (*string, bool)`

GetProcessorIdOk returns a tuple with the ProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorId

`func (o *PaymentMethodMappingExternal) SetProcessorId(v string)`

SetProcessorId sets ProcessorId field to given value.


### GetProcessorName

`func (o *PaymentMethodMappingExternal) GetProcessorName() PaymentProcessorName`

GetProcessorName returns the ProcessorName field if non-nil, zero value otherwise.

### GetProcessorNameOk

`func (o *PaymentMethodMappingExternal) GetProcessorNameOk() (*PaymentProcessorName, bool)`

GetProcessorNameOk returns a tuple with the ProcessorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorName

`func (o *PaymentMethodMappingExternal) SetProcessorName(v PaymentProcessorName)`

SetProcessorName sets ProcessorName field to given value.


### GetTheirPaymentMethodId

`func (o *PaymentMethodMappingExternal) GetTheirPaymentMethodId() string`

GetTheirPaymentMethodId returns the TheirPaymentMethodId field if non-nil, zero value otherwise.

### GetTheirPaymentMethodIdOk

`func (o *PaymentMethodMappingExternal) GetTheirPaymentMethodIdOk() (*string, bool)`

GetTheirPaymentMethodIdOk returns a tuple with the TheirPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheirPaymentMethodId

`func (o *PaymentMethodMappingExternal) SetTheirPaymentMethodId(v string)`

SetTheirPaymentMethodId sets TheirPaymentMethodId field to given value.


### GetUpdatedAt

`func (o *PaymentMethodMappingExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentMethodMappingExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentMethodMappingExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


