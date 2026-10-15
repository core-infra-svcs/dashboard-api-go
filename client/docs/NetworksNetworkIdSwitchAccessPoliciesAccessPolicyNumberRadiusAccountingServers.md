# NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** | Unique ID of the RADIUS accounting server to update | [optional] 
**OrganizationRadiusServerId** | Pointer to **string** | Organization wide RADIUS server ID. Omit host, port and secret when this field is provided. If supplied, host and port must match the organization-wide server; secret is ignored | [optional] 
**Host** | Pointer to **string** | Public IP address or fully-qualified domain name of the RADIUS accounting server | [optional] 
**Port** | Pointer to **int32** | Port that the RADIUS Accounting server listens on for access requests | [optional] 
**Secret** | Pointer to **string** | RADIUS client shared secret | [optional] 
**Radsec** | Pointer to [**NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadsec1**](NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadsec1.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers

`func NewNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers() *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers`

NewNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers instantiates a new NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServersWithDefaults

`func NewNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServersWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers`

NewNetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServersWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetOrganizationRadiusServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetOrganizationRadiusServerId() string`

GetOrganizationRadiusServerId returns the OrganizationRadiusServerId field if non-nil, zero value otherwise.

### GetOrganizationRadiusServerIdOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetOrganizationRadiusServerIdOk() (*string, bool)`

GetOrganizationRadiusServerIdOk returns a tuple with the OrganizationRadiusServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRadiusServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) SetOrganizationRadiusServerId(v string)`

SetOrganizationRadiusServerId sets OrganizationRadiusServerId field to given value.

### HasOrganizationRadiusServerId

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) HasOrganizationRadiusServerId() bool`

HasOrganizationRadiusServerId returns a boolean if a field has been set.

### GetHost

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRadsec

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetRadsec() NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadsec1`

GetRadsec returns the Radsec field if non-nil, zero value otherwise.

### GetRadsecOk

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) GetRadsecOk() (*NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadsec1, bool)`

GetRadsecOk returns a tuple with the Radsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadsec

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) SetRadsec(v NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadsec1)`

SetRadsec sets Radsec field to given value.

### HasRadsec

`func (o *NetworksNetworkIdSwitchAccessPoliciesAccessPolicyNumberRadiusAccountingServers) HasRadsec() bool`

HasRadsec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


