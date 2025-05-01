# ListResponseBillingMeterSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BillingMeterSummary**](BillingMeterSummary.md) |  | 
**PageNumber** | **int32** |  | 
**PageSize** | **int32** |  | 
**TotalObjects** | **int32** |  | 

## Methods

### NewListResponseBillingMeterSummary

`func NewListResponseBillingMeterSummary(data []BillingMeterSummary, pageNumber int32, pageSize int32, totalObjects int32, ) *ListResponseBillingMeterSummary`

NewListResponseBillingMeterSummary instantiates a new ListResponseBillingMeterSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResponseBillingMeterSummaryWithDefaults

`func NewListResponseBillingMeterSummaryWithDefaults() *ListResponseBillingMeterSummary`

NewListResponseBillingMeterSummaryWithDefaults instantiates a new ListResponseBillingMeterSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListResponseBillingMeterSummary) GetData() []BillingMeterSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListResponseBillingMeterSummary) GetDataOk() (*[]BillingMeterSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListResponseBillingMeterSummary) SetData(v []BillingMeterSummary)`

SetData sets Data field to given value.


### GetPageNumber

`func (o *ListResponseBillingMeterSummary) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ListResponseBillingMeterSummary) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ListResponseBillingMeterSummary) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *ListResponseBillingMeterSummary) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListResponseBillingMeterSummary) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListResponseBillingMeterSummary) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalObjects

`func (o *ListResponseBillingMeterSummary) GetTotalObjects() int32`

GetTotalObjects returns the TotalObjects field if non-nil, zero value otherwise.

### GetTotalObjectsOk

`func (o *ListResponseBillingMeterSummary) GetTotalObjectsOk() (*int32, bool)`

GetTotalObjectsOk returns a tuple with the TotalObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjects

`func (o *ListResponseBillingMeterSummary) SetTotalObjects(v int32)`

SetTotalObjects sets TotalObjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


