# InlineResponse20037

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**Enabled** | Pointer to **bool** | Whether the port is configured to be enabled. | [optional] 
**Status** | Pointer to **string** | The current connection status of the port. | [optional] 
**IsUplink** | Pointer to **bool** | Whether the port is the switch&#39;s uplink. | [optional] 
**Errors** | Pointer to **[]string** | All errors present on the port. | [optional] 
**Warnings** | Pointer to **[]string** | All warnings present on the port. | [optional] 
**Speed** | Pointer to **string** | The current data transfer rate which the port is operating at. | [optional] 
**Duplex** | Pointer to **string** | The current duplex of a connected port. | [optional] 
**SpanningTree** | Pointer to [**DevicesSerialSwitchPortsStatusesSpanningTree**](DevicesSerialSwitchPortsStatusesSpanningTree.md) |  | [optional] 
**Poe** | Pointer to [**DevicesSerialSwitchPortsStatusesPoe**](DevicesSerialSwitchPortsStatusesPoe.md) |  | [optional] 
**UsageInKb** | Pointer to [**DevicesSerialSwitchPortsStatusesUsageInKb**](DevicesSerialSwitchPortsStatusesUsageInKb.md) |  | [optional] 
**Cdp** | Pointer to [**DevicesSerialSwitchPortsStatusesCdp**](DevicesSerialSwitchPortsStatusesCdp.md) |  | [optional] 
**Lldp** | Pointer to [**DevicesSerialSwitchPortsStatusesLldp**](DevicesSerialSwitchPortsStatusesLldp.md) |  | [optional] 
**ClientCount** | Pointer to **int32** | The number of clients connected through this port. | [optional] 
**PowerUsageInWh** | Pointer to **float32** | How much power (in watt-hours) has been delivered by this port during the timespan. | [optional] 
**TrafficInKbps** | Pointer to [**DevicesSerialSwitchPortsStatusesTrafficInKbps**](DevicesSerialSwitchPortsStatusesTrafficInKbps.md) |  | [optional] 
**SecurePort** | Pointer to [**DevicesSerialSwitchPortsStatusesSecurePort**](DevicesSerialSwitchPortsStatusesSecurePort.md) |  | [optional] 

## Methods

### NewInlineResponse20037

`func NewInlineResponse20037() *InlineResponse20037`

NewInlineResponse20037 instantiates a new InlineResponse20037 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20037WithDefaults

`func NewInlineResponse20037WithDefaults() *InlineResponse20037`

NewInlineResponse20037WithDefaults instantiates a new InlineResponse20037 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse20037) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse20037) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse20037) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse20037) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20037) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20037) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20037) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20037) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20037) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20037) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20037) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20037) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsUplink

`func (o *InlineResponse20037) GetIsUplink() bool`

GetIsUplink returns the IsUplink field if non-nil, zero value otherwise.

### GetIsUplinkOk

`func (o *InlineResponse20037) GetIsUplinkOk() (*bool, bool)`

GetIsUplinkOk returns a tuple with the IsUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUplink

`func (o *InlineResponse20037) SetIsUplink(v bool)`

SetIsUplink sets IsUplink field to given value.

### HasIsUplink

`func (o *InlineResponse20037) HasIsUplink() bool`

HasIsUplink returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse20037) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse20037) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse20037) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse20037) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *InlineResponse20037) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *InlineResponse20037) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *InlineResponse20037) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *InlineResponse20037) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetSpeed

`func (o *InlineResponse20037) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *InlineResponse20037) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *InlineResponse20037) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *InlineResponse20037) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetDuplex

`func (o *InlineResponse20037) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *InlineResponse20037) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *InlineResponse20037) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *InlineResponse20037) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetSpanningTree

`func (o *InlineResponse20037) GetSpanningTree() DevicesSerialSwitchPortsStatusesSpanningTree`

GetSpanningTree returns the SpanningTree field if non-nil, zero value otherwise.

### GetSpanningTreeOk

`func (o *InlineResponse20037) GetSpanningTreeOk() (*DevicesSerialSwitchPortsStatusesSpanningTree, bool)`

GetSpanningTreeOk returns a tuple with the SpanningTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTree

`func (o *InlineResponse20037) SetSpanningTree(v DevicesSerialSwitchPortsStatusesSpanningTree)`

SetSpanningTree sets SpanningTree field to given value.

### HasSpanningTree

`func (o *InlineResponse20037) HasSpanningTree() bool`

