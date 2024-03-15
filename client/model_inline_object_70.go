/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject70 struct for InlineObject70
type InlineObject70 struct {
	// The name of the VLAN
	Name *string `json:"name,omitempty"`
	// The subnet of the VLAN
	Subnet *string `json:"subnet,omitempty"`
	// The local IP of the appliance on the VLAN
	ApplianceIp *string `json:"applianceIp,omitempty"`
	// The id of the desired group policy to apply to the VLAN
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	// The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN
	VpnNatSubnet *string `json:"vpnNatSubnet,omitempty"`
	// The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'
	DhcpHandling *string `json:"dhcpHandling,omitempty"`
	// The IPs of the DHCP servers that DHCP requests should be relayed to
	DhcpRelayServerIps []string `json:"dhcpRelayServerIps,omitempty"`
	// The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'
	DhcpLeaseTime *string `json:"dhcpLeaseTime,omitempty"`
	// Use DHCP boot options specified in other properties
	DhcpBootOptionsEnabled *bool `json:"dhcpBootOptionsEnabled,omitempty"`
	// DHCP boot option to direct boot clients to the server to load the boot file from
	DhcpBootNextServer *string `json:"dhcpBootNextServer,omitempty"`
	// DHCP boot option for boot filename
	DhcpBootFilename *string `json:"dhcpBootFilename,omitempty"`
	// The DHCP fixed IP assignments on the VLAN. This should be an object that contains mappings from MAC addresses to objects that themselves each contain \"ip\" and \"name\" string fields. See the sample request/response for more details.
	FixedIpAssignments map[string]interface{} `json:"fixedIpAssignments,omitempty"`
	// The DHCP reserved IP ranges on the VLAN
	ReservedIpRanges []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges `json:"reservedIpRanges,omitempty"`
	// The DNS nameservers used for DHCP responses, either \"upstream_dns\", \"google_dns\", \"opendns\", or a newline seperated string of IP addresses or domain names
	DnsNameservers *string `json:"dnsNameservers,omitempty"`
	// The list of DHCP options that will be included in DHCP responses. Each object in the list should have \"code\", \"type\", and \"value\" properties.
	DhcpOptions []NetworksNetworkIdApplianceVlansDhcpOptions `json:"dhcpOptions,omitempty"`
	// Type of subnetting of the VLAN. Applicable only for template network.
	TemplateVlanType *string `json:"templateVlanType,omitempty"`
	// CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	Cidr *string `json:"cidr,omitempty"`
	// Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Mask *int32 `json:"mask,omitempty"`
	Ipv6 *NetworksNetworkIdApplianceSingleLanIpv6 `json:"ipv6,omitempty"`
	MandatoryDhcp *NetworksNetworkIdApplianceVlansMandatoryDhcp `json:"mandatoryDhcp,omitempty"`
}

// NewInlineObject70 instantiates a new InlineObject70 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject70() *InlineObject70 {
	this := InlineObject70{}
	return &this
}

// NewInlineObject70WithDefaults instantiates a new InlineObject70 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject70WithDefaults() *InlineObject70 {
	this := InlineObject70{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject70) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject70) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject70) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineObject70) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject70) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineObject70) SetSubnet(v string) {
	o.Subnet = &v
}

// GetApplianceIp returns the ApplianceIp field value if set, zero value otherwise.
func (o *InlineObject70) GetApplianceIp() string {
	if o == nil || isNil(o.ApplianceIp) {
		var ret string
		return ret
	}
	return *o.ApplianceIp
}

// GetApplianceIpOk returns a tuple with the ApplianceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetApplianceIpOk() (*string, bool) {
	if o == nil || isNil(o.ApplianceIp) {
    return nil, false
	}
	return o.ApplianceIp, true
}

// HasApplianceIp returns a boolean if a field has been set.
func (o *InlineObject70) HasApplianceIp() bool {
	if o != nil && !isNil(o.ApplianceIp) {
		return true
	}

	return false
}

