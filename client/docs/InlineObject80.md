# InlineObject80

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigTemplateId** | **string** | The ID of the template to which the network should be bound. | 
**AutoBind** | Pointer to **bool** | Optional boolean indicating whether the network&#39;s switches should automatically bind to profiles of the same model. Defaults to false if left unspecified. This option only affects switch networks and switch templates. Auto-bind is not valid unless the switch template has at least one profile and has at most one profile per switch model. | [optional] 

## Methods

### NewInlineObject80

`func NewInlineObject80(configTemplateId string, ) *InlineObject80`

NewInlineObject80 instantiates a new InlineObject80 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject80WithDefaults

`func NewInlineObject80WithDefaults() *InlineObject80`

NewInlineObject80WithDefaults instantiates a new InlineObject80 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigTemplateId

`func (o *InlineObject80) GetConfigTemplateId() string`

GetConfigTemplateId returns the ConfigTemplateId field if non-nil, zero value otherwise.

### GetConfigTemplateIdOk

`func (o *InlineObject80) GetConfigTemplateIdOk() (*string, bool)`

GetConfigTemplateIdOk returns a tuple with the ConfigTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplateId

`func (o *InlineObject80) SetConfigTemplateId(v string)`

SetConfigTemplateId sets ConfigTemplateId field to given value.


### GetAutoBind

`func (o *InlineObject80) GetAutoBind() bool`

GetAutoBind returns the AutoBind field if non-nil, zero value otherwise.

### GetAutoBindOk

`func (o *InlineObject80) GetAutoBindOk() (*bool, bool)`

GetAutoBindOk returns a tuple with the AutoBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBind

`func (o *InlineObject80) SetAutoBind(v bool)`

SetAutoBind sets AutoBind field to given value.

### HasAutoBind

`func (o *InlineObject80) HasAutoBind() bool`

HasAutoBind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


