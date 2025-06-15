# InlineResponse200285

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of Early Access Feature | [optional] 
**ShortName** | Pointer to **string** | Name of Early Access Feature | [optional] 
**LimitScopeToNetworks** | Pointer to [**[]InlineResponse200285LimitScopeToNetworks**](InlineResponse200285LimitScopeToNetworks.md) | Networks assigned to the Early Access Feature | [optional] 
**OptOutEligibility** | Pointer to [**InlineResponse200285OptOutEligibility**](InlineResponse200285OptOutEligibility.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when Early Access Feature was created | [optional] 

## Methods

### NewInlineResponse200285

`func NewInlineResponse200285() *InlineResponse200285`

NewInlineResponse200285 instantiates a new InlineResponse200285 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200285WithDefaults

`func NewInlineResponse200285WithDefaults() *InlineResponse200285`

NewInlineResponse200285WithDefaults instantiates a new InlineResponse200285 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200285) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200285) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200285) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200285) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShortName

`func (o *InlineResponse200285) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *InlineResponse200285) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *InlineResponse200285) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *InlineResponse200285) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetLimitScopeToNetworks

`func (o *InlineResponse200285) GetLimitScopeToNetworks() []InlineResponse200285LimitScopeToNetworks`

GetLimitScopeToNetworks returns the LimitScopeToNetworks field if non-nil, zero value otherwise.

### GetLimitScopeToNetworksOk

`func (o *InlineResponse200285) GetLimitScopeToNetworksOk() (*[]InlineResponse200285LimitScopeToNetworks, bool)`

GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitScopeToNetworks

`func (o *InlineResponse200285) SetLimitScopeToNetworks(v []InlineResponse200285LimitScopeToNetworks)`

SetLimitScopeToNetworks sets LimitScopeToNetworks field to given value.

### HasLimitScopeToNetworks

`func (o *InlineResponse200285) HasLimitScopeToNetworks() bool`

HasLimitScopeToNetworks returns a boolean if a field has been set.

### GetOptOutEligibility

`func (o *InlineResponse200285) GetOptOutEligibility() InlineResponse200285OptOutEligibility`

GetOptOutEligibility returns the OptOutEligibility field if non-nil, zero value otherwise.

### GetOptOutEligibilityOk

`func (o *InlineResponse200285) GetOptOutEligibilityOk() (*InlineResponse200285OptOutEligibility, bool)`

GetOptOutEligibilityOk returns a tuple with the OptOutEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptOutEligibility

`func (o *InlineResponse200285) SetOptOutEligibility(v InlineResponse200285OptOutEligibility)`

SetOptOutEligibility sets OptOutEligibility field to given value.

### HasOptOutEligibility

`func (o *InlineResponse200285) HasOptOutEligibility() bool`

HasOptOutEligibility returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200285) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200285) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200285) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200285) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


