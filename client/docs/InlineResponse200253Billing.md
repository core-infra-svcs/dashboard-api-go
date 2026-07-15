# InlineResponse200253Billing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeAccess** | Pointer to [**InlineResponse200253BillingFreeAccess**](InlineResponse200253BillingFreeAccess.md) |  | [optional] 
**PrepaidAccessFastLoginEnabled** | Pointer to **bool** | Whether or not billing uses the fast login prepaid access option. | [optional] 
**ReplyToEmailAddress** | Pointer to **string** | The email address that reeceives replies from clients | [optional] 

## Methods

### NewInlineResponse200253Billing

`func NewInlineResponse200253Billing() *InlineResponse200253Billing`

NewInlineResponse200253Billing instantiates a new InlineResponse200253Billing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200253BillingWithDefaults

`func NewInlineResponse200253BillingWithDefaults() *InlineResponse200253Billing`

NewInlineResponse200253BillingWithDefaults instantiates a new InlineResponse200253Billing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeAccess

`func (o *InlineResponse200253Billing) GetFreeAccess() InlineResponse200253BillingFreeAccess`

GetFreeAccess returns the FreeAccess field if non-nil, zero value otherwise.

### GetFreeAccessOk

`func (o *InlineResponse200253Billing) GetFreeAccessOk() (*InlineResponse200253BillingFreeAccess, bool)`

GetFreeAccessOk returns a tuple with the FreeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeAccess

`func (o *InlineResponse200253Billing) SetFreeAccess(v InlineResponse200253BillingFreeAccess)`

SetFreeAccess sets FreeAccess field to given value.

### HasFreeAccess

`func (o *InlineResponse200253Billing) HasFreeAccess() bool`

HasFreeAccess returns a boolean if a field has been set.

### GetPrepaidAccessFastLoginEnabled

`func (o *InlineResponse200253Billing) GetPrepaidAccessFastLoginEnabled() bool`

GetPrepaidAccessFastLoginEnabled returns the PrepaidAccessFastLoginEnabled field if non-nil, zero value otherwise.

### GetPrepaidAccessFastLoginEnabledOk

`func (o *InlineResponse200253Billing) GetPrepaidAccessFastLoginEnabledOk() (*bool, bool)`

GetPrepaidAccessFastLoginEnabledOk returns a tuple with the PrepaidAccessFastLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidAccessFastLoginEnabled

`func (o *InlineResponse200253Billing) SetPrepaidAccessFastLoginEnabled(v bool)`

SetPrepaidAccessFastLoginEnabled sets PrepaidAccessFastLoginEnabled field to given value.

### HasPrepaidAccessFastLoginEnabled

`func (o *InlineResponse200253Billing) HasPrepaidAccessFastLoginEnabled() bool`

HasPrepaidAccessFastLoginEnabled returns a boolean if a field has been set.

### GetReplyToEmailAddress

`func (o *InlineResponse200253Billing) GetReplyToEmailAddress() string`

GetReplyToEmailAddress returns the ReplyToEmailAddress field if non-nil, zero value otherwise.

### GetReplyToEmailAddressOk

`func (o *InlineResponse200253Billing) GetReplyToEmailAddressOk() (*string, bool)`

GetReplyToEmailAddressOk returns a tuple with the ReplyToEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmailAddress

`func (o *InlineResponse200253Billing) SetReplyToEmailAddress(v string)`

SetReplyToEmailAddress sets ReplyToEmailAddress field to given value.

### HasReplyToEmailAddress

`func (o *InlineResponse200253Billing) HasReplyToEmailAddress() bool`

HasReplyToEmailAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


