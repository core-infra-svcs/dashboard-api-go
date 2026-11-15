# InlineResponse200327Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200327Network**](InlineResponse200327Network.md) |  | [optional] 
**Successes** | Pointer to [**InlineResponse200327Successes**](InlineResponse200327Successes.md) |  | [optional] 
**Failures** | Pointer to [**InlineResponse200327Failures**](InlineResponse200327Failures.md) |  | [optional] 
**Clients** | Pointer to [**InlineResponse200327Clients**](InlineResponse200327Clients.md) |  | [optional] 
**ByContributor** | Pointer to [**[]InlineResponse200327ByContributor**](InlineResponse200327ByContributor.md) | Connection failures grouped by contributor. Maximum 5 entries. No ordering is guaranteed. | [optional] 

## Methods

### NewInlineResponse200327Items

`func NewInlineResponse200327Items() *InlineResponse200327Items`

NewInlineResponse200327Items instantiates a new InlineResponse200327Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200327ItemsWithDefaults

`func NewInlineResponse200327ItemsWithDefaults() *InlineResponse200327Items`

NewInlineResponse200327ItemsWithDefaults instantiates a new InlineResponse200327Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200327Items) GetNetwork() InlineResponse200327Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200327Items) GetNetworkOk() (*InlineResponse200327Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200327Items) SetNetwork(v InlineResponse200327Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200327Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSuccesses

`func (o *InlineResponse200327Items) GetSuccesses() InlineResponse200327Successes`

GetSuccesses returns the Successes field if non-nil, zero value otherwise.

### GetSuccessesOk

`func (o *InlineResponse200327Items) GetSuccessesOk() (*InlineResponse200327Successes, bool)`

GetSuccessesOk returns a tuple with the Successes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccesses

`func (o *InlineResponse200327Items) SetSuccesses(v InlineResponse200327Successes)`

SetSuccesses sets Successes field to given value.

### HasSuccesses

`func (o *InlineResponse200327Items) HasSuccesses() bool`

HasSuccesses returns a boolean if a field has been set.

### GetFailures

`func (o *InlineResponse200327Items) GetFailures() InlineResponse200327Failures`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *InlineResponse200327Items) GetFailuresOk() (*InlineResponse200327Failures, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *InlineResponse200327Items) SetFailures(v InlineResponse200327Failures)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *InlineResponse200327Items) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200327Items) GetClients() InlineResponse200327Clients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200327Items) GetClientsOk() (*InlineResponse200327Clients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200327Items) SetClients(v InlineResponse200327Clients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200327Items) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetByContributor

`func (o *InlineResponse200327Items) GetByContributor() []InlineResponse200327ByContributor`

GetByContributor returns the ByContributor field if non-nil, zero value otherwise.

### GetByContributorOk

`func (o *InlineResponse200327Items) GetByContributorOk() (*[]InlineResponse200327ByContributor, bool)`

GetByContributorOk returns a tuple with the ByContributor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByContributor

`func (o *InlineResponse200327Items) SetByContributor(v []InlineResponse200327ByContributor)`

SetByContributor sets ByContributor field to given value.

### HasByContributor

`func (o *InlineResponse200327Items) HasByContributor() bool`

HasByContributor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


