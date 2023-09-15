# DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | The IP protocol used for the address | [optional] 
**AssignmentMode** | Pointer to **string** | The type of address assignment. Either static or dynamic. | [optional] 
**Address** | Pointer to **string** | The IP address configured for the alternate management interface | [optional] 
**Gateway** | Pointer to **string** | The gateway address configured for the alternate managment interface | [optional] 
**Prefix** | Pointer to **string** | The IPv6 prefix length of the IPv6 interface. Required if IPv6 object is included. | [optional] 
**Nameservers** | Pointer to [**DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers**](DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers.md) |  | [optional] 

## Methods

### NewDevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses

`func NewDevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses() *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses`

NewDevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses instantiates a new DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialWirelessAlternateManagementInterfaceIpv6AddressesWithDefaults

`func NewDevicesSerialWirelessAlternateManagementInterfaceIpv6AddressesWithDefaults() *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses`

NewDevicesSerialWirelessAlternateManagementInterfaceIpv6AddressesWithDefaults instantiates a new DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAssignmentMode

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetAssignmentMode() string`

GetAssignmentMode returns the AssignmentMode field if non-nil, zero value otherwise.

### GetAssignmentModeOk

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetAssignmentModeOk() (*string, bool)`

GetAssignmentModeOk returns a tuple with the AssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentMode

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) SetAssignmentMode(v string)`

SetAssignmentMode sets AssignmentMode field to given value.

### HasAssignmentMode

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) HasAssignmentMode() bool`

HasAssignmentMode returns a boolean if a field has been set.

### GetAddress

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPrefix

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetNameservers

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetNameservers() DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) GetNameserversOk() (*DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) SetNameservers(v DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *DevicesSerialWirelessAlternateManagementInterfaceIpv6Addresses) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


