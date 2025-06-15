# InlineResponse200237Peers

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
**IpsecPolicies** | Pointer to [**InlineResponse200237IpsecPolicies**](InlineResponse200237IpsecPolicies.md) |  | [optional] 
**SlaPolicy** | Pointer to [**InlineResponse200237SlaPolicy**](InlineResponse200237SlaPolicy.md) |  | [optional] 
**IpsecPoliciesPreset** | Pointer to **string** | One of the following available presets: &#39;default&#39;, &#39;aws&#39;, &#39;azure&#39;, &#39;umbrella&#39;, &#39;zscaler&#39;. If this is provided, the &#39;ipsecPolicies&#39; parameter is ignored. | [optional] 
**IkeVersion** | Pointer to **string** | [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to &#39;1&#39; when omitted. | [optional] [default to "1"]
**NetworkTags** | Pointer to **[]string** | A list of network tags that will connect with this peer. Use [&#39;all&#39;] for all networks. Use [&#39;none&#39;] for no networks. If not included, the default is [&#39;all&#39;]. | [optional] 
**Network** | Pointer to [**InlineResponse200237Network**](InlineResponse200237Network.md) |  | [optional] 
**IsRouteBased** | Pointer to **bool** | [optional] If true, the VPN peer is route-based. If not included, the default is false. Supported only for MX 19.1 and above. | [optional] 
**EbgpNeighbor** | Pointer to [**InlineResponse200237EbgpNeighbor**](InlineResponse200237EbgpNeighbor.md) |  | [optional] 
**PriorityInGroup** | Pointer to **int32** | [optional] Represents the order of peer inside a group. | [optional] 
**Group** | Pointer to [**InlineResponse200237Group**](InlineResponse200237Group.md) |  | [optional] 

## Methods

### NewInlineResponse200237Peers

`func NewInlineResponse200237Peers() *InlineResponse200237Peers`

NewInlineResponse200237Peers instantiates a new InlineResponse200237Peers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200237PeersWithDefaults

`func NewInlineResponse200237PeersWithDefaults() *InlineResponse200237Peers`

NewInlineResponse200237PeersWithDefaults instantiates a new InlineResponse200237Peers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerId

`func (o *InlineResponse200237Peers) GetPeerId() string`

GetPeerId returns the PeerId field if non-nil, zero value otherwise.

### GetPeerIdOk

`func (o *InlineResponse200237Peers) GetPeerIdOk() (*string, bool)`

GetPeerIdOk returns a tuple with the PeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerId

`func (o *InlineResponse200237Peers) SetPeerId(v string)`

SetPeerId sets PeerId field to given value.

### HasPeerId

`func (o *InlineResponse200237Peers) HasPeerId() bool`

HasPeerId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200237Peers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200237Peers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200237Peers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200237Peers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *InlineResponse200237Peers) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InlineResponse200237Peers) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InlineResponse200237Peers) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InlineResponse200237Peers) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRemoteId

`func (o *InlineResponse200237Peers) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *InlineResponse200237Peers) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *InlineResponse200237Peers) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *InlineResponse200237Peers) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetLocalId

`func (o *InlineResponse200237Peers) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *InlineResponse200237Peers) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *InlineResponse200237Peers) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *InlineResponse200237Peers) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetSecret

