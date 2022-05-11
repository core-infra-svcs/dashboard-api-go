# NetworksNetworkIdSwitchLinkAggregationsSwitchPorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | **string** | Serial number of the switch. | 
**PortId** | **string** | Port identifier of switch port. For modules, the identifier is \&quot;SlotNumber_ModuleType_PortNumber\&quot; (Ex: \&quot;1_8X10G_1\&quot;), otherwise it is just the port number (Ex: \&quot;8\&quot;). | 

## Methods

### NewNetworksNetworkIdSwitchLinkAggregationsSwitchPorts

`func NewNetworksNetworkIdSwitchLinkAggregationsSwitchPorts(serial string, portId string, ) *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts`

NewNetworksNetworkIdSwitchLinkAggregationsSwitchPorts instantiates a new NetworksNetworkIdSwitchLinkAggregationsSwitchPorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchLinkAggregationsSwitchPortsWithDefaults

`func NewNetworksNetworkIdSwitchLinkAggregationsSwitchPortsWithDefaults() *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts`

NewNetworksNetworkIdSwitchLinkAggregationsSwitchPortsWithDefaults instantiates a new NetworksNetworkIdSwitchLinkAggregationsSwitchPorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetPortId

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) SetPortId(v string)`

SetPortId sets PortId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


