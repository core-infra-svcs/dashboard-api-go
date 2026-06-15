# InlineResponse200103

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200103Network**](InlineResponse200103Network.md) |  | [optional] 
**Policies** | Pointer to [**[]InlineResponse200103Policies**](InlineResponse200103Policies.md) | Umbrella policies applied to the network. The maximum number of policies is determined by Cisco Umbrella account limits. | [optional] 

## Methods

### NewInlineResponse200103

`func NewInlineResponse200103() *InlineResponse200103`

NewInlineResponse200103 instantiates a new InlineResponse200103 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200103WithDefaults

`func NewInlineResponse200103WithDefaults() *InlineResponse200103`

NewInlineResponse200103WithDefaults instantiates a new InlineResponse200103 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200103) GetNetwork() InlineResponse200103Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200103) GetNetworkOk() (*InlineResponse200103Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200103) SetNetwork(v InlineResponse200103Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200103) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse200103) GetPolicies() []InlineResponse200103Policies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse200103) GetPoliciesOk() (*[]InlineResponse200103Policies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse200103) SetPolicies(v []InlineResponse200103Policies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse200103) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


