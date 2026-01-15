# OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhysicalInterfaceDirection** | Pointer to **string** | Direction for physical interface capture | [optional] 
**ControlPlaneDirection** | Pointer to **string** | Direction for control plane capture | [optional] 
**InnerFilterMacs** | Pointer to **[]string** | Inner MAC address filter for tunneled traffic (up to 5 MAC addresses for Campus Gateway devices) | [optional] 
**BufferFiles** | Pointer to **int32** | Number of buffer files for circular buffer capture (1-5 for Campus Gateway devices) | [optional] 
**MaxFilesize** | Pointer to **int32** | Maximum file size in megabytes (MB). Range: 1-100 MB when bufferFiles&#x3D;1, 1-500 MB when bufferFiles&#x3D;2-5 | [optional] 
**CaptureType** | Pointer to **string** | Type of capture. Possible values: linear (default), circular | [optional] 
**PacketsPerSecond** | Pointer to **int32** | Packets per second limit for Campus Gateway devices (1-1000000). | [optional] 

## Methods

### NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced

`func NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced() *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced`

NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced instantiates a new OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvancedWithDefaults

`func NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvancedWithDefaults() *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced`

NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvancedWithDefaults instantiates a new OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhysicalInterfaceDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetPhysicalInterfaceDirection() string`

GetPhysicalInterfaceDirection returns the PhysicalInterfaceDirection field if non-nil, zero value otherwise.

### GetPhysicalInterfaceDirectionOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetPhysicalInterfaceDirectionOk() (*string, bool)`

GetPhysicalInterfaceDirectionOk returns a tuple with the PhysicalInterfaceDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalInterfaceDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) SetPhysicalInterfaceDirection(v string)`

SetPhysicalInterfaceDirection sets PhysicalInterfaceDirection field to given value.

### HasPhysicalInterfaceDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) HasPhysicalInterfaceDirection() bool`

HasPhysicalInterfaceDirection returns a boolean if a field has been set.

### GetControlPlaneDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetControlPlaneDirection() string`

GetControlPlaneDirection returns the ControlPlaneDirection field if non-nil, zero value otherwise.

### GetControlPlaneDirectionOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetControlPlaneDirectionOk() (*string, bool)`

GetControlPlaneDirectionOk returns a tuple with the ControlPlaneDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) SetControlPlaneDirection(v string)`

SetControlPlaneDirection sets ControlPlaneDirection field to given value.

### HasControlPlaneDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) HasControlPlaneDirection() bool`

HasControlPlaneDirection returns a boolean if a field has been set.

### GetInnerFilterMacs

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetInnerFilterMacs() []string`

GetInnerFilterMacs returns the InnerFilterMacs field if non-nil, zero value otherwise.

### GetInnerFilterMacsOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetInnerFilterMacsOk() (*[]string, bool)`

GetInnerFilterMacsOk returns a tuple with the InnerFilterMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerFilterMacs

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) SetInnerFilterMacs(v []string)`

SetInnerFilterMacs sets InnerFilterMacs field to given value.

### HasInnerFilterMacs

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) HasInnerFilterMacs() bool`

HasInnerFilterMacs returns a boolean if a field has been set.

### GetBufferFiles

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetBufferFiles() int32`

GetBufferFiles returns the BufferFiles field if non-nil, zero value otherwise.

### GetBufferFilesOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetBufferFilesOk() (*int32, bool)`

GetBufferFilesOk returns a tuple with the BufferFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferFiles

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) SetBufferFiles(v int32)`

SetBufferFiles sets BufferFiles field to given value.

### HasBufferFiles

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) HasBufferFiles() bool`

HasBufferFiles returns a boolean if a field has been set.

### GetMaxFilesize

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetMaxFilesize() int32`

GetMaxFilesize returns the MaxFilesize field if non-nil, zero value otherwise.

### GetMaxFilesizeOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetMaxFilesizeOk() (*int32, bool)`

GetMaxFilesizeOk returns a tuple with the MaxFilesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFilesize

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) SetMaxFilesize(v int32)`

SetMaxFilesize sets MaxFilesize field to given value.

### HasMaxFilesize

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) HasMaxFilesize() bool`

HasMaxFilesize returns a boolean if a field has been set.

### GetCaptureType

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetCaptureType() string`

GetCaptureType returns the CaptureType field if non-nil, zero value otherwise.

### GetCaptureTypeOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetCaptureTypeOk() (*string, bool)`

GetCaptureTypeOk returns a tuple with the CaptureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureType

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) SetCaptureType(v string)`

SetCaptureType sets CaptureType field to given value.

### HasCaptureType

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) HasCaptureType() bool`

HasCaptureType returns a boolean if a field has been set.

### GetPacketsPerSecond

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetPacketsPerSecond() int32`

GetPacketsPerSecond returns the PacketsPerSecond field if non-nil, zero value otherwise.

### GetPacketsPerSecondOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) GetPacketsPerSecondOk() (*int32, bool)`

GetPacketsPerSecondOk returns a tuple with the PacketsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsPerSecond

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) SetPacketsPerSecond(v int32)`

SetPacketsPerSecond sets PacketsPerSecond field to given value.

### HasPacketsPerSecond

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced) HasPacketsPerSecond() bool`

HasPacketsPerSecond returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


