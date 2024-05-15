# InlineResponse200100

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial | [optional] 
**Model** | Pointer to **string** | Device model. | [optional] 
**Tags** | Pointer to **string** | Device tags. | [optional] 
**Wifi0** | Pointer to [**[]NetworksNetworkIdNetworkHealthChannelUtilizationWifi0**](NetworksNetworkIdNetworkHealthChannelUtilizationWifi0.md) | Channel utilization for first wifi radio of device. | [optional] 
**Wifi1** | Pointer to [**[]NetworksNetworkIdNetworkHealthChannelUtilizationWifi0**](NetworksNetworkIdNetworkHealthChannelUtilizationWifi0.md) | Channel utilization for second wifi radio of device. | [optional] 

## Methods

### NewInlineResponse200100

`func NewInlineResponse200100() *InlineResponse200100`

NewInlineResponse200100 instantiates a new InlineResponse200100 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200100WithDefaults

`func NewInlineResponse200100WithDefaults() *InlineResponse200100`

NewInlineResponse200100WithDefaults instantiates a new InlineResponse200100 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200100) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200100) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200100) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200100) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200100) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200100) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200100) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200100) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200100) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200100) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200100) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200100) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWifi0

`func (o *InlineResponse200100) GetWifi0() []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0`

GetWifi0 returns the Wifi0 field if non-nil, zero value otherwise.

### GetWifi0Ok

`func (o *InlineResponse200100) GetWifi0Ok() (*[]NetworksNetworkIdNetworkHealthChannelUtilizationWifi0, bool)`

GetWifi0Ok returns a tuple with the Wifi0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifi0

`func (o *InlineResponse200100) SetWifi0(v []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0)`

SetWifi0 sets Wifi0 field to given value.

### HasWifi0

`func (o *InlineResponse200100) HasWifi0() bool`

HasWifi0 returns a boolean if a field has been set.

### GetWifi1

`func (o *InlineResponse200100) GetWifi1() []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0`

GetWifi1 returns the Wifi1 field if non-nil, zero value otherwise.

### GetWifi1Ok

`func (o *InlineResponse200100) GetWifi1Ok() (*[]NetworksNetworkIdNetworkHealthChannelUtilizationWifi0, bool)`

GetWifi1Ok returns a tuple with the Wifi1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifi1

`func (o *InlineResponse200100) SetWifi1(v []NetworksNetworkIdNetworkHealthChannelUtilizationWifi0)`

SetWifi1 sets Wifi1 field to given value.

### HasWifi1

`func (o *InlineResponse200100) HasWifi1() bool`

HasWifi1 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


