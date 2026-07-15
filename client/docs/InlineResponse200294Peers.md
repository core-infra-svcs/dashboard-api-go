# InlineResponse200294Peers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerId** | Pointer to **string** | The ID of the IPsec peer | [optional] 
**Name** | Pointer to **string** | The name of the VPN peer | [optional] 
**PublicIp** | Pointer to **string** | [optional] The public IP of the VPN peer | [optional] 
**RemoteId** | Pointer to **string** | [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN. | [optional] 
**LocalId** | Pointer to **string** | [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to. | [optional] 
**Secret** | Pointer to **string** | The shared secret with the VPN peer | [optional] 
**PrivateSubnets** | Pointer to **[]string** | The list of the private subnets of the VPN peer | [optional] 
**IpsecPolicies** | Pointer to [**InlineResponse200294IpsecPolicies**](InlineResponse200294IpsecPolicies.md) |  | [optional] 
**SlaPolicy** | Pointer to [**InlineResponse200294SlaPolicy**](InlineResponse200294SlaPolicy.md) |  | [optional] 
**IpsecPoliciesPreset** | Pointer to **string** | One of the following available presets: &#39;default&#39;, &#39;aws&#39;, &#39;azure&#39;, &#39;umbrella&#39;, &#39;umbrella_short_lived&#39;, &#39;secure&#39;, &#39;zscaler&#39;. If this is provided, the &#39;ipsecPolicies&#39; parameter is ignored. Note that &#39;umbrella_short_lived&#39; maps to the Dashboard UI label &#39;Umbrella&#39;, while &#39;umbrella&#39; maps to the Dashboard UI label &#39;Umbrella (Deprecated)&#39;. | [optional] 
**IkeVersion** | Pointer to **string** | [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to &#39;1&#39; when omitted. | [optional] [default to "1"]
**NetworkTags** | Pointer to **[]string** | A list of network tags that will connect with this peer. Use [&#39;all&#39;] for all networks. Use [&#39;none&#39;] for no networks. If not included, the default is [&#39;all&#39;]. | [optional] 
**Network** | Pointer to [**InlineResponse200294Network**](InlineResponse200294Network.md) |  | [optional] 
**IsRouteBased** | Pointer to **bool** | [optional] If true, the VPN peer is route-based. If not included, the default is false. Supported only for MX 19.1 and above. | [optional] 
**EbgpNeighbor** | Pointer to [**InlineResponse200294EbgpNeighbor**](InlineResponse200294EbgpNeighbor.md) |  | [optional] 
**EcmpUplinkConfigs** | Pointer to [**[]InlineResponse200294EcmpUplinkConfigs**](InlineResponse200294EcmpUplinkConfigs.md) | [optional] The ECMP per-uplink BGP-over-IPsec configuration for the VPN peer. | [optional] 
**PriorityInGroup** | Pointer to **int32** | [optional] Represents the order of peer inside a group. | [optional] 
**Group** | Pointer to [**InlineResponse200294Group**](InlineResponse200294Group.md) |  | [optional] 

## Methods

### NewInlineResponse200294Peers

`func NewInlineResponse200294Peers() *InlineResponse200294Peers`

NewInlineResponse200294Peers instantiates a new InlineResponse200294Peers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200294PeersWithDefaults

`func NewInlineResponse200294PeersWithDefaults() *InlineResponse200294Peers`

NewInlineResponse200294PeersWithDefaults instantiates a new InlineResponse200294Peers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerId

`func (o *InlineResponse200294Peers) GetPeerId() string`

GetPeerId returns the PeerId field if non-nil, zero value otherwise.

### GetPeerIdOk

`func (o *InlineResponse200294Peers) GetPeerIdOk() (*string, bool)`

GetPeerIdOk returns a tuple with the PeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerId

`func (o *InlineResponse200294Peers) SetPeerId(v string)`

SetPeerId sets PeerId field to given value.

### HasPeerId

`func (o *InlineResponse200294Peers) HasPeerId() bool`

HasPeerId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200294Peers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200294Peers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200294Peers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200294Peers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *InlineResponse200294Peers) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InlineResponse200294Peers) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InlineResponse200294Peers) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InlineResponse200294Peers) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRemoteId

