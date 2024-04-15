# NetworksNetworkIdApplianceSdwanInternetPoliciesValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol of the traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | [**NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource**](NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource.md) |  | 
**Destination** | [**NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination**](NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination.md) |  | 

## Methods

### NewNetworksNetworkIdApplianceSdwanInternetPoliciesValue

`func NewNetworksNetworkIdApplianceSdwanInternetPoliciesValue(source NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource, destination NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination, ) *NetworksNetworkIdApplianceSdwanInternetPoliciesValue`

NewNetworksNetworkIdApplianceSdwanInternetPoliciesValue instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueWithDefaults

`func NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueWithDefaults() *NetworksNetworkIdApplianceSdwanInternetPoliciesValue`

NewNetworksNetworkIdApplianceSdwanInternetPoliciesValueWithDefaults instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetSource() NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetSourceOk() (*NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) SetSource(v NetworksNetworkIdApplianceSdwanInternetPoliciesValueSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetDestination() NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) GetDestinationOk() (*NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesValue) SetDestination(v NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


