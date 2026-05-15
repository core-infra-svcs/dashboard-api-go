# InlineObject110

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Staged Upgrade Group. Length must be 1 to 255 characters | 
**Description** | Pointer to **string** | Description of the Staged Upgrade Group. Length must be 1 to 255 characters | [optional] 
**IsDefault** | **bool** | Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group | 
**AssignedDevices** | Pointer to [**NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1**](NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1.md) |  | [optional] 

## Methods

### NewInlineObject110

`func NewInlineObject110(name string, isDefault bool, ) *InlineObject110`

NewInlineObject110 instantiates a new InlineObject110 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject110WithDefaults

`func NewInlineObject110WithDefaults() *InlineObject110`

NewInlineObject110WithDefaults instantiates a new InlineObject110 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject110) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject110) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject110) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InlineObject110) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject110) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject110) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject110) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsDefault

`func (o *InlineObject110) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *InlineObject110) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *InlineObject110) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetAssignedDevices

`func (o *InlineObject110) GetAssignedDevices() NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1`

GetAssignedDevices returns the AssignedDevices field if non-nil, zero value otherwise.

### GetAssignedDevicesOk

`func (o *InlineObject110) GetAssignedDevicesOk() (*NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1, bool)`

GetAssignedDevicesOk returns a tuple with the AssignedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedDevices

`func (o *InlineObject110) SetAssignedDevices(v NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1)`

SetAssignedDevices sets AssignedDevices field to given value.

### HasAssignedDevices

`func (o *InlineObject110) HasAssignedDevices() bool`

HasAssignedDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


