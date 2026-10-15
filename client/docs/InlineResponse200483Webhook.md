# InlineResponse200483Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The webhook receiver URL where the callback will be sent | [optional] 
**HttpServer** | Pointer to [**InlineResponse200483WebhookHttpServer**](InlineResponse200483WebhookHttpServer.md) |  | [optional] 
**PayloadTemplate** | Pointer to [**InlineResponse200483WebhookPayloadTemplate**](InlineResponse200483WebhookPayloadTemplate.md) |  | [optional] 
**SentAt** | Pointer to **time.Time** | The timestamp the callback was sent to the webhook receiver | [optional] 

## Methods

### NewInlineResponse200483Webhook

`func NewInlineResponse200483Webhook() *InlineResponse200483Webhook`

NewInlineResponse200483Webhook instantiates a new InlineResponse200483Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200483WebhookWithDefaults

`func NewInlineResponse200483WebhookWithDefaults() *InlineResponse200483Webhook`

NewInlineResponse200483WebhookWithDefaults instantiates a new InlineResponse200483Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *InlineResponse200483Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200483Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200483Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200483Webhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHttpServer

`func (o *InlineResponse200483Webhook) GetHttpServer() InlineResponse200483WebhookHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *InlineResponse200483Webhook) GetHttpServerOk() (*InlineResponse200483WebhookHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *InlineResponse200483Webhook) SetHttpServer(v InlineResponse200483WebhookHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *InlineResponse200483Webhook) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *InlineResponse200483Webhook) GetPayloadTemplate() InlineResponse200483WebhookPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *InlineResponse200483Webhook) GetPayloadTemplateOk() (*InlineResponse200483WebhookPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *InlineResponse200483Webhook) SetPayloadTemplate(v InlineResponse200483WebhookPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *InlineResponse200483Webhook) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### GetSentAt

`func (o *InlineResponse200483Webhook) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *InlineResponse200483Webhook) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *InlineResponse200483Webhook) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *InlineResponse200483Webhook) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


