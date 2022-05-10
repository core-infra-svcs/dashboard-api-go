# NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | The IP address to test connectivity with | 
**Description** | Pointer to **string** | Description of the testing destination. Optional, defaults to null | [optional] 
**Default** | Pointer to **bool** | Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed | [optional] 

## Methods

### NewNetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations

`func NewNetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations(ip string, ) *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations`

NewNetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations instantiates a new NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinationsWithDefaults

`func NewNetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinationsWithDefaults() *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations`

NewNetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinationsWithDefaults instantiates a new NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) SetIp(v string)`

SetIp sets Ip field to given value.


### GetDescription

`func (o *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefault

`func (o *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *NetworksNetworkIdApplianceConnectivityMonitoringDestinationsDestinations) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


