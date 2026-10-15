# NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** | Unique ID of an existing Access Policy accounting server to copy | [optional] 
**OrganizationRadiusServerId** | Pointer to **string** | Organization wide RADIUS server ID. Omit host, port and secret when this field is provided. If supplied, host and port must match the organization-wide server; secret is ignored | [optional] 
**Host** | Pointer to **string** | Public IP address or fully-qualified domain name of the RADIUS accounting server | [optional] 
**Port** | Pointer to **int32** | Port that the RADIUS Accounting server listens on for access requests | [optional] 
**Secret** | Pointer to **string** | RADIUS client shared secret | [optional] 
**Radsec** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesRadsec3**](NetworksNetworkIdSwitchAccessPoliciesRadsec3.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1

`func NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1() *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1`

NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1WithDefaults

`func NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1WithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1`

NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1WithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetOrganizationRadiusServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetOrganizationRadiusServerId() string`

GetOrganizationRadiusServerId returns the OrganizationRadiusServerId field if non-nil, zero value otherwise.

### GetOrganizationRadiusServerIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetOrganizationRadiusServerIdOk() (*string, bool)`

GetOrganizationRadiusServerIdOk returns a tuple with the OrganizationRadiusServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRadiusServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetOrganizationRadiusServerId(v string)`

SetOrganizationRadiusServerId sets OrganizationRadiusServerId field to given value.

### HasOrganizationRadiusServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) HasOrganizationRadiusServerId() bool`

HasOrganizationRadiusServerId returns a boolean if a field has been set.

### GetHost

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRadsec

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetRadsec() NetworksNetworkIdSwitchAccessPoliciesRadsec3`

GetRadsec returns the Radsec field if non-nil, zero value otherwise.

### GetRadsecOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetRadsecOk() (*NetworksNetworkIdSwitchAccessPoliciesRadsec3, bool)`

GetRadsecOk returns a tuple with the Radsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadsec

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetRadsec(v NetworksNetworkIdSwitchAccessPoliciesRadsec3)`

SetRadsec sets Radsec field to given value.

### HasRadsec

`func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) HasRadsec() bool`

HasRadsec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


