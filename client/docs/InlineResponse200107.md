# InlineResponse200107

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200107Network**](InlineResponse200107Network.md) |  | [optional] 
**Policies** | Pointer to [**[]InlineResponse200107Policies**](InlineResponse200107Policies.md) | Umbrella policies applied to the network. The maximum number of policies is determined by Cisco Umbrella account limits. | [optional] 

## Methods

### NewInlineResponse200107

`func NewInlineResponse200107() *InlineResponse200107`

NewInlineResponse200107 instantiates a new InlineResponse200107 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200107WithDefaults

`func NewInlineResponse200107WithDefaults() *InlineResponse200107`

NewInlineResponse200107WithDefaults instantiates a new InlineResponse200107 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200107) GetNetwork() InlineResponse200107Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200107) GetNetworkOk() (*InlineResponse200107Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200107) SetNetwork(v InlineResponse200107Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200107) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse200107) GetPolicies() []InlineResponse200107Policies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse200107) GetPoliciesOk() (*[]InlineResponse200107Policies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse200107) SetPolicies(v []InlineResponse200107Policies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse200107) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


