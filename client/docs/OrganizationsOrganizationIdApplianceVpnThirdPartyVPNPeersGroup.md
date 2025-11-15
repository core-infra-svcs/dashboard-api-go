# OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | [optional] Represents the ordering of primary and backup tunnels group. primary and backup tunnels are grouped by this number. If you submit a request with the numbers [1, 9, 999], these numbers will be automatically adjusted to a sequential order starting from 1. So, they will be changed to [1, 2, 3] to reflect their positions in the sequence. | [optional] 
**Failover** | Pointer to [**InlineResponse200242GroupFailover**](InlineResponse200242GroupFailover.md) |  | [optional] 
**ActiveActiveTunnel** | Pointer to **bool** | [optional] Both primary and backup tunnels are active. | [optional] 

## Methods

### NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup

`func NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup() *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup`

NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup instantiates a new OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroupWithDefaults

`func NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroupWithDefaults() *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup`

NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroupWithDefaults instantiates a new OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetFailover

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) GetFailover() InlineResponse200242GroupFailover`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) GetFailoverOk() (*InlineResponse200242GroupFailover, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) SetFailover(v InlineResponse200242GroupFailover)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetActiveActiveTunnel

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) GetActiveActiveTunnel() bool`

GetActiveActiveTunnel returns the ActiveActiveTunnel field if non-nil, zero value otherwise.

### GetActiveActiveTunnelOk

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) GetActiveActiveTunnelOk() (*bool, bool)`

GetActiveActiveTunnelOk returns a tuple with the ActiveActiveTunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveActiveTunnel

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) SetActiveActiveTunnel(v bool)`

SetActiveActiveTunnel sets ActiveActiveTunnel field to given value.

### HasActiveActiveTunnel

`func (o *OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersGroup) HasActiveActiveTunnel() bool`

HasActiveActiveTunnel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


