# InlineResponse200272

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V2cEnabled** | Pointer to **bool** | Boolean indicating whether SNMP version 2c is enabled for the organization. | [optional] 
**V2CommunityString** | Pointer to **string** | The community string for SNMP version 2c, if enabled. | [optional] 
**V3Enabled** | Pointer to **bool** | Boolean indicating whether SNMP version 3 is enabled for the organization. | [optional] 
**V3User** | Pointer to **string** | The user for SNMP version 3, if enabled. | [optional] 
**V3AuthMode** | Pointer to **string** | The SNMP version 3 authentication mode. Can be either &#39;MD5&#39; or &#39;SHA&#39;. | [optional] 
**V3PrivMode** | Pointer to **string** | The SNMP version 3 privacy mode. Can be either &#39;DES&#39; or &#39;AES128&#39;. | [optional] 
**PeerIps** | Pointer to **[]string** | The list of IPv4 addresses that are allowed to access the SNMP server. | [optional] 
**Hostname** | Pointer to **string** | The hostname of the SNMP server. | [optional] 
**Port** | Pointer to **int32** | The port of the SNMP server. | [optional] 

## Methods

### NewInlineResponse200272

`func NewInlineResponse200272() *InlineResponse200272`

NewInlineResponse200272 instantiates a new InlineResponse200272 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200272WithDefaults

`func NewInlineResponse200272WithDefaults() *InlineResponse200272`

NewInlineResponse200272WithDefaults instantiates a new InlineResponse200272 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetV2cEnabled

`func (o *InlineResponse200272) GetV2cEnabled() bool`

GetV2cEnabled returns the V2cEnabled field if non-nil, zero value otherwise.

### GetV2cEnabledOk

`func (o *InlineResponse200272) GetV2cEnabledOk() (*bool, bool)`

GetV2cEnabledOk returns a tuple with the V2cEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2cEnabled

`func (o *InlineResponse200272) SetV2cEnabled(v bool)`

SetV2cEnabled sets V2cEnabled field to given value.

### HasV2cEnabled

`func (o *InlineResponse200272) HasV2cEnabled() bool`

HasV2cEnabled returns a boolean if a field has been set.

### GetV2CommunityString

`func (o *InlineResponse200272) GetV2CommunityString() string`

GetV2CommunityString returns the V2CommunityString field if non-nil, zero value otherwise.

### GetV2CommunityStringOk

`func (o *InlineResponse200272) GetV2CommunityStringOk() (*string, bool)`

GetV2CommunityStringOk returns a tuple with the V2CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2CommunityString

`func (o *InlineResponse200272) SetV2CommunityString(v string)`

SetV2CommunityString sets V2CommunityString field to given value.

### HasV2CommunityString

`func (o *InlineResponse200272) HasV2CommunityString() bool`

HasV2CommunityString returns a boolean if a field has been set.

### GetV3Enabled

`func (o *InlineResponse200272) GetV3Enabled() bool`

GetV3Enabled returns the V3Enabled field if non-nil, zero value otherwise.

### GetV3EnabledOk

`func (o *InlineResponse200272) GetV3EnabledOk() (*bool, bool)`

GetV3EnabledOk returns a tuple with the V3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Enabled

`func (o *InlineResponse200272) SetV3Enabled(v bool)`

SetV3Enabled sets V3Enabled field to given value.

### HasV3Enabled

`func (o *InlineResponse200272) HasV3Enabled() bool`

HasV3Enabled returns a boolean if a field has been set.

### GetV3User

`func (o *InlineResponse200272) GetV3User() string`

GetV3User returns the V3User field if non-nil, zero value otherwise.

### GetV3UserOk

`func (o *InlineResponse200272) GetV3UserOk() (*string, bool)`

GetV3UserOk returns a tuple with the V3User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3User

`func (o *InlineResponse200272) SetV3User(v string)`

SetV3User sets V3User field to given value.

### HasV3User

`func (o *InlineResponse200272) HasV3User() bool`

HasV3User returns a boolean if a field has been set.

### GetV3AuthMode

`func (o *InlineResponse200272) GetV3AuthMode() string`

GetV3AuthMode returns the V3AuthMode field if non-nil, zero value otherwise.

### GetV3AuthModeOk

`func (o *InlineResponse200272) GetV3AuthModeOk() (*string, bool)`

GetV3AuthModeOk returns a tuple with the V3AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3AuthMode

`func (o *InlineResponse200272) SetV3AuthMode(v string)`

SetV3AuthMode sets V3AuthMode field to given value.

### HasV3AuthMode

`func (o *InlineResponse200272) HasV3AuthMode() bool`

HasV3AuthMode returns a boolean if a field has been set.

### GetV3PrivMode

`func (o *InlineResponse200272) GetV3PrivMode() string`

GetV3PrivMode returns the V3PrivMode field if non-nil, zero value otherwise.

### GetV3PrivModeOk

`func (o *InlineResponse200272) GetV3PrivModeOk() (*string, bool)`

GetV3PrivModeOk returns a tuple with the V3PrivMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3PrivMode

`func (o *InlineResponse200272) SetV3PrivMode(v string)`

SetV3PrivMode sets V3PrivMode field to given value.

### HasV3PrivMode

`func (o *InlineResponse200272) HasV3PrivMode() bool`

HasV3PrivMode returns a boolean if a field has been set.

### GetPeerIps

`func (o *InlineResponse200272) GetPeerIps() []string`

GetPeerIps returns the PeerIps field if non-nil, zero value otherwise.

### GetPeerIpsOk

`func (o *InlineResponse200272) GetPeerIpsOk() (*[]string, bool)`

GetPeerIpsOk returns a tuple with the PeerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIps

`func (o *InlineResponse200272) SetPeerIps(v []string)`

SetPeerIps sets PeerIps field to given value.

### HasPeerIps

`func (o *InlineResponse200272) HasPeerIps() bool`

HasPeerIps returns a boolean if a field has been set.

### GetHostname

`func (o *InlineResponse200272) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InlineResponse200272) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InlineResponse200272) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InlineResponse200272) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *InlineResponse200272) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse200272) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse200272) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse200272) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


