# InlineResponse200155

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**InlineResponse200155Alerts**](InlineResponse200155Alerts.md) |  | [optional] 
**DefaultPolicy** | Pointer to **string** | &#39;allow&#39; or &#39;block&#39; new DHCP servers. Default value is &#39;allow&#39;. | [optional] 
**BlockedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set       to allow.An empty array will clear the entries. | [optional] 
**AllowedServers** | Pointer to **[]string** | List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set       to block.An empty array will clear the entries. | [optional] 
**ArpInspection** | Pointer to [**InlineResponse200155ArpInspection**](InlineResponse200155ArpInspection.md) |  | [optional] 

## Methods

### NewInlineResponse200155

`func NewInlineResponse200155() *InlineResponse200155`

NewInlineResponse200155 instantiates a new InlineResponse200155 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200155WithDefaults

`func NewInlineResponse200155WithDefaults() *InlineResponse200155`

NewInlineResponse200155WithDefaults instantiates a new InlineResponse200155 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *InlineResponse200155) GetAlerts() InlineResponse200155Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineResponse200155) GetAlertsOk() (*InlineResponse200155Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineResponse200155) SetAlerts(v InlineResponse200155Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineResponse200155) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDefaultPolicy

`func (o *InlineResponse200155) GetDefaultPolicy() string`

GetDefaultPolicy returns the DefaultPolicy field if non-nil, zero value otherwise.

### GetDefaultPolicyOk

`func (o *InlineResponse200155) GetDefaultPolicyOk() (*string, bool)`

GetDefaultPolicyOk returns a tuple with the DefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPolicy

`func (o *InlineResponse200155) SetDefaultPolicy(v string)`

SetDefaultPolicy sets DefaultPolicy field to given value.

### HasDefaultPolicy

`func (o *InlineResponse200155) HasDefaultPolicy() bool`

HasDefaultPolicy returns a boolean if a field has been set.

### GetBlockedServers

`func (o *InlineResponse200155) GetBlockedServers() []string`

GetBlockedServers returns the BlockedServers field if non-nil, zero value otherwise.

### GetBlockedServersOk

`func (o *InlineResponse200155) GetBlockedServersOk() (*[]string, bool)`

GetBlockedServersOk returns a tuple with the BlockedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedServers

`func (o *InlineResponse200155) SetBlockedServers(v []string)`

SetBlockedServers sets BlockedServers field to given value.

### HasBlockedServers

`func (o *InlineResponse200155) HasBlockedServers() bool`

HasBlockedServers returns a boolean if a field has been set.

### GetAllowedServers

`func (o *InlineResponse200155) GetAllowedServers() []string`

GetAllowedServers returns the AllowedServers field if non-nil, zero value otherwise.

### GetAllowedServersOk

`func (o *InlineResponse200155) GetAllowedServersOk() (*[]string, bool)`

GetAllowedServersOk returns a tuple with the AllowedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServers

`func (o *InlineResponse200155) SetAllowedServers(v []string)`

SetAllowedServers sets AllowedServers field to given value.

### HasAllowedServers

`func (o *InlineResponse200155) HasAllowedServers() bool`

HasAllowedServers returns a boolean if a field has been set.

### GetArpInspection

`func (o *InlineResponse200155) GetArpInspection() InlineResponse200155ArpInspection`

GetArpInspection returns the ArpInspection field if non-nil, zero value otherwise.

### GetArpInspectionOk

`func (o *InlineResponse200155) GetArpInspectionOk() (*InlineResponse200155ArpInspection, bool)`

GetArpInspectionOk returns a tuple with the ArpInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpInspection

`func (o *InlineResponse200155) SetArpInspection(v InlineResponse200155ArpInspection)`

SetArpInspection sets ArpInspection field to given value.

### HasArpInspection

`func (o *InlineResponse200155) HasArpInspection() bool`

HasArpInspection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


