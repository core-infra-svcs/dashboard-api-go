# InlineResponse200319Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | ID of the primary cluster | [optional] 
**Name** | Pointer to **string** | Name of the primary cluster | [optional] 
**Network** | Pointer to [**InlineResponse200319Network**](InlineResponse200319Network.md) |  | [optional] 
**Model** | Pointer to **string** | Device model of the primary cluster | [optional] 
**Available** | Pointer to [**[]InlineResponse200319Available**](InlineResponse200319Available.md) | Available backup cluster targets for this primary cluster | [optional] 

## Methods

### NewInlineResponse200319Items

`func NewInlineResponse200319Items() *InlineResponse200319Items`

NewInlineResponse200319Items instantiates a new InlineResponse200319Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200319ItemsWithDefaults

`func NewInlineResponse200319ItemsWithDefaults() *InlineResponse200319Items`

NewInlineResponse200319ItemsWithDefaults instantiates a new InlineResponse200319Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *InlineResponse200319Items) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineResponse200319Items) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineResponse200319Items) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *InlineResponse200319Items) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200319Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200319Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200319Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200319Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200319Items) GetNetwork() InlineResponse200319Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200319Items) GetNetworkOk() (*InlineResponse200319Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200319Items) SetNetwork(v InlineResponse200319Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200319Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200319Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200319Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200319Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200319Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetAvailable

`func (o *InlineResponse200319Items) GetAvailable() []InlineResponse200319Available`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *InlineResponse200319Items) GetAvailableOk() (*[]InlineResponse200319Available, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *InlineResponse200319Items) SetAvailable(v []InlineResponse200319Available)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *InlineResponse200319Items) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


