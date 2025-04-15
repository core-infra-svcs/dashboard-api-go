# InlineResponse200331Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The webhook receiver URL where the callback will be sent | [optional] 
**HttpServer** | Pointer to [**InlineResponse200331WebhookHttpServer**](InlineResponse200331WebhookHttpServer.md) |  | [optional] 
**PayloadTemplate** | Pointer to [**InlineResponse200331WebhookPayloadTemplate**](InlineResponse200331WebhookPayloadTemplate.md) |  | [optional] 
**SentAt** | Pointer to **time.Time** | The timestamp the callback was sent to the webhook receiver | [optional] 

## Methods

### NewInlineResponse200331Webhook

`func NewInlineResponse200331Webhook() *InlineResponse200331Webhook`

NewInlineResponse200331Webhook instantiates a new InlineResponse200331Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200331WebhookWithDefaults

`func NewInlineResponse200331WebhookWithDefaults() *InlineResponse200331Webhook`

NewInlineResponse200331WebhookWithDefaults instantiates a new InlineResponse200331Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *InlineResponse200331Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200331Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200331Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200331Webhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHttpServer

`func (o *InlineResponse200331Webhook) GetHttpServer() InlineResponse200331WebhookHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *InlineResponse200331Webhook) GetHttpServerOk() (*InlineResponse200331WebhookHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *InlineResponse200331Webhook) SetHttpServer(v InlineResponse200331WebhookHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *InlineResponse200331Webhook) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *InlineResponse200331Webhook) GetPayloadTemplate() InlineResponse200331WebhookPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *InlineResponse200331Webhook) GetPayloadTemplateOk() (*InlineResponse200331WebhookPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *InlineResponse200331Webhook) SetPayloadTemplate(v InlineResponse200331WebhookPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *InlineResponse200331Webhook) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### GetSentAt

`func (o *InlineResponse200331Webhook) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *InlineResponse200331Webhook) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *InlineResponse200331Webhook) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *InlineResponse200331Webhook) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


