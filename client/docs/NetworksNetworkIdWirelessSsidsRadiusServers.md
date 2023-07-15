# NetworksNetworkIdWirelessSsidsRadiusServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | IP address (or FQDN) of your RADIUS server | [optional] 
**Port** | Pointer to **int32** | UDP port the RADIUS server listens on for Access-requests | [optional] 
**OpenRoamingCertificateId** | Pointer to **int32** | The ID of the Openroaming Certificate attached to radius server | [optional] 
**CaCertificate** | Pointer to **string** | Certificate used for authorization for the RADSEC Server | [optional] 

## Methods

### NewNetworksNetworkIdWirelessSsidsRadiusServers

`func NewNetworksNetworkIdWirelessSsidsRadiusServers() *NetworksNetworkIdWirelessSsidsRadiusServers`

NewNetworksNetworkIdWirelessSsidsRadiusServers instantiates a new NetworksNetworkIdWirelessSsidsRadiusServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessSsidsRadiusServersWithDefaults

`func NewNetworksNetworkIdWirelessSsidsRadiusServersWithDefaults() *NetworksNetworkIdWirelessSsidsRadiusServers`

NewNetworksNetworkIdWirelessSsidsRadiusServersWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsRadiusServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetOpenRoamingCertificateId

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetOpenRoamingCertificateId() int32`

GetOpenRoamingCertificateId returns the OpenRoamingCertificateId field if non-nil, zero value otherwise.

### GetOpenRoamingCertificateIdOk

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetOpenRoamingCertificateIdOk() (*int32, bool)`

GetOpenRoamingCertificateIdOk returns a tuple with the OpenRoamingCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenRoamingCertificateId

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) SetOpenRoamingCertificateId(v int32)`

SetOpenRoamingCertificateId sets OpenRoamingCertificateId field to given value.

### HasOpenRoamingCertificateId

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) HasOpenRoamingCertificateId() bool`

HasOpenRoamingCertificateId returns a boolean if a field has been set.

### GetCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsRadiusServers) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


