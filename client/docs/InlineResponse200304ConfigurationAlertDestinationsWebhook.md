# InlineResponse200304ConfigurationAlertDestinationsWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not the webhook channel is enabled for this alert type | [optional] 
**Recipients** | Pointer to [**[]InlineResponse200304ConfigurationAlertDestinationsWebhookRecipients**](InlineResponse200304ConfigurationAlertDestinationsWebhookRecipients.md) | List of webhook recipients | [optional] 

## Methods

### NewInlineResponse200304ConfigurationAlertDestinationsWebhook

`func NewInlineResponse200304ConfigurationAlertDestinationsWebhook() *InlineResponse200304ConfigurationAlertDestinationsWebhook`

NewInlineResponse200304ConfigurationAlertDestinationsWebhook instantiates a new InlineResponse200304ConfigurationAlertDestinationsWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200304ConfigurationAlertDestinationsWebhookWithDefaults

`func NewInlineResponse200304ConfigurationAlertDestinationsWebhookWithDefaults() *InlineResponse200304ConfigurationAlertDestinationsWebhook`

NewInlineResponse200304ConfigurationAlertDestinationsWebhookWithDefaults instantiates a new InlineResponse200304ConfigurationAlertDestinationsWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200304ConfigurationAlertDestinationsWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200304ConfigurationAlertDestinationsWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200304ConfigurationAlertDestinationsWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200304ConfigurationAlertDestinationsWebhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRecipients

`func (o *InlineResponse200304ConfigurationAlertDestinationsWebhook) GetRecipients() []InlineResponse200304ConfigurationAlertDestinationsWebhookRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineResponse200304ConfigurationAlertDestinationsWebhook) GetRecipientsOk() (*[]InlineResponse200304ConfigurationAlertDestinationsWebhookRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineResponse200304ConfigurationAlertDestinationsWebhook) SetRecipients(v []InlineResponse200304ConfigurationAlertDestinationsWebhookRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InlineResponse200304ConfigurationAlertDestinationsWebhook) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


