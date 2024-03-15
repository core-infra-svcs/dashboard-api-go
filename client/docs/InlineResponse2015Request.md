# InlineResponse2015Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial number | [optional] 
**VlanId** | Pointer to **int32** | The target&#39;s VLAN (1 to 4094) | [optional] 
**Mac** | Pointer to **string** | The target&#39;s MAC address | [optional] 

## Methods

### NewInlineResponse2015Request

`func NewInlineResponse2015Request() *InlineResponse2015Request`

NewInlineResponse2015Request instantiates a new InlineResponse2015Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2015RequestWithDefaults

`func NewInlineResponse2015RequestWithDefaults() *InlineResponse2015Request`

NewInlineResponse2015RequestWithDefaults instantiates a new InlineResponse2015Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse2015Request) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse2015Request) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse2015Request) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse2015Request) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineResponse2015Request) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineResponse2015Request) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineResponse2015Request) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InlineResponse2015Request) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse2015Request) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse2015Request) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse2015Request) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse2015Request) HasMac() bool`

HasMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


