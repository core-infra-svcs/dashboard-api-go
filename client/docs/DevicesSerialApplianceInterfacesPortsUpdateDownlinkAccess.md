# DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to **string** | The VLAN for which this port will accept and pass traffic in &#39;access&#39; mode. All untagged traffic will automatically be treated as if it belonged to this VLAN | [optional] 
**Policy** | Pointer to [**DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccessPolicy**](DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccessPolicy.md) |  | [optional] 

## Methods

### NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess

`func NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess() *DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess`

NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess instantiates a new DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkAccessWithDefaults

`func NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkAccessWithDefaults() *DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess`

NewDevicesSerialApplianceInterfacesPortsUpdateDownlinkAccessWithDefaults instantiates a new DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetPolicy

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess) GetPolicy() DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccessPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess) GetPolicyOk() (*DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccessPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess) SetPolicy(v DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccessPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *DevicesSerialApplianceInterfacesPortsUpdateDownlinkAccess) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


