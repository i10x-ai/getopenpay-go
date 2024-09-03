# ListResponseInvoiceExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]InvoiceExternal**](InvoiceExternal.md) |  | 
**TotalObjects** | **int32** |  | 
**PageNumber** | **int32** |  | 
**PageSize** | **int32** |  | 

## Methods

### NewListResponseInvoiceExternal

`func NewListResponseInvoiceExternal(data []InvoiceExternal, totalObjects int32, pageNumber int32, pageSize int32, ) *ListResponseInvoiceExternal`

NewListResponseInvoiceExternal instantiates a new ListResponseInvoiceExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResponseInvoiceExternalWithDefaults

`func NewListResponseInvoiceExternalWithDefaults() *ListResponseInvoiceExternal`

NewListResponseInvoiceExternalWithDefaults instantiates a new ListResponseInvoiceExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListResponseInvoiceExternal) GetData() []InvoiceExternal`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListResponseInvoiceExternal) GetDataOk() (*[]InvoiceExternal, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListResponseInvoiceExternal) SetData(v []InvoiceExternal)`

SetData sets Data field to given value.


### GetTotalObjects

`func (o *ListResponseInvoiceExternal) GetTotalObjects() int32`

GetTotalObjects returns the TotalObjects field if non-nil, zero value otherwise.

### GetTotalObjectsOk

`func (o *ListResponseInvoiceExternal) GetTotalObjectsOk() (*int32, bool)`

GetTotalObjectsOk returns a tuple with the TotalObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjects

`func (o *ListResponseInvoiceExternal) SetTotalObjects(v int32)`

SetTotalObjects sets TotalObjects field to given value.


### GetPageNumber

`func (o *ListResponseInvoiceExternal) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ListResponseInvoiceExternal) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ListResponseInvoiceExternal) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *ListResponseInvoiceExternal) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListResponseInvoiceExternal) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListResponseInvoiceExternal) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


