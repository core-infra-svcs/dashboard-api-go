# NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | E.g.: \&quot;any\&quot;, \&quot;0\&quot; (also means \&quot;any\&quot;), \&quot;8080\&quot;, \&quot;1-1024\&quot; | [optional] 
**Cidr** | Pointer to **string** | CIDR format address, or \&quot;any\&quot;. E.g.: \&quot;192.168.10.0/24\&quot;,  \&quot;192.168.10.1\&quot; (same as \&quot;192.168.10.1/32\&quot;), \&quot;0.0.0.0/0\&quot; (same as \&quot;any\&quot;) | [optional] 

## Methods

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestinationWithDefaults

`func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestinationWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination`

NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestinationWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCidr

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueDestination) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


