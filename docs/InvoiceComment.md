# InvoiceComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | Status update. | 
**CreatedAt** | **time.Time** | Creation date of the comment. | 
**UserEmail** | **NullableString** |  | 

## Methods

### NewInvoiceComment

`func NewInvoiceComment(comment string, createdAt time.Time, userEmail NullableString, ) *InvoiceComment`

NewInvoiceComment instantiates a new InvoiceComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCommentWithDefaults

`func NewInvoiceCommentWithDefaults() *InvoiceComment`

NewInvoiceCommentWithDefaults instantiates a new InvoiceComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *InvoiceComment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InvoiceComment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InvoiceComment) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCreatedAt

`func (o *InvoiceComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserEmail

`func (o *InvoiceComment) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *InvoiceComment) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *InvoiceComment) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### SetUserEmailNil

`func (o *InvoiceComment) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *InvoiceComment) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


