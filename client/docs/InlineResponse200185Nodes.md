# InlineResponse200185Nodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DerivedId** | Pointer to **string** | The derived identifier of the node | [optional] 
**Mac** | Pointer to **string** | The MAC address of the node | [optional] 
**Type** | Pointer to **string** | The type of the node. Can be &#39;device&#39;, &#39;discovered&#39;, &#39;stack&#39;, or &#39;unknown&#39; | [optional] 
**Root** | Pointer to **bool** | Whether the node is the root of the topology | [optional] 
**Device** | Pointer to [**InlineResponse200185Device**](InlineResponse200185Device.md) |  | [optional] 
**Discovered** | Pointer to [**InlineResponse200185Discovered**](InlineResponse200185Discovered.md) |  | [optional] 
**Stack** | Pointer to [**InlineResponse200185Stack**](InlineResponse200185Stack.md) |  | [optional] 

## Methods

### NewInlineResponse200185Nodes

`func NewInlineResponse200185Nodes() *InlineResponse200185Nodes`

NewInlineResponse200185Nodes instantiates a new InlineResponse200185Nodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200185NodesWithDefaults

`func NewInlineResponse200185NodesWithDefaults() *InlineResponse200185Nodes`

NewInlineResponse200185NodesWithDefaults instantiates a new InlineResponse200185Nodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDerivedId

`func (o *InlineResponse200185Nodes) GetDerivedId() string`

GetDerivedId returns the DerivedId field if non-nil, zero value otherwise.

### GetDerivedIdOk

`func (o *InlineResponse200185Nodes) GetDerivedIdOk() (*string, bool)`

GetDerivedIdOk returns a tuple with the DerivedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedId

`func (o *InlineResponse200185Nodes) SetDerivedId(v string)`

SetDerivedId sets DerivedId field to given value.

### HasDerivedId

`func (o *InlineResponse200185Nodes) HasDerivedId() bool`

HasDerivedId returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200185Nodes) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200185Nodes) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200185Nodes) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200185Nodes) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200185Nodes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200185Nodes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200185Nodes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200185Nodes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoot

`func (o *InlineResponse200185Nodes) GetRoot() bool`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *InlineResponse200185Nodes) GetRootOk() (*bool, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *InlineResponse200185Nodes) SetRoot(v bool)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *InlineResponse200185Nodes) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200185Nodes) GetDevice() InlineResponse200185Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200185Nodes) GetDeviceOk() (*InlineResponse200185Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200185Nodes) SetDevice(v InlineResponse200185Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200185Nodes) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDiscovered

`func (o *InlineResponse200185Nodes) GetDiscovered() InlineResponse200185Discovered`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *InlineResponse200185Nodes) GetDiscoveredOk() (*InlineResponse200185Discovered, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *InlineResponse200185Nodes) SetDiscovered(v InlineResponse200185Discovered)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *InlineResponse200185Nodes) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetStack

`func (o *InlineResponse200185Nodes) GetStack() InlineResponse200185Stack`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *InlineResponse200185Nodes) GetStackOk() (*InlineResponse200185Stack, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *InlineResponse200185Nodes) SetStack(v InlineResponse200185Stack)`

SetStack sets Stack field to given value.

### HasStack

`func (o *InlineResponse200185Nodes) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


