# CreatePortalSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The ID of an existing customer. | 
**ReturnUrl** | **string** | The default URL to redirect customers to when they click on the portal&#39;s link to return to your website. | 

## Methods

### NewCreatePortalSessionRequest

`func NewCreatePortalSessionRequest(customerId string, returnUrl string, ) *CreatePortalSessionRequest`

NewCreatePortalSessionRequest instantiates a new CreatePortalSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePortalSessionRequestWithDefaults

`func NewCreatePortalSessionRequestWithDefaults() *CreatePortalSessionRequest`

NewCreatePortalSessionRequestWithDefaults instantiates a new CreatePortalSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CreatePortalSessionRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreatePortalSessionRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreatePortalSessionRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetReturnUrl

`func (o *CreatePortalSessionRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreatePortalSessionRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreatePortalSessionRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


