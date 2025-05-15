# InlineResponse200110

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

### NewInlineResponse200110

`func NewInlineResponse200110() *InlineResponse200110`

NewInlineResponse200110 instantiates a new InlineResponse200110 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200110WithDefaults

`func NewInlineResponse200110WithDefaults() *InlineResponse200110`

NewInlineResponse200110WithDefaults instantiates a new InlineResponse200110 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200110) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200110) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200110) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200110) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200110) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200110) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200110) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200110) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHost

`func (o *InlineResponse200110) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineResponse200110) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineResponse200110) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InlineResponse200110) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *InlineResponse200110) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse200110) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse200110) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse200110) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurity

`func (o *InlineResponse200110) GetSecurity() NetworksNetworkIdMqttBrokersSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *InlineResponse200110) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *InlineResponse200110) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *InlineResponse200110) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineResponse200110) GetAuthentication() NetworksNetworkIdMqttBrokersAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineResponse200110) GetAuthenticationOk() (*NetworksNetworkIdMqttBrokersAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineResponse200110) SetAuthentication(v NetworksNetworkIdMqttBrokersAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineResponse200110) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


