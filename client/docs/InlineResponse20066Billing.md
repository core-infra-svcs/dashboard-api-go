# InlineResponse20066Billing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeAccess** | Pointer to [**InlineResponse20066BillingFreeAccess**](InlineResponse20066BillingFreeAccess.md) |  | [optional] 
**PrepaidAccessFastLoginEnabled** | Pointer to **bool** | Whether or not billing uses the fast login prepaid access option. | [optional] 
**ReplyToEmailAddress** | Pointer to **string** | The email address that reeceives replies from clients | [optional] 

## Methods

### NewInlineResponse20066Billing

`func NewInlineResponse20066Billing() *InlineResponse20066Billing`

NewInlineResponse20066Billing instantiates a new InlineResponse20066Billing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20066BillingWithDefaults

`func NewInlineResponse20066BillingWithDefaults() *InlineResponse20066Billing`

NewInlineResponse20066BillingWithDefaults instantiates a new InlineResponse20066Billing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeAccess

`func (o *InlineResponse20066Billing) GetFreeAccess() InlineResponse20066BillingFreeAccess`

GetFreeAccess returns the FreeAccess field if non-nil, zero value otherwise.

### GetFreeAccessOk

`func (o *InlineResponse20066Billing) GetFreeAccessOk() (*InlineResponse20066BillingFreeAccess, bool)`

GetFreeAccessOk returns a tuple with the FreeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeAccess

`func (o *InlineResponse20066Billing) SetFreeAccess(v InlineResponse20066BillingFreeAccess)`

SetFreeAccess sets FreeAccess field to given value.

### HasFreeAccess

`func (o *InlineResponse20066Billing) HasFreeAccess() bool`

HasFreeAccess returns a boolean if a field has been set.

### GetPrepaidAccessFastLoginEnabled

`func (o *InlineResponse20066Billing) GetPrepaidAccessFastLoginEnabled() bool`

GetPrepaidAccessFastLoginEnabled returns the PrepaidAccessFastLoginEnabled field if non-nil, zero value otherwise.

### GetPrepaidAccessFastLoginEnabledOk

`func (o *InlineResponse20066Billing) GetPrepaidAccessFastLoginEnabledOk() (*bool, bool)`

GetPrepaidAccessFastLoginEnabledOk returns a tuple with the PrepaidAccessFastLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidAccessFastLoginEnabled

`func (o *InlineResponse20066Billing) SetPrepaidAccessFastLoginEnabled(v bool)`

SetPrepaidAccessFastLoginEnabled sets PrepaidAccessFastLoginEnabled field to given value.

### HasPrepaidAccessFastLoginEnabled

`func (o *InlineResponse20066Billing) HasPrepaidAccessFastLoginEnabled() bool`

HasPrepaidAccessFastLoginEnabled returns a boolean if a field has been set.

### GetReplyToEmailAddress

`func (o *InlineResponse20066Billing) GetReplyToEmailAddress() string`

GetReplyToEmailAddress returns the ReplyToEmailAddress field if non-nil, zero value otherwise.

### GetReplyToEmailAddressOk

`func (o *InlineResponse20066Billing) GetReplyToEmailAddressOk() (*string, bool)`

GetReplyToEmailAddressOk returns a tuple with the ReplyToEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmailAddress

`func (o *InlineResponse20066Billing) SetReplyToEmailAddress(v string)`

SetReplyToEmailAddress sets ReplyToEmailAddress field to given value.

### HasReplyToEmailAddress

`func (o *InlineResponse20066Billing) HasReplyToEmailAddress() bool`

HasReplyToEmailAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


