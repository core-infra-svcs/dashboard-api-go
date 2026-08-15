# InlineResponse200302ConfigurationAlertDestinationsWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not the webhook channel is enabled for this alert type | [optional] 
**Recipients** | Pointer to [**[]InlineResponse200302ConfigurationAlertDestinationsWebhookRecipients**](InlineResponse200302ConfigurationAlertDestinationsWebhookRecipients.md) | List of webhook recipients | [optional] 

## Methods

### NewInlineResponse200302ConfigurationAlertDestinationsWebhook

`func NewInlineResponse200302ConfigurationAlertDestinationsWebhook() *InlineResponse200302ConfigurationAlertDestinationsWebhook`

NewInlineResponse200302ConfigurationAlertDestinationsWebhook instantiates a new InlineResponse200302ConfigurationAlertDestinationsWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200302ConfigurationAlertDestinationsWebhookWithDefaults

`func NewInlineResponse200302ConfigurationAlertDestinationsWebhookWithDefaults() *InlineResponse200302ConfigurationAlertDestinationsWebhook`

NewInlineResponse200302ConfigurationAlertDestinationsWebhookWithDefaults instantiates a new InlineResponse200302ConfigurationAlertDestinationsWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200302ConfigurationAlertDestinationsWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200302ConfigurationAlertDestinationsWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200302ConfigurationAlertDestinationsWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200302ConfigurationAlertDestinationsWebhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRecipients

`func (o *InlineResponse200302ConfigurationAlertDestinationsWebhook) GetRecipients() []InlineResponse200302ConfigurationAlertDestinationsWebhookRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineResponse200302ConfigurationAlertDestinationsWebhook) GetRecipientsOk() (*[]InlineResponse200302ConfigurationAlertDestinationsWebhookRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineResponse200302ConfigurationAlertDestinationsWebhook) SetRecipients(v []InlineResponse200302ConfigurationAlertDestinationsWebhookRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InlineResponse200302ConfigurationAlertDestinationsWebhook) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


