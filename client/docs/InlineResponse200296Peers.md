# InlineResponse200296Peers

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
**IpsecPolicies** | Pointer to [**InlineResponse200296IpsecPolicies**](InlineResponse200296IpsecPolicies.md) |  | [optional] 
**SlaPolicy** | Pointer to [**InlineResponse200296SlaPolicy**](InlineResponse200296SlaPolicy.md) |  | [optional] 
**IpsecPoliciesPreset** | Pointer to **string** | One of the following available presets: &#39;default&#39;, &#39;aws&#39;, &#39;azure&#39;, &#39;umbrella&#39;, &#39;umbrella_short_lived&#39;, &#39;secure&#39;, &#39;zscaler&#39;. If this is provided, the &#39;ipsecPolicies&#39; parameter is ignored. Note that &#39;umbrella_short_lived&#39; maps to the Dashboard UI label &#39;Umbrella&#39;, while &#39;umbrella&#39; maps to the Dashboard UI label &#39;Umbrella (Deprecated)&#39;. | [optional] 
**IkeVersion** | Pointer to **string** | [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to &#39;1&#39; when omitted. | [optional] [default to "1"]
**NetworkTags** | Pointer to **[]string** | A list of network tags that will connect with this peer. Use [&#39;all&#39;] for all networks. Use [&#39;none&#39;] for no networks. If not included, the default is [&#39;all&#39;]. | [optional] 
**Network** | Pointer to [**InlineResponse200296Network**](InlineResponse200296Network.md) |  | [optional] 
**IsRouteBased** | Pointer to **bool** | [optional] If true, the VPN peer is route-based. If not included, the default is false. Supported only for MX 19.1 and above. | [optional] 
**EbgpNeighbor** | Pointer to [**InlineResponse200296EbgpNeighbor**](InlineResponse200296EbgpNeighbor.md) |  | [optional] 
**EcmpUplinkConfigs** | Pointer to [**[]InlineResponse200296EcmpUplinkConfigs**](InlineResponse200296EcmpUplinkConfigs.md) | [optional] The ECMP per-uplink BGP-over-IPsec configuration for the VPN peer. | [optional] 
**PriorityInGroup** | Pointer to **int32** | [optional] Represents the order of peer inside a group. | [optional] 
**Group** | Pointer to [**InlineResponse200296Group**](InlineResponse200296Group.md) |  | [optional] 

## Methods

### NewInlineResponse200296Peers

`func NewInlineResponse200296Peers() *InlineResponse200296Peers`

NewInlineResponse200296Peers instantiates a new InlineResponse200296Peers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200296PeersWithDefaults

`func NewInlineResponse200296PeersWithDefaults() *InlineResponse200296Peers`

NewInlineResponse200296PeersWithDefaults instantiates a new InlineResponse200296Peers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerId

`func (o *InlineResponse200296Peers) GetPeerId() string`

GetPeerId returns the PeerId field if non-nil, zero value otherwise.

### GetPeerIdOk

`func (o *InlineResponse200296Peers) GetPeerIdOk() (*string, bool)`

GetPeerIdOk returns a tuple with the PeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerId

`func (o *InlineResponse200296Peers) SetPeerId(v string)`

SetPeerId sets PeerId field to given value.

### HasPeerId

`func (o *InlineResponse200296Peers) HasPeerId() bool`

HasPeerId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200296Peers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200296Peers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200296Peers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200296Peers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *InlineResponse200296Peers) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InlineResponse200296Peers) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InlineResponse200296Peers) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InlineResponse200296Peers) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRemoteId

`func (o *InlineResponse200296Peers) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *InlineResponse200296Peers) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *InlineResponse200296Peers) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *InlineResponse200296Peers) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetLocalId

`func (o *InlineResponse200296Peers) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *InlineResponse200296Peers) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *InlineResponse200296Peers) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *InlineResponse200296Peers) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetSecret

