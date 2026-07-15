# InlineResponse200192

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**InlineResponse200192Alerts**](InlineResponse200192Alerts.md) |  | [optional] 
**DefaultPolicy** | Pointer to **string** | &#39;allow&#39; or &#39;block&#39; new DHCP servers. Default value is &#39;allow&#39;. | [optional] 
**BlockedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set       to allow.An empty array will clear the entries. | [optional] 
**AllowedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set       to block.An empty array will clear the entries. | [optional] 
**AlwaysAllowedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers that are always allowed on the network. These are the MAC addresses of switches, switch stacks, and VRRP groups in the network. These servers cannot be blocked. | [optional] 
**ArpInspection** | Pointer to [**InlineResponse200192ArpInspection**](InlineResponse200192ArpInspection.md) |  | [optional] 

## Methods

### NewInlineResponse200192

`func NewInlineResponse200192() *InlineResponse200192`

NewInlineResponse200192 instantiates a new InlineResponse200192 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200192WithDefaults

`func NewInlineResponse200192WithDefaults() *InlineResponse200192`

NewInlineResponse200192WithDefaults instantiates a new InlineResponse200192 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *InlineResponse200192) GetAlerts() InlineResponse200192Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineResponse200192) GetAlertsOk() (*InlineResponse200192Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineResponse200192) SetAlerts(v InlineResponse200192Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineResponse200192) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDefaultPolicy

`func (o *InlineResponse200192) GetDefaultPolicy() string`

GetDefaultPolicy returns the DefaultPolicy field if non-nil, zero value otherwise.

### GetDefaultPolicyOk

`func (o *InlineResponse200192) GetDefaultPolicyOk() (*string, bool)`

GetDefaultPolicyOk returns a tuple with the DefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPolicy

`func (o *InlineResponse200192) SetDefaultPolicy(v string)`

SetDefaultPolicy sets DefaultPolicy field to given value.

### HasDefaultPolicy

`func (o *InlineResponse200192) HasDefaultPolicy() bool`

HasDefaultPolicy returns a boolean if a field has been set.

### GetBlockedServers

`func (o *InlineResponse200192) GetBlockedServers() []string`

GetBlockedServers returns the BlockedServers field if non-nil, zero value otherwise.

### GetBlockedServersOk

`func (o *InlineResponse200192) GetBlockedServersOk() (*[]string, bool)`

GetBlockedServersOk returns a tuple with the BlockedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedServers

`func (o *InlineResponse200192) SetBlockedServers(v []string)`

SetBlockedServers sets BlockedServers field to given value.

### HasBlockedServers

`func (o *InlineResponse200192) HasBlockedServers() bool`

HasBlockedServers returns a boolean if a field has been set.

### GetAllowedServers

`func (o *InlineResponse200192) GetAllowedServers() []string`

GetAllowedServers returns the AllowedServers field if non-nil, zero value otherwise.

### GetAllowedServersOk

`func (o *InlineResponse200192) GetAllowedServersOk() (*[]string, bool)`

GetAllowedServersOk returns a tuple with the AllowedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServers

`func (o *InlineResponse200192) SetAllowedServers(v []string)`

SetAllowedServers sets AllowedServers field to given value.

### HasAllowedServers

`func (o *InlineResponse200192) HasAllowedServers() bool`

HasAllowedServers returns a boolean if a field has been set.

### GetAlwaysAllowedServers

`func (o *InlineResponse200192) GetAlwaysAllowedServers() []string`

GetAlwaysAllowedServers returns the AlwaysAllowedServers field if non-nil, zero value otherwise.

### GetAlwaysAllowedServersOk

`func (o *InlineResponse200192) GetAlwaysAllowedServersOk() (*[]string, bool)`

GetAlwaysAllowedServersOk returns a tuple with the AlwaysAllowedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysAllowedServers

`func (o *InlineResponse200192) SetAlwaysAllowedServers(v []string)`

SetAlwaysAllowedServers sets AlwaysAllowedServers field to given value.

### HasAlwaysAllowedServers

`func (o *InlineResponse200192) HasAlwaysAllowedServers() bool`

HasAlwaysAllowedServers returns a boolean if a field has been set.

### GetArpInspection

`func (o *InlineResponse200192) GetArpInspection() InlineResponse200192ArpInspection`

GetArpInspection returns the ArpInspection field if non-nil, zero value otherwise.

### GetArpInspectionOk

`func (o *InlineResponse200192) GetArpInspectionOk() (*InlineResponse200192ArpInspection, bool)`

GetArpInspectionOk returns a tuple with the ArpInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpInspection

`func (o *InlineResponse200192) SetArpInspection(v InlineResponse200192ArpInspection)`

SetArpInspection sets ArpInspection field to given value.

### HasArpInspection

`func (o *InlineResponse200192) HasArpInspection() bool`

HasArpInspection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


