# InlineResponse20026Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | Interface IP address | [optional] 
**Name** | Pointer to **string** | Interface name | [optional] 
**Subnet** | Pointer to **string** | Subnet containing the interface | [optional] 
**Flags** | Pointer to **[]string** | List of flags (unordered) | [optional] 
**Neighbors** | Pointer to **[]string** | Array of PIM Neighbor IP Addresses | [optional] 

## Methods

### NewInlineResponse20026Interfaces

`func NewInlineResponse20026Interfaces() *InlineResponse20026Interfaces`

NewInlineResponse20026Interfaces instantiates a new InlineResponse20026Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20026InterfacesWithDefaults

`func NewInlineResponse20026InterfacesWithDefaults() *InlineResponse20026Interfaces`

NewInlineResponse20026InterfacesWithDefaults instantiates a new InlineResponse20026Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InlineResponse20026Interfaces) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse20026Interfaces) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse20026Interfaces) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse20026Interfaces) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20026Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20026Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20026Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20026Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20026Interfaces) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20026Interfaces) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20026Interfaces) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20026Interfaces) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetFlags

`func (o *InlineResponse20026Interfaces) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *InlineResponse20026Interfaces) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *InlineResponse20026Interfaces) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *InlineResponse20026Interfaces) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetNeighbors

`func (o *InlineResponse20026Interfaces) GetNeighbors() []string`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *InlineResponse20026Interfaces) GetNeighborsOk() (*[]string, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *InlineResponse20026Interfaces) SetNeighbors(v []string)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *InlineResponse20026Interfaces) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