`func (o *InlineResponse200294Peers) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *InlineResponse200294Peers) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *InlineResponse200294Peers) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *InlineResponse200294Peers) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetLocalId

`func (o *InlineResponse200294Peers) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *InlineResponse200294Peers) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *InlineResponse200294Peers) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *InlineResponse200294Peers) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetSecret

`func (o *InlineResponse200294Peers) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *InlineResponse200294Peers) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *InlineResponse200294Peers) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *InlineResponse200294Peers) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *InlineResponse200294Peers) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *InlineResponse200294Peers) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *InlineResponse200294Peers) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *InlineResponse200294Peers) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetIpsecPolicies

`func (o *InlineResponse200294Peers) GetIpsecPolicies() InlineResponse200294IpsecPolicies`

GetIpsecPolicies returns the IpsecPolicies field if non-nil, zero value otherwise.

### GetIpsecPoliciesOk

`func (o *InlineResponse200294Peers) GetIpsecPoliciesOk() (*InlineResponse200294IpsecPolicies, bool)`

GetIpsecPoliciesOk returns a tuple with the IpsecPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicies

`func (o *InlineResponse200294Peers) SetIpsecPolicies(v InlineResponse200294IpsecPolicies)`

SetIpsecPolicies sets IpsecPolicies field to given value.

### HasIpsecPolicies

`func (o *InlineResponse200294Peers) HasIpsecPolicies() bool`

HasIpsecPolicies returns a boolean if a field has been set.

### GetSlaPolicy

`func (o *InlineResponse200294Peers) GetSlaPolicy() InlineResponse200294SlaPolicy`

GetSlaPolicy returns the SlaPolicy field if non-nil, zero value otherwise.

### GetSlaPolicyOk

`func (o *InlineResponse200294Peers) GetSlaPolicyOk() (*InlineResponse200294SlaPolicy, bool)`

GetSlaPolicyOk returns a tuple with the SlaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaPolicy

`func (o *InlineResponse200294Peers) SetSlaPolicy(v InlineResponse200294SlaPolicy)`

SetSlaPolicy sets SlaPolicy field to given value.

### HasSlaPolicy

`func (o *InlineResponse200294Peers) HasSlaPolicy() bool`

HasSlaPolicy returns a boolean if a field has been set.

### GetIpsecPoliciesPreset

`func (o *InlineResponse200294Peers) GetIpsecPoliciesPreset() string`

GetIpsecPoliciesPreset returns the IpsecPoliciesPreset field if non-nil, zero value otherwise.

### GetIpsecPoliciesPresetOk

`func (o *InlineResponse200294Peers) GetIpsecPoliciesPresetOk() (*string, bool)`

GetIpsecPoliciesPresetOk returns a tuple with the IpsecPoliciesPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPoliciesPreset

`func (o *InlineResponse200294Peers) SetIpsecPoliciesPreset(v string)`

SetIpsecPoliciesPreset sets IpsecPoliciesPreset field to given value.

### HasIpsecPoliciesPreset

`func (o *InlineResponse200294Peers) HasIpsecPoliciesPreset() bool`

HasIpsecPoliciesPreset returns a boolean if a field has been set.

### GetIkeVersion

`func (o *InlineResponse200294Peers) GetIkeVersion() string`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *InlineResponse200294Peers) GetIkeVersionOk() (*string, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *InlineResponse200294Peers) SetIkeVersion(v string)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *InlineResponse200294Peers) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetNetworkTags

`func (o *InlineResponse200294Peers) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InlineResponse200294Peers) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InlineResponse200294Peers) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *InlineResponse200294Peers) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200294Peers) GetNetwork() InlineResponse200294Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200294Peers) GetNetworkOk() (*InlineResponse200294Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200294Peers) SetNetwork(v InlineResponse200294Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200294Peers) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIsRouteBased

