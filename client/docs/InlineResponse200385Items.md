# InlineResponse200385Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Name** | Pointer to **string** | Name of the device | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device | [optional] 
**Provisioned** | Pointer to **int32** | The total RAM size provisioned on the device, in kB | [optional] 
**Used** | Pointer to [**InlineResponse200385Used**](InlineResponse200385Used.md) |  | [optional] 
**Free** | Pointer to [**InlineResponse200385Free**](InlineResponse200385Free.md) |  | [optional] 
**Network** | Pointer to [**InlineResponse200385Network**](InlineResponse200385Network.md) |  | [optional] 
**Intervals** | Pointer to [**[]InlineResponse200385Intervals**](InlineResponse200385Intervals.md) | Time interval snapshots of system memory utilization on the device with the most recent snapshot first | [optional] 

## Methods

### NewInlineResponse200385Items

`func NewInlineResponse200385Items() *InlineResponse200385Items`

NewInlineResponse200385Items instantiates a new InlineResponse200385Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200385ItemsWithDefaults

`func NewInlineResponse200385ItemsWithDefaults() *InlineResponse200385Items`

NewInlineResponse200385ItemsWithDefaults instantiates a new InlineResponse200385Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200385Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200385Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200385Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200385Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200385Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200385Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200385Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200385Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200385Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200385Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200385Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200385Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200385Items) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200385Items) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200385Items) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200385Items) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200385Items) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200385Items) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200385Items) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200385Items) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProvisioned

`func (o *InlineResponse200385Items) GetProvisioned() int32`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *InlineResponse200385Items) GetProvisionedOk() (*int32, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *InlineResponse200385Items) SetProvisioned(v int32)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *InlineResponse200385Items) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetUsed

`func (o *InlineResponse200385Items) GetUsed() InlineResponse200385Used`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *InlineResponse200385Items) GetUsedOk() (*InlineResponse200385Used, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *InlineResponse200385Items) SetUsed(v InlineResponse200385Used)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *InlineResponse200385Items) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetFree

`func (o *InlineResponse200385Items) GetFree() InlineResponse200385Free`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *InlineResponse200385Items) GetFreeOk() (*InlineResponse200385Free, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *InlineResponse200385Items) SetFree(v InlineResponse200385Free)`

SetFree sets Free field to given value.

### HasFree

`func (o *InlineResponse200385Items) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200385Items) GetNetwork() InlineResponse200385Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200385Items) GetNetworkOk() (*InlineResponse200385Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200385Items) SetNetwork(v InlineResponse200385Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200385Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIntervals

`func (o *InlineResponse200385Items) GetIntervals() []InlineResponse200385Intervals`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *InlineResponse200385Items) GetIntervalsOk() (*[]InlineResponse200385Intervals, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *InlineResponse200385Items) SetIntervals(v []InlineResponse200385Intervals)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *InlineResponse200385Items) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


