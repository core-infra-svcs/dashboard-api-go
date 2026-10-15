# InlineResponse200475

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the switch. | 
**Serial** | **string** | The serial number of the switch. | 
**Mac** | **string** | The MAC address of the switch. | 
**Network** | [**OrganizationsOrganizationIdSwitchPortsBySwitchNetwork**](OrganizationsOrganizationIdSwitchPortsBySwitchNetwork.md) |  | 
**Model** | **string** | The model of the switch. | 
**Ports** | [**[]OrganizationsOrganizationIdSwitchPortsBySwitchPorts**](OrganizationsOrganizationIdSwitchPortsBySwitchPorts.md) | Ports belonging to the switch | 

## Methods

### NewInlineResponse200475

`func NewInlineResponse200475(name string, serial string, mac string, network OrganizationsOrganizationIdSwitchPortsBySwitchNetwork, model string, ports []OrganizationsOrganizationIdSwitchPortsBySwitchPorts, ) *InlineResponse200475`

NewInlineResponse200475 instantiates a new InlineResponse200475 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200475WithDefaults

`func NewInlineResponse200475WithDefaults() *InlineResponse200475`

NewInlineResponse200475WithDefaults instantiates a new InlineResponse200475 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200475) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200475) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200475) SetName(v string)`

SetName sets Name field to given value.


### GetSerial

`func (o *InlineResponse200475) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200475) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200475) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetMac

`func (o *InlineResponse200475) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200475) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200475) SetMac(v string)`

SetMac sets Mac field to given value.


### GetNetwork

`func (o *InlineResponse200475) GetNetwork() OrganizationsOrganizationIdSwitchPortsBySwitchNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200475) GetNetworkOk() (*OrganizationsOrganizationIdSwitchPortsBySwitchNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200475) SetNetwork(v OrganizationsOrganizationIdSwitchPortsBySwitchNetwork)`

SetNetwork sets Network field to given value.


### GetModel

`func (o *InlineResponse200475) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200475) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200475) SetModel(v string)`

SetModel sets Model field to given value.


### GetPorts

`func (o *InlineResponse200475) GetPorts() []OrganizationsOrganizationIdSwitchPortsBySwitchPorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200475) GetPortsOk() (*[]OrganizationsOrganizationIdSwitchPortsBySwitchPorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200475) SetPorts(v []OrganizationsOrganizationIdSwitchPortsBySwitchPorts)`

SetPorts sets Ports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


