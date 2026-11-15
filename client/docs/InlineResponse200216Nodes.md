# InlineResponse200216Nodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DerivedId** | Pointer to **string** | The derived identifier of the node | [optional] 
**Mac** | Pointer to **string** | The MAC address of the node | [optional] 
**Type** | Pointer to **string** | The type of the node. Can be &#39;device&#39;, &#39;discovered&#39;, &#39;stack&#39;, or &#39;unknown&#39; | [optional] 
**Root** | Pointer to **bool** | Whether the node is the root of the topology | [optional] 
**Device** | Pointer to [**InlineResponse200216Device**](InlineResponse200216Device.md) |  | [optional] 
**Discovered** | Pointer to [**InlineResponse200216Discovered**](InlineResponse200216Discovered.md) |  | [optional] 
**Stack** | Pointer to [**InlineResponse200216Stack**](InlineResponse200216Stack.md) |  | [optional] 

## Methods

### NewInlineResponse200216Nodes

`func NewInlineResponse200216Nodes() *InlineResponse200216Nodes`

NewInlineResponse200216Nodes instantiates a new InlineResponse200216Nodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200216NodesWithDefaults

`func NewInlineResponse200216NodesWithDefaults() *InlineResponse200216Nodes`

NewInlineResponse200216NodesWithDefaults instantiates a new InlineResponse200216Nodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDerivedId

`func (o *InlineResponse200216Nodes) GetDerivedId() string`

GetDerivedId returns the DerivedId field if non-nil, zero value otherwise.

### GetDerivedIdOk

`func (o *InlineResponse200216Nodes) GetDerivedIdOk() (*string, bool)`

GetDerivedIdOk returns a tuple with the DerivedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedId

`func (o *InlineResponse200216Nodes) SetDerivedId(v string)`

SetDerivedId sets DerivedId field to given value.

### HasDerivedId

`func (o *InlineResponse200216Nodes) HasDerivedId() bool`

HasDerivedId returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200216Nodes) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200216Nodes) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200216Nodes) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200216Nodes) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200216Nodes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200216Nodes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200216Nodes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200216Nodes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoot

`func (o *InlineResponse200216Nodes) GetRoot() bool`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *InlineResponse200216Nodes) GetRootOk() (*bool, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *InlineResponse200216Nodes) SetRoot(v bool)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *InlineResponse200216Nodes) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200216Nodes) GetDevice() InlineResponse200216Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200216Nodes) GetDeviceOk() (*InlineResponse200216Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200216Nodes) SetDevice(v InlineResponse200216Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200216Nodes) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDiscovered

`func (o *InlineResponse200216Nodes) GetDiscovered() InlineResponse200216Discovered`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *InlineResponse200216Nodes) GetDiscoveredOk() (*InlineResponse200216Discovered, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *InlineResponse200216Nodes) SetDiscovered(v InlineResponse200216Discovered)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *InlineResponse200216Nodes) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetStack

`func (o *InlineResponse200216Nodes) GetStack() InlineResponse200216Stack`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *InlineResponse200216Nodes) GetStackOk() (*InlineResponse200216Stack, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *InlineResponse200216Nodes) SetStack(v InlineResponse200216Stack)`

SetStack sets Stack field to given value.

### HasStack

`func (o *InlineResponse200216Nodes) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


