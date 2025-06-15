# InlineResponse20030Lldp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemName** | Pointer to **string** | Device system name | [optional] 
**SystemDescription** | Pointer to **string** | Device system description | [optional] 
**PortId** | Pointer to **string** | ID for the port | [optional] 
**ChassisId** | Pointer to **string** | Chassis ID | [optional] 
**ManagementVlan** | Pointer to **int32** | Management VLAN | [optional] 
**PortVlan** | Pointer to **int32** | Port VLAN | [optional] 
**ManagementAddress** | Pointer to **string** | Management IP address | [optional] 
**PortDescription** | Pointer to **string** | Port description | [optional] 
**SystemCapabilities** | Pointer to **string** | System capabilities | [optional] 
**SourcePort** | Pointer to **string** | Source port | [optional] 

## Methods

### NewInlineResponse20030Lldp

`func NewInlineResponse20030Lldp() *InlineResponse20030Lldp`

NewInlineResponse20030Lldp instantiates a new InlineResponse20030Lldp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20030LldpWithDefaults

`func NewInlineResponse20030LldpWithDefaults() *InlineResponse20030Lldp`

NewInlineResponse20030LldpWithDefaults instantiates a new InlineResponse20030Lldp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemName

`func (o *InlineResponse20030Lldp) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *InlineResponse20030Lldp) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *InlineResponse20030Lldp) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *InlineResponse20030Lldp) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetSystemDescription

`func (o *InlineResponse20030Lldp) GetSystemDescription() string`

GetSystemDescription returns the SystemDescription field if non-nil, zero value otherwise.

### GetSystemDescriptionOk

`func (o *InlineResponse20030Lldp) GetSystemDescriptionOk() (*string, bool)`

GetSystemDescriptionOk returns a tuple with the SystemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDescription

`func (o *InlineResponse20030Lldp) SetSystemDescription(v string)`

SetSystemDescription sets SystemDescription field to given value.

### HasSystemDescription

`func (o *InlineResponse20030Lldp) HasSystemDescription() bool`

HasSystemDescription returns a boolean if a field has been set.

### GetPortId

`func (o *InlineResponse20030Lldp) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse20030Lldp) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse20030Lldp) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse20030Lldp) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetChassisId

`func (o *InlineResponse20030Lldp) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *InlineResponse20030Lldp) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *InlineResponse20030Lldp) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *InlineResponse20030Lldp) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetManagementVlan

`func (o *InlineResponse20030Lldp) GetManagementVlan() int32`

GetManagementVlan returns the ManagementVlan field if non-nil, zero value otherwise.

### GetManagementVlanOk

`func (o *InlineResponse20030Lldp) GetManagementVlanOk() (*int32, bool)`

GetManagementVlanOk returns a tuple with the ManagementVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementVlan

`func (o *InlineResponse20030Lldp) SetManagementVlan(v int32)`

SetManagementVlan sets ManagementVlan field to given value.

### HasManagementVlan

`func (o *InlineResponse20030Lldp) HasManagementVlan() bool`

HasManagementVlan returns a boolean if a field has been set.

### GetPortVlan

`func (o *InlineResponse20030Lldp) GetPortVlan() int32`

GetPortVlan returns the PortVlan field if non-nil, zero value otherwise.

### GetPortVlanOk

`func (o *InlineResponse20030Lldp) GetPortVlanOk() (*int32, bool)`

GetPortVlanOk returns a tuple with the PortVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlan

`func (o *InlineResponse20030Lldp) SetPortVlan(v int32)`

SetPortVlan sets PortVlan field to given value.

### HasPortVlan

`func (o *InlineResponse20030Lldp) HasPortVlan() bool`

HasPortVlan returns a boolean if a field has been set.

### GetManagementAddress

`func (o *InlineResponse20030Lldp) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *InlineResponse20030Lldp) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *InlineResponse20030Lldp) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *InlineResponse20030Lldp) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetPortDescription

`func (o *InlineResponse20030Lldp) GetPortDescription() string`

GetPortDescription returns the PortDescription field if non-nil, zero value otherwise.

### GetPortDescriptionOk

`func (o *InlineResponse20030Lldp) GetPortDescriptionOk() (*string, bool)`

GetPortDescriptionOk returns a tuple with the PortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDescription

`func (o *InlineResponse20030Lldp) SetPortDescription(v string)`

SetPortDescription sets PortDescription field to given value.

### HasPortDescription

`func (o *InlineResponse20030Lldp) HasPortDescription() bool`

HasPortDescription returns a boolean if a field has been set.

### GetSystemCapabilities

`func (o *InlineResponse20030Lldp) GetSystemCapabilities() string`

GetSystemCapabilities returns the SystemCapabilities field if non-nil, zero value otherwise.

### GetSystemCapabilitiesOk

`func (o *InlineResponse20030Lldp) GetSystemCapabilitiesOk() (*string, bool)`

GetSystemCapabilitiesOk returns a tuple with the SystemCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCapabilities

`func (o *InlineResponse20030Lldp) SetSystemCapabilities(v string)`

SetSystemCapabilities sets SystemCapabilities field to given value.

### HasSystemCapabilities

`func (o *InlineResponse20030Lldp) HasSystemCapabilities() bool`

HasSystemCapabilities returns a boolean if a field has been set.

### GetSourcePort

`func (o *InlineResponse20030Lldp) GetSourcePort() string`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *InlineResponse20030Lldp) GetSourcePortOk() (*string, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *InlineResponse20030Lldp) SetSourcePort(v string)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *InlineResponse20030Lldp) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


