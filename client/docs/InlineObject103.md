# InlineObject103

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the MQTT broker. | 
**Host** | **string** | Host name/IP address where the MQTT broker runs. | 
**Port** | **int32** | Host port though which the MQTT broker can be reached. | 
**Security** | Pointer to [**NetworksNetworkIdMqttBrokersSecurity**](NetworksNetworkIdMqttBrokersSecurity.md) |  | [optional] 
**Authentication** | Pointer to [**NetworksNetworkIdMqttBrokersAuthentication**](NetworksNetworkIdMqttBrokersAuthentication.md) |  | [optional] 

## Methods

### NewInlineObject103

`func NewInlineObject103(name string, host string, port int32, ) *InlineObject103`

NewInlineObject103 instantiates a new InlineObject103 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject103WithDefaults

`func NewInlineObject103WithDefaults() *InlineObject103`

NewInlineObject103WithDefaults instantiates a new InlineObject103 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject103) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject103) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject103) SetName(v string)`

SetName sets Name field to given value.


### GetHost

`func (o *InlineObject103) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineObject103) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineObject103) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *InlineObject103) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineObject103) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineObject103) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSecurity

`func (o *InlineObject103) GetSecurity() NetworksNetworkIdMqttBrokersSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *InlineObject103) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *InlineObject103) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *InlineObject103) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineObject103) GetAuthentication() NetworksNetworkIdMqttBrokersAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineObject103) GetAuthenticationOk() (*NetworksNetworkIdMqttBrokersAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineObject103) SetAuthentication(v NetworksNetworkIdMqttBrokersAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineObject103) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


