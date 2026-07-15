# InlineResponse200359

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of Early Access Feature | [optional] 
**ShortName** | Pointer to **string** | Name of Early Access Feature | [optional] 
**LimitScopeToNetworks** | Pointer to [**[]InlineResponse200359LimitScopeToNetworks**](InlineResponse200359LimitScopeToNetworks.md) | Networks assigned to the Early Access Feature | [optional] 
**OptOutEligibility** | Pointer to [**InlineResponse200359OptOutEligibility**](InlineResponse200359OptOutEligibility.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when Early Access Feature was created | [optional] 

## Methods

### NewInlineResponse200359

`func NewInlineResponse200359() *InlineResponse200359`

NewInlineResponse200359 instantiates a new InlineResponse200359 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200359WithDefaults

`func NewInlineResponse200359WithDefaults() *InlineResponse200359`

NewInlineResponse200359WithDefaults instantiates a new InlineResponse200359 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200359) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200359) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200359) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200359) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShortName

`func (o *InlineResponse200359) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *InlineResponse200359) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *InlineResponse200359) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *InlineResponse200359) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetLimitScopeToNetworks

`func (o *InlineResponse200359) GetLimitScopeToNetworks() []InlineResponse200359LimitScopeToNetworks`

GetLimitScopeToNetworks returns the LimitScopeToNetworks field if non-nil, zero value otherwise.

### GetLimitScopeToNetworksOk

`func (o *InlineResponse200359) GetLimitScopeToNetworksOk() (*[]InlineResponse200359LimitScopeToNetworks, bool)`

GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitScopeToNetworks

`func (o *InlineResponse200359) SetLimitScopeToNetworks(v []InlineResponse200359LimitScopeToNetworks)`

SetLimitScopeToNetworks sets LimitScopeToNetworks field to given value.

### HasLimitScopeToNetworks

`func (o *InlineResponse200359) HasLimitScopeToNetworks() bool`

HasLimitScopeToNetworks returns a boolean if a field has been set.

### GetOptOutEligibility

`func (o *InlineResponse200359) GetOptOutEligibility() InlineResponse200359OptOutEligibility`

GetOptOutEligibility returns the OptOutEligibility field if non-nil, zero value otherwise.

### GetOptOutEligibilityOk

`func (o *InlineResponse200359) GetOptOutEligibilityOk() (*InlineResponse200359OptOutEligibility, bool)`

GetOptOutEligibilityOk returns a tuple with the OptOutEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptOutEligibility

`func (o *InlineResponse200359) SetOptOutEligibility(v InlineResponse200359OptOutEligibility)`

SetOptOutEligibility sets OptOutEligibility field to given value.

### HasOptOutEligibility

`func (o *InlineResponse200359) HasOptOutEligibility() bool`

HasOptOutEligibility returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200359) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200359) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200359) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200359) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


