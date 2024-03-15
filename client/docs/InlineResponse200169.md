# InlineResponse200169

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial of the device | [optional] 
**Name** | Pointer to **string** | Name assigned to the device | [optional] 
**DeviceStatus** | Pointer to **string** | Status of the device upgrade | [optional] 
**Upgrade** | Pointer to [**OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade**](OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade.md) |  | [optional] 

## Methods

### NewInlineResponse200169

`func NewInlineResponse200169() *InlineResponse200169`

NewInlineResponse200169 instantiates a new InlineResponse200169 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200169WithDefaults

`func NewInlineResponse200169WithDefaults() *InlineResponse200169`

NewInlineResponse200169WithDefaults instantiates a new InlineResponse200169 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200169) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200169) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200169) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200169) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200169) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200169) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200169) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200169) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *InlineResponse200169) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *InlineResponse200169) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *InlineResponse200169) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *InlineResponse200169) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetUpgrade

`func (o *InlineResponse200169) GetUpgrade() OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *InlineResponse200169) GetUpgradeOk() (*OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *InlineResponse200169) SetUpgrade(v OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgrade)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *InlineResponse200169) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


