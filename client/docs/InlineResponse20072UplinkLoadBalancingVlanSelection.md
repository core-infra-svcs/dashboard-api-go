# InlineResponse20072UplinkLoadBalancingVlanSelection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Selection mode (auto|device) | [optional] 
**ByDevice** | Pointer to [**[]InlineResponse20072UplinkLoadBalancingVlanSelectionByDevice**](InlineResponse20072UplinkLoadBalancingVlanSelectionByDevice.md) | Per-device VLAN selection | [optional] 

## Methods

### NewInlineResponse20072UplinkLoadBalancingVlanSelection

`func NewInlineResponse20072UplinkLoadBalancingVlanSelection() *InlineResponse20072UplinkLoadBalancingVlanSelection`

NewInlineResponse20072UplinkLoadBalancingVlanSelection instantiates a new InlineResponse20072UplinkLoadBalancingVlanSelection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20072UplinkLoadBalancingVlanSelectionWithDefaults

`func NewInlineResponse20072UplinkLoadBalancingVlanSelectionWithDefaults() *InlineResponse20072UplinkLoadBalancingVlanSelection`

NewInlineResponse20072UplinkLoadBalancingVlanSelectionWithDefaults instantiates a new InlineResponse20072UplinkLoadBalancingVlanSelection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20072UplinkLoadBalancingVlanSelection) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20072UplinkLoadBalancingVlanSelection) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20072UplinkLoadBalancingVlanSelection) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20072UplinkLoadBalancingVlanSelection) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetByDevice

`func (o *InlineResponse20072UplinkLoadBalancingVlanSelection) GetByDevice() []InlineResponse20072UplinkLoadBalancingVlanSelectionByDevice`

GetByDevice returns the ByDevice field if non-nil, zero value otherwise.

### GetByDeviceOk

`func (o *InlineResponse20072UplinkLoadBalancingVlanSelection) GetByDeviceOk() (*[]InlineResponse20072UplinkLoadBalancingVlanSelectionByDevice, bool)`

GetByDeviceOk returns a tuple with the ByDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByDevice

`func (o *InlineResponse20072UplinkLoadBalancingVlanSelection) SetByDevice(v []InlineResponse20072UplinkLoadBalancingVlanSelectionByDevice)`

SetByDevice sets ByDevice field to given value.

### HasByDevice

`func (o *InlineResponse20072UplinkLoadBalancingVlanSelection) HasByDevice() bool`

HasByDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


