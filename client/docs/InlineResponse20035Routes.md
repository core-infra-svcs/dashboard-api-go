# InlineResponse20035Routes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | Source IP or \&quot;Any\&quot; | [optional] 
**Group** | Pointer to **string** | Group IP address | [optional] 
**Vrf** | Pointer to **string** | VRF name for the route. Included on networks with IOS XE 17.18 or higher | [optional] 
**IpVersion** | Pointer to **string** | IP version for the route. Included on networks with IOS XE 17.18 or higher | [optional] 
**RendezvousPoint** | Pointer to **string** | IP address of the rendezvous point | [optional] 
**IncomingInterfaceName** | Pointer to **string** | Name of the Virtual Interface traffic is arriving on | [optional] 
**OutgoingInterfaceNames** | Pointer to **[]string** | Array of outgoing Virtual Interface names | [optional] 
**Flags** | Pointer to **[]string** | List of flags (unordered) | [optional] 

## Methods

### NewInlineResponse20035Routes

`func NewInlineResponse20035Routes() *InlineResponse20035Routes`

NewInlineResponse20035Routes instantiates a new InlineResponse20035Routes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20035RoutesWithDefaults

`func NewInlineResponse20035RoutesWithDefaults() *InlineResponse20035Routes`

NewInlineResponse20035RoutesWithDefaults instantiates a new InlineResponse20035Routes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *InlineResponse20035Routes) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse20035Routes) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse20035Routes) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InlineResponse20035Routes) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetGroup

`func (o *InlineResponse20035Routes) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse20035Routes) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse20035Routes) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse20035Routes) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse20035Routes) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse20035Routes) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse20035Routes) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse20035Routes) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetIpVersion

`func (o *InlineResponse20035Routes) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineResponse20035Routes) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineResponse20035Routes) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *InlineResponse20035Routes) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetRendezvousPoint

`func (o *InlineResponse20035Routes) GetRendezvousPoint() string`

GetRendezvousPoint returns the RendezvousPoint field if non-nil, zero value otherwise.

### GetRendezvousPointOk

`func (o *InlineResponse20035Routes) GetRendezvousPointOk() (*string, bool)`

GetRendezvousPointOk returns a tuple with the RendezvousPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendezvousPoint

`func (o *InlineResponse20035Routes) SetRendezvousPoint(v string)`

SetRendezvousPoint sets RendezvousPoint field to given value.

### HasRendezvousPoint

`func (o *InlineResponse20035Routes) HasRendezvousPoint() bool`

HasRendezvousPoint returns a boolean if a field has been set.

### GetIncomingInterfaceName

`func (o *InlineResponse20035Routes) GetIncomingInterfaceName() string`

GetIncomingInterfaceName returns the IncomingInterfaceName field if non-nil, zero value otherwise.

### GetIncomingInterfaceNameOk

`func (o *InlineResponse20035Routes) GetIncomingInterfaceNameOk() (*string, bool)`

GetIncomingInterfaceNameOk returns a tuple with the IncomingInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingInterfaceName

`func (o *InlineResponse20035Routes) SetIncomingInterfaceName(v string)`

SetIncomingInterfaceName sets IncomingInterfaceName field to given value.

### HasIncomingInterfaceName

`func (o *InlineResponse20035Routes) HasIncomingInterfaceName() bool`

HasIncomingInterfaceName returns a boolean if a field has been set.

### GetOutgoingInterfaceNames

`func (o *InlineResponse20035Routes) GetOutgoingInterfaceNames() []string`

GetOutgoingInterfaceNames returns the OutgoingInterfaceNames field if non-nil, zero value otherwise.

### GetOutgoingInterfaceNamesOk

`func (o *InlineResponse20035Routes) GetOutgoingInterfaceNamesOk() (*[]string, bool)`

GetOutgoingInterfaceNamesOk returns a tuple with the OutgoingInterfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingInterfaceNames

`func (o *InlineResponse20035Routes) SetOutgoingInterfaceNames(v []string)`

SetOutgoingInterfaceNames sets OutgoingInterfaceNames field to given value.

### HasOutgoingInterfaceNames

`func (o *InlineResponse20035Routes) HasOutgoingInterfaceNames() bool`

HasOutgoingInterfaceNames returns a boolean if a field has been set.

### GetFlags

`func (o *InlineResponse20035Routes) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *InlineResponse20035Routes) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *InlineResponse20035Routes) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *InlineResponse20035Routes) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


