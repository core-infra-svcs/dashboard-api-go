# InlineResponse200111NamedVlansPoolDhcpMonitoring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP | [optional] 
**Duration** | Pointer to **int32** | The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool. | [optional] 

## Methods

### NewInlineResponse200111NamedVlansPoolDhcpMonitoring

`func NewInlineResponse200111NamedVlansPoolDhcpMonitoring() *InlineResponse200111NamedVlansPoolDhcpMonitoring`

NewInlineResponse200111NamedVlansPoolDhcpMonitoring instantiates a new InlineResponse200111NamedVlansPoolDhcpMonitoring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200111NamedVlansPoolDhcpMonitoringWithDefaults

`func NewInlineResponse200111NamedVlansPoolDhcpMonitoringWithDefaults() *InlineResponse200111NamedVlansPoolDhcpMonitoring`

NewInlineResponse200111NamedVlansPoolDhcpMonitoringWithDefaults instantiates a new InlineResponse200111NamedVlansPoolDhcpMonitoring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200111NamedVlansPoolDhcpMonitoring) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200111NamedVlansPoolDhcpMonitoring) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200111NamedVlansPoolDhcpMonitoring) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200111NamedVlansPoolDhcpMonitoring) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDuration

`func (o *InlineResponse200111NamedVlansPoolDhcpMonitoring) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineResponse200111NamedVlansPoolDhcpMonitoring) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineResponse200111NamedVlansPoolDhcpMonitoring) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineResponse200111NamedVlansPoolDhcpMonitoring) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