// SetApplianceIp gets a reference to the given string and assigns it to the ApplianceIp field.
func (o *InlineObject70) SetApplianceIp(v string) {
	o.ApplianceIp = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *InlineObject70) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *InlineObject70) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *InlineObject70) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetVpnNatSubnet returns the VpnNatSubnet field value if set, zero value otherwise.
func (o *InlineObject70) GetVpnNatSubnet() string {
	if o == nil || isNil(o.VpnNatSubnet) {
		var ret string
		return ret
	}
	return *o.VpnNatSubnet
}

// GetVpnNatSubnetOk returns a tuple with the VpnNatSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetVpnNatSubnetOk() (*string, bool) {
	if o == nil || isNil(o.VpnNatSubnet) {
    return nil, false
	}
	return o.VpnNatSubnet, true
}

// HasVpnNatSubnet returns a boolean if a field has been set.
func (o *InlineObject70) HasVpnNatSubnet() bool {
	if o != nil && !isNil(o.VpnNatSubnet) {
		return true
	}

	return false
}

// SetVpnNatSubnet gets a reference to the given string and assigns it to the VpnNatSubnet field.
func (o *InlineObject70) SetVpnNatSubnet(v string) {
	o.VpnNatSubnet = &v
}

// GetDhcpHandling returns the DhcpHandling field value if set, zero value otherwise.
func (o *InlineObject70) GetDhcpHandling() string {
	if o == nil || isNil(o.DhcpHandling) {
		var ret string
		return ret
	}
	return *o.DhcpHandling
}

// GetDhcpHandlingOk returns a tuple with the DhcpHandling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetDhcpHandlingOk() (*string, bool) {
	if o == nil || isNil(o.DhcpHandling) {
    return nil, false
	}
	return o.DhcpHandling, true
}

// HasDhcpHandling returns a boolean if a field has been set.
func (o *InlineObject70) HasDhcpHandling() bool {
	if o != nil && !isNil(o.DhcpHandling) {
		return true
	}

	return false
}

// SetDhcpHandling gets a reference to the given string and assigns it to the DhcpHandling field.
func (o *InlineObject70) SetDhcpHandling(v string) {
	o.DhcpHandling = &v
}

// GetDhcpRelayServerIps returns the DhcpRelayServerIps field value if set, zero value otherwise.
func (o *InlineObject70) GetDhcpRelayServerIps() []string {
	if o == nil || isNil(o.DhcpRelayServerIps) {
		var ret []string
		return ret
	}
	return o.DhcpRelayServerIps
}

// GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetDhcpRelayServerIpsOk() ([]string, bool) {
	if o == nil || isNil(o.DhcpRelayServerIps) {
    return nil, false
	}
	return o.DhcpRelayServerIps, true
}

// HasDhcpRelayServerIps returns a boolean if a field has been set.
func (o *InlineObject70) HasDhcpRelayServerIps() bool {
	if o != nil && !isNil(o.DhcpRelayServerIps) {
		return true
	}

	return false
}

// SetDhcpRelayServerIps gets a reference to the given []string and assigns it to the DhcpRelayServerIps field.
func (o *InlineObject70) SetDhcpRelayServerIps(v []string) {
	o.DhcpRelayServerIps = v
}

// GetDhcpLeaseTime returns the DhcpLeaseTime field value if set, zero value otherwise.
func (o *InlineObject70) GetDhcpLeaseTime() string {
	if o == nil || isNil(o.DhcpLeaseTime) {
		var ret string
		return ret
	}
	return *o.DhcpLeaseTime
}

// GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || isNil(o.DhcpLeaseTime) {
    return nil, false
	}
	return o.DhcpLeaseTime, true
}

// HasDhcpLeaseTime returns a boolean if a field has been set.
func (o *InlineObject70) HasDhcpLeaseTime() bool {
	if o != nil && !isNil(o.DhcpLeaseTime) {
		return true
	}

	return false
}

