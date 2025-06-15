# InlineResponse200343Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The webhook receiver URL where the callback will be sent | [optional] 
**HttpServer** | Pointer to [**InlineResponse200343WebhookHttpServer**](InlineResponse200343WebhookHttpServer.md) |  | [optional] 
**PayloadTemplate** | Pointer to [**InlineResponse200343WebhookPayloadTemplate**](InlineResponse200343WebhookPayloadTemplate.md) |  | [optional] 
**SentAt** | Pointer to **time.Time** | The timestamp the callback was sent to the webhook receiver | [optional] 

## Methods

### NewInlineResponse200343Webhook

`func NewInlineResponse200343Webhook() *InlineResponse200343Webhook`

NewInlineResponse200343Webhook instantiates a new InlineResponse200343Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200343WebhookWithDefaults

`func NewInlineResponse200343WebhookWithDefaults() *InlineResponse200343Webhook`

NewInlineResponse200343WebhookWithDefaults instantiates a new InlineResponse200343Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *InlineResponse200343Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200343Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200343Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200343Webhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHttpServer

`func (o *InlineResponse200343Webhook) GetHttpServer() InlineResponse200343WebhookHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *InlineResponse200343Webhook) GetHttpServerOk() (*InlineResponse200343WebhookHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *InlineResponse200343Webhook) SetHttpServer(v InlineResponse200343WebhookHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *InlineResponse200343Webhook) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *InlineResponse200343Webhook) GetPayloadTemplate() InlineResponse200343WebhookPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *InlineResponse200343Webhook) GetPayloadTemplateOk() (*InlineResponse200343WebhookPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *InlineResponse200343Webhook) SetPayloadTemplate(v InlineResponse200343WebhookPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *InlineResponse200343Webhook) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### GetSentAt

`func (o *InlineResponse200343Webhook) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *InlineResponse200343Webhook) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *InlineResponse200343Webhook) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *InlineResponse200343Webhook) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


