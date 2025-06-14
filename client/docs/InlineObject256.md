# InlineObject256

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

## Methods

### NewInlineObject256

`func NewInlineObject256(serials []string, name string, ) *InlineObject256`

NewInlineObject256 instantiates a new InlineObject256 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject256WithDefaults

`func NewInlineObject256WithDefaults() *InlineObject256`

NewInlineObject256WithDefaults instantiates a new InlineObject256 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerials

`func (o *InlineObject256) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject256) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject256) SetSerials(v []string)`

SetSerials sets Serials field to given value.


### GetName

`func (o *InlineObject256) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject256) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject256) SetName(v string)`

SetName sets Name field to given value.


### GetOutputType

`func (o *InlineObject256) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *InlineObject256) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *InlineObject256) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *InlineObject256) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### GetDestination

`func (o *InlineObject256) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineObject256) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineObject256) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *InlineObject256) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetPorts

`func (o *InlineObject256) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineObject256) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineObject256) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineObject256) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetNotes

`func (o *InlineObject256) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject256) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject256) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject256) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDuration

`func (o *InlineObject256) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineObject256) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineObject256) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineObject256) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterExpression

`func (o *InlineObject256) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *InlineObject256) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *InlineObject256) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *InlineObject256) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetInterface

`func (o *InlineObject256) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineObject256) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineObject256) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineObject256) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


