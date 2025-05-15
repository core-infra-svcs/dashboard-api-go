# InlineResponse200268Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptureId** | Pointer to **string** | Id of packet capture file | [optional] 
**Network** | Pointer to [**InlineResponse200268Network**](InlineResponse200268Network.md) |  | [optional] 
**Devices** | Pointer to **[]map[string]interface{}** | Device(s) of the packet capture file | [optional] 
**Device** | Pointer to [**InlineResponse200268Device**](InlineResponse200268Device.md) |  | [optional] 
**Admin** | Pointer to [**InlineResponse200268Admin**](InlineResponse200268Admin.md) |  | [optional] 
**Client** | Pointer to [**InlineResponse200268Client**](InlineResponse200268Client.md) |  | [optional] 
**Details** | Pointer to [**[]InlineResponse200268Details**](InlineResponse200268Details.md) | Array of device specific details | [optional] 
**Name** | Pointer to **string** | Name of packet capture file | [optional] 
**StartTs** | Pointer to **string** | Start time of creation of packet capture file | [optional] 
**Ports** | Pointer to **string** | Ports of packet capture file | [optional] 
**Status** | Pointer to **string** | Status of packet capture file | [optional] 
**ErrorMessage** | Pointer to **string** | Error log of packet capture file | [optional] 
**Destination** | Pointer to **string** | Destination of packet capture file | [optional] 
**Process** | Pointer to **string** | Source of packet capture file | [optional] 
**File** | Pointer to [**InlineResponse200268File**](InlineResponse200268File.md) |  | [optional] 
**Duration** | Pointer to **int32** | Duration of packet capture file | [optional] 
**FilterExpression** | Pointer to **string** | Filter expression for the packet capture | [optional] 
**Counts** | Pointer to [**InlineResponse200268Counts**](InlineResponse200268Counts.md) |  | [optional] 
**Interface** | Pointer to **string** | Interface of the packet capture | [optional] 

## Methods

### NewInlineResponse200268Items

`func NewInlineResponse200268Items() *InlineResponse200268Items`

NewInlineResponse200268Items instantiates a new InlineResponse200268Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200268ItemsWithDefaults

`func NewInlineResponse200268ItemsWithDefaults() *InlineResponse200268Items`

NewInlineResponse200268ItemsWithDefaults instantiates a new InlineResponse200268Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptureId

`func (o *InlineResponse200268Items) GetCaptureId() string`

GetCaptureId returns the CaptureId field if non-nil, zero value otherwise.

### GetCaptureIdOk

`func (o *InlineResponse200268Items) GetCaptureIdOk() (*string, bool)`

GetCaptureIdOk returns a tuple with the CaptureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureId

`func (o *InlineResponse200268Items) SetCaptureId(v string)`

SetCaptureId sets CaptureId field to given value.

### HasCaptureId

`func (o *InlineResponse200268Items) HasCaptureId() bool`

HasCaptureId returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200268Items) GetNetwork() InlineResponse200268Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200268Items) GetNetworkOk() (*InlineResponse200268Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200268Items) SetNetwork(v InlineResponse200268Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200268Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200268Items) GetDevices() []map[string]interface{}`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200268Items) GetDevicesOk() (*[]map[string]interface{}, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200268Items) SetDevices(v []map[string]interface{})`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200268Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200268Items) GetDevice() InlineResponse200268Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200268Items) GetDeviceOk() (*InlineResponse200268Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200268Items) SetDevice(v InlineResponse200268Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200268Items) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetAdmin

`func (o *InlineResponse200268Items) GetAdmin() InlineResponse200268Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *InlineResponse200268Items) GetAdminOk() (*InlineResponse200268Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *InlineResponse200268Items) SetAdmin(v InlineResponse200268Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *InlineResponse200268Items) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetClient

`func (o *InlineResponse200268Items) GetClient() InlineResponse200268Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *InlineResponse200268Items) GetClientOk() (*InlineResponse200268Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *InlineResponse200268Items) SetClient(v InlineResponse200268Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *InlineResponse200268Items) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200268Items) GetDetails() []InlineResponse200268Details`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200268Items) GetDetailsOk() (*[]InlineResponse200268Details, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200268Items) SetDetails(v []InlineResponse200268Details)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200268Items) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200268Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200268Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200268Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200268Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTs

`func (o *InlineResponse200268Items) GetStartTs() string`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200268Items) GetStartTsOk() (*string, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200268Items) SetStartTs(v string)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200268Items) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200268Items) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200268Items) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200268Items) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200268Items) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200268Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200268Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200268Items) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200268Items) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *InlineResponse200268Items) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InlineResponse200268Items) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InlineResponse200268Items) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *InlineResponse200268Items) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetDestination

`func (o *InlineResponse200268Items) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse200268Items) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse200268Items) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *InlineResponse200268Items) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetProcess

`func (o *InlineResponse200268Items) GetProcess() string`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *InlineResponse200268Items) GetProcessOk() (*string, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *InlineResponse200268Items) SetProcess(v string)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *InlineResponse200268Items) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetFile

`func (o *InlineResponse200268Items) GetFile() InlineResponse200268File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *InlineResponse200268Items) GetFileOk() (*InlineResponse200268File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *InlineResponse200268Items) SetFile(v InlineResponse200268File)`

SetFile sets File field to given value.

### HasFile

`func (o *InlineResponse200268Items) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetDuration

`func (o *InlineResponse200268Items) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineResponse200268Items) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineResponse200268Items) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineResponse200268Items) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterExpression

`func (o *InlineResponse200268Items) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *InlineResponse200268Items) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *InlineResponse200268Items) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *InlineResponse200268Items) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200268Items) GetCounts() InlineResponse200268Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200268Items) GetCountsOk() (*InlineResponse200268Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200268Items) SetCounts(v InlineResponse200268Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200268Items) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetInterface

`func (o *InlineResponse200268Items) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse200268Items) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse200268Items) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse200268Items) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