// SetDhcpLeaseTime gets a reference to the given string and assigns it to the DhcpLeaseTime field.
func (o *InlineObject70) SetDhcpLeaseTime(v string) {
	o.DhcpLeaseTime = &v
}

// GetDhcpBootOptionsEnabled returns the DhcpBootOptionsEnabled field value if set, zero value otherwise.
func (o *InlineObject70) GetDhcpBootOptionsEnabled() bool {
	if o == nil || isNil(o.DhcpBootOptionsEnabled) {
		var ret bool
		return ret
	}
	return *o.DhcpBootOptionsEnabled
}

// GetDhcpBootOptionsEnabledOk returns a tuple with the DhcpBootOptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetDhcpBootOptionsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.DhcpBootOptionsEnabled) {
    return nil, false
	}
	return o.DhcpBootOptionsEnabled, true
}

// HasDhcpBootOptionsEnabled returns a boolean if a field has been set.
func (o *InlineObject70) HasDhcpBootOptionsEnabled() bool {
	if o != nil && !isNil(o.DhcpBootOptionsEnabled) {
		return true
	}

	return false
}

// SetDhcpBootOptionsEnabled gets a reference to the given bool and assigns it to the DhcpBootOptionsEnabled field.
func (o *InlineObject70) SetDhcpBootOptionsEnabled(v bool) {
	o.DhcpBootOptionsEnabled = &v
}

// GetDhcpBootNextServer returns the DhcpBootNextServer field value if set, zero value otherwise.
func (o *InlineObject70) GetDhcpBootNextServer() string {
	if o == nil || isNil(o.DhcpBootNextServer) {
		var ret string
		return ret
	}
	return *o.DhcpBootNextServer
}

// GetDhcpBootNextServerOk returns a tuple with the DhcpBootNextServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetDhcpBootNextServerOk() (*string, bool) {
	if o == nil || isNil(o.DhcpBootNextServer) {
    return nil, false
	}
	return o.DhcpBootNextServer, true
}

// HasDhcpBootNextServer returns a boolean if a field has been set.
func (o *InlineObject70) HasDhcpBootNextServer() bool {
	if o != nil && !isNil(o.DhcpBootNextServer) {
		return true
	}

	return false
}

// SetDhcpBootNextServer gets a reference to the given string and assigns it to the DhcpBootNextServer field.
func (o *InlineObject70) SetDhcpBootNextServer(v string) {
	o.DhcpBootNextServer = &v
}

// GetDhcpBootFilename returns the DhcpBootFilename field value if set, zero value otherwise.
func (o *InlineObject70) GetDhcpBootFilename() string {
	if o == nil || isNil(o.DhcpBootFilename) {
		var ret string
		return ret
	}
	return *o.DhcpBootFilename
}

// GetDhcpBootFilenameOk returns a tuple with the DhcpBootFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetDhcpBootFilenameOk() (*string, bool) {
	if o == nil || isNil(o.DhcpBootFilename) {
    return nil, false
	}
	return o.DhcpBootFilename, true
}

// HasDhcpBootFilename returns a boolean if a field has been set.
func (o *InlineObject70) HasDhcpBootFilename() bool {
	if o != nil && !isNil(o.DhcpBootFilename) {
		return true
	}

	return false
}

