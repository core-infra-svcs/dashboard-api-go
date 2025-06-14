# InlineObject115

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the MQTT broker. | 
**Host** | **string** | Host name/IP address where the MQTT broker runs. | 
**Port** | **int32** | Host port though which the MQTT broker can be reached. | 
**Security** | Pointer to [**NetworksNetworkIdMqttBrokersSecurity1**](NetworksNetworkIdMqttBrokersSecurity1.md) |  | [optional] 
**Authentication** | Pointer to [**NetworksNetworkIdMqttBrokersAuthentication1**](NetworksNetworkIdMqttBrokersAuthentication1.md) |  | [optional] 

## Methods

### NewInlineObject115

`func NewInlineObject115(name string, host string, port int32, ) *InlineObject115`

NewInlineObject115 instantiates a new InlineObject115 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject115WithDefaults

`func NewInlineObject115WithDefaults() *InlineObject115`

NewInlineObject115WithDefaults instantiates a new InlineObject115 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject115) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject115) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject115) SetName(v string)`

SetName sets Name field to given value.


### GetHost

`func (o *InlineObject115) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineObject115) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineObject115) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *InlineObject115) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineObject115) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineObject115) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSecurity

`func (o *InlineObject115) GetSecurity() NetworksNetworkIdMqttBrokersSecurity1`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *InlineObject115) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity1, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *InlineObject115) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity1)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *InlineObject115) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineObject115) GetAuthentication() NetworksNetworkIdMqttBrokersAuthentication1`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineObject115) GetAuthenticationOk() (*NetworksNetworkIdMqttBrokersAuthentication1, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineObject115) SetAuthentication(v NetworksNetworkIdMqttBrokersAuthentication1)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineObject115) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


