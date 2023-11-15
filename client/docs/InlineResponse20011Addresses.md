# InlineResponse20011Addresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | The IP protocol used for the address | [optional] 
**AssignmentMode** | Pointer to **string** | The type of address assignment. Either static or dynamic. | [optional] 
**Address** | Pointer to **string** | The IP address configured for the alternate management interface | [optional] 
**Gateway** | Pointer to **string** | The gateway address configured for the alternate managment interface | [optional] 
**Prefix** | Pointer to **string** | The IPv6 prefix of the interface. Required if IPv6 object is included. | [optional] 
**Nameservers** | Pointer to [**DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers**](DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers.md) |  | [optional] 

## Methods

### NewInlineResponse20011Addresses

`func NewInlineResponse20011Addresses() *InlineResponse20011Addresses`

NewInlineResponse20011Addresses instantiates a new InlineResponse20011Addresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20011AddressesWithDefaults

`func NewInlineResponse20011AddressesWithDefaults() *InlineResponse20011Addresses`

NewInlineResponse20011AddressesWithDefaults instantiates a new InlineResponse20011Addresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse20011Addresses) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20011Addresses) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20011Addresses) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20011Addresses) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAssignmentMode

`func (o *InlineResponse20011Addresses) GetAssignmentMode() string`

GetAssignmentMode returns the AssignmentMode field if non-nil, zero value otherwise.

### GetAssignmentModeOk

`func (o *InlineResponse20011Addresses) GetAssignmentModeOk() (*string, bool)`

GetAssignmentModeOk returns a tuple with the AssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentMode

`func (o *InlineResponse20011Addresses) SetAssignmentMode(v string)`

SetAssignmentMode sets AssignmentMode field to given value.

### HasAssignmentMode

`func (o *InlineResponse20011Addresses) HasAssignmentMode() bool`

HasAssignmentMode returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse20011Addresses) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse20011Addresses) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse20011Addresses) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse20011Addresses) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *InlineResponse20011Addresses) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse20011Addresses) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse20011Addresses) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InlineResponse20011Addresses) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPrefix

`func (o *InlineResponse20011Addresses) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *InlineResponse20011Addresses) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *InlineResponse20011Addresses) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *InlineResponse20011Addresses) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetNameservers

`func (o *InlineResponse20011Addresses) GetNameservers() DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineResponse20011Addresses) GetNameserversOk() (*DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineResponse20011Addresses) SetNameservers(v DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *InlineResponse20011Addresses) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

