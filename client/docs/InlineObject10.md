# InlineObject10

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **string** | FQDN, IPv4 or IPv6 address | 
**Count** | Pointer to **int32** | Count parameter to pass to ping. [1..5], default 5 | [optional] 

## Methods

### NewInlineObject10

`func NewInlineObject10(target string, ) *InlineObject10`

NewInlineObject10 instantiates a new InlineObject10 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject10WithDefaults

`func NewInlineObject10WithDefaults() *InlineObject10`

NewInlineObject10WithDefaults instantiates a new InlineObject10 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *InlineObject10) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *InlineObject10) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *InlineObject10) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetCount

`func (o *InlineObject10) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineObject10) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineObject10) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InlineObject10) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


