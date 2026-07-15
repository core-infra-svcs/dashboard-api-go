# NetworksNetworkIdDevicesSyslogServersServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | The IP address or FQDN of the syslog server | 
**Port** | **int32** | The port of the syslog server | 
**Roles** | **[]string** | A list of roles for the syslog server, specific to Product Type. Options can be found by querying GET api/v1/organizations/:id/devices/syslog/servers/roles/byNetwork and using the value property. Maximum array size is limited by the number of unique roles, which is about 10 | 
**TransportProtocol** | Pointer to **string** | The transport protocol for the syslog server, defaults to UDP | [optional] 
**Encryption** | Pointer to [**NetworksNetworkIdDevicesSyslogServersEncryption**](NetworksNetworkIdDevicesSyslogServersEncryption.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdDevicesSyslogServersServers

`func NewNetworksNetworkIdDevicesSyslogServersServers(host string, port int32, roles []string, ) *NetworksNetworkIdDevicesSyslogServersServers`

NewNetworksNetworkIdDevicesSyslogServersServers instantiates a new NetworksNetworkIdDevicesSyslogServersServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdDevicesSyslogServersServersWithDefaults

`func NewNetworksNetworkIdDevicesSyslogServersServersWithDefaults() *NetworksNetworkIdDevicesSyslogServersServers`

NewNetworksNetworkIdDevicesSyslogServersServersWithDefaults instantiates a new NetworksNetworkIdDevicesSyslogServersServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *NetworksNetworkIdDevicesSyslogServersServers) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworksNetworkIdDevicesSyslogServersServers) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworksNetworkIdDevicesSyslogServersServers) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *NetworksNetworkIdDevicesSyslogServersServers) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworksNetworkIdDevicesSyslogServersServers) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworksNetworkIdDevicesSyslogServersServers) SetPort(v int32)`

SetPort sets Port field to given value.


### GetRoles

`func (o *NetworksNetworkIdDevicesSyslogServersServers) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *NetworksNetworkIdDevicesSyslogServersServers) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *NetworksNetworkIdDevicesSyslogServersServers) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetTransportProtocol

`func (o *NetworksNetworkIdDevicesSyslogServersServers) GetTransportProtocol() string`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *NetworksNetworkIdDevicesSyslogServersServers) GetTransportProtocolOk() (*string, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *NetworksNetworkIdDevicesSyslogServersServers) SetTransportProtocol(v string)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *NetworksNetworkIdDevicesSyslogServersServers) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetEncryption

`func (o *NetworksNetworkIdDevicesSyslogServersServers) GetEncryption() NetworksNetworkIdDevicesSyslogServersEncryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *NetworksNetworkIdDevicesSyslogServersServers) GetEncryptionOk() (*NetworksNetworkIdDevicesSyslogServersEncryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *NetworksNetworkIdDevicesSyslogServersServers) SetEncryption(v NetworksNetworkIdDevicesSyslogServersEncryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *NetworksNetworkIdDevicesSyslogServersServers) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


