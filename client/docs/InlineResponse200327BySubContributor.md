# InlineResponse200327BySubContributor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Failure sub-contributor name. | [optional] 
**Failures** | Pointer to [**InlineResponse200327Failures2**](InlineResponse200327Failures2.md) |  | [optional] 
**FailureContribution** | Pointer to [**InlineResponse200327FailureContribution1**](InlineResponse200327FailureContribution1.md) |  | [optional] 
**Clients** | Pointer to [**InlineResponse200327Clients2**](InlineResponse200327Clients2.md) |  | [optional] 
**ClientFailureContribution** | Pointer to [**InlineResponse200327ClientFailureContribution1**](InlineResponse200327ClientFailureContribution1.md) |  | [optional] 
**ByReason** | Pointer to [**[]InlineResponse200327ByReason**](InlineResponse200327ByReason.md) | Failure reasons. Maximum 100 entries. Sub-contributor totals include reasons not listed. No ordering is guaranteed. | [optional] 

## Methods

### NewInlineResponse200327BySubContributor

`func NewInlineResponse200327BySubContributor() *InlineResponse200327BySubContributor`

NewInlineResponse200327BySubContributor instantiates a new InlineResponse200327BySubContributor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200327BySubContributorWithDefaults

`func NewInlineResponse200327BySubContributorWithDefaults() *InlineResponse200327BySubContributor`

NewInlineResponse200327BySubContributorWithDefaults instantiates a new InlineResponse200327BySubContributor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200327BySubContributor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200327BySubContributor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200327BySubContributor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200327BySubContributor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFailures

`func (o *InlineResponse200327BySubContributor) GetFailures() InlineResponse200327Failures2`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *InlineResponse200327BySubContributor) GetFailuresOk() (*InlineResponse200327Failures2, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *InlineResponse200327BySubContributor) SetFailures(v InlineResponse200327Failures2)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *InlineResponse200327BySubContributor) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetFailureContribution

`func (o *InlineResponse200327BySubContributor) GetFailureContribution() InlineResponse200327FailureContribution1`

GetFailureContribution returns the FailureContribution field if non-nil, zero value otherwise.

### GetFailureContributionOk

`func (o *InlineResponse200327BySubContributor) GetFailureContributionOk() (*InlineResponse200327FailureContribution1, bool)`

GetFailureContributionOk returns a tuple with the FailureContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureContribution

`func (o *InlineResponse200327BySubContributor) SetFailureContribution(v InlineResponse200327FailureContribution1)`

SetFailureContribution sets FailureContribution field to given value.

### HasFailureContribution

`func (o *InlineResponse200327BySubContributor) HasFailureContribution() bool`

HasFailureContribution returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200327BySubContributor) GetClients() InlineResponse200327Clients2`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200327BySubContributor) GetClientsOk() (*InlineResponse200327Clients2, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200327BySubContributor) SetClients(v InlineResponse200327Clients2)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200327BySubContributor) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetClientFailureContribution

`func (o *InlineResponse200327BySubContributor) GetClientFailureContribution() InlineResponse200327ClientFailureContribution1`

GetClientFailureContribution returns the ClientFailureContribution field if non-nil, zero value otherwise.

### GetClientFailureContributionOk

`func (o *InlineResponse200327BySubContributor) GetClientFailureContributionOk() (*InlineResponse200327ClientFailureContribution1, bool)`

GetClientFailureContributionOk returns a tuple with the ClientFailureContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFailureContribution

`func (o *InlineResponse200327BySubContributor) SetClientFailureContribution(v InlineResponse200327ClientFailureContribution1)`

SetClientFailureContribution sets ClientFailureContribution field to given value.

### HasClientFailureContribution

`func (o *InlineResponse200327BySubContributor) HasClientFailureContribution() bool`

HasClientFailureContribution returns a boolean if a field has been set.

### GetByReason

`func (o *InlineResponse200327BySubContributor) GetByReason() []InlineResponse200327ByReason`

GetByReason returns the ByReason field if non-nil, zero value otherwise.

### GetByReasonOk

`func (o *InlineResponse200327BySubContributor) GetByReasonOk() (*[]InlineResponse200327ByReason, bool)`

GetByReasonOk returns a tuple with the ByReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByReason

`func (o *InlineResponse200327BySubContributor) SetByReason(v []InlineResponse200327ByReason)`

SetByReason sets ByReason field to given value.

### HasByReason

`func (o *InlineResponse200327BySubContributor) HasByReason() bool`

HasByReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


