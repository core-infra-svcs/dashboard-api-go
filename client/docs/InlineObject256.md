# InlineObject256

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serials** | **[]string** | A list of Meraki Serials to migrate | 
**Target** | **string** | The controller or management mode to which the devices will be migrated | 

## Methods

### NewInlineObject256

`func NewInlineObject256(serials []string, target string, ) *InlineObject256`

NewInlineObject256 instantiates a new InlineObject256 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject256WithDefaults

`func NewInlineObject256WithDefaults() *InlineObject256`

NewInlineObject256WithDefaults instantiates a new InlineObject256 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerials

`func (o *InlineObject256) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject256) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject256) SetSerials(v []string)`

SetSerials sets Serials field to given value.


### GetTarget

`func (o *InlineObject256) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *InlineObject256) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *InlineObject256) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


