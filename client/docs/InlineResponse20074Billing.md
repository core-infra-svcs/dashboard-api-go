# InlineResponse20074Billing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeAccess** | Pointer to [**InlineResponse20074BillingFreeAccess**](InlineResponse20074BillingFreeAccess.md) |  | [optional] 
**PrepaidAccessFastLoginEnabled** | Pointer to **bool** | Whether or not billing uses the fast login prepaid access option. | [optional] 
**ReplyToEmailAddress** | Pointer to **string** | The email address that reeceives replies from clients | [optional] 

## Methods

### NewInlineResponse20074Billing

`func NewInlineResponse20074Billing() *InlineResponse20074Billing`

NewInlineResponse20074Billing instantiates a new InlineResponse20074Billing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20074BillingWithDefaults

`func NewInlineResponse20074BillingWithDefaults() *InlineResponse20074Billing`

NewInlineResponse20074BillingWithDefaults instantiates a new InlineResponse20074Billing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeAccess

`func (o *InlineResponse20074Billing) GetFreeAccess() InlineResponse20074BillingFreeAccess`

GetFreeAccess returns the FreeAccess field if non-nil, zero value otherwise.

### GetFreeAccessOk

`func (o *InlineResponse20074Billing) GetFreeAccessOk() (*InlineResponse20074BillingFreeAccess, bool)`

GetFreeAccessOk returns a tuple with the FreeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeAccess

`func (o *InlineResponse20074Billing) SetFreeAccess(v InlineResponse20074BillingFreeAccess)`

SetFreeAccess sets FreeAccess field to given value.

### HasFreeAccess

`func (o *InlineResponse20074Billing) HasFreeAccess() bool`

HasFreeAccess returns a boolean if a field has been set.

### GetPrepaidAccessFastLoginEnabled

`func (o *InlineResponse20074Billing) GetPrepaidAccessFastLoginEnabled() bool`

GetPrepaidAccessFastLoginEnabled returns the PrepaidAccessFastLoginEnabled field if non-nil, zero value otherwise.

### GetPrepaidAccessFastLoginEnabledOk

`func (o *InlineResponse20074Billing) GetPrepaidAccessFastLoginEnabledOk() (*bool, bool)`

GetPrepaidAccessFastLoginEnabledOk returns a tuple with the PrepaidAccessFastLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidAccessFastLoginEnabled

`func (o *InlineResponse20074Billing) SetPrepaidAccessFastLoginEnabled(v bool)`

SetPrepaidAccessFastLoginEnabled sets PrepaidAccessFastLoginEnabled field to given value.

### HasPrepaidAccessFastLoginEnabled

`func (o *InlineResponse20074Billing) HasPrepaidAccessFastLoginEnabled() bool`

HasPrepaidAccessFastLoginEnabled returns a boolean if a field has been set.

### GetReplyToEmailAddress

`func (o *InlineResponse20074Billing) GetReplyToEmailAddress() string`

GetReplyToEmailAddress returns the ReplyToEmailAddress field if non-nil, zero value otherwise.

### GetReplyToEmailAddressOk

`func (o *InlineResponse20074Billing) GetReplyToEmailAddressOk() (*string, bool)`

GetReplyToEmailAddressOk returns a tuple with the ReplyToEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmailAddress

`func (o *InlineResponse20074Billing) SetReplyToEmailAddress(v string)`

SetReplyToEmailAddress sets ReplyToEmailAddress field to given value.

### HasReplyToEmailAddress

`func (o *InlineResponse20074Billing) HasReplyToEmailAddress() bool`

HasReplyToEmailAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


