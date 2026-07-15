# InlineObject197

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the template | [optional] 
**Body** | Pointer to **string** | The liquid template used for the body of the webhook message. | [optional] 
**Headers** | Pointer to [**[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders1**](NetworksNetworkIdWebhooksPayloadTemplatesHeaders1.md) | The liquid template used with the webhook headers. | [optional] 
**BodyFile** | Pointer to **string** | A file containing liquid template used for the body of the webhook message. | [optional] 
**HeadersFile** | Pointer to **string** | A file containing the liquid template used with the webhook headers. | [optional] 

## Methods

### NewInlineObject197

`func NewInlineObject197() *InlineObject197`

NewInlineObject197 instantiates a new InlineObject197 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject197WithDefaults

`func NewInlineObject197WithDefaults() *InlineObject197`

NewInlineObject197WithDefaults instantiates a new InlineObject197 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject197) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject197) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject197) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject197) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBody

`func (o *InlineObject197) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InlineObject197) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InlineObject197) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *InlineObject197) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *InlineObject197) GetHeaders() []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineObject197) GetHeadersOk() (*[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders1, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineObject197) SetHeaders(v []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineObject197) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *InlineObject197) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *InlineObject197) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetBodyFile

`func (o *InlineObject197) GetBodyFile() string`

GetBodyFile returns the BodyFile field if non-nil, zero value otherwise.

### GetBodyFileOk

`func (o *InlineObject197) GetBodyFileOk() (*string, bool)`

GetBodyFileOk returns a tuple with the BodyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyFile

`func (o *InlineObject197) SetBodyFile(v string)`

SetBodyFile sets BodyFile field to given value.

### HasBodyFile

`func (o *InlineObject197) HasBodyFile() bool`

HasBodyFile returns a boolean if a field has been set.

### GetHeadersFile

`func (o *InlineObject197) GetHeadersFile() string`

GetHeadersFile returns the HeadersFile field if non-nil, zero value otherwise.

### GetHeadersFileOk

`func (o *InlineObject197) GetHeadersFileOk() (*string, bool)`

GetHeadersFileOk returns a tuple with the HeadersFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersFile

`func (o *InlineObject197) SetHeadersFile(v string)`

SetHeadersFile sets HeadersFile field to given value.

### HasHeadersFile

`func (o *InlineObject197) HasHeadersFile() bool`

HasHeadersFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


