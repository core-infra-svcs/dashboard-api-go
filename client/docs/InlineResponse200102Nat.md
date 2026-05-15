# InlineResponse200102Nat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not VPN subnet translation is enabled for the subnet | [optional] 
**RemoteSubnet** | Pointer to **string** | The translated subnet to be used in the VPN | [optional] 

## Methods

### NewInlineResponse200102Nat

`func NewInlineResponse200102Nat() *InlineResponse200102Nat`

NewInlineResponse200102Nat instantiates a new InlineResponse200102Nat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200102NatWithDefaults

`func NewInlineResponse200102NatWithDefaults() *InlineResponse200102Nat`

NewInlineResponse200102NatWithDefaults instantiates a new InlineResponse200102Nat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200102Nat) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200102Nat) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200102Nat) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200102Nat) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *InlineResponse200102Nat) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *InlineResponse200102Nat) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *InlineResponse200102Nat) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *InlineResponse200102Nat) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


