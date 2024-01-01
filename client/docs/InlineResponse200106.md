# InlineResponse200106

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | AP port profile ID | [optional] 
**Name** | Pointer to **string** | AP port profile name | [optional] 
**IsDefault** | Pointer to **bool** | Is default profile | [optional] 
**Ports** | Pointer to [**[]NetworksNetworkIdWirelessEthernetPortsProfilesPorts**](NetworksNetworkIdWirelessEthernetPortsProfilesPorts.md) | Ports config | [optional] 
**UsbPorts** | Pointer to [**[]NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts**](NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts.md) | Usb ports config | [optional] 

## Methods

### NewInlineResponse200106

`func NewInlineResponse200106() *InlineResponse200106`

NewInlineResponse200106 instantiates a new InlineResponse200106 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200106WithDefaults

`func NewInlineResponse200106WithDefaults() *InlineResponse200106`

NewInlineResponse200106WithDefaults instantiates a new InlineResponse200106 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *InlineResponse200106) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InlineResponse200106) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InlineResponse200106) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *InlineResponse200106) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200106) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200106) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200106) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200106) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDefault

`func (o *InlineResponse200106) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *InlineResponse200106) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *InlineResponse200106) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *InlineResponse200106) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200106) GetPorts() []NetworksNetworkIdWirelessEthernetPortsProfilesPorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200106) GetPortsOk() (*[]NetworksNetworkIdWirelessEthernetPortsProfilesPorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200106) SetPorts(v []NetworksNetworkIdWirelessEthernetPortsProfilesPorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200106) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetUsbPorts

`func (o *InlineResponse200106) GetUsbPorts() []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts`

GetUsbPorts returns the UsbPorts field if non-nil, zero value otherwise.

### GetUsbPortsOk

`func (o *InlineResponse200106) GetUsbPortsOk() (*[]NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts, bool)`

GetUsbPortsOk returns a tuple with the UsbPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPorts

`func (o *InlineResponse200106) SetUsbPorts(v []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts)`

SetUsbPorts sets UsbPorts field to given value.

### HasUsbPorts

`func (o *InlineResponse200106) HasUsbPorts() bool`

HasUsbPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


