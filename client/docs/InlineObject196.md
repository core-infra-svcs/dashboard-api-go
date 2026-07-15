# InlineObject196

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new template | 
**Body** | Pointer to **string** | The liquid template used for the body of the webhook message. Either &#x60;body&#x60; or &#x60;bodyFile&#x60; must be specified. | [optional] 
**Headers** | Pointer to [**[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders1**](NetworksNetworkIdWebhooksPayloadTemplatesHeaders1.md) | The liquid template used with the webhook headers. | [optional] 
**BodyFile** | Pointer to **string** | A Base64 encoded file containing liquid template used for the body of the webhook message. Either &#x60;body&#x60; or &#x60;bodyFile&#x60; must be specified. | [optional] 
**HeadersFile** | Pointer to **string** | A Base64 encoded file containing the liquid template used with the webhook headers. | [optional] 

## Methods

### NewInlineObject196

`func NewInlineObject196(name string, ) *InlineObject196`

NewInlineObject196 instantiates a new InlineObject196 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject196WithDefaults

`func NewInlineObject196WithDefaults() *InlineObject196`

NewInlineObject196WithDefaults instantiates a new InlineObject196 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject196) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject196) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject196) SetName(v string)`

SetName sets Name field to given value.


### GetBody

`func (o *InlineObject196) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InlineObject196) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InlineObject196) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *InlineObject196) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *InlineObject196) GetHeaders() []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineObject196) GetHeadersOk() (*[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders1, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineObject196) SetHeaders(v []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineObject196) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *InlineObject196) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *InlineObject196) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetBodyFile

`func (o *InlineObject196) GetBodyFile() string`

GetBodyFile returns the BodyFile field if non-nil, zero value otherwise.

### GetBodyFileOk

`func (o *InlineObject196) GetBodyFileOk() (*string, bool)`

GetBodyFileOk returns a tuple with the BodyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyFile

`func (o *InlineObject196) SetBodyFile(v string)`

SetBodyFile sets BodyFile field to given value.

### HasBodyFile

`func (o *InlineObject196) HasBodyFile() bool`

HasBodyFile returns a boolean if a field has been set.

### GetHeadersFile

`func (o *InlineObject196) GetHeadersFile() string`

GetHeadersFile returns the HeadersFile field if non-nil, zero value otherwise.

### GetHeadersFileOk

`func (o *InlineObject196) GetHeadersFileOk() (*string, bool)`

GetHeadersFileOk returns a tuple with the HeadersFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersFile

`func (o *InlineObject196) SetHeadersFile(v string)`

SetHeadersFile sets HeadersFile field to given value.

### HasHeadersFile

`func (o *InlineObject196) HasHeadersFile() bool`

HasHeadersFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


