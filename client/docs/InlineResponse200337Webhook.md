# InlineResponse200337Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The webhook receiver URL where the callback will be sent | [optional] 
**HttpServer** | Pointer to [**InlineResponse200337WebhookHttpServer**](InlineResponse200337WebhookHttpServer.md) |  | [optional] 
**PayloadTemplate** | Pointer to [**InlineResponse200337WebhookPayloadTemplate**](InlineResponse200337WebhookPayloadTemplate.md) |  | [optional] 
**SentAt** | Pointer to **time.Time** | The timestamp the callback was sent to the webhook receiver | [optional] 

## Methods

### NewInlineResponse200337Webhook

`func NewInlineResponse200337Webhook() *InlineResponse200337Webhook`

NewInlineResponse200337Webhook instantiates a new InlineResponse200337Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200337WebhookWithDefaults

`func NewInlineResponse200337WebhookWithDefaults() *InlineResponse200337Webhook`

NewInlineResponse200337WebhookWithDefaults instantiates a new InlineResponse200337Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *InlineResponse200337Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200337Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200337Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200337Webhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHttpServer

`func (o *InlineResponse200337Webhook) GetHttpServer() InlineResponse200337WebhookHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *InlineResponse200337Webhook) GetHttpServerOk() (*InlineResponse200337WebhookHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *InlineResponse200337Webhook) SetHttpServer(v InlineResponse200337WebhookHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *InlineResponse200337Webhook) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *InlineResponse200337Webhook) GetPayloadTemplate() InlineResponse200337WebhookPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *InlineResponse200337Webhook) GetPayloadTemplateOk() (*InlineResponse200337WebhookPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *InlineResponse200337Webhook) SetPayloadTemplate(v InlineResponse200337WebhookPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *InlineResponse200337Webhook) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### GetSentAt

`func (o *InlineResponse200337Webhook) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *InlineResponse200337Webhook) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *InlineResponse200337Webhook) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *InlineResponse200337Webhook) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


