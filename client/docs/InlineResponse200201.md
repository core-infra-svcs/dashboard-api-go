# InlineResponse200201

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the switch. | [optional] 
**Serial** | Pointer to **string** | The serial number of the switch. | [optional] 
**Mac** | Pointer to **string** | The MAC address of the switch. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdSwitchPortsBySwitchNetwork**](OrganizationsOrganizationIdSwitchPortsBySwitchNetwork.md) |  | [optional] 
**Model** | Pointer to **string** | The model of the switch. | [optional] 
**Ports** | Pointer to [**[]OrganizationsOrganizationIdSwitchPortsBySwitchPorts**](OrganizationsOrganizationIdSwitchPortsBySwitchPorts.md) | Ports belonging to the switch | [optional] 

## Methods

### NewInlineResponse200201

`func NewInlineResponse200201() *InlineResponse200201`

NewInlineResponse200201 instantiates a new InlineResponse200201 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200201WithDefaults

`func NewInlineResponse200201WithDefaults() *InlineResponse200201`

NewInlineResponse200201WithDefaults instantiates a new InlineResponse200201 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200201) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200201) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200201) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200201) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200201) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200201) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200201) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200201) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200201) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200201) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200201) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200201) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200201) GetNetwork() OrganizationsOrganizationIdSwitchPortsBySwitchNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200201) GetNetworkOk() (*OrganizationsOrganizationIdSwitchPortsBySwitchNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200201) SetNetwork(v OrganizationsOrganizationIdSwitchPortsBySwitchNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200201) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200201) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200201) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200201) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200201) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200201) GetPorts() []OrganizationsOrganizationIdSwitchPortsBySwitchPorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200201) GetPortsOk() (*[]OrganizationsOrganizationIdSwitchPortsBySwitchPorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200201) SetPorts(v []OrganizationsOrganizationIdSwitchPortsBySwitchPorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200201) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


