# DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NativeVlan** | Pointer to **string** | The Native VLAN for the port. All untagged traffic that comes in on this port will be treated as if it belonged to this VLAN. This can also be set to 0 to drop untagged traffic. | [optional] 
**AllowedVlans** | Pointer to **[]string** | The VLANs for which this port will accept and pass traffic in &#39;trunk&#39; mode. This must include the Native VLAN if one is set | [optional] 
**Sgt** | Pointer to [**DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunkSgt**](DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunkSgt.md) |  | [optional] 

## Methods

### NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk

`func NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk() *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk`

NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk instantiates a new DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunkWithDefaults

`func NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunkWithDefaults() *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk`

NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunkWithDefaults instantiates a new DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNativeVlan

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) GetNativeVlan() string`

GetNativeVlan returns the NativeVlan field if non-nil, zero value otherwise.

### GetNativeVlanOk

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) GetNativeVlanOk() (*string, bool)`

GetNativeVlanOk returns a tuple with the NativeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVlan

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) SetNativeVlan(v string)`

SetNativeVlan sets NativeVlan field to given value.

### HasNativeVlan

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) HasNativeVlan() bool`

HasNativeVlan returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) GetAllowedVlans() []string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) GetAllowedVlansOk() (*[]string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) SetAllowedVlans(v []string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetSgt

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) GetSgt() DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunkSgt`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) GetSgtOk() (*DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunkSgt, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) SetSgt(v DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunkSgt)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkTrunk) HasSgt() bool`

HasSgt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


