# InlineResponse200334Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | ID of the primary cluster | [optional] 
**Name** | Pointer to **string** | Name of the primary cluster | [optional] 
**Network** | Pointer to [**InlineResponse200334Network**](InlineResponse200334Network.md) |  | [optional] 
**Model** | Pointer to **string** | Device model of the primary cluster | [optional] 
**Available** | Pointer to [**[]InlineResponse200334Available**](InlineResponse200334Available.md) | Available backup cluster targets for this primary cluster | [optional] 

## Methods

### NewInlineResponse200334Items

`func NewInlineResponse200334Items() *InlineResponse200334Items`

NewInlineResponse200334Items instantiates a new InlineResponse200334Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200334ItemsWithDefaults

`func NewInlineResponse200334ItemsWithDefaults() *InlineResponse200334Items`

NewInlineResponse200334ItemsWithDefaults instantiates a new InlineResponse200334Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *InlineResponse200334Items) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineResponse200334Items) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineResponse200334Items) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *InlineResponse200334Items) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200334Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200334Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200334Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200334Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200334Items) GetNetwork() InlineResponse200334Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200334Items) GetNetworkOk() (*InlineResponse200334Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200334Items) SetNetwork(v InlineResponse200334Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200334Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200334Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200334Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200334Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200334Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetAvailable

`func (o *InlineResponse200334Items) GetAvailable() []InlineResponse200334Available`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *InlineResponse200334Items) GetAvailableOk() (*[]InlineResponse200334Available, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *InlineResponse200334Items) SetAvailable(v []InlineResponse200334Available)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *InlineResponse200334Items) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


