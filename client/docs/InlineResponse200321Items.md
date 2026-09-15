# InlineResponse200321Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | ID of the primary cluster | [optional] 
**Name** | Pointer to **string** | Name of the primary cluster | [optional] 
**Network** | Pointer to [**InlineResponse200321Network**](InlineResponse200321Network.md) |  | [optional] 
**Model** | Pointer to **string** | Device model of the primary cluster | [optional] 
**Available** | Pointer to [**[]InlineResponse200321Available**](InlineResponse200321Available.md) | Available backup cluster targets for this primary cluster | [optional] 

## Methods

### NewInlineResponse200321Items

`func NewInlineResponse200321Items() *InlineResponse200321Items`

NewInlineResponse200321Items instantiates a new InlineResponse200321Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200321ItemsWithDefaults

`func NewInlineResponse200321ItemsWithDefaults() *InlineResponse200321Items`

NewInlineResponse200321ItemsWithDefaults instantiates a new InlineResponse200321Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *InlineResponse200321Items) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineResponse200321Items) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineResponse200321Items) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *InlineResponse200321Items) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200321Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200321Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200321Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200321Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200321Items) GetNetwork() InlineResponse200321Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200321Items) GetNetworkOk() (*InlineResponse200321Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200321Items) SetNetwork(v InlineResponse200321Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200321Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200321Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200321Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200321Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200321Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetAvailable

`func (o *InlineResponse200321Items) GetAvailable() []InlineResponse200321Available`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *InlineResponse200321Items) GetAvailableOk() (*[]InlineResponse200321Available, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *InlineResponse200321Items) SetAvailable(v []InlineResponse200321Available)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *InlineResponse200321Items) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


