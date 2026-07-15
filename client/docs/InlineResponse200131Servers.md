# InlineResponse200131Servers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The IP address or FQDN of the syslog server | [optional] 
**Port** | Pointer to **int32** | The port of the syslog server | [optional] 
**Roles** | Pointer to **[]string** | A list of roles for the syslog server, specific to Product Type. Options can be found by querying GET api/v1/organizations/:id/devices/syslog/servers/roles/byNetwork and using the value property. Maximum array size is limited by the number of unique roles, which is about 10 | [optional] 
**TransportProtocol** | Pointer to **string** | The transport protocol for the syslog server | [optional] 
**Encryption** | Pointer to [**InlineResponse200131Encryption**](InlineResponse200131Encryption.md) |  | [optional] 

## Methods

### NewInlineResponse200131Servers

`func NewInlineResponse200131Servers() *InlineResponse200131Servers`

NewInlineResponse200131Servers instantiates a new InlineResponse200131Servers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200131ServersWithDefaults

`func NewInlineResponse200131ServersWithDefaults() *InlineResponse200131Servers`

NewInlineResponse200131ServersWithDefaults instantiates a new InlineResponse200131Servers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *InlineResponse200131Servers) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineResponse200131Servers) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineResponse200131Servers) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InlineResponse200131Servers) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *InlineResponse200131Servers) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse200131Servers) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse200131Servers) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse200131Servers) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRoles

`func (o *InlineResponse200131Servers) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *InlineResponse200131Servers) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *InlineResponse200131Servers) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *InlineResponse200131Servers) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTransportProtocol

`func (o *InlineResponse200131Servers) GetTransportProtocol() string`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *InlineResponse200131Servers) GetTransportProtocolOk() (*string, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *InlineResponse200131Servers) SetTransportProtocol(v string)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *InlineResponse200131Servers) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetEncryption

`func (o *InlineResponse200131Servers) GetEncryption() InlineResponse200131Encryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *InlineResponse200131Servers) GetEncryptionOk() (*InlineResponse200131Encryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *InlineResponse200131Servers) SetEncryption(v InlineResponse200131Encryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *InlineResponse200131Servers) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


