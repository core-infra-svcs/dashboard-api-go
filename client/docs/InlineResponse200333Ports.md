# InlineResponse200333Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**LastUpdatedAt** | Pointer to **string** | Timestamp for most recent discovery info on this port. | [optional] 
**Cdp** | Pointer to [**[]InlineResponse200333Cdp**](InlineResponse200333Cdp.md) | The Cisco Discovery Protocol (CDP) information of the connected device. | [optional] 
**Lldp** | Pointer to [**[]InlineResponse200333Lldp**](InlineResponse200333Lldp.md) | The Link Layer Discovery Protocol (LLDP) information of the connected device. | [optional] 

## Methods

### NewInlineResponse200333Ports

`func NewInlineResponse200333Ports() *InlineResponse200333Ports`

NewInlineResponse200333Ports instantiates a new InlineResponse200333Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200333PortsWithDefaults

`func NewInlineResponse200333PortsWithDefaults() *InlineResponse200333Ports`

NewInlineResponse200333PortsWithDefaults instantiates a new InlineResponse200333Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse200333Ports) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse200333Ports) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse200333Ports) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse200333Ports) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200333Ports) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200333Ports) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200333Ports) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200333Ports) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse200333Ports) GetCdp() []InlineResponse200333Cdp`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse200333Ports) GetCdpOk() (*[]InlineResponse200333Cdp, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse200333Ports) SetCdp(v []InlineResponse200333Cdp)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse200333Ports) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetLldp

`func (o *InlineResponse200333Ports) GetLldp() []InlineResponse200333Lldp`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse200333Ports) GetLldpOk() (*[]InlineResponse200333Lldp, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse200333Ports) SetLldp(v []InlineResponse200333Lldp)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse200333Ports) HasLldp() bool`

HasLldp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


