# InlineResponse200430Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentId** | **string** | ID of the network enforcement target policy assignment | 
**Policy** | [**InlineResponse200430Policy**](InlineResponse200430Policy.md) |  | 
**Network** | [**InlineResponse200429Network**](InlineResponse200429Network.md) |  | 
**Firewall** | [**InlineResponse200429Firewall**](InlineResponse200429Firewall.md) |  | 

## Methods

### NewInlineResponse200430Items

`func NewInlineResponse200430Items(assignmentId string, policy InlineResponse200430Policy, network InlineResponse200429Network, firewall InlineResponse200429Firewall, ) *InlineResponse200430Items`

NewInlineResponse200430Items instantiates a new InlineResponse200430Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200430ItemsWithDefaults

`func NewInlineResponse200430ItemsWithDefaults() *InlineResponse200430Items`

NewInlineResponse200430ItemsWithDefaults instantiates a new InlineResponse200430Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentId

`func (o *InlineResponse200430Items) GetAssignmentId() string`

GetAssignmentId returns the AssignmentId field if non-nil, zero value otherwise.

### GetAssignmentIdOk

`func (o *InlineResponse200430Items) GetAssignmentIdOk() (*string, bool)`

GetAssignmentIdOk returns a tuple with the AssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentId

`func (o *InlineResponse200430Items) SetAssignmentId(v string)`

SetAssignmentId sets AssignmentId field to given value.


### GetPolicy

`func (o *InlineResponse200430Items) GetPolicy() InlineResponse200430Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200430Items) GetPolicyOk() (*InlineResponse200430Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200430Items) SetPolicy(v InlineResponse200430Policy)`

SetPolicy sets Policy field to given value.


### GetNetwork

`func (o *InlineResponse200430Items) GetNetwork() InlineResponse200429Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200430Items) GetNetworkOk() (*InlineResponse200429Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200430Items) SetNetwork(v InlineResponse200429Network)`

SetNetwork sets Network field to given value.


### GetFirewall

`func (o *InlineResponse200430Items) GetFirewall() InlineResponse200429Firewall`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *InlineResponse200430Items) GetFirewallOk() (*InlineResponse200429Firewall, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *InlineResponse200430Items) SetFirewall(v InlineResponse200429Firewall)`

SetFirewall sets Firewall field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


