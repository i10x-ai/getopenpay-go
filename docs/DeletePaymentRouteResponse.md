# DeletePaymentRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message describing result of API call. | [optional] [default to "Payment route deleted successfully"]
**PaymentRouteId** | **string** | Unique identifier of the payment route. | 

## Methods

### NewDeletePaymentRouteResponse

`func NewDeletePaymentRouteResponse(paymentRouteId string, ) *DeletePaymentRouteResponse`

NewDeletePaymentRouteResponse instantiates a new DeletePaymentRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePaymentRouteResponseWithDefaults

`func NewDeletePaymentRouteResponseWithDefaults() *DeletePaymentRouteResponse`

NewDeletePaymentRouteResponseWithDefaults instantiates a new DeletePaymentRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeletePaymentRouteResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeletePaymentRouteResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeletePaymentRouteResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeletePaymentRouteResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPaymentRouteId

`func (o *DeletePaymentRouteResponse) GetPaymentRouteId() string`

GetPaymentRouteId returns the PaymentRouteId field if non-nil, zero value otherwise.

### GetPaymentRouteIdOk

`func (o *DeletePaymentRouteResponse) GetPaymentRouteIdOk() (*string, bool)`

GetPaymentRouteIdOk returns a tuple with the PaymentRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRouteId

`func (o *DeletePaymentRouteResponse) SetPaymentRouteId(v string)`

SetPaymentRouteId sets PaymentRouteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


