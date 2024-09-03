# DeleteProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message describing result of API call. | [optional] [default to "Product deleted Successfully."]
**ProductId** | **string** | Unique identifier of the product. | 

## Methods

### NewDeleteProductResponse

`func NewDeleteProductResponse(productId string, ) *DeleteProductResponse`

NewDeleteProductResponse instantiates a new DeleteProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteProductResponseWithDefaults

`func NewDeleteProductResponseWithDefaults() *DeleteProductResponse`

NewDeleteProductResponseWithDefaults instantiates a new DeleteProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeleteProductResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteProductResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteProductResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeleteProductResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProductId

`func (o *DeleteProductResponse) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *DeleteProductResponse) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *DeleteProductResponse) SetProductId(v string)`

SetProductId sets ProductId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