`func (o *InlineResponse200296Peers) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *InlineResponse200296Peers) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *InlineResponse200296Peers) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *InlineResponse200296Peers) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *InlineResponse200296Peers) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *InlineResponse200296Peers) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *InlineResponse200296Peers) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *InlineResponse200296Peers) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetIpsecPolicies

`func (o *InlineResponse200296Peers) GetIpsecPolicies() InlineResponse200296IpsecPolicies`

GetIpsecPolicies returns the IpsecPolicies field if non-nil, zero value otherwise.

### GetIpsecPoliciesOk

`func (o *InlineResponse200296Peers) GetIpsecPoliciesOk() (*InlineResponse200296IpsecPolicies, bool)`

GetIpsecPoliciesOk returns a tuple with the IpsecPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicies

`func (o *InlineResponse200296Peers) SetIpsecPolicies(v InlineResponse200296IpsecPolicies)`

SetIpsecPolicies sets IpsecPolicies field to given value.

### HasIpsecPolicies

`func (o *InlineResponse200296Peers) HasIpsecPolicies() bool`

HasIpsecPolicies returns a boolean if a field has been set.

### GetSlaPolicy

`func (o *InlineResponse200296Peers) GetSlaPolicy() InlineResponse200296SlaPolicy`

GetSlaPolicy returns the SlaPolicy field if non-nil, zero value otherwise.

### GetSlaPolicyOk

`func (o *InlineResponse200296Peers) GetSlaPolicyOk() (*InlineResponse200296SlaPolicy, bool)`

GetSlaPolicyOk returns a tuple with the SlaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaPolicy

`func (o *InlineResponse200296Peers) SetSlaPolicy(v InlineResponse200296SlaPolicy)`

SetSlaPolicy sets SlaPolicy field to given value.

### HasSlaPolicy

`func (o *InlineResponse200296Peers) HasSlaPolicy() bool`

HasSlaPolicy returns a boolean if a field has been set.

### GetIpsecPoliciesPreset

`func (o *InlineResponse200296Peers) GetIpsecPoliciesPreset() string`

GetIpsecPoliciesPreset returns the IpsecPoliciesPreset field if non-nil, zero value otherwise.

### GetIpsecPoliciesPresetOk

`func (o *InlineResponse200296Peers) GetIpsecPoliciesPresetOk() (*string, bool)`

GetIpsecPoliciesPresetOk returns a tuple with the IpsecPoliciesPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPoliciesPreset

`func (o *InlineResponse200296Peers) SetIpsecPoliciesPreset(v string)`

SetIpsecPoliciesPreset sets IpsecPoliciesPreset field to given value.

### HasIpsecPoliciesPreset

`func (o *InlineResponse200296Peers) HasIpsecPoliciesPreset() bool`

HasIpsecPoliciesPreset returns a boolean if a field has been set.

### GetIkeVersion

`func (o *InlineResponse200296Peers) GetIkeVersion() string`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *InlineResponse200296Peers) GetIkeVersionOk() (*string, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *InlineResponse200296Peers) SetIkeVersion(v string)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *InlineResponse200296Peers) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetNetworkTags

`func (o *InlineResponse200296Peers) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InlineResponse200296Peers) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InlineResponse200296Peers) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *InlineResponse200296Peers) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200296Peers) GetNetwork() InlineResponse200296Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200296Peers) GetNetworkOk() (*InlineResponse200296Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200296Peers) SetNetwork(v InlineResponse200296Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200296Peers) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIsRouteBased

`func (o *InlineResponse200296Peers) GetIsRouteBased() bool`

GetIsRouteBased returns the IsRouteBased field if non-nil, zero value otherwise.

### GetIsRouteBasedOk

`func (o *InlineResponse200296Peers) GetIsRouteBasedOk() (*bool, bool)`

GetIsRouteBasedOk returns a tuple with the IsRouteBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRouteBased

`func (o *InlineResponse200296Peers) SetIsRouteBased(v bool)`

SetIsRouteBased sets IsRouteBased field to given value.

### HasIsRouteBased

`func (o *InlineResponse200296Peers) HasIsRouteBased() bool`

HasIsRouteBased returns a boolean if a field has been set.

### GetEbgpNeighbor

`func (o *InlineResponse200296Peers) GetEbgpNeighbor() InlineResponse200296EbgpNeighbor`

GetEbgpNeighbor returns the EbgpNeighbor field if non-nil, zero value otherwise.

### GetEbgpNeighborOk

`func (o *InlineResponse200296Peers) GetEbgpNeighborOk() (*InlineResponse200296EbgpNeighbor, bool)`

GetEbgpNeighborOk returns a tuple with the EbgpNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpNeighbor

`func (o *InlineResponse200296Peers) SetEbgpNeighbor(v InlineResponse200296EbgpNeighbor)`

SetEbgpNeighbor sets EbgpNeighbor field to given value.

### HasEbgpNeighbor

`func (o *InlineResponse200296Peers) HasEbgpNeighbor() bool`

HasEbgpNeighbor returns a boolean if a field has been set.

### GetEcmpUplinkConfigs

`func (o *InlineResponse200296Peers) GetEcmpUplinkConfigs() []InlineResponse200296EcmpUplinkConfigs`

GetEcmpUplinkConfigs returns the EcmpUplinkConfigs field if non-nil, zero value otherwise.

### GetEcmpUplinkConfigsOk

`func (o *InlineResponse200296Peers) GetEcmpUplinkConfigsOk() (*[]InlineResponse200296EcmpUplinkConfigs, bool)`

GetEcmpUplinkConfigsOk returns a tuple with the EcmpUplinkConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcmpUplinkConfigs

`func (o *InlineResponse200296Peers) SetEcmpUplinkConfigs(v []InlineResponse200296EcmpUplinkConfigs)`

SetEcmpUplinkConfigs sets EcmpUplinkConfigs field to given value.

### HasEcmpUplinkConfigs

`func (o *InlineResponse200296Peers) HasEcmpUplinkConfigs() bool`

HasEcmpUplinkConfigs returns a boolean if a field has been set.

### GetPriorityInGroup

`func (o *InlineResponse200296Peers) GetPriorityInGroup() int32`

GetPriorityInGroup returns the PriorityInGroup field if non-nil, zero value otherwise.

### GetPriorityInGroupOk

`func (o *InlineResponse200296Peers) GetPriorityInGroupOk() (*int32, bool)`

GetPriorityInGroupOk returns a tuple with the PriorityInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityInGroup

`func (o *InlineResponse200296Peers) SetPriorityInGroup(v int32)`

SetPriorityInGroup sets PriorityInGroup field to given value.

### HasPriorityInGroup

`func (o *InlineResponse200296Peers) HasPriorityInGroup() bool`

HasPriorityInGroup returns a boolean if a field has been set.

### GetGroup

`func (o *InlineResponse200296Peers) GetGroup() InlineResponse200296Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse200296Peers) GetGroupOk() (*InlineResponse200296Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse200296Peers) SetGroup(v InlineResponse200296Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse200296Peers) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


