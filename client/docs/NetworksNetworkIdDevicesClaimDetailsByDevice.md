# NetworksNetworkIdDevicesClaimDetailsByDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | **string** | The serial of the device these details relate to | 
**Details** | [**[]NetworksNetworkIdDevicesClaimDetails**](NetworksNetworkIdDevicesClaimDetails.md) | An array of details. Supported list         of details includes: \&quot;device mode\&quot;, \&quot;username\&quot;, \&quot;password\&quot;, \&quot;enable password\&quot;, \&quot;ap mapping type\&quot; and         \&quot;ap network id\&quot;. For onboarding into hybrid mode, the value of the device mode detail must be \&quot;monitored\&quot; | 

## Methods

### NewNetworksNetworkIdDevicesClaimDetailsByDevice

`func NewNetworksNetworkIdDevicesClaimDetailsByDevice(serial string, details []NetworksNetworkIdDevicesClaimDetails, ) *NetworksNetworkIdDevicesClaimDetailsByDevice`

NewNetworksNetworkIdDevicesClaimDetailsByDevice instantiates a new NetworksNetworkIdDevicesClaimDetailsByDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdDevicesClaimDetailsByDeviceWithDefaults

`func NewNetworksNetworkIdDevicesClaimDetailsByDeviceWithDefaults() *NetworksNetworkIdDevicesClaimDetailsByDevice`

NewNetworksNetworkIdDevicesClaimDetailsByDeviceWithDefaults instantiates a new NetworksNetworkIdDevicesClaimDetailsByDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *NetworksNetworkIdDevicesClaimDetailsByDevice) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NetworksNetworkIdDevicesClaimDetailsByDevice) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NetworksNetworkIdDevicesClaimDetailsByDevice) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetDetails

`func (o *NetworksNetworkIdDevicesClaimDetailsByDevice) GetDetails() []NetworksNetworkIdDevicesClaimDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NetworksNetworkIdDevicesClaimDetailsByDevice) GetDetailsOk() (*[]NetworksNetworkIdDevicesClaimDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NetworksNetworkIdDevicesClaimDetailsByDevice) SetDetails(v []NetworksNetworkIdDevicesClaimDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


