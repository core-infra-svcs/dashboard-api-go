# InlineResponse20045Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of routes in the routing table | [optional] 
**ByVrf** | Pointer to [**[]InlineResponse20045CountsByVrf**](InlineResponse20045CountsByVrf.md) | Routing table counts by VRF | [optional] 

## Methods

### NewInlineResponse20045Counts

`func NewInlineResponse20045Counts() *InlineResponse20045Counts`

NewInlineResponse20045Counts instantiates a new InlineResponse20045Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20045CountsWithDefaults

`func NewInlineResponse20045CountsWithDefaults() *InlineResponse20045Counts`

NewInlineResponse20045CountsWithDefaults instantiates a new InlineResponse20045Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse20045Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse20045Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse20045Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse20045Counts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByVrf

`func (o *InlineResponse20045Counts) GetByVrf() []InlineResponse20045CountsByVrf`

GetByVrf returns the ByVrf field if non-nil, zero value otherwise.

### GetByVrfOk

`func (o *InlineResponse20045Counts) GetByVrfOk() (*[]InlineResponse20045CountsByVrf, bool)`

GetByVrfOk returns a tuple with the ByVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByVrf

`func (o *InlineResponse20045Counts) SetByVrf(v []InlineResponse20045CountsByVrf)`

SetByVrf sets ByVrf field to given value.

### HasByVrf

`func (o *InlineResponse20045Counts) HasByVrf() bool`

HasByVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


