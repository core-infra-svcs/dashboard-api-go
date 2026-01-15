# InlineResponse200361

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackId** | Pointer to **string** | The ID of the callback | [optional] 
**Status** | Pointer to **string** | The status of the callback | [optional] 
**Errors** | Pointer to **[]string** | The errors returned by the callback | [optional] 
**CreatedBy** | Pointer to [**InlineResponse200361CreatedBy**](InlineResponse200361CreatedBy.md) |  | [optional] 
**Webhook** | Pointer to [**InlineResponse200361Webhook**](InlineResponse200361Webhook.md) |  | [optional] 

## Methods

### NewInlineResponse200361

`func NewInlineResponse200361() *InlineResponse200361`

NewInlineResponse200361 instantiates a new InlineResponse200361 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200361WithDefaults

`func NewInlineResponse200361WithDefaults() *InlineResponse200361`

NewInlineResponse200361WithDefaults instantiates a new InlineResponse200361 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackId

`func (o *InlineResponse200361) GetCallbackId() string`

GetCallbackId returns the CallbackId field if non-nil, zero value otherwise.

### GetCallbackIdOk

`func (o *InlineResponse200361) GetCallbackIdOk() (*string, bool)`

GetCallbackIdOk returns a tuple with the CallbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackId

`func (o *InlineResponse200361) SetCallbackId(v string)`

SetCallbackId sets CallbackId field to given value.

### HasCallbackId

`func (o *InlineResponse200361) HasCallbackId() bool`

HasCallbackId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200361) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200361) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200361) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200361) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200361) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200361) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200361) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200361) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetCreatedBy

`func (o *InlineResponse200361) GetCreatedBy() InlineResponse200361CreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InlineResponse200361) GetCreatedByOk() (*InlineResponse200361CreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InlineResponse200361) SetCreatedBy(v InlineResponse200361CreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InlineResponse200361) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetWebhook

`func (o *InlineResponse200361) GetWebhook() InlineResponse200361Webhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *InlineResponse200361) GetWebhookOk() (*InlineResponse200361Webhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *InlineResponse200361) SetWebhook(v InlineResponse200361Webhook)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *InlineResponse200361) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