// SetDhcpBootFilename gets a reference to the given string and assigns it to the DhcpBootFilename field.
func (o *InlineObject70) SetDhcpBootFilename(v string) {
	o.DhcpBootFilename = &v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *InlineObject70) GetFixedIpAssignments() map[string]interface{} {
	if o == nil || isNil(o.FixedIpAssignments) {
		var ret map[string]interface{}
		return ret
	}
	return o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetFixedIpAssignmentsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.FixedIpAssignments) {
    return map[string]interface{}{}, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *InlineObject70) HasFixedIpAssignments() bool {
	if o != nil && !isNil(o.FixedIpAssignments) {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given map[string]interface{} and assigns it to the FixedIpAssignments field.
func (o *InlineObject70) SetFixedIpAssignments(v map[string]interface{}) {
	o.FixedIpAssignments = v
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *InlineObject70) GetReservedIpRanges() []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges {
	if o == nil || isNil(o.ReservedIpRanges) {
		var ret []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges
		return ret
	}
	return o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetReservedIpRangesOk() ([]NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges, bool) {
	if o == nil || isNil(o.ReservedIpRanges) {
    return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *InlineObject70) HasReservedIpRanges() bool {
	if o != nil && !isNil(o.ReservedIpRanges) {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges and assigns it to the ReservedIpRanges field.
func (o *InlineObject70) SetReservedIpRanges(v []NetworksNetworkIdApplianceStaticRoutesStaticRouteIdReservedIpRanges) {
	o.ReservedIpRanges = v
}

// GetDnsNameservers returns the DnsNameservers field value if set, zero value otherwise.
func (o *InlineObject70) GetDnsNameservers() string {
	if o == nil || isNil(o.DnsNameservers) {
		var ret string
		return ret
	}
	return *o.DnsNameservers
}

// GetDnsNameserversOk returns a tuple with the DnsNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetDnsNameserversOk() (*string, bool) {
	if o == nil || isNil(o.DnsNameservers) {
    return nil, false
	}
	return o.DnsNameservers, true
}

// HasDnsNameservers returns a boolean if a field has been set.
func (o *InlineObject70) HasDnsNameservers() bool {
	if o != nil && !isNil(o.DnsNameservers) {
		return true
	}

	return false
}

// SetDnsNameservers gets a reference to the given string and assigns it to the DnsNameservers field.
func (o *InlineObject70) SetDnsNameservers(v string) {
	o.DnsNameservers = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *InlineObject70) GetDhcpOptions() []NetworksNetworkIdApplianceVlansDhcpOptions {
	if o == nil || isNil(o.DhcpOptions) {
		var ret []NetworksNetworkIdApplianceVlansDhcpOptions
		return ret
	}
	return o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetDhcpOptionsOk() ([]NetworksNetworkIdApplianceVlansDhcpOptions, bool) {
	if o == nil || isNil(o.DhcpOptions) {
    return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *InlineObject70) HasDhcpOptions() bool {
	if o != nil && !isNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given []NetworksNetworkIdApplianceVlansDhcpOptions and assigns it to the DhcpOptions field.
func (o *InlineObject70) SetDhcpOptions(v []NetworksNetworkIdApplianceVlansDhcpOptions) {
	o.DhcpOptions = v
}

// GetTemplateVlanType returns the TemplateVlanType field value if set, zero value otherwise.
func (o *InlineObject70) GetTemplateVlanType() string {
	if o == nil || isNil(o.TemplateVlanType) {
		var ret string
		return ret
	}
	return *o.TemplateVlanType
}

// GetTemplateVlanTypeOk returns a tuple with the TemplateVlanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetTemplateVlanTypeOk() (*string, bool) {
	if o == nil || isNil(o.TemplateVlanType) {
    return nil, false
	}
	return o.TemplateVlanType, true
}

// HasTemplateVlanType returns a boolean if a field has been set.
func (o *InlineObject70) HasTemplateVlanType() bool {
	if o != nil && !isNil(o.TemplateVlanType) {
		return true
	}

	return false
}

// SetTemplateVlanType gets a reference to the given string and assigns it to the TemplateVlanType field.
func (o *InlineObject70) SetTemplateVlanType(v string) {
	o.TemplateVlanType = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineObject70) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineObject70) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineObject70) SetCidr(v string) {
	o.Cidr = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *InlineObject70) GetMask() int32 {
	if o == nil || isNil(o.Mask) {
		var ret int32
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetMaskOk() (*int32, bool) {
	if o == nil || isNil(o.Mask) {
    return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *InlineObject70) HasMask() bool {
	if o != nil && !isNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int32 and assigns it to the Mask field.
func (o *InlineObject70) SetMask(v int32) {
	o.Mask = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *InlineObject70) GetIpv6() NetworksNetworkIdApplianceSingleLanIpv6 {
	if o == nil || isNil(o.Ipv6) {
		var ret NetworksNetworkIdApplianceSingleLanIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetIpv6Ok() (*NetworksNetworkIdApplianceSingleLanIpv6, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *InlineObject70) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given NetworksNetworkIdApplianceSingleLanIpv6 and assigns it to the Ipv6 field.
func (o *InlineObject70) SetIpv6(v NetworksNetworkIdApplianceSingleLanIpv6) {
	o.Ipv6 = &v
}

// GetMandatoryDhcp returns the MandatoryDhcp field value if set, zero value otherwise.
func (o *InlineObject70) GetMandatoryDhcp() NetworksNetworkIdApplianceVlansMandatoryDhcp {
	if o == nil || isNil(o.MandatoryDhcp) {
		var ret NetworksNetworkIdApplianceVlansMandatoryDhcp
		return ret
	}
	return *o.MandatoryDhcp
}

// GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetMandatoryDhcpOk() (*NetworksNetworkIdApplianceVlansMandatoryDhcp, bool) {
	if o == nil || isNil(o.MandatoryDhcp) {
    return nil, false
	}
	return o.MandatoryDhcp, true
}

// HasMandatoryDhcp returns a boolean if a field has been set.
func (o *InlineObject70) HasMandatoryDhcp() bool {
	if o != nil && !isNil(o.MandatoryDhcp) {
		return true
	}

	return false
}

// SetMandatoryDhcp gets a reference to the given NetworksNetworkIdApplianceVlansMandatoryDhcp and assigns it to the MandatoryDhcp field.
func (o *InlineObject70) SetMandatoryDhcp(v NetworksNetworkIdApplianceVlansMandatoryDhcp) {
	o.MandatoryDhcp = &v
}

func (o InlineObject70) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.ApplianceIp) {
		toSerialize["applianceIp"] = o.ApplianceIp
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !isNil(o.VpnNatSubnet) {
		toSerialize["vpnNatSubnet"] = o.VpnNatSubnet
	}
	if !isNil(o.DhcpHandling) {
		toSerialize["dhcpHandling"] = o.DhcpHandling
	}
	if !isNil(o.DhcpRelayServerIps) {
		toSerialize["dhcpRelayServerIps"] = o.DhcpRelayServerIps
	}
	if !isNil(o.DhcpLeaseTime) {
		toSerialize["dhcpLeaseTime"] = o.DhcpLeaseTime
	}
	if !isNil(o.DhcpBootOptionsEnabled) {
		toSerialize["dhcpBootOptionsEnabled"] = o.DhcpBootOptionsEnabled
	}
	if !isNil(o.DhcpBootNextServer) {
		toSerialize["dhcpBootNextServer"] = o.DhcpBootNextServer
	}
	if !isNil(o.DhcpBootFilename) {
		toSerialize["dhcpBootFilename"] = o.DhcpBootFilename
	}
	if !isNil(o.FixedIpAssignments) {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	if !isNil(o.ReservedIpRanges) {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	if !isNil(o.DnsNameservers) {
		toSerialize["dnsNameservers"] = o.DnsNameservers
	}
	if !isNil(o.DhcpOptions) {
		toSerialize["dhcpOptions"] = o.DhcpOptions
	}
	if !isNil(o.TemplateVlanType) {
		toSerialize["templateVlanType"] = o.TemplateVlanType
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	if !isNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	if !isNil(o.MandatoryDhcp) {
		toSerialize["mandatoryDhcp"] = o.MandatoryDhcp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject70 struct {
	value *InlineObject70
	isSet bool
}

func (v NullableInlineObject70) Get() *InlineObject70 {
	return v.value
}

func (v *NullableInlineObject70) Set(val *InlineObject70) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject70) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject70) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject70(val *InlineObject70) *NullableInlineObject70 {
	return &NullableInlineObject70{value: val, isSet: true}
}

func (v NullableInlineObject70) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject70) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


