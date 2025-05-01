# ListResponsePaymentLinkExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PaymentLinkExternal**](PaymentLinkExternal.md) |  | 
**PageNumber** | **int32** |  | 
**PageSize** | **int32** |  | 
**TotalObjects** | **int32** |  | 

## Methods

### NewListResponsePaymentLinkExternal

`func NewListResponsePaymentLinkExternal(data []PaymentLinkExternal, pageNumber int32, pageSize int32, totalObjects int32, ) *ListResponsePaymentLinkExternal`

NewListResponsePaymentLinkExternal instantiates a new ListResponsePaymentLinkExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResponsePaymentLinkExternalWithDefaults

`func NewListResponsePaymentLinkExternalWithDefaults() *ListResponsePaymentLinkExternal`

NewListResponsePaymentLinkExternalWithDefaults instantiates a new ListResponsePaymentLinkExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListResponsePaymentLinkExternal) GetData() []PaymentLinkExternal`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListResponsePaymentLinkExternal) GetDataOk() (*[]PaymentLinkExternal, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListResponsePaymentLinkExternal) SetData(v []PaymentLinkExternal)`

SetData sets Data field to given value.


### GetPageNumber

`func (o *ListResponsePaymentLinkExternal) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ListResponsePaymentLinkExternal) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ListResponsePaymentLinkExternal) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *ListResponsePaymentLinkExternal) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListResponsePaymentLinkExternal) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListResponsePaymentLinkExternal) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalObjects

`func (o *ListResponsePaymentLinkExternal) GetTotalObjects() int32`

GetTotalObjects returns the TotalObjects field if non-nil, zero value otherwise.

### GetTotalObjectsOk

`func (o *ListResponsePaymentLinkExternal) GetTotalObjectsOk() (*int32, bool)`

GetTotalObjectsOk returns a tuple with the TotalObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjects

`func (o *ListResponsePaymentLinkExternal) SetTotalObjects(v int32)`

SetTotalObjects sets TotalObjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


