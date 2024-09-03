# PriceTierParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitsFrom** | **int32** | The starting unit for the price tier. | 
**UnitsUpto** | Pointer to **NullableInt32** |  | [optional] 
**UnitAmountAtom** | **int32** | The price for the unit in the smallest denomination.It is in atomic units (in USD this is cents). | 
**FlatAmountAtom** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPriceTierParams

`func NewPriceTierParams(unitsFrom int32, unitAmountAtom int32, ) *PriceTierParams`

NewPriceTierParams instantiates a new PriceTierParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTierParamsWithDefaults

`func NewPriceTierParamsWithDefaults() *PriceTierParams`

NewPriceTierParamsWithDefaults instantiates a new PriceTierParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitsFrom

`func (o *PriceTierParams) GetUnitsFrom() int32`

GetUnitsFrom returns the UnitsFrom field if non-nil, zero value otherwise.

### GetUnitsFromOk

`func (o *PriceTierParams) GetUnitsFromOk() (*int32, bool)`

GetUnitsFromOk returns a tuple with the UnitsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsFrom

`func (o *PriceTierParams) SetUnitsFrom(v int32)`

SetUnitsFrom sets UnitsFrom field to given value.


### GetUnitsUpto

`func (o *PriceTierParams) GetUnitsUpto() int32`

GetUnitsUpto returns the UnitsUpto field if non-nil, zero value otherwise.

### GetUnitsUptoOk

`func (o *PriceTierParams) GetUnitsUptoOk() (*int32, bool)`

GetUnitsUptoOk returns a tuple with the UnitsUpto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsUpto

`func (o *PriceTierParams) SetUnitsUpto(v int32)`

SetUnitsUpto sets UnitsUpto field to given value.

### HasUnitsUpto

`func (o *PriceTierParams) HasUnitsUpto() bool`

HasUnitsUpto returns a boolean if a field has been set.

### SetUnitsUptoNil

`func (o *PriceTierParams) SetUnitsUptoNil(b bool)`

 SetUnitsUptoNil sets the value for UnitsUpto to be an explicit nil

### UnsetUnitsUpto
`func (o *PriceTierParams) UnsetUnitsUpto()`

UnsetUnitsUpto ensures that no value is present for UnitsUpto, not even an explicit nil
### GetUnitAmountAtom

`func (o *PriceTierParams) GetUnitAmountAtom() int32`

GetUnitAmountAtom returns the UnitAmountAtom field if non-nil, zero value otherwise.

### GetUnitAmountAtomOk

`func (o *PriceTierParams) GetUnitAmountAtomOk() (*int32, bool)`

GetUnitAmountAtomOk returns a tuple with the UnitAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountAtom

`func (o *PriceTierParams) SetUnitAmountAtom(v int32)`

SetUnitAmountAtom sets UnitAmountAtom field to given value.


### GetFlatAmountAtom

`func (o *PriceTierParams) GetFlatAmountAtom() int32`

GetFlatAmountAtom returns the FlatAmountAtom field if non-nil, zero value otherwise.

### GetFlatAmountAtomOk

`func (o *PriceTierParams) GetFlatAmountAtomOk() (*int32, bool)`

GetFlatAmountAtomOk returns a tuple with the FlatAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatAmountAtom

`func (o *PriceTierParams) SetFlatAmountAtom(v int32)`

SetFlatAmountAtom sets FlatAmountAtom field to given value.

### HasFlatAmountAtom

`func (o *PriceTierParams) HasFlatAmountAtom() bool`

HasFlatAmountAtom returns a boolean if a field has been set.

### SetFlatAmountAtomNil

`func (o *PriceTierParams) SetFlatAmountAtomNil(b bool)`

 SetFlatAmountAtomNil sets the value for FlatAmountAtom to be an explicit nil

### UnsetFlatAmountAtom
`func (o *PriceTierParams) UnsetFlatAmountAtom()`

UnsetFlatAmountAtom ensures that no value is present for FlatAmountAtom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


