# InvoiceQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**SortKey** | Pointer to **string** | Key name based on which data is sorted. | [optional] [default to "created_at"]
**SortDescending** | Pointer to **bool** | Sort direction. | [optional] [default to false]
**CreatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**UpdatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**Expand** | Pointer to **[]string** |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullableInvoiceStatusEnum**](InvoiceStatusEnum.md) |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**BillingReason** | Pointer to [**NullableBillingReasonEnum**](BillingReasonEnum.md) |  | [optional] 
**TotalAmountAtom** | Pointer to [**NullableIntRangeFilter**](IntRangeFilter.md) |  | [optional] 
**CollectionMethod** | Pointer to [**NullableCollectionMethodEnum**](CollectionMethodEnum.md) |  | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) |  | [optional] 
**PeriodStart** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**PeriodEnd** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**CouponId** | Pointer to **NullableString** |  | [optional] 
**Search** | Pointer to [**NullableSearchFilter**](SearchFilter.md) |  | [optional] 

## Methods

### NewInvoiceQueryParams

`func NewInvoiceQueryParams() *InvoiceQueryParams`

NewInvoiceQueryParams instantiates a new InvoiceQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceQueryParamsWithDefaults

`func NewInvoiceQueryParamsWithDefaults() *InvoiceQueryParams`

NewInvoiceQueryParamsWithDefaults instantiates a new InvoiceQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *InvoiceQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *InvoiceQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *InvoiceQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *InvoiceQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *InvoiceQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *InvoiceQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *InvoiceQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *InvoiceQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortKey

`func (o *InvoiceQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *InvoiceQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *InvoiceQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *InvoiceQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetSortDescending

`func (o *InvoiceQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *InvoiceQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *InvoiceQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *InvoiceQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InvoiceQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *InvoiceQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *InvoiceQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *InvoiceQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InvoiceQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *InvoiceQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *InvoiceQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetExpand

`func (o *InvoiceQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *InvoiceQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *InvoiceQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *InvoiceQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetCustomerId

`func (o *InvoiceQueryParams) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InvoiceQueryParams) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InvoiceQueryParams) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InvoiceQueryParams) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *InvoiceQueryParams) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *InvoiceQueryParams) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetStatus

`func (o *InvoiceQueryParams) GetStatus() InvoiceStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceQueryParams) GetStatusOk() (*InvoiceStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceQueryParams) SetStatus(v InvoiceStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvoiceQueryParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *InvoiceQueryParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *InvoiceQueryParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubscriptionId

`func (o *InvoiceQueryParams) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *InvoiceQueryParams) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *InvoiceQueryParams) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *InvoiceQueryParams) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *InvoiceQueryParams) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *InvoiceQueryParams) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetBillingReason

`func (o *InvoiceQueryParams) GetBillingReason() BillingReasonEnum`

GetBillingReason returns the BillingReason field if non-nil, zero value otherwise.

### GetBillingReasonOk

`func (o *InvoiceQueryParams) GetBillingReasonOk() (*BillingReasonEnum, bool)`

GetBillingReasonOk returns a tuple with the BillingReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingReason

`func (o *InvoiceQueryParams) SetBillingReason(v BillingReasonEnum)`

SetBillingReason sets BillingReason field to given value.

### HasBillingReason

`func (o *InvoiceQueryParams) HasBillingReason() bool`

HasBillingReason returns a boolean if a field has been set.

### SetBillingReasonNil

`func (o *InvoiceQueryParams) SetBillingReasonNil(b bool)`

 SetBillingReasonNil sets the value for BillingReason to be an explicit nil

### UnsetBillingReason
`func (o *InvoiceQueryParams) UnsetBillingReason()`

UnsetBillingReason ensures that no value is present for BillingReason, not even an explicit nil
### GetTotalAmountAtom

`func (o *InvoiceQueryParams) GetTotalAmountAtom() IntRangeFilter`

GetTotalAmountAtom returns the TotalAmountAtom field if non-nil, zero value otherwise.

