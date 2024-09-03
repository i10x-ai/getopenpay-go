# DeletePriceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message describing result of API call. | [optional] [default to "Price deleted Successfully."]
**PriceId** | **string** | Unique identifier of the price. | 

## Methods

### NewDeletePriceResponse

`func NewDeletePriceResponse(priceId string, ) *DeletePriceResponse`

NewDeletePriceResponse instantiates a new DeletePriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePriceResponseWithDefaults

`func NewDeletePriceResponseWithDefaults() *DeletePriceResponse`

NewDeletePriceResponseWithDefaults instantiates a new DeletePriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeletePriceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeletePriceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeletePriceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeletePriceResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPriceId

`func (o *DeletePriceResponse) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *DeletePriceResponse) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *DeletePriceResponse) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


