# InlineResponse20043ResultsBgp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The BGP route type reported by the device. | [optional] 
**Metrics** | Pointer to [**InlineResponse20043ResultsBgpMetrics**](InlineResponse20043ResultsBgpMetrics.md) |  | [optional] 
**AutonomousSystem** | Pointer to [**InlineResponse20043ResultsBgpAutonomousSystem**](InlineResponse20043ResultsBgpAutonomousSystem.md) |  | [optional] 
**Communities** | Pointer to [**InlineResponse20043ResultsBgpCommunities**](InlineResponse20043ResultsBgpCommunities.md) |  | [optional] 

## Methods

### NewInlineResponse20043ResultsBgp

`func NewInlineResponse20043ResultsBgp() *InlineResponse20043ResultsBgp`

NewInlineResponse20043ResultsBgp instantiates a new InlineResponse20043ResultsBgp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20043ResultsBgpWithDefaults

`func NewInlineResponse20043ResultsBgpWithDefaults() *InlineResponse20043ResultsBgp`

NewInlineResponse20043ResultsBgpWithDefaults instantiates a new InlineResponse20043ResultsBgp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse20043ResultsBgp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20043ResultsBgp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20043ResultsBgp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20043ResultsBgp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetrics

`func (o *InlineResponse20043ResultsBgp) GetMetrics() InlineResponse20043ResultsBgpMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *InlineResponse20043ResultsBgp) GetMetricsOk() (*InlineResponse20043ResultsBgpMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *InlineResponse20043ResultsBgp) SetMetrics(v InlineResponse20043ResultsBgpMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *InlineResponse20043ResultsBgp) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetAutonomousSystem

`func (o *InlineResponse20043ResultsBgp) GetAutonomousSystem() InlineResponse20043ResultsBgpAutonomousSystem`

GetAutonomousSystem returns the AutonomousSystem field if non-nil, zero value otherwise.

### GetAutonomousSystemOk

`func (o *InlineResponse20043ResultsBgp) GetAutonomousSystemOk() (*InlineResponse20043ResultsBgpAutonomousSystem, bool)`

GetAutonomousSystemOk returns a tuple with the AutonomousSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomousSystem

`func (o *InlineResponse20043ResultsBgp) SetAutonomousSystem(v InlineResponse20043ResultsBgpAutonomousSystem)`

SetAutonomousSystem sets AutonomousSystem field to given value.

### HasAutonomousSystem

`func (o *InlineResponse20043ResultsBgp) HasAutonomousSystem() bool`

HasAutonomousSystem returns a boolean if a field has been set.

### GetCommunities

`func (o *InlineResponse20043ResultsBgp) GetCommunities() InlineResponse20043ResultsBgpCommunities`

GetCommunities returns the Communities field if non-nil, zero value otherwise.

### GetCommunitiesOk

`func (o *InlineResponse20043ResultsBgp) GetCommunitiesOk() (*InlineResponse20043ResultsBgpCommunities, bool)`

GetCommunitiesOk returns a tuple with the Communities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunities

`func (o *InlineResponse20043ResultsBgp) SetCommunities(v InlineResponse20043ResultsBgpCommunities)`

SetCommunities sets Communities field to given value.

### HasCommunities

`func (o *InlineResponse20043ResultsBgp) HasCommunities() bool`

HasCommunities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


