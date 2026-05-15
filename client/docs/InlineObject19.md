# InlineObject19

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | **[]string** | A list of ports for which to perform the cable test.  For Catalyst switches, IOS interface names are also supported, such as \&quot;GigabitEthernet1/0/8\&quot;, \&quot;Gi1/0/8\&quot;, or even \&quot;1/0/8\&quot;. | 
**Callback** | Pointer to [**DevicesSerialLiveToolsArpTableCallback**](DevicesSerialLiveToolsArpTableCallback.md) |  | [optional] 

## Methods

### NewInlineObject19

`func NewInlineObject19(ports []string, ) *InlineObject19`

NewInlineObject19 instantiates a new InlineObject19 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject19WithDefaults

`func NewInlineObject19WithDefaults() *InlineObject19`

NewInlineObject19WithDefaults instantiates a new InlineObject19 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *InlineObject19) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineObject19) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineObject19) SetPorts(v []string)`

SetPorts sets Ports field to given value.


### GetCallback

`func (o *InlineObject19) GetCallback() DevicesSerialLiveToolsArpTableCallback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineObject19) GetCallbackOk() (*DevicesSerialLiveToolsArpTableCallback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineObject19) SetCallback(v DevicesSerialLiveToolsArpTableCallback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineObject19) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


