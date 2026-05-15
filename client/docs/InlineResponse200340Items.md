# InlineResponse200340Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Name** | Pointer to **string** | Name of the device | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device | [optional] 
**Provisioned** | Pointer to **int32** | The total RAM size provisioned on the device, in kB | [optional] 
**Used** | Pointer to [**InlineResponse200340Used**](InlineResponse200340Used.md) |  | [optional] 
**Free** | Pointer to [**InlineResponse200340Free**](InlineResponse200340Free.md) |  | [optional] 
**Network** | Pointer to [**InlineResponse200340Network**](InlineResponse200340Network.md) |  | [optional] 
**Intervals** | Pointer to [**[]InlineResponse200340Intervals**](InlineResponse200340Intervals.md) | Time interval snapshots of system memory utilization on the device with the most recent snapshot first | [optional] 

## Methods

### NewInlineResponse200340Items

`func NewInlineResponse200340Items() *InlineResponse200340Items`

NewInlineResponse200340Items instantiates a new InlineResponse200340Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200340ItemsWithDefaults

`func NewInlineResponse200340ItemsWithDefaults() *InlineResponse200340Items`

NewInlineResponse200340ItemsWithDefaults instantiates a new InlineResponse200340Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200340Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200340Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200340Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200340Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200340Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200340Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200340Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200340Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200340Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200340Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200340Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200340Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200340Items) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200340Items) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200340Items) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200340Items) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200340Items) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200340Items) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200340Items) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200340Items) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProvisioned

`func (o *InlineResponse200340Items) GetProvisioned() int32`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *InlineResponse200340Items) GetProvisionedOk() (*int32, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *InlineResponse200340Items) SetProvisioned(v int32)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *InlineResponse200340Items) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetUsed

`func (o *InlineResponse200340Items) GetUsed() InlineResponse200340Used`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *InlineResponse200340Items) GetUsedOk() (*InlineResponse200340Used, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *InlineResponse200340Items) SetUsed(v InlineResponse200340Used)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *InlineResponse200340Items) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetFree

`func (o *InlineResponse200340Items) GetFree() InlineResponse200340Free`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *InlineResponse200340Items) GetFreeOk() (*InlineResponse200340Free, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *InlineResponse200340Items) SetFree(v InlineResponse200340Free)`

SetFree sets Free field to given value.

### HasFree

`func (o *InlineResponse200340Items) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200340Items) GetNetwork() InlineResponse200340Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200340Items) GetNetworkOk() (*InlineResponse200340Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200340Items) SetNetwork(v InlineResponse200340Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200340Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIntervals

`func (o *InlineResponse200340Items) GetIntervals() []InlineResponse200340Intervals`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *InlineResponse200340Items) GetIntervalsOk() (*[]InlineResponse200340Intervals, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *InlineResponse200340Items) SetIntervals(v []InlineResponse200340Intervals)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *InlineResponse200340Items) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