HasSpanningTree returns a boolean if a field has been set.

### GetPoe

`func (o *InlineResponse20037) GetPoe() DevicesSerialSwitchPortsStatusesPoe`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *InlineResponse20037) GetPoeOk() (*DevicesSerialSwitchPortsStatusesPoe, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *InlineResponse20037) SetPoe(v DevicesSerialSwitchPortsStatusesPoe)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *InlineResponse20037) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetUsageInKb

`func (o *InlineResponse20037) GetUsageInKb() DevicesSerialSwitchPortsStatusesUsageInKb`

GetUsageInKb returns the UsageInKb field if non-nil, zero value otherwise.

### GetUsageInKbOk

`func (o *InlineResponse20037) GetUsageInKbOk() (*DevicesSerialSwitchPortsStatusesUsageInKb, bool)`

GetUsageInKbOk returns a tuple with the UsageInKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageInKb

`func (o *InlineResponse20037) SetUsageInKb(v DevicesSerialSwitchPortsStatusesUsageInKb)`

SetUsageInKb sets UsageInKb field to given value.

### HasUsageInKb

`func (o *InlineResponse20037) HasUsageInKb() bool`

HasUsageInKb returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse20037) GetCdp() DevicesSerialSwitchPortsStatusesCdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse20037) GetCdpOk() (*DevicesSerialSwitchPortsStatusesCdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse20037) SetCdp(v DevicesSerialSwitchPortsStatusesCdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse20037) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetLldp

`func (o *InlineResponse20037) GetLldp() DevicesSerialSwitchPortsStatusesLldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse20037) GetLldpOk() (*DevicesSerialSwitchPortsStatusesLldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse20037) SetLldp(v DevicesSerialSwitchPortsStatusesLldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse20037) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetClientCount

`func (o *InlineResponse20037) GetClientCount() int32`

GetClientCount returns the ClientCount field if non-nil, zero value otherwise.

### GetClientCountOk

`func (o *InlineResponse20037) GetClientCountOk() (*int32, bool)`

GetClientCountOk returns a tuple with the ClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCount

`func (o *InlineResponse20037) SetClientCount(v int32)`

SetClientCount sets ClientCount field to given value.

### HasClientCount

`func (o *InlineResponse20037) HasClientCount() bool`

HasClientCount returns a boolean if a field has been set.

### GetPowerUsageInWh

`func (o *InlineResponse20037) GetPowerUsageInWh() float32`

GetPowerUsageInWh returns the PowerUsageInWh field if non-nil, zero value otherwise.

### GetPowerUsageInWhOk

`func (o *InlineResponse20037) GetPowerUsageInWhOk() (*float32, bool)`

GetPowerUsageInWhOk returns a tuple with the PowerUsageInWh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerUsageInWh

`func (o *InlineResponse20037) SetPowerUsageInWh(v float32)`

SetPowerUsageInWh sets PowerUsageInWh field to given value.

### HasPowerUsageInWh

`func (o *InlineResponse20037) HasPowerUsageInWh() bool`

HasPowerUsageInWh returns a boolean if a field has been set.

### GetTrafficInKbps

`func (o *InlineResponse20037) GetTrafficInKbps() DevicesSerialSwitchPortsStatusesTrafficInKbps`

GetTrafficInKbps returns the TrafficInKbps field if non-nil, zero value otherwise.

### GetTrafficInKbpsOk

`func (o *InlineResponse20037) GetTrafficInKbpsOk() (*DevicesSerialSwitchPortsStatusesTrafficInKbps, bool)`

GetTrafficInKbpsOk returns a tuple with the TrafficInKbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficInKbps

`func (o *InlineResponse20037) SetTrafficInKbps(v DevicesSerialSwitchPortsStatusesTrafficInKbps)`

SetTrafficInKbps sets TrafficInKbps field to given value.

### HasTrafficInKbps

`func (o *InlineResponse20037) HasTrafficInKbps() bool`

HasTrafficInKbps returns a boolean if a field has been set.

### GetSecurePort

`func (o *InlineResponse20037) GetSecurePort() DevicesSerialSwitchPortsStatusesSecurePort`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *InlineResponse20037) GetSecurePortOk() (*DevicesSerialSwitchPortsStatusesSecurePort, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *InlineResponse20037) SetSecurePort(v DevicesSerialSwitchPortsStatusesSecurePort)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *InlineResponse20037) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


