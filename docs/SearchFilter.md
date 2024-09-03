# SearchFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ilike** | Pointer to **NullableString** |  | [optional] 
**Eq** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSearchFilter

`func NewSearchFilter() *SearchFilter`

NewSearchFilter instantiates a new SearchFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFilterWithDefaults

`func NewSearchFilterWithDefaults() *SearchFilter`

NewSearchFilterWithDefaults instantiates a new SearchFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIlike

`func (o *SearchFilter) GetIlike() string`

GetIlike returns the Ilike field if non-nil, zero value otherwise.

### GetIlikeOk

`func (o *SearchFilter) GetIlikeOk() (*string, bool)`

GetIlikeOk returns a tuple with the Ilike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIlike

`func (o *SearchFilter) SetIlike(v string)`

SetIlike sets Ilike field to given value.

### HasIlike

`func (o *SearchFilter) HasIlike() bool`

HasIlike returns a boolean if a field has been set.

### SetIlikeNil

`func (o *SearchFilter) SetIlikeNil(b bool)`

 SetIlikeNil sets the value for Ilike to be an explicit nil

### UnsetIlike
`func (o *SearchFilter) UnsetIlike()`

UnsetIlike ensures that no value is present for Ilike, not even an explicit nil
### GetEq

`func (o *SearchFilter) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *SearchFilter) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *SearchFilter) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *SearchFilter) HasEq() bool`

HasEq returns a boolean if a field has been set.

### SetEqNil

`func (o *SearchFilter) SetEqNil(b bool)`

 SetEqNil sets the value for Eq to be an explicit nil

### UnsetEq
`func (o *SearchFilter) UnsetEq()`

UnsetEq ensures that no value is present for Eq, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


