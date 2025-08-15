# InlineResponse200344Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**LastUpdatedAt** | Pointer to **string** | Timestamp for most recent discovery info on this port. | [optional] 
**Cdp** | Pointer to [**[]InlineResponse200344Cdp**](InlineResponse200344Cdp.md) | The Cisco Discovery Protocol (CDP) information of the connected device. | [optional] 
**Lldp** | Pointer to [**[]InlineResponse200344Lldp**](InlineResponse200344Lldp.md) | The Link Layer Discovery Protocol (LLDP) information of the connected device. | [optional] 

## Methods

### NewInlineResponse200344Ports

`func NewInlineResponse200344Ports() *InlineResponse200344Ports`

NewInlineResponse200344Ports instantiates a new InlineResponse200344Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200344PortsWithDefaults

`func NewInlineResponse200344PortsWithDefaults() *InlineResponse200344Ports`

NewInlineResponse200344PortsWithDefaults instantiates a new InlineResponse200344Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse200344Ports) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse200344Ports) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse200344Ports) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse200344Ports) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200344Ports) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200344Ports) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200344Ports) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200344Ports) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse200344Ports) GetCdp() []InlineResponse200344Cdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse200344Ports) GetCdpOk() (*[]InlineResponse200344Cdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse200344Ports) SetCdp(v []InlineResponse200344Cdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse200344Ports) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetLldp

`func (o *InlineResponse200344Ports) GetLldp() []InlineResponse200344Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse200344Ports) GetLldpOk() (*[]InlineResponse200344Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse200344Ports) SetLldp(v []InlineResponse200344Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse200344Ports) HasLldp() bool`

HasLldp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


