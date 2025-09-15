# InlineResponse200278Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptureId** | Pointer to **string** | Id of packet capture file | [optional] 
**Network** | Pointer to [**InlineResponse200278Network**](InlineResponse200278Network.md) |  | [optional] 
**Devices** | Pointer to **[]map[string]interface{}** | Device(s) of the packet capture file | [optional] 
**Device** | Pointer to [**InlineResponse200278Device**](InlineResponse200278Device.md) |  | [optional] 
**Admin** | Pointer to [**InlineResponse200278Admin**](InlineResponse200278Admin.md) |  | [optional] 
**Client** | Pointer to [**InlineResponse200278Client**](InlineResponse200278Client.md) |  | [optional] 
**Details** | Pointer to [**[]InlineResponse200278Details**](InlineResponse200278Details.md) | Array of device specific details | [optional] 
**Name** | Pointer to **string** | Name of packet capture file | [optional] 
**StartTs** | Pointer to **string** | Start time of creation of packet capture file | [optional] 
**Ports** | Pointer to **string** | Ports of packet capture file | [optional] 
**Status** | Pointer to **string** | Status of packet capture file | [optional] 
**ErrorMessage** | Pointer to **string** | Error log of packet capture file | [optional] 
**Destination** | Pointer to **string** | Destination of packet capture file | [optional] 
**Process** | Pointer to **string** | Source of packet capture file | [optional] 
**File** | Pointer to [**InlineResponse200278File**](InlineResponse200278File.md) |  | [optional] 
**Duration** | Pointer to **int32** | Duration of packet capture file | [optional] 
**FilterExpression** | Pointer to **string** | Filter expression for the packet capture | [optional] 
**Counts** | Pointer to [**InlineResponse200278Counts**](InlineResponse200278Counts.md) |  | [optional] 
**Interface** | Pointer to **string** | Interface of the packet capture | [optional] 

## Methods

### NewInlineResponse200278Items

`func NewInlineResponse200278Items() *InlineResponse200278Items`

NewInlineResponse200278Items instantiates a new InlineResponse200278Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200278ItemsWithDefaults

`func NewInlineResponse200278ItemsWithDefaults() *InlineResponse200278Items`

NewInlineResponse200278ItemsWithDefaults instantiates a new InlineResponse200278Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptureId

`func (o *InlineResponse200278Items) GetCaptureId() string`

GetCaptureId returns the CaptureId field if non-nil, zero value otherwise.

### GetCaptureIdOk

`func (o *InlineResponse200278Items) GetCaptureIdOk() (*string, bool)`

GetCaptureIdOk returns a tuple with the CaptureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureId

`func (o *InlineResponse200278Items) SetCaptureId(v string)`

SetCaptureId sets CaptureId field to given value.

### HasCaptureId

`func (o *InlineResponse200278Items) HasCaptureId() bool`

HasCaptureId returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200278Items) GetNetwork() InlineResponse200278Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200278Items) GetNetworkOk() (*InlineResponse200278Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200278Items) SetNetwork(v InlineResponse200278Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200278Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200278Items) GetDevices() []map[string]interface{}`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200278Items) GetDevicesOk() (*[]map[string]interface{}, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200278Items) SetDevices(v []map[string]interface{})`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200278Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200278Items) GetDevice() InlineResponse200278Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200278Items) GetDeviceOk() (*InlineResponse200278Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200278Items) SetDevice(v InlineResponse200278Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200278Items) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetAdmin

`func (o *InlineResponse200278Items) GetAdmin() InlineResponse200278Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *InlineResponse200278Items) GetAdminOk() (*InlineResponse200278Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *InlineResponse200278Items) SetAdmin(v InlineResponse200278Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *InlineResponse200278Items) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetClient

`func (o *InlineResponse200278Items) GetClient() InlineResponse200278Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *InlineResponse200278Items) GetClientOk() (*InlineResponse200278Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *InlineResponse200278Items) SetClient(v InlineResponse200278Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *InlineResponse200278Items) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200278Items) GetDetails() []InlineResponse200278Details`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200278Items) GetDetailsOk() (*[]InlineResponse200278Details, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200278Items) SetDetails(v []InlineResponse200278Details)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200278Items) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200278Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200278Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200278Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200278Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTs

`func (o *InlineResponse200278Items) GetStartTs() string`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200278Items) GetStartTsOk() (*string, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200278Items) SetStartTs(v string)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200278Items) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200278Items) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200278Items) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200278Items) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200278Items) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200278Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200278Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200278Items) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200278Items) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *InlineResponse200278Items) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InlineResponse200278Items) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InlineResponse200278Items) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *InlineResponse200278Items) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetDestination

`func (o *InlineResponse200278Items) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse200278Items) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse200278Items) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *InlineResponse200278Items) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetProcess

`func (o *InlineResponse200278Items) GetProcess() string`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *InlineResponse200278Items) GetProcessOk() (*string, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *InlineResponse200278Items) SetProcess(v string)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *InlineResponse200278Items) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetFile

`func (o *InlineResponse200278Items) GetFile() InlineResponse200278File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *InlineResponse200278Items) GetFileOk() (*InlineResponse200278File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *InlineResponse200278Items) SetFile(v InlineResponse200278File)`

SetFile sets File field to given value.

### HasFile

`func (o *InlineResponse200278Items) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetDuration

`func (o *InlineResponse200278Items) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineResponse200278Items) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineResponse200278Items) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InlineResponse200278Items) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterExpression

`func (o *InlineResponse200278Items) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *InlineResponse200278Items) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *InlineResponse200278Items) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *InlineResponse200278Items) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200278Items) GetCounts() InlineResponse200278Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200278Items) GetCountsOk() (*InlineResponse200278Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200278Items) SetCounts(v InlineResponse200278Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200278Items) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetInterface

`func (o *InlineResponse200278Items) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse200278Items) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse200278Items) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse200278Items) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


