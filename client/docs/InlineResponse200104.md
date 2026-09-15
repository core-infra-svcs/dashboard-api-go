# InlineResponse200104

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200104Network**](InlineResponse200104Network.md) |  | [optional] 
**Policies** | Pointer to [**[]InlineResponse200104Policies**](InlineResponse200104Policies.md) | Umbrella policies applied to the network. The maximum number of policies is determined by Cisco Umbrella account limits. | [optional] 

## Methods

### NewInlineResponse200104

`func NewInlineResponse200104() *InlineResponse200104`

NewInlineResponse200104 instantiates a new InlineResponse200104 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200104WithDefaults

`func NewInlineResponse200104WithDefaults() *InlineResponse200104`

NewInlineResponse200104WithDefaults instantiates a new InlineResponse200104 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200104) GetNetwork() InlineResponse200104Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200104) GetNetworkOk() (*InlineResponse200104Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200104) SetNetwork(v InlineResponse200104Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200104) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse200104) GetPolicies() []InlineResponse200104Policies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse200104) GetPoliciesOk() (*[]InlineResponse200104Policies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse200104) SetPolicies(v []InlineResponse200104Policies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse200104) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


