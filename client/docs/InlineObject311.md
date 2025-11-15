# InlineObject311

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enrolled** | **bool** | Parameter to enroll or unenroll the zigbee devices | 
**Channel** | Pointer to **string** | The new channel for the zigbee device | [optional] 

## Methods

### NewInlineObject311

`func NewInlineObject311(enrolled bool, ) *InlineObject311`

NewInlineObject311 instantiates a new InlineObject311 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject311WithDefaults

`func NewInlineObject311WithDefaults() *InlineObject311`

NewInlineObject311WithDefaults instantiates a new InlineObject311 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnrolled

`func (o *InlineObject311) GetEnrolled() bool`

GetEnrolled returns the Enrolled field if non-nil, zero value otherwise.

### GetEnrolledOk

`func (o *InlineObject311) GetEnrolledOk() (*bool, bool)`

GetEnrolledOk returns a tuple with the Enrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolled

`func (o *InlineObject311) SetEnrolled(v bool)`

SetEnrolled sets Enrolled field to given value.


### GetChannel

`func (o *InlineObject311) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineObject311) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineObject311) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineObject311) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


