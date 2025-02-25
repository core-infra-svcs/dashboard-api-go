# InlineResponse200350Readings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Mac** | Pointer to **string** | The MAC address of the wireless controller interface | [optional] 
**Recv** | Pointer to **int32** | The volume of data, in bytes/sec, received by wireless controller interface | [optional] 
**Send** | Pointer to **int32** | The volume of data, in bytes/sec, transmitted by wireless controller interface | [optional] 

## Methods

### NewInlineResponse200350Readings

`func NewInlineResponse200350Readings() *InlineResponse200350Readings`

NewInlineResponse200350Readings instantiates a new InlineResponse200350Readings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200350ReadingsWithDefaults

`func NewInlineResponse200350ReadingsWithDefaults() *InlineResponse200350Readings`

NewInlineResponse200350ReadingsWithDefaults instantiates a new InlineResponse200350Readings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200350Readings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200350Readings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200350Readings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200350Readings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200350Readings) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200350Readings) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200350Readings) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200350Readings) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetRecv

`func (o *InlineResponse200350Readings) GetRecv() int32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *InlineResponse200350Readings) GetRecvOk() (*int32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *InlineResponse200350Readings) SetRecv(v int32)`

SetRecv sets Recv field to given value.

### HasRecv

`func (o *InlineResponse200350Readings) HasRecv() bool`

HasRecv returns a boolean if a field has been set.

### GetSend

`func (o *InlineResponse200350Readings) GetSend() int32`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *InlineResponse200350Readings) GetSendOk() (*int32, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *InlineResponse200350Readings) SetSend(v int32)`

SetSend sets Send field to given value.

### HasSend

`func (o *InlineResponse200350Readings) HasSend() bool`

HasSend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


