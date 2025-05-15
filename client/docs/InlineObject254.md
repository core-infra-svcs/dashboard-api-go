# InlineObject254

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**[]OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices**](OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices.md) | Device details (maximum of 20 devices allowed) | 
**Notes** | Pointer to **string** | Reason for capture | [optional] 
**Duration** | Pointer to **int32** | Duration of the capture in seconds | [optional] 
**FilterExpression** | Pointer to **string** | Filter expression for the capture | [optional] 
**Name** | **string** | Name of packet capture file | 

## Methods

### NewInlineObject254

`func NewInlineObject254(devices []OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices, name string, ) *InlineObject254`

NewInlineObject254 instantiates a new InlineObject254 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject254WithDefaults

`func NewInlineObject254WithDefaults() *InlineObject254`

NewInlineObject254WithDefaults instantiates a new InlineObject254 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *InlineObject254) GetDevices() []OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineObject254) GetDevicesOk() (*[]OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineObject254) SetDevices(v []OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices)`

SetDevices sets Devices field to given value.


### GetNotes

`func (o *InlineObject254) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject254) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject254) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject254) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDuration

`func (o *InlineObject254) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineObject254) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineObject254) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineObject254) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterExpression

`func (o *InlineObject254) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *InlineObject254) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *InlineObject254) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *InlineObject254) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetName

`func (o *InlineObject254) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject254) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject254) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


