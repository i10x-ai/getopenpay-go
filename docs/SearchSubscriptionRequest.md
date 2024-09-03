# SearchSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**Query** | **string** | The search query string in limited Lucene Query Syntax. Query is case insensitive.     Limitations:        1. You can combine multiple query clauses in a search by either separating them with a space, or using the AND or OR keywords (case insensitive). By default, the API combines clauses with AND logic. AND and OR operators are mutually exclusive.       2. Nesting of Queries by ( ) is not supported.     Example searches.       1. search by exact field. example: &#x60;name:&#39;field_name&#39;&#x60;       2. search by wildcard. example: &#x60;name:&#39;*sub_string_1*&#39;&#x60;, &#x60;name:&#39;prefix_string*&#39;&#x60;       3. range search using [], both bounds are included in result. example: &#x60;amount:[100 TO *]&#x60;, &#x60;amount:[10 TO 100]&#x60;.      | 

## Methods

### NewSearchSubscriptionRequest

`func NewSearchSubscriptionRequest(query string, ) *SearchSubscriptionRequest`

NewSearchSubscriptionRequest instantiates a new SearchSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSubscriptionRequestWithDefaults

`func NewSearchSubscriptionRequestWithDefaults() *SearchSubscriptionRequest`

NewSearchSubscriptionRequestWithDefaults instantiates a new SearchSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *SearchSubscriptionRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *SearchSubscriptionRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *SearchSubscriptionRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *SearchSubscriptionRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *SearchSubscriptionRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SearchSubscriptionRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SearchSubscriptionRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *SearchSubscriptionRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetQuery

`func (o *SearchSubscriptionRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchSubscriptionRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchSubscriptionRequest) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


