# InlineObject9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | [optional] The snapshot will be taken from this time on the camera. The timestamp is expected to be in ISO 8601 format. If no timestamp is specified, we will assume current time. | [optional] 
**Fullframe** | Pointer to **bool** | [optional] If set to \&quot;true\&quot; the snapshot will be taken at full sensor resolution. This will error if used with timestamp. | [optional] 

## Methods

### NewInlineObject9

`func NewInlineObject9() *InlineObject9`

NewInlineObject9 instantiates a new InlineObject9 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject9WithDefaults

`func NewInlineObject9WithDefaults() *InlineObject9`

NewInlineObject9WithDefaults instantiates a new InlineObject9 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *InlineObject9) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineObject9) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineObject9) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InlineObject9) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetFullframe

`func (o *InlineObject9) GetFullframe() bool`

GetFullframe returns the Fullframe field if non-nil, zero value otherwise.

### GetFullframeOk

`func (o *InlineObject9) GetFullframeOk() (*bool, bool)`

GetFullframeOk returns a tuple with the Fullframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullframe

`func (o *InlineObject9) SetFullframe(v bool)`

SetFullframe sets Fullframe field to given value.

### HasFullframe

`func (o *InlineObject9) HasFullframe() bool`

HasFullframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


