# InlineResponse200167

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayloadTemplateId** | Pointer to **string** | Webhook payload template Id | [optional] 
**Type** | Pointer to **string** | The type of the payload template | [optional] 
**Name** | Pointer to **string** | The name of the payload template | [optional] 
**Headers** | Pointer to [**[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders**](NetworksNetworkIdWebhooksPayloadTemplatesHeaders.md) | The payload template headers, will be rendered as a key-value pair in the webhook. | [optional] 
**Body** | Pointer to **string** | The body of the payload template, in liquid template | [optional] 
**Sharing** | Pointer to [**NetworksNetworkIdWebhooksPayloadTemplatesSharing**](NetworksNetworkIdWebhooksPayloadTemplatesSharing.md) |  | [optional] 

## Methods

### NewInlineResponse200167

`func NewInlineResponse200167() *InlineResponse200167`

NewInlineResponse200167 instantiates a new InlineResponse200167 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200167WithDefaults

`func NewInlineResponse200167WithDefaults() *InlineResponse200167`

NewInlineResponse200167WithDefaults instantiates a new InlineResponse200167 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayloadTemplateId

`func (o *InlineResponse200167) GetPayloadTemplateId() string`

GetPayloadTemplateId returns the PayloadTemplateId field if non-nil, zero value otherwise.

### GetPayloadTemplateIdOk

`func (o *InlineResponse200167) GetPayloadTemplateIdOk() (*string, bool)`

GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplateId

`func (o *InlineResponse200167) SetPayloadTemplateId(v string)`

SetPayloadTemplateId sets PayloadTemplateId field to given value.

### HasPayloadTemplateId

`func (o *InlineResponse200167) HasPayloadTemplateId() bool`

HasPayloadTemplateId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200167) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200167) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200167) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200167) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200167) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200167) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200167) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200167) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHeaders

`func (o *InlineResponse200167) GetHeaders() []NetworksNetworkIdWebhooksPayloadTemplatesHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineResponse200167) GetHeadersOk() (*[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineResponse200167) SetHeaders(v []NetworksNetworkIdWebhooksPayloadTemplatesHeaders)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineResponse200167) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *InlineResponse200167) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InlineResponse200167) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InlineResponse200167) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *InlineResponse200167) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSharing

`func (o *InlineResponse200167) GetSharing() NetworksNetworkIdWebhooksPayloadTemplatesSharing`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *InlineResponse200167) GetSharingOk() (*NetworksNetworkIdWebhooksPayloadTemplatesSharing, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *InlineResponse200167) SetSharing(v NetworksNetworkIdWebhooksPayloadTemplatesSharing)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *InlineResponse200167) HasSharing() bool`

HasSharing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


