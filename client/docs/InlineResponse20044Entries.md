# InlineResponse20044Entries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of route defined | [optional] 
**Subnet** | Pointer to **string** | The subnet of the route in CIDR format | [optional] 
**NextHops** | Pointer to [**[]InlineResponse20044NextHops**](InlineResponse20044NextHops.md) | Ordered list of ECMP(Equal cost multiple path) next hops for the destination subnet | [optional] 
**IpVersion** | Pointer to **string** | The IP version of the route&#39; | [optional] 
**Vrf** | Pointer to [**InlineResponse20044Vrf**](InlineResponse20044Vrf.md) |  | [optional] 

## Methods

### NewInlineResponse20044Entries

`func NewInlineResponse20044Entries() *InlineResponse20044Entries`

NewInlineResponse20044Entries instantiates a new InlineResponse20044Entries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20044EntriesWithDefaults

`func NewInlineResponse20044EntriesWithDefaults() *InlineResponse20044Entries`

NewInlineResponse20044EntriesWithDefaults instantiates a new InlineResponse20044Entries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse20044Entries) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20044Entries) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20044Entries) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20044Entries) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20044Entries) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20044Entries) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20044Entries) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20044Entries) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetNextHops

`func (o *InlineResponse20044Entries) GetNextHops() []InlineResponse20044NextHops`

GetNextHops returns the NextHops field if non-nil, zero value otherwise.

### GetNextHopsOk

`func (o *InlineResponse20044Entries) GetNextHopsOk() (*[]InlineResponse20044NextHops, bool)`

GetNextHopsOk returns a tuple with the NextHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHops

`func (o *InlineResponse20044Entries) SetNextHops(v []InlineResponse20044NextHops)`

SetNextHops sets NextHops field to given value.

### HasNextHops

`func (o *InlineResponse20044Entries) HasNextHops() bool`

HasNextHops returns a boolean if a field has been set.

### GetIpVersion

`func (o *InlineResponse20044Entries) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *InlineResponse20044Entries) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *InlineResponse20044Entries) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *InlineResponse20044Entries) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse20044Entries) GetVrf() InlineResponse20044Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse20044Entries) GetVrfOk() (*InlineResponse20044Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse20044Entries) SetVrf(v InlineResponse20044Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse20044Entries) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


