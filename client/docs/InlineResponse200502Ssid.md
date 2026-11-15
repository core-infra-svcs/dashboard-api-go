# InlineResponse200502Ssid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the SSID | 
**Number** | **int32** | The number of the SSID | 
**Name** | **string** | The name of the SSID | 
**TunneledTo** | [**InlineResponse200502SsidTunneledTo**](InlineResponse200502SsidTunneledTo.md) |  | 

## Methods

### NewInlineResponse200502Ssid

`func NewInlineResponse200502Ssid(id string, number int32, name string, tunneledTo InlineResponse200502SsidTunneledTo, ) *InlineResponse200502Ssid`

NewInlineResponse200502Ssid instantiates a new InlineResponse200502Ssid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200502SsidWithDefaults

`func NewInlineResponse200502SsidWithDefaults() *InlineResponse200502Ssid`

NewInlineResponse200502SsidWithDefaults instantiates a new InlineResponse200502Ssid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200502Ssid) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200502Ssid) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200502Ssid) SetId(v string)`

SetId sets Id field to given value.


### GetNumber

`func (o *InlineResponse200502Ssid) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200502Ssid) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200502Ssid) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetName

`func (o *InlineResponse200502Ssid) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200502Ssid) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200502Ssid) SetName(v string)`

SetName sets Name field to given value.


### GetTunneledTo

`func (o *InlineResponse200502Ssid) GetTunneledTo() InlineResponse200502SsidTunneledTo`

GetTunneledTo returns the TunneledTo field if non-nil, zero value otherwise.

### GetTunneledToOk

`func (o *InlineResponse200502Ssid) GetTunneledToOk() (*InlineResponse200502SsidTunneledTo, bool)`

GetTunneledToOk returns a tuple with the TunneledTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledTo

`func (o *InlineResponse200502Ssid) SetTunneledTo(v InlineResponse200502SsidTunneledTo)`

SetTunneledTo sets TunneledTo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


