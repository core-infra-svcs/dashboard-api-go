# InlineObject267

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new network | 
**ProductTypes** | **[]string** | The product type(s) of the new network. If more than one type is included, the network will be a combined network. | 
**Tags** | Pointer to **[]string** | A list of tags to be applied to the network | [optional] 
**TimeZone** | Pointer to **string** | The timezone of the network. For a list of allowed timezones, please see the &#39;TZ&#39; column in the table in &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones&#39;&gt;this article.&lt;/a&gt; | [optional] 
**CopyFromNetworkId** | Pointer to **string** | The ID of the network to copy configuration from. Other provided parameters will override the copied configuration, except type which must match this network&#39;s type exactly. | [optional] 
**Notes** | Pointer to **string** | Add any notes or additional information about this network here. | [optional] 

## Methods

### NewInlineObject267

`func NewInlineObject267(name string, productTypes []string, ) *InlineObject267`

NewInlineObject267 instantiates a new InlineObject267 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject267WithDefaults

`func NewInlineObject267WithDefaults() *InlineObject267`

NewInlineObject267WithDefaults instantiates a new InlineObject267 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject267) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject267) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject267) SetName(v string)`

SetName sets Name field to given value.


### GetProductTypes

`func (o *InlineObject267) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *InlineObject267) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *InlineObject267) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.


### GetTags

`func (o *InlineObject267) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject267) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject267) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject267) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeZone

`func (o *InlineObject267) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *InlineObject267) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *InlineObject267) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *InlineObject267) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetCopyFromNetworkId

`func (o *InlineObject267) GetCopyFromNetworkId() string`

GetCopyFromNetworkId returns the CopyFromNetworkId field if non-nil, zero value otherwise.

### GetCopyFromNetworkIdOk

`func (o *InlineObject267) GetCopyFromNetworkIdOk() (*string, bool)`

GetCopyFromNetworkIdOk returns a tuple with the CopyFromNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyFromNetworkId

`func (o *InlineObject267) SetCopyFromNetworkId(v string)`

SetCopyFromNetworkId sets CopyFromNetworkId field to given value.

### HasCopyFromNetworkId

`func (o *InlineObject267) HasCopyFromNetworkId() bool`

HasCopyFromNetworkId returns a boolean if a field has been set.

### GetNotes

`func (o *InlineObject267) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject267) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject267) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject267) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


