# InlineResponse200241Peers

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
**IpsecPolicies** | Pointer to [**InlineResponse200241IpsecPolicies**](InlineResponse200241IpsecPolicies.md) |  | [optional] 
**SlaPolicy** | Pointer to [**InlineResponse200241SlaPolicy**](InlineResponse200241SlaPolicy.md) |  | [optional] 
**IpsecPoliciesPreset** | Pointer to **string** | One of the following available presets: &#39;default&#39;, &#39;aws&#39;, &#39;azure&#39;, &#39;umbrella&#39;, &#39;zscaler&#39;. If this is provided, the &#39;ipsecPolicies&#39; parameter is ignored. | [optional] 
**IkeVersion** | Pointer to **string** | [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to &#39;1&#39; when omitted. | [optional] [default to "1"]
**NetworkTags** | Pointer to **[]string** | A list of network tags that will connect with this peer. Use [&#39;all&#39;] for all networks. Use [&#39;none&#39;] for no networks. If not included, the default is [&#39;all&#39;]. | [optional] 
**Network** | Pointer to [**InlineResponse200241Network**](InlineResponse200241Network.md) |  | [optional] 
**IsRouteBased** | Pointer to **bool** | [optional] If true, the VPN peer is route-based. If not included, the default is false. Supported only for MX 19.1 and above. | [optional] 
**EbgpNeighbor** | Pointer to [**InlineResponse200241EbgpNeighbor**](InlineResponse200241EbgpNeighbor.md) |  | [optional] 
**PriorityInGroup** | Pointer to **int32** | [optional] Represents the order of peer inside a group. | [optional] 
**Group** | Pointer to [**InlineResponse200241Group**](InlineResponse200241Group.md) |  | [optional] 

## Methods

### NewInlineResponse200241Peers

`func NewInlineResponse200241Peers() *InlineResponse200241Peers`

NewInlineResponse200241Peers instantiates a new InlineResponse200241Peers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200241PeersWithDefaults

`func NewInlineResponse200241PeersWithDefaults() *InlineResponse200241Peers`

NewInlineResponse200241PeersWithDefaults instantiates a new InlineResponse200241Peers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerId

`func (o *InlineResponse200241Peers) GetPeerId() string`

GetPeerId returns the PeerId field if non-nil, zero value otherwise.

### GetPeerIdOk

`func (o *InlineResponse200241Peers) GetPeerIdOk() (*string, bool)`

GetPeerIdOk returns a tuple with the PeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerId

`func (o *InlineResponse200241Peers) SetPeerId(v string)`

SetPeerId sets PeerId field to given value.

### HasPeerId

`func (o *InlineResponse200241Peers) HasPeerId() bool`

HasPeerId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200241Peers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200241Peers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200241Peers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200241Peers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *InlineResponse200241Peers) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InlineResponse200241Peers) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InlineResponse200241Peers) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InlineResponse200241Peers) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRemoteId

`func (o *InlineResponse200241Peers) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *InlineResponse200241Peers) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *InlineResponse200241Peers) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *InlineResponse200241Peers) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetLocalId

`func (o *InlineResponse200241Peers) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *InlineResponse200241Peers) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *InlineResponse200241Peers) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *InlineResponse200241Peers) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetSecret

