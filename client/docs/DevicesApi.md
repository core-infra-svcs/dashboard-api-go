# \DevicesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchNetworkFloorPlansDevicesUpdate**](DevicesApi.md#BatchNetworkFloorPlansDevicesUpdate) | **Post** /networks/{networkId}/floorPlans/devices/batchUpdate | Update floorplan assignments for a batch of devices
[**BlinkDeviceLeds**](DevicesApi.md#BlinkDeviceLeds) | **Post** /devices/{serial}/blinkLeds | Blink the LEDs on a device
[**BulkOrganizationDevicesPacketCaptureCapturesCreate**](DevicesApi.md#BulkOrganizationDevicesPacketCaptureCapturesCreate) | **Post** /organizations/{organizationId}/devices/packetCapture/captures/bulkCreate | Perform a packet capture on multiple devices and store in Meraki Cloud.
[**BulkOrganizationDevicesPacketCaptureCapturesDelete**](DevicesApi.md#BulkOrganizationDevicesPacketCaptureCapturesDelete) | **Post** /organizations/{organizationId}/devices/packetCapture/captures/bulkDelete | BulkDelete packet captures from cloud
[**BulkUpdateOrganizationDevicesDetails**](DevicesApi.md#BulkUpdateOrganizationDevicesDetails) | **Post** /organizations/{organizationId}/devices/details/bulkUpdate | Updating device details (currently only used for Catalyst devices)
[**CheckinNetworkSmDevices**](DevicesApi.md#CheckinNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/checkin | Force check-in a set of devices
[**ClaimNetworkDevices**](DevicesApi.md#ClaimNetworkDevices) | **Post** /networks/{networkId}/devices/claim | Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requests against that device to succeed)
[**CloneOrganizationSwitchDevices**](DevicesApi.md#CloneOrganizationSwitchDevices) | **Post** /organizations/{organizationId}/switch/devices/clone | Clone port-level and some switch-level configuration settings from a source switch to one or more target switches
[**CreateDeviceLiveToolsArpTable**](DevicesApi.md#CreateDeviceLiveToolsArpTable) | **Post** /devices/{serial}/liveTools/arpTable | Enqueue a job to perform a ARP table request for the device
[**CreateDeviceLiveToolsCableTest**](DevicesApi.md#CreateDeviceLiveToolsCableTest) | **Post** /devices/{serial}/liveTools/cableTest | Enqueue a job to perform a cable test for the device on the specified ports
[**CreateDeviceLiveToolsLedsBlink**](DevicesApi.md#CreateDeviceLiveToolsLedsBlink) | **Post** /devices/{serial}/liveTools/leds/blink | Enqueue a job to blink LEDs on a device
[**CreateDeviceLiveToolsMacTable**](DevicesApi.md#CreateDeviceLiveToolsMacTable) | **Post** /devices/{serial}/liveTools/macTable | Enqueue a job to request the MAC table from the device
[**CreateDeviceLiveToolsMulticastRouting**](DevicesApi.md#CreateDeviceLiveToolsMulticastRouting) | **Post** /devices/{serial}/liveTools/multicastRouting | Enqueue a job to perform a Multicast routing request for the device
[**CreateDeviceLiveToolsPing**](DevicesApi.md#CreateDeviceLiveToolsPing) | **Post** /devices/{serial}/liveTools/ping | Enqueue a job to ping a target host from the device
[**CreateDeviceLiveToolsPingDevice**](DevicesApi.md#CreateDeviceLiveToolsPingDevice) | **Post** /devices/{serial}/liveTools/pingDevice | Enqueue a job to check connectivity status to the device
[**CreateDeviceLiveToolsThroughputTest**](DevicesApi.md#CreateDeviceLiveToolsThroughputTest) | **Post** /devices/{serial}/liveTools/throughputTest | Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput
[**CreateDeviceLiveToolsWakeOnLan**](DevicesApi.md#CreateDeviceLiveToolsWakeOnLan) | **Post** /devices/{serial}/liveTools/wakeOnLan | Enqueue a job to send a Wake-on-LAN packet from the device
[**CreateOrganizationDevicesControllerMigration**](DevicesApi.md#CreateOrganizationDevicesControllerMigration) | **Post** /organizations/{organizationId}/devices/controller/migrations | Migrate devices to another controller or management mode
[**CreateOrganizationDevicesPacketCaptureCapture**](DevicesApi.md#CreateOrganizationDevicesPacketCaptureCapture) | **Post** /organizations/{organizationId}/devices/packetCapture/captures | Perform a packet capture on a device and store in Meraki Cloud
[**CreateOrganizationDevicesPacketCaptureSchedule**](DevicesApi.md#CreateOrganizationDevicesPacketCaptureSchedule) | **Post** /organizations/{organizationId}/devices/packetCapture/schedules | Create a schedule for packet capture
[**CreateOrganizationInventoryDevicesSwapsBulk**](DevicesApi.md#CreateOrganizationInventoryDevicesSwapsBulk) | **Post** /organizations/{organizationId}/inventory/devices/swaps/bulk | Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.
[**CreateOrganizationWirelessDevicesRadsecCertificatesAuthority**](DevicesApi.md#CreateOrganizationWirelessDevicesRadsecCertificatesAuthority) | **Post** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Create an organization&#39;s RADSEC device Certificate Authority (CA)
[**DeleteOrganizationDevicesPacketCaptureCapture**](DevicesApi.md#DeleteOrganizationDevicesPacketCaptureCapture) | **Delete** /organizations/{organizationId}/devices/packetCapture/captures/{captureId} | Delete a single packet capture from cloud using captureId
[**DeleteOrganizationDevicesPacketCaptureSchedule**](DevicesApi.md#DeleteOrganizationDevicesPacketCaptureSchedule) | **Delete** /organizations/{organizationId}/devices/packetCapture/schedules/{scheduleId} | Delete schedule from cloud
[**GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl**](DevicesApi.md#GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl) | **Post** /organizations/{organizationId}/devices/packetCapture/captures/{captureId}/downloadUrl/generate | Get presigned download URL for given packet capture id
[**GetDevice**](DevicesApi.md#GetDevice) | **Get** /devices/{serial} | Return a single device
[**GetDeviceCellularSims**](DevicesApi.md#GetDeviceCellularSims) | **Get** /devices/{serial}/cellular/sims | Return the SIM and APN configurations for a cellular device.
[**GetDeviceClients**](DevicesApi.md#GetDeviceClients) | **Get** /devices/{serial}/clients | List the clients of a device, up to a maximum of a month ago
[**GetDeviceLiveToolsArpTable**](DevicesApi.md#GetDeviceLiveToolsArpTable) | **Get** /devices/{serial}/liveTools/arpTable/{arpTableId} | Return an ARP table live tool job.
[**GetDeviceLiveToolsCableTest**](DevicesApi.md#GetDeviceLiveToolsCableTest) | **Get** /devices/{serial}/liveTools/cableTest/{id} | Return a cable test live tool job.
[**GetDeviceLiveToolsLedsBlink**](DevicesApi.md#GetDeviceLiveToolsLedsBlink) | **Get** /devices/{serial}/liveTools/leds/blink/{ledsBlinkId} | Return a blink LEDs job
[**GetDeviceLiveToolsMacTable**](DevicesApi.md#GetDeviceLiveToolsMacTable) | **Get** /devices/{serial}/liveTools/macTable/{macTableId} | Return a MAC table live tool job.
[**GetDeviceLiveToolsMulticastRouting**](DevicesApi.md#GetDeviceLiveToolsMulticastRouting) | **Get** /devices/{serial}/liveTools/multicastRouting/{multicastRoutingId} | Return a Multicast routing live tool job.
[**GetDeviceLiveToolsPing**](DevicesApi.md#GetDeviceLiveToolsPing) | **Get** /devices/{serial}/liveTools/ping/{id} | Return a ping job
[**GetDeviceLiveToolsPingDevice**](DevicesApi.md#GetDeviceLiveToolsPingDevice) | **Get** /devices/{serial}/liveTools/pingDevice/{id} | Return a ping device job
[**GetDeviceLiveToolsThroughputTest**](DevicesApi.md#GetDeviceLiveToolsThroughputTest) | **Get** /devices/{serial}/liveTools/throughputTest/{throughputTestId} | Return a throughput test job
[**GetDeviceLiveToolsWakeOnLan**](DevicesApi.md#GetDeviceLiveToolsWakeOnLan) | **Get** /devices/{serial}/liveTools/wakeOnLan/{wakeOnLanId} | Return a Wake-on-LAN job
[**GetDeviceLldpCdp**](DevicesApi.md#GetDeviceLldpCdp) | **Get** /devices/{serial}/lldpCdp | List LLDP and CDP information for a device
[**GetDeviceLossAndLatencyHistory**](DevicesApi.md#GetDeviceLossAndLatencyHistory) | **Get** /devices/{serial}/lossAndLatencyHistory | Get the uplink loss percentage and latency in milliseconds, and goodput in kilobits per second for MX, MG and Z devices.
[**GetDeviceManagementInterface**](DevicesApi.md#GetDeviceManagementInterface) | **Get** /devices/{serial}/managementInterface | Return the management interface settings for a device
[**GetNetworkDevices**](DevicesApi.md#GetNetworkDevices) | **Get** /networks/{networkId}/devices | List the devices in a network
[**GetNetworkSmDeviceCellularUsageHistory**](DevicesApi.md#GetNetworkSmDeviceCellularUsageHistory) | **Get** /networks/{networkId}/sm/devices/{deviceId}/cellularUsageHistory | Return the client&#39;s daily cellular data usage history
[**GetNetworkSmDeviceCerts**](DevicesApi.md#GetNetworkSmDeviceCerts) | **Get** /networks/{networkId}/sm/devices/{deviceId}/certs | List the certs on a device
[**GetNetworkSmDeviceConnectivity**](DevicesApi.md#GetNetworkSmDeviceConnectivity) | **Get** /networks/{networkId}/sm/devices/{deviceId}/connectivity | Returns historical connectivity data (whether a device is regularly checking in to Dashboard).
[**GetNetworkSmDeviceDesktopLogs**](DevicesApi.md#GetNetworkSmDeviceDesktopLogs) | **Get** /networks/{networkId}/sm/devices/{deviceId}/desktopLogs | Return historical records of various Systems Manager network connection details for desktop devices.
[**GetNetworkSmDeviceDeviceCommandLogs**](DevicesApi.md#GetNetworkSmDeviceDeviceCommandLogs) | **Get** /networks/{networkId}/sm/devices/{deviceId}/deviceCommandLogs | Return historical records of commands sent to Systems Manager devices
[**GetNetworkSmDeviceDeviceProfiles**](DevicesApi.md#GetNetworkSmDeviceDeviceProfiles) | **Get** /networks/{networkId}/sm/devices/{deviceId}/deviceProfiles | Get the installed profiles associated with a device
[**GetNetworkSmDeviceNetworkAdapters**](DevicesApi.md#GetNetworkSmDeviceNetworkAdapters) | **Get** /networks/{networkId}/sm/devices/{deviceId}/networkAdapters | List the network adapters of a device
[**GetNetworkSmDevicePerformanceHistory**](DevicesApi.md#GetNetworkSmDevicePerformanceHistory) | **Get** /networks/{networkId}/sm/devices/{deviceId}/performanceHistory | Return historical records of various Systems Manager client metrics for desktop devices.
[**GetNetworkSmDeviceRestrictions**](DevicesApi.md#GetNetworkSmDeviceRestrictions) | **Get** /networks/{networkId}/sm/devices/{deviceId}/restrictions | List the restrictions on a device
[**GetNetworkSmDeviceSecurityCenters**](DevicesApi.md#GetNetworkSmDeviceSecurityCenters) | **Get** /networks/{networkId}/sm/devices/{deviceId}/securityCenters | List the security centers on a device
[**GetNetworkSmDeviceSoftwares**](DevicesApi.md#GetNetworkSmDeviceSoftwares) | **Get** /networks/{networkId}/sm/devices/{deviceId}/softwares | Get a list of softwares associated with a device
[**GetNetworkSmDeviceWlanLists**](DevicesApi.md#GetNetworkSmDeviceWlanLists) | **Get** /networks/{networkId}/sm/devices/{deviceId}/wlanLists | List the saved SSID names on a device
[**GetNetworkSmDevices**](DevicesApi.md#GetNetworkSmDevices) | **Get** /networks/{networkId}/sm/devices | List the devices enrolled in an SM network with various specified fields and filters
[**GetNetworkWirelessDevicesConnectionStats**](DevicesApi.md#GetNetworkWirelessDevicesConnectionStats) | **Get** /networks/{networkId}/wireless/devices/connectionStats | Aggregated connectivity info for this network, grouped by node
[**GetNetworkWirelessDevicesLatencyStats**](DevicesApi.md#GetNetworkWirelessDevicesLatencyStats) | **Get** /networks/{networkId}/wireless/devices/latencyStats | Aggregated latency info for this network, grouped by node
[**GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice**](DevicesApi.md#GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice) | **Get** /organizations/{organizationId}/campusGateway/devices/uplinks/localOverrides/byDevice | Uplink overrides configured locally on Campus Gateway devices in an organization.
[**GetOrganizationDevices**](DevicesApi.md#GetOrganizationDevices) | **Get** /organizations/{organizationId}/devices | List the devices in an organization that have been assigned to a network.
[**GetOrganizationDevicesAvailabilities**](DevicesApi.md#GetOrganizationDevicesAvailabilities) | **Get** /organizations/{organizationId}/devices/availabilities | List the availability information for devices in an organization
[**GetOrganizationDevicesAvailabilitiesChangeHistory**](DevicesApi.md#GetOrganizationDevicesAvailabilitiesChangeHistory) | **Get** /organizations/{organizationId}/devices/availabilities/changeHistory | List the availability history information for devices in an organization.
[**GetOrganizationDevicesControllerMigrations**](DevicesApi.md#GetOrganizationDevicesControllerMigrations) | **Get** /organizations/{organizationId}/devices/controller/migrations | Retrieve device migration statuses in an organization
[**GetOrganizationDevicesOverviewByModel**](DevicesApi.md#GetOrganizationDevicesOverviewByModel) | **Get** /organizations/{organizationId}/devices/overview/byModel | Lists the count for each device model
[**GetOrganizationDevicesPacketCaptureCaptures**](DevicesApi.md#GetOrganizationDevicesPacketCaptureCaptures) | **Get** /organizations/{organizationId}/devices/packetCapture/captures | List Packet Captures
[**GetOrganizationDevicesPacketCaptureSchedules**](DevicesApi.md#GetOrganizationDevicesPacketCaptureSchedules) | **Get** /organizations/{organizationId}/devices/packetCapture/schedules | List the Packet Capture Schedules
[**GetOrganizationDevicesPowerModulesStatusesByDevice**](DevicesApi.md#GetOrganizationDevicesPowerModulesStatusesByDevice) | **Get** /organizations/{organizationId}/devices/powerModules/statuses/byDevice | List the most recent status information for power modules in rackmount MX and MS devices that support them
[**GetOrganizationDevicesProvisioningStatuses**](DevicesApi.md#GetOrganizationDevicesProvisioningStatuses) | **Get** /organizations/{organizationId}/devices/provisioning/statuses | List the provisioning statuses information for devices in an organization.
[**GetOrganizationDevicesStatuses**](DevicesApi.md#GetOrganizationDevicesStatuses) | **Get** /organizations/{organizationId}/devices/statuses | List the status of every Meraki device in the organization
[**GetOrganizationDevicesStatusesOverview**](DevicesApi.md#GetOrganizationDevicesStatusesOverview) | **Get** /organizations/{organizationId}/devices/statuses/overview | Return an overview of current device statuses
[**GetOrganizationDevicesSystemMemoryUsageHistoryByInterval**](DevicesApi.md#GetOrganizationDevicesSystemMemoryUsageHistoryByInterval) | **Get** /organizations/{organizationId}/devices/system/memory/usage/history/byInterval | Return the memory utilization history in kB for devices in the organization.
[**GetOrganizationDevicesUplinksAddressesByDevice**](DevicesApi.md#GetOrganizationDevicesUplinksAddressesByDevice) | **Get** /organizations/{organizationId}/devices/uplinks/addresses/byDevice | List the current uplink addresses for devices in an organization.
[**GetOrganizationDevicesUplinksLossAndLatency**](DevicesApi.md#GetOrganizationDevicesUplinksLossAndLatency) | **Get** /organizations/{organizationId}/devices/uplinksLossAndLatency | Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago
[**GetOrganizationFloorPlansAutoLocateDevices**](DevicesApi.md#GetOrganizationFloorPlansAutoLocateDevices) | **Get** /organizations/{organizationId}/floorPlans/autoLocate/devices | List auto locate details for each device in your organization
[**GetOrganizationInventoryDevice**](DevicesApi.md#GetOrganizationInventoryDevice) | **Get** /organizations/{organizationId}/inventory/devices/{serial} | Return a single device from the inventory of an organization
[**GetOrganizationInventoryDevices**](DevicesApi.md#GetOrganizationInventoryDevices) | **Get** /organizations/{organizationId}/inventory/devices | Return the device inventory for an organization
[**GetOrganizationInventoryDevicesSwapsBulk**](DevicesApi.md#GetOrganizationInventoryDevicesSwapsBulk) | **Get** /organizations/{organizationId}/inventory/devices/swaps/bulk/{id} | List of device swaps for a given request ID ({id}).
[**GetOrganizationSummaryTopDevicesByUsage**](DevicesApi.md#GetOrganizationSummaryTopDevicesByUsage) | **Get** /organizations/{organizationId}/summary/top/devices/byUsage | Return metrics for organization&#39;s top 10 devices sorted by data usage over given time range
[**GetOrganizationSummaryTopDevicesModelsByUsage**](DevicesApi.md#GetOrganizationSummaryTopDevicesModelsByUsage) | **Get** /organizations/{organizationId}/summary/top/devices/models/byUsage | Return metrics for organization&#39;s top 10 device models sorted by data usage over given time range
[**GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice**](DevicesApi.md#GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/l2/byDevice | List wireless LAN controller layer 2 interfaces in an organization
[**GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice**](DevicesApi.md#GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/l2/statuses/changeHistory/byDevice | List wireless LAN controller layer 2 interfaces history status in an organization
[**GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval**](DevicesApi.md#GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/l2/usage/history/byInterval | List wireless LAN controller layer 2 interfaces history usage in an organization
[**GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice**](DevicesApi.md#GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/l3/byDevice | List wireless LAN controller layer 3 interfaces in an organization
[**GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice**](DevicesApi.md#GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/l3/statuses/changeHistory/byDevice | List wireless LAN controller layer 3 interfaces history status in an organization
[**GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval**](DevicesApi.md#GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/l3/usage/history/byInterval | List wireless LAN controller layer 3 interfaces history usage in an organization
[**GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice**](DevicesApi.md#GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/packets/overview/byDevice | Retrieve the packet counters for the interfaces of a Wireless LAN controller
[**GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval**](DevicesApi.md#GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/usage/history/byInterval | Retrieve the traffic for the interfaces of a Wireless LAN controller
[**GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory**](DevicesApi.md#GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory) | **Get** /organizations/{organizationId}/wirelessController/devices/redundancy/failover/history | List the failover events of wireless LAN controllers in an organization
[**GetOrganizationWirelessControllerDevicesRedundancyStatuses**](DevicesApi.md#GetOrganizationWirelessControllerDevicesRedundancyStatuses) | **Get** /organizations/{organizationId}/wirelessController/devices/redundancy/statuses | List redundancy details of wireless LAN controllers in an organization
[**GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval**](DevicesApi.md#GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval) | **Get** /organizations/{organizationId}/wirelessController/devices/system/utilization/history/byInterval | List cpu utilization data of wireless LAN controllers in an organization
[**GetOrganizationWirelessDevicesChannelUtilizationByDevice**](DevicesApi.md#GetOrganizationWirelessDevicesChannelUtilizationByDevice) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/byDevice | Get average channel utilization for all bands in a network, split by AP
[**GetOrganizationWirelessDevicesChannelUtilizationByNetwork**](DevicesApi.md#GetOrganizationWirelessDevicesChannelUtilizationByNetwork) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/byNetwork | Get average channel utilization across all bands for all networks in the organization
[**GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval**](DevicesApi.md#GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/history/byDevice/byInterval | Get a time-series of average channel utilization for all bands, segmented by device.
[**GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval**](DevicesApi.md#GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/history/byNetwork/byInterval | Get a time-series of average channel utilization for all bands
[**GetOrganizationWirelessDevicesEthernetStatuses**](DevicesApi.md#GetOrganizationWirelessDevicesEthernetStatuses) | **Get** /organizations/{organizationId}/wireless/devices/ethernet/statuses | List the most recent Ethernet link speed, duplex, aggregation and power mode and status information for wireless devices.
[**GetOrganizationWirelessDevicesPacketLossByClient**](DevicesApi.md#GetOrganizationWirelessDevicesPacketLossByClient) | **Get** /organizations/{organizationId}/wireless/devices/packetLoss/byClient | Get average packet loss for the given timespan for all clients in the organization.
[**GetOrganizationWirelessDevicesPacketLossByDevice**](DevicesApi.md#GetOrganizationWirelessDevicesPacketLossByDevice) | **Get** /organizations/{organizationId}/wireless/devices/packetLoss/byDevice | Get average packet loss for the given timespan for all devices in the organization
[**GetOrganizationWirelessDevicesPacketLossByNetwork**](DevicesApi.md#GetOrganizationWirelessDevicesPacketLossByNetwork) | **Get** /organizations/{organizationId}/wireless/devices/packetLoss/byNetwork | Get average packet loss for the given timespan for all networks in the organization.
[**GetOrganizationWirelessDevicesPowerModeHistory**](DevicesApi.md#GetOrganizationWirelessDevicesPowerModeHistory) | **Get** /organizations/{organizationId}/wireless/devices/power/mode/history | Return a record of power mode changes for wireless devices in the organization
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthorities**](DevicesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthorities) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Query for details on the organization&#39;s RADSEC device Certificate Authority certificates (CAs)
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls**](DevicesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities/crls | Query for certificate revocation list (CRL) for the organization&#39;s RADSEC device Certificate Authorities (CAs).
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas**](DevicesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities/crls/deltas | Query for all delta certificate revocation list (CRL) for the organization&#39;s RADSEC device Certificate Authority (CA) with the given id.
[**GetOrganizationWirelessDevicesSystemCpuLoadHistory**](DevicesApi.md#GetOrganizationWirelessDevicesSystemCpuLoadHistory) | **Get** /organizations/{organizationId}/wireless/devices/system/cpu/load/history | Return the CPU Load history for a list of wireless devices in the organization.
[**GetOrganizationWirelessDevicesWirelessControllersByDevice**](DevicesApi.md#GetOrganizationWirelessDevicesWirelessControllersByDevice) | **Get** /organizations/{organizationId}/wireless/devices/wirelessControllers/byDevice | List of Catalyst access points information
[**GetOrganizationWirelessZigbeeDevices**](DevicesApi.md#GetOrganizationWirelessZigbeeDevices) | **Get** /organizations/{organizationId}/wireless/zigbee/devices | List the Zigbee wireless devices for an organization or the supplied network(s)
[**InstallNetworkSmDeviceApps**](DevicesApi.md#InstallNetworkSmDeviceApps) | **Post** /networks/{networkId}/sm/devices/{deviceId}/installApps | Install applications on a device
[**LockNetworkSmDevices**](DevicesApi.md#LockNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/lock | Lock a set of devices
[**ModifyNetworkSmDevicesTags**](DevicesApi.md#ModifyNetworkSmDevicesTags) | **Post** /networks/{networkId}/sm/devices/modifyTags | Add, delete, or update the tags of a set of devices
[**MoveNetworkSmDevices**](DevicesApi.md#MoveNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/move | Move a set of devices to a new network
[**RebootDevice**](DevicesApi.md#RebootDevice) | **Post** /devices/{serial}/reboot | Reboot a device
[**RebootNetworkSmDevices**](DevicesApi.md#RebootNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/reboot | Reboot a set of endpoints
[**RefreshNetworkSmDeviceDetails**](DevicesApi.md#RefreshNetworkSmDeviceDetails) | **Post** /networks/{networkId}/sm/devices/{deviceId}/refreshDetails | Refresh the details of a device
[**RemoveNetworkDevices**](DevicesApi.md#RemoveNetworkDevices) | **Post** /networks/{networkId}/devices/remove | Remove a single device
[**ReorderOrganizationDevicesPacketCaptureSchedules**](DevicesApi.md#ReorderOrganizationDevicesPacketCaptureSchedules) | **Post** /organizations/{organizationId}/devices/packetCapture/schedules/reorder | Bulk update priorities of pcap schedules
[**ShutdownNetworkSmDevices**](DevicesApi.md#ShutdownNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/shutdown | Shutdown a set of endpoints
[**StopOrganizationDevicesPacketCaptureCapture**](DevicesApi.md#StopOrganizationDevicesPacketCaptureCapture) | **Post** /organizations/{organizationId}/devices/packetCapture/captures/{captureId}/stop | Stop a specific packet capture (not supported for Catalyst devices)
[**UnenrollNetworkSmDevice**](DevicesApi.md#UnenrollNetworkSmDevice) | **Post** /networks/{networkId}/sm/devices/{deviceId}/unenroll | Unenroll a device
[**UninstallNetworkSmDeviceApps**](DevicesApi.md#UninstallNetworkSmDeviceApps) | **Post** /networks/{networkId}/sm/devices/{deviceId}/uninstallApps | Uninstall applications on a device
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Put** /devices/{serial} | Update the attributes of a device
[**UpdateDeviceCellularSims**](DevicesApi.md#UpdateDeviceCellularSims) | **Put** /devices/{serial}/cellular/sims | Updates the SIM and APN configurations for a cellular device.
[**UpdateDeviceManagementInterface**](DevicesApi.md#UpdateDeviceManagementInterface) | **Put** /devices/{serial}/managementInterface | Update the management interface settings for a device
[**UpdateNetworkSmDevicesFields**](DevicesApi.md#UpdateNetworkSmDevicesFields) | **Put** /networks/{networkId}/sm/devices/fields | Modify the fields of a device
[**UpdateOrganizationDevicesPacketCaptureSchedule**](DevicesApi.md#UpdateOrganizationDevicesPacketCaptureSchedule) | **Put** /organizations/{organizationId}/devices/packetCapture/schedules/{scheduleId} | Update a schedule for packet capture
[**UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities**](DevicesApi.md#UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities) | **Put** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Update an organization&#39;s RADSEC device Certificate Authority (CA) state
[**UpdateOrganizationWirelessZigbeeDevice**](DevicesApi.md#UpdateOrganizationWirelessZigbeeDevice) | **Put** /organizations/{organizationId}/wireless/zigbee/devices/{id} | Endpoint to update zigbee gateways
[**VmxNetworkDevicesClaim**](DevicesApi.md#VmxNetworkDevicesClaim) | **Post** /networks/{networkId}/devices/claim/vmx | Claim a vMX into a network
[**WipeNetworkSmDevices**](DevicesApi.md#WipeNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/wipe | Wipe a device



## BatchNetworkFloorPlansDevicesUpdate

> InlineResponse200108 BatchNetworkFloorPlansDevicesUpdate(ctx, networkId).BatchNetworkFloorPlansDevicesUpdate(batchNetworkFloorPlansDevicesUpdate).Execute()

Update floorplan assignments for a batch of devices



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
    batchNetworkFloorPlansDevicesUpdate := *openapiclient.NewInlineObject110([]openapiclient.NetworksNetworkIdFloorPlansDevicesBatchUpdateAssignments{*openapiclient.NewNetworksNetworkIdFloorPlansDevicesBatchUpdateAssignments("Serial_example", *openapiclient.NewNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan("Id_example"))}) // InlineObject110 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.BatchNetworkFloorPlansDevicesUpdate(context.Background(), networkId).BatchNetworkFloorPlansDevicesUpdate(batchNetworkFloorPlansDevicesUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.BatchNetworkFloorPlansDevicesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchNetworkFloorPlansDevicesUpdate`: InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.BatchNetworkFloorPlansDevicesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchNetworkFloorPlansDevicesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchNetworkFloorPlansDevicesUpdate** | [**InlineObject110**](InlineObject110.md) |  | 

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


## BlinkDeviceLeds

> InlineResponse2021 BlinkDeviceLeds(ctx, serial).BlinkDeviceLeds(blinkDeviceLeds).Execute()

Blink the LEDs on a device



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
    blinkDeviceLeds := *openapiclient.NewInlineObject6() // InlineObject6 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.BlinkDeviceLeds(context.Background(), serial).BlinkDeviceLeds(blinkDeviceLeds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.BlinkDeviceLeds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlinkDeviceLeds`: InlineResponse2021
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.BlinkDeviceLeds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlinkDeviceLedsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blinkDeviceLeds** | [**InlineObject6**](InlineObject6.md) |  | 

### Return type

[**InlineResponse2021**](InlineResponse2021.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkOrganizationDevicesPacketCaptureCapturesCreate

> InlineResponse20121 BulkOrganizationDevicesPacketCaptureCapturesCreate(ctx, organizationId).BulkOrganizationDevicesPacketCaptureCapturesCreate(bulkOrganizationDevicesPacketCaptureCapturesCreate).Execute()

Perform a packet capture on multiple devices and store in Meraki Cloud.



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
    bulkOrganizationDevicesPacketCaptureCapturesCreate := *openapiclient.NewInlineObject259([]openapiclient.OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices{*openapiclient.NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices()}, "Name_example") // InlineObject259 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.BulkOrganizationDevicesPacketCaptureCapturesCreate(context.Background(), organizationId).BulkOrganizationDevicesPacketCaptureCapturesCreate(bulkOrganizationDevicesPacketCaptureCapturesCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.BulkOrganizationDevicesPacketCaptureCapturesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkOrganizationDevicesPacketCaptureCapturesCreate`: InlineResponse20121
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.BulkOrganizationDevicesPacketCaptureCapturesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkOrganizationDevicesPacketCaptureCapturesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkOrganizationDevicesPacketCaptureCapturesCreate** | [**InlineObject259**](InlineObject259.md) |  | 

### Return type

[**InlineResponse20121**](InlineResponse20121.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkOrganizationDevicesPacketCaptureCapturesDelete

> BulkOrganizationDevicesPacketCaptureCapturesDelete(ctx, organizationId).BulkOrganizationDevicesPacketCaptureCapturesDelete(bulkOrganizationDevicesPacketCaptureCapturesDelete).Execute()

BulkDelete packet captures from cloud



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
    bulkOrganizationDevicesPacketCaptureCapturesDelete := *openapiclient.NewInlineObject260([]string{"CaptureIds_example"}) // InlineObject260 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.BulkOrganizationDevicesPacketCaptureCapturesDelete(context.Background(), organizationId).BulkOrganizationDevicesPacketCaptureCapturesDelete(bulkOrganizationDevicesPacketCaptureCapturesDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.BulkOrganizationDevicesPacketCaptureCapturesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkOrganizationDevicesPacketCaptureCapturesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkOrganizationDevicesPacketCaptureCapturesDelete** | [**InlineObject260**](InlineObject260.md) |  | 

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


## BulkUpdateOrganizationDevicesDetails

> InlineResponse200276 BulkUpdateOrganizationDevicesDetails(ctx, organizationId).BulkUpdateOrganizationDevicesDetails(bulkUpdateOrganizationDevicesDetails).Execute()

Updating device details (currently only used for Catalyst devices)



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
    bulkUpdateOrganizationDevicesDetails := *openapiclient.NewInlineObject257([]string{"Serials_example"}, []openapiclient.NetworksNetworkIdDevicesClaimDetails{*openapiclient.NewNetworksNetworkIdDevicesClaimDetails("Name_example")}) // InlineObject257 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.BulkUpdateOrganizationDevicesDetails(context.Background(), organizationId).BulkUpdateOrganizationDevicesDetails(bulkUpdateOrganizationDevicesDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.BulkUpdateOrganizationDevicesDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateOrganizationDevicesDetails`: InlineResponse200276
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.BulkUpdateOrganizationDevicesDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateOrganizationDevicesDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkUpdateOrganizationDevicesDetails** | [**InlineObject257**](InlineObject257.md) |  | 

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


## CheckinNetworkSmDevices

> InlineResponse200126 CheckinNetworkSmDevices(ctx, networkId).CheckinNetworkSmDevices(checkinNetworkSmDevices).Execute()

Force check-in a set of devices



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
    checkinNetworkSmDevices := *openapiclient.NewInlineObject125() // InlineObject125 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CheckinNetworkSmDevices(context.Background(), networkId).CheckinNetworkSmDevices(checkinNetworkSmDevices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CheckinNetworkSmDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckinNetworkSmDevices`: InlineResponse200126
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CheckinNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckinNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkinNetworkSmDevices** | [**InlineObject125**](InlineObject125.md) |  | 

### Return type

[**InlineResponse200126**](InlineResponse200126.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimNetworkDevices

> InlineResponse20095 ClaimNetworkDevices(ctx, networkId).ClaimNetworkDevices(claimNetworkDevices).AddAtomically(addAtomically).Execute()

Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requests against that device to succeed)



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
    claimNetworkDevices := *openapiclient.NewInlineObject95([]string{"Serials_example"}) // InlineObject95 | 
    addAtomically := true // bool | Whether to claim devices atomically. If true, all devices will be claimed or none will be claimed. Default is true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ClaimNetworkDevices(context.Background(), networkId).ClaimNetworkDevices(claimNetworkDevices).AddAtomically(addAtomically).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ClaimNetworkDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimNetworkDevices`: InlineResponse20095
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.ClaimNetworkDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimNetworkDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimNetworkDevices** | [**InlineObject95**](InlineObject95.md) |  | 
 **addAtomically** | **bool** | Whether to claim devices atomically. If true, all devices will be claimed or none will be claimed. Default is true. | 

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


## CloneOrganizationSwitchDevices

> InlineResponse200342 CloneOrganizationSwitchDevices(ctx, organizationId).CloneOrganizationSwitchDevices(cloneOrganizationSwitchDevices).Execute()

Clone port-level and some switch-level configuration settings from a source switch to one or more target switches



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
    cloneOrganizationSwitchDevices := *openapiclient.NewInlineObject302("SourceSerial_example", []string{"TargetSerials_example"}) // InlineObject302 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CloneOrganizationSwitchDevices(context.Background(), organizationId).CloneOrganizationSwitchDevices(cloneOrganizationSwitchDevices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CloneOrganizationSwitchDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneOrganizationSwitchDevices`: InlineResponse200342
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CloneOrganizationSwitchDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneOrganizationSwitchDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneOrganizationSwitchDevices** | [**InlineObject302**](InlineObject302.md) |  | 

### Return type

[**InlineResponse200342**](InlineResponse200342.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsArpTable

> InlineResponse2011 CreateDeviceLiveToolsArpTable(ctx, serial).CreateDeviceLiveToolsArpTable(createDeviceLiveToolsArpTable).Execute()

Enqueue a job to perform a ARP table request for the device



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
    createDeviceLiveToolsArpTable := *openapiclient.NewInlineObject16() // InlineObject16 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDeviceLiveToolsArpTable(context.Background(), serial).CreateDeviceLiveToolsArpTable(createDeviceLiveToolsArpTable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDeviceLiveToolsArpTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsArpTable`: InlineResponse2011
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDeviceLiveToolsArpTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsArpTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsArpTable** | [**InlineObject16**](InlineObject16.md) |  | 

### Return type

[**InlineResponse2011**](InlineResponse2011.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsCableTest

> InlineResponse2012 CreateDeviceLiveToolsCableTest(ctx, serial).CreateDeviceLiveToolsCableTest(createDeviceLiveToolsCableTest).Execute()

Enqueue a job to perform a cable test for the device on the specified ports



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
    createDeviceLiveToolsCableTest := *openapiclient.NewInlineObject17([]string{"Ports_example"}) // InlineObject17 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDeviceLiveToolsCableTest(context.Background(), serial).CreateDeviceLiveToolsCableTest(createDeviceLiveToolsCableTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDeviceLiveToolsCableTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsCableTest`: InlineResponse2012
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDeviceLiveToolsCableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsCableTest** | [**InlineObject17**](InlineObject17.md) |  | 

### Return type

[**InlineResponse2012**](InlineResponse2012.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsLedsBlink

> InlineResponse2013 CreateDeviceLiveToolsLedsBlink(ctx, serial).CreateDeviceLiveToolsLedsBlink(createDeviceLiveToolsLedsBlink).Execute()

Enqueue a job to blink LEDs on a device



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
    createDeviceLiveToolsLedsBlink := *openapiclient.NewInlineObject18(int32(123)) // InlineObject18 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDeviceLiveToolsLedsBlink(context.Background(), serial).CreateDeviceLiveToolsLedsBlink(createDeviceLiveToolsLedsBlink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDeviceLiveToolsLedsBlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsLedsBlink`: InlineResponse2013
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDeviceLiveToolsLedsBlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsLedsBlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsLedsBlink** | [**InlineObject18**](InlineObject18.md) |  | 

### Return type

[**InlineResponse2013**](InlineResponse2013.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsMacTable

> InlineResponse2014 CreateDeviceLiveToolsMacTable(ctx, serial).CreateDeviceLiveToolsMacTable(createDeviceLiveToolsMacTable).Execute()

Enqueue a job to request the MAC table from the device



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
    createDeviceLiveToolsMacTable := *openapiclient.NewInlineObject19() // InlineObject19 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDeviceLiveToolsMacTable(context.Background(), serial).CreateDeviceLiveToolsMacTable(createDeviceLiveToolsMacTable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDeviceLiveToolsMacTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsMacTable`: InlineResponse2014
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDeviceLiveToolsMacTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsMacTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsMacTable** | [**InlineObject19**](InlineObject19.md) |  | 

### Return type

[**InlineResponse2014**](InlineResponse2014.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsMulticastRouting

> InlineResponse2015 CreateDeviceLiveToolsMulticastRouting(ctx, serial).CreateDeviceLiveToolsMulticastRouting(createDeviceLiveToolsMulticastRouting).Execute()

Enqueue a job to perform a Multicast routing request for the device



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
    createDeviceLiveToolsMulticastRouting := *openapiclient.NewInlineObject20() // InlineObject20 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDeviceLiveToolsMulticastRouting(context.Background(), serial).CreateDeviceLiveToolsMulticastRouting(createDeviceLiveToolsMulticastRouting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDeviceLiveToolsMulticastRouting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsMulticastRouting`: InlineResponse2015
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDeviceLiveToolsMulticastRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsMulticastRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsMulticastRouting** | [**InlineObject20**](InlineObject20.md) |  | 

### Return type

[**InlineResponse2015**](InlineResponse2015.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsPing

> InlineResponse2016 CreateDeviceLiveToolsPing(ctx, serial).CreateDeviceLiveToolsPing(createDeviceLiveToolsPing).Execute()

Enqueue a job to ping a target host from the device



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
    createDeviceLiveToolsPing := *openapiclient.NewInlineObject21("Target_example") // InlineObject21 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDeviceLiveToolsPing(context.Background(), serial).CreateDeviceLiveToolsPing(createDeviceLiveToolsPing).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDeviceLiveToolsPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsPing`: InlineResponse2016
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDeviceLiveToolsPing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPing** | [**InlineObject21**](InlineObject21.md) |  | 

### Return type

[**InlineResponse2016**](InlineResponse2016.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsPingDevice

> InlineResponse2017 CreateDeviceLiveToolsPingDevice(ctx, serial).CreateDeviceLiveToolsPingDevice(createDeviceLiveToolsPingDevice).Execute()

Enqueue a job to check connectivity status to the device



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
    createDeviceLiveToolsPingDevice := *openapiclient.NewInlineObject22() // InlineObject22 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDeviceLiveToolsPingDevice(context.Background(), serial).CreateDeviceLiveToolsPingDevice(createDeviceLiveToolsPingDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDeviceLiveToolsPingDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsPingDevice`: InlineResponse2017
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDeviceLiveToolsPingDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPingDevice** | [**InlineObject22**](InlineObject22.md) |  | 

### Return type

[**InlineResponse2017**](InlineResponse2017.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsThroughputTest

> InlineResponse2018 CreateDeviceLiveToolsThroughputTest(ctx, serial).CreateDeviceLiveToolsThroughputTest(createDeviceLiveToolsThroughputTest).Execute()

Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput



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
    createDeviceLiveToolsThroughputTest := *openapiclient.NewInlineObject23() // InlineObject23 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDeviceLiveToolsThroughputTest(context.Background(), serial).CreateDeviceLiveToolsThroughputTest(createDeviceLiveToolsThroughputTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDeviceLiveToolsThroughputTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsThroughputTest`: InlineResponse2018
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDeviceLiveToolsThroughputTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsThroughputTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsThroughputTest** | [**InlineObject23**](InlineObject23.md) |  | 

### Return type

[**InlineResponse2018**](InlineResponse2018.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsWakeOnLan

> InlineResponse2019 CreateDeviceLiveToolsWakeOnLan(ctx, serial).CreateDeviceLiveToolsWakeOnLan(createDeviceLiveToolsWakeOnLan).Execute()

Enqueue a job to send a Wake-on-LAN packet from the device



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
    createDeviceLiveToolsWakeOnLan := *openapiclient.NewInlineObject24(int32(123), "Mac_example") // InlineObject24 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDeviceLiveToolsWakeOnLan(context.Background(), serial).CreateDeviceLiveToolsWakeOnLan(createDeviceLiveToolsWakeOnLan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDeviceLiveToolsWakeOnLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsWakeOnLan`: InlineResponse2019
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDeviceLiveToolsWakeOnLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsWakeOnLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsWakeOnLan** | [**InlineObject24**](InlineObject24.md) |  | 

### Return type

[**InlineResponse2019**](InlineResponse2019.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationDevicesControllerMigration

> []InlineResponse200275Items CreateOrganizationDevicesControllerMigration(ctx, organizationId).CreateOrganizationDevicesControllerMigration(createOrganizationDevicesControllerMigration).Execute()

Migrate devices to another controller or management mode



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
    createOrganizationDevicesControllerMigration := *openapiclient.NewInlineObject256([]string{"Serials_example"}, "Target_example") // InlineObject256 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateOrganizationDevicesControllerMigration(context.Background(), organizationId).CreateOrganizationDevicesControllerMigration(createOrganizationDevicesControllerMigration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateOrganizationDevicesControllerMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationDevicesControllerMigration`: []InlineResponse200275Items
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateOrganizationDevicesControllerMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationDevicesControllerMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationDevicesControllerMigration** | [**InlineObject256**](InlineObject256.md) |  | 

### Return type

[**[]InlineResponse200275Items**](InlineResponse200275Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationDevicesPacketCaptureCapture

> InlineResponse200278Items CreateOrganizationDevicesPacketCaptureCapture(ctx, organizationId).CreateOrganizationDevicesPacketCaptureCapture(createOrganizationDevicesPacketCaptureCapture).Execute()

Perform a packet capture on a device and store in Meraki Cloud



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
    createOrganizationDevicesPacketCaptureCapture := *openapiclient.NewInlineObject258([]string{"Serials_example"}, "Name_example") // InlineObject258 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateOrganizationDevicesPacketCaptureCapture(context.Background(), organizationId).CreateOrganizationDevicesPacketCaptureCapture(createOrganizationDevicesPacketCaptureCapture).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateOrganizationDevicesPacketCaptureCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationDevicesPacketCaptureCapture`: InlineResponse200278Items
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateOrganizationDevicesPacketCaptureCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationDevicesPacketCaptureCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationDevicesPacketCaptureCapture** | [**InlineObject258**](InlineObject258.md) |  | 

### Return type

[**InlineResponse200278Items**](InlineResponse200278Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationDevicesPacketCaptureSchedule

> InlineResponse200280Items CreateOrganizationDevicesPacketCaptureSchedule(ctx, organizationId).CreateOrganizationDevicesPacketCaptureSchedule(createOrganizationDevicesPacketCaptureSchedule).Execute()

Create a schedule for packet capture



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
    createOrganizationDevicesPacketCaptureSchedule := *openapiclient.NewInlineObject262([]openapiclient.OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices{*openapiclient.NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices()}) // InlineObject262 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateOrganizationDevicesPacketCaptureSchedule(context.Background(), organizationId).CreateOrganizationDevicesPacketCaptureSchedule(createOrganizationDevicesPacketCaptureSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateOrganizationDevicesPacketCaptureSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationDevicesPacketCaptureSchedule`: InlineResponse200280Items
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateOrganizationDevicesPacketCaptureSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationDevicesPacketCaptureScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationDevicesPacketCaptureSchedule** | [**InlineObject262**](InlineObject262.md) |  | 

### Return type

[**InlineResponse200280Items**](InlineResponse200280Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryDevicesSwapsBulk

> InlineResponse207 CreateOrganizationInventoryDevicesSwapsBulk(ctx, organizationId).CreateOrganizationInventoryDevicesSwapsBulk(createOrganizationInventoryDevicesSwapsBulk).Execute()

Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.



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
    createOrganizationInventoryDevicesSwapsBulk := *openapiclient.NewInlineObject273([]openapiclient.OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps{*openapiclient.NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps(*openapiclient.NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices("Old_example", "New_example"), "AfterAction_example")}) // InlineObject273 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateOrganizationInventoryDevicesSwapsBulk(context.Background(), organizationId).CreateOrganizationInventoryDevicesSwapsBulk(createOrganizationInventoryDevicesSwapsBulk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateOrganizationInventoryDevicesSwapsBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryDevicesSwapsBulk`: InlineResponse207
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateOrganizationInventoryDevicesSwapsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryDevicesSwapsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryDevicesSwapsBulk** | [**InlineObject273**](InlineObject273.md) |  | 

### Return type

[**InlineResponse207**](InlineResponse207.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationWirelessDevicesRadsecCertificatesAuthority

> InlineResponse200366 CreateOrganizationWirelessDevicesRadsecCertificatesAuthority(ctx, organizationId).Execute()

Create an organization's RADSEC device Certificate Authority (CA)



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
    resp, r, err := apiClient.DevicesApi.CreateOrganizationWirelessDevicesRadsecCertificatesAuthority(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateOrganizationWirelessDevicesRadsecCertificatesAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationWirelessDevicesRadsecCertificatesAuthority`: InlineResponse200366
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateOrganizationWirelessDevicesRadsecCertificatesAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationWirelessDevicesRadsecCertificatesAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200366**](InlineResponse200366.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationDevicesPacketCaptureCapture

> DeleteOrganizationDevicesPacketCaptureCapture(ctx, organizationId, captureId).Execute()

Delete a single packet capture from cloud using captureId



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
    captureId := "captureId_example" // string | Capture ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DeleteOrganizationDevicesPacketCaptureCapture(context.Background(), organizationId, captureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeleteOrganizationDevicesPacketCaptureCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**captureId** | **string** | Capture ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationDevicesPacketCaptureCaptureRequest struct via the builder pattern


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


## DeleteOrganizationDevicesPacketCaptureSchedule

> DeleteOrganizationDevicesPacketCaptureSchedule(ctx, organizationId, scheduleId).DeleteOrganizationDevicesPacketCaptureSchedule(deleteOrganizationDevicesPacketCaptureSchedule).Execute()

Delete schedule from cloud



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
    scheduleId := "scheduleId_example" // string | Schedule ID
    deleteOrganizationDevicesPacketCaptureSchedule := *openapiclient.NewInlineObject265("ScheduleId_example") // InlineObject265 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DeleteOrganizationDevicesPacketCaptureSchedule(context.Background(), organizationId, scheduleId).DeleteOrganizationDevicesPacketCaptureSchedule(deleteOrganizationDevicesPacketCaptureSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeleteOrganizationDevicesPacketCaptureSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**scheduleId** | **string** | Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationDevicesPacketCaptureScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteOrganizationDevicesPacketCaptureSchedule** | [**InlineObject265**](InlineObject265.md) |  | 

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


## GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl

> InlineResponse200279 GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl(ctx, organizationId, captureId).Execute()

Get presigned download URL for given packet capture id



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
    captureId := "captureId_example" // string | Capture ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl(context.Background(), organizationId, captureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl`: InlineResponse200279
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**captureId** | **string** | Capture ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateOrganizationDevicesPacketCaptureCaptureDownloadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200279**](InlineResponse200279.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevice

> InlineResponse2006 GetDevice(ctx, serial).Execute()

Return a single device



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
    resp, r, err := apiClient.DevicesApi.GetDevice(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevice`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCellularSims

> InlineResponse20018 GetDeviceCellularSims(ctx, serial).Execute()

Return the SIM and APN configurations for a cellular device.



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
    resp, r, err := apiClient.DevicesApi.GetDeviceCellularSims(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceCellularSims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCellularSims`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceCellularSims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceClients

> []InlineResponse20021 GetDeviceClients(ctx, serial).T0(t0).Timespan(timespan).Execute()

List the clients of a device, up to a maximum of a month ago



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceClients(context.Background(), serial).T0(t0).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceClients`: []InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]InlineResponse20021**](InlineResponse20021.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsArpTable

> InlineResponse20022 GetDeviceLiveToolsArpTable(ctx, serial, arpTableId).Execute()

Return an ARP table live tool job.



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
    arpTableId := "arpTableId_example" // string | Arp table ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceLiveToolsArpTable(context.Background(), serial, arpTableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceLiveToolsArpTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsArpTable`: InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceLiveToolsArpTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**arpTableId** | **string** | Arp table ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsArpTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20022**](InlineResponse20022.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsCableTest

> InlineResponse20023 GetDeviceLiveToolsCableTest(ctx, serial, id).Execute()

Return a cable test live tool job.



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceLiveToolsCableTest(context.Background(), serial, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceLiveToolsCableTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsCableTest`: InlineResponse20023
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceLiveToolsCableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20023**](InlineResponse20023.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsLedsBlink

> InlineResponse20024 GetDeviceLiveToolsLedsBlink(ctx, serial, ledsBlinkId).Execute()

Return a blink LEDs job



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
    ledsBlinkId := "ledsBlinkId_example" // string | Leds blink ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceLiveToolsLedsBlink(context.Background(), serial, ledsBlinkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceLiveToolsLedsBlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsLedsBlink`: InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceLiveToolsLedsBlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**ledsBlinkId** | **string** | Leds blink ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsLedsBlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20024**](InlineResponse20024.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsMacTable

> InlineResponse20025 GetDeviceLiveToolsMacTable(ctx, serial, macTableId).Execute()

Return a MAC table live tool job.



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
    macTableId := "macTableId_example" // string | Mac table ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceLiveToolsMacTable(context.Background(), serial, macTableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceLiveToolsMacTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsMacTable`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceLiveToolsMacTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**macTableId** | **string** | Mac table ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsMacTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20025**](InlineResponse20025.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsMulticastRouting

> InlineResponse20026 GetDeviceLiveToolsMulticastRouting(ctx, serial, multicastRoutingId).Execute()

Return a Multicast routing live tool job.



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
    multicastRoutingId := "multicastRoutingId_example" // string | Multicast routing ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceLiveToolsMulticastRouting(context.Background(), serial, multicastRoutingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceLiveToolsMulticastRouting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsMulticastRouting`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceLiveToolsMulticastRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**multicastRoutingId** | **string** | Multicast routing ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsMulticastRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPing

> InlineResponse20027 GetDeviceLiveToolsPing(ctx, serial, id).Execute()

Return a ping job



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceLiveToolsPing(context.Background(), serial, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceLiveToolsPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsPing`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceLiveToolsPing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPingDevice

> InlineResponse20028 GetDeviceLiveToolsPingDevice(ctx, serial, id).Execute()

Return a ping device job



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceLiveToolsPingDevice(context.Background(), serial, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceLiveToolsPingDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsPingDevice`: InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceLiveToolsPingDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsThroughputTest

> InlineResponse20029 GetDeviceLiveToolsThroughputTest(ctx, serial, throughputTestId).Execute()

Return a throughput test job



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
    throughputTestId := "throughputTestId_example" // string | Throughput test ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceLiveToolsThroughputTest(context.Background(), serial, throughputTestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceLiveToolsThroughputTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsThroughputTest`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceLiveToolsThroughputTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**throughputTestId** | **string** | Throughput test ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsThroughputTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20029**](InlineResponse20029.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsWakeOnLan

> InlineResponse20030 GetDeviceLiveToolsWakeOnLan(ctx, serial, wakeOnLanId).Execute()

Return a Wake-on-LAN job



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
    wakeOnLanId := "wakeOnLanId_example" // string | Wake on lan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceLiveToolsWakeOnLan(context.Background(), serial, wakeOnLanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceLiveToolsWakeOnLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsWakeOnLan`: InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceLiveToolsWakeOnLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**wakeOnLanId** | **string** | Wake on lan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsWakeOnLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20030**](InlineResponse20030.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLldpCdp

> InlineResponse20031 GetDeviceLldpCdp(ctx, serial).Execute()

List LLDP and CDP information for a device



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
    resp, r, err := apiClient.DevicesApi.GetDeviceLldpCdp(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceLldpCdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLldpCdp`: InlineResponse20031
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceLldpCdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLldpCdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20031**](InlineResponse20031.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLossAndLatencyHistory

> []InlineResponse20032 GetDeviceLossAndLatencyHistory(ctx, serial).Ip(ip).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Uplink(uplink).Execute()

Get the uplink loss percentage and latency in milliseconds, and goodput in kilobits per second for MX, MG and Z devices.



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
    ip := "ip_example" // string | The destination IP used to obtain the requested stats. This is required.
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 60 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 60, 600, 3600, 86400. The default is 60. (optional)
    uplink := "uplink_example" // string | The WAN uplink used to obtain the requested stats. Valid uplinks are wan1, wan2, wan3, cellular. The default is wan1. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceLossAndLatencyHistory(context.Background(), serial).Ip(ip).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Uplink(uplink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceLossAndLatencyHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLossAndLatencyHistory`: []InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceLossAndLatencyHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLossAndLatencyHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ip** | **string** | The destination IP used to obtain the requested stats. This is required. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 60 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 60, 600, 3600, 86400. The default is 60. | 
 **uplink** | **string** | The WAN uplink used to obtain the requested stats. Valid uplinks are wan1, wan2, wan3, cellular. The default is wan1. | 

### Return type

[**[]InlineResponse20032**](InlineResponse20032.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceManagementInterface

> InlineResponse20033 GetDeviceManagementInterface(ctx, serial).Execute()

Return the management interface settings for a device



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
    resp, r, err := apiClient.DevicesApi.GetDeviceManagementInterface(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceManagementInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceManagementInterface`: InlineResponse20033
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20033**](InlineResponse20033.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevices

> []InlineResponse2006 GetNetworkDevices(ctx, networkId).Execute()

List the devices in a network



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
    resp, r, err := apiClient.DevicesApi.GetNetworkDevices(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDevices`: []InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse2006**](InlineResponse2006.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceCellularUsageHistory

> []InlineResponse200132 GetNetworkSmDeviceCellularUsageHistory(ctx, networkId, deviceId).Execute()

Return the client's daily cellular data usage history



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDeviceCellularUsageHistory(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDeviceCellularUsageHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceCellularUsageHistory`: []InlineResponse200132
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDeviceCellularUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceCellularUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse200132**](InlineResponse200132.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceCerts

> []InlineResponse200133 GetNetworkSmDeviceCerts(ctx, networkId, deviceId).Execute()

List the certs on a device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDeviceCerts(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDeviceCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceCerts`: []InlineResponse200133
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDeviceCerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse200133**](InlineResponse200133.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceConnectivity

> []InlineResponse200134 GetNetworkSmDeviceConnectivity(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns historical connectivity data (whether a device is regularly checking in to Dashboard).



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
    deviceId := "deviceId_example" // string | Device ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDeviceConnectivity(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDeviceConnectivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceConnectivity`: []InlineResponse200134
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDeviceConnectivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceConnectivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse200134**](InlineResponse200134.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceDesktopLogs

> []InlineResponse200135 GetNetworkSmDeviceDesktopLogs(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return historical records of various Systems Manager network connection details for desktop devices.



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
    deviceId := "deviceId_example" // string | Device ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDeviceDesktopLogs(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDeviceDesktopLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceDesktopLogs`: []InlineResponse200135
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDeviceDesktopLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceDesktopLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse200135**](InlineResponse200135.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceDeviceCommandLogs

> []InlineResponse200136 GetNetworkSmDeviceDeviceCommandLogs(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return historical records of commands sent to Systems Manager devices



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
    deviceId := "deviceId_example" // string | Device ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDeviceDeviceCommandLogs(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDeviceDeviceCommandLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceDeviceCommandLogs`: []InlineResponse200136
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDeviceDeviceCommandLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceDeviceCommandLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse200136**](InlineResponse200136.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceDeviceProfiles

> []InlineResponse200137 GetNetworkSmDeviceDeviceProfiles(ctx, networkId, deviceId).Execute()

Get the installed profiles associated with a device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDeviceDeviceProfiles(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDeviceDeviceProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceDeviceProfiles`: []InlineResponse200137
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDeviceDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse200137**](InlineResponse200137.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceNetworkAdapters

> []InlineResponse200138 GetNetworkSmDeviceNetworkAdapters(ctx, networkId, deviceId).Execute()

List the network adapters of a device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDeviceNetworkAdapters(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDeviceNetworkAdapters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceNetworkAdapters`: []InlineResponse200138
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDeviceNetworkAdapters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceNetworkAdaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse200138**](InlineResponse200138.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDevicePerformanceHistory

> []InlineResponse200139 GetNetworkSmDevicePerformanceHistory(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return historical records of various Systems Manager client metrics for desktop devices.



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
    deviceId := "deviceId_example" // string | Device ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDevicePerformanceHistory(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDevicePerformanceHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDevicePerformanceHistory`: []InlineResponse200139
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDevicePerformanceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDevicePerformanceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse200139**](InlineResponse200139.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceRestrictions

> InlineResponse200140 GetNetworkSmDeviceRestrictions(ctx, networkId, deviceId).Execute()

List the restrictions on a device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDeviceRestrictions(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDeviceRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceRestrictions`: InlineResponse200140
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDeviceRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200140**](InlineResponse200140.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceSecurityCenters

> []InlineResponse200141 GetNetworkSmDeviceSecurityCenters(ctx, networkId, deviceId).Execute()

List the security centers on a device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDeviceSecurityCenters(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDeviceSecurityCenters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceSecurityCenters`: []InlineResponse200141
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDeviceSecurityCenters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceSecurityCentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse200141**](InlineResponse200141.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceSoftwares

> []InlineResponse200142 GetNetworkSmDeviceSoftwares(ctx, networkId, deviceId).Execute()

Get a list of softwares associated with a device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDeviceSoftwares(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDeviceSoftwares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceSoftwares`: []InlineResponse200142
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDeviceSoftwares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceSoftwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse200142**](InlineResponse200142.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceWlanLists

> []InlineResponse200144 GetNetworkSmDeviceWlanLists(ctx, networkId, deviceId).Execute()

List the saved SSID names on a device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDeviceWlanLists(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDeviceWlanLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceWlanLists`: []InlineResponse200144
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDeviceWlanLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceWlanListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse200144**](InlineResponse200144.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDevices

> []InlineResponse200125 GetNetworkSmDevices(ctx, networkId).Fields(fields).WifiMacs(wifiMacs).Serials(serials).Ids(ids).Uuids(uuids).SystemTypes(systemTypes).Scope(scope).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the devices enrolled in an SM network with various specified fields and filters



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
    fields := []string{"Inner_example"} // []string | Additional fields that will be displayed for each device.     The default fields are: id, name, tags, ssid, wifiMac, osName, systemModel, uuid, and serialNumber. The additional fields are: ip,     systemType, availableDeviceCapacity, kioskAppName, biosVersion, lastConnected, missingAppsCount, userSuppliedAddress, location, lastUser,     ownerEmail, ownerUsername, osBuild, publicIp, phoneNumber, diskInfoJson, deviceCapacity, isManaged, hadMdm, isSupervised, meid, imei, iccid,     simCarrierNetwork, cellularDataUsed, isHotspotEnabled, createdAt, batteryEstCharge, quarantined, avName, avRunning, asName, fwName,     isRooted, loginRequired, screenLockEnabled, screenLockDelay, autoLoginDisabled, autoTags, hasMdm, hasDesktopAgent, diskEncryptionEnabled,     hardwareEncryptionCaps, passCodeLock, usesHardwareKeystore, androidSecurityPatchVersion, cellular, and url. (optional)
    wifiMacs := []string{"Inner_example"} // []string | Filter devices by wifi mac(s). (optional)
    serials := []string{"Inner_example"} // []string | Filter devices by serial(s). (optional)
    ids := []string{"Inner_example"} // []string | Filter devices by id(s). (optional)
    uuids := []string{"Inner_example"} // []string | Filter devices by uuid(s). (optional)
    systemTypes := []string{"Inner_example"} // []string | Filter devices by system type(s). (optional)
    scope := []string{"Inner_example"} // []string | Specify a scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkSmDevices(context.Background(), networkId).Fields(fields).WifiMacs(wifiMacs).Serials(serials).Ids(ids).Uuids(uuids).SystemTypes(systemTypes).Scope(scope).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkSmDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDevices`: []InlineResponse200125
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Additional fields that will be displayed for each device.     The default fields are: id, name, tags, ssid, wifiMac, osName, systemModel, uuid, and serialNumber. The additional fields are: ip,     systemType, availableDeviceCapacity, kioskAppName, biosVersion, lastConnected, missingAppsCount, userSuppliedAddress, location, lastUser,     ownerEmail, ownerUsername, osBuild, publicIp, phoneNumber, diskInfoJson, deviceCapacity, isManaged, hadMdm, isSupervised, meid, imei, iccid,     simCarrierNetwork, cellularDataUsed, isHotspotEnabled, createdAt, batteryEstCharge, quarantined, avName, avRunning, asName, fwName,     isRooted, loginRequired, screenLockEnabled, screenLockDelay, autoLoginDisabled, autoTags, hasMdm, hasDesktopAgent, diskEncryptionEnabled,     hardwareEncryptionCaps, passCodeLock, usesHardwareKeystore, androidSecurityPatchVersion, cellular, and url. | 
 **wifiMacs** | **[]string** | Filter devices by wifi mac(s). | 
 **serials** | **[]string** | Filter devices by serial(s). | 
 **ids** | **[]string** | Filter devices by id(s). | 
 **uuids** | **[]string** | Filter devices by uuid(s). | 
 **systemTypes** | **[]string** | Filter devices by system type(s). | 
 **scope** | **[]string** | Specify a scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse200125**](InlineResponse200125.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessDevicesConnectionStats

> []InlineResponse20046 GetNetworkWirelessDevicesConnectionStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()

Aggregated connectivity info for this network, grouped by node



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkWirelessDevicesConnectionStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkWirelessDevicesConnectionStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessDevicesConnectionStats`: []InlineResponse20046
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkWirelessDevicesConnectionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessDevicesConnectionStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 

### Return type

[**[]InlineResponse20046**](InlineResponse20046.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessDevicesLatencyStats

> []map[string]interface{} GetNetworkWirelessDevicesLatencyStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()

Aggregated latency info for this network, grouped by node



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)
    fields := "fields_example" // string | Partial selection: If present, this call will return only the selected fields of [\"rawDistribution\", \"avg\"]. All fields will be returned by default. Selected fields must be entered as a comma separated string. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetNetworkWirelessDevicesLatencyStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetNetworkWirelessDevicesLatencyStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessDevicesLatencyStats`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetNetworkWirelessDevicesLatencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessDevicesLatencyStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 
 **fields** | **string** | Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string. | 

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


## GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice

> InlineResponse200257 GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Uplink overrides configured locally on Campus Gateway devices in an organization.



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
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned devices will be filtered to only include these serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice`: InlineResponse200257
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned devices will be filtered to only include these serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200257**](InlineResponse200257.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevices

> []InlineResponse20096 GetOrganizationDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).NetworkIds(networkIds).ProductTypes(productTypes).Tags(tags).TagsFilterType(tagsFilterType).Name(name).Mac(mac).Serial(serial).Model(model).Macs(macs).Serials(serials).SensorMetrics(sensorMetrics).SensorAlertProfileIds(sensorAlertProfileIds).Models(models).Execute()

List the devices in an organization that have been assigned to a network.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 5000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    configurationUpdatedAfter := "configurationUpdatedAfter_example" // string | Filter results by whether or not the device's configuration has been updated after the given timestamp (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by network. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. (optional)
    tags := []string{"Inner_example"} // []string | Optional parameter to filter devices by tags. (optional)
    tagsFilterType := "tagsFilterType_example" // string | Optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)
    name := "name_example" // string | Optional parameter to filter devices by name. All returned devices will have a name that contains the search term or is an exact match. (optional)
    mac := "mac_example" // string | Optional parameter to filter devices by MAC address. All returned devices will have a MAC address that contains the search term or is an exact match. (optional)
    serial := "serial_example" // string | Optional parameter to filter devices by serial number. All returned devices will have a serial number that contains the search term or is an exact match. (optional)
    model := "model_example" // string | Optional parameter to filter devices by model. All returned devices will have a model that contains the search term or is an exact match. (optional)
    macs := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more MAC addresses. All returned devices will have a MAC address that is an exact match. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more serial numbers. All returned devices will have a serial number that is an exact match. (optional)
    sensorMetrics := []string{"SensorMetrics_example"} // []string | Optional parameter to filter devices by the metrics that they provide. Only applies to sensor devices. (optional)
    sensorAlertProfileIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by the alert profiles that are bound to them. Only applies to sensor devices. (optional)
    models := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).NetworkIds(networkIds).ProductTypes(productTypes).Tags(tags).TagsFilterType(tagsFilterType).Name(name).Mac(mac).Serial(serial).Model(model).Macs(macs).Serials(serials).SensorMetrics(sensorMetrics).SensorAlertProfileIds(sensorAlertProfileIds).Models(models).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevices`: []InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 5000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **configurationUpdatedAfter** | **string** | Filter results by whether or not the device&#39;s configuration has been updated after the given timestamp | 
 **networkIds** | **[]string** | Optional parameter to filter devices by network. | 
 **productTypes** | **[]string** | Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. | 
 **tags** | **[]string** | Optional parameter to filter devices by tags. | 
 **tagsFilterType** | **string** | Optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 
 **name** | **string** | Optional parameter to filter devices by name. All returned devices will have a name that contains the search term or is an exact match. | 
 **mac** | **string** | Optional parameter to filter devices by MAC address. All returned devices will have a MAC address that contains the search term or is an exact match. | 
 **serial** | **string** | Optional parameter to filter devices by serial number. All returned devices will have a serial number that contains the search term or is an exact match. | 
 **model** | **string** | Optional parameter to filter devices by model. All returned devices will have a model that contains the search term or is an exact match. | 
 **macs** | **[]string** | Optional parameter to filter devices by one or more MAC addresses. All returned devices will have a MAC address that is an exact match. | 
 **serials** | **[]string** | Optional parameter to filter devices by one or more serial numbers. All returned devices will have a serial number that is an exact match. | 
 **sensorMetrics** | **[]string** | Optional parameter to filter devices by the metrics that they provide. Only applies to sensor devices. | 
 **sensorAlertProfileIds** | **[]string** | Optional parameter to filter devices by the alert profiles that are bound to them. Only applies to sensor devices. | 
 **models** | **[]string** | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. | 

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


## GetOrganizationDevicesAvailabilities

> []InlineResponse200273 GetOrganizationDevicesAvailabilities(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Statuses(statuses).Execute()

List the availability information for devices in an organization



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. (optional)
    productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. Valid types are wireless, appliance, switch, camera, cellularGateway, sensor, wirelessController, and campusGateway (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)
    statuses := []string{"Statuses_example"} // []string | Optional parameter to filter device availabilities by device status. This filter uses multiple exact matches. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesAvailabilities(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesAvailabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesAvailabilities`: []InlineResponse200273
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesAvailabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. Valid types are wireless, appliance, switch, camera, cellularGateway, sensor, wirelessController, and campusGateway | 
 **serials** | **[]string** | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 
 **statuses** | **[]string** | Optional parameter to filter device availabilities by device status. This filter uses multiple exact matches. | 

### Return type

[**[]InlineResponse200273**](InlineResponse200273.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesAvailabilitiesChangeHistory

> []InlineResponse200274 GetOrganizationDevicesAvailabilitiesChangeHistory(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Serials(serials).ProductTypes(productTypes).NetworkIds(networkIds).Statuses(statuses).Execute()

List the availability history information for devices in an organization.



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities history by device serial numbers (optional)
    productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities history by device product types (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities history by network IDs (optional)
    statuses := []string{"Statuses_example"} // []string | Optional parameter to filter device availabilities history by device statuses (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesAvailabilitiesChangeHistory(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Serials(serials).ProductTypes(productTypes).NetworkIds(networkIds).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesAvailabilitiesChangeHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesAvailabilitiesChangeHistory`: []InlineResponse200274
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesAvailabilitiesChangeHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **serials** | **[]string** | Optional parameter to filter device availabilities history by device serial numbers | 
 **productTypes** | **[]string** | Optional parameter to filter device availabilities history by device product types | 
 **networkIds** | **[]string** | Optional parameter to filter device availabilities history by network IDs | 
 **statuses** | **[]string** | Optional parameter to filter device availabilities history by device statuses | 

### Return type

[**[]InlineResponse200274**](InlineResponse200274.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesControllerMigrations

> InlineResponse200275 GetOrganizationDevicesControllerMigrations(ctx, organizationId).Serials(serials).NetworkIds(networkIds).Target(target).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Retrieve device migration statuses in an organization



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
    serials := []string{"Inner_example"} // []string | A list of Meraki Serials for which to retrieve migrations (optional)
    networkIds := []string{"Inner_example"} // []string | Filter device migrations by network IDs (optional)
    target := "target_example" // string | Filter device migrations by target destination (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesControllerMigrations(context.Background(), organizationId).Serials(serials).NetworkIds(networkIds).Target(target).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesControllerMigrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesControllerMigrations`: InlineResponse200275
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesControllerMigrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesControllerMigrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of Meraki Serials for which to retrieve migrations | 
 **networkIds** | **[]string** | Filter device migrations by network IDs | 
 **target** | **string** | Filter device migrations by target destination | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

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


## GetOrganizationDevicesOverviewByModel

> InlineResponse200277 GetOrganizationDevicesOverviewByModel(ctx, organizationId).Models(models).NetworkIds(networkIds).ProductTypes(productTypes).Execute()

Lists the count for each device model



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
    models := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by networkId. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter device by device product types. This filter uses multiple exact matches. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesOverviewByModel(context.Background(), organizationId).Models(models).NetworkIds(networkIds).ProductTypes(productTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesOverviewByModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesOverviewByModel`: InlineResponse200277
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesOverviewByModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesOverviewByModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **models** | **[]string** | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. | 
 **networkIds** | **[]string** | Optional parameter to filter devices by networkId. | 
 **productTypes** | **[]string** | Optional parameter to filter device by device product types. This filter uses multiple exact matches. | 

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


## GetOrganizationDevicesPacketCaptureCaptures

> InlineResponse200278 GetOrganizationDevicesPacketCaptureCaptures(ctx, organizationId).CaptureIds(captureIds).NetworkIds(networkIds).Serials(serials).Process(process).CaptureStatus(captureStatus).Name(name).ClientMac(clientMac).Notes(notes).DeviceName(deviceName).AdminName(adminName).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List Packet Captures



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
    captureIds := []string{"Inner_example"} // []string | Return the packet captures of the specified capture ids (optional)
    networkIds := []string{"Inner_example"} // []string | Return the packet captures of the specified network(s) (optional)
    serials := []string{"Inner_example"} // []string | Return the packet captures of the specified device(s) (optional)
    process := []string{"Inner_example"} // []string | Return the packet captures of the specified process (optional)
    captureStatus := []string{"Inner_example"} // []string | Return the packet captures of the specified capture status (optional)
    name := []string{"Inner_example"} // []string | Return the packet captures matching the specified name (optional)
    clientMac := []string{"Inner_example"} // []string | Return the packet captures matching the specified client macs (optional)
    notes := "notes_example" // string | Return the packet captures matching the specified notes (optional)
    deviceName := "deviceName_example" // string | Return the packet captures matching the specified device name (optional)
    adminName := "adminName_example" // string | Return the packet captures matching the admin name (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 365 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'descending'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesPacketCaptureCaptures(context.Background(), organizationId).CaptureIds(captureIds).NetworkIds(networkIds).Serials(serials).Process(process).CaptureStatus(captureStatus).Name(name).ClientMac(clientMac).Notes(notes).DeviceName(deviceName).AdminName(adminName).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesPacketCaptureCaptures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesPacketCaptureCaptures`: InlineResponse200278
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesPacketCaptureCaptures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesPacketCaptureCapturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **captureIds** | **[]string** | Return the packet captures of the specified capture ids | 
 **networkIds** | **[]string** | Return the packet captures of the specified network(s) | 
 **serials** | **[]string** | Return the packet captures of the specified device(s) | 
 **process** | **[]string** | Return the packet captures of the specified process | 
 **captureStatus** | **[]string** | Return the packet captures of the specified capture status | 
 **name** | **[]string** | Return the packet captures matching the specified name | 
 **clientMac** | **[]string** | Return the packet captures matching the specified client macs | 
 **notes** | **string** | Return the packet captures matching the specified notes | 
 **deviceName** | **string** | Return the packet captures matching the specified device name | 
 **adminName** | **string** | Return the packet captures matching the admin name | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 365 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;descending&#39;. | 

### Return type

[**InlineResponse200278**](InlineResponse200278.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesPacketCaptureSchedules

> InlineResponse200280 GetOrganizationDevicesPacketCaptureSchedules(ctx, organizationId).ScheduleIds(scheduleIds).NetworkIds(networkIds).DeviceIds(deviceIds).Execute()

List the Packet Capture Schedules



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
    scheduleIds := []string{"Inner_example"} // []string | Return the packet captures schedules of the specified packet capture schedule ids (optional)
    networkIds := []string{"Inner_example"} // []string | Return the scheduled packet captures of the specified network(s) (optional)
    deviceIds := []string{"Inner_example"} // []string | Return the scheduled packet captures of the specified device(s) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesPacketCaptureSchedules(context.Background(), organizationId).ScheduleIds(scheduleIds).NetworkIds(networkIds).DeviceIds(deviceIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesPacketCaptureSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesPacketCaptureSchedules`: InlineResponse200280
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesPacketCaptureSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesPacketCaptureSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduleIds** | **[]string** | Return the packet captures schedules of the specified packet capture schedule ids | 
 **networkIds** | **[]string** | Return the scheduled packet captures of the specified network(s) | 
 **deviceIds** | **[]string** | Return the scheduled packet captures of the specified device(s) | 

### Return type

[**InlineResponse200280**](InlineResponse200280.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesPowerModulesStatusesByDevice

> []InlineResponse200282 GetOrganizationDevicesPowerModulesStatusesByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the most recent status information for power modules in rackmount MX and MS devices that support them



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. (optional)
    productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesPowerModulesStatusesByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesPowerModulesStatusesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesPowerModulesStatusesByDevice`: []InlineResponse200282
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesPowerModulesStatusesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]InlineResponse200282**](InlineResponse200282.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesProvisioningStatuses

> []InlineResponse200283 GetOrganizationDevicesProvisioningStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Status(status).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the provisioning statuses information for devices in an organization.



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device by network ID. This filter uses multiple exact matches. (optional)
    productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device by device product types. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device by device serial numbers. This filter uses multiple exact matches. (optional)
    status := "status_example" // string | An optional parameter to filter devices by the provisioning status. Accepted statuses: unprovisioned, incomplete, complete. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesProvisioningStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Status(status).Tags(tags).TagsFilterType(tagsFilterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesProvisioningStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesProvisioningStatuses`: []InlineResponse200283
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesProvisioningStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesProvisioningStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device by device serial numbers. This filter uses multiple exact matches. | 
 **status** | **string** | An optional parameter to filter devices by the provisioning status. Accepted statuses: unprovisioned, incomplete, complete. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]InlineResponse200283**](InlineResponse200283.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesStatuses

> []InlineResponse200284 GetOrganizationDevicesStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Statuses(statuses).ProductTypes(productTypes).Models(models).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the status of every Meraki device in the organization



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by network ids. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter devices by serials. (optional)
    statuses := []string{"Statuses_example"} // []string | Optional parameter to filter devices by statuses. Valid statuses are [\"online\", \"alerting\", \"offline\", \"dormant\"]. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. (optional)
    models := []string{"Inner_example"} // []string | Optional parameter to filter devices by models. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Statuses(statuses).ProductTypes(productTypes).Models(models).Tags(tags).TagsFilterType(tagsFilterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesStatuses`: []InlineResponse200284
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter devices by network ids. | 
 **serials** | **[]string** | Optional parameter to filter devices by serials. | 
 **statuses** | **[]string** | Optional parameter to filter devices by statuses. Valid statuses are [\&quot;online\&quot;, \&quot;alerting\&quot;, \&quot;offline\&quot;, \&quot;dormant\&quot;]. | 
 **productTypes** | **[]string** | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. | 
 **models** | **[]string** | Optional parameter to filter devices by models. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]InlineResponse200284**](InlineResponse200284.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesStatusesOverview

> InlineResponse200285 GetOrganizationDevicesStatusesOverview(ctx, organizationId).ProductTypes(productTypes).NetworkIds(networkIds).Execute()

Return an overview of current device statuses



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
    productTypes := []string{"ProductTypes_example"} // []string | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. (optional)
    networkIds := []string{"Inner_example"} // []string | An optional parameter to filter device statuses by network. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesStatusesOverview(context.Background(), organizationId).ProductTypes(productTypes).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesStatusesOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesStatusesOverview`: InlineResponse200285
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesStatusesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesStatusesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productTypes** | **[]string** | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. | 
 **networkIds** | **[]string** | An optional parameter to filter device statuses by network. | 

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


## GetOrganizationDevicesSystemMemoryUsageHistoryByInterval

> InlineResponse200286 GetOrganizationDevicesSystemMemoryUsageHistoryByInterval(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).NetworkIds(networkIds).Serials(serials).ProductTypes(productTypes).Execute()

Return the memory utilization history in kB for devices in the organization.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 20. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 2 hours. If interval is provided, the timespan will be autocalculated. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 3600, 14400. The default is 300. Interval is calculated if time params are provided. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the result set by the included set of network IDs (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities history by device serial numbers (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesSystemMemoryUsageHistoryByInterval(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).NetworkIds(networkIds).Serials(serials).ProductTypes(productTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesSystemMemoryUsageHistoryByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesSystemMemoryUsageHistoryByInterval`: InlineResponse200286
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesSystemMemoryUsageHistoryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 20. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 2 hours. If interval is provided, the timespan will be autocalculated. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 3600, 14400. The default is 300. Interval is calculated if time params are provided. | 
 **networkIds** | **[]string** | Optional parameter to filter the result set by the included set of network IDs | 
 **serials** | **[]string** | Optional parameter to filter device availabilities history by device serial numbers | 
 **productTypes** | **[]string** | Optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. | 

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


## GetOrganizationDevicesUplinksAddressesByDevice

> []InlineResponse200287 GetOrganizationDevicesUplinksAddressesByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the current uplink addresses for devices in an organization.



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device uplinks by network ID. This filter uses multiple exact matches. (optional)
    productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device uplinks by device product types. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesUplinksAddressesByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesUplinksAddressesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesUplinksAddressesByDevice`: []InlineResponse200287
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesUplinksAddressesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesUplinksAddressesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device uplinks by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device uplinks by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]InlineResponse200287**](InlineResponse200287.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesUplinksLossAndLatency

> []InlineResponse200288 GetOrganizationDevicesUplinksLossAndLatency(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Uplink(uplink).Ip(ip).Execute()

Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 60 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 5 minutes after t0. The latest possible time that t1 can be is 2 minutes into the past. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 5 minutes. The default is 5 minutes. (optional)
    uplink := "uplink_example" // string | Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, wan3, cellular. Default will return all uplinks. (optional)
    ip := "ip_example" // string | Optional filter for a specific destination IP. Default will return all destination IPs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationDevicesUplinksLossAndLatency(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Uplink(uplink).Ip(ip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationDevicesUplinksLossAndLatency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesUplinksLossAndLatency`: []InlineResponse200288
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationDevicesUplinksLossAndLatency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesUplinksLossAndLatencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 60 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 5 minutes after t0. The latest possible time that t1 can be is 2 minutes into the past. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 5 minutes. The default is 5 minutes. | 
 **uplink** | **string** | Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, wan3, cellular. Default will return all uplinks. | 
 **ip** | **string** | Optional filter for a specific destination IP. Default will return all destination IPs. | 

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


## GetOrganizationFloorPlansAutoLocateDevices

> []InlineResponse200293 GetOrganizationFloorPlansAutoLocateDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).FloorPlanIds(floorPlanIds).Execute()

List auto locate details for each device in your organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 10000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more network IDs (optional)
    floorPlanIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more floorplan IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationFloorPlansAutoLocateDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).FloorPlanIds(floorPlanIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationFloorPlansAutoLocateDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationFloorPlansAutoLocateDevices`: []InlineResponse200293
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationFloorPlansAutoLocateDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationFloorPlansAutoLocateDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 10000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter devices by one or more network IDs | 
 **floorPlanIds** | **[]string** | Optional parameter to filter devices by one or more floorplan IDs | 

### Return type

[**[]InlineResponse200293**](InlineResponse200293.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevice

> InlineResponse200300 GetOrganizationInventoryDevice(ctx, organizationId, serial).Execute()

Return a single device from the inventory of an organization



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationInventoryDevice(context.Background(), organizationId, serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationInventoryDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryDevice`: InlineResponse200300
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationInventoryDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200300**](InlineResponse200300.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevices

> []InlineResponse200300 GetOrganizationInventoryDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).UsedState(usedState).Search(search).Macs(macs).NetworkIds(networkIds).Serials(serials).Models(models).OrderNumbers(orderNumbers).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).Execute()

Return the device inventory for an organization



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
    usedState := "usedState_example" // string | Filter results by used or unused inventory. Accepted values are 'used' or 'unused'. (optional)
    search := "search_example" // string | Search for devices in inventory based on serial number, mac address, or model. (optional)
    macs := []string{"Inner_example"} // []string | Search for devices in inventory based on mac addresses. (optional)
    networkIds := []string{"Inner_example"} // []string | Search for devices in inventory based on network ids. Use explicit 'null' value to get available devices only. (optional)
    serials := []string{"Inner_example"} // []string | Search for devices in inventory based on serials. (optional)
    models := []string{"Inner_example"} // []string | Search for devices in inventory based on model. (optional)
    orderNumbers := []string{"Inner_example"} // []string | Search for devices in inventory based on order numbers. (optional)
    tags := []string{"Inner_example"} // []string | Filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). (optional)
    tagsFilterType := "tagsFilterType_example" // string | To use with 'tags' parameter, to filter devices which contain ANY or ALL given tags. Accepted values are 'withAnyTags' or 'withAllTags', default is 'withAnyTags'. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Filter devices by product type. Accepted values are appliance, camera, campusGateway, cellularGateway, secureConnect, sensor, switch, systemsManager, wireless, and wirelessController. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationInventoryDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).UsedState(usedState).Search(search).Macs(macs).NetworkIds(networkIds).Serials(serials).Models(models).OrderNumbers(orderNumbers).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationInventoryDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryDevices`: []InlineResponse200300
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationInventoryDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **usedState** | **string** | Filter results by used or unused inventory. Accepted values are &#39;used&#39; or &#39;unused&#39;. | 
 **search** | **string** | Search for devices in inventory based on serial number, mac address, or model. | 
 **macs** | **[]string** | Search for devices in inventory based on mac addresses. | 
 **networkIds** | **[]string** | Search for devices in inventory based on network ids. Use explicit &#39;null&#39; value to get available devices only. | 
 **serials** | **[]string** | Search for devices in inventory based on serials. | 
 **models** | **[]string** | Search for devices in inventory based on model. | 
 **orderNumbers** | **[]string** | Search for devices in inventory based on order numbers. | 
 **tags** | **[]string** | Filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). | 
 **tagsFilterType** | **string** | To use with &#39;tags&#39; parameter, to filter devices which contain ANY or ALL given tags. Accepted values are &#39;withAnyTags&#39; or &#39;withAllTags&#39;, default is &#39;withAnyTags&#39;. | 
 **productTypes** | **[]string** | Filter devices by product type. Accepted values are appliance, camera, campusGateway, cellularGateway, secureConnect, sensor, switch, systemsManager, wireless, and wirelessController. | 

### Return type

[**[]InlineResponse200300**](InlineResponse200300.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevicesSwapsBulk

> InlineResponse207 GetOrganizationInventoryDevicesSwapsBulk(ctx, organizationId, id).Execute()

List of device swaps for a given request ID ({id}).



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationInventoryDevicesSwapsBulk(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationInventoryDevicesSwapsBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryDevicesSwapsBulk`: InlineResponse207
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationInventoryDevicesSwapsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDevicesSwapsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse207**](InlineResponse207.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopDevicesByUsage

> []InlineResponse200337 GetOrganizationSummaryTopDevicesByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 devices sorted by data usage over given time range



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
    networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
    deviceTag := "deviceTag_example" // string | Match result to an exact device tag (optional)
    quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. Maximum is 50 (optional)
    ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
    usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationSummaryTopDevicesByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationSummaryTopDevicesByUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummaryTopDevicesByUsage`: []InlineResponse200337
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationSummaryTopDevicesByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopDevicesByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. Maximum is 50 | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200337**](InlineResponse200337.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopDevicesModelsByUsage

> []InlineResponse200338 GetOrganizationSummaryTopDevicesModelsByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 device models sorted by data usage over given time range



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
    networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
    deviceTag := "deviceTag_example" // string | Match result to an exact device tag (optional)
    quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. Maximum is 50 (optional)
    ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
    usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationSummaryTopDevicesModelsByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationSummaryTopDevicesModelsByUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummaryTopDevicesModelsByUsage`: []InlineResponse200338
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationSummaryTopDevicesModelsByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopDevicesModelsByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. Maximum is 50 | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200338**](InlineResponse200338.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice

> InlineResponse200381 GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice(ctx, organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List wireless LAN controller layer 2 interfaces in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice(context.Background(), organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice`: InlineResponse200381
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200381**](InlineResponse200381.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice

> InlineResponse200382 GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice(ctx, organizationId).Serials(serials).IncludeInterfacesWithoutChanges(includeInterfacesWithoutChanges).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List wireless LAN controller layer 2 interfaces history status in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    includeInterfacesWithoutChanges := true // bool | By default, interfaces without changes are omitted from the response for brevity. If you want to include the interfaces even if they have no changes, set to true. (default: false) (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice(context.Background(), organizationId).Serials(serials).IncludeInterfacesWithoutChanges(includeInterfacesWithoutChanges).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice`: InlineResponse200382
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **includeInterfacesWithoutChanges** | **bool** | By default, interfaces without changes are omitted from the response for brevity. If you want to include the interfaces even if they have no changes, set to true. (default: false) | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200382**](InlineResponse200382.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval

> InlineResponse200383 GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval(ctx, organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List wireless LAN controller layer 2 interfaces history usage in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval(context.Background(), organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval`: InlineResponse200383
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200383**](InlineResponse200383.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice

> InlineResponse200384 GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice(ctx, organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List wireless LAN controller layer 3 interfaces in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice(context.Background(), organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice`: InlineResponse200384
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL3ByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesL3ByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200384**](InlineResponse200384.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice

> InlineResponse200385 GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice(ctx, organizationId).Serials(serials).IncludeInterfacesWithoutChanges(includeInterfacesWithoutChanges).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List wireless LAN controller layer 3 interfaces history status in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    includeInterfacesWithoutChanges := true // bool | By default, interfaces without changes are omitted from the response for brevity. If you want to include the interfaces even if they have no changes, set to true. (default: false) (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice(context.Background(), organizationId).Serials(serials).IncludeInterfacesWithoutChanges(includeInterfacesWithoutChanges).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice`: InlineResponse200385
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **includeInterfacesWithoutChanges** | **bool** | By default, interfaces without changes are omitted from the response for brevity. If you want to include the interfaces even if they have no changes, set to true. (default: false) | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200385**](InlineResponse200385.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval

> InlineResponse200386 GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval(ctx, organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List wireless LAN controller layer 3 interfaces history usage in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval(context.Background(), organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval`: InlineResponse200386
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200386**](InlineResponse200386.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice

> InlineResponse200387 GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice(ctx, organizationId).Serials(serials).Names(names).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Retrieve the packet counters for the interfaces of a Wireless LAN controller



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    names := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its interface name. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 1 day after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 hour. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice(context.Background(), organizationId).Serials(serials).Names(names).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice`: InlineResponse200387
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **names** | **[]string** | Optional parameter to filter wireless LAN controller by its interface name. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 1 day after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 hour. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200387**](InlineResponse200387.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval

> InlineResponse200388 GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval(ctx, organizationId).Serials(serials).Names(names).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Retrieve the traffic for the interfaces of a Wireless LAN controller



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    names := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its interface name. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval(context.Background(), organizationId).Serials(serials).Names(names).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval`: InlineResponse200388
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **names** | **[]string** | Optional parameter to filter wireless LAN controller by its interface name. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200388**](InlineResponse200388.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory

> []InlineResponse200389 GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory(ctx, organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the failover events of wireless LAN controllers in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory(context.Background(), organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory`: []InlineResponse200389
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessControllerDevicesRedundancyFailoverHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesRedundancyFailoverHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse200389**](InlineResponse200389.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesRedundancyStatuses

> InlineResponse200390 GetOrganizationWirelessControllerDevicesRedundancyStatuses(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List redundancy details of wireless LAN controllers in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud IDs. This filter uses multiple exact matches. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessControllerDevicesRedundancyStatuses(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessControllerDevicesRedundancyStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesRedundancyStatuses`: InlineResponse200390
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessControllerDevicesRedundancyStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesRedundancyStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud IDs. This filter uses multiple exact matches. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200390**](InlineResponse200390.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval

> InlineResponse200391 GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval(ctx, organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List cpu utilization data of wireless LAN controllers in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval(context.Background(), organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval`: InlineResponse200391
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200391**](InlineResponse200391.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationByDevice

> []InlineResponse200356 GetOrganizationWirelessDevicesChannelUtilizationByDevice(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get average channel utilization for all bands in a network, split by AP



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationByDevice(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesChannelUtilizationByDevice`: []InlineResponse200356
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]InlineResponse200356**](InlineResponse200356.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationByNetwork

> []InlineResponse200357 GetOrganizationWirelessDevicesChannelUtilizationByNetwork(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get average channel utilization across all bands for all networks in the organization



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationByNetwork(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesChannelUtilizationByNetwork`: []InlineResponse200357
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]InlineResponse200357**](InlineResponse200357.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval

> []InlineResponse200358 GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get a time-series of average channel utilization for all bands, segmented by device.



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval`: []InlineResponse200358
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]InlineResponse200358**](InlineResponse200358.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval

> []InlineResponse200359 GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get a time-series of average channel utilization for all bands



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval`: []InlineResponse200359
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]InlineResponse200359**](InlineResponse200359.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesEthernetStatuses

> []InlineResponse200360 GetOrganizationWirelessDevicesEthernetStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

List the most recent Ethernet link speed, duplex, aggregation and power mode and status information for wireless devices.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesEthernetStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesEthernetStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesEthernetStatuses`: []InlineResponse200360
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesEthernetStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesEthernetStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;L_3456 | 

### Return type

[**[]InlineResponse200360**](InlineResponse200360.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesPacketLossByClient

> []InlineResponse200361 GetOrganizationWirelessDevicesPacketLossByClient(ctx, organizationId).NetworkIds(networkIds).Ssids(ssids).Bands(bands).Macs(macs).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()

Get average packet loss for the given timespan for all clients in the organization.



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    ssids := []int32{int32(123)} // []int32 | Filter results by SSID number. (optional)
    bands := []string{"Inner_example"} // []string | Filter results by band. Valid bands are: 2.4, 5, and 6. (optional)
    macs := []string{"Inner_example"} // []string | Filter results by client mac address(es). (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesPacketLossByClient(context.Background(), organizationId).NetworkIds(networkIds).Ssids(ssids).Bands(bands).Macs(macs).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesPacketLossByClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesPacketLossByClient`: []InlineResponse200361
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesPacketLossByClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesPacketLossByClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **ssids** | **[]int32** | Filter results by SSID number. | 
 **bands** | **[]string** | Filter results by band. Valid bands are: 2.4, 5, and 6. | 
 **macs** | **[]string** | Filter results by client mac address(es). | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. | 

### Return type

[**[]InlineResponse200361**](InlineResponse200361.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesPacketLossByDevice

> []InlineResponse200362 GetOrganizationWirelessDevicesPacketLossByDevice(ctx, organizationId).NetworkIds(networkIds).Serials(serials).Ssids(ssids).Bands(bands).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()

Get average packet loss for the given timespan for all devices in the organization



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
    ssids := []int32{int32(123)} // []int32 | Filter results by SSID number. (optional)
    bands := []string{"Inner_example"} // []string | Filter results by band. Valid bands are: 2.4, 5, and 6. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesPacketLossByDevice(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).Ssids(ssids).Bands(bands).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesPacketLossByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesPacketLossByDevice`: []InlineResponse200362
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesPacketLossByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesPacketLossByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **ssids** | **[]int32** | Filter results by SSID number. | 
 **bands** | **[]string** | Filter results by band. Valid bands are: 2.4, 5, and 6. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. | 

### Return type

[**[]InlineResponse200362**](InlineResponse200362.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesPacketLossByNetwork

> []InlineResponse200363 GetOrganizationWirelessDevicesPacketLossByNetwork(ctx, organizationId).NetworkIds(networkIds).Serials(serials).Ssids(ssids).Bands(bands).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()

Get average packet loss for the given timespan for all networks in the organization.



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
    ssids := []int32{int32(123)} // []int32 | Filter results by SSID number. (optional)
    bands := []string{"Inner_example"} // []string | Filter results by band. Valid bands are: 2.4, 5, and 6. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesPacketLossByNetwork(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).Ssids(ssids).Bands(bands).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesPacketLossByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesPacketLossByNetwork`: []InlineResponse200363
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesPacketLossByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesPacketLossByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **ssids** | **[]int32** | Filter results by SSID number. | 
 **bands** | **[]string** | Filter results by band. Valid bands are: 2.4, 5, and 6. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. | 

### Return type

[**[]InlineResponse200363**](InlineResponse200363.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesPowerModeHistory

> InlineResponse200364 GetOrganizationWirelessDevicesPowerModeHistory(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()

Return a record of power mode changes for wireless devices in the organization



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 1 day after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 day. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 20. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the result set by the included set of network IDs (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities history by device serial numbers (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesPowerModeHistory(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesPowerModeHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesPowerModeHistory`: InlineResponse200364
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesPowerModeHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesPowerModeHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 1 day after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 20. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter the result set by the included set of network IDs | 
 **serials** | **[]string** | Optional parameter to filter device availabilities history by device serial numbers | 

### Return type

[**InlineResponse200364**](InlineResponse200364.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthorities

> []InlineResponse200365 GetOrganizationWirelessDevicesRadsecCertificatesAuthorities(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

Query for details on the organization's RADSEC device Certificate Authority certificates (CAs)



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
    certificateAuthorityIds := []string{"Inner_example"} // []string | Optional parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthorities(context.Background(), organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthorities`: []InlineResponse200365
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateAuthorityIds** | **[]string** | Optional parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. | 

### Return type

[**[]InlineResponse200365**](InlineResponse200365.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls

> InlineResponse200367 GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

Query for certificate revocation list (CRL) for the organization's RADSEC device Certificate Authorities (CAs).



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
    certificateAuthorityIds := []string{"Inner_example"} // []string | Optional parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls(context.Background(), organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls`: InlineResponse200367
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateAuthorityIds** | **[]string** | Optional parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. | 

### Return type

[**InlineResponse200367**](InlineResponse200367.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas

> InlineResponse200367 GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

Query for all delta certificate revocation list (CRL) for the organization's RADSEC device Certificate Authority (CA) with the given id.



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
    certificateAuthorityIds := []string{"Inner_example"} // []string | Parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas(context.Background(), organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas`: InlineResponse200367
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateAuthorityIds** | **[]string** | Parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. | 

### Return type

[**InlineResponse200367**](InlineResponse200367.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesSystemCpuLoadHistory

> InlineResponse200368 GetOrganizationWirelessDevicesSystemCpuLoadHistory(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()

Return the CPU Load history for a list of wireless devices in the organization.



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 1 day after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 day. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 20. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the result set by the included set of network IDs (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities history by device serial numbers (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesSystemCpuLoadHistory(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesSystemCpuLoadHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesSystemCpuLoadHistory`: InlineResponse200368
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesSystemCpuLoadHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 1 day after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 20. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter the result set by the included set of network IDs | 
 **serials** | **[]string** | Optional parameter to filter device availabilities history by device serial numbers | 

### Return type

[**InlineResponse200368**](InlineResponse200368.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesWirelessControllersByDevice

> InlineResponse200369 GetOrganizationWirelessDevicesWirelessControllersByDevice(ctx, organizationId).NetworkIds(networkIds).Serials(serials).ControllerSerials(controllerSerials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List of Catalyst access points information



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter access points by network ID. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter access points by its cloud ID. This filter uses multiple exact matches. (optional)
    controllerSerials := []string{"Inner_example"} // []string | Optional parameter to filter access points by its wireless LAN controller cloud ID. This filter uses multiple exact matches. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessDevicesWirelessControllersByDevice(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).ControllerSerials(controllerSerials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessDevicesWirelessControllersByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesWirelessControllersByDevice`: InlineResponse200369
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessDevicesWirelessControllersByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter access points by network ID. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter access points by its cloud ID. This filter uses multiple exact matches. | 
 **controllerSerials** | **[]string** | Optional parameter to filter access points by its wireless LAN controller cloud ID. This filter uses multiple exact matches. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200369**](InlineResponse200369.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessZigbeeDevices

> []InlineResponse200376 GetOrganizationWirelessZigbeeDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).IsEnrolled(isEnrolled).Search(search).Execute()

List the Zigbee wireless devices for an organization or the supplied network(s)



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Parameter of networks you want the zigbee devices for. E.g.: networkIds[]=N_12345678&networkIds[]=N_3456 (optional)
    isEnrolled := true // bool | Filter devices based on if they are enrolled or not (optional)
    search := "search_example" // string | Filter devices by their name, tag or serial (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetOrganizationWirelessZigbeeDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).IsEnrolled(isEnrolled).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetOrganizationWirelessZigbeeDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessZigbeeDevices`: []InlineResponse200376
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetOrganizationWirelessZigbeeDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessZigbeeDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Parameter of networks you want the zigbee devices for. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;N_3456 | 
 **isEnrolled** | **bool** | Filter devices based on if they are enrolled or not | 
 **search** | **string** | Filter devices by their name, tag or serial | 

### Return type

[**[]InlineResponse200376**](InlineResponse200376.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallNetworkSmDeviceApps

> map[string]interface{} InstallNetworkSmDeviceApps(ctx, networkId, deviceId).InstallNetworkSmDeviceApps(installNetworkSmDeviceApps).Execute()

Install applications on a device



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
    deviceId := "deviceId_example" // string | Device ID
    installNetworkSmDeviceApps := *openapiclient.NewInlineObject133([]string{"AppIds_example"}) // InlineObject133 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.InstallNetworkSmDeviceApps(context.Background(), networkId, deviceId).InstallNetworkSmDeviceApps(installNetworkSmDeviceApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.InstallNetworkSmDeviceApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallNetworkSmDeviceApps`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.InstallNetworkSmDeviceApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallNetworkSmDeviceAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **installNetworkSmDeviceApps** | [**InlineObject133**](InlineObject133.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockNetworkSmDevices

> InlineResponse200126 LockNetworkSmDevices(ctx, networkId).LockNetworkSmDevices(lockNetworkSmDevices).Execute()

Lock a set of devices



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
    lockNetworkSmDevices := *openapiclient.NewInlineObject127() // InlineObject127 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.LockNetworkSmDevices(context.Background(), networkId).LockNetworkSmDevices(lockNetworkSmDevices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.LockNetworkSmDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockNetworkSmDevices`: InlineResponse200126
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.LockNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockNetworkSmDevices** | [**InlineObject127**](InlineObject127.md) |  | 

### Return type

[**InlineResponse200126**](InlineResponse200126.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyNetworkSmDevicesTags

> []InlineResponse200128 ModifyNetworkSmDevicesTags(ctx, networkId).ModifyNetworkSmDevicesTags(modifyNetworkSmDevicesTags).Execute()

Add, delete, or update the tags of a set of devices



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
    modifyNetworkSmDevicesTags := *openapiclient.NewInlineObject128([]string{"Tags_example"}, "UpdateAction_example") // InlineObject128 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ModifyNetworkSmDevicesTags(context.Background(), networkId).ModifyNetworkSmDevicesTags(modifyNetworkSmDevicesTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ModifyNetworkSmDevicesTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyNetworkSmDevicesTags`: []InlineResponse200128
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.ModifyNetworkSmDevicesTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyNetworkSmDevicesTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyNetworkSmDevicesTags** | [**InlineObject128**](InlineObject128.md) |  | 

### Return type

[**[]InlineResponse200128**](InlineResponse200128.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveNetworkSmDevices

> InlineResponse200129 MoveNetworkSmDevices(ctx, networkId).MoveNetworkSmDevices(moveNetworkSmDevices).Execute()

Move a set of devices to a new network



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
    moveNetworkSmDevices := *openapiclient.NewInlineObject129("NewNetwork_example") // InlineObject129 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.MoveNetworkSmDevices(context.Background(), networkId).MoveNetworkSmDevices(moveNetworkSmDevices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.MoveNetworkSmDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveNetworkSmDevices`: InlineResponse200129
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.MoveNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveNetworkSmDevices** | [**InlineObject129**](InlineObject129.md) |  | 

### Return type

[**InlineResponse200129**](InlineResponse200129.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootDevice

> InlineResponse2023 RebootDevice(ctx, serial).Execute()

Reboot a device



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
    resp, r, err := apiClient.DevicesApi.RebootDevice(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.RebootDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootDevice`: InlineResponse2023
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.RebootDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2023**](InlineResponse2023.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootNetworkSmDevices

> InlineResponse200130 RebootNetworkSmDevices(ctx, networkId).RebootNetworkSmDevices(rebootNetworkSmDevices).Execute()

Reboot a set of endpoints



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
    rebootNetworkSmDevices := *openapiclient.NewInlineObject130() // InlineObject130 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.RebootNetworkSmDevices(context.Background(), networkId).RebootNetworkSmDevices(rebootNetworkSmDevices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.RebootNetworkSmDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootNetworkSmDevices`: InlineResponse200130
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.RebootNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rebootNetworkSmDevices** | [**InlineObject130**](InlineObject130.md) |  | 

### Return type

[**InlineResponse200130**](InlineResponse200130.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshNetworkSmDeviceDetails

> map[string]interface{} RefreshNetworkSmDeviceDetails(ctx, networkId, deviceId).Execute()

Refresh the details of a device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.RefreshNetworkSmDeviceDetails(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.RefreshNetworkSmDeviceDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshNetworkSmDeviceDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.RefreshNetworkSmDeviceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshNetworkSmDeviceDetailsRequest struct via the builder pattern


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


## RemoveNetworkDevices

> RemoveNetworkDevices(ctx, networkId).RemoveNetworkDevices(removeNetworkDevices).Execute()

Remove a single device



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
    removeNetworkDevices := *openapiclient.NewInlineObject97("Serial_example") // InlineObject97 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.RemoveNetworkDevices(context.Background(), networkId).RemoveNetworkDevices(removeNetworkDevices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.RemoveNetworkDevices``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveNetworkDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeNetworkDevices** | [**InlineObject97**](InlineObject97.md) |  | 

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


## ReorderOrganizationDevicesPacketCaptureSchedules

> InlineResponse200281 ReorderOrganizationDevicesPacketCaptureSchedules(ctx, organizationId).ReorderOrganizationDevicesPacketCaptureSchedules(reorderOrganizationDevicesPacketCaptureSchedules).Execute()

Bulk update priorities of pcap schedules



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
    reorderOrganizationDevicesPacketCaptureSchedules := *openapiclient.NewInlineObject263([]openapiclient.OrganizationsOrganizationIdDevicesPacketCaptureSchedulesReorderOrder{*openapiclient.NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesReorderOrder()}) // InlineObject263 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ReorderOrganizationDevicesPacketCaptureSchedules(context.Background(), organizationId).ReorderOrganizationDevicesPacketCaptureSchedules(reorderOrganizationDevicesPacketCaptureSchedules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ReorderOrganizationDevicesPacketCaptureSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReorderOrganizationDevicesPacketCaptureSchedules`: InlineResponse200281
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.ReorderOrganizationDevicesPacketCaptureSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReorderOrganizationDevicesPacketCaptureSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reorderOrganizationDevicesPacketCaptureSchedules** | [**InlineObject263**](InlineObject263.md) |  | 

### Return type

[**InlineResponse200281**](InlineResponse200281.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShutdownNetworkSmDevices

> InlineResponse200130 ShutdownNetworkSmDevices(ctx, networkId).ShutdownNetworkSmDevices(shutdownNetworkSmDevices).Execute()

Shutdown a set of endpoints



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
    shutdownNetworkSmDevices := *openapiclient.NewInlineObject131() // InlineObject131 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ShutdownNetworkSmDevices(context.Background(), networkId).ShutdownNetworkSmDevices(shutdownNetworkSmDevices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ShutdownNetworkSmDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShutdownNetworkSmDevices`: InlineResponse200130
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.ShutdownNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shutdownNetworkSmDevices** | [**InlineObject131**](InlineObject131.md) |  | 

### Return type

[**InlineResponse200130**](InlineResponse200130.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopOrganizationDevicesPacketCaptureCapture

> InlineResponse200278Items StopOrganizationDevicesPacketCaptureCapture(ctx, organizationId, captureId).StopOrganizationDevicesPacketCaptureCapture(stopOrganizationDevicesPacketCaptureCapture).Execute()

Stop a specific packet capture (not supported for Catalyst devices)



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
    captureId := "captureId_example" // string | Capture ID
    stopOrganizationDevicesPacketCaptureCapture := *openapiclient.NewInlineObject261([]string{"Serials_example"}) // InlineObject261 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.StopOrganizationDevicesPacketCaptureCapture(context.Background(), organizationId, captureId).StopOrganizationDevicesPacketCaptureCapture(stopOrganizationDevicesPacketCaptureCapture).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.StopOrganizationDevicesPacketCaptureCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopOrganizationDevicesPacketCaptureCapture`: InlineResponse200278Items
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.StopOrganizationDevicesPacketCaptureCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**captureId** | **string** | Capture ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopOrganizationDevicesPacketCaptureCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stopOrganizationDevicesPacketCaptureCapture** | [**InlineObject261**](InlineObject261.md) |  | 

### Return type

[**InlineResponse200278Items**](InlineResponse200278Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnenrollNetworkSmDevice

> InlineResponse200143 UnenrollNetworkSmDevice(ctx, networkId, deviceId).Execute()

Unenroll a device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.UnenrollNetworkSmDevice(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UnenrollNetworkSmDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnenrollNetworkSmDevice`: InlineResponse200143
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UnenrollNetworkSmDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnenrollNetworkSmDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200143**](InlineResponse200143.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallNetworkSmDeviceApps

> map[string]interface{} UninstallNetworkSmDeviceApps(ctx, networkId, deviceId).UninstallNetworkSmDeviceApps(uninstallNetworkSmDeviceApps).Execute()

Uninstall applications on a device



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
    deviceId := "deviceId_example" // string | Device ID
    uninstallNetworkSmDeviceApps := *openapiclient.NewInlineObject134([]string{"AppIds_example"}) // InlineObject134 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.UninstallNetworkSmDeviceApps(context.Background(), networkId, deviceId).UninstallNetworkSmDeviceApps(uninstallNetworkSmDeviceApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UninstallNetworkSmDeviceApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UninstallNetworkSmDeviceApps`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UninstallNetworkSmDeviceApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallNetworkSmDeviceAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uninstallNetworkSmDeviceApps** | [**InlineObject134**](InlineObject134.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> InlineResponse2006 UpdateDevice(ctx, serial).UpdateDevice(updateDevice).Execute()

Update the attributes of a device



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
    updateDevice := *openapiclient.NewInlineObject3() // InlineObject3 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.UpdateDevice(context.Background(), serial).UpdateDevice(updateDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDevice`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDevice** | [**InlineObject3**](InlineObject3.md) |  | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularSims

> InlineResponse20018 UpdateDeviceCellularSims(ctx, serial).UpdateDeviceCellularSims(updateDeviceCellularSims).Execute()

Updates the SIM and APN configurations for a cellular device.



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
    updateDeviceCellularSims := *openapiclient.NewInlineObject13() // InlineObject13 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.UpdateDeviceCellularSims(context.Background(), serial).UpdateDeviceCellularSims(updateDeviceCellularSims).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateDeviceCellularSims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCellularSims`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateDeviceCellularSims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularSims** | [**InlineObject13**](InlineObject13.md) |  | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceManagementInterface

> InlineResponse20033 UpdateDeviceManagementInterface(ctx, serial).UpdateDeviceManagementInterface(updateDeviceManagementInterface).Execute()

Update the management interface settings for a device



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
    updateDeviceManagementInterface := *openapiclient.NewInlineObject25() // InlineObject25 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.UpdateDeviceManagementInterface(context.Background(), serial).UpdateDeviceManagementInterface(updateDeviceManagementInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateDeviceManagementInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceManagementInterface`: InlineResponse20033
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateDeviceManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceManagementInterface** | [**InlineObject25**](InlineObject25.md) |  | 

### Return type

[**InlineResponse20033**](InlineResponse20033.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSmDevicesFields

> []InlineResponse200127 UpdateNetworkSmDevicesFields(ctx, networkId).UpdateNetworkSmDevicesFields(updateNetworkSmDevicesFields).Execute()

Modify the fields of a device



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
    updateNetworkSmDevicesFields := *openapiclient.NewInlineObject126(*openapiclient.NewNetworksNetworkIdSmDevicesFieldsDeviceFields()) // InlineObject126 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.UpdateNetworkSmDevicesFields(context.Background(), networkId).UpdateNetworkSmDevicesFields(updateNetworkSmDevicesFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateNetworkSmDevicesFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSmDevicesFields`: []InlineResponse200127
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateNetworkSmDevicesFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSmDevicesFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSmDevicesFields** | [**InlineObject126**](InlineObject126.md) |  | 

### Return type

[**[]InlineResponse200127**](InlineResponse200127.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationDevicesPacketCaptureSchedule

> InlineResponse200280Items UpdateOrganizationDevicesPacketCaptureSchedule(ctx, organizationId, scheduleId).UpdateOrganizationDevicesPacketCaptureSchedule(updateOrganizationDevicesPacketCaptureSchedule).Execute()

Update a schedule for packet capture



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
    scheduleId := "scheduleId_example" // string | Schedule ID
    updateOrganizationDevicesPacketCaptureSchedule := *openapiclient.NewInlineObject264([]openapiclient.OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices{*openapiclient.NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices()}) // InlineObject264 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.UpdateOrganizationDevicesPacketCaptureSchedule(context.Background(), organizationId, scheduleId).UpdateOrganizationDevicesPacketCaptureSchedule(updateOrganizationDevicesPacketCaptureSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateOrganizationDevicesPacketCaptureSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationDevicesPacketCaptureSchedule`: InlineResponse200280Items
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateOrganizationDevicesPacketCaptureSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**scheduleId** | **string** | Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationDevicesPacketCaptureScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationDevicesPacketCaptureSchedule** | [**InlineObject264**](InlineObject264.md) |  | 

### Return type

[**InlineResponse200280Items**](InlineResponse200280Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities

> InlineResponse200366 UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(ctx, organizationId).UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(updateOrganizationWirelessDevicesRadsecCertificatesAuthorities).Execute()

Update an organization's RADSEC device Certificate Authority (CA) state



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
    updateOrganizationWirelessDevicesRadsecCertificatesAuthorities := *openapiclient.NewInlineObject303() // InlineObject303 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(context.Background(), organizationId).UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(updateOrganizationWirelessDevicesRadsecCertificatesAuthorities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities`: InlineResponse200366
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationWirelessDevicesRadsecCertificatesAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationWirelessDevicesRadsecCertificatesAuthorities** | [**InlineObject303**](InlineObject303.md) |  | 

### Return type

[**InlineResponse200366**](InlineResponse200366.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessZigbeeDevice

> InlineResponse200376 UpdateOrganizationWirelessZigbeeDevice(ctx, organizationId, id).UpdateOrganizationWirelessZigbeeDevice(updateOrganizationWirelessZigbeeDevice).Execute()

Endpoint to update zigbee gateways



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
    id := "id_example" // string | ID
    updateOrganizationWirelessZigbeeDevice := *openapiclient.NewInlineObject309(false) // InlineObject309 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.UpdateOrganizationWirelessZigbeeDevice(context.Background(), organizationId, id).UpdateOrganizationWirelessZigbeeDevice(updateOrganizationWirelessZigbeeDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateOrganizationWirelessZigbeeDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessZigbeeDevice`: InlineResponse200376
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateOrganizationWirelessZigbeeDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationWirelessZigbeeDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationWirelessZigbeeDevice** | [**InlineObject309**](InlineObject309.md) |  | 

### Return type

[**InlineResponse200376**](InlineResponse200376.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmxNetworkDevicesClaim

> InlineResponse20096 VmxNetworkDevicesClaim(ctx, networkId).VmxNetworkDevicesClaim(vmxNetworkDevicesClaim).Execute()

Claim a vMX into a network



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
    vmxNetworkDevicesClaim := *openapiclient.NewInlineObject96("Size_example") // InlineObject96 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.VmxNetworkDevicesClaim(context.Background(), networkId).VmxNetworkDevicesClaim(vmxNetworkDevicesClaim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.VmxNetworkDevicesClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmxNetworkDevicesClaim`: InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.VmxNetworkDevicesClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmxNetworkDevicesClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmxNetworkDevicesClaim** | [**InlineObject96**](InlineObject96.md) |  | 

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


## WipeNetworkSmDevices

> InlineResponse200131 WipeNetworkSmDevices(ctx, networkId).WipeNetworkSmDevices(wipeNetworkSmDevices).Execute()

Wipe a device



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
    wipeNetworkSmDevices := *openapiclient.NewInlineObject132() // InlineObject132 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.WipeNetworkSmDevices(context.Background(), networkId).WipeNetworkSmDevices(wipeNetworkSmDevices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.WipeNetworkSmDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WipeNetworkSmDevices`: InlineResponse200131
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.WipeNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWipeNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wipeNetworkSmDevices** | [**InlineObject132**](InlineObject132.md) |  | 

### Return type

[**InlineResponse200131**](InlineResponse200131.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

