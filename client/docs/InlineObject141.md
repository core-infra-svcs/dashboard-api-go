# InlineObject141

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for easy reference to the HTTP server | 
**Url** | **string** | The URL of the HTTP server. Once set, cannot be updated. | 
**SharedSecret** | Pointer to **string** | A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki. | [optional] 
**PayloadTemplate** | Pointer to [**NetworksNetworkIdWebhooksHttpServersPayloadTemplate1**](NetworksNetworkIdWebhooksHttpServersPayloadTemplate1.md) |  | [optional] 

## Methods

### NewInlineObject141

`func NewInlineObject141(name string, url string, ) *InlineObject141`

NewInlineObject141 instantiates a new InlineObject141 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject141WithDefaults

`func NewInlineObject141WithDefaults() *InlineObject141`

NewInlineObject141WithDefaults instantiates a new InlineObject141 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject141) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject141) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject141) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *InlineObject141) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineObject141) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineObject141) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSharedSecret

`func (o *InlineObject141) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *InlineObject141) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *InlineObject141) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *InlineObject141) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *InlineObject141) GetPayloadTemplate() NetworksNetworkIdWebhooksHttpServersPayloadTemplate1`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *InlineObject141) GetPayloadTemplateOk() (*NetworksNetworkIdWebhooksHttpServersPayloadTemplate1, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *InlineObject141) SetPayloadTemplate(v NetworksNetworkIdWebhooksHttpServersPayloadTemplate1)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *InlineObject141) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


