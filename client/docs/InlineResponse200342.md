# InlineResponse200342

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number for the device. | [optional] 
**Mac** | Pointer to **string** | The MAC address of the device. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork**](OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork.md) |  | [optional] 
**ByBand** | Pointer to [**[]OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand**](OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand.md) | Channel utilization broken down by band. | [optional] 

## Methods

### NewInlineResponse200342

`func NewInlineResponse200342() *InlineResponse200342`

NewInlineResponse200342 instantiates a new InlineResponse200342 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200342WithDefaults

`func NewInlineResponse200342WithDefaults() *InlineResponse200342`

NewInlineResponse200342WithDefaults instantiates a new InlineResponse200342 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200342) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200342) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200342) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200342) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200342) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200342) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200342) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200342) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200342) GetNetwork() OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200342) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200342) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200342) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetByBand

`func (o *InlineResponse200342) GetByBand() []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand`

GetByBand returns the ByBand field if non-nil, zero value otherwise.

### GetByBandOk

`func (o *InlineResponse200342) GetByBandOk() (*[]OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand, bool)`

GetByBandOk returns a tuple with the ByBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByBand

`func (o *InlineResponse200342) SetByBand(v []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand)`

SetByBand sets ByBand field to given value.

### HasByBand

`func (o *InlineResponse200342) HasByBand() bool`

HasByBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


