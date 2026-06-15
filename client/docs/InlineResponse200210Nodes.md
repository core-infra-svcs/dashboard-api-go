# InlineResponse200210Nodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DerivedId** | Pointer to **string** | The derived identifier of the node | [optional] 
**Mac** | Pointer to **string** | The MAC address of the node | [optional] 
**Type** | Pointer to **string** | The type of the node. Can be &#39;device&#39;, &#39;discovered&#39;, &#39;stack&#39;, or &#39;unknown&#39; | [optional] 
**Root** | Pointer to **bool** | Whether the node is the root of the topology | [optional] 
**Device** | Pointer to [**InlineResponse200210Device**](InlineResponse200210Device.md) |  | [optional] 
**Discovered** | Pointer to [**InlineResponse200210Discovered**](InlineResponse200210Discovered.md) |  | [optional] 
**Stack** | Pointer to [**InlineResponse200210Stack**](InlineResponse200210Stack.md) |  | [optional] 

## Methods

### NewInlineResponse200210Nodes

`func NewInlineResponse200210Nodes() *InlineResponse200210Nodes`

NewInlineResponse200210Nodes instantiates a new InlineResponse200210Nodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200210NodesWithDefaults

`func NewInlineResponse200210NodesWithDefaults() *InlineResponse200210Nodes`

NewInlineResponse200210NodesWithDefaults instantiates a new InlineResponse200210Nodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDerivedId

`func (o *InlineResponse200210Nodes) GetDerivedId() string`

GetDerivedId returns the DerivedId field if non-nil, zero value otherwise.

### GetDerivedIdOk

`func (o *InlineResponse200210Nodes) GetDerivedIdOk() (*string, bool)`

GetDerivedIdOk returns a tuple with the DerivedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedId

`func (o *InlineResponse200210Nodes) SetDerivedId(v string)`

SetDerivedId sets DerivedId field to given value.

### HasDerivedId

`func (o *InlineResponse200210Nodes) HasDerivedId() bool`

HasDerivedId returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200210Nodes) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200210Nodes) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200210Nodes) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200210Nodes) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200210Nodes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200210Nodes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200210Nodes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200210Nodes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoot

`func (o *InlineResponse200210Nodes) GetRoot() bool`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *InlineResponse200210Nodes) GetRootOk() (*bool, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *InlineResponse200210Nodes) SetRoot(v bool)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *InlineResponse200210Nodes) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200210Nodes) GetDevice() InlineResponse200210Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200210Nodes) GetDeviceOk() (*InlineResponse200210Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200210Nodes) SetDevice(v InlineResponse200210Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200210Nodes) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDiscovered

`func (o *InlineResponse200210Nodes) GetDiscovered() InlineResponse200210Discovered`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *InlineResponse200210Nodes) GetDiscoveredOk() (*InlineResponse200210Discovered, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *InlineResponse200210Nodes) SetDiscovered(v InlineResponse200210Discovered)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *InlineResponse200210Nodes) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetStack

`func (o *InlineResponse200210Nodes) GetStack() InlineResponse200210Stack`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *InlineResponse200210Nodes) GetStackOk() (*InlineResponse200210Stack, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *InlineResponse200210Nodes) SetStack(v InlineResponse200210Stack)`

SetStack sets Stack field to given value.

### HasStack

`func (o *InlineResponse200210Nodes) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


