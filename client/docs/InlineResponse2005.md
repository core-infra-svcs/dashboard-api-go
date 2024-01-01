# InlineResponse2005

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalRtspEnabled** | Pointer to **bool** | Boolean indicating if external rtsp stream is exposed | [optional] 
**RtspUrl** | Pointer to **string** | External rstp url. Will only be returned if external rtsp stream is exposed | [optional] 

## Methods

### NewInlineResponse2005

`func NewInlineResponse2005() *InlineResponse2005`

NewInlineResponse2005 instantiates a new InlineResponse2005 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2005WithDefaults

`func NewInlineResponse2005WithDefaults() *InlineResponse2005`

NewInlineResponse2005WithDefaults instantiates a new InlineResponse2005 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalRtspEnabled

`func (o *InlineResponse2005) GetExternalRtspEnabled() bool`

GetExternalRtspEnabled returns the ExternalRtspEnabled field if non-nil, zero value otherwise.

### GetExternalRtspEnabledOk

`func (o *InlineResponse2005) GetExternalRtspEnabledOk() (*bool, bool)`

GetExternalRtspEnabledOk returns a tuple with the ExternalRtspEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRtspEnabled

`func (o *InlineResponse2005) SetExternalRtspEnabled(v bool)`

SetExternalRtspEnabled sets ExternalRtspEnabled field to given value.

### HasExternalRtspEnabled

`func (o *InlineResponse2005) HasExternalRtspEnabled() bool`

HasExternalRtspEnabled returns a boolean if a field has been set.

### GetRtspUrl

`func (o *InlineResponse2005) GetRtspUrl() string`

GetRtspUrl returns the RtspUrl field if non-nil, zero value otherwise.

### GetRtspUrlOk

`func (o *InlineResponse2005) GetRtspUrlOk() (*string, bool)`

GetRtspUrlOk returns a tuple with the RtspUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtspUrl

`func (o *InlineResponse2005) SetRtspUrl(v string)`

SetRtspUrl sets RtspUrl field to given value.

### HasRtspUrl

`func (o *InlineResponse2005) HasRtspUrl() bool`

HasRtspUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


