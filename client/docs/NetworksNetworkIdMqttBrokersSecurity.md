# NetworksNetworkIdMqttBrokersSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Security protocol of the MQTT broker. | [optional] 
**Tls** | Pointer to [**NetworksNetworkIdMqttBrokersSecurityTls**](NetworksNetworkIdMqttBrokersSecurityTls.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdMqttBrokersSecurity

`func NewNetworksNetworkIdMqttBrokersSecurity() *NetworksNetworkIdMqttBrokersSecurity`

NewNetworksNetworkIdMqttBrokersSecurity instantiates a new NetworksNetworkIdMqttBrokersSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdMqttBrokersSecurityWithDefaults

`func NewNetworksNetworkIdMqttBrokersSecurityWithDefaults() *NetworksNetworkIdMqttBrokersSecurity`

NewNetworksNetworkIdMqttBrokersSecurityWithDefaults instantiates a new NetworksNetworkIdMqttBrokersSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *NetworksNetworkIdMqttBrokersSecurity) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NetworksNetworkIdMqttBrokersSecurity) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NetworksNetworkIdMqttBrokersSecurity) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *NetworksNetworkIdMqttBrokersSecurity) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTls

`func (o *NetworksNetworkIdMqttBrokersSecurity) GetTls() NetworksNetworkIdMqttBrokersSecurityTls`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *NetworksNetworkIdMqttBrokersSecurity) GetTlsOk() (*NetworksNetworkIdMqttBrokersSecurityTls, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *NetworksNetworkIdMqttBrokersSecurity) SetTls(v NetworksNetworkIdMqttBrokersSecurityTls)`

SetTls sets Tls field to given value.

### HasTls

`func (o *NetworksNetworkIdMqttBrokersSecurity) HasTls() bool`

HasTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


