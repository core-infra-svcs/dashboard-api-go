# InlineResponse200442Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentId** | **string** | ID of the network enforcement target policy assignment | 
**Policy** | [**InlineResponse200442Policy**](InlineResponse200442Policy.md) |  | 
**Network** | [**InlineResponse200441Network**](InlineResponse200441Network.md) |  | 
**Firewall** | [**InlineResponse200441Firewall**](InlineResponse200441Firewall.md) |  | 

## Methods

### NewInlineResponse200442Items

`func NewInlineResponse200442Items(assignmentId string, policy InlineResponse200442Policy, network InlineResponse200441Network, firewall InlineResponse200441Firewall, ) *InlineResponse200442Items`

NewInlineResponse200442Items instantiates a new InlineResponse200442Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200442ItemsWithDefaults

`func NewInlineResponse200442ItemsWithDefaults() *InlineResponse200442Items`

NewInlineResponse200442ItemsWithDefaults instantiates a new InlineResponse200442Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentId

`func (o *InlineResponse200442Items) GetAssignmentId() string`

GetAssignmentId returns the AssignmentId field if non-nil, zero value otherwise.

### GetAssignmentIdOk

`func (o *InlineResponse200442Items) GetAssignmentIdOk() (*string, bool)`

GetAssignmentIdOk returns a tuple with the AssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentId

`func (o *InlineResponse200442Items) SetAssignmentId(v string)`

SetAssignmentId sets AssignmentId field to given value.


### GetPolicy

`func (o *InlineResponse200442Items) GetPolicy() InlineResponse200442Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse200442Items) GetPolicyOk() (*InlineResponse200442Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse200442Items) SetPolicy(v InlineResponse200442Policy)`

SetPolicy sets Policy field to given value.


### GetNetwork

`func (o *InlineResponse200442Items) GetNetwork() InlineResponse200441Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200442Items) GetNetworkOk() (*InlineResponse200441Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200442Items) SetNetwork(v InlineResponse200441Network)`

SetNetwork sets Network field to given value.


### GetFirewall

`func (o *InlineResponse200442Items) GetFirewall() InlineResponse200441Firewall`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *InlineResponse200442Items) GetFirewallOk() (*InlineResponse200441Firewall, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *InlineResponse200442Items) SetFirewall(v InlineResponse200441Firewall)`

SetFirewall sets Firewall field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


