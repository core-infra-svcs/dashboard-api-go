# InlineObject301

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serials** | **[]string** | The serial(s) of the device(s) | 
**Name** | **string** | Name of packet capture file | 
**OutputType** | Pointer to **string** | Output type of packet capture file. Possible values: text, pcap, cloudshark, or upload_to_cloud | [optional] 
**Destination** | Pointer to **string** | Destination of packet capture file. Possible values: [upload_to_cloud] | [optional] 
**Ports** | Pointer to **string** | Ports of packet capture file, comma-separated | [optional] 
**Notes** | Pointer to **string** | Reason for taking the packet capture | [optional] 
**Duration** | Pointer to **int32** | Duration in seconds of packet capture | [optional] 
**FilterExpression** | Pointer to **string** | Filter expression for packet capture | [optional] 
**Interface** | Pointer to **string** | Interface of the device | [optional] 
**Advanced** | Pointer to [**OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced**](OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced.md) |  | [optional] 

## Methods

### NewInlineObject301

`func NewInlineObject301(serials []string, name string, ) *InlineObject301`

NewInlineObject301 instantiates a new InlineObject301 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject301WithDefaults

`func NewInlineObject301WithDefaults() *InlineObject301`

NewInlineObject301WithDefaults instantiates a new InlineObject301 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerials

`func (o *InlineObject301) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject301) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject301) SetSerials(v []string)`

SetSerials sets Serials field to given value.


### GetName

`func (o *InlineObject301) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject301) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject301) SetName(v string)`

SetName sets Name field to given value.


### GetOutputType

`func (o *InlineObject301) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *InlineObject301) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *InlineObject301) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *InlineObject301) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### GetDestination

`func (o *InlineObject301) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineObject301) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineObject301) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *InlineObject301) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetPorts

`func (o *InlineObject301) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineObject301) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineObject301) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineObject301) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetNotes

`func (o *InlineObject301) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject301) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject301) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject301) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDuration

`func (o *InlineObject301) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineObject301) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineObject301) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineObject301) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterExpression

`func (o *InlineObject301) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *InlineObject301) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *InlineObject301) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *InlineObject301) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetInterface

`func (o *InlineObject301) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineObject301) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineObject301) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineObject301) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetAdvanced

`func (o *InlineObject301) GetAdvanced() OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *InlineObject301) GetAdvancedOk() (*OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *InlineObject301) SetAdvanced(v OrganizationsOrganizationIdDevicesPacketCaptureCapturesAdvanced)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *InlineObject301) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


