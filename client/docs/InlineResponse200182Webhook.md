# InlineResponse200182Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The webhook receiver URL where the callback will be sent | [optional] 
**HttpServer** | Pointer to [**InlineResponse200182WebhookHttpServer**](InlineResponse200182WebhookHttpServer.md) |  | [optional] 
**PayloadTemplate** | Pointer to [**InlineResponse200182WebhookPayloadTemplate**](InlineResponse200182WebhookPayloadTemplate.md) |  | [optional] 
**SentAt** | Pointer to **time.Time** | The timestamp the callback was sent to the webhook receiver | [optional] 

## Methods

### NewInlineResponse200182Webhook

`func NewInlineResponse200182Webhook() *InlineResponse200182Webhook`

NewInlineResponse200182Webhook instantiates a new InlineResponse200182Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200182WebhookWithDefaults

`func NewInlineResponse200182WebhookWithDefaults() *InlineResponse200182Webhook`

NewInlineResponse200182WebhookWithDefaults instantiates a new InlineResponse200182Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *InlineResponse200182Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200182Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200182Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200182Webhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHttpServer

`func (o *InlineResponse200182Webhook) GetHttpServer() InlineResponse200182WebhookHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *InlineResponse200182Webhook) GetHttpServerOk() (*InlineResponse200182WebhookHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *InlineResponse200182Webhook) SetHttpServer(v InlineResponse200182WebhookHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *InlineResponse200182Webhook) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *InlineResponse200182Webhook) GetPayloadTemplate() InlineResponse200182WebhookPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *InlineResponse200182Webhook) GetPayloadTemplateOk() (*InlineResponse200182WebhookPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *InlineResponse200182Webhook) SetPayloadTemplate(v InlineResponse200182WebhookPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *InlineResponse200182Webhook) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### GetSentAt

`func (o *InlineResponse200182Webhook) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *InlineResponse200182Webhook) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *InlineResponse200182Webhook) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *InlineResponse200182Webhook) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


