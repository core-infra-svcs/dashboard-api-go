# InlineResponse200248ServiceProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Service Provider name | [optional] 
**Plans** | Pointer to [**[]InlineResponse200248ServiceProviderPlans**](InlineResponse200248ServiceProviderPlans.md) | Plans currently active on the eSIM | [optional] 

## Methods

### NewInlineResponse200248ServiceProvider

`func NewInlineResponse200248ServiceProvider() *InlineResponse200248ServiceProvider`

NewInlineResponse200248ServiceProvider instantiates a new InlineResponse200248ServiceProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200248ServiceProviderWithDefaults

`func NewInlineResponse200248ServiceProviderWithDefaults() *InlineResponse200248ServiceProvider`

NewInlineResponse200248ServiceProviderWithDefaults instantiates a new InlineResponse200248ServiceProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200248ServiceProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200248ServiceProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200248ServiceProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200248ServiceProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlans

`func (o *InlineResponse200248ServiceProvider) GetPlans() []InlineResponse200248ServiceProviderPlans`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *InlineResponse200248ServiceProvider) GetPlansOk() (*[]InlineResponse200248ServiceProviderPlans, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *InlineResponse200248ServiceProvider) SetPlans(v []InlineResponse200248ServiceProviderPlans)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *InlineResponse200248ServiceProvider) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


