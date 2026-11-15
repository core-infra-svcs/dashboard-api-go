# InlineResponse20043Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** | Whether the device reports the route as active. | [optional] 
**Bgp** | Pointer to [**InlineResponse20043ResultsBgp**](InlineResponse20043ResultsBgp.md) |  | [optional] 

## Methods

### NewInlineResponse20043Results

`func NewInlineResponse20043Results() *InlineResponse20043Results`

NewInlineResponse20043Results instantiates a new InlineResponse20043Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20043ResultsWithDefaults

`func NewInlineResponse20043ResultsWithDefaults() *InlineResponse20043Results`

NewInlineResponse20043ResultsWithDefaults instantiates a new InlineResponse20043Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *InlineResponse20043Results) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *InlineResponse20043Results) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *InlineResponse20043Results) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *InlineResponse20043Results) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetBgp

`func (o *InlineResponse20043Results) GetBgp() InlineResponse20043ResultsBgp`

GetBgp returns the Bgp field if non-nil, zero value otherwise.

### GetBgpOk

`func (o *InlineResponse20043Results) GetBgpOk() (*InlineResponse20043ResultsBgp, bool)`

GetBgpOk returns a tuple with the Bgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgp

`func (o *InlineResponse20043Results) SetBgp(v InlineResponse20043ResultsBgp)`

SetBgp sets Bgp field to given value.

### HasBgp

`func (o *InlineResponse20043Results) HasBgp() bool`

HasBgp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


