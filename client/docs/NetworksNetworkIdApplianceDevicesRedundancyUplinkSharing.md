# NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable uplink sharing | [optional] 
**VlanId** | Pointer to **NullableString** | Uplink sharing VLAN ID | [optional] 
**ByInterface** | Pointer to [**[]NetworksNetworkIdApplianceDevicesRedundancyUplinkSharingByInterface**](NetworksNetworkIdApplianceDevicesRedundancyUplinkSharingByInterface.md) | Per-interface sharing configuration | [optional] 

## Methods

### NewNetworksNetworkIdApplianceDevicesRedundancyUplinkSharing

`func NewNetworksNetworkIdApplianceDevicesRedundancyUplinkSharing() *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing`

NewNetworksNetworkIdApplianceDevicesRedundancyUplinkSharing instantiates a new NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceDevicesRedundancyUplinkSharingWithDefaults

`func NewNetworksNetworkIdApplianceDevicesRedundancyUplinkSharingWithDefaults() *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing`

NewNetworksNetworkIdApplianceDevicesRedundancyUplinkSharingWithDefaults instantiates a new NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVlanId

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetByInterface

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) GetByInterface() []NetworksNetworkIdApplianceDevicesRedundancyUplinkSharingByInterface`

GetByInterface returns the ByInterface field if non-nil, zero value otherwise.

### GetByInterfaceOk

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) GetByInterfaceOk() (*[]NetworksNetworkIdApplianceDevicesRedundancyUplinkSharingByInterface, bool)`

GetByInterfaceOk returns a tuple with the ByInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByInterface

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) SetByInterface(v []NetworksNetworkIdApplianceDevicesRedundancyUplinkSharingByInterface)`

SetByInterface sets ByInterface field to given value.

### HasByInterface

`func (o *NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing) HasByInterface() bool`

HasByInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


