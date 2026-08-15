# InlineResponse200324Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByTunnelStatus** | [**InlineResponse200324CountsByTunnelStatus**](InlineResponse200324CountsByTunnelStatus.md) |  | 
**Total** | **int32** | Total connections | 

## Methods

### NewInlineResponse200324Counts

`func NewInlineResponse200324Counts(byTunnelStatus InlineResponse200324CountsByTunnelStatus, total int32, ) *InlineResponse200324Counts`

NewInlineResponse200324Counts instantiates a new InlineResponse200324Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200324CountsWithDefaults

`func NewInlineResponse200324CountsWithDefaults() *InlineResponse200324Counts`

NewInlineResponse200324CountsWithDefaults instantiates a new InlineResponse200324Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByTunnelStatus

`func (o *InlineResponse200324Counts) GetByTunnelStatus() InlineResponse200324CountsByTunnelStatus`

GetByTunnelStatus returns the ByTunnelStatus field if non-nil, zero value otherwise.

### GetByTunnelStatusOk

`func (o *InlineResponse200324Counts) GetByTunnelStatusOk() (*InlineResponse200324CountsByTunnelStatus, bool)`

GetByTunnelStatusOk returns a tuple with the ByTunnelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByTunnelStatus

`func (o *InlineResponse200324Counts) SetByTunnelStatus(v InlineResponse200324CountsByTunnelStatus)`

SetByTunnelStatus sets ByTunnelStatus field to given value.


### GetTotal

`func (o *InlineResponse200324Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200324Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200324Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


