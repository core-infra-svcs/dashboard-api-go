# InlineObject251

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the configuration template | 
**TimeZone** | Pointer to **string** | The timezone of the configuration template. For a list of allowed timezones, please see the &#39;TZ&#39; column in the table in &lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones&#39;&gt;this article&lt;/a&gt;. Not applicable if copying from existing network or template | [optional] 
**CopyFromNetworkId** | Pointer to **string** | The ID of the network or config template to copy configuration from | [optional] 

## Methods

### NewInlineObject251

`func NewInlineObject251(name string, ) *InlineObject251`

NewInlineObject251 instantiates a new InlineObject251 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject251WithDefaults

`func NewInlineObject251WithDefaults() *InlineObject251`

NewInlineObject251WithDefaults instantiates a new InlineObject251 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject251) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject251) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject251) SetName(v string)`

SetName sets Name field to given value.


### GetTimeZone

`func (o *InlineObject251) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *InlineObject251) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *InlineObject251) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *InlineObject251) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetCopyFromNetworkId

`func (o *InlineObject251) GetCopyFromNetworkId() string`

GetCopyFromNetworkId returns the CopyFromNetworkId field if non-nil, zero value otherwise.

### GetCopyFromNetworkIdOk

`func (o *InlineObject251) GetCopyFromNetworkIdOk() (*string, bool)`

GetCopyFromNetworkIdOk returns a tuple with the CopyFromNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyFromNetworkId

`func (o *InlineObject251) SetCopyFromNetworkId(v string)`

SetCopyFromNetworkId sets CopyFromNetworkId field to given value.

### HasCopyFromNetworkId

`func (o *InlineObject251) HasCopyFromNetworkId() bool`

HasCopyFromNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


