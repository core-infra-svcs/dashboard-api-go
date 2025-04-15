# InlineResponse200345Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Name** | Pointer to **string** | Name of the device | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device | [optional] 
**Network** | Pointer to [**InlineResponse200270Network**](InlineResponse200270Network.md) |  | [optional] 
**CpuCount** | Pointer to **int32** | Number of CPU cores on the device | [optional] 
**Series** | Pointer to [**[]InlineResponse200345Series**](InlineResponse200345Series.md) | Series of cpu load average measurements on the device | [optional] 

## Methods

### NewInlineResponse200345Items

`func NewInlineResponse200345Items() *InlineResponse200345Items`

NewInlineResponse200345Items instantiates a new InlineResponse200345Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200345ItemsWithDefaults

`func NewInlineResponse200345ItemsWithDefaults() *InlineResponse200345Items`

NewInlineResponse200345ItemsWithDefaults instantiates a new InlineResponse200345Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200345Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200345Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200345Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200345Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200345Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200345Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200345Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200345Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200345Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200345Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200345Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200345Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200345Items) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200345Items) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200345Items) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200345Items) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200345Items) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200345Items) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200345Items) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200345Items) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200345Items) GetNetwork() InlineResponse200270Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200345Items) GetNetworkOk() (*InlineResponse200270Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200345Items) SetNetwork(v InlineResponse200270Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200345Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCpuCount

`func (o *InlineResponse200345Items) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *InlineResponse200345Items) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *InlineResponse200345Items) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *InlineResponse200345Items) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetSeries

`func (o *InlineResponse200345Items) GetSeries() []InlineResponse200345Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *InlineResponse200345Items) GetSeriesOk() (*[]InlineResponse200345Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *InlineResponse200345Items) SetSeries(v []InlineResponse200345Series)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *InlineResponse200345Items) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


