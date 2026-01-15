# OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhysicalInterfaceDirection** | Pointer to **string** | Direction for physical interface capture | [optional] 
**ControlPlaneDirection** | Pointer to **string** | Direction for control plane capture | [optional] 
**InnerFilterMac** | Pointer to **[]string** | Inner MAC address filter for tunneled traffic (up to 5 MAC addresses for Campus Gateway devices) | [optional] 
**BufferFiles** | Pointer to **int32** | Number of buffer files for circular buffer capture (1-5 for Campus Gateway devices) | [optional] 
**MaxFilesize** | Pointer to **int32** | Maximum file size in megabytes (MB). Range: 1-100 MB when bufferFiles&#x3D;1, 1-500 MB when bufferFiles&#x3D;2-5 | [optional] 
**CaptureType** | Pointer to **string** | Type of capture. Possible values: linear (default), circular | [optional] 

## Methods

### NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced

`func NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced() *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced`

NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced instantiates a new OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvancedWithDefaults

`func NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvancedWithDefaults() *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced`

NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvancedWithDefaults instantiates a new OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhysicalInterfaceDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetPhysicalInterfaceDirection() string`

GetPhysicalInterfaceDirection returns the PhysicalInterfaceDirection field if non-nil, zero value otherwise.

### GetPhysicalInterfaceDirectionOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetPhysicalInterfaceDirectionOk() (*string, bool)`

GetPhysicalInterfaceDirectionOk returns a tuple with the PhysicalInterfaceDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalInterfaceDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) SetPhysicalInterfaceDirection(v string)`

SetPhysicalInterfaceDirection sets PhysicalInterfaceDirection field to given value.

### HasPhysicalInterfaceDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) HasPhysicalInterfaceDirection() bool`

HasPhysicalInterfaceDirection returns a boolean if a field has been set.

### GetControlPlaneDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetControlPlaneDirection() string`

GetControlPlaneDirection returns the ControlPlaneDirection field if non-nil, zero value otherwise.

### GetControlPlaneDirectionOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetControlPlaneDirectionOk() (*string, bool)`

GetControlPlaneDirectionOk returns a tuple with the ControlPlaneDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) SetControlPlaneDirection(v string)`

SetControlPlaneDirection sets ControlPlaneDirection field to given value.

### HasControlPlaneDirection

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) HasControlPlaneDirection() bool`

HasControlPlaneDirection returns a boolean if a field has been set.

### GetInnerFilterMac

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetInnerFilterMac() []string`

GetInnerFilterMac returns the InnerFilterMac field if non-nil, zero value otherwise.

### GetInnerFilterMacOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetInnerFilterMacOk() (*[]string, bool)`

GetInnerFilterMacOk returns a tuple with the InnerFilterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerFilterMac

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) SetInnerFilterMac(v []string)`

SetInnerFilterMac sets InnerFilterMac field to given value.

### HasInnerFilterMac

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) HasInnerFilterMac() bool`

HasInnerFilterMac returns a boolean if a field has been set.

### GetBufferFiles

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetBufferFiles() int32`

GetBufferFiles returns the BufferFiles field if non-nil, zero value otherwise.

### GetBufferFilesOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetBufferFilesOk() (*int32, bool)`

GetBufferFilesOk returns a tuple with the BufferFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferFiles

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) SetBufferFiles(v int32)`

SetBufferFiles sets BufferFiles field to given value.

### HasBufferFiles

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) HasBufferFiles() bool`

HasBufferFiles returns a boolean if a field has been set.

### GetMaxFilesize

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetMaxFilesize() int32`

GetMaxFilesize returns the MaxFilesize field if non-nil, zero value otherwise.

### GetMaxFilesizeOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetMaxFilesizeOk() (*int32, bool)`

GetMaxFilesizeOk returns a tuple with the MaxFilesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFilesize

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) SetMaxFilesize(v int32)`

SetMaxFilesize sets MaxFilesize field to given value.

### HasMaxFilesize

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) HasMaxFilesize() bool`

HasMaxFilesize returns a boolean if a field has been set.

### GetCaptureType

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetCaptureType() string`

GetCaptureType returns the CaptureType field if non-nil, zero value otherwise.

### GetCaptureTypeOk

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) GetCaptureTypeOk() (*string, bool)`

GetCaptureTypeOk returns a tuple with the CaptureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureType

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) SetCaptureType(v string)`

SetCaptureType sets CaptureType field to given value.

### HasCaptureType

`func (o *OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced) HasCaptureType() bool`

HasCaptureType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


