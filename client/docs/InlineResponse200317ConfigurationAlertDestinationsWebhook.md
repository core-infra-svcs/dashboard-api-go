# InlineResponse200317ConfigurationAlertDestinationsWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not the webhook channel is enabled for this alert type | [optional] 
**Recipients** | Pointer to [**[]InlineResponse200317ConfigurationAlertDestinationsWebhookRecipients**](InlineResponse200317ConfigurationAlertDestinationsWebhookRecipients.md) | List of webhook recipients | [optional] 

## Methods

### NewInlineResponse200317ConfigurationAlertDestinationsWebhook

`func NewInlineResponse200317ConfigurationAlertDestinationsWebhook() *InlineResponse200317ConfigurationAlertDestinationsWebhook`

NewInlineResponse200317ConfigurationAlertDestinationsWebhook instantiates a new InlineResponse200317ConfigurationAlertDestinationsWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200317ConfigurationAlertDestinationsWebhookWithDefaults

`func NewInlineResponse200317ConfigurationAlertDestinationsWebhookWithDefaults() *InlineResponse200317ConfigurationAlertDestinationsWebhook`

NewInlineResponse200317ConfigurationAlertDestinationsWebhookWithDefaults instantiates a new InlineResponse200317ConfigurationAlertDestinationsWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200317ConfigurationAlertDestinationsWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200317ConfigurationAlertDestinationsWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200317ConfigurationAlertDestinationsWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200317ConfigurationAlertDestinationsWebhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRecipients

`func (o *InlineResponse200317ConfigurationAlertDestinationsWebhook) GetRecipients() []InlineResponse200317ConfigurationAlertDestinationsWebhookRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineResponse200317ConfigurationAlertDestinationsWebhook) GetRecipientsOk() (*[]InlineResponse200317ConfigurationAlertDestinationsWebhookRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineResponse200317ConfigurationAlertDestinationsWebhook) SetRecipients(v []InlineResponse200317ConfigurationAlertDestinationsWebhookRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InlineResponse200317ConfigurationAlertDestinationsWebhook) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


