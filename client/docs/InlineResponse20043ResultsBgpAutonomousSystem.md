# InlineResponse20043ResultsBgpAutonomousSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **[]int32** | The BGP autonomous system path reported by the device. When multiple values are present, they are ordered from the most recently prepended AS — typically the neighboring AS — toward the originating AS. | [optional] 

## Methods

### NewInlineResponse20043ResultsBgpAutonomousSystem

`func NewInlineResponse20043ResultsBgpAutonomousSystem() *InlineResponse20043ResultsBgpAutonomousSystem`

NewInlineResponse20043ResultsBgpAutonomousSystem instantiates a new InlineResponse20043ResultsBgpAutonomousSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20043ResultsBgpAutonomousSystemWithDefaults

`func NewInlineResponse20043ResultsBgpAutonomousSystemWithDefaults() *InlineResponse20043ResultsBgpAutonomousSystem`

NewInlineResponse20043ResultsBgpAutonomousSystemWithDefaults instantiates a new InlineResponse20043ResultsBgpAutonomousSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *InlineResponse20043ResultsBgpAutonomousSystem) GetPath() []int32`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *InlineResponse20043ResultsBgpAutonomousSystem) GetPathOk() (*[]int32, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *InlineResponse20043ResultsBgpAutonomousSystem) SetPath(v []int32)`

SetPath sets Path field to given value.

### HasPath

`func (o *InlineResponse20043ResultsBgpAutonomousSystem) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


