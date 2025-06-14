# InlineResponse200338Ports

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
**SecurePort** | Pointer to [**InlineResponse200338SecurePort**](InlineResponse200338SecurePort.md) |  | [optional] 

## Methods

### NewInlineResponse200338Ports

`func NewInlineResponse200338Ports() *InlineResponse200338Ports`

NewInlineResponse200338Ports instantiates a new InlineResponse200338Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200338PortsWithDefaults

`func NewInlineResponse200338PortsWithDefaults() *InlineResponse200338Ports`

NewInlineResponse200338PortsWithDefaults instantiates a new InlineResponse200338Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse200338Ports) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse200338Ports) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse200338Ports) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse200338Ports) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200338Ports) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200338Ports) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200338Ports) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200338Ports) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200338Ports) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200338Ports) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200338Ports) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200338Ports) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsUplink

`func (o *InlineResponse200338Ports) GetIsUplink() bool`

GetIsUplink returns the IsUplink field if non-nil, zero value otherwise.

### GetIsUplinkOk

`func (o *InlineResponse200338Ports) GetIsUplinkOk() (*bool, bool)`

GetIsUplinkOk returns a tuple with the IsUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUplink

`func (o *InlineResponse200338Ports) SetIsUplink(v bool)`

SetIsUplink sets IsUplink field to given value.

### HasIsUplink

`func (o *InlineResponse200338Ports) HasIsUplink() bool`

HasIsUplink returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200338Ports) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200338Ports) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200338Ports) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200338Ports) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *InlineResponse200338Ports) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *InlineResponse200338Ports) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *InlineResponse200338Ports) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *InlineResponse200338Ports) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetSpeed

`func (o *InlineResponse200338Ports) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *InlineResponse200338Ports) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *InlineResponse200338Ports) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *InlineResponse200338Ports) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetDuplex

`func (o *InlineResponse200338Ports) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *InlineResponse200338Ports) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *InlineResponse200338Ports) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *InlineResponse200338Ports) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetSpanningTree

`func (o *InlineResponse200338Ports) GetSpanningTree() DevicesSerialSwitchPortsStatusesSpanningTree`

GetSpanningTree returns the SpanningTree field if non-nil, zero value otherwise.

### GetSpanningTreeOk

`func (o *InlineResponse200338Ports) GetSpanningTreeOk() (*DevicesSerialSwitchPortsStatusesSpanningTree, bool)`

GetSpanningTreeOk returns a tuple with the SpanningTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanningTree

`func (o *InlineResponse200338Ports) SetSpanningTree(v DevicesSerialSwitchPortsStatusesSpanningTree)`

SetSpanningTree sets SpanningTree field to given value.

### HasSpanningTree

`func (o *InlineResponse200338Ports) HasSpanningTree() bool`

HasSpanningTree returns a boolean if a field has been set.

### GetPoe

`func (o *InlineResponse200338Ports) GetPoe() DevicesSerialSwitchPortsStatusesPoe`

GetPoe returns the Poe field if non-nil, zero value otherwise.

### GetPoeOk

`func (o *InlineResponse200338Ports) GetPoeOk() (*DevicesSerialSwitchPortsStatusesPoe, bool)`

GetPoeOk returns a tuple with the Poe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoe

`func (o *InlineResponse200338Ports) SetPoe(v DevicesSerialSwitchPortsStatusesPoe)`

SetPoe sets Poe field to given value.

### HasPoe

`func (o *InlineResponse200338Ports) HasPoe() bool`

HasPoe returns a boolean if a field has been set.

### GetSecurePort

`func (o *InlineResponse200338Ports) GetSecurePort() InlineResponse200338SecurePort`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *InlineResponse200338Ports) GetSecurePortOk() (*InlineResponse200338SecurePort, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *InlineResponse200338Ports) SetSecurePort(v InlineResponse200338SecurePort)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *InlineResponse200338Ports) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


