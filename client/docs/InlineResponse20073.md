# InlineResponse20073

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**InlineResponse20073Alerts**](InlineResponse20073Alerts.md) |  | [optional] 
**DefaultPolicy** | Pointer to **string** | &#39;allow&#39; or &#39;block&#39; new DHCP servers. Default value is &#39;allow&#39;. | [optional] 
**BlockedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set       to allow.An empty array will clear the entries. | [optional] 
**AllowedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set       to block.An empty array will clear the entries. | [optional] 
**ArpInspection** | Pointer to [**InlineResponse20073ArpInspection**](InlineResponse20073ArpInspection.md) |  | [optional] 

## Methods

### NewInlineResponse20073

`func NewInlineResponse20073() *InlineResponse20073`

NewInlineResponse20073 instantiates a new InlineResponse20073 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20073WithDefaults

`func NewInlineResponse20073WithDefaults() *InlineResponse20073`

NewInlineResponse20073WithDefaults instantiates a new InlineResponse20073 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *InlineResponse20073) GetAlerts() InlineResponse20073Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineResponse20073) GetAlertsOk() (*InlineResponse20073Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineResponse20073) SetAlerts(v InlineResponse20073Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineResponse20073) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDefaultPolicy

`func (o *InlineResponse20073) GetDefaultPolicy() string`

GetDefaultPolicy returns the DefaultPolicy field if non-nil, zero value otherwise.

### GetDefaultPolicyOk

`func (o *InlineResponse20073) GetDefaultPolicyOk() (*string, bool)`

GetDefaultPolicyOk returns a tuple with the DefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPolicy

`func (o *InlineResponse20073) SetDefaultPolicy(v string)`

SetDefaultPolicy sets DefaultPolicy field to given value.

### HasDefaultPolicy

`func (o *InlineResponse20073) HasDefaultPolicy() bool`

HasDefaultPolicy returns a boolean if a field has been set.

### GetBlockedServers

`func (o *InlineResponse20073) GetBlockedServers() []string`

GetBlockedServers returns the BlockedServers field if non-nil, zero value otherwise.

### GetBlockedServersOk

`func (o *InlineResponse20073) GetBlockedServersOk() (*[]string, bool)`

GetBlockedServersOk returns a tuple with the BlockedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedServers

`func (o *InlineResponse20073) SetBlockedServers(v []string)`

SetBlockedServers sets BlockedServers field to given value.

### HasBlockedServers

`func (o *InlineResponse20073) HasBlockedServers() bool`

HasBlockedServers returns a boolean if a field has been set.

### GetAllowedServers

`func (o *InlineResponse20073) GetAllowedServers() []string`

GetAllowedServers returns the AllowedServers field if non-nil, zero value otherwise.

### GetAllowedServersOk

`func (o *InlineResponse20073) GetAllowedServersOk() (*[]string, bool)`

GetAllowedServersOk returns a tuple with the AllowedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServers

`func (o *InlineResponse20073) SetAllowedServers(v []string)`

SetAllowedServers sets AllowedServers field to given value.

### HasAllowedServers

`func (o *InlineResponse20073) HasAllowedServers() bool`

HasAllowedServers returns a boolean if a field has been set.

### GetArpInspection

`func (o *InlineResponse20073) GetArpInspection() InlineResponse20073ArpInspection`

GetArpInspection returns the ArpInspection field if non-nil, zero value otherwise.

### GetArpInspectionOk

`func (o *InlineResponse20073) GetArpInspectionOk() (*InlineResponse20073ArpInspection, bool)`

GetArpInspectionOk returns a tuple with the ArpInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpInspection

`func (o *InlineResponse20073) SetArpInspection(v InlineResponse20073ArpInspection)`

SetArpInspection sets ArpInspection field to given value.

### HasArpInspection

`func (o *InlineResponse20073) HasArpInspection() bool`

HasArpInspection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


