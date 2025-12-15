# NetworksNetworkIdWirelessSsidsAccessControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encryption** | Pointer to [**NetworksNetworkIdWirelessSsidsAccessControlEncryption**](NetworksNetworkIdWirelessSsidsAccessControlEncryption.md) |  | [optional] 
**Bandwidth** | Pointer to [**NetworksNetworkIdWirelessSsidsAccessControlBandwidth**](NetworksNetworkIdWirelessSsidsAccessControlBandwidth.md) |  | [optional] 
**ClientIpAssignment** | Pointer to [**NetworksNetworkIdWirelessSsidsAccessControlClientIpAssignment**](NetworksNetworkIdWirelessSsidsAccessControlClientIpAssignment.md) |  | [optional] 
**ClientsBlockedFromUsingLan** | Pointer to **bool** | Whether clients are blocked from using the LAN | [optional] 
**WiredClientsPartOfWifiNetwork** | Pointer to **bool** | Whether wired SSID clients are part of the Wi-Fi network | [optional] 
**Tunnel** | Pointer to [**NetworksNetworkIdWirelessSsidsAccessControlTunnel**](NetworksNetworkIdWirelessSsidsAccessControlTunnel.md) |  | [optional] 
**Vlan** | Pointer to [**NetworksNetworkIdWirelessSsidsAccessControlVlan**](NetworksNetworkIdWirelessSsidsAccessControlVlan.md) |  | [optional] 
**SplashPage** | Pointer to [**NetworksNetworkIdWirelessSsidsAccessControlSplashPage**](NetworksNetworkIdWirelessSsidsAccessControlSplashPage.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdWirelessSsidsAccessControl

`func NewNetworksNetworkIdWirelessSsidsAccessControl() *NetworksNetworkIdWirelessSsidsAccessControl`

NewNetworksNetworkIdWirelessSsidsAccessControl instantiates a new NetworksNetworkIdWirelessSsidsAccessControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessSsidsAccessControlWithDefaults

`func NewNetworksNetworkIdWirelessSsidsAccessControlWithDefaults() *NetworksNetworkIdWirelessSsidsAccessControl`

NewNetworksNetworkIdWirelessSsidsAccessControlWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsAccessControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryption

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetEncryption() NetworksNetworkIdWirelessSsidsAccessControlEncryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetEncryptionOk() (*NetworksNetworkIdWirelessSsidsAccessControlEncryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) SetEncryption(v NetworksNetworkIdWirelessSsidsAccessControlEncryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetBandwidth

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetBandwidth() NetworksNetworkIdWirelessSsidsAccessControlBandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetBandwidthOk() (*NetworksNetworkIdWirelessSsidsAccessControlBandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) SetBandwidth(v NetworksNetworkIdWirelessSsidsAccessControlBandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetClientIpAssignment

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetClientIpAssignment() NetworksNetworkIdWirelessSsidsAccessControlClientIpAssignment`

GetClientIpAssignment returns the ClientIpAssignment field if non-nil, zero value otherwise.

### GetClientIpAssignmentOk

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetClientIpAssignmentOk() (*NetworksNetworkIdWirelessSsidsAccessControlClientIpAssignment, bool)`

GetClientIpAssignmentOk returns a tuple with the ClientIpAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpAssignment

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) SetClientIpAssignment(v NetworksNetworkIdWirelessSsidsAccessControlClientIpAssignment)`

SetClientIpAssignment sets ClientIpAssignment field to given value.

### HasClientIpAssignment

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) HasClientIpAssignment() bool`

HasClientIpAssignment returns a boolean if a field has been set.

### GetClientsBlockedFromUsingLan

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetClientsBlockedFromUsingLan() bool`

GetClientsBlockedFromUsingLan returns the ClientsBlockedFromUsingLan field if non-nil, zero value otherwise.

### GetClientsBlockedFromUsingLanOk

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetClientsBlockedFromUsingLanOk() (*bool, bool)`

GetClientsBlockedFromUsingLanOk returns a tuple with the ClientsBlockedFromUsingLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsBlockedFromUsingLan

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) SetClientsBlockedFromUsingLan(v bool)`

SetClientsBlockedFromUsingLan sets ClientsBlockedFromUsingLan field to given value.

### HasClientsBlockedFromUsingLan

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) HasClientsBlockedFromUsingLan() bool`

HasClientsBlockedFromUsingLan returns a boolean if a field has been set.

### GetWiredClientsPartOfWifiNetwork

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetWiredClientsPartOfWifiNetwork() bool`

GetWiredClientsPartOfWifiNetwork returns the WiredClientsPartOfWifiNetwork field if non-nil, zero value otherwise.

### GetWiredClientsPartOfWifiNetworkOk

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetWiredClientsPartOfWifiNetworkOk() (*bool, bool)`

GetWiredClientsPartOfWifiNetworkOk returns a tuple with the WiredClientsPartOfWifiNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredClientsPartOfWifiNetwork

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) SetWiredClientsPartOfWifiNetwork(v bool)`

SetWiredClientsPartOfWifiNetwork sets WiredClientsPartOfWifiNetwork field to given value.

### HasWiredClientsPartOfWifiNetwork

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) HasWiredClientsPartOfWifiNetwork() bool`

HasWiredClientsPartOfWifiNetwork returns a boolean if a field has been set.

### GetTunnel

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetTunnel() NetworksNetworkIdWirelessSsidsAccessControlTunnel`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetTunnelOk() (*NetworksNetworkIdWirelessSsidsAccessControlTunnel, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) SetTunnel(v NetworksNetworkIdWirelessSsidsAccessControlTunnel)`

SetTunnel sets Tunnel field to given value.

### HasTunnel

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) HasTunnel() bool`

HasTunnel returns a boolean if a field has been set.

### GetVlan

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetVlan() NetworksNetworkIdWirelessSsidsAccessControlVlan`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetVlanOk() (*NetworksNetworkIdWirelessSsidsAccessControlVlan, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) SetVlan(v NetworksNetworkIdWirelessSsidsAccessControlVlan)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetSplashPage

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetSplashPage() NetworksNetworkIdWirelessSsidsAccessControlSplashPage`

GetSplashPage returns the SplashPage field if non-nil, zero value otherwise.

### GetSplashPageOk

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) GetSplashPageOk() (*NetworksNetworkIdWirelessSsidsAccessControlSplashPage, bool)`

GetSplashPageOk returns a tuple with the SplashPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPage

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) SetSplashPage(v NetworksNetworkIdWirelessSsidsAccessControlSplashPage)`

SetSplashPage sets SplashPage field to given value.

### HasSplashPage

`func (o *NetworksNetworkIdWirelessSsidsAccessControl) HasSplashPage() bool`

HasSplashPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


