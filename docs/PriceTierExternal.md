# PriceTierExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**FlatAmountAtom** | **int32** |  | 
**Id** | **string** | Unique identifier for the object. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**UnitAmountAtom** | **int32** |  | 
**UnitsUpto** | Pointer to **NullableInt32** |  | [optional] 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewPriceTierExternal

`func NewPriceTierExternal(createdAt time.Time, flatAmountAtom int32, id string, unitAmountAtom int32, updatedAt time.Time, ) *PriceTierExternal`

NewPriceTierExternal instantiates a new PriceTierExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTierExternalWithDefaults

`func NewPriceTierExternalWithDefaults() *PriceTierExternal`

NewPriceTierExternalWithDefaults instantiates a new PriceTierExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PriceTierExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PriceTierExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PriceTierExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFlatAmountAtom

`func (o *PriceTierExternal) GetFlatAmountAtom() int32`

GetFlatAmountAtom returns the FlatAmountAtom field if non-nil, zero value otherwise.

### GetFlatAmountAtomOk

`func (o *PriceTierExternal) GetFlatAmountAtomOk() (*int32, bool)`

GetFlatAmountAtomOk returns a tuple with the FlatAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatAmountAtom

`func (o *PriceTierExternal) SetFlatAmountAtom(v int32)`

SetFlatAmountAtom sets FlatAmountAtom field to given value.


### GetId

`func (o *PriceTierExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceTierExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceTierExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *PriceTierExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PriceTierExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PriceTierExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PriceTierExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetObject

`func (o *PriceTierExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PriceTierExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PriceTierExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *PriceTierExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUnitAmountAtom

`func (o *PriceTierExternal) GetUnitAmountAtom() int32`

GetUnitAmountAtom returns the UnitAmountAtom field if non-nil, zero value otherwise.

### GetUnitAmountAtomOk

`func (o *PriceTierExternal) GetUnitAmountAtomOk() (*int32, bool)`

GetUnitAmountAtomOk returns a tuple with the UnitAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountAtom

`func (o *PriceTierExternal) SetUnitAmountAtom(v int32)`

SetUnitAmountAtom sets UnitAmountAtom field to given value.


### GetUnitsUpto

`func (o *PriceTierExternal) GetUnitsUpto() int32`

GetUnitsUpto returns the UnitsUpto field if non-nil, zero value otherwise.

### GetUnitsUptoOk

`func (o *PriceTierExternal) GetUnitsUptoOk() (*int32, bool)`

GetUnitsUptoOk returns a tuple with the UnitsUpto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsUpto

`func (o *PriceTierExternal) SetUnitsUpto(v int32)`

SetUnitsUpto sets UnitsUpto field to given value.

### HasUnitsUpto

`func (o *PriceTierExternal) HasUnitsUpto() bool`

HasUnitsUpto returns a boolean if a field has been set.

### SetUnitsUptoNil

`func (o *PriceTierExternal) SetUnitsUptoNil(b bool)`

 SetUnitsUptoNil sets the value for UnitsUpto to be an explicit nil

### UnsetUnitsUpto
`func (o *PriceTierExternal) UnsetUnitsUpto()`

UnsetUnitsUpto ensures that no value is present for UnitsUpto, not even an explicit nil
### GetUpdatedAt

`func (o *PriceTierExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PriceTierExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PriceTierExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


