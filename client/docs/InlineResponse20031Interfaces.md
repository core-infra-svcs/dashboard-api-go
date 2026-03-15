# InlineResponse20031Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | Interface IP address | [optional] 
**Name** | Pointer to **string** | Interface name | [optional] 
**Vrf** | Pointer to **string** | VRF name for the interface. Included on networks with IOS XE 17.18 or higher | [optional] 
**VrfType** | Pointer to **string** | VRF type for the interface. Included on networks with IOS XE 17.18 or higher | [optional] 
**IpVersion** | Pointer to **string** | IP version for the interface. Included on networks with IOS XE 17.18 or higher | [optional] 
**Subnet** | Pointer to **string** | Subnet containing the interface | [optional] 
**Flags** | Pointer to **[]string** | List of flags (unordered) | [optional] 
**Neighbors** | Pointer to **[]string** | Array of PIM Neighbor IP Addresses | [optional] 

## Methods

### NewInlineResponse20031Interfaces

`func NewInlineResponse20031Interfaces() *InlineResponse20031Interfaces`

NewInlineResponse20031Interfaces instantiates a new InlineResponse20031Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20031InterfacesWithDefaults

`func NewInlineResponse20031InterfacesWithDefaults() *InlineResponse20031Interfaces`

NewInlineResponse20031InterfacesWithDefaults instantiates a new InlineResponse20031Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InlineResponse20031Interfaces) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse20031Interfaces) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse20031Interfaces) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse20031Interfaces) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20031Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20031Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20031Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20031Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse20031Interfaces) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse20031Interfaces) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse20031Interfaces) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse20031Interfaces) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetVrfType

`func (o *InlineResponse20031Interfaces) GetVrfType() string`

GetVrfType returns the VrfType field if non-nil, zero value otherwise.

### GetVrfTypeOk

`func (o *InlineResponse20031Interfaces) GetVrfTypeOk() (*string, bool)`

GetVrfTypeOk returns a tuple with the VrfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfType

`func (o *InlineResponse20031Interfaces) SetVrfType(v string)`

SetVrfType sets VrfType field to given value.

### HasVrfType

`func (o *InlineResponse20031Interfaces) HasVrfType() bool`

HasVrfType returns a boolean if a field has been set.

### GetIpVersion

`func (o *InlineResponse20031Interfaces) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineResponse20031Interfaces) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineResponse20031Interfaces) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *InlineResponse20031Interfaces) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20031Interfaces) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20031Interfaces) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20031Interfaces) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20031Interfaces) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetFlags

`func (o *InlineResponse20031Interfaces) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *InlineResponse20031Interfaces) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *InlineResponse20031Interfaces) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *InlineResponse20031Interfaces) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetNeighbors

`func (o *InlineResponse20031Interfaces) GetNeighbors() []string`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *InlineResponse20031Interfaces) GetNeighborsOk() (*[]string, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *InlineResponse20031Interfaces) SetNeighbors(v []string)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *InlineResponse20031Interfaces) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


