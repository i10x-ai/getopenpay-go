# PriceTierPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitsUpto** | Pointer to **NullableInt32** |  | [optional] 
**UnitAmountAtom** | **int32** | Price per unit in the currency atom. | 
**FlatAmountAtom** | **int32** | Flat price in the currency atom. | 

## Methods

### NewPriceTierPublic

`func NewPriceTierPublic(unitAmountAtom int32, flatAmountAtom int32, ) *PriceTierPublic`

NewPriceTierPublic instantiates a new PriceTierPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTierPublicWithDefaults

`func NewPriceTierPublicWithDefaults() *PriceTierPublic`

NewPriceTierPublicWithDefaults instantiates a new PriceTierPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitsUpto

`func (o *PriceTierPublic) GetUnitsUpto() int32`

GetUnitsUpto returns the UnitsUpto field if non-nil, zero value otherwise.

### GetUnitsUptoOk

`func (o *PriceTierPublic) GetUnitsUptoOk() (*int32, bool)`

GetUnitsUptoOk returns a tuple with the UnitsUpto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsUpto

`func (o *PriceTierPublic) SetUnitsUpto(v int32)`

SetUnitsUpto sets UnitsUpto field to given value.

### HasUnitsUpto

`func (o *PriceTierPublic) HasUnitsUpto() bool`

HasUnitsUpto returns a boolean if a field has been set.

### SetUnitsUptoNil

`func (o *PriceTierPublic) SetUnitsUptoNil(b bool)`

 SetUnitsUptoNil sets the value for UnitsUpto to be an explicit nil

### UnsetUnitsUpto
`func (o *PriceTierPublic) UnsetUnitsUpto()`

UnsetUnitsUpto ensures that no value is present for UnitsUpto, not even an explicit nil
### GetUnitAmountAtom

`func (o *PriceTierPublic) GetUnitAmountAtom() int32`

GetUnitAmountAtom returns the UnitAmountAtom field if non-nil, zero value otherwise.

### GetUnitAmountAtomOk

`func (o *PriceTierPublic) GetUnitAmountAtomOk() (*int32, bool)`

GetUnitAmountAtomOk returns a tuple with the UnitAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountAtom

`func (o *PriceTierPublic) SetUnitAmountAtom(v int32)`

SetUnitAmountAtom sets UnitAmountAtom field to given value.


### GetFlatAmountAtom

`func (o *PriceTierPublic) GetFlatAmountAtom() int32`

GetFlatAmountAtom returns the FlatAmountAtom field if non-nil, zero value otherwise.

### GetFlatAmountAtomOk

`func (o *PriceTierPublic) GetFlatAmountAtomOk() (*int32, bool)`

GetFlatAmountAtomOk returns a tuple with the FlatAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatAmountAtom

`func (o *PriceTierPublic) SetFlatAmountAtom(v int32)`

SetFlatAmountAtom sets FlatAmountAtom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


