# InlineResponse20080

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InlineResponse20080Network**](InlineResponse20080Network.md) |  | 
**Rules** | [**[]NetworksNetworkIdApplianceFirewallMulticastForwardingRules**](NetworksNetworkIdApplianceFirewallMulticastForwardingRules.md) | Static multicast forwarding rules. | 

## Methods

### NewInlineResponse20080

`func NewInlineResponse20080(network InlineResponse20080Network, rules []NetworksNetworkIdApplianceFirewallMulticastForwardingRules, ) *InlineResponse20080`

NewInlineResponse20080 instantiates a new InlineResponse20080 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20080WithDefaults

`func NewInlineResponse20080WithDefaults() *InlineResponse20080`

NewInlineResponse20080WithDefaults instantiates a new InlineResponse20080 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse20080) GetNetwork() InlineResponse20080Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20080) GetNetworkOk() (*InlineResponse20080Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20080) SetNetwork(v InlineResponse20080Network)`

SetNetwork sets Network field to given value.


### GetRules

`func (o *InlineResponse20080) GetRules() []NetworksNetworkIdApplianceFirewallMulticastForwardingRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse20080) GetRulesOk() (*[]NetworksNetworkIdApplianceFirewallMulticastForwardingRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse20080) SetRules(v []NetworksNetworkIdApplianceFirewallMulticastForwardingRules)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


