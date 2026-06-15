# InlineResponse20042Entries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of route defined | [optional] 
**Subnet** | Pointer to **string** | The subnet of the route in CIDR format | [optional] 
**NextHops** | Pointer to [**[]InlineResponse20042NextHops**](InlineResponse20042NextHops.md) | Ordered list of ECMP(Equal cost multiple path) next hops for the destination subnet | [optional] 
**IpVersion** | Pointer to **string** | The IP version of the route&#39; | [optional] 
**Vrf** | Pointer to [**InlineResponse20042Vrf**](InlineResponse20042Vrf.md) |  | [optional] 

## Methods

### NewInlineResponse20042Entries

`func NewInlineResponse20042Entries() *InlineResponse20042Entries`

NewInlineResponse20042Entries instantiates a new InlineResponse20042Entries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20042EntriesWithDefaults

`func NewInlineResponse20042EntriesWithDefaults() *InlineResponse20042Entries`

NewInlineResponse20042EntriesWithDefaults instantiates a new InlineResponse20042Entries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse20042Entries) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20042Entries) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20042Entries) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20042Entries) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20042Entries) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20042Entries) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20042Entries) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20042Entries) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetNextHops

`func (o *InlineResponse20042Entries) GetNextHops() []InlineResponse20042NextHops`

GetNextHops returns the NextHops field if non-nil, zero value otherwise.

### GetNextHopsOk

`func (o *InlineResponse20042Entries) GetNextHopsOk() (*[]InlineResponse20042NextHops, bool)`

GetNextHopsOk returns a tuple with the NextHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHops

`func (o *InlineResponse20042Entries) SetNextHops(v []InlineResponse20042NextHops)`

SetNextHops sets NextHops field to given value.

### HasNextHops

`func (o *InlineResponse20042Entries) HasNextHops() bool`

HasNextHops returns a boolean if a field has been set.

### GetIpVersion

`func (o *InlineResponse20042Entries) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineResponse20042Entries) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineResponse20042Entries) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *InlineResponse20042Entries) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse20042Entries) GetVrf() InlineResponse20042Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse20042Entries) GetVrfOk() (*InlineResponse20042Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse20042Entries) SetVrf(v InlineResponse20042Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse20042Entries) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


