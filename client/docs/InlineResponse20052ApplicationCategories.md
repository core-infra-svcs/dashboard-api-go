# InlineResponse20052ApplicationCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the category | [optional] 
**Name** | Pointer to **string** | The name of the category | [optional] 
**Applications** | Pointer to [**[]InlineResponse20052Applications**](InlineResponse20052Applications.md) | Details of the associated applications | [optional] 

## Methods

### NewInlineResponse20052ApplicationCategories

`func NewInlineResponse20052ApplicationCategories() *InlineResponse20052ApplicationCategories`

NewInlineResponse20052ApplicationCategories instantiates a new InlineResponse20052ApplicationCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20052ApplicationCategoriesWithDefaults

`func NewInlineResponse20052ApplicationCategoriesWithDefaults() *InlineResponse20052ApplicationCategories`

NewInlineResponse20052ApplicationCategoriesWithDefaults instantiates a new InlineResponse20052ApplicationCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20052ApplicationCategories) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20052ApplicationCategories) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20052ApplicationCategories) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20052ApplicationCategories) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20052ApplicationCategories) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20052ApplicationCategories) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20052ApplicationCategories) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20052ApplicationCategories) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApplications

`func (o *InlineResponse20052ApplicationCategories) GetApplications() []InlineResponse20052Applications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *InlineResponse20052ApplicationCategories) GetApplicationsOk() (*[]InlineResponse20052Applications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *InlineResponse20052ApplicationCategories) SetApplications(v []InlineResponse20052Applications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *InlineResponse20052ApplicationCategories) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


