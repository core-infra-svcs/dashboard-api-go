# InlineResponse200108

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the MQTT Broker. | [optional] 
**Name** | Pointer to **string** | Name of the MQTT Broker. | [optional] 
**Host** | Pointer to **string** | Host name/IP address where the MQTT broker runs. | [optional] 
**Port** | Pointer to **int32** | Host port though which the MQTT broker can be reached. | [optional] 
**Security** | Pointer to [**NetworksNetworkIdMqttBrokersSecurity**](NetworksNetworkIdMqttBrokersSecurity.md) |  | [optional] 
**Authentication** | Pointer to [**NetworksNetworkIdMqttBrokersAuthentication**](NetworksNetworkIdMqttBrokersAuthentication.md) |  | [optional] 

## Methods

### NewInlineResponse200108

`func NewInlineResponse200108() *InlineResponse200108`

NewInlineResponse200108 instantiates a new InlineResponse200108 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200108WithDefaults

`func NewInlineResponse200108WithDefaults() *InlineResponse200108`

NewInlineResponse200108WithDefaults instantiates a new InlineResponse200108 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200108) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200108) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200108) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200108) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200108) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200108) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200108) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200108) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHost

`func (o *InlineResponse200108) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineResponse200108) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineResponse200108) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InlineResponse200108) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *InlineResponse200108) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse200108) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse200108) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse200108) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurity

`func (o *InlineResponse200108) GetSecurity() NetworksNetworkIdMqttBrokersSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *InlineResponse200108) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *InlineResponse200108) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *InlineResponse200108) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineResponse200108) GetAuthentication() NetworksNetworkIdMqttBrokersAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineResponse200108) GetAuthenticationOk() (*NetworksNetworkIdMqttBrokersAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineResponse200108) SetAuthentication(v NetworksNetworkIdMqttBrokersAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineResponse200108) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


