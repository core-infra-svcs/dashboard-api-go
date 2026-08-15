# \ApplianceApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNetworkApplianceUmbrellaPolicies**](ApplianceApi.md#AddNetworkApplianceUmbrellaPolicies) | **Post** /networks/{networkId}/appliance/umbrella/policies/add | Add one Cisco Umbrella DNS security policy to an MX network by policy ID
[**AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans**](ApplianceApi.md#AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans) | **Post** /organizations/{organizationId}/policies/global/group/policies/appliance/vlans/assign | Assign VLANs to a policy
[**BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate**](ApplianceApi.md#BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate) | **Post** /organizations/{organizationId}/appliance/dns/local/profiles/assignments/bulkCreate | Assign the local DNS profile to networks in the organization
[**ConnectNetworkApplianceUmbrellaAccount**](ApplianceApi.md#ConnectNetworkApplianceUmbrellaAccount) | **Post** /networks/{networkId}/appliance/umbrella/account/connect | Connect a Cisco Umbrella account to this network
[**CreateDeviceApplianceInterfacesPortsUpdate**](ApplianceApi.md#CreateDeviceApplianceInterfacesPortsUpdate) | **Post** /devices/{serial}/appliance/interfaces/ports/update | Update configurations for an appliance&#39;s specified port
[**CreateDeviceApplianceVmxAuthenticationToken**](ApplianceApi.md#CreateDeviceApplianceVmxAuthenticationToken) | **Post** /devices/{serial}/appliance/vmx/authenticationToken | Generate a new vMX authentication token
[**CreateNetworkApplianceDevicesRedundancySwap**](ApplianceApi.md#CreateNetworkApplianceDevicesRedundancySwap) | **Post** /networks/{networkId}/appliance/devices/redundancy/swap | Swap MX primary and warm spare appliances
[**CreateNetworkApplianceInterfacesL3**](ApplianceApi.md#CreateNetworkApplianceInterfacesL3) | **Post** /networks/{networkId}/appliance/interfaces/l3 | Create wired L3 interface
[**CreateNetworkAppliancePrefixesDelegatedStatic**](ApplianceApi.md#CreateNetworkAppliancePrefixesDelegatedStatic) | **Post** /networks/{networkId}/appliance/prefixes/delegated/statics | Add a static delegated prefix from a network
[**CreateNetworkApplianceRfProfile**](ApplianceApi.md#CreateNetworkApplianceRfProfile) | **Post** /networks/{networkId}/appliance/rfProfiles | Creates new RF profile for this network
[**CreateNetworkApplianceStaticRoute**](ApplianceApi.md#CreateNetworkApplianceStaticRoute) | **Post** /networks/{networkId}/appliance/staticRoutes | Add a static route for an MX or teleworker network
[**CreateNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceApi.md#CreateNetworkApplianceTrafficShapingCustomPerformanceClass) | **Post** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses | Add a custom performance class for an MX network
[**CreateNetworkApplianceVlan**](ApplianceApi.md#CreateNetworkApplianceVlan) | **Post** /networks/{networkId}/appliance/vlans | Add a VLAN
[**CreateOrganizationApplianceDnsLocalProfile**](ApplianceApi.md#CreateOrganizationApplianceDnsLocalProfile) | **Post** /organizations/{organizationId}/appliance/dns/local/profiles | Create a new local DNS profile
[**CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete**](ApplianceApi.md#CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete) | **Post** /organizations/{organizationId}/appliance/dns/local/profiles/assignments/bulkDelete | Unassign the local DNS profile to networks in the organization
[**CreateOrganizationApplianceDnsLocalRecord**](ApplianceApi.md#CreateOrganizationApplianceDnsLocalRecord) | **Post** /organizations/{organizationId}/appliance/dns/local/records | Create a new local DNS record
[**CreateOrganizationApplianceDnsSplitProfile**](ApplianceApi.md#CreateOrganizationApplianceDnsSplitProfile) | **Post** /organizations/{organizationId}/appliance/dns/split/profiles | Create a new split DNS profile
[**CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate**](ApplianceApi.md#CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate) | **Post** /organizations/{organizationId}/appliance/dns/split/profiles/assignments/bulkCreate | Assign the split DNS profile to networks in the organization
[**CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete**](ApplianceApi.md#CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete) | **Post** /organizations/{organizationId}/appliance/dns/split/profiles/assignments/bulkDelete | Unassign the split DNS profile to networks in the organization
[**DeleteNetworkApplianceInterfacesL3**](ApplianceApi.md#DeleteNetworkApplianceInterfacesL3) | **Delete** /networks/{networkId}/appliance/interfaces/l3/{interfaceId} | Delete wired L3 interface
[**DeleteNetworkAppliancePrefixesDelegatedStatic**](ApplianceApi.md#DeleteNetworkAppliancePrefixesDelegatedStatic) | **Delete** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Delete a static delegated prefix from a network
[**DeleteNetworkApplianceRfProfile**](ApplianceApi.md#DeleteNetworkApplianceRfProfile) | **Delete** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Delete a RF Profile
[**DeleteNetworkApplianceStaticRoute**](ApplianceApi.md#DeleteNetworkApplianceStaticRoute) | **Delete** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Delete a static route from an MX or teleworker network
[**DeleteNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceApi.md#DeleteNetworkApplianceTrafficShapingCustomPerformanceClass) | **Delete** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Delete a custom performance class from an MX network
[**DeleteNetworkApplianceVlan**](ApplianceApi.md#DeleteNetworkApplianceVlan) | **Delete** /networks/{networkId}/appliance/vlans/{vlanId} | Delete a VLAN from a network
[**DeleteOrganizationApplianceDnsLocalProfile**](ApplianceApi.md#DeleteOrganizationApplianceDnsLocalProfile) | **Delete** /organizations/{organizationId}/appliance/dns/local/profiles/{profileId} | Deletes a local DNS profile
[**DeleteOrganizationApplianceDnsLocalRecord**](ApplianceApi.md#DeleteOrganizationApplianceDnsLocalRecord) | **Delete** /organizations/{organizationId}/appliance/dns/local/records/{recordId} | Deletes a local DNS record
[**DeleteOrganizationApplianceDnsSplitProfile**](ApplianceApi.md#DeleteOrganizationApplianceDnsSplitProfile) | **Delete** /organizations/{organizationId}/appliance/dns/split/profiles/{profileId} | Deletes a split DNS profile
[**DisconnectNetworkApplianceUmbrellaAccount**](ApplianceApi.md#DisconnectNetworkApplianceUmbrellaAccount) | **Post** /networks/{networkId}/appliance/umbrella/account/disconnect | Disconnect Umbrella account from this network
[**ExclusionsNetworkApplianceUmbrellaDomains**](ApplianceApi.md#ExclusionsNetworkApplianceUmbrellaDomains) | **Put** /networks/{networkId}/appliance/umbrella/domains/exclusions | Specify one or more domain names to be excluded from being routed to Cisco Umbrella.
[**GetDeviceApplianceDhcpSubnets**](ApplianceApi.md#GetDeviceApplianceDhcpSubnets) | **Get** /devices/{serial}/appliance/dhcp/subnets | Return the DHCP subnet information for an appliance
[**GetDeviceAppliancePerformance**](ApplianceApi.md#GetDeviceAppliancePerformance) | **Get** /devices/{serial}/appliance/performance | Return the performance score for a single Secure Appliance or Secure Router
[**GetDeviceAppliancePrefixesDelegated**](ApplianceApi.md#GetDeviceAppliancePrefixesDelegated) | **Get** /devices/{serial}/appliance/prefixes/delegated | Return current delegated IPv6 prefixes on an appliance.
[**GetDeviceAppliancePrefixesDelegatedVlanAssignments**](ApplianceApi.md#GetDeviceAppliancePrefixesDelegatedVlanAssignments) | **Get** /devices/{serial}/appliance/prefixes/delegated/vlanAssignments | Return prefixes assigned to all IPv6 enabled VLANs on an appliance.
[**GetDeviceApplianceRadioSettings**](ApplianceApi.md#GetDeviceApplianceRadioSettings) | **Get** /devices/{serial}/appliance/radio/settings | Return the radio settings of an appliance
[**GetDeviceApplianceUplinksSettings**](ApplianceApi.md#GetDeviceApplianceUplinksSettings) | **Get** /devices/{serial}/appliance/uplinks/settings | Return the uplink settings for a secure router or security appliance
[**GetNetworkApplianceClientSecurityEvents**](ApplianceApi.md#GetNetworkApplianceClientSecurityEvents) | **Get** /networks/{networkId}/appliance/clients/{clientId}/security/events | List the security events for a client
[**GetNetworkApplianceConnectivityMonitoringDestinations**](ApplianceApi.md#GetNetworkApplianceConnectivityMonitoringDestinations) | **Get** /networks/{networkId}/appliance/connectivityMonitoringDestinations | Return the connectivity testing destinations for an MX network
[**GetNetworkApplianceContentFiltering**](ApplianceApi.md#GetNetworkApplianceContentFiltering) | **Get** /networks/{networkId}/appliance/contentFiltering | Return the content filtering settings for an MX network
[**GetNetworkApplianceContentFilteringCategories**](ApplianceApi.md#GetNetworkApplianceContentFilteringCategories) | **Get** /networks/{networkId}/appliance/contentFiltering/categories | List all available content filtering categories for an MX network
[**GetNetworkApplianceFirewallCellularFirewallRules**](ApplianceApi.md#GetNetworkApplianceFirewallCellularFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/cellularFirewallRules | Return the cellular firewall rules for an MX network
[**GetNetworkApplianceFirewallFirewalledService**](ApplianceApi.md#GetNetworkApplianceFirewallFirewalledService) | **Get** /networks/{networkId}/appliance/firewall/firewalledServices/{service} | Return the accessibility settings of the given service (&#39;ICMP&#39;, &#39;web&#39;, or &#39;SNMP&#39;)
[**GetNetworkApplianceFirewallFirewalledServices**](ApplianceApi.md#GetNetworkApplianceFirewallFirewalledServices) | **Get** /networks/{networkId}/appliance/firewall/firewalledServices | List the appliance services and their accessibility rules
[**GetNetworkApplianceFirewallInboundCellularFirewallRules**](ApplianceApi.md#GetNetworkApplianceFirewallInboundCellularFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/inboundCellularFirewallRules | Return the inbound cellular firewall rules for an MX network
[**GetNetworkApplianceFirewallInboundFirewallRules**](ApplianceApi.md#GetNetworkApplianceFirewallInboundFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/inboundFirewallRules | Return the inbound firewall rules for an MX network
[**GetNetworkApplianceFirewallL3FirewallRules**](ApplianceApi.md#GetNetworkApplianceFirewallL3FirewallRules) | **Get** /networks/{networkId}/appliance/firewall/l3FirewallRules | Return the L3 firewall rules for an MX network
[**GetNetworkApplianceFirewallL7FirewallRules**](ApplianceApi.md#GetNetworkApplianceFirewallL7FirewallRules) | **Get** /networks/{networkId}/appliance/firewall/l7FirewallRules | List the MX L7 firewall rules for an MX network
[**GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories**](ApplianceApi.md#GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories) | **Get** /networks/{networkId}/appliance/firewall/l7FirewallRules/applicationCategories | Return the L7 firewall application categories and their associated applications for an MX network
[**GetNetworkApplianceFirewallOneToManyNatRules**](ApplianceApi.md#GetNetworkApplianceFirewallOneToManyNatRules) | **Get** /networks/{networkId}/appliance/firewall/oneToManyNatRules | Return the 1:Many NAT mapping rules for an MX network
[**GetNetworkApplianceFirewallOneToOneNatRules**](ApplianceApi.md#GetNetworkApplianceFirewallOneToOneNatRules) | **Get** /networks/{networkId}/appliance/firewall/oneToOneNatRules | Return the 1:1 NAT mapping rules for an MX network
[**GetNetworkApplianceFirewallPortForwardingRules**](ApplianceApi.md#GetNetworkApplianceFirewallPortForwardingRules) | **Get** /networks/{networkId}/appliance/firewall/portForwardingRules | Return the port forwarding rules for an MX network
[**GetNetworkApplianceFirewallSettings**](ApplianceApi.md#GetNetworkApplianceFirewallSettings) | **Get** /networks/{networkId}/appliance/firewall/settings | Return the firewall settings for this network
[**GetNetworkAppliancePort**](ApplianceApi.md#GetNetworkAppliancePort) | **Get** /networks/{networkId}/appliance/ports/{portId} | Return per-port VLAN settings for a single secure router or security appliance port.
[**GetNetworkAppliancePorts**](ApplianceApi.md#GetNetworkAppliancePorts) | **Get** /networks/{networkId}/appliance/ports | List per-port VLAN settings for all ports of a secure router or security appliance.
[**GetNetworkAppliancePrefixesDelegatedStatic**](ApplianceApi.md#GetNetworkAppliancePrefixesDelegatedStatic) | **Get** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Return a static delegated prefix from a network
[**GetNetworkAppliancePrefixesDelegatedStatics**](ApplianceApi.md#GetNetworkAppliancePrefixesDelegatedStatics) | **Get** /networks/{networkId}/appliance/prefixes/delegated/statics | List static delegated prefixes for a network
[**GetNetworkApplianceRfProfile**](ApplianceApi.md#GetNetworkApplianceRfProfile) | **Get** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Return a RF profile
[**GetNetworkApplianceRfProfiles**](ApplianceApi.md#GetNetworkApplianceRfProfiles) | **Get** /networks/{networkId}/appliance/rfProfiles | List the RF profiles for this network
[**GetNetworkApplianceSecurityEvents**](ApplianceApi.md#GetNetworkApplianceSecurityEvents) | **Get** /networks/{networkId}/appliance/security/events | List the security events for a network
[**GetNetworkApplianceSecurityIntrusion**](ApplianceApi.md#GetNetworkApplianceSecurityIntrusion) | **Get** /networks/{networkId}/appliance/security/intrusion | Returns all supported intrusion settings for an MX network
[**GetNetworkApplianceSecurityMalware**](ApplianceApi.md#GetNetworkApplianceSecurityMalware) | **Get** /networks/{networkId}/appliance/security/malware | Returns all supported malware settings for an MX network
[**GetNetworkApplianceSettings**](ApplianceApi.md#GetNetworkApplianceSettings) | **Get** /networks/{networkId}/appliance/settings | Return the appliance settings for a network
[**GetNetworkApplianceSingleLan**](ApplianceApi.md#GetNetworkApplianceSingleLan) | **Get** /networks/{networkId}/appliance/singleLan | Return single LAN configuration
[**GetNetworkApplianceSsid**](ApplianceApi.md#GetNetworkApplianceSsid) | **Get** /networks/{networkId}/appliance/ssids/{number} | Return a single MX SSID
[**GetNetworkApplianceSsids**](ApplianceApi.md#GetNetworkApplianceSsids) | **Get** /networks/{networkId}/appliance/ssids | List the MX SSIDs in a network
[**GetNetworkApplianceStaticRoute**](ApplianceApi.md#GetNetworkApplianceStaticRoute) | **Get** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Return a static route for an MX or teleworker network
[**GetNetworkApplianceStaticRoutes**](ApplianceApi.md#GetNetworkApplianceStaticRoutes) | **Get** /networks/{networkId}/appliance/staticRoutes | List the static routes for an MX or teleworker network
[**GetNetworkApplianceTrafficShaping**](ApplianceApi.md#GetNetworkApplianceTrafficShaping) | **Get** /networks/{networkId}/appliance/trafficShaping | Display the traffic shaping settings for an MX network
[**GetNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceApi.md#GetNetworkApplianceTrafficShapingCustomPerformanceClass) | **Get** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Return a custom performance class for an MX network
[**GetNetworkApplianceTrafficShapingCustomPerformanceClasses**](ApplianceApi.md#GetNetworkApplianceTrafficShapingCustomPerformanceClasses) | **Get** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses | List all custom performance classes for an MX network
[**GetNetworkApplianceTrafficShapingRules**](ApplianceApi.md#GetNetworkApplianceTrafficShapingRules) | **Get** /networks/{networkId}/appliance/trafficShaping/rules | Display the traffic shaping settings rules for an MX network
[**GetNetworkApplianceTrafficShapingUplinkBandwidth**](ApplianceApi.md#GetNetworkApplianceTrafficShapingUplinkBandwidth) | **Get** /networks/{networkId}/appliance/trafficShaping/uplinkBandwidth | Returns the uplink bandwidth limits for your MX network
[**GetNetworkApplianceTrafficShapingUplinkSelection**](ApplianceApi.md#GetNetworkApplianceTrafficShapingUplinkSelection) | **Get** /networks/{networkId}/appliance/trafficShaping/uplinkSelection | Show uplink selection settings for an MX network
[**GetNetworkApplianceUplinksUsageHistory**](ApplianceApi.md#GetNetworkApplianceUplinksUsageHistory) | **Get** /networks/{networkId}/appliance/uplinks/usageHistory | Get the sent and received bytes for each uplink of a network.
[**GetNetworkApplianceVlan**](ApplianceApi.md#GetNetworkApplianceVlan) | **Get** /networks/{networkId}/appliance/vlans/{vlanId} | Return a VLAN
[**GetNetworkApplianceVlans**](ApplianceApi.md#GetNetworkApplianceVlans) | **Get** /networks/{networkId}/appliance/vlans | List the VLANs for a Security Appliance network
[**GetNetworkApplianceVlansSettings**](ApplianceApi.md#GetNetworkApplianceVlansSettings) | **Get** /networks/{networkId}/appliance/vlans/settings | Returns the enabled status of VLANs for the network
[**GetNetworkApplianceVpnBgp**](ApplianceApi.md#GetNetworkApplianceVpnBgp) | **Get** /networks/{networkId}/appliance/vpn/bgp | Return a Hub BGP Configuration
[**GetNetworkApplianceVpnSiteToSiteVpn**](ApplianceApi.md#GetNetworkApplianceVpnSiteToSiteVpn) | **Get** /networks/{networkId}/appliance/vpn/siteToSiteVpn | Return the site-to-site VPN settings of a network
[**GetNetworkApplianceWarmSpare**](ApplianceApi.md#GetNetworkApplianceWarmSpare) | **Get** /networks/{networkId}/appliance/warmSpare | Return MX warm spare settings
[**GetOrganizationApplianceDevicesInterfacesL3**](ApplianceApi.md#GetOrganizationApplianceDevicesInterfacesL3) | **Get** /organizations/{organizationId}/appliance/devices/interfaces/l3 | List L3 interfaces across networks for the organization
[**GetOrganizationApplianceDevicesInterfacesPortsByDevice**](ApplianceApi.md#GetOrganizationApplianceDevicesInterfacesPortsByDevice) | **Get** /organizations/{organizationId}/appliance/devices/interfaces/ports/byDevice | Returns port configurations for appliances in a given organization
[**GetOrganizationApplianceDevicesPortsTransceiversReadingsHistoryByDevice**](ApplianceApi.md#GetOrganizationApplianceDevicesPortsTransceiversReadingsHistoryByDevice) | **Get** /organizations/{organizationId}/appliance/devices/ports/transceivers/readings/history/byDevice | Return time-series digital optical monitoring (DOM) readings for ports on each DOM-enabled Catalyst appliance in an organization.
[**GetOrganizationApplianceDevicesRedundancyByNetwork**](ApplianceApi.md#GetOrganizationApplianceDevicesRedundancyByNetwork) | **Get** /organizations/{organizationId}/appliance/devices/redundancy/byNetwork | Return MX warm spare settings
[**GetOrganizationApplianceDnsLocalProfiles**](ApplianceApi.md#GetOrganizationApplianceDnsLocalProfiles) | **Get** /organizations/{organizationId}/appliance/dns/local/profiles | Fetch the local DNS profiles used in the organization
[**GetOrganizationApplianceDnsLocalProfilesAssignments**](ApplianceApi.md#GetOrganizationApplianceDnsLocalProfilesAssignments) | **Get** /organizations/{organizationId}/appliance/dns/local/profiles/assignments | Fetch the local DNS profile assignments in the organization
[**GetOrganizationApplianceDnsLocalRecords**](ApplianceApi.md#GetOrganizationApplianceDnsLocalRecords) | **Get** /organizations/{organizationId}/appliance/dns/local/records | Fetch the DNS records used in local DNS profiles
[**GetOrganizationApplianceDnsSplitProfiles**](ApplianceApi.md#GetOrganizationApplianceDnsSplitProfiles) | **Get** /organizations/{organizationId}/appliance/dns/split/profiles | Fetch the split DNS profiles used in the organization
[**GetOrganizationApplianceDnsSplitProfilesAssignments**](ApplianceApi.md#GetOrganizationApplianceDnsSplitProfilesAssignments) | **Get** /organizations/{organizationId}/appliance/dns/split/profiles/assignments | Fetch the split DNS profile assignments in the organization
[**GetOrganizationApplianceFirewallMulticastForwardingByNetwork**](ApplianceApi.md#GetOrganizationApplianceFirewallMulticastForwardingByNetwork) | **Get** /organizations/{organizationId}/appliance/firewall/multicastForwarding/byNetwork | List Static Multicasting forwarding settings for MX networks
[**GetOrganizationApplianceInterfacesPacketsOverviewsByDevice**](ApplianceApi.md#GetOrganizationApplianceInterfacesPacketsOverviewsByDevice) | **Get** /organizations/{organizationId}/appliance/interfaces/packets/overviews/byDevice | Returns packet counter overviews for all interfaces on Secure Routers in the organization, including totals and average rates by packet type over the requested timespan.
[**GetOrganizationApplianceRoutingVrfsSettings**](ApplianceApi.md#GetOrganizationApplianceRoutingVrfsSettings) | **Get** /organizations/{organizationId}/appliance/routing/vrfs/settings | Return the VRF setting for an organization.
[**GetOrganizationApplianceSecurityEvents**](ApplianceApi.md#GetOrganizationApplianceSecurityEvents) | **Get** /organizations/{organizationId}/appliance/security/events | List the security events for an organization
[**GetOrganizationApplianceSecurityIntrusion**](ApplianceApi.md#GetOrganizationApplianceSecurityIntrusion) | **Get** /organizations/{organizationId}/appliance/security/intrusion | Returns all supported intrusion settings for an organization
[**GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork**](ApplianceApi.md#GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork) | **Get** /organizations/{organizationId}/appliance/trafficShaping/vpnExclusions/byNetwork | Display VPN exclusion rules for MX networks.
[**GetOrganizationApplianceUplinkStatuses**](ApplianceApi.md#GetOrganizationApplianceUplinkStatuses) | **Get** /organizations/{organizationId}/appliance/uplink/statuses | List the uplink status of every Meraki MX and Z series appliances in the organization
[**GetOrganizationApplianceUplinksNatByNetwork**](ApplianceApi.md#GetOrganizationApplianceUplinksNatByNetwork) | **Get** /organizations/{organizationId}/appliance/uplinks/nat/byNetwork | Fetch uplink NAT settings of each network in the organization
[**GetOrganizationApplianceUplinksStatusesOverview**](ApplianceApi.md#GetOrganizationApplianceUplinksStatusesOverview) | **Get** /organizations/{organizationId}/appliance/uplinks/statuses/overview | Returns an overview of uplink statuses
[**GetOrganizationApplianceUplinksUsageByNetwork**](ApplianceApi.md#GetOrganizationApplianceUplinksUsageByNetwork) | **Get** /organizations/{organizationId}/appliance/uplinks/usage/byNetwork | Get the sent and received bytes for each uplink of all MX and Z networks within an organization
[**GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas**](ApplianceApi.md#GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas) | **Get** /organizations/{organizationId}/appliance/vpn/siteToSite/ipsec/peers/slas | Get the list of available IPsec SLA policies for an organization
[**GetOrganizationApplianceVpnStats**](ApplianceApi.md#GetOrganizationApplianceVpnStats) | **Get** /organizations/{organizationId}/appliance/vpn/stats | Show VPN history stats for networks in an organization
[**GetOrganizationApplianceVpnStatuses**](ApplianceApi.md#GetOrganizationApplianceVpnStatuses) | **Get** /organizations/{organizationId}/appliance/vpn/statuses | Show VPN status for networks in an organization
[**GetOrganizationApplianceVpnThirdPartyVPNPeers**](ApplianceApi.md#GetOrganizationApplianceVpnThirdPartyVPNPeers) | **Get** /organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers | Return the third party VPN peers for an organization
[**GetOrganizationApplianceVpnVpnFirewallRules**](ApplianceApi.md#GetOrganizationApplianceVpnVpnFirewallRules) | **Get** /organizations/{organizationId}/appliance/vpn/vpnFirewallRules | Return the firewall rules for an organization&#39;s site-to-site VPN
[**GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments**](ApplianceApi.md#GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments) | **Get** /organizations/{organizationId}/policies/global/group/policies/appliance/vlans/assignments | List appliance VLAN policy assignments
[**GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan**](ApplianceApi.md#GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan) | **Get** /organizations/{organizationId}/policies/global/group/policies/appliance/vlans/assignments/byVlan | List policies by appliance VLANs
[**ProtectionNetworkApplianceUmbrella**](ApplianceApi.md#ProtectionNetworkApplianceUmbrella) | **Put** /networks/{networkId}/appliance/umbrella/protection | Enable or disable umbrella protection for an appliance network
[**RemoveNetworkApplianceUmbrellaPolicies**](ApplianceApi.md#RemoveNetworkApplianceUmbrellaPolicies) | **Post** /networks/{networkId}/appliance/umbrella/policies/remove | Remove one Cisco Umbrella DNS security policy from an MX network by policy ID
[**RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans**](ApplianceApi.md#RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans) | **Post** /organizations/{organizationId}/policies/global/group/policies/appliance/vlans/remove | Remove VLANs from a policy
[**SwapNetworkApplianceWarmSpare**](ApplianceApi.md#SwapNetworkApplianceWarmSpare) | **Post** /networks/{networkId}/appliance/warmSpare/swap | Swap MX primary and warm spare appliances
[**UpdateDeviceApplianceRadioSettings**](ApplianceApi.md#UpdateDeviceApplianceRadioSettings) | **Put** /devices/{serial}/appliance/radio/settings | Update the radio settings of an appliance
[**UpdateDeviceApplianceUplinksSettings**](ApplianceApi.md#UpdateDeviceApplianceUplinksSettings) | **Put** /devices/{serial}/appliance/uplinks/settings | Update the uplink settings for a secure router or security appliance
[**UpdateNetworkApplianceConnectivityMonitoringDestinations**](ApplianceApi.md#UpdateNetworkApplianceConnectivityMonitoringDestinations) | **Put** /networks/{networkId}/appliance/connectivityMonitoringDestinations | Update the connectivity testing destinations for an MX network
[**UpdateNetworkApplianceContentFiltering**](ApplianceApi.md#UpdateNetworkApplianceContentFiltering) | **Put** /networks/{networkId}/appliance/contentFiltering | Update the content filtering settings for an MX network
[**UpdateNetworkApplianceDevicesRedundancy**](ApplianceApi.md#UpdateNetworkApplianceDevicesRedundancy) | **Put** /networks/{networkId}/appliance/devices/redundancy | Update MX warm spare settings
[**UpdateNetworkApplianceFirewallCellularFirewallRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallCellularFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/cellularFirewallRules | Update the cellular firewall rules of an MX network
[**UpdateNetworkApplianceFirewallFirewalledService**](ApplianceApi.md#UpdateNetworkApplianceFirewallFirewalledService) | **Put** /networks/{networkId}/appliance/firewall/firewalledServices/{service} | Updates the accessibility settings for the given service (&#39;ICMP&#39;, &#39;web&#39;, or &#39;SNMP&#39;)
[**UpdateNetworkApplianceFirewallInboundCellularFirewallRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallInboundCellularFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/inboundCellularFirewallRules | Update the inbound cellular firewall rules of an MX network
[**UpdateNetworkApplianceFirewallInboundFirewallRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallInboundFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/inboundFirewallRules | Update the inbound firewall rules of an MX network
[**UpdateNetworkApplianceFirewallL3FirewallRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallL3FirewallRules) | **Put** /networks/{networkId}/appliance/firewall/l3FirewallRules | Update the L3 firewall rules of an MX network
[**UpdateNetworkApplianceFirewallL7FirewallRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallL7FirewallRules) | **Put** /networks/{networkId}/appliance/firewall/l7FirewallRules | Update the MX L7 firewall rules for an MX network
[**UpdateNetworkApplianceFirewallMulticastForwarding**](ApplianceApi.md#UpdateNetworkApplianceFirewallMulticastForwarding) | **Put** /networks/{networkId}/appliance/firewall/multicastForwarding | Update static multicast forward rules for a network
[**UpdateNetworkApplianceFirewallOneToManyNatRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallOneToManyNatRules) | **Put** /networks/{networkId}/appliance/firewall/oneToManyNatRules | Set the 1:Many NAT mapping rules for an MX network
[**UpdateNetworkApplianceFirewallOneToOneNatRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallOneToOneNatRules) | **Put** /networks/{networkId}/appliance/firewall/oneToOneNatRules | Set the 1:1 NAT mapping rules for an MX network
[**UpdateNetworkApplianceFirewallPortForwardingRules**](ApplianceApi.md#UpdateNetworkApplianceFirewallPortForwardingRules) | **Put** /networks/{networkId}/appliance/firewall/portForwardingRules | Update the port forwarding rules for an MX network
[**UpdateNetworkApplianceFirewallSettings**](ApplianceApi.md#UpdateNetworkApplianceFirewallSettings) | **Put** /networks/{networkId}/appliance/firewall/settings | Update the firewall settings for this network
[**UpdateNetworkApplianceInterfacesL3**](ApplianceApi.md#UpdateNetworkApplianceInterfacesL3) | **Put** /networks/{networkId}/appliance/interfaces/l3/{interfaceId} | Update wired L3 interface
[**UpdateNetworkAppliancePort**](ApplianceApi.md#UpdateNetworkAppliancePort) | **Put** /networks/{networkId}/appliance/ports/{portId} | Update the per-port VLAN settings for a single secure router or security appliance port.
[**UpdateNetworkAppliancePrefixesDelegatedStatic**](ApplianceApi.md#UpdateNetworkAppliancePrefixesDelegatedStatic) | **Put** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Update a static delegated prefix from a network
[**UpdateNetworkApplianceRfProfile**](ApplianceApi.md#UpdateNetworkApplianceRfProfile) | **Put** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Updates specified RF profile for this network
[**UpdateNetworkApplianceSdwanInternetPolicies**](ApplianceApi.md#UpdateNetworkApplianceSdwanInternetPolicies) | **Put** /networks/{networkId}/appliance/sdwan/internetPolicies | Update SDWAN internet traffic preferences for an MX network
[**UpdateNetworkApplianceSecurityIntrusion**](ApplianceApi.md#UpdateNetworkApplianceSecurityIntrusion) | **Put** /networks/{networkId}/appliance/security/intrusion | Set the supported intrusion settings for an MX network
[**UpdateNetworkApplianceSecurityMalware**](ApplianceApi.md#UpdateNetworkApplianceSecurityMalware) | **Put** /networks/{networkId}/appliance/security/malware | Set the supported malware settings for an MX network
[**UpdateNetworkApplianceSettings**](ApplianceApi.md#UpdateNetworkApplianceSettings) | **Put** /networks/{networkId}/appliance/settings | Update the appliance settings for a network
[**UpdateNetworkApplianceSingleLan**](ApplianceApi.md#UpdateNetworkApplianceSingleLan) | **Put** /networks/{networkId}/appliance/singleLan | Update single LAN configuration
[**UpdateNetworkApplianceSsid**](ApplianceApi.md#UpdateNetworkApplianceSsid) | **Put** /networks/{networkId}/appliance/ssids/{number} | Update the attributes of an MX SSID
[**UpdateNetworkApplianceStaticRoute**](ApplianceApi.md#UpdateNetworkApplianceStaticRoute) | **Put** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Update a static route for an MX or teleworker network
[**UpdateNetworkApplianceTrafficShaping**](ApplianceApi.md#UpdateNetworkApplianceTrafficShaping) | **Put** /networks/{networkId}/appliance/trafficShaping | Update the traffic shaping settings for an MX network
[**UpdateNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceApi.md#UpdateNetworkApplianceTrafficShapingCustomPerformanceClass) | **Put** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Update a custom performance class for an MX network
[**UpdateNetworkApplianceTrafficShapingRules**](ApplianceApi.md#UpdateNetworkApplianceTrafficShapingRules) | **Put** /networks/{networkId}/appliance/trafficShaping/rules | Update the traffic shaping settings rules for an MX network
[**UpdateNetworkApplianceTrafficShapingUplinkBandwidth**](ApplianceApi.md#UpdateNetworkApplianceTrafficShapingUplinkBandwidth) | **Put** /networks/{networkId}/appliance/trafficShaping/uplinkBandwidth | Updates the uplink bandwidth settings for your MX network.
[**UpdateNetworkApplianceTrafficShapingUplinkSelection**](ApplianceApi.md#UpdateNetworkApplianceTrafficShapingUplinkSelection) | **Put** /networks/{networkId}/appliance/trafficShaping/uplinkSelection | Update uplink selection settings for an MX network
[**UpdateNetworkApplianceTrafficShapingVpnExclusions**](ApplianceApi.md#UpdateNetworkApplianceTrafficShapingVpnExclusions) | **Put** /networks/{networkId}/appliance/trafficShaping/vpnExclusions | Update VPN exclusion rules for an MX network.
[**UpdateNetworkApplianceUplinksNat**](ApplianceApi.md#UpdateNetworkApplianceUplinksNat) | **Put** /networks/{networkId}/appliance/uplinks/nat | Update uplink NAT settings of the specified network
[**UpdateNetworkApplianceVlan**](ApplianceApi.md#UpdateNetworkApplianceVlan) | **Put** /networks/{networkId}/appliance/vlans/{vlanId} | Update a VLAN
[**UpdateNetworkApplianceVlansSettings**](ApplianceApi.md#UpdateNetworkApplianceVlansSettings) | **Put** /networks/{networkId}/appliance/vlans/settings | Enable/Disable VLANs for the given network
[**UpdateNetworkApplianceVpnBgp**](ApplianceApi.md#UpdateNetworkApplianceVpnBgp) | **Put** /networks/{networkId}/appliance/vpn/bgp | Update a Hub BGP Configuration
[**UpdateNetworkApplianceVpnSiteToSiteVpn**](ApplianceApi.md#UpdateNetworkApplianceVpnSiteToSiteVpn) | **Put** /networks/{networkId}/appliance/vpn/siteToSiteVpn | Update the site-to-site VPN settings of a network
[**UpdateNetworkApplianceWarmSpare**](ApplianceApi.md#UpdateNetworkApplianceWarmSpare) | **Put** /networks/{networkId}/appliance/warmSpare | Update MX warm spare settings
[**UpdateOrganizationApplianceDnsLocalProfile**](ApplianceApi.md#UpdateOrganizationApplianceDnsLocalProfile) | **Put** /organizations/{organizationId}/appliance/dns/local/profiles/{profileId} | Update a local DNS profile
[**UpdateOrganizationApplianceDnsLocalRecord**](ApplianceApi.md#UpdateOrganizationApplianceDnsLocalRecord) | **Put** /organizations/{organizationId}/appliance/dns/local/records/{recordId} | Updates a local DNS record
[**UpdateOrganizationApplianceDnsSplitProfile**](ApplianceApi.md#UpdateOrganizationApplianceDnsSplitProfile) | **Put** /organizations/{organizationId}/appliance/dns/split/profiles/{profileId} | Update a split DNS profile
[**UpdateOrganizationApplianceRoutingVrfsSettings**](ApplianceApi.md#UpdateOrganizationApplianceRoutingVrfsSettings) | **Put** /organizations/{organizationId}/appliance/routing/vrfs/settings | Update the VRF setting for an organization.
[**UpdateOrganizationApplianceSecurityIntrusion**](ApplianceApi.md#UpdateOrganizationApplianceSecurityIntrusion) | **Put** /organizations/{organizationId}/appliance/security/intrusion | Sets supported intrusion settings for an organization
[**UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas**](ApplianceApi.md#UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas) | **Put** /organizations/{organizationId}/appliance/vpn/siteToSite/ipsec/peers/slas | Update the IPsec SLA policies for an organization
[**UpdateOrganizationApplianceVpnThirdPartyVPNPeers**](ApplianceApi.md#UpdateOrganizationApplianceVpnThirdPartyVPNPeers) | **Put** /organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers | Update the third party VPN peers for an organization
[**UpdateOrganizationApplianceVpnVpnFirewallRules**](ApplianceApi.md#UpdateOrganizationApplianceVpnVpnFirewallRules) | **Put** /organizations/{organizationId}/appliance/vpn/vpnFirewallRules | Update the firewall rules of an organization&#39;s site-to-site VPN



## AddNetworkApplianceUmbrellaPolicies

> InlineResponse200103 AddNetworkApplianceUmbrellaPolicies(ctx, networkId).AddNetworkApplianceUmbrellaPolicies(addNetworkApplianceUmbrellaPolicies).Execute()

Add one Cisco Umbrella DNS security policy to an MX network by policy ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    addNetworkApplianceUmbrellaPolicies := *openapiclient.NewInlineObject88(*openapiclient.NewNetworksNetworkIdApplianceUmbrellaPoliciesAddPolicy("Id_example")) // InlineObject88 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.AddNetworkApplianceUmbrellaPolicies(context.Background(), networkId).AddNetworkApplianceUmbrellaPolicies(addNetworkApplianceUmbrellaPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.AddNetworkApplianceUmbrellaPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddNetworkApplianceUmbrellaPolicies`: InlineResponse200103
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.AddNetworkApplianceUmbrellaPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNetworkApplianceUmbrellaPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addNetworkApplianceUmbrellaPolicies** | [**InlineObject88**](InlineObject88.md) |  | 

### Return type

[**InlineResponse200103**](InlineResponse200103.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans

> InlineResponse200406 AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(ctx, organizationId).AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(assignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans).Execute()

Assign VLANs to a policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    assignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans := *openapiclient.NewInlineObject327(*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansAssignPolicy("Id_example"), []openapiclient.OrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansAssignVlans{*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansAssignVlans("InterfaceId_example")}) // InlineObject327 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(context.Background(), organizationId).AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(assignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans`: InlineResponse200406
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans** | [**InlineObject327**](InlineObject327.md) |  | 

### Return type

[**InlineResponse200406**](InlineResponse200406.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate

> InlineResponse200278 BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate(ctx, organizationId).BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate(bulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate).Execute()

Assign the local DNS profile to networks in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    bulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate := *openapiclient.NewInlineObject246([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems()}) // InlineObject246 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate(context.Background(), organizationId).BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate(bulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate`: InlineResponse200278
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate** | [**InlineObject246**](InlineObject246.md) |  | 

### Return type

[**InlineResponse200278**](InlineResponse200278.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectNetworkApplianceUmbrellaAccount

> InlineResponse200101 ConnectNetworkApplianceUmbrellaAccount(ctx, networkId).ConnectNetworkApplianceUmbrellaAccount(connectNetworkApplianceUmbrellaAccount).Execute()

Connect a Cisco Umbrella account to this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    connectNetworkApplianceUmbrellaAccount := *openapiclient.NewInlineObject86(*openapiclient.NewNetworksNetworkIdApplianceUmbrellaAccountConnectApi("Key_example", "Secret_example")) // InlineObject86 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.ConnectNetworkApplianceUmbrellaAccount(context.Background(), networkId).ConnectNetworkApplianceUmbrellaAccount(connectNetworkApplianceUmbrellaAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.ConnectNetworkApplianceUmbrellaAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectNetworkApplianceUmbrellaAccount`: InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.ConnectNetworkApplianceUmbrellaAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectNetworkApplianceUmbrellaAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectNetworkApplianceUmbrellaAccount** | [**InlineObject86**](InlineObject86.md) |  | 

### Return type

[**InlineResponse200101**](InlineResponse200101.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceApplianceInterfacesPortsUpdate

> InlineResponse2008 CreateDeviceApplianceInterfacesPortsUpdate(ctx, serial).CreateDeviceApplianceInterfacesPortsUpdate(createDeviceApplianceInterfacesPortsUpdate).Execute()

Update configurations for an appliance's specified port



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Serial
    createDeviceApplianceInterfacesPortsUpdate := *openapiclient.NewInlineObject4() // InlineObject4 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateDeviceApplianceInterfacesPortsUpdate(context.Background(), serial).CreateDeviceApplianceInterfacesPortsUpdate(createDeviceApplianceInterfacesPortsUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateDeviceApplianceInterfacesPortsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceApplianceInterfacesPortsUpdate`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateDeviceApplianceInterfacesPortsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceApplianceInterfacesPortsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceApplianceInterfacesPortsUpdate** | [**InlineObject4**](InlineObject4.md) |  | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceApplianceVmxAuthenticationToken

> InlineResponse201 CreateDeviceApplianceVmxAuthenticationToken(ctx, serial).Execute()

Generate a new vMX authentication token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateDeviceApplianceVmxAuthenticationToken(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateDeviceApplianceVmxAuthenticationToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceApplianceVmxAuthenticationToken`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateDeviceApplianceVmxAuthenticationToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceApplianceVmxAuthenticationTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkApplianceDevicesRedundancySwap

> InlineResponse20072 CreateNetworkApplianceDevicesRedundancySwap(ctx, networkId).Execute()

Swap MX primary and warm spare appliances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkApplianceDevicesRedundancySwap(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkApplianceDevicesRedundancySwap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceDevicesRedundancySwap`: InlineResponse20072
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkApplianceDevicesRedundancySwap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceDevicesRedundancySwapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20072**](InlineResponse20072.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkApplianceInterfacesL3

> InlineResponse20116 CreateNetworkApplianceInterfacesL3(ctx, networkId).CreateNetworkApplianceInterfacesL3(createNetworkApplianceInterfacesL3).Execute()

Create wired L3 interface



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkApplianceInterfacesL3 := *openapiclient.NewInlineObject64(*openapiclient.NewNetworksNetworkIdApplianceInterfacesL3Ipv4("Address_example", "Subnet_example")) // InlineObject64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkApplianceInterfacesL3(context.Background(), networkId).CreateNetworkApplianceInterfacesL3(createNetworkApplianceInterfacesL3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkApplianceInterfacesL3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceInterfacesL3`: InlineResponse20116
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkApplianceInterfacesL3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceInterfacesL3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceInterfacesL3** | [**InlineObject64**](InlineObject64.md) |  | 

### Return type

[**InlineResponse20116**](InlineResponse20116.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkAppliancePrefixesDelegatedStatic

> InlineResponse20085 CreateNetworkAppliancePrefixesDelegatedStatic(ctx, networkId).CreateNetworkAppliancePrefixesDelegatedStatic(createNetworkAppliancePrefixesDelegatedStatic).Execute()

Add a static delegated prefix from a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkAppliancePrefixesDelegatedStatic := *openapiclient.NewInlineObject67("Prefix_example", *openapiclient.NewNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1("Type_example")) // InlineObject67 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId).CreateNetworkAppliancePrefixesDelegatedStatic(createNetworkAppliancePrefixesDelegatedStatic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkAppliancePrefixesDelegatedStatic`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkAppliancePrefixesDelegatedStatic** | [**InlineObject67**](InlineObject67.md) |  | 

### Return type

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkApplianceRfProfile

> InlineResponse20086Assigned CreateNetworkApplianceRfProfile(ctx, networkId).CreateNetworkApplianceRfProfile(createNetworkApplianceRfProfile).Execute()

Creates new RF profile for this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkApplianceRfProfile := *openapiclient.NewInlineObject69("Name_example") // InlineObject69 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkApplianceRfProfile(context.Background(), networkId).CreateNetworkApplianceRfProfile(createNetworkApplianceRfProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceRfProfile`: InlineResponse20086Assigned
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceRfProfile** | [**InlineObject69**](InlineObject69.md) |  | 

### Return type

[**InlineResponse20086Assigned**](InlineResponse20086Assigned.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkApplianceStaticRoute

> InlineResponse20094 CreateNetworkApplianceStaticRoute(ctx, networkId).CreateNetworkApplianceStaticRoute(createNetworkApplianceStaticRoute).Execute()

Add a static route for an MX or teleworker network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkApplianceStaticRoute := *openapiclient.NewInlineObject77("Name_example", "Subnet_example", "GatewayIp_example") // InlineObject77 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkApplianceStaticRoute(context.Background(), networkId).CreateNetworkApplianceStaticRoute(createNetworkApplianceStaticRoute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkApplianceStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceStaticRoute`: InlineResponse20094
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceStaticRoute** | [**InlineObject77**](InlineObject77.md) |  | 

### Return type

[**InlineResponse20094**](InlineResponse20094.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkApplianceTrafficShapingCustomPerformanceClass

> InlineResponse20096 CreateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId).CreateNetworkApplianceTrafficShapingCustomPerformanceClass(createNetworkApplianceTrafficShapingCustomPerformanceClass).Execute()

Add a custom performance class for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkApplianceTrafficShapingCustomPerformanceClass := *openapiclient.NewInlineObject80("Name_example") // InlineObject80 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId).CreateNetworkApplianceTrafficShapingCustomPerformanceClass(createNetworkApplianceTrafficShapingCustomPerformanceClass).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceTrafficShapingCustomPerformanceClass`: InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceTrafficShapingCustomPerformanceClass** | [**InlineObject80**](InlineObject80.md) |  | 

### Return type

[**InlineResponse20096**](InlineResponse20096.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkApplianceVlan

> InlineResponse20117 CreateNetworkApplianceVlan(ctx, networkId).CreateNetworkApplianceVlan(createNetworkApplianceVlan).Execute()

Add a VLAN



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkApplianceVlan := *openapiclient.NewInlineObject92("Id_example", "Name_example") // InlineObject92 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateNetworkApplianceVlan(context.Background(), networkId).CreateNetworkApplianceVlan(createNetworkApplianceVlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceVlan`: InlineResponse20117
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceVlan** | [**InlineObject92**](InlineObject92.md) |  | 

### Return type

[**InlineResponse20117**](InlineResponse20117.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApplianceDnsLocalProfile

> InlineResponse200276 CreateOrganizationApplianceDnsLocalProfile(ctx, organizationId).CreateOrganizationApplianceDnsLocalProfile(createOrganizationApplianceDnsLocalProfile).Execute()

Create a new local DNS profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationApplianceDnsLocalProfile := *openapiclient.NewInlineObject245("Name_example") // InlineObject245 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateOrganizationApplianceDnsLocalProfile(context.Background(), organizationId).CreateOrganizationApplianceDnsLocalProfile(createOrganizationApplianceDnsLocalProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateOrganizationApplianceDnsLocalProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsLocalProfile`: InlineResponse200276
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateOrganizationApplianceDnsLocalProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceDnsLocalProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceDnsLocalProfile** | [**InlineObject245**](InlineObject245.md) |  | 

### Return type

[**InlineResponse200276**](InlineResponse200276.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete

> InlineResponse200278 CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(ctx, organizationId).CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete).Execute()

Unassign the local DNS profile to networks in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete := *openapiclient.NewInlineObject247([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems()}) // InlineObject247 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(context.Background(), organizationId).CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete`: InlineResponse200278
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete** | [**InlineObject247**](InlineObject247.md) |  | 

### Return type

[**InlineResponse200278**](InlineResponse200278.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApplianceDnsLocalRecord

> []InlineResponse200279 CreateOrganizationApplianceDnsLocalRecord(ctx, organizationId).CreateOrganizationApplianceDnsLocalRecord(createOrganizationApplianceDnsLocalRecord).Execute()

Create a new local DNS record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationApplianceDnsLocalRecord := *openapiclient.NewInlineObject249("Hostname_example", "Address_example", *openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile1("Id_example")) // InlineObject249 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateOrganizationApplianceDnsLocalRecord(context.Background(), organizationId).CreateOrganizationApplianceDnsLocalRecord(createOrganizationApplianceDnsLocalRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateOrganizationApplianceDnsLocalRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsLocalRecord`: []InlineResponse200279
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateOrganizationApplianceDnsLocalRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceDnsLocalRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceDnsLocalRecord** | [**InlineObject249**](InlineObject249.md) |  | 

### Return type

[**[]InlineResponse200279**](InlineResponse200279.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApplianceDnsSplitProfile

> InlineResponse200280 CreateOrganizationApplianceDnsSplitProfile(ctx, organizationId).CreateOrganizationApplianceDnsSplitProfile(createOrganizationApplianceDnsSplitProfile).Execute()

Create a new split DNS profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationApplianceDnsSplitProfile := *openapiclient.NewInlineObject251("Name_example", []string{"Hostnames_example"}, *openapiclient.NewOrganizationsOrganizationIdApplianceDnsSplitProfilesNameservers1([]string{"Addresses_example"})) // InlineObject251 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateOrganizationApplianceDnsSplitProfile(context.Background(), organizationId).CreateOrganizationApplianceDnsSplitProfile(createOrganizationApplianceDnsSplitProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateOrganizationApplianceDnsSplitProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsSplitProfile`: InlineResponse200280
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateOrganizationApplianceDnsSplitProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceDnsSplitProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceDnsSplitProfile** | [**InlineObject251**](InlineObject251.md) |  | 

### Return type

[**InlineResponse200280**](InlineResponse200280.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate

> InlineResponse200282 CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(ctx, organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate).Execute()

Assign the split DNS profile to networks in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate := *openapiclient.NewInlineObject252([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems()}) // InlineObject252 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(context.Background(), organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate`: InlineResponse200282
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate** | [**InlineObject252**](InlineObject252.md) |  | 

### Return type

[**InlineResponse200282**](InlineResponse200282.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete

> InlineResponse200282 CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(ctx, organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete).Execute()

Unassign the split DNS profile to networks in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete := *openapiclient.NewInlineObject253([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems()}) // InlineObject253 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(context.Background(), organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete`: InlineResponse200282
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete** | [**InlineObject253**](InlineObject253.md) |  | 

### Return type

[**InlineResponse200282**](InlineResponse200282.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceInterfacesL3

> DeleteNetworkApplianceInterfacesL3(ctx, networkId, interfaceId).Execute()

Delete wired L3 interface



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    interfaceId := "interfaceId_example" // string | Interface ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.DeleteNetworkApplianceInterfacesL3(context.Background(), networkId, interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteNetworkApplianceInterfacesL3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceInterfacesL3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkAppliancePrefixesDelegatedStatic

> map[string]interface{} DeleteNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).Execute()

Delete a static delegated prefix from a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.DeleteNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkAppliancePrefixesDelegatedStatic`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.DeleteNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceRfProfile

> DeleteNetworkApplianceRfProfile(ctx, networkId, rfProfileId).Execute()

Delete a RF Profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    rfProfileId := "rfProfileId_example" // string | Rf profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.DeleteNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceStaticRoute

> DeleteNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).Execute()

Delete a static route from an MX or teleworker network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    staticRouteId := "staticRouteId_example" // string | Static route ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.DeleteNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteNetworkApplianceStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceTrafficShapingCustomPerformanceClass

> DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).Execute()

Delete a custom performance class from an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceVlan

> DeleteNetworkApplianceVlan(ctx, networkId, vlanId).Execute()

Delete a VLAN from a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    vlanId := "vlanId_example" // string | Vlan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.DeleteNetworkApplianceVlan(context.Background(), networkId, vlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationApplianceDnsLocalProfile

> DeleteOrganizationApplianceDnsLocalProfile(ctx, organizationId, profileId).Execute()

Deletes a local DNS profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    profileId := "profileId_example" // string | Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.DeleteOrganizationApplianceDnsLocalProfile(context.Background(), organizationId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteOrganizationApplianceDnsLocalProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationApplianceDnsLocalProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationApplianceDnsLocalRecord

> DeleteOrganizationApplianceDnsLocalRecord(ctx, organizationId, recordId).Execute()

Deletes a local DNS record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    recordId := "recordId_example" // string | Record ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.DeleteOrganizationApplianceDnsLocalRecord(context.Background(), organizationId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteOrganizationApplianceDnsLocalRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**recordId** | **string** | Record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationApplianceDnsLocalRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationApplianceDnsSplitProfile

> DeleteOrganizationApplianceDnsSplitProfile(ctx, organizationId, profileId).Execute()

Deletes a split DNS profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    profileId := "profileId_example" // string | Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.DeleteOrganizationApplianceDnsSplitProfile(context.Background(), organizationId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DeleteOrganizationApplianceDnsSplitProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationApplianceDnsSplitProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectNetworkApplianceUmbrellaAccount

> DisconnectNetworkApplianceUmbrellaAccount(ctx, networkId).Execute()

Disconnect Umbrella account from this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.DisconnectNetworkApplianceUmbrellaAccount(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.DisconnectNetworkApplianceUmbrellaAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectNetworkApplianceUmbrellaAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExclusionsNetworkApplianceUmbrellaDomains

> InlineResponse200102 ExclusionsNetworkApplianceUmbrellaDomains(ctx, networkId).ExclusionsNetworkApplianceUmbrellaDomains(exclusionsNetworkApplianceUmbrellaDomains).Execute()

Specify one or more domain names to be excluded from being routed to Cisco Umbrella.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    exclusionsNetworkApplianceUmbrellaDomains := *openapiclient.NewInlineObject87([]string{"Domains_example"}) // InlineObject87 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.ExclusionsNetworkApplianceUmbrellaDomains(context.Background(), networkId).ExclusionsNetworkApplianceUmbrellaDomains(exclusionsNetworkApplianceUmbrellaDomains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.ExclusionsNetworkApplianceUmbrellaDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExclusionsNetworkApplianceUmbrellaDomains`: InlineResponse200102
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.ExclusionsNetworkApplianceUmbrellaDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExclusionsNetworkApplianceUmbrellaDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exclusionsNetworkApplianceUmbrellaDomains** | [**InlineObject87**](InlineObject87.md) |  | 

### Return type

[**InlineResponse200102**](InlineResponse200102.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceApplianceDhcpSubnets

> []InlineResponse2007 GetDeviceApplianceDhcpSubnets(ctx, serial).Execute()

Return the DHCP subnet information for an appliance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceApplianceDhcpSubnets(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceApplianceDhcpSubnets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceApplianceDhcpSubnets`: []InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceApplianceDhcpSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceDhcpSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse2007**](InlineResponse2007.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceAppliancePerformance

> InlineResponse2009 GetDeviceAppliancePerformance(ctx, serial).T0(t0).T1(t1).Timespan(timespan).Execute()

Return the performance score for a single Secure Appliance or Secure Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Serial
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 30 minutes and be less than or equal to 14 days. The default is 30 minutes. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceAppliancePerformance(context.Background(), serial).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceAppliancePerformance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceAppliancePerformance`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceAppliancePerformance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 30 minutes and be less than or equal to 14 days. The default is 30 minutes. | 

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceAppliancePrefixesDelegated

> []InlineResponse20010 GetDeviceAppliancePrefixesDelegated(ctx, serial).Execute()

Return current delegated IPv6 prefixes on an appliance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceAppliancePrefixesDelegated(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceAppliancePrefixesDelegated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceAppliancePrefixesDelegated`: []InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceAppliancePrefixesDelegated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePrefixesDelegatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20010**](InlineResponse20010.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceAppliancePrefixesDelegatedVlanAssignments

> []InlineResponse20011 GetDeviceAppliancePrefixesDelegatedVlanAssignments(ctx, serial).Execute()

Return prefixes assigned to all IPv6 enabled VLANs on an appliance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceAppliancePrefixesDelegatedVlanAssignments(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceAppliancePrefixesDelegatedVlanAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceAppliancePrefixesDelegatedVlanAssignments`: []InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceAppliancePrefixesDelegatedVlanAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20011**](InlineResponse20011.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceApplianceRadioSettings

> InlineResponse20012 GetDeviceApplianceRadioSettings(ctx, serial).Execute()

Return the radio settings of an appliance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceApplianceRadioSettings(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceApplianceRadioSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceApplianceRadioSettings`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceApplianceRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceApplianceUplinksSettings

> InlineResponse20013 GetDeviceApplianceUplinksSettings(ctx, serial).Execute()

Return the uplink settings for a secure router or security appliance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetDeviceApplianceUplinksSettings(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetDeviceApplianceUplinksSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceApplianceUplinksSettings`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetDeviceApplianceUplinksSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceUplinksSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceClientSecurityEvents

> []map[string]interface{} GetNetworkApplianceClientSecurityEvents(ctx, networkId, clientId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for a client



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    clientId := "clientId_example" // string | Client ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceClientSecurityEvents(context.Background(), networkId, clientId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceClientSecurityEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceClientSecurityEvents`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceClientSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceClientSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceConnectivityMonitoringDestinations

> InlineResponse20069 GetNetworkApplianceConnectivityMonitoringDestinations(ctx, networkId).Execute()

Return the connectivity testing destinations for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceConnectivityMonitoringDestinations(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceConnectivityMonitoringDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceConnectivityMonitoringDestinations`: InlineResponse20069
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20069**](InlineResponse20069.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceContentFiltering

> InlineResponse20070 GetNetworkApplianceContentFiltering(ctx, networkId).Execute()

Return the content filtering settings for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceContentFiltering(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceContentFiltering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceContentFiltering`: InlineResponse20070
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceContentFiltering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceContentFilteringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20070**](InlineResponse20070.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceContentFilteringCategories

> InlineResponse20071 GetNetworkApplianceContentFilteringCategories(ctx, networkId).Execute()

List all available content filtering categories for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceContentFilteringCategories(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceContentFilteringCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceContentFilteringCategories`: InlineResponse20071
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceContentFilteringCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceContentFilteringCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20071**](InlineResponse20071.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallCellularFirewallRules

> InlineResponse20073 GetNetworkApplianceFirewallCellularFirewallRules(ctx, networkId).Execute()

Return the cellular firewall rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallCellularFirewallRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallCellularFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallCellularFirewallRules`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallFirewalledService

> InlineResponse20075 GetNetworkApplianceFirewallFirewalledService(ctx, networkId, service).Execute()

Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP')



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    service := "service_example" // string | Service

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallFirewalledService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallFirewalledService`: InlineResponse20075
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallFirewalledService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**service** | **string** | Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallFirewalledServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20075**](InlineResponse20075.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallFirewalledServices

> []InlineResponse20074 GetNetworkApplianceFirewallFirewalledServices(ctx, networkId).Execute()

List the appliance services and their accessibility rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallFirewalledServices(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallFirewalledServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallFirewalledServices`: []InlineResponse20074
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallFirewalledServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallFirewalledServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20074**](InlineResponse20074.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallInboundCellularFirewallRules

> InlineResponse20073 GetNetworkApplianceFirewallInboundCellularFirewallRules(ctx, networkId).Execute()

Return the inbound cellular firewall rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallInboundCellularFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallInboundCellularFirewallRules`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallInboundCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallInboundCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallInboundFirewallRules

> InlineResponse20076 GetNetworkApplianceFirewallInboundFirewallRules(ctx, networkId).Execute()

Return the inbound firewall rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallInboundFirewallRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallInboundFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallInboundFirewallRules`: InlineResponse20076
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallInboundFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallInboundFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20076**](InlineResponse20076.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallL3FirewallRules

> InlineResponse20073 GetNetworkApplianceFirewallL3FirewallRules(ctx, networkId).Execute()

Return the L3 firewall rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallL3FirewallRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallL3FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallL3FirewallRules`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallL3FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallL7FirewallRules

> InlineResponse20077 GetNetworkApplianceFirewallL7FirewallRules(ctx, networkId).Execute()

List the MX L7 firewall rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallL7FirewallRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallL7FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallL7FirewallRules`: InlineResponse20077
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallL7FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallL7FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20077**](InlineResponse20077.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories

> InlineResponse20078 GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(ctx, networkId).Execute()

Return the L7 firewall application categories and their associated applications for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories`: InlineResponse20078
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20078**](InlineResponse20078.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallOneToManyNatRules

> InlineResponse20080 GetNetworkApplianceFirewallOneToManyNatRules(ctx, networkId).Execute()

Return the 1:Many NAT mapping rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallOneToManyNatRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallOneToManyNatRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallOneToManyNatRules`: InlineResponse20080
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallOneToManyNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallOneToManyNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20080**](InlineResponse20080.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallOneToOneNatRules

> InlineResponse20081 GetNetworkApplianceFirewallOneToOneNatRules(ctx, networkId).Execute()

Return the 1:1 NAT mapping rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallOneToOneNatRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallOneToOneNatRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallOneToOneNatRules`: InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallOneToOneNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallOneToOneNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20081**](InlineResponse20081.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallPortForwardingRules

> InlineResponse20082 GetNetworkApplianceFirewallPortForwardingRules(ctx, networkId).Execute()

Return the port forwarding rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallPortForwardingRules`: InlineResponse20082
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20082**](InlineResponse20082.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallSettings

> InlineResponse20083 GetNetworkApplianceFirewallSettings(ctx, networkId).Execute()

Return the firewall settings for this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceFirewallSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceFirewallSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallSettings`: InlineResponse20083
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceFirewallSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20083**](InlineResponse20083.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePort

> InlineResponse20084 GetNetworkAppliancePort(ctx, networkId, portId).Execute()

Return per-port VLAN settings for a single secure router or security appliance port.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    portId := "portId_example" // string | Port ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkAppliancePort(context.Background(), networkId, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkAppliancePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAppliancePort`: InlineResponse20084
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkAppliancePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20084**](InlineResponse20084.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePorts

> []InlineResponse20084 GetNetworkAppliancePorts(ctx, networkId).Execute()

List per-port VLAN settings for all ports of a secure router or security appliance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkAppliancePorts(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkAppliancePorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAppliancePorts`: []InlineResponse20084
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkAppliancePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20084**](InlineResponse20084.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePrefixesDelegatedStatic

> InlineResponse20085 GetNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).Execute()

Return a static delegated prefix from a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAppliancePrefixesDelegatedStatic`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePrefixesDelegatedStatics

> []InlineResponse20085 GetNetworkAppliancePrefixesDelegatedStatics(ctx, networkId).Execute()

List static delegated prefixes for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatics(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAppliancePrefixesDelegatedStatics`: []InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkAppliancePrefixesDelegatedStatics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePrefixesDelegatedStaticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20085**](InlineResponse20085.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceRfProfile

> InlineResponse20086Assigned GetNetworkApplianceRfProfile(ctx, networkId, rfProfileId).Execute()

Return a RF profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    rfProfileId := "rfProfileId_example" // string | Rf profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceRfProfile`: InlineResponse20086Assigned
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20086Assigned**](InlineResponse20086Assigned.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceRfProfiles

> InlineResponse20086 GetNetworkApplianceRfProfiles(ctx, networkId).Execute()

List the RF profiles for this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceRfProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceRfProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceRfProfiles`: InlineResponse20086
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceRfProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceRfProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20086**](InlineResponse20086.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSecurityEvents

> []map[string]interface{} GetNetworkApplianceSecurityEvents(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSecurityEvents(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSecurityEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSecurityEvents`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSecurityIntrusion

> InlineResponse20088 GetNetworkApplianceSecurityIntrusion(ctx, networkId).Execute()

Returns all supported intrusion settings for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSecurityIntrusion(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSecurityIntrusion`: InlineResponse20088
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20088**](InlineResponse20088.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSecurityMalware

> InlineResponse20090 GetNetworkApplianceSecurityMalware(ctx, networkId).Execute()

Returns all supported malware settings for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSecurityMalware(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSecurityMalware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSecurityMalware`: InlineResponse20090
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSecurityMalware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityMalwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20090**](InlineResponse20090.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSettings

> InlineResponse20091 GetNetworkApplianceSettings(ctx, networkId).Execute()

Return the appliance settings for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSettings`: InlineResponse20091
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20091**](InlineResponse20091.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSingleLan

> InlineResponse20092 GetNetworkApplianceSingleLan(ctx, networkId).Execute()

Return single LAN configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSingleLan(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSingleLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSingleLan`: InlineResponse20092
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSingleLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSingleLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20092**](InlineResponse20092.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSsid

> InlineResponse20093 GetNetworkApplianceSsid(ctx, networkId, number).Execute()

Return a single MX SSID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSsid(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSsid`: InlineResponse20093
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20093**](InlineResponse20093.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSsids

> []InlineResponse20093 GetNetworkApplianceSsids(ctx, networkId).Execute()

List the MX SSIDs in a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceSsids(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceSsids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSsids`: []InlineResponse20093
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceSsids`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSsidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20093**](InlineResponse20093.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceStaticRoute

> InlineResponse20094 GetNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).Execute()

Return a static route for an MX or teleworker network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    staticRouteId := "staticRouteId_example" // string | Static route ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceStaticRoute`: InlineResponse20094
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20094**](InlineResponse20094.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceStaticRoutes

> []InlineResponse20094 GetNetworkApplianceStaticRoutes(ctx, networkId).Execute()

List the static routes for an MX or teleworker network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceStaticRoutes(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceStaticRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceStaticRoutes`: []InlineResponse20094
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20094**](InlineResponse20094.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShaping

> InlineResponse20095 GetNetworkApplianceTrafficShaping(ctx, networkId).Execute()

Display the traffic shaping settings for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShaping(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShaping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShaping`: InlineResponse20095
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShaping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20095**](InlineResponse20095.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingCustomPerformanceClass

> InlineResponse20096 GetNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).Execute()

Return a custom performance class for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingCustomPerformanceClass`: InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20096**](InlineResponse20096.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingCustomPerformanceClasses

> []InlineResponse20096 GetNetworkApplianceTrafficShapingCustomPerformanceClasses(ctx, networkId).Execute()

List all custom performance classes for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClasses(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClasses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingCustomPerformanceClasses`: []InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShapingCustomPerformanceClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingCustomPerformanceClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20096**](InlineResponse20096.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingRules

> InlineResponse20097 GetNetworkApplianceTrafficShapingRules(ctx, networkId).Execute()

Display the traffic shaping settings rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShapingRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingRules`: InlineResponse20097
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20097**](InlineResponse20097.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingUplinkBandwidth

> InlineResponse20098 GetNetworkApplianceTrafficShapingUplinkBandwidth(ctx, networkId).Execute()

Returns the uplink bandwidth limits for your MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShapingUplinkBandwidth(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShapingUplinkBandwidth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingUplinkBandwidth`: InlineResponse20098
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShapingUplinkBandwidth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingUplinkBandwidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20098**](InlineResponse20098.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingUplinkSelection

> InlineResponse20099 GetNetworkApplianceTrafficShapingUplinkSelection(ctx, networkId).Execute()

Show uplink selection settings for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceTrafficShapingUplinkSelection(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceTrafficShapingUplinkSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingUplinkSelection`: InlineResponse20099
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceTrafficShapingUplinkSelection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingUplinkSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20099**](InlineResponse20099.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceUplinksUsageHistory

> []InlineResponse200106 GetNetworkApplianceUplinksUsageHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()

Get the sent and received bytes for each uplink of a network.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 10 minutes. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 600, 1800, 3600, 86400. The default is 60. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceUplinksUsageHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceUplinksUsageHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceUplinksUsageHistory`: []InlineResponse200106
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceUplinksUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceUplinksUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 10 minutes. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 600, 1800, 3600, 86400. The default is 60. | 

### Return type

[**[]InlineResponse200106**](InlineResponse200106.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlan

> InlineResponse200107 GetNetworkApplianceVlan(ctx, networkId, vlanId).Execute()

Return a VLAN



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    vlanId := "vlanId_example" // string | Vlan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceVlan(context.Background(), networkId, vlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVlan`: InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200107**](InlineResponse200107.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlans

> []InlineResponse200107 GetNetworkApplianceVlans(ctx, networkId).Execute()

List the VLANs for a Security Appliance network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceVlans(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceVlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVlans`: []InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200107**](InlineResponse200107.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlansSettings

> InlineResponse200108 GetNetworkApplianceVlansSettings(ctx, networkId).Execute()

Returns the enabled status of VLANs for the network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceVlansSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceVlansSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVlansSettings`: InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceVlansSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlansSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200108**](InlineResponse200108.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVpnBgp

> InlineResponse200109 GetNetworkApplianceVpnBgp(ctx, networkId).Execute()

Return a Hub BGP Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceVpnBgp(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceVpnBgp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVpnBgp`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceVpnBgp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVpnBgpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVpnSiteToSiteVpn

> InlineResponse200110 GetNetworkApplianceVpnSiteToSiteVpn(ctx, networkId).Execute()

Return the site-to-site VPN settings of a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceVpnSiteToSiteVpn(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceVpnSiteToSiteVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVpnSiteToSiteVpn`: InlineResponse200110
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceVpnSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVpnSiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200110**](InlineResponse200110.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceWarmSpare

> InlineResponse200111 GetNetworkApplianceWarmSpare(ctx, networkId).Execute()

Return MX warm spare settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetNetworkApplianceWarmSpare(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetNetworkApplianceWarmSpare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceWarmSpare`: InlineResponse200111
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200111**](InlineResponse200111.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceDevicesInterfacesL3

> InlineResponse200273 GetOrganizationApplianceDevicesInterfacesL3(ctx, organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List L3 interfaces across networks for the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    networkIds := []string{"Inner_example"} // []string | Optional Network IDs to filter results (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceDevicesInterfacesL3(context.Background(), organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceDevicesInterfacesL3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDevicesInterfacesL3`: InlineResponse200273
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceDevicesInterfacesL3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDevicesInterfacesL3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional Network IDs to filter results | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200273**](InlineResponse200273.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceDevicesInterfacesPortsByDevice

> InlineResponse200274 GetOrganizationApplianceDevicesInterfacesPortsByDevice(ctx, organizationId).Serials(serials).Numbers(numbers).Execute()

Returns port configurations for appliances in a given organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    serials := []string{"Inner_example"} // []string | Parameter to filter the results by device serials (optional)
    numbers := []string{"Inner_example"} // []string | Parameter to filter the results by specific ports (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceDevicesInterfacesPortsByDevice(context.Background(), organizationId).Serials(serials).Numbers(numbers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceDevicesInterfacesPortsByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDevicesInterfacesPortsByDevice`: InlineResponse200274
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceDevicesInterfacesPortsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDevicesInterfacesPortsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Parameter to filter the results by device serials | 
 **numbers** | **[]string** | Parameter to filter the results by specific ports | 

### Return type

[**InlineResponse200274**](InlineResponse200274.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceDevicesPortsTransceiversReadingsHistoryByDevice

> InlineResponse200275 GetOrganizationApplianceDevicesPortsTransceiversReadingsHistoryByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).NetworkIds(networkIds).Serials(serials).PortIds(portIds).Execute()

Return time-series digital optical monitoring (DOM) readings for ports on each DOM-enabled Catalyst appliance in an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 10. Default is 5. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 30 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 30 days. The default is 1 day. If interval is provided, the timespan will be autocalculated. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 14400, 86400. The default is 1200. Interval is calculated if time params are provided. (optional)
    networkIds := []string{"Inner_example"} // []string | Networks for which information should be gathered. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter usage by appliance serial. (optional)
    portIds := []string{"Inner_example"} // []string | Optional parameter to filter usage by port ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceDevicesPortsTransceiversReadingsHistoryByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).NetworkIds(networkIds).Serials(serials).PortIds(portIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceDevicesPortsTransceiversReadingsHistoryByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDevicesPortsTransceiversReadingsHistoryByDevice`: InlineResponse200275
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceDevicesPortsTransceiversReadingsHistoryByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDevicesPortsTransceiversReadingsHistoryByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 10. Default is 5. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 30 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 30 days. The default is 1 day. If interval is provided, the timespan will be autocalculated. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 14400, 86400. The default is 1200. Interval is calculated if time params are provided. | 
 **networkIds** | **[]string** | Networks for which information should be gathered. | 
 **serials** | **[]string** | Optional parameter to filter usage by appliance serial. | 
 **portIds** | **[]string** | Optional parameter to filter usage by port ID. | 

### Return type

[**InlineResponse200275**](InlineResponse200275.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceDevicesRedundancyByNetwork

> InlineResponse20072 GetOrganizationApplianceDevicesRedundancyByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return MX warm spare settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 5 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceDevicesRedundancyByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceDevicesRedundancyByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDevicesRedundancyByNetwork`: InlineResponse20072
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceDevicesRedundancyByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDevicesRedundancyByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 5 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse20072**](InlineResponse20072.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceDnsLocalProfiles

> []InlineResponse200276 GetOrganizationApplianceDnsLocalProfiles(ctx, organizationId).ProfileIds(profileIds).Execute()

Fetch the local DNS profiles used in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    profileIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by profile IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceDnsLocalProfiles(context.Background(), organizationId).ProfileIds(profileIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceDnsLocalProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDnsLocalProfiles`: []InlineResponse200276
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceDnsLocalProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDnsLocalProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileIds** | **[]string** | Optional parameter to filter the results by profile IDs | 

### Return type

[**[]InlineResponse200276**](InlineResponse200276.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceDnsLocalProfilesAssignments

> InlineResponse200277 GetOrganizationApplianceDnsLocalProfilesAssignments(ctx, organizationId).ProfileIds(profileIds).NetworkIds(networkIds).Execute()

Fetch the local DNS profile assignments in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    profileIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by profile IDs (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by network IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceDnsLocalProfilesAssignments(context.Background(), organizationId).ProfileIds(profileIds).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceDnsLocalProfilesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDnsLocalProfilesAssignments`: InlineResponse200277
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceDnsLocalProfilesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDnsLocalProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileIds** | **[]string** | Optional parameter to filter the results by profile IDs | 
 **networkIds** | **[]string** | Optional parameter to filter the results by network IDs | 

### Return type

[**InlineResponse200277**](InlineResponse200277.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceDnsLocalRecords

> []InlineResponse200279 GetOrganizationApplianceDnsLocalRecords(ctx, organizationId).ProfileIds(profileIds).Execute()

Fetch the DNS records used in local DNS profiles



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    profileIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by profile IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceDnsLocalRecords(context.Background(), organizationId).ProfileIds(profileIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceDnsLocalRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDnsLocalRecords`: []InlineResponse200279
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceDnsLocalRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDnsLocalRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileIds** | **[]string** | Optional parameter to filter the results by profile IDs | 

### Return type

[**[]InlineResponse200279**](InlineResponse200279.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceDnsSplitProfiles

> []InlineResponse200280 GetOrganizationApplianceDnsSplitProfiles(ctx, organizationId).ProfileIds(profileIds).Execute()

Fetch the split DNS profiles used in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    profileIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by profile IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceDnsSplitProfiles(context.Background(), organizationId).ProfileIds(profileIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceDnsSplitProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDnsSplitProfiles`: []InlineResponse200280
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceDnsSplitProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDnsSplitProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileIds** | **[]string** | Optional parameter to filter the results by profile IDs | 

### Return type

[**[]InlineResponse200280**](InlineResponse200280.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceDnsSplitProfilesAssignments

> InlineResponse200281 GetOrganizationApplianceDnsSplitProfilesAssignments(ctx, organizationId).ProfileIds(profileIds).NetworkIds(networkIds).Execute()

Fetch the split DNS profile assignments in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    profileIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by profile IDs (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by network IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceDnsSplitProfilesAssignments(context.Background(), organizationId).ProfileIds(profileIds).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceDnsSplitProfilesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDnsSplitProfilesAssignments`: InlineResponse200281
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceDnsSplitProfilesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileIds** | **[]string** | Optional parameter to filter the results by profile IDs | 
 **networkIds** | **[]string** | Optional parameter to filter the results by network IDs | 

### Return type

[**InlineResponse200281**](InlineResponse200281.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceFirewallMulticastForwardingByNetwork

> InlineResponse200283 GetOrganizationApplianceFirewallMulticastForwardingByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

List Static Multicasting forwarding settings for MX networks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by network IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceFirewallMulticastForwardingByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceFirewallMulticastForwardingByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceFirewallMulticastForwardingByNetwork`: InlineResponse200283
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceFirewallMulticastForwardingByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceFirewallMulticastForwardingByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter the results by network IDs | 

### Return type

[**InlineResponse200283**](InlineResponse200283.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceInterfacesPacketsOverviewsByDevice

> InlineResponse200284 GetOrganizationApplianceInterfacesPacketsOverviewsByDevice(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()

Returns packet counter overviews for all interfaces on Secure Routers in the organization, including totals and average rates by packet type over the requested timespan.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 14 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 50. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter Secure Routers in the provided networks (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter Secure Routers by their serial numbers (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceInterfacesPacketsOverviewsByDevice(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceInterfacesPacketsOverviewsByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceInterfacesPacketsOverviewsByDevice`: InlineResponse200284
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceInterfacesPacketsOverviewsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceInterfacesPacketsOverviewsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 14 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 50. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter Secure Routers in the provided networks | 
 **serials** | **[]string** | Optional parameter to filter Secure Routers by their serial numbers | 

### Return type

[**InlineResponse200284**](InlineResponse200284.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceRoutingVrfsSettings

> InlineResponse200285 GetOrganizationApplianceRoutingVrfsSettings(ctx, organizationId).Execute()

Return the VRF setting for an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceRoutingVrfsSettings(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceRoutingVrfsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceRoutingVrfsSettings`: InlineResponse200285
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceRoutingVrfsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceRoutingVrfsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200285**](InlineResponse200285.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityEvents

> []map[string]interface{} GetOrganizationApplianceSecurityEvents(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceSecurityEvents(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceSecurityEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityEvents`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityIntrusion

> InlineResponse200286 GetOrganizationApplianceSecurityIntrusion(ctx, organizationId).Execute()

Returns all supported intrusion settings for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceSecurityIntrusion(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusion`: InlineResponse200286
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200286**](InlineResponse200286.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork

> InlineResponse200287 GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Display VPN exclusion rules for MX networks.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by network IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork`: InlineResponse200287
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter the results by network IDs | 

### Return type

[**InlineResponse200287**](InlineResponse200287.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceUplinkStatuses

> []InlineResponse200288 GetOrganizationApplianceUplinkStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()

List the uplink status of every Meraki MX and Z series appliances in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned devices will be filtered to only include these networks. (optional)
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned devices will be filtered to only include these serials. (optional)
    iccids := []string{"Inner_example"} // []string | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceUplinkStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceUplinkStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceUplinkStatuses`: []InlineResponse200288
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceUplinkStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceUplinkStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of network IDs. The returned devices will be filtered to only include these networks. | 
 **serials** | **[]string** | A list of serial numbers. The returned devices will be filtered to only include these serials. | 
 **iccids** | **[]string** | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. | 

### Return type

[**[]InlineResponse200288**](InlineResponse200288.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceUplinksNatByNetwork

> []InlineResponse200289 GetOrganizationApplianceUplinksNatByNetwork(ctx, organizationId).NetworkIds(networkIds).Interfaces(interfaces).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Fetch uplink NAT settings of each network in the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by the included set of network IDs (optional)
    interfaces := []string{"Inner_example"} // []string | Optional parameter to filter the results by the included set of interfaces (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceUplinksNatByNetwork(context.Background(), organizationId).NetworkIds(networkIds).Interfaces(interfaces).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceUplinksNatByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceUplinksNatByNetwork`: []InlineResponse200289
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceUplinksNatByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceUplinksNatByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter the results by the included set of network IDs | 
 **interfaces** | **[]string** | Optional parameter to filter the results by the included set of interfaces | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse200289**](InlineResponse200289.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceUplinksStatusesOverview

> InlineResponse200290 GetOrganizationApplianceUplinksStatusesOverview(ctx, organizationId).NetworkIds(networkIds).Execute()

Returns an overview of uplink statuses



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned devices will be filtered to only include these networks. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceUplinksStatusesOverview(context.Background(), organizationId).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceUplinksStatusesOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceUplinksStatusesOverview`: InlineResponse200290
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceUplinksStatusesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceUplinksStatusesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | A list of network IDs. The returned devices will be filtered to only include these networks. | 

### Return type

[**InlineResponse200290**](InlineResponse200290.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceUplinksUsageByNetwork

> []InlineResponse200291 GetOrganizationApplianceUplinksUsageByNetwork(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Get the sent and received bytes for each uplink of all MX and Z networks within an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceUplinksUsageByNetwork(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceUplinksUsageByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceUplinksUsageByNetwork`: []InlineResponse200291
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceUplinksUsageByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceUplinksUsageByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200291**](InlineResponse200291.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas

> InlineResponse200292 GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(ctx, organizationId).Execute()

Get the list of available IPsec SLA policies for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: InlineResponse200292
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200292**](InlineResponse200292.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceVpnStats

> []InlineResponse200294 GetOrganizationApplianceVpnStats(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).T0(t0).T1(t1).Timespan(timespan).Execute()

Show VPN history stats for networks in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456 (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceVpnStats(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceVpnStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnStats`: []InlineResponse200294
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceVpnStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;L_3456 | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200294**](InlineResponse200294.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceVpnStatuses

> []InlineResponse200295 GetOrganizationApplianceVpnStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Show VPN status for networks in an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceVpnStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceVpnStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnStatuses`: []InlineResponse200295
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceVpnStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;L_3456 | 

### Return type

[**[]InlineResponse200295**](InlineResponse200295.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceVpnThirdPartyVPNPeers

> InlineResponse200296 GetOrganizationApplianceVpnThirdPartyVPNPeers(ctx, organizationId).Execute()

Return the third party VPN peers for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceVpnThirdPartyVPNPeers(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceVpnThirdPartyVPNPeers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnThirdPartyVPNPeers`: InlineResponse200296
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceVpnThirdPartyVPNPeers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnThirdPartyVPNPeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200296**](InlineResponse200296.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceVpnVpnFirewallRules

> InlineResponse20073 GetOrganizationApplianceVpnVpnFirewallRules(ctx, organizationId).Execute()

Return the firewall rules for an organization's site-to-site VPN



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationApplianceVpnVpnFirewallRules(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationApplianceVpnVpnFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnVpnFirewallRules`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationApplianceVpnVpnFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnVpnFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments

> InlineResponse200408 GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments(ctx, organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List appliance VLAN policy assignments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    assignmentIds := []string{"Inner_example"} // []string | Filter assignments by assignment IDs (optional)
    policyIds := []string{"Inner_example"} // []string | Filter assignments by policy IDs (optional)
    interfaceIds := []string{"Inner_example"} // []string | Filter assignments by interface IDs (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments(context.Background(), organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments`: InlineResponse200408
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignmentIds** | **[]string** | Filter assignments by assignment IDs | 
 **policyIds** | **[]string** | Filter assignments by policy IDs | 
 **interfaceIds** | **[]string** | Filter assignments by interface IDs | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200408**](InlineResponse200408.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan

> InlineResponse200409 GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan(ctx, organizationId).Search(search).VlanIds(vlanIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List policies by appliance VLANs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    search := "search_example" // string | Search term for filtering policies (optional)
    vlanIds := []string{"Inner_example"} // []string | Filter by VLAN IDs (optional)
    interfaceIds := []string{"Inner_example"} // []string | Filter by interface IDs (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan(context.Background(), organizationId).Search(search).VlanIds(vlanIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan`: InlineResponse200409
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Search term for filtering policies | 
 **vlanIds** | **[]string** | Filter by VLAN IDs | 
 **interfaceIds** | **[]string** | Filter by interface IDs | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200409**](InlineResponse200409.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProtectionNetworkApplianceUmbrella

> InlineResponse200104 ProtectionNetworkApplianceUmbrella(ctx, networkId).ProtectionNetworkApplianceUmbrella(protectionNetworkApplianceUmbrella).Execute()

Enable or disable umbrella protection for an appliance network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    protectionNetworkApplianceUmbrella := *openapiclient.NewInlineObject90(false) // InlineObject90 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.ProtectionNetworkApplianceUmbrella(context.Background(), networkId).ProtectionNetworkApplianceUmbrella(protectionNetworkApplianceUmbrella).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.ProtectionNetworkApplianceUmbrella``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProtectionNetworkApplianceUmbrella`: InlineResponse200104
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.ProtectionNetworkApplianceUmbrella`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProtectionNetworkApplianceUmbrellaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **protectionNetworkApplianceUmbrella** | [**InlineObject90**](InlineObject90.md) |  | 

### Return type

[**InlineResponse200104**](InlineResponse200104.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNetworkApplianceUmbrellaPolicies

> RemoveNetworkApplianceUmbrellaPolicies(ctx, networkId).RemoveNetworkApplianceUmbrellaPolicies(removeNetworkApplianceUmbrellaPolicies).Execute()

Remove one Cisco Umbrella DNS security policy from an MX network by policy ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    removeNetworkApplianceUmbrellaPolicies := *openapiclient.NewInlineObject89(*openapiclient.NewNetworksNetworkIdApplianceUmbrellaPoliciesRemovePolicy("Id_example")) // InlineObject89 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.RemoveNetworkApplianceUmbrellaPolicies(context.Background(), networkId).RemoveNetworkApplianceUmbrellaPolicies(removeNetworkApplianceUmbrellaPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.RemoveNetworkApplianceUmbrellaPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNetworkApplianceUmbrellaPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeNetworkApplianceUmbrellaPolicies** | [**InlineObject89**](InlineObject89.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans

> InlineResponse200406 RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(ctx, organizationId).RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(removeOrganizationPoliciesGlobalGroupPoliciesApplianceVlans).Execute()

Remove VLANs from a policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    removeOrganizationPoliciesGlobalGroupPoliciesApplianceVlans := *openapiclient.NewInlineObject328(*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansRemovePolicy("Id_example"), []openapiclient.OrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansAssignVlans{*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansAssignVlans("InterfaceId_example")}) // InlineObject328 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(context.Background(), organizationId).RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(removeOrganizationPoliciesGlobalGroupPoliciesApplianceVlans).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans`: InlineResponse200406
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeOrganizationPoliciesGlobalGroupPoliciesApplianceVlans** | [**InlineObject328**](InlineObject328.md) |  | 

### Return type

[**InlineResponse200406**](InlineResponse200406.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwapNetworkApplianceWarmSpare

> InlineResponse200111 SwapNetworkApplianceWarmSpare(ctx, networkId).Execute()

Swap MX primary and warm spare appliances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.SwapNetworkApplianceWarmSpare(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.SwapNetworkApplianceWarmSpare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwapNetworkApplianceWarmSpare`: InlineResponse200111
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.SwapNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwapNetworkApplianceWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200111**](InlineResponse200111.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceApplianceRadioSettings

> InlineResponse20012 UpdateDeviceApplianceRadioSettings(ctx, serial).UpdateDeviceApplianceRadioSettings(updateDeviceApplianceRadioSettings).Execute()

Update the radio settings of an appliance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Serial
    updateDeviceApplianceRadioSettings := *openapiclient.NewInlineObject5() // InlineObject5 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateDeviceApplianceRadioSettings(context.Background(), serial).UpdateDeviceApplianceRadioSettings(updateDeviceApplianceRadioSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateDeviceApplianceRadioSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceApplianceRadioSettings`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateDeviceApplianceRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceApplianceRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceApplianceRadioSettings** | [**InlineObject5**](InlineObject5.md) |  | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceApplianceUplinksSettings

> InlineResponse20013 UpdateDeviceApplianceUplinksSettings(ctx, serial).UpdateDeviceApplianceUplinksSettings(updateDeviceApplianceUplinksSettings).Execute()

Update the uplink settings for a secure router or security appliance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | Serial
    updateDeviceApplianceUplinksSettings := *openapiclient.NewInlineObject6(*openapiclient.NewDevicesSerialApplianceUplinksSettingsInterfaces()) // InlineObject6 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateDeviceApplianceUplinksSettings(context.Background(), serial).UpdateDeviceApplianceUplinksSettings(updateDeviceApplianceUplinksSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateDeviceApplianceUplinksSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceApplianceUplinksSettings`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateDeviceApplianceUplinksSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceApplianceUplinksSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceApplianceUplinksSettings** | [**InlineObject6**](InlineObject6.md) |  | 

### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceConnectivityMonitoringDestinations

> InlineResponse20069 UpdateNetworkApplianceConnectivityMonitoringDestinations(ctx, networkId).UpdateNetworkApplianceConnectivityMonitoringDestinations(updateNetworkApplianceConnectivityMonitoringDestinations).Execute()

Update the connectivity testing destinations for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceConnectivityMonitoringDestinations := *openapiclient.NewInlineObject50() // InlineObject50 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceConnectivityMonitoringDestinations(context.Background(), networkId).UpdateNetworkApplianceConnectivityMonitoringDestinations(updateNetworkApplianceConnectivityMonitoringDestinations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceConnectivityMonitoringDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceConnectivityMonitoringDestinations`: InlineResponse20069
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceConnectivityMonitoringDestinations** | [**InlineObject50**](InlineObject50.md) |  | 

### Return type

[**InlineResponse20069**](InlineResponse20069.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceContentFiltering

> InlineResponse20070 UpdateNetworkApplianceContentFiltering(ctx, networkId).UpdateNetworkApplianceContentFiltering(updateNetworkApplianceContentFiltering).Execute()

Update the content filtering settings for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceContentFiltering := *openapiclient.NewInlineObject51() // InlineObject51 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceContentFiltering(context.Background(), networkId).UpdateNetworkApplianceContentFiltering(updateNetworkApplianceContentFiltering).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceContentFiltering``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceContentFiltering`: InlineResponse20070
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceContentFiltering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceContentFilteringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceContentFiltering** | [**InlineObject51**](InlineObject51.md) |  | 

### Return type

[**InlineResponse20070**](InlineResponse20070.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceDevicesRedundancy

> InlineResponse20072 UpdateNetworkApplianceDevicesRedundancy(ctx, networkId).UpdateNetworkApplianceDevicesRedundancy(updateNetworkApplianceDevicesRedundancy).Execute()

Update MX warm spare settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceDevicesRedundancy := *openapiclient.NewInlineObject52(false) // InlineObject52 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceDevicesRedundancy(context.Background(), networkId).UpdateNetworkApplianceDevicesRedundancy(updateNetworkApplianceDevicesRedundancy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceDevicesRedundancy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceDevicesRedundancy`: InlineResponse20072
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceDevicesRedundancy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceDevicesRedundancyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceDevicesRedundancy** | [**InlineObject52**](InlineObject52.md) |  | 

### Return type

[**InlineResponse20072**](InlineResponse20072.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallCellularFirewallRules

> InlineResponse20073 UpdateNetworkApplianceFirewallCellularFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallCellularFirewallRules(updateNetworkApplianceFirewallCellularFirewallRules).Execute()

Update the cellular firewall rules of an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceFirewallCellularFirewallRules := *openapiclient.NewInlineObject53() // InlineObject53 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallCellularFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallCellularFirewallRules(updateNetworkApplianceFirewallCellularFirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallCellularFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallCellularFirewallRules`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallCellularFirewallRules** | [**InlineObject53**](InlineObject53.md) |  | 

### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallFirewalledService

> InlineResponse20074 UpdateNetworkApplianceFirewallFirewalledService(ctx, networkId, service).UpdateNetworkApplianceFirewallFirewalledService(updateNetworkApplianceFirewallFirewalledService).Execute()

Updates the accessibility settings for the given service ('ICMP', 'web', or 'SNMP')



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    service := "service_example" // string | Service
    updateNetworkApplianceFirewallFirewalledService := *openapiclient.NewInlineObject54("Access_example") // InlineObject54 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).UpdateNetworkApplianceFirewallFirewalledService(updateNetworkApplianceFirewallFirewalledService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallFirewalledService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallFirewalledService`: InlineResponse20074
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallFirewalledService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**service** | **string** | Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallFirewalledServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceFirewallFirewalledService** | [**InlineObject54**](InlineObject54.md) |  | 

### Return type

[**InlineResponse20074**](InlineResponse20074.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallInboundCellularFirewallRules

> InlineResponse20073 UpdateNetworkApplianceFirewallInboundCellularFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallInboundCellularFirewallRules(updateNetworkApplianceFirewallInboundCellularFirewallRules).Execute()

Update the inbound cellular firewall rules of an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceFirewallInboundCellularFirewallRules := *openapiclient.NewInlineObject55() // InlineObject55 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallInboundCellularFirewallRules(updateNetworkApplianceFirewallInboundCellularFirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallInboundCellularFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallInboundCellularFirewallRules`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallInboundCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallInboundCellularFirewallRules** | [**InlineObject55**](InlineObject55.md) |  | 

### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallInboundFirewallRules

> InlineResponse20076 UpdateNetworkApplianceFirewallInboundFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallInboundFirewallRules(updateNetworkApplianceFirewallInboundFirewallRules).Execute()

Update the inbound firewall rules of an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceFirewallInboundFirewallRules := *openapiclient.NewInlineObject56() // InlineObject56 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallInboundFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallInboundFirewallRules(updateNetworkApplianceFirewallInboundFirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallInboundFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallInboundFirewallRules`: InlineResponse20076
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallInboundFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallInboundFirewallRules** | [**InlineObject56**](InlineObject56.md) |  | 

### Return type

[**InlineResponse20076**](InlineResponse20076.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallL3FirewallRules

> InlineResponse20073 UpdateNetworkApplianceFirewallL3FirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallL3FirewallRules(updateNetworkApplianceFirewallL3FirewallRules).Execute()

Update the L3 firewall rules of an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceFirewallL3FirewallRules := *openapiclient.NewInlineObject57() // InlineObject57 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallL3FirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallL3FirewallRules(updateNetworkApplianceFirewallL3FirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallL3FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallL3FirewallRules`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallL3FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallL3FirewallRules** | [**InlineObject57**](InlineObject57.md) |  | 

### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallL7FirewallRules

> InlineResponse20077 UpdateNetworkApplianceFirewallL7FirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallL7FirewallRules(updateNetworkApplianceFirewallL7FirewallRules).Execute()

Update the MX L7 firewall rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceFirewallL7FirewallRules := *openapiclient.NewInlineObject58() // InlineObject58 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallL7FirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallL7FirewallRules(updateNetworkApplianceFirewallL7FirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallL7FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallL7FirewallRules`: InlineResponse20077
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallL7FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallL7FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallL7FirewallRules** | [**InlineObject58**](InlineObject58.md) |  | 

### Return type

[**InlineResponse20077**](InlineResponse20077.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallMulticastForwarding

> InlineResponse20079 UpdateNetworkApplianceFirewallMulticastForwarding(ctx, networkId).UpdateNetworkApplianceFirewallMulticastForwarding(updateNetworkApplianceFirewallMulticastForwarding).Execute()

Update static multicast forward rules for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceFirewallMulticastForwarding := *openapiclient.NewInlineObject59([]openapiclient.NetworksNetworkIdApplianceFirewallMulticastForwardingRules{*openapiclient.NewNetworksNetworkIdApplianceFirewallMulticastForwardingRules("Description_example", "Address_example", []string{"VlanIds_example"})}) // InlineObject59 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallMulticastForwarding(context.Background(), networkId).UpdateNetworkApplianceFirewallMulticastForwarding(updateNetworkApplianceFirewallMulticastForwarding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallMulticastForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallMulticastForwarding`: InlineResponse20079
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallMulticastForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallMulticastForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallMulticastForwarding** | [**InlineObject59**](InlineObject59.md) |  | 

### Return type

[**InlineResponse20079**](InlineResponse20079.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallOneToManyNatRules

> InlineResponse20080 UpdateNetworkApplianceFirewallOneToManyNatRules(ctx, networkId).UpdateNetworkApplianceFirewallOneToManyNatRules(updateNetworkApplianceFirewallOneToManyNatRules).Execute()

Set the 1:Many NAT mapping rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceFirewallOneToManyNatRules := *openapiclient.NewInlineObject60([]openapiclient.NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules{*openapiclient.NewNetworksNetworkIdApplianceFirewallOneToManyNatRulesRules("PublicIp_example", "Uplink_example", []openapiclient.InlineResponse20080PortRules{*openapiclient.NewInlineResponse20080PortRules()})}) // InlineObject60 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallOneToManyNatRules(context.Background(), networkId).UpdateNetworkApplianceFirewallOneToManyNatRules(updateNetworkApplianceFirewallOneToManyNatRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallOneToManyNatRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallOneToManyNatRules`: InlineResponse20080
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallOneToManyNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallOneToManyNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallOneToManyNatRules** | [**InlineObject60**](InlineObject60.md) |  | 

### Return type

[**InlineResponse20080**](InlineResponse20080.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallOneToOneNatRules

> InlineResponse20081 UpdateNetworkApplianceFirewallOneToOneNatRules(ctx, networkId).UpdateNetworkApplianceFirewallOneToOneNatRules(updateNetworkApplianceFirewallOneToOneNatRules).Execute()

Set the 1:1 NAT mapping rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceFirewallOneToOneNatRules := *openapiclient.NewInlineObject61([]openapiclient.NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules{*openapiclient.NewNetworksNetworkIdApplianceFirewallOneToOneNatRulesRules("LanIp_example")}) // InlineObject61 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallOneToOneNatRules(context.Background(), networkId).UpdateNetworkApplianceFirewallOneToOneNatRules(updateNetworkApplianceFirewallOneToOneNatRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallOneToOneNatRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallOneToOneNatRules`: InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallOneToOneNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallOneToOneNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallOneToOneNatRules** | [**InlineObject61**](InlineObject61.md) |  | 

### Return type

[**InlineResponse20081**](InlineResponse20081.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallPortForwardingRules

> InlineResponse20082 UpdateNetworkApplianceFirewallPortForwardingRules(ctx, networkId).UpdateNetworkApplianceFirewallPortForwardingRules(updateNetworkApplianceFirewallPortForwardingRules).Execute()

Update the port forwarding rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceFirewallPortForwardingRules := *openapiclient.NewInlineObject62([]openapiclient.NetworksNetworkIdApplianceFirewallPortForwardingRulesRules{*openapiclient.NewNetworksNetworkIdApplianceFirewallPortForwardingRulesRules("LanIp_example", "PublicPort_example", "LocalPort_example", []string{"AllowedIps_example"}, "Protocol_example")}) // InlineObject62 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).UpdateNetworkApplianceFirewallPortForwardingRules(updateNetworkApplianceFirewallPortForwardingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallPortForwardingRules`: InlineResponse20082
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallPortForwardingRules** | [**InlineObject62**](InlineObject62.md) |  | 

### Return type

[**InlineResponse20082**](InlineResponse20082.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallSettings

> InlineResponse20083 UpdateNetworkApplianceFirewallSettings(ctx, networkId).UpdateNetworkApplianceFirewallSettings(updateNetworkApplianceFirewallSettings).Execute()

Update the firewall settings for this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceFirewallSettings := *openapiclient.NewInlineObject63() // InlineObject63 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceFirewallSettings(context.Background(), networkId).UpdateNetworkApplianceFirewallSettings(updateNetworkApplianceFirewallSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceFirewallSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallSettings`: InlineResponse20083
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceFirewallSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallSettings** | [**InlineObject63**](InlineObject63.md) |  | 

### Return type

[**InlineResponse20083**](InlineResponse20083.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceInterfacesL3

> InlineResponse20116 UpdateNetworkApplianceInterfacesL3(ctx, networkId, interfaceId).UpdateNetworkApplianceInterfacesL3(updateNetworkApplianceInterfacesL3).Execute()

Update wired L3 interface



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    interfaceId := "interfaceId_example" // string | Interface ID
    updateNetworkApplianceInterfacesL3 := *openapiclient.NewInlineObject65() // InlineObject65 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceInterfacesL3(context.Background(), networkId, interfaceId).UpdateNetworkApplianceInterfacesL3(updateNetworkApplianceInterfacesL3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceInterfacesL3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceInterfacesL3`: InlineResponse20116
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceInterfacesL3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceInterfacesL3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceInterfacesL3** | [**InlineObject65**](InlineObject65.md) |  | 

### Return type

[**InlineResponse20116**](InlineResponse20116.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAppliancePort

> InlineResponse20084 UpdateNetworkAppliancePort(ctx, networkId, portId).UpdateNetworkAppliancePort(updateNetworkAppliancePort).Execute()

Update the per-port VLAN settings for a single secure router or security appliance port.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    portId := "portId_example" // string | Port ID
    updateNetworkAppliancePort := *openapiclient.NewInlineObject66() // InlineObject66 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkAppliancePort(context.Background(), networkId, portId).UpdateNetworkAppliancePort(updateNetworkAppliancePort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkAppliancePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkAppliancePort`: InlineResponse20084
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkAppliancePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAppliancePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkAppliancePort** | [**InlineObject66**](InlineObject66.md) |  | 

### Return type

[**InlineResponse20084**](InlineResponse20084.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAppliancePrefixesDelegatedStatic

> InlineResponse20085 UpdateNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).UpdateNetworkAppliancePrefixesDelegatedStatic(updateNetworkAppliancePrefixesDelegatedStatic).Execute()

Update a static delegated prefix from a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID
    updateNetworkAppliancePrefixesDelegatedStatic := *openapiclient.NewInlineObject68() // InlineObject68 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).UpdateNetworkAppliancePrefixesDelegatedStatic(updateNetworkAppliancePrefixesDelegatedStatic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkAppliancePrefixesDelegatedStatic`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkAppliancePrefixesDelegatedStatic** | [**InlineObject68**](InlineObject68.md) |  | 

### Return type

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceRfProfile

> InlineResponse20086Assigned UpdateNetworkApplianceRfProfile(ctx, networkId, rfProfileId).UpdateNetworkApplianceRfProfile(updateNetworkApplianceRfProfile).Execute()

Updates specified RF profile for this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    rfProfileId := "rfProfileId_example" // string | Rf profile ID
    updateNetworkApplianceRfProfile := *openapiclient.NewInlineObject70() // InlineObject70 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).UpdateNetworkApplianceRfProfile(updateNetworkApplianceRfProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceRfProfile`: InlineResponse20086Assigned
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceRfProfile** | [**InlineObject70**](InlineObject70.md) |  | 

### Return type

[**InlineResponse20086Assigned**](InlineResponse20086Assigned.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSdwanInternetPolicies

> InlineResponse20087 UpdateNetworkApplianceSdwanInternetPolicies(ctx, networkId).UpdateNetworkApplianceSdwanInternetPolicies(updateNetworkApplianceSdwanInternetPolicies).Execute()

Update SDWAN internet traffic preferences for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceSdwanInternetPolicies := *openapiclient.NewInlineObject71() // InlineObject71 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceSdwanInternetPolicies(context.Background(), networkId).UpdateNetworkApplianceSdwanInternetPolicies(updateNetworkApplianceSdwanInternetPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceSdwanInternetPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSdwanInternetPolicies`: InlineResponse20087
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceSdwanInternetPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSdwanInternetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSdwanInternetPolicies** | [**InlineObject71**](InlineObject71.md) |  | 

### Return type

[**InlineResponse20087**](InlineResponse20087.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSecurityIntrusion

> InlineResponse20089 UpdateNetworkApplianceSecurityIntrusion(ctx, networkId).UpdateNetworkApplianceSecurityIntrusion(updateNetworkApplianceSecurityIntrusion).Execute()

Set the supported intrusion settings for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceSecurityIntrusion := *openapiclient.NewInlineObject72() // InlineObject72 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceSecurityIntrusion(context.Background(), networkId).UpdateNetworkApplianceSecurityIntrusion(updateNetworkApplianceSecurityIntrusion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSecurityIntrusion`: InlineResponse20089
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSecurityIntrusion** | [**InlineObject72**](InlineObject72.md) |  | 

### Return type

[**InlineResponse20089**](InlineResponse20089.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSecurityMalware

> InlineResponse20090 UpdateNetworkApplianceSecurityMalware(ctx, networkId).UpdateNetworkApplianceSecurityMalware(updateNetworkApplianceSecurityMalware).Execute()

Set the supported malware settings for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceSecurityMalware := *openapiclient.NewInlineObject73("Mode_example") // InlineObject73 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceSecurityMalware(context.Background(), networkId).UpdateNetworkApplianceSecurityMalware(updateNetworkApplianceSecurityMalware).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceSecurityMalware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSecurityMalware`: InlineResponse20090
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceSecurityMalware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSecurityMalwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSecurityMalware** | [**InlineObject73**](InlineObject73.md) |  | 

### Return type

[**InlineResponse20090**](InlineResponse20090.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSettings

> InlineResponse20091 UpdateNetworkApplianceSettings(ctx, networkId).UpdateNetworkApplianceSettings(updateNetworkApplianceSettings).Execute()

Update the appliance settings for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceSettings := *openapiclient.NewInlineObject74() // InlineObject74 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceSettings(context.Background(), networkId).UpdateNetworkApplianceSettings(updateNetworkApplianceSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSettings`: InlineResponse20091
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSettings** | [**InlineObject74**](InlineObject74.md) |  | 

### Return type

[**InlineResponse20091**](InlineResponse20091.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSingleLan

> InlineResponse20092 UpdateNetworkApplianceSingleLan(ctx, networkId).UpdateNetworkApplianceSingleLan(updateNetworkApplianceSingleLan).Execute()

Update single LAN configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceSingleLan := *openapiclient.NewInlineObject75() // InlineObject75 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceSingleLan(context.Background(), networkId).UpdateNetworkApplianceSingleLan(updateNetworkApplianceSingleLan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceSingleLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSingleLan`: InlineResponse20092
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceSingleLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSingleLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSingleLan** | [**InlineObject75**](InlineObject75.md) |  | 

### Return type

[**InlineResponse20092**](InlineResponse20092.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSsid

> InlineResponse20093 UpdateNetworkApplianceSsid(ctx, networkId, number).UpdateNetworkApplianceSsid(updateNetworkApplianceSsid).Execute()

Update the attributes of an MX SSID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    number := "number_example" // string | Number
    updateNetworkApplianceSsid := *openapiclient.NewInlineObject76() // InlineObject76 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceSsid(context.Background(), networkId, number).UpdateNetworkApplianceSsid(updateNetworkApplianceSsid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceSsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSsid`: InlineResponse20093
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceSsid** | [**InlineObject76**](InlineObject76.md) |  | 

### Return type

[**InlineResponse20093**](InlineResponse20093.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceStaticRoute

> InlineResponse20094 UpdateNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).UpdateNetworkApplianceStaticRoute(updateNetworkApplianceStaticRoute).Execute()

Update a static route for an MX or teleworker network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    staticRouteId := "staticRouteId_example" // string | Static route ID
    updateNetworkApplianceStaticRoute := *openapiclient.NewInlineObject78() // InlineObject78 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).UpdateNetworkApplianceStaticRoute(updateNetworkApplianceStaticRoute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceStaticRoute`: InlineResponse20094
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceStaticRoute** | [**InlineObject78**](InlineObject78.md) |  | 

### Return type

[**InlineResponse20094**](InlineResponse20094.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShaping

> InlineResponse20095 UpdateNetworkApplianceTrafficShaping(ctx, networkId).UpdateNetworkApplianceTrafficShaping(updateNetworkApplianceTrafficShaping).Execute()

Update the traffic shaping settings for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceTrafficShaping := *openapiclient.NewInlineObject79(*openapiclient.NewInlineResponse20095GlobalBandwidthLimits(int32(123), int32(123))) // InlineObject79 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceTrafficShaping(context.Background(), networkId).UpdateNetworkApplianceTrafficShaping(updateNetworkApplianceTrafficShaping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceTrafficShaping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShaping`: InlineResponse20095
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceTrafficShaping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShaping** | [**InlineObject79**](InlineObject79.md) |  | 

### Return type

[**InlineResponse20095**](InlineResponse20095.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingCustomPerformanceClass

> InlineResponse20096 UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(updateNetworkApplianceTrafficShapingCustomPerformanceClass).Execute()

Update a custom performance class for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID
    updateNetworkApplianceTrafficShapingCustomPerformanceClass := *openapiclient.NewInlineObject81() // InlineObject81 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(updateNetworkApplianceTrafficShapingCustomPerformanceClass).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingCustomPerformanceClass`: InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceTrafficShapingCustomPerformanceClass** | [**InlineObject81**](InlineObject81.md) |  | 

### Return type

[**InlineResponse20096**](InlineResponse20096.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingRules

> InlineResponse20097 UpdateNetworkApplianceTrafficShapingRules(ctx, networkId).UpdateNetworkApplianceTrafficShapingRules(updateNetworkApplianceTrafficShapingRules).Execute()

Update the traffic shaping settings rules for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceTrafficShapingRules := *openapiclient.NewInlineObject82() // InlineObject82 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceTrafficShapingRules(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingRules(updateNetworkApplianceTrafficShapingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingRules`: InlineResponse20097
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingRules** | [**InlineObject82**](InlineObject82.md) |  | 

### Return type

[**InlineResponse20097**](InlineResponse20097.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingUplinkBandwidth

> InlineResponse20098 UpdateNetworkApplianceTrafficShapingUplinkBandwidth(ctx, networkId).UpdateNetworkApplianceTrafficShapingUplinkBandwidth(updateNetworkApplianceTrafficShapingUplinkBandwidth).Execute()

Updates the uplink bandwidth settings for your MX network.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceTrafficShapingUplinkBandwidth := *openapiclient.NewInlineObject83() // InlineObject83 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkBandwidth(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingUplinkBandwidth(updateNetworkApplianceTrafficShapingUplinkBandwidth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkBandwidth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingUplinkBandwidth`: InlineResponse20098
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkBandwidth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingUplinkBandwidth** | [**InlineObject83**](InlineObject83.md) |  | 

### Return type

[**InlineResponse20098**](InlineResponse20098.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingUplinkSelection

> InlineResponse20099 UpdateNetworkApplianceTrafficShapingUplinkSelection(ctx, networkId).UpdateNetworkApplianceTrafficShapingUplinkSelection(updateNetworkApplianceTrafficShapingUplinkSelection).Execute()

Update uplink selection settings for an MX network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceTrafficShapingUplinkSelection := *openapiclient.NewInlineObject84() // InlineObject84 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkSelection(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingUplinkSelection(updateNetworkApplianceTrafficShapingUplinkSelection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingUplinkSelection`: InlineResponse20099
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceTrafficShapingUplinkSelection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingUplinkSelection** | [**InlineObject84**](InlineObject84.md) |  | 

### Return type

[**InlineResponse20099**](InlineResponse20099.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingVpnExclusions

> InlineResponse200100 UpdateNetworkApplianceTrafficShapingVpnExclusions(ctx, networkId).UpdateNetworkApplianceTrafficShapingVpnExclusions(updateNetworkApplianceTrafficShapingVpnExclusions).Execute()

Update VPN exclusion rules for an MX network.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceTrafficShapingVpnExclusions := *openapiclient.NewInlineObject85() // InlineObject85 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceTrafficShapingVpnExclusions(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingVpnExclusions(updateNetworkApplianceTrafficShapingVpnExclusions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceTrafficShapingVpnExclusions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingVpnExclusions`: InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceTrafficShapingVpnExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingVpnExclusions** | [**InlineObject85**](InlineObject85.md) |  | 

### Return type

[**InlineResponse200100**](InlineResponse200100.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceUplinksNat

> InlineResponse200105 UpdateNetworkApplianceUplinksNat(ctx, networkId).UpdateNetworkApplianceUplinksNat(updateNetworkApplianceUplinksNat).Execute()

Update uplink NAT settings of the specified network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceUplinksNat := *openapiclient.NewInlineObject91([]openapiclient.NetworksNetworkIdApplianceUplinksNatUplinks{*openapiclient.NewNetworksNetworkIdApplianceUplinksNatUplinks("Interface_example", *openapiclient.NewNetworksNetworkIdApplianceUplinksNatNat(false))}) // InlineObject91 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceUplinksNat(context.Background(), networkId).UpdateNetworkApplianceUplinksNat(updateNetworkApplianceUplinksNat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceUplinksNat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceUplinksNat`: InlineResponse200105
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceUplinksNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceUplinksNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceUplinksNat** | [**InlineObject91**](InlineObject91.md) |  | 

### Return type

[**InlineResponse200105**](InlineResponse200105.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVlan

> InlineResponse200107 UpdateNetworkApplianceVlan(ctx, networkId, vlanId).UpdateNetworkApplianceVlan(updateNetworkApplianceVlan).Execute()

Update a VLAN



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    vlanId := "vlanId_example" // string | Vlan ID
    updateNetworkApplianceVlan := *openapiclient.NewInlineObject94() // InlineObject94 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceVlan(context.Background(), networkId, vlanId).UpdateNetworkApplianceVlan(updateNetworkApplianceVlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVlan`: InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceVlan** | [**InlineObject94**](InlineObject94.md) |  | 

### Return type

[**InlineResponse200107**](InlineResponse200107.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVlansSettings

> InlineResponse200108 UpdateNetworkApplianceVlansSettings(ctx, networkId).UpdateNetworkApplianceVlansSettings(updateNetworkApplianceVlansSettings).Execute()

Enable/Disable VLANs for the given network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceVlansSettings := *openapiclient.NewInlineObject93() // InlineObject93 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceVlansSettings(context.Background(), networkId).UpdateNetworkApplianceVlansSettings(updateNetworkApplianceVlansSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceVlansSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVlansSettings`: InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceVlansSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVlansSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVlansSettings** | [**InlineObject93**](InlineObject93.md) |  | 

### Return type

[**InlineResponse200108**](InlineResponse200108.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVpnBgp

> InlineResponse200109 UpdateNetworkApplianceVpnBgp(ctx, networkId).UpdateNetworkApplianceVpnBgp(updateNetworkApplianceVpnBgp).Execute()

Update a Hub BGP Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceVpnBgp := *openapiclient.NewInlineObject95(false) // InlineObject95 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceVpnBgp(context.Background(), networkId).UpdateNetworkApplianceVpnBgp(updateNetworkApplianceVpnBgp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceVpnBgp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVpnBgp`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceVpnBgp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVpnBgpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVpnBgp** | [**InlineObject95**](InlineObject95.md) |  | 

### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVpnSiteToSiteVpn

> InlineResponse200110 UpdateNetworkApplianceVpnSiteToSiteVpn(ctx, networkId).UpdateNetworkApplianceVpnSiteToSiteVpn(updateNetworkApplianceVpnSiteToSiteVpn).Execute()

Update the site-to-site VPN settings of a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceVpnSiteToSiteVpn := *openapiclient.NewInlineObject96("Mode_example") // InlineObject96 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceVpnSiteToSiteVpn(context.Background(), networkId).UpdateNetworkApplianceVpnSiteToSiteVpn(updateNetworkApplianceVpnSiteToSiteVpn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceVpnSiteToSiteVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVpnSiteToSiteVpn`: InlineResponse200110
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceVpnSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVpnSiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVpnSiteToSiteVpn** | [**InlineObject96**](InlineObject96.md) |  | 

### Return type

[**InlineResponse200110**](InlineResponse200110.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceWarmSpare

> InlineResponse200111 UpdateNetworkApplianceWarmSpare(ctx, networkId).UpdateNetworkApplianceWarmSpare(updateNetworkApplianceWarmSpare).Execute()

Update MX warm spare settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceWarmSpare := *openapiclient.NewInlineObject97(false) // InlineObject97 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateNetworkApplianceWarmSpare(context.Background(), networkId).UpdateNetworkApplianceWarmSpare(updateNetworkApplianceWarmSpare).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateNetworkApplianceWarmSpare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceWarmSpare`: InlineResponse200111
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceWarmSpare** | [**InlineObject97**](InlineObject97.md) |  | 

### Return type

[**InlineResponse200111**](InlineResponse200111.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceDnsLocalProfile

> InlineResponse200276 UpdateOrganizationApplianceDnsLocalProfile(ctx, organizationId, profileId).UpdateOrganizationApplianceDnsLocalProfile(updateOrganizationApplianceDnsLocalProfile).Execute()

Update a local DNS profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    profileId := "profileId_example" // string | Profile ID
    updateOrganizationApplianceDnsLocalProfile := *openapiclient.NewInlineObject248("Name_example") // InlineObject248 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateOrganizationApplianceDnsLocalProfile(context.Background(), organizationId, profileId).UpdateOrganizationApplianceDnsLocalProfile(updateOrganizationApplianceDnsLocalProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateOrganizationApplianceDnsLocalProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceDnsLocalProfile`: InlineResponse200276
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateOrganizationApplianceDnsLocalProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceDnsLocalProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationApplianceDnsLocalProfile** | [**InlineObject248**](InlineObject248.md) |  | 

### Return type

[**InlineResponse200276**](InlineResponse200276.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceDnsLocalRecord

> InlineResponse200279 UpdateOrganizationApplianceDnsLocalRecord(ctx, organizationId, recordId).UpdateOrganizationApplianceDnsLocalRecord(updateOrganizationApplianceDnsLocalRecord).Execute()

Updates a local DNS record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    recordId := "recordId_example" // string | Record ID
    updateOrganizationApplianceDnsLocalRecord := *openapiclient.NewInlineObject250() // InlineObject250 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateOrganizationApplianceDnsLocalRecord(context.Background(), organizationId, recordId).UpdateOrganizationApplianceDnsLocalRecord(updateOrganizationApplianceDnsLocalRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateOrganizationApplianceDnsLocalRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceDnsLocalRecord`: InlineResponse200279
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateOrganizationApplianceDnsLocalRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**recordId** | **string** | Record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceDnsLocalRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationApplianceDnsLocalRecord** | [**InlineObject250**](InlineObject250.md) |  | 

### Return type

[**InlineResponse200279**](InlineResponse200279.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceDnsSplitProfile

> InlineResponse200280 UpdateOrganizationApplianceDnsSplitProfile(ctx, organizationId, profileId).UpdateOrganizationApplianceDnsSplitProfile(updateOrganizationApplianceDnsSplitProfile).Execute()

Update a split DNS profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    profileId := "profileId_example" // string | Profile ID
    updateOrganizationApplianceDnsSplitProfile := *openapiclient.NewInlineObject254() // InlineObject254 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateOrganizationApplianceDnsSplitProfile(context.Background(), organizationId, profileId).UpdateOrganizationApplianceDnsSplitProfile(updateOrganizationApplianceDnsSplitProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateOrganizationApplianceDnsSplitProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceDnsSplitProfile`: InlineResponse200280
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateOrganizationApplianceDnsSplitProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceDnsSplitProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationApplianceDnsSplitProfile** | [**InlineObject254**](InlineObject254.md) |  | 

### Return type

[**InlineResponse200280**](InlineResponse200280.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceRoutingVrfsSettings

> InlineResponse200285 UpdateOrganizationApplianceRoutingVrfsSettings(ctx, organizationId).UpdateOrganizationApplianceRoutingVrfsSettings(updateOrganizationApplianceRoutingVrfsSettings).Execute()

Update the VRF setting for an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    updateOrganizationApplianceRoutingVrfsSettings := *openapiclient.NewInlineObject255(false) // InlineObject255 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateOrganizationApplianceRoutingVrfsSettings(context.Background(), organizationId).UpdateOrganizationApplianceRoutingVrfsSettings(updateOrganizationApplianceRoutingVrfsSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateOrganizationApplianceRoutingVrfsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceRoutingVrfsSettings`: InlineResponse200285
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateOrganizationApplianceRoutingVrfsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceRoutingVrfsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceRoutingVrfsSettings** | [**InlineObject255**](InlineObject255.md) |  | 

### Return type

[**InlineResponse200285**](InlineResponse200285.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceSecurityIntrusion

> InlineResponse200286 UpdateOrganizationApplianceSecurityIntrusion(ctx, organizationId).UpdateOrganizationApplianceSecurityIntrusion(updateOrganizationApplianceSecurityIntrusion).Execute()

Sets supported intrusion settings for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    updateOrganizationApplianceSecurityIntrusion := *openapiclient.NewInlineObject256([]openapiclient.OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules{*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules("RuleId_example")}) // InlineObject256 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateOrganizationApplianceSecurityIntrusion(context.Background(), organizationId).UpdateOrganizationApplianceSecurityIntrusion(updateOrganizationApplianceSecurityIntrusion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateOrganizationApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceSecurityIntrusion`: InlineResponse200286
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateOrganizationApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceSecurityIntrusion** | [**InlineObject256**](InlineObject256.md) |  | 

### Return type

[**InlineResponse200286**](InlineResponse200286.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas

> InlineResponse200293 UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(ctx, organizationId).UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas).Execute()

Update the IPsec SLA policies for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas := *openapiclient.NewInlineObject257() // InlineObject257 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(context.Background(), organizationId).UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: InlineResponse200293
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas** | [**InlineObject257**](InlineObject257.md) |  | 

### Return type

[**InlineResponse200293**](InlineResponse200293.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceVpnThirdPartyVPNPeers

> InlineResponse200296 UpdateOrganizationApplianceVpnThirdPartyVPNPeers(ctx, organizationId).UpdateOrganizationApplianceVpnThirdPartyVPNPeers(updateOrganizationApplianceVpnThirdPartyVPNPeers).Execute()

Update the third party VPN peers for an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    updateOrganizationApplianceVpnThirdPartyVPNPeers := *openapiclient.NewInlineObject258([]openapiclient.OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers{*openapiclient.NewOrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers("Name_example", []string{"PrivateSubnets_example"}, "Secret_example")}) // InlineObject258 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateOrganizationApplianceVpnThirdPartyVPNPeers(context.Background(), organizationId).UpdateOrganizationApplianceVpnThirdPartyVPNPeers(updateOrganizationApplianceVpnThirdPartyVPNPeers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateOrganizationApplianceVpnThirdPartyVPNPeers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceVpnThirdPartyVPNPeers`: InlineResponse200296
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateOrganizationApplianceVpnThirdPartyVPNPeers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceVpnThirdPartyVPNPeers** | [**InlineObject258**](InlineObject258.md) |  | 

### Return type

[**InlineResponse200296**](InlineResponse200296.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceVpnVpnFirewallRules

> InlineResponse20073 UpdateOrganizationApplianceVpnVpnFirewallRules(ctx, organizationId).UpdateOrganizationApplianceVpnVpnFirewallRules(updateOrganizationApplianceVpnVpnFirewallRules).Execute()

Update the firewall rules of an organization's site-to-site VPN



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    updateOrganizationApplianceVpnVpnFirewallRules := *openapiclient.NewInlineObject259() // InlineObject259 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceApi.UpdateOrganizationApplianceVpnVpnFirewallRules(context.Background(), organizationId).UpdateOrganizationApplianceVpnVpnFirewallRules(updateOrganizationApplianceVpnVpnFirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceApi.UpdateOrganizationApplianceVpnVpnFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceVpnVpnFirewallRules`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `ApplianceApi.UpdateOrganizationApplianceVpnVpnFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceVpnVpnFirewallRules** | [**InlineObject259**](InlineObject259.md) |  | 

### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

