# InlineResponse20030

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceMac** | Pointer to **string** | Source MAC address | [optional] 
**Ports** | Pointer to [**map[string]InlineResponse20030Ports**](InlineResponse20030Ports.md) | Mapping of ports to lldp and/or cdp information | [optional] 

## Methods

### NewInlineResponse20030

`func NewInlineResponse20030() *InlineResponse20030`

NewInlineResponse20030 instantiates a new InlineResponse20030 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20030WithDefaults

`func NewInlineResponse20030WithDefaults() *InlineResponse20030`

NewInlineResponse20030WithDefaults instantiates a new InlineResponse20030 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceMac

`func (o *InlineResponse20030) GetSourceMac() string`

GetSourceMac returns the SourceMac field if non-nil, zero value otherwise.

### GetSourceMacOk

`func (o *InlineResponse20030) GetSourceMacOk() (*string, bool)`

GetSourceMacOk returns a tuple with the SourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMac

`func (o *InlineResponse20030) SetSourceMac(v string)`

SetSourceMac sets SourceMac field to given value.

### HasSourceMac

`func (o *InlineResponse20030) HasSourceMac() bool`

HasSourceMac returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse20030) GetPorts() map[string]InlineResponse20030Ports`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse20030) GetPortsOk() (*map[string]InlineResponse20030Ports, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse20030) SetPorts(v map[string]InlineResponse20030Ports)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse20030) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


