# InlineResponse200125

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Meraki Id of the device record. | [optional] 
**Tags** | Pointer to **[]string** | An array of tags associated with the device. | [optional] 
**WifiMac** | Pointer to **string** | The MAC of the device. | [optional] 
**Serial** | Pointer to **string** | The device serial. | [optional] 

## Methods

### NewInlineResponse200125

`func NewInlineResponse200125() *InlineResponse200125`

NewInlineResponse200125 instantiates a new InlineResponse200125 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200125WithDefaults

`func NewInlineResponse200125WithDefaults() *InlineResponse200125`

NewInlineResponse200125WithDefaults instantiates a new InlineResponse200125 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200125) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200125) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200125) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200125) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200125) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200125) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200125) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200125) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWifiMac

`func (o *InlineResponse200125) GetWifiMac() string`

GetWifiMac returns the WifiMac field if non-nil, zero value otherwise.

### GetWifiMacOk

`func (o *InlineResponse200125) GetWifiMacOk() (*string, bool)`

GetWifiMacOk returns a tuple with the WifiMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMac

`func (o *InlineResponse200125) SetWifiMac(v string)`

SetWifiMac sets WifiMac field to given value.

### HasWifiMac

`func (o *InlineResponse200125) HasWifiMac() bool`

HasWifiMac returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200125) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200125) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200125) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200125) HasSerial() bool`

HasSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


