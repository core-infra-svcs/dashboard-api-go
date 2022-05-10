# NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional) | [optional] 
**Policy** | **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule | 
**Protocol** | **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39;, &#39;icmp6&#39; or &#39;any&#39;) | 
**DestPort** | Pointer to **string** | Comma-separated list of destination port(s) (integer in the range 1-65535), or &#39;any&#39; | [optional] 
**DestCidr** | **string** | Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or &#39;any&#39; | 

## Methods

### NewNetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules

`func NewNetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules(policy string, protocol string, destCidr string, ) *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules`

NewNetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules instantiates a new NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRulesWithDefaults

`func NewNetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRulesWithDefaults() *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules`

NewNetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRulesWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetProtocol

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestPort

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


