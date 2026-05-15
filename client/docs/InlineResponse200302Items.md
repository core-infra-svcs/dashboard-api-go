# InlineResponse200302Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Service provider name. | [optional] 
**Logo** | Pointer to [**InlineResponse200302Logo**](InlineResponse200302Logo.md) |  | [optional] 
**IsBootstrap** | Pointer to **bool** | Indicates if service provider is the bootstrap provider. | [optional] 
**Terms** | Pointer to [**InlineResponse200302Terms**](InlineResponse200302Terms.md) |  | [optional] 

## Methods

### NewInlineResponse200302Items

`func NewInlineResponse200302Items() *InlineResponse200302Items`

NewInlineResponse200302Items instantiates a new InlineResponse200302Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200302ItemsWithDefaults

`func NewInlineResponse200302ItemsWithDefaults() *InlineResponse200302Items`

NewInlineResponse200302ItemsWithDefaults instantiates a new InlineResponse200302Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200302Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200302Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200302Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200302Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogo

`func (o *InlineResponse200302Items) GetLogo() InlineResponse200302Logo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *InlineResponse200302Items) GetLogoOk() (*InlineResponse200302Logo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *InlineResponse200302Items) SetLogo(v InlineResponse200302Logo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *InlineResponse200302Items) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetIsBootstrap

`func (o *InlineResponse200302Items) GetIsBootstrap() bool`

GetIsBootstrap returns the IsBootstrap field if non-nil, zero value otherwise.

### GetIsBootstrapOk

`func (o *InlineResponse200302Items) GetIsBootstrapOk() (*bool, bool)`

GetIsBootstrapOk returns a tuple with the IsBootstrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootstrap

`func (o *InlineResponse200302Items) SetIsBootstrap(v bool)`

SetIsBootstrap sets IsBootstrap field to given value.

### HasIsBootstrap

`func (o *InlineResponse200302Items) HasIsBootstrap() bool`

HasIsBootstrap returns a boolean if a field has been set.

### GetTerms

`func (o *InlineResponse200302Items) GetTerms() InlineResponse200302Terms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *InlineResponse200302Items) GetTermsOk() (*InlineResponse200302Terms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *InlineResponse200302Items) SetTerms(v InlineResponse200302Terms)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *InlineResponse200302Items) HasTerms() bool`

HasTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


