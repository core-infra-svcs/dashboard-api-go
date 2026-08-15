# InlineResponse200302ConfigurationAlertDestinationsSms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not the sms channel is enabled for this alert type | [optional] 
**Recipients** | Pointer to **[]string** | List of sms recipient phone numbers to receive the notifications | [optional] 

## Methods

### NewInlineResponse200302ConfigurationAlertDestinationsSms

`func NewInlineResponse200302ConfigurationAlertDestinationsSms() *InlineResponse200302ConfigurationAlertDestinationsSms`

NewInlineResponse200302ConfigurationAlertDestinationsSms instantiates a new InlineResponse200302ConfigurationAlertDestinationsSms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200302ConfigurationAlertDestinationsSmsWithDefaults

`func NewInlineResponse200302ConfigurationAlertDestinationsSmsWithDefaults() *InlineResponse200302ConfigurationAlertDestinationsSms`

NewInlineResponse200302ConfigurationAlertDestinationsSmsWithDefaults instantiates a new InlineResponse200302ConfigurationAlertDestinationsSms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200302ConfigurationAlertDestinationsSms) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200302ConfigurationAlertDestinationsSms) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200302ConfigurationAlertDestinationsSms) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200302ConfigurationAlertDestinationsSms) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRecipients

`func (o *InlineResponse200302ConfigurationAlertDestinationsSms) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineResponse200302ConfigurationAlertDestinationsSms) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineResponse200302ConfigurationAlertDestinationsSms) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InlineResponse200302ConfigurationAlertDestinationsSms) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


