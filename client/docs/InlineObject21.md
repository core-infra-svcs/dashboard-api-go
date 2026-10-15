# InlineObject21

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | **[]string** | A list of ports for which to perform the cable test.  For Catalyst switches, IOS interface names are also supported, such as \&quot;GigabitEthernet1/0/8\&quot;, \&quot;Gi1/0/8\&quot;, or even \&quot;1/0/8\&quot;. | 
**Callback** | Pointer to [**DevicesSerialLiveToolsArpTableCallback**](DevicesSerialLiveToolsArpTableCallback.md) |  | [optional] 

## Methods

### NewInlineObject21

`func NewInlineObject21(ports []string, ) *InlineObject21`

NewInlineObject21 instantiates a new InlineObject21 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject21WithDefaults

`func NewInlineObject21WithDefaults() *InlineObject21`

NewInlineObject21WithDefaults instantiates a new InlineObject21 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *InlineObject21) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineObject21) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineObject21) SetPorts(v []string)`

SetPorts sets Ports field to given value.


### GetCallback

`func (o *InlineObject21) GetCallback() DevicesSerialLiveToolsArpTableCallback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineObject21) GetCallbackOk() (*DevicesSerialLiveToolsArpTableCallback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineObject21) SetCallback(v DevicesSerialLiveToolsArpTableCallback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineObject21) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


