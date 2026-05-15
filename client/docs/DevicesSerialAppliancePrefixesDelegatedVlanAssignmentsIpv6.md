# DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **NullableString** | Assigned IPv6 subnet prefix. | [optional] 
**Address** | **string** | IPv6 address assigned to the MX on the VLAN. | 
**LinkLocal** | Pointer to [**DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6LinkLocal**](DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6LinkLocal.md) |  | [optional] 
**SolicitedNodeMulticast** | Pointer to [**DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6SolicitedNodeMulticast**](DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6SolicitedNodeMulticast.md) |  | [optional] 

## Methods

### NewDevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6

`func NewDevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6(address string, ) *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6`

NewDevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6 instantiates a new DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6WithDefaults

`func NewDevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6WithDefaults() *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6`

NewDevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6WithDefaults instantiates a new DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetAddress

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetLinkLocal

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) GetLinkLocal() DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6LinkLocal`

GetLinkLocal returns the LinkLocal field if non-nil, zero value otherwise.

### GetLinkLocalOk

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) GetLinkLocalOk() (*DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6LinkLocal, bool)`

GetLinkLocalOk returns a tuple with the LinkLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkLocal

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) SetLinkLocal(v DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6LinkLocal)`

SetLinkLocal sets LinkLocal field to given value.

### HasLinkLocal

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) HasLinkLocal() bool`

HasLinkLocal returns a boolean if a field has been set.

### GetSolicitedNodeMulticast

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) GetSolicitedNodeMulticast() DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6SolicitedNodeMulticast`

GetSolicitedNodeMulticast returns the SolicitedNodeMulticast field if non-nil, zero value otherwise.

### GetSolicitedNodeMulticastOk

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) GetSolicitedNodeMulticastOk() (*DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6SolicitedNodeMulticast, bool)`

GetSolicitedNodeMulticastOk returns a tuple with the SolicitedNodeMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolicitedNodeMulticast

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) SetSolicitedNodeMulticast(v DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6SolicitedNodeMulticast)`

SetSolicitedNodeMulticast sets SolicitedNodeMulticast field to given value.

### HasSolicitedNodeMulticast

`func (o *DevicesSerialAppliancePrefixesDelegatedVlanAssignmentsIpv6) HasSolicitedNodeMulticast() bool`

HasSolicitedNodeMulticast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