`func (o *InlineResponse200294Peers) GetIsRouteBased() bool`

GetIsRouteBased returns the IsRouteBased field if non-nil, zero value otherwise.

### GetIsRouteBasedOk

`func (o *InlineResponse200294Peers) GetIsRouteBasedOk() (*bool, bool)`

GetIsRouteBasedOk returns a tuple with the IsRouteBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRouteBased

`func (o *InlineResponse200294Peers) SetIsRouteBased(v bool)`

SetIsRouteBased sets IsRouteBased field to given value.

### HasIsRouteBased

`func (o *InlineResponse200294Peers) HasIsRouteBased() bool`

HasIsRouteBased returns a boolean if a field has been set.

### GetEbgpNeighbor

`func (o *InlineResponse200294Peers) GetEbgpNeighbor() InlineResponse200294EbgpNeighbor`

GetEbgpNeighbor returns the EbgpNeighbor field if non-nil, zero value otherwise.

### GetEbgpNeighborOk

`func (o *InlineResponse200294Peers) GetEbgpNeighborOk() (*InlineResponse200294EbgpNeighbor, bool)`

GetEbgpNeighborOk returns a tuple with the EbgpNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpNeighbor

`func (o *InlineResponse200294Peers) SetEbgpNeighbor(v InlineResponse200294EbgpNeighbor)`

SetEbgpNeighbor sets EbgpNeighbor field to given value.

### HasEbgpNeighbor

`func (o *InlineResponse200294Peers) HasEbgpNeighbor() bool`

HasEbgpNeighbor returns a boolean if a field has been set.

### GetEcmpUplinkConfigs

`func (o *InlineResponse200294Peers) GetEcmpUplinkConfigs() []InlineResponse200294EcmpUplinkConfigs`

GetEcmpUplinkConfigs returns the EcmpUplinkConfigs field if non-nil, zero value otherwise.

### GetEcmpUplinkConfigsOk

`func (o *InlineResponse200294Peers) GetEcmpUplinkConfigsOk() (*[]InlineResponse200294EcmpUplinkConfigs, bool)`

GetEcmpUplinkConfigsOk returns a tuple with the EcmpUplinkConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcmpUplinkConfigs

`func (o *InlineResponse200294Peers) SetEcmpUplinkConfigs(v []InlineResponse200294EcmpUplinkConfigs)`

SetEcmpUplinkConfigs sets EcmpUplinkConfigs field to given value.

### HasEcmpUplinkConfigs

`func (o *InlineResponse200294Peers) HasEcmpUplinkConfigs() bool`

HasEcmpUplinkConfigs returns a boolean if a field has been set.

### GetPriorityInGroup

`func (o *InlineResponse200294Peers) GetPriorityInGroup() int32`

GetPriorityInGroup returns the PriorityInGroup field if non-nil, zero value otherwise.

### GetPriorityInGroupOk

`func (o *InlineResponse200294Peers) GetPriorityInGroupOk() (*int32, bool)`

GetPriorityInGroupOk returns a tuple with the PriorityInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityInGroup

`func (o *InlineResponse200294Peers) SetPriorityInGroup(v int32)`

SetPriorityInGroup sets PriorityInGroup field to given value.

### HasPriorityInGroup

`func (o *InlineResponse200294Peers) HasPriorityInGroup() bool`

HasPriorityInGroup returns a boolean if a field has been set.

### GetGroup

`func (o *InlineResponse200294Peers) GetGroup() InlineResponse200294Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse200294Peers) GetGroupOk() (*InlineResponse200294Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse200294Peers) SetGroup(v InlineResponse200294Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse200294Peers) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


