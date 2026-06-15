# InlineResponse20043Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of routes in the routing table | [optional] 
**ByVrf** | Pointer to [**[]InlineResponse20043CountsByVrf**](InlineResponse20043CountsByVrf.md) | Routing table counts by VRF | [optional] 

## Methods

### NewInlineResponse20043Counts

`func NewInlineResponse20043Counts() *InlineResponse20043Counts`

NewInlineResponse20043Counts instantiates a new InlineResponse20043Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20043CountsWithDefaults

`func NewInlineResponse20043CountsWithDefaults() *InlineResponse20043Counts`

NewInlineResponse20043CountsWithDefaults instantiates a new InlineResponse20043Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse20043Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse20043Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse20043Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse20043Counts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByVrf

`func (o *InlineResponse20043Counts) GetByVrf() []InlineResponse20043CountsByVrf`

GetByVrf returns the ByVrf field if non-nil, zero value otherwise.

### GetByVrfOk

`func (o *InlineResponse20043Counts) GetByVrfOk() (*[]InlineResponse20043CountsByVrf, bool)`

GetByVrfOk returns a tuple with the ByVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByVrf

`func (o *InlineResponse20043Counts) SetByVrf(v []InlineResponse20043CountsByVrf)`

SetByVrf sets ByVrf field to given value.

### HasByVrf

`func (o *InlineResponse20043Counts) HasByVrf() bool`

HasByVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


