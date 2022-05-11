# NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | E.g.: \&quot;any\&quot;, \&quot;0\&quot; (also means \&quot;any\&quot;), \&quot;8080\&quot;, \&quot;1-1024\&quot; | [optional] 
**Cidr** | Pointer to **string** | CIDR format address, or \&quot;any\&quot;. E.g.: \&quot;192.168.10.0/24\&quot;,  \&quot;192.168.10.1\&quot; (same as \&quot;192.168.10.1/32\&quot;), \&quot;0.0.0.0/0\&quot; (same as \&quot;any\&quot;) | [optional] 
**Network** | Pointer to **string** | Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: \&quot;L_12345678\&quot;. | [optional] 
**Vlan** | Pointer to **int32** | VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network. | [optional] 
**Host** | Pointer to **int32** | Host ID in the VLAN, should be used along with &#39;vlan&#39;, and not exceed the vlan subnet capacity. Currently only available under a template network. | [optional] 
**Fqdn** | Pointer to **string** | FQDN format address. Currently only availabe in &#39;destination&#39; of &#39;vpnTrafficUplinkPreference&#39; object. E.g.: &#39;www.google.com&#39; | [optional] 

## Methods

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1DestinationWithDefaults

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1DestinationWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1DestinationWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCidr

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVlan

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetHost

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetHost() int32`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetHostOk() (*int32, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetHost(v int32)`

SetHost sets Host field to given value.

### HasHost

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetFqdn

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