`func (o *InlineResponse200241Peers) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *InlineResponse200241Peers) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *InlineResponse200241Peers) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *InlineResponse200241Peers) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *InlineResponse200241Peers) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *InlineResponse200241Peers) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *InlineResponse200241Peers) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *InlineResponse200241Peers) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetIpsecPolicies

`func (o *InlineResponse200241Peers) GetIpsecPolicies() InlineResponse200241IpsecPolicies`

GetIpsecPolicies returns the IpsecPolicies field if non-nil, zero value otherwise.

### GetIpsecPoliciesOk

`func (o *InlineResponse200241Peers) GetIpsecPoliciesOk() (*InlineResponse200241IpsecPolicies, bool)`

GetIpsecPoliciesOk returns a tuple with the IpsecPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicies

`func (o *InlineResponse200241Peers) SetIpsecPolicies(v InlineResponse200241IpsecPolicies)`

SetIpsecPolicies sets IpsecPolicies field to given value.

### HasIpsecPolicies

`func (o *InlineResponse200241Peers) HasIpsecPolicies() bool`

HasIpsecPolicies returns a boolean if a field has been set.

### GetSlaPolicy

`func (o *InlineResponse200241Peers) GetSlaPolicy() InlineResponse200241SlaPolicy`

GetSlaPolicy returns the SlaPolicy field if non-nil, zero value otherwise.

### GetSlaPolicyOk

`func (o *InlineResponse200241Peers) GetSlaPolicyOk() (*InlineResponse200241SlaPolicy, bool)`

GetSlaPolicyOk returns a tuple with the SlaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaPolicy

`func (o *InlineResponse200241Peers) SetSlaPolicy(v InlineResponse200241SlaPolicy)`

SetSlaPolicy sets SlaPolicy field to given value.

### HasSlaPolicy

`func (o *InlineResponse200241Peers) HasSlaPolicy() bool`

HasSlaPolicy returns a boolean if a field has been set.

### GetIpsecPoliciesPreset

`func (o *InlineResponse200241Peers) GetIpsecPoliciesPreset() string`

GetIpsecPoliciesPreset returns the IpsecPoliciesPreset field if non-nil, zero value otherwise.

### GetIpsecPoliciesPresetOk

`func (o *InlineResponse200241Peers) GetIpsecPoliciesPresetOk() (*string, bool)`

GetIpsecPoliciesPresetOk returns a tuple with the IpsecPoliciesPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPoliciesPreset

`func (o *InlineResponse200241Peers) SetIpsecPoliciesPreset(v string)`

SetIpsecPoliciesPreset sets IpsecPoliciesPreset field to given value.

### HasIpsecPoliciesPreset

`func (o *InlineResponse200241Peers) HasIpsecPoliciesPreset() bool`

HasIpsecPoliciesPreset returns a boolean if a field has been set.

### GetIkeVersion

`func (o *InlineResponse200241Peers) GetIkeVersion() string`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *InlineResponse200241Peers) GetIkeVersionOk() (*string, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *InlineResponse200241Peers) SetIkeVersion(v string)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *InlineResponse200241Peers) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetNetworkTags

`func (o *InlineResponse200241Peers) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InlineResponse200241Peers) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InlineResponse200241Peers) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *InlineResponse200241Peers) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200241Peers) GetNetwork() InlineResponse200241Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200241Peers) GetNetworkOk() (*InlineResponse200241Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200241Peers) SetNetwork(v InlineResponse200241Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200241Peers) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIsRouteBased

`func (o *InlineResponse200241Peers) GetIsRouteBased() bool`

GetIsRouteBased returns the IsRouteBased field if non-nil, zero value otherwise.

### GetIsRouteBasedOk

`func (o *InlineResponse200241Peers) GetIsRouteBasedOk() (*bool, bool)`

GetIsRouteBasedOk returns a tuple with the IsRouteBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRouteBased

`func (o *InlineResponse200241Peers) SetIsRouteBased(v bool)`

SetIsRouteBased sets IsRouteBased field to given value.

### HasIsRouteBased

`func (o *InlineResponse200241Peers) HasIsRouteBased() bool`

HasIsRouteBased returns a boolean if a field has been set.

### GetEbgpNeighbor

`func (o *InlineResponse200241Peers) GetEbgpNeighbor() InlineResponse200241EbgpNeighbor`

GetEbgpNeighbor returns the EbgpNeighbor field if non-nil, zero value otherwise.

### GetEbgpNeighborOk

`func (o *InlineResponse200241Peers) GetEbgpNeighborOk() (*InlineResponse200241EbgpNeighbor, bool)`

GetEbgpNeighborOk returns a tuple with the EbgpNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpNeighbor

`func (o *InlineResponse200241Peers) SetEbgpNeighbor(v InlineResponse200241EbgpNeighbor)`

SetEbgpNeighbor sets EbgpNeighbor field to given value.

### HasEbgpNeighbor

`func (o *InlineResponse200241Peers) HasEbgpNeighbor() bool`

HasEbgpNeighbor returns a boolean if a field has been set.

### GetPriorityInGroup

`func (o *InlineResponse200241Peers) GetPriorityInGroup() int32`

GetPriorityInGroup returns the PriorityInGroup field if non-nil, zero value otherwise.

### GetPriorityInGroupOk

`func (o *InlineResponse200241Peers) GetPriorityInGroupOk() (*int32, bool)`

GetPriorityInGroupOk returns a tuple with the PriorityInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityInGroup

`func (o *InlineResponse200241Peers) SetPriorityInGroup(v int32)`

SetPriorityInGroup sets PriorityInGroup field to given value.

### HasPriorityInGroup

`func (o *InlineResponse200241Peers) HasPriorityInGroup() bool`

HasPriorityInGroup returns a boolean if a field has been set.

### GetGroup

`func (o *InlineResponse200241Peers) GetGroup() InlineResponse200241Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse200241Peers) GetGroupOk() (*InlineResponse200241Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse200241Peers) SetGroup(v InlineResponse200241Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse200241Peers) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


