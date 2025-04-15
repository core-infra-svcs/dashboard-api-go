# InlineResponse200338

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time of the channel utilization interval. | [optional] 
**EndTs** | Pointer to **time.Time** | The end time of the channel utilization interval. | [optional] 
**Serial** | Pointer to **string** | The serial number for the device. | [optional] 
**Mac** | Pointer to **string** | The MAC address of the device. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork**](OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork.md) |  | [optional] 
**ByBand** | Pointer to [**[]OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand**](OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand.md) | Channel utilization broken down by band. | [optional] 

## Methods

### NewInlineResponse200338

`func NewInlineResponse200338() *InlineResponse200338`

NewInlineResponse200338 instantiates a new InlineResponse200338 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200338WithDefaults

`func NewInlineResponse200338WithDefaults() *InlineResponse200338`

NewInlineResponse200338WithDefaults instantiates a new InlineResponse200338 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200338) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200338) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200338) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200338) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200338) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200338) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200338) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200338) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200338) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200338) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200338) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200338) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200338) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200338) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200338) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200338) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200338) GetNetwork() OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200338) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200338) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200338) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetByBand

`func (o *InlineResponse200338) GetByBand() []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand`

GetByBand returns the ByBand field if non-nil, zero value otherwise.

### GetByBandOk

`func (o *InlineResponse200338) GetByBandOk() (*[]OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand, bool)`

GetByBandOk returns a tuple with the ByBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByBand

`func (o *InlineResponse200338) SetByBand(v []OrganizationsOrganizationIdWirelessDevicesChannelUtilizationByDeviceByBand)`

SetByBand sets ByBand field to given value.

### HasByBand

`func (o *InlineResponse200338) HasByBand() bool`

HasByBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


