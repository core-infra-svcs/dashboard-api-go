# InlineResponse200179

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A Base64 encoded ID. | [optional] 
**Name** | Pointer to **string** | A name for easy reference to the HTTP server | [optional] 
**Url** | Pointer to **string** | The URL of the HTTP server. | [optional] 
**NetworkId** | Pointer to **string** | A Meraki network ID. | [optional] 
**PayloadTemplate** | Pointer to [**NetworksNetworkIdWebhooksHttpServersPayloadTemplate**](NetworksNetworkIdWebhooksHttpServersPayloadTemplate.md) |  | [optional] 

## Methods

### NewInlineResponse200179

`func NewInlineResponse200179() *InlineResponse200179`

NewInlineResponse200179 instantiates a new InlineResponse200179 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200179WithDefaults

`func NewInlineResponse200179WithDefaults() *InlineResponse200179`

NewInlineResponse200179WithDefaults instantiates a new InlineResponse200179 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200179) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200179) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200179) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200179) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200179) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200179) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200179) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200179) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200179) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200179) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200179) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200179) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200179) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200179) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200179) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200179) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *InlineResponse200179) GetPayloadTemplate() NetworksNetworkIdWebhooksHttpServersPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *InlineResponse200179) GetPayloadTemplateOk() (*NetworksNetworkIdWebhooksHttpServersPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *InlineResponse200179) SetPayloadTemplate(v NetworksNetworkIdWebhooksHttpServersPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *InlineResponse200179) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


