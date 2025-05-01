# ListResponseInvoiceItemExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]InvoiceItemExternal**](InvoiceItemExternal.md) |  | 
**PageNumber** | **int32** |  | 
**PageSize** | **int32** |  | 
**TotalObjects** | **int32** |  | 

## Methods

### NewListResponseInvoiceItemExternal

`func NewListResponseInvoiceItemExternal(data []InvoiceItemExternal, pageNumber int32, pageSize int32, totalObjects int32, ) *ListResponseInvoiceItemExternal`

NewListResponseInvoiceItemExternal instantiates a new ListResponseInvoiceItemExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResponseInvoiceItemExternalWithDefaults

`func NewListResponseInvoiceItemExternalWithDefaults() *ListResponseInvoiceItemExternal`

NewListResponseInvoiceItemExternalWithDefaults instantiates a new ListResponseInvoiceItemExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListResponseInvoiceItemExternal) GetData() []InvoiceItemExternal`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListResponseInvoiceItemExternal) GetDataOk() (*[]InvoiceItemExternal, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListResponseInvoiceItemExternal) SetData(v []InvoiceItemExternal)`

SetData sets Data field to given value.


### GetPageNumber

`func (o *ListResponseInvoiceItemExternal) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ListResponseInvoiceItemExternal) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ListResponseInvoiceItemExternal) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *ListResponseInvoiceItemExternal) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListResponseInvoiceItemExternal) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListResponseInvoiceItemExternal) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalObjects

`func (o *ListResponseInvoiceItemExternal) GetTotalObjects() int32`

GetTotalObjects returns the TotalObjects field if non-nil, zero value otherwise.

### GetTotalObjectsOk

`func (o *ListResponseInvoiceItemExternal) GetTotalObjectsOk() (*int32, bool)`

GetTotalObjectsOk returns a tuple with the TotalObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjects

`func (o *ListResponseInvoiceItemExternal) SetTotalObjects(v int32)`

SetTotalObjects sets TotalObjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


