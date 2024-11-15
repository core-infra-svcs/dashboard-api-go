# InlineObject170

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for easy reference to the HTTP server | [optional] 
**SharedSecret** | Pointer to **string** | A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki. | [optional] 
**PayloadTemplate** | Pointer to [**NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate**](NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate.md) |  | [optional] 

## Methods

### NewInlineObject170

`func NewInlineObject170() *InlineObject170`

NewInlineObject170 instantiates a new InlineObject170 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject170WithDefaults

`func NewInlineObject170WithDefaults() *InlineObject170`

NewInlineObject170WithDefaults instantiates a new InlineObject170 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject170) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject170) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject170) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject170) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSharedSecret

`func (o *InlineObject170) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *InlineObject170) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *InlineObject170) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *InlineObject170) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *InlineObject170) GetPayloadTemplate() NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *InlineObject170) GetPayloadTemplateOk() (*NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *InlineObject170) SetPayloadTemplate(v NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *InlineObject170) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


