# InlineObject302

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**[]OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices**](OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices.md) | Device details (maximum of 20 devices allowed) | 
**Notes** | Pointer to **string** | Reason for capture | [optional] 
**Duration** | Pointer to **int32** | Duration of the capture in seconds | [optional] 
**FilterExpression** | Pointer to **string** | Filter expression for the capture | [optional] 
**Name** | **string** | Name of packet capture file | 
**Advanced** | Pointer to [**OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced**](OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced.md) |  | [optional] 

## Methods

### NewInlineObject302

`func NewInlineObject302(devices []OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices, name string, ) *InlineObject302`

NewInlineObject302 instantiates a new InlineObject302 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject302WithDefaults

`func NewInlineObject302WithDefaults() *InlineObject302`

NewInlineObject302WithDefaults instantiates a new InlineObject302 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *InlineObject302) GetDevices() []OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineObject302) GetDevicesOk() (*[]OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineObject302) SetDevices(v []OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices)`

SetDevices sets Devices field to given value.


### GetNotes

`func (o *InlineObject302) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject302) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject302) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject302) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDuration

`func (o *InlineObject302) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineObject302) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineObject302) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineObject302) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterExpression

`func (o *InlineObject302) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *InlineObject302) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *InlineObject302) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *InlineObject302) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetName

`func (o *InlineObject302) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject302) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject302) SetName(v string)`

SetName sets Name field to given value.


### GetAdvanced

`func (o *InlineObject302) GetAdvanced() OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *InlineObject302) GetAdvancedOk() (*OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *InlineObject302) SetAdvanced(v OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateAdvanced)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *InlineObject302) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


