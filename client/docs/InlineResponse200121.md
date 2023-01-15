# InlineResponse200121

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number of the AP | [optional] 
**Name** | Pointer to **string** | The name of the AP | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork**](OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork.md) |  | [optional] 
**Power** | Pointer to [**OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower**](OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower.md) |  | [optional] 
**Ports** | Pointer to [**[]OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts**](OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts.md) | List of port details | [optional] 
**Aggregation** | Pointer to [**OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation**](OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation.md) |  | [optional] 

## Methods

### NewInlineResponse200121

`func NewInlineResponse200121() *InlineResponse200121`

NewInlineResponse200121 instantiates a new InlineResponse200121 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200121WithDefaults

`func NewInlineResponse200121WithDefaults() *InlineResponse200121`

NewInlineResponse200121WithDefaults instantiates a new InlineResponse200121 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200121) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200121) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200121) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200121) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200121) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200121) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200121) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200121) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200121) GetNetwork() OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200121) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200121) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesEthernetStatusesNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200121) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPower

`func (o *InlineResponse200121) GetPower() OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *InlineResponse200121) GetPowerOk() (*OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *InlineResponse200121) SetPower(v OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower)`

SetPower sets Power field to given value.

### HasPower

`func (o *InlineResponse200121) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200121) GetPorts() []OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200121) GetPortsOk() (*[]OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200121) SetPorts(v []OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200121) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetAggregation

`func (o *InlineResponse200121) GetAggregation() OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *InlineResponse200121) GetAggregationOk() (*OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *InlineResponse200121) SetAggregation(v OrganizationsOrganizationIdWirelessDevicesEthernetStatusesAggregation)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *InlineResponse200121) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


