# InlineResponse200114Nat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not VPN subnet translation is enabled for the subnet | [optional] 
**RemoteSubnet** | Pointer to **string** | The translated subnet to be used in the VPN | [optional] 
**Legacy** | Pointer to **map[string]interface{}** | DEPRECATED: Legacy VPN NAT settings. For VLAN entries, when the corresponding NFO is enabled, the related field is returned for every VLAN, including false values. | [optional] 

## Methods

### NewInlineResponse200114Nat

`func NewInlineResponse200114Nat() *InlineResponse200114Nat`

NewInlineResponse200114Nat instantiates a new InlineResponse200114Nat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200114NatWithDefaults

`func NewInlineResponse200114NatWithDefaults() *InlineResponse200114Nat`

NewInlineResponse200114NatWithDefaults instantiates a new InlineResponse200114Nat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200114Nat) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200114Nat) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200114Nat) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200114Nat) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *InlineResponse200114Nat) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *InlineResponse200114Nat) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *InlineResponse200114Nat) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *InlineResponse200114Nat) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetLegacy

`func (o *InlineResponse200114Nat) GetLegacy() map[string]interface{}`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *InlineResponse200114Nat) GetLegacyOk() (*map[string]interface{}, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *InlineResponse200114Nat) SetLegacy(v map[string]interface{})`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *InlineResponse200114Nat) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


