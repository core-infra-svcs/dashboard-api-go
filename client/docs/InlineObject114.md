# InlineObject114

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the MQTT broker. | [optional] 
**Host** | Pointer to **string** | Host name/IP address where the MQTT broker runs. | [optional] 
**Port** | Pointer to **int32** | Host port though which the MQTT broker can be reached. | [optional] 
**Security** | Pointer to [**NetworksNetworkIdMqttBrokersSecurity1**](NetworksNetworkIdMqttBrokersSecurity1.md) |  | [optional] 
**Authentication** | Pointer to [**NetworksNetworkIdMqttBrokersAuthentication1**](NetworksNetworkIdMqttBrokersAuthentication1.md) |  | [optional] 

## Methods

### NewInlineObject114

`func NewInlineObject114() *InlineObject114`

NewInlineObject114 instantiates a new InlineObject114 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject114WithDefaults

`func NewInlineObject114WithDefaults() *InlineObject114`

NewInlineObject114WithDefaults instantiates a new InlineObject114 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject114) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject114) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject114) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject114) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHost

`func (o *InlineObject114) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineObject114) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineObject114) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InlineObject114) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *InlineObject114) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineObject114) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineObject114) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineObject114) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurity

`func (o *InlineObject114) GetSecurity() NetworksNetworkIdMqttBrokersSecurity1`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *InlineObject114) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity1, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *InlineObject114) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity1)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *InlineObject114) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineObject114) GetAuthentication() NetworksNetworkIdMqttBrokersAuthentication1`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineObject114) GetAuthenticationOk() (*NetworksNetworkIdMqttBrokersAuthentication1, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineObject114) SetAuthentication(v NetworksNetworkIdMqttBrokersAuthentication1)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineObject114) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


