# InlineObject33

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanId** | **int32** | The target&#39;s VLAN (1 to 4094) | 
**Mac** | **string** | The target&#39;s MAC address | 
**Callback** | Pointer to [**DevicesSerialLiveToolsArpTableCallback**](DevicesSerialLiveToolsArpTableCallback.md) |  | [optional] 

## Methods

### NewInlineObject33

`func NewInlineObject33(vlanId int32, mac string, ) *InlineObject33`

NewInlineObject33 instantiates a new InlineObject33 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject33WithDefaults

`func NewInlineObject33WithDefaults() *InlineObject33`

NewInlineObject33WithDefaults instantiates a new InlineObject33 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanId

`func (o *InlineObject33) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineObject33) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineObject33) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### GetMac

`func (o *InlineObject33) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineObject33) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineObject33) SetMac(v string)`

SetMac sets Mac field to given value.


### GetCallback

`func (o *InlineObject33) GetCallback() DevicesSerialLiveToolsArpTableCallback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineObject33) GetCallbackOk() (*DevicesSerialLiveToolsArpTableCallback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineObject33) SetCallback(v DevicesSerialLiveToolsArpTableCallback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineObject33) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


