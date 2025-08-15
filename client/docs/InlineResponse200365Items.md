# InlineResponse200365Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Name** | Pointer to **string** | Name of the device | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device | [optional] 
**Network** | Pointer to [**InlineResponse200284Network**](InlineResponse200284Network.md) |  | [optional] 
**CpuCount** | Pointer to **int32** | Number of CPU cores on the device | [optional] 
**Series** | Pointer to [**[]InlineResponse200365Series**](InlineResponse200365Series.md) | Series of cpu load average measurements on the device | [optional] 

## Methods

### NewInlineResponse200365Items

`func NewInlineResponse200365Items() *InlineResponse200365Items`

NewInlineResponse200365Items instantiates a new InlineResponse200365Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200365ItemsWithDefaults

`func NewInlineResponse200365ItemsWithDefaults() *InlineResponse200365Items`

NewInlineResponse200365ItemsWithDefaults instantiates a new InlineResponse200365Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200365Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200365Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200365Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200365Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200365Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200365Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200365Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200365Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200365Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200365Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200365Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200365Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200365Items) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200365Items) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200365Items) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200365Items) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200365Items) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200365Items) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200365Items) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200365Items) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200365Items) GetNetwork() InlineResponse200284Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200365Items) GetNetworkOk() (*InlineResponse200284Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200365Items) SetNetwork(v InlineResponse200284Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200365Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCpuCount

`func (o *InlineResponse200365Items) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *InlineResponse200365Items) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *InlineResponse200365Items) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *InlineResponse200365Items) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetSeries

`func (o *InlineResponse200365Items) GetSeries() []InlineResponse200365Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *InlineResponse200365Items) GetSeriesOk() (*[]InlineResponse200365Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *InlineResponse200365Items) SetSeries(v []InlineResponse200365Series)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *InlineResponse200365Items) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


