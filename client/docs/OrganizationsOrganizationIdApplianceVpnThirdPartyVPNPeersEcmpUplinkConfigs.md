# OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the ECMP uplink configuration | [optional] 
**Wan** | Pointer to **string** | The WAN uplink associated with this ECMP configuration. | [optional] 
**PrivateSubnets** | Pointer to **[]string** | The list of the private subnets of the VPN peer | [optional] 
**EbgpNeighbor** | Pointer to [**OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEbgpNeighbor1**](OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEbgpNeighbor1.md) |  | [optional] 

## Methods

### NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs

`func NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs() *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs`

NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs instantiates a new OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigsWithDefaults

`func NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigsWithDefaults() *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs`

NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigsWithDefaults instantiates a new OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWan

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) GetWan() string`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) GetWanOk() (*string, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) SetWan(v string)`

SetWan sets Wan field to given value.

### HasWan

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) HasWan() bool`

HasWan returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetEbgpNeighbor

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) GetEbgpNeighbor() OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEbgpNeighbor1`

GetEbgpNeighbor returns the EbgpNeighbor field if non-nil, zero value otherwise.

### GetEbgpNeighborOk

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) GetEbgpNeighborOk() (*OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEbgpNeighbor1, bool)`

GetEbgpNeighborOk returns a tuple with the EbgpNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpNeighbor

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) SetEbgpNeighbor(v OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEbgpNeighbor1)`

SetEbgpNeighbor sets EbgpNeighbor field to given value.

### HasEbgpNeighbor

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersEcmpUplinkConfigs) HasEbgpNeighbor() bool`

HasEbgpNeighbor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


