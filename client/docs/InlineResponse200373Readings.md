# InlineResponse200373Readings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Mac** | Pointer to **string** | The MAC address of the wireless controller interface | [optional] 
**Recv** | Pointer to **int32** | The volume of data, in bytes/sec, received by wireless controller interface | [optional] 
**Send** | Pointer to **int32** | The volume of data, in bytes/sec, transmitted by wireless controller interface | [optional] 

## Methods

### NewInlineResponse200373Readings

`func NewInlineResponse200373Readings() *InlineResponse200373Readings`

NewInlineResponse200373Readings instantiates a new InlineResponse200373Readings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200373ReadingsWithDefaults

`func NewInlineResponse200373ReadingsWithDefaults() *InlineResponse200373Readings`

NewInlineResponse200373ReadingsWithDefaults instantiates a new InlineResponse200373Readings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200373Readings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200373Readings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200373Readings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200373Readings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200373Readings) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200373Readings) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200373Readings) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200373Readings) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetRecv

`func (o *InlineResponse200373Readings) GetRecv() int32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *InlineResponse200373Readings) GetRecvOk() (*int32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *InlineResponse200373Readings) SetRecv(v int32)`

SetRecv sets Recv field to given value.

### HasRecv

`func (o *InlineResponse200373Readings) HasRecv() bool`

HasRecv returns a boolean if a field has been set.

### GetSend

`func (o *InlineResponse200373Readings) GetSend() int32`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *InlineResponse200373Readings) GetSendOk() (*int32, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *InlineResponse200373Readings) SetSend(v int32)`

SetSend sets Send field to given value.

### HasSend

`func (o *InlineResponse200373Readings) HasSend() bool`

HasSend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