### GetTotalAmountAtomOk

`func (o *InvoiceQueryParams) GetTotalAmountAtomOk() (*IntRangeFilter, bool)`

GetTotalAmountAtomOk returns a tuple with the TotalAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountAtom

`func (o *InvoiceQueryParams) SetTotalAmountAtom(v IntRangeFilter)`

SetTotalAmountAtom sets TotalAmountAtom field to given value.

### HasTotalAmountAtom

`func (o *InvoiceQueryParams) HasTotalAmountAtom() bool`

HasTotalAmountAtom returns a boolean if a field has been set.

### SetTotalAmountAtomNil

`func (o *InvoiceQueryParams) SetTotalAmountAtomNil(b bool)`

 SetTotalAmountAtomNil sets the value for TotalAmountAtom to be an explicit nil

### UnsetTotalAmountAtom
`func (o *InvoiceQueryParams) UnsetTotalAmountAtom()`

UnsetTotalAmountAtom ensures that no value is present for TotalAmountAtom, not even an explicit nil
### GetCollectionMethod

`func (o *InvoiceQueryParams) GetCollectionMethod() CollectionMethodEnum`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *InvoiceQueryParams) GetCollectionMethodOk() (*CollectionMethodEnum, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *InvoiceQueryParams) SetCollectionMethod(v CollectionMethodEnum)`

SetCollectionMethod sets CollectionMethod field to given value.

### HasCollectionMethod

`func (o *InvoiceQueryParams) HasCollectionMethod() bool`

HasCollectionMethod returns a boolean if a field has been set.

### SetCollectionMethodNil

`func (o *InvoiceQueryParams) SetCollectionMethodNil(b bool)`

 SetCollectionMethodNil sets the value for CollectionMethod to be an explicit nil

### UnsetCollectionMethod
`func (o *InvoiceQueryParams) UnsetCollectionMethod()`

UnsetCollectionMethod ensures that no value is present for CollectionMethod, not even an explicit nil
### GetCurrency

`func (o *InvoiceQueryParams) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceQueryParams) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceQueryParams) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceQueryParams) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *InvoiceQueryParams) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *InvoiceQueryParams) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPeriodStart

`func (o *InvoiceQueryParams) GetPeriodStart() DateTimeFilter`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *InvoiceQueryParams) GetPeriodStartOk() (*DateTimeFilter, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *InvoiceQueryParams) SetPeriodStart(v DateTimeFilter)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *InvoiceQueryParams) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### SetPeriodStartNil

`func (o *InvoiceQueryParams) SetPeriodStartNil(b bool)`

 SetPeriodStartNil sets the value for PeriodStart to be an explicit nil

### UnsetPeriodStart
`func (o *InvoiceQueryParams) UnsetPeriodStart()`

UnsetPeriodStart ensures that no value is present for PeriodStart, not even an explicit nil
### GetPeriodEnd

`func (o *InvoiceQueryParams) GetPeriodEnd() DateTimeFilter`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *InvoiceQueryParams) GetPeriodEndOk() (*DateTimeFilter, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *InvoiceQueryParams) SetPeriodEnd(v DateTimeFilter)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *InvoiceQueryParams) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### SetPeriodEndNil

`func (o *InvoiceQueryParams) SetPeriodEndNil(b bool)`

 SetPeriodEndNil sets the value for PeriodEnd to be an explicit nil

### UnsetPeriodEnd
`func (o *InvoiceQueryParams) UnsetPeriodEnd()`

UnsetPeriodEnd ensures that no value is present for PeriodEnd, not even an explicit nil
### GetCouponId

`func (o *InvoiceQueryParams) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *InvoiceQueryParams) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *InvoiceQueryParams) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *InvoiceQueryParams) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *InvoiceQueryParams) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *InvoiceQueryParams) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetSearch

`func (o *InvoiceQueryParams) GetSearch() SearchFilter`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *InvoiceQueryParams) GetSearchOk() (*SearchFilter, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *InvoiceQueryParams) SetSearch(v SearchFilter)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *InvoiceQueryParams) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *InvoiceQueryParams) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *InvoiceQueryParams) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


