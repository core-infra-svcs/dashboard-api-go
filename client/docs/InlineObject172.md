# InlineObject172

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | AP port profile name | 
**Ports** | [**[]NetworksNetworkIdWirelessEthernetPortsProfilesPorts1**](NetworksNetworkIdWirelessEthernetPortsProfilesPorts1.md) | AP ports configuration | 
**UsbPorts** | Pointer to [**[]NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1**](NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1.md) | AP usb ports configuration | [optional] 

## Methods

### NewInlineObject172

`func NewInlineObject172(name string, ports []NetworksNetworkIdWirelessEthernetPortsProfilesPorts1, ) *InlineObject172`

NewInlineObject172 instantiates a new InlineObject172 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject172WithDefaults

`func NewInlineObject172WithDefaults() *InlineObject172`

NewInlineObject172WithDefaults instantiates a new InlineObject172 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject172) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject172) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject172) SetName(v string)`

SetName sets Name field to given value.


### GetPorts

`func (o *InlineObject172) GetPorts() []NetworksNetworkIdWirelessEthernetPortsProfilesPorts1`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineObject172) GetPortsOk() (*[]NetworksNetworkIdWirelessEthernetPortsProfilesPorts1, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineObject172) SetPorts(v []NetworksNetworkIdWirelessEthernetPortsProfilesPorts1)`

SetPorts sets Ports field to given value.


### GetUsbPorts

`func (o *InlineObject172) GetUsbPorts() []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1`

GetUsbPorts returns the UsbPorts field if non-nil, zero value otherwise.

### GetUsbPortsOk

`func (o *InlineObject172) GetUsbPortsOk() (*[]NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1, bool)`

GetUsbPortsOk returns a tuple with the UsbPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPorts

`func (o *InlineObject172) SetUsbPorts(v []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts1)`

SetUsbPorts sets UsbPorts field to given value.

### HasUsbPorts

`func (o *InlineObject172) HasUsbPorts() bool`

HasUsbPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


