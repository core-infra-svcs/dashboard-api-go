# InlineResponse200384Readings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The type of packets being counted | [optional] 
**Total** | Pointer to **int32** | The total count of sent and received packets during the timespan | [optional] 
**Recv** | Pointer to **int32** | The total count of packets received by the interface during the timespan | [optional] 
**Send** | Pointer to **int32** | The total count of packets sent by the interface during the timespan | [optional] 
**Rate** | Pointer to [**InlineResponse200384Rate**](InlineResponse200384Rate.md) |  | [optional] 

## Methods

### NewInlineResponse200384Readings

`func NewInlineResponse200384Readings() *InlineResponse200384Readings`

NewInlineResponse200384Readings instantiates a new InlineResponse200384Readings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200384ReadingsWithDefaults

`func NewInlineResponse200384ReadingsWithDefaults() *InlineResponse200384Readings`

NewInlineResponse200384ReadingsWithDefaults instantiates a new InlineResponse200384Readings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200384Readings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200384Readings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200384Readings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200384Readings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse200384Readings) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200384Readings) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200384Readings) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200384Readings) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRecv

`func (o *InlineResponse200384Readings) GetRecv() int32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *InlineResponse200384Readings) GetRecvOk() (*int32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *InlineResponse200384Readings) SetRecv(v int32)`

SetRecv sets Recv field to given value.

### HasRecv

`func (o *InlineResponse200384Readings) HasRecv() bool`

HasRecv returns a boolean if a field has been set.

### GetSend

`func (o *InlineResponse200384Readings) GetSend() int32`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *InlineResponse200384Readings) GetSendOk() (*int32, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *InlineResponse200384Readings) SetSend(v int32)`

SetSend sets Send field to given value.

### HasSend

`func (o *InlineResponse200384Readings) HasSend() bool`

HasSend returns a boolean if a field has been set.

### GetRate

`func (o *InlineResponse200384Readings) GetRate() InlineResponse200384Rate`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *InlineResponse200384Readings) GetRateOk() (*InlineResponse200384Rate, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *InlineResponse200384Readings) SetRate(v InlineResponse200384Rate)`

SetRate sets Rate field to given value.

### HasRate

`func (o *InlineResponse200384Readings) HasRate() bool`

HasRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


