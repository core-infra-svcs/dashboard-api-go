# InlineResponse2005

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Desired UUID of the beacon. If the value is set to null it will reset to Dashboard&#39;s automatically generated value. | [optional] 
**Major** | Pointer to **int32** | Desired major value of the beacon. If the value is set to null it will reset to Dashboard&#39;s automatically generated value. | [optional] 
**Minor** | Pointer to **int32** | Desired minor value of the beacon. If the value is set to null it will reset to Dashboard&#39;s automatically generated value. | [optional] 

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

### GetUuid

`func (o *InlineResponse2005) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InlineResponse2005) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InlineResponse2005) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InlineResponse2005) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMajor

`func (o *InlineResponse2005) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *InlineResponse2005) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *InlineResponse2005) SetMajor(v int32)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *InlineResponse2005) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMinor

`func (o *InlineResponse2005) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *InlineResponse2005) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *InlineResponse2005) SetMinor(v int32)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *InlineResponse2005) HasMinor() bool`

HasMinor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


