# InlineObject258

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**[]OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices**](OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices.md) | Device details (maximum of 20 devices allowed) | 
**Notes** | Pointer to **string** | Reason for capture | [optional] 
**Duration** | Pointer to **int32** | Duration of the capture in seconds | [optional] 
**FilterExpression** | Pointer to **string** | Filter expression for the capture | [optional] 
**Name** | **string** | Name of packet capture file | 

## Methods

### NewInlineObject258

`func NewInlineObject258(devices []OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices, name string, ) *InlineObject258`

NewInlineObject258 instantiates a new InlineObject258 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject258WithDefaults

`func NewInlineObject258WithDefaults() *InlineObject258`

NewInlineObject258WithDefaults instantiates a new InlineObject258 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *InlineObject258) GetDevices() []OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineObject258) GetDevicesOk() (*[]OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineObject258) SetDevices(v []OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices)`

SetDevices sets Devices field to given value.


### GetNotes

`func (o *InlineObject258) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject258) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject258) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject258) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDuration

`func (o *InlineObject258) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineObject258) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineObject258) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineObject258) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterExpression

`func (o *InlineObject258) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *InlineObject258) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *InlineObject258) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *InlineObject258) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetName

`func (o *InlineObject258) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject258) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject258) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


