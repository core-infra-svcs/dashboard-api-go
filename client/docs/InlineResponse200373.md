# InlineResponse200373

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**OrganizationsOrganizationIdWirelessZigbeeDevicesNetwork**](OrganizationsOrganizationIdWirelessZigbeeDevicesNetwork.md) |  | [optional] 
**PanId** | Pointer to **string** | Unique id for the zigbee device node | [optional] 
**Channel** | Pointer to **string** | Channel the zigbee node will communicate on either auto or between 11-25 | [optional] 
**TransmitPowerLevel** | Pointer to **int32** | Power level for the zigbee devices will be a int between 10-20 | [optional] 
**Enrolled** | Pointer to **bool** | Whether this zigbee node is enrolled or not | [optional] 
**Status** | Pointer to **string** | The node status, either online or offline | [optional] 
**Gateway** | Pointer to [**OrganizationsOrganizationIdWirelessZigbeeDevicesGateway**](OrganizationsOrganizationIdWirelessZigbeeDevicesGateway.md) |  | [optional] 
**Counts** | Pointer to [**OrganizationsOrganizationIdWirelessZigbeeDevicesCounts**](OrganizationsOrganizationIdWirelessZigbeeDevicesCounts.md) |  | [optional] 

## Methods

### NewInlineResponse200373

`func NewInlineResponse200373() *InlineResponse200373`

NewInlineResponse200373 instantiates a new InlineResponse200373 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200373WithDefaults

`func NewInlineResponse200373WithDefaults() *InlineResponse200373`

NewInlineResponse200373WithDefaults instantiates a new InlineResponse200373 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200373) GetNetwork() OrganizationsOrganizationIdWirelessZigbeeDevicesNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200373) GetNetworkOk() (*OrganizationsOrganizationIdWirelessZigbeeDevicesNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200373) SetNetwork(v OrganizationsOrganizationIdWirelessZigbeeDevicesNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200373) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPanId

`func (o *InlineResponse200373) GetPanId() string`

GetPanId returns the PanId field if non-nil, zero value otherwise.

### GetPanIdOk

`func (o *InlineResponse200373) GetPanIdOk() (*string, bool)`

GetPanIdOk returns a tuple with the PanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanId

`func (o *InlineResponse200373) SetPanId(v string)`

SetPanId sets PanId field to given value.

### HasPanId

`func (o *InlineResponse200373) HasPanId() bool`

HasPanId returns a boolean if a field has been set.

### GetChannel

`func (o *InlineResponse200373) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineResponse200373) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineResponse200373) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineResponse200373) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTransmitPowerLevel

`func (o *InlineResponse200373) GetTransmitPowerLevel() int32`

GetTransmitPowerLevel returns the TransmitPowerLevel field if non-nil, zero value otherwise.

### GetTransmitPowerLevelOk

`func (o *InlineResponse200373) GetTransmitPowerLevelOk() (*int32, bool)`

GetTransmitPowerLevelOk returns a tuple with the TransmitPowerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitPowerLevel

`func (o *InlineResponse200373) SetTransmitPowerLevel(v int32)`

SetTransmitPowerLevel sets TransmitPowerLevel field to given value.

### HasTransmitPowerLevel

`func (o *InlineResponse200373) HasTransmitPowerLevel() bool`

HasTransmitPowerLevel returns a boolean if a field has been set.

### GetEnrolled

`func (o *InlineResponse200373) GetEnrolled() bool`

GetEnrolled returns the Enrolled field if non-nil, zero value otherwise.

### GetEnrolledOk

`func (o *InlineResponse200373) GetEnrolledOk() (*bool, bool)`

GetEnrolledOk returns a tuple with the Enrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolled

`func (o *InlineResponse200373) SetEnrolled(v bool)`

SetEnrolled sets Enrolled field to given value.

### HasEnrolled

`func (o *InlineResponse200373) HasEnrolled() bool`

HasEnrolled returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200373) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200373) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200373) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200373) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGateway

`func (o *InlineResponse200373) GetGateway() OrganizationsOrganizationIdWirelessZigbeeDevicesGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse200373) GetGatewayOk() (*OrganizationsOrganizationIdWirelessZigbeeDevicesGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse200373) SetGateway(v OrganizationsOrganizationIdWirelessZigbeeDevicesGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InlineResponse200373) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200373) GetCounts() OrganizationsOrganizationIdWirelessZigbeeDevicesCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200373) GetCountsOk() (*OrganizationsOrganizationIdWirelessZigbeeDevicesCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200373) SetCounts(v OrganizationsOrganizationIdWirelessZigbeeDevicesCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200373) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


