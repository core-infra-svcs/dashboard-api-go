# InlineResponse200331

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackId** | Pointer to **string** | The ID of the callback | [optional] 
**Status** | Pointer to **string** | The status of the callback | [optional] 
**Errors** | Pointer to **[]string** | The errors returned by the callback | [optional] 
**CreatedBy** | Pointer to [**InlineResponse200331CreatedBy**](InlineResponse200331CreatedBy.md) |  | [optional] 
**Webhook** | Pointer to [**InlineResponse200331Webhook**](InlineResponse200331Webhook.md) |  | [optional] 

## Methods

### NewInlineResponse200331

`func NewInlineResponse200331() *InlineResponse200331`

NewInlineResponse200331 instantiates a new InlineResponse200331 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200331WithDefaults

`func NewInlineResponse200331WithDefaults() *InlineResponse200331`

NewInlineResponse200331WithDefaults instantiates a new InlineResponse200331 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackId

`func (o *InlineResponse200331) GetCallbackId() string`

GetCallbackId returns the CallbackId field if non-nil, zero value otherwise.

### GetCallbackIdOk

`func (o *InlineResponse200331) GetCallbackIdOk() (*string, bool)`

GetCallbackIdOk returns a tuple with the CallbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackId

`func (o *InlineResponse200331) SetCallbackId(v string)`

SetCallbackId sets CallbackId field to given value.

### HasCallbackId

`func (o *InlineResponse200331) HasCallbackId() bool`

HasCallbackId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200331) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200331) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200331) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200331) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200331) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200331) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200331) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200331) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetCreatedBy

`func (o *InlineResponse200331) GetCreatedBy() InlineResponse200331CreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InlineResponse200331) GetCreatedByOk() (*InlineResponse200331CreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InlineResponse200331) SetCreatedBy(v InlineResponse200331CreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InlineResponse200331) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetWebhook

`func (o *InlineResponse200331) GetWebhook() InlineResponse200331Webhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *InlineResponse200331) GetWebhookOk() (*InlineResponse200331Webhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *InlineResponse200331) SetWebhook(v InlineResponse200331Webhook)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *InlineResponse200331) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


