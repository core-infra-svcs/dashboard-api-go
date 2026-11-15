# InlineResponse200327ByContributor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Failure contributor name. | [optional] 
**Successes** | Pointer to [**InlineResponse200327Successes1**](InlineResponse200327Successes1.md) |  | [optional] 
**Failures** | Pointer to [**InlineResponse200327Failures1**](InlineResponse200327Failures1.md) |  | [optional] 
**FailureContribution** | Pointer to [**InlineResponse200327FailureContribution**](InlineResponse200327FailureContribution.md) |  | [optional] 
**ClientFailureContribution** | Pointer to [**InlineResponse200327ClientFailureContribution**](InlineResponse200327ClientFailureContribution.md) |  | [optional] 
**Clients** | Pointer to [**InlineResponse200327Clients1**](InlineResponse200327Clients1.md) |  | [optional] 
**BySubContributor** | Pointer to [**[]InlineResponse200327BySubContributor**](InlineResponse200327BySubContributor.md) | Failure subcontributors. Maximum 10 entries. Contributor totals include entries not listed. No ordering is guaranteed. | [optional] 

## Methods

### NewInlineResponse200327ByContributor

`func NewInlineResponse200327ByContributor() *InlineResponse200327ByContributor`

NewInlineResponse200327ByContributor instantiates a new InlineResponse200327ByContributor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200327ByContributorWithDefaults

`func NewInlineResponse200327ByContributorWithDefaults() *InlineResponse200327ByContributor`

NewInlineResponse200327ByContributorWithDefaults instantiates a new InlineResponse200327ByContributor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200327ByContributor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200327ByContributor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200327ByContributor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200327ByContributor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSuccesses

`func (o *InlineResponse200327ByContributor) GetSuccesses() InlineResponse200327Successes1`

GetSuccesses returns the Successes field if non-nil, zero value otherwise.

### GetSuccessesOk

`func (o *InlineResponse200327ByContributor) GetSuccessesOk() (*InlineResponse200327Successes1, bool)`

GetSuccessesOk returns a tuple with the Successes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccesses

`func (o *InlineResponse200327ByContributor) SetSuccesses(v InlineResponse200327Successes1)`

SetSuccesses sets Successes field to given value.

### HasSuccesses

`func (o *InlineResponse200327ByContributor) HasSuccesses() bool`

HasSuccesses returns a boolean if a field has been set.

### GetFailures

`func (o *InlineResponse200327ByContributor) GetFailures() InlineResponse200327Failures1`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *InlineResponse200327ByContributor) GetFailuresOk() (*InlineResponse200327Failures1, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *InlineResponse200327ByContributor) SetFailures(v InlineResponse200327Failures1)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *InlineResponse200327ByContributor) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetFailureContribution

`func (o *InlineResponse200327ByContributor) GetFailureContribution() InlineResponse200327FailureContribution`

GetFailureContribution returns the FailureContribution field if non-nil, zero value otherwise.

### GetFailureContributionOk

`func (o *InlineResponse200327ByContributor) GetFailureContributionOk() (*InlineResponse200327FailureContribution, bool)`

GetFailureContributionOk returns a tuple with the FailureContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureContribution

`func (o *InlineResponse200327ByContributor) SetFailureContribution(v InlineResponse200327FailureContribution)`

SetFailureContribution sets FailureContribution field to given value.

### HasFailureContribution

`func (o *InlineResponse200327ByContributor) HasFailureContribution() bool`

HasFailureContribution returns a boolean if a field has been set.

### GetClientFailureContribution

`func (o *InlineResponse200327ByContributor) GetClientFailureContribution() InlineResponse200327ClientFailureContribution`

GetClientFailureContribution returns the ClientFailureContribution field if non-nil, zero value otherwise.

### GetClientFailureContributionOk

`func (o *InlineResponse200327ByContributor) GetClientFailureContributionOk() (*InlineResponse200327ClientFailureContribution, bool)`

GetClientFailureContributionOk returns a tuple with the ClientFailureContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFailureContribution

`func (o *InlineResponse200327ByContributor) SetClientFailureContribution(v InlineResponse200327ClientFailureContribution)`

SetClientFailureContribution sets ClientFailureContribution field to given value.

### HasClientFailureContribution

`func (o *InlineResponse200327ByContributor) HasClientFailureContribution() bool`

HasClientFailureContribution returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200327ByContributor) GetClients() InlineResponse200327Clients1`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200327ByContributor) GetClientsOk() (*InlineResponse200327Clients1, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200327ByContributor) SetClients(v InlineResponse200327Clients1)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200327ByContributor) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetBySubContributor

`func (o *InlineResponse200327ByContributor) GetBySubContributor() []InlineResponse200327BySubContributor`

GetBySubContributor returns the BySubContributor field if non-nil, zero value otherwise.

### GetBySubContributorOk

`func (o *InlineResponse200327ByContributor) GetBySubContributorOk() (*[]InlineResponse200327BySubContributor, bool)`

GetBySubContributorOk returns a tuple with the BySubContributor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySubContributor

`func (o *InlineResponse200327ByContributor) SetBySubContributor(v []InlineResponse200327BySubContributor)`

SetBySubContributor sets BySubContributor field to given value.

### HasBySubContributor

`func (o *InlineResponse200327ByContributor) HasBySubContributor() bool`

HasBySubContributor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


