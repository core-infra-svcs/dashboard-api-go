# InlineObject175

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL where the test webhook will be sent | 
**SharedSecret** | Pointer to **string** | The shared secret the test webhook will send. Optional. Defaults to an empty string. | [optional] [default to ""]
**PayloadTemplateId** | Pointer to **string** | The ID of the payload template of the test webhook. Defaults to the HTTP server&#39;s template ID if one exists for the given URL, or Generic template ID otherwise | [optional] 
**PayloadTemplateName** | Pointer to **string** | The name of the payload template. | [optional] 
**AlertTypeId** | Pointer to **string** | The type of alert which the test webhook will send. Optional. Defaults to power_supply_down. | [optional] [default to "power_supply_down"]

## Methods

### NewInlineObject175

`func NewInlineObject175(url string, ) *InlineObject175`

NewInlineObject175 instantiates a new InlineObject175 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject175WithDefaults

`func NewInlineObject175WithDefaults() *InlineObject175`

NewInlineObject175WithDefaults instantiates a new InlineObject175 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *InlineObject175) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineObject175) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineObject175) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSharedSecret

`func (o *InlineObject175) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *InlineObject175) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *InlineObject175) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *InlineObject175) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetPayloadTemplateId

`func (o *InlineObject175) GetPayloadTemplateId() string`

GetPayloadTemplateId returns the PayloadTemplateId field if non-nil, zero value otherwise.

### GetPayloadTemplateIdOk

`func (o *InlineObject175) GetPayloadTemplateIdOk() (*string, bool)`

GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplateId

`func (o *InlineObject175) SetPayloadTemplateId(v string)`

SetPayloadTemplateId sets PayloadTemplateId field to given value.

### HasPayloadTemplateId

`func (o *InlineObject175) HasPayloadTemplateId() bool`

HasPayloadTemplateId returns a boolean if a field has been set.

### GetPayloadTemplateName

`func (o *InlineObject175) GetPayloadTemplateName() string`

GetPayloadTemplateName returns the PayloadTemplateName field if non-nil, zero value otherwise.

### GetPayloadTemplateNameOk

`func (o *InlineObject175) GetPayloadTemplateNameOk() (*string, bool)`

GetPayloadTemplateNameOk returns a tuple with the PayloadTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplateName

`func (o *InlineObject175) SetPayloadTemplateName(v string)`

SetPayloadTemplateName sets PayloadTemplateName field to given value.

### HasPayloadTemplateName

`func (o *InlineObject175) HasPayloadTemplateName() bool`

HasPayloadTemplateName returns a boolean if a field has been set.

### GetAlertTypeId

`func (o *InlineObject175) GetAlertTypeId() string`

GetAlertTypeId returns the AlertTypeId field if non-nil, zero value otherwise.

### GetAlertTypeIdOk

`func (o *InlineObject175) GetAlertTypeIdOk() (*string, bool)`

GetAlertTypeIdOk returns a tuple with the AlertTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypeId

`func (o *InlineObject175) SetAlertTypeId(v string)`

SetAlertTypeId sets AlertTypeId field to given value.

### HasAlertTypeId

`func (o *InlineObject175) HasAlertTypeId() bool`

HasAlertTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


