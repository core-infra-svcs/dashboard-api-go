# InlineResponse20079Peers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the VPN peer | [optional] 
**PublicIp** | Pointer to **string** | [optional] The public IP of the VPN peer | [optional] 
**RemoteId** | Pointer to **string** | [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN. | [optional] 
**LocalId** | Pointer to **string** | [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to. | [optional] 
**Secret** | Pointer to **string** | The shared secret with the VPN peer | [optional] 
**PrivateSubnets** | Pointer to **[]string** | The list of the private subnets of the VPN peer | [optional] 
**IpsecPolicies** | Pointer to [**InlineResponse20079IpsecPolicies**](InlineResponse20079IpsecPolicies.md) |  | [optional] 
**IpsecPoliciesPreset** | Pointer to **string** | One of the following available presets: &#39;default&#39;, &#39;aws&#39;, &#39;azure&#39;. If this is provided, the &#39;ipsecPolicies&#39; parameter is ignored. | [optional] 
**IkeVersion** | Pointer to **string** | [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to &#39;1&#39; when omitted. | [optional] [default to "1"]
**NetworkTags** | Pointer to **[]string** | A list of network tags that will connect with this peer. Use [&#39;all&#39;] for all networks. Use [&#39;none&#39;] for no networks. If not included, the default is [&#39;all&#39;]. | [optional] 

## Methods

### NewInlineResponse20079Peers

`func NewInlineResponse20079Peers() *InlineResponse20079Peers`

NewInlineResponse20079Peers instantiates a new InlineResponse20079Peers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20079PeersWithDefaults

`func NewInlineResponse20079PeersWithDefaults() *InlineResponse20079Peers`

NewInlineResponse20079PeersWithDefaults instantiates a new InlineResponse20079Peers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse20079Peers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20079Peers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20079Peers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20079Peers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *InlineResponse20079Peers) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InlineResponse20079Peers) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InlineResponse20079Peers) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InlineResponse20079Peers) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRemoteId

`func (o *InlineResponse20079Peers) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *InlineResponse20079Peers) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *InlineResponse20079Peers) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *InlineResponse20079Peers) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetLocalId

`func (o *InlineResponse20079Peers) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *InlineResponse20079Peers) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *InlineResponse20079Peers) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *InlineResponse20079Peers) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetSecret

`func (o *InlineResponse20079Peers) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *InlineResponse20079Peers) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *InlineResponse20079Peers) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *InlineResponse20079Peers) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *InlineResponse20079Peers) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *InlineResponse20079Peers) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *InlineResponse20079Peers) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *InlineResponse20079Peers) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetIpsecPolicies

`func (o *InlineResponse20079Peers) GetIpsecPolicies() InlineResponse20079IpsecPolicies`

GetIpsecPolicies returns the IpsecPolicies field if non-nil, zero value otherwise.

### GetIpsecPoliciesOk

`func (o *InlineResponse20079Peers) GetIpsecPoliciesOk() (*InlineResponse20079IpsecPolicies, bool)`

GetIpsecPoliciesOk returns a tuple with the IpsecPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicies

`func (o *InlineResponse20079Peers) SetIpsecPolicies(v InlineResponse20079IpsecPolicies)`

SetIpsecPolicies sets IpsecPolicies field to given value.

### HasIpsecPolicies

`func (o *InlineResponse20079Peers) HasIpsecPolicies() bool`

HasIpsecPolicies returns a boolean if a field has been set.

### GetIpsecPoliciesPreset

`func (o *InlineResponse20079Peers) GetIpsecPoliciesPreset() string`

GetIpsecPoliciesPreset returns the IpsecPoliciesPreset field if non-nil, zero value otherwise.

### GetIpsecPoliciesPresetOk

`func (o *InlineResponse20079Peers) GetIpsecPoliciesPresetOk() (*string, bool)`

GetIpsecPoliciesPresetOk returns a tuple with the IpsecPoliciesPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPoliciesPreset

`func (o *InlineResponse20079Peers) SetIpsecPoliciesPreset(v string)`

SetIpsecPoliciesPreset sets IpsecPoliciesPreset field to given value.

### HasIpsecPoliciesPreset

`func (o *InlineResponse20079Peers) HasIpsecPoliciesPreset() bool`

HasIpsecPoliciesPreset returns a boolean if a field has been set.

### GetIkeVersion

`func (o *InlineResponse20079Peers) GetIkeVersion() string`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *InlineResponse20079Peers) GetIkeVersionOk() (*string, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *InlineResponse20079Peers) SetIkeVersion(v string)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *InlineResponse20079Peers) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetNetworkTags

`func (o *InlineResponse20079Peers) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InlineResponse20079Peers) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InlineResponse20079Peers) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *InlineResponse20079Peers) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