`func (o *InlineResponse200237Peers) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *InlineResponse200237Peers) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *InlineResponse200237Peers) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *InlineResponse200237Peers) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *InlineResponse200237Peers) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *InlineResponse200237Peers) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *InlineResponse200237Peers) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *InlineResponse200237Peers) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetIpsecPolicies

`func (o *InlineResponse200237Peers) GetIpsecPolicies() InlineResponse200237IpsecPolicies`

GetIpsecPolicies returns the IpsecPolicies field if non-nil, zero value otherwise.

### GetIpsecPoliciesOk

`func (o *InlineResponse200237Peers) GetIpsecPoliciesOk() (*InlineResponse200237IpsecPolicies, bool)`

GetIpsecPoliciesOk returns a tuple with the IpsecPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicies

`func (o *InlineResponse200237Peers) SetIpsecPolicies(v InlineResponse200237IpsecPolicies)`

SetIpsecPolicies sets IpsecPolicies field to given value.

### HasIpsecPolicies

`func (o *InlineResponse200237Peers) HasIpsecPolicies() bool`

HasIpsecPolicies returns a boolean if a field has been set.

### GetSlaPolicy

`func (o *InlineResponse200237Peers) GetSlaPolicy() InlineResponse200237SlaPolicy`

GetSlaPolicy returns the SlaPolicy field if non-nil, zero value otherwise.

### GetSlaPolicyOk

`func (o *InlineResponse200237Peers) GetSlaPolicyOk() (*InlineResponse200237SlaPolicy, bool)`

GetSlaPolicyOk returns a tuple with the SlaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaPolicy

`func (o *InlineResponse200237Peers) SetSlaPolicy(v InlineResponse200237SlaPolicy)`

SetSlaPolicy sets SlaPolicy field to given value.

### HasSlaPolicy

`func (o *InlineResponse200237Peers) HasSlaPolicy() bool`

HasSlaPolicy returns a boolean if a field has been set.

### GetIpsecPoliciesPreset

`func (o *InlineResponse200237Peers) GetIpsecPoliciesPreset() string`

GetIpsecPoliciesPreset returns the IpsecPoliciesPreset field if non-nil, zero value otherwise.

### GetIpsecPoliciesPresetOk

`func (o *InlineResponse200237Peers) GetIpsecPoliciesPresetOk() (*string, bool)`

GetIpsecPoliciesPresetOk returns a tuple with the IpsecPoliciesPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPoliciesPreset

`func (o *InlineResponse200237Peers) SetIpsecPoliciesPreset(v string)`

SetIpsecPoliciesPreset sets IpsecPoliciesPreset field to given value.

### HasIpsecPoliciesPreset

`func (o *InlineResponse200237Peers) HasIpsecPoliciesPreset() bool`

HasIpsecPoliciesPreset returns a boolean if a field has been set.

### GetIkeVersion

`func (o *InlineResponse200237Peers) GetIkeVersion() string`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *InlineResponse200237Peers) GetIkeVersionOk() (*string, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *InlineResponse200237Peers) SetIkeVersion(v string)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *InlineResponse200237Peers) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetNetworkTags

`func (o *InlineResponse200237Peers) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InlineResponse200237Peers) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InlineResponse200237Peers) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *InlineResponse200237Peers) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200237Peers) GetNetwork() InlineResponse200237Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200237Peers) GetNetworkOk() (*InlineResponse200237Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200237Peers) SetNetwork(v InlineResponse200237Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200237Peers) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIsRouteBased

`func (o *InlineResponse200237Peers) GetIsRouteBased() bool`

GetIsRouteBased returns the IsRouteBased field if non-nil, zero value otherwise.

### GetIsRouteBasedOk

`func (o *InlineResponse200237Peers) GetIsRouteBasedOk() (*bool, bool)`

GetIsRouteBasedOk returns a tuple with the IsRouteBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRouteBased

`func (o *InlineResponse200237Peers) SetIsRouteBased(v bool)`

SetIsRouteBased sets IsRouteBased field to given value.

### HasIsRouteBased

`func (o *InlineResponse200237Peers) HasIsRouteBased() bool`

HasIsRouteBased returns a boolean if a field has been set.

### GetEbgpNeighbor

`func (o *InlineResponse200237Peers) GetEbgpNeighbor() InlineResponse200237EbgpNeighbor`

GetEbgpNeighbor returns the EbgpNeighbor field if non-nil, zero value otherwise.

### GetEbgpNeighborOk

`func (o *InlineResponse200237Peers) GetEbgpNeighborOk() (*InlineResponse200237EbgpNeighbor, bool)`

GetEbgpNeighborOk returns a tuple with the EbgpNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpNeighbor

`func (o *InlineResponse200237Peers) SetEbgpNeighbor(v InlineResponse200237EbgpNeighbor)`

SetEbgpNeighbor sets EbgpNeighbor field to given value.

### HasEbgpNeighbor

`func (o *InlineResponse200237Peers) HasEbgpNeighbor() bool`

HasEbgpNeighbor returns a boolean if a field has been set.

### GetPriorityInGroup

`func (o *InlineResponse200237Peers) GetPriorityInGroup() int32`

GetPriorityInGroup returns the PriorityInGroup field if non-nil, zero value otherwise.

### GetPriorityInGroupOk

`func (o *InlineResponse200237Peers) GetPriorityInGroupOk() (*int32, bool)`

GetPriorityInGroupOk returns a tuple with the PriorityInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityInGroup

`func (o *InlineResponse200237Peers) SetPriorityInGroup(v int32)`

SetPriorityInGroup sets PriorityInGroup field to given value.

### HasPriorityInGroup

`func (o *InlineResponse200237Peers) HasPriorityInGroup() bool`

HasPriorityInGroup returns a boolean if a field has been set.

### GetGroup

`func (o *InlineResponse200237Peers) GetGroup() InlineResponse200237Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse200237Peers) GetGroupOk() (*InlineResponse200237Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse200237Peers) SetGroup(v InlineResponse200237Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse200237Peers) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


