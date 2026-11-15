# InlineResponse200221

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Device | [optional] 
**Serial** | Pointer to **string** | Serial of the Device | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**ProductType** | Pointer to **string** | The product type | [optional] 
**ConfigurationSource** | Pointer to **string** | Indicates whether the device configuration is managed by the cloud or the device | [optional] 
**VlanProfile** | Pointer to [**NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile**](NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile.md) |  | [optional] 
**Stack** | Pointer to [**NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack**](NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack.md) |  | [optional] 

## Methods

### NewInlineResponse200221

`func NewInlineResponse200221() *InlineResponse200221`

NewInlineResponse200221 instantiates a new InlineResponse200221 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200221WithDefaults

`func NewInlineResponse200221WithDefaults() *InlineResponse200221`

NewInlineResponse200221WithDefaults instantiates a new InlineResponse200221 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200221) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200221) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200221) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200221) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200221) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200221) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200221) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200221) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200221) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200221) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200221) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200221) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetProductType

`func (o *InlineResponse200221) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InlineResponse200221) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InlineResponse200221) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InlineResponse200221) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetConfigurationSource

`func (o *InlineResponse200221) GetConfigurationSource() string`

GetConfigurationSource returns the ConfigurationSource field if non-nil, zero value otherwise.

### GetConfigurationSourceOk

`func (o *InlineResponse200221) GetConfigurationSourceOk() (*string, bool)`

GetConfigurationSourceOk returns a tuple with the ConfigurationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationSource

`func (o *InlineResponse200221) SetConfigurationSource(v string)`

SetConfigurationSource sets ConfigurationSource field to given value.

### HasConfigurationSource

`func (o *InlineResponse200221) HasConfigurationSource() bool`

HasConfigurationSource returns a boolean if a field has been set.

### GetVlanProfile

`func (o *InlineResponse200221) GetVlanProfile() NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile`

GetVlanProfile returns the VlanProfile field if non-nil, zero value otherwise.

### GetVlanProfileOk

`func (o *InlineResponse200221) GetVlanProfileOk() (*NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile, bool)`

GetVlanProfileOk returns a tuple with the VlanProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanProfile

`func (o *InlineResponse200221) SetVlanProfile(v NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile)`

SetVlanProfile sets VlanProfile field to given value.

### HasVlanProfile

`func (o *InlineResponse200221) HasVlanProfile() bool`

HasVlanProfile returns a boolean if a field has been set.

### GetStack

`func (o *InlineResponse200221) GetStack() NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *InlineResponse200221) GetStackOk() (*NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *InlineResponse200221) SetStack(v NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack)`

SetStack sets Stack field to given value.

### HasStack

`func (o *InlineResponse200221) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


