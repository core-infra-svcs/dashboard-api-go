# InlineResponse200303Override

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier for the override resource | [optional] 
**Security** | Pointer to [**InlineResponse200303OverrideSecurity**](InlineResponse200303OverrideSecurity.md) |  | [optional] 

## Methods

### NewInlineResponse200303Override

`func NewInlineResponse200303Override() *InlineResponse200303Override`

NewInlineResponse200303Override instantiates a new InlineResponse200303Override object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200303OverrideWithDefaults

`func NewInlineResponse200303OverrideWithDefaults() *InlineResponse200303Override`

NewInlineResponse200303OverrideWithDefaults instantiates a new InlineResponse200303Override object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200303Override) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200303Override) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200303Override) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200303Override) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSecurity

`func (o *InlineResponse200303Override) GetSecurity() InlineResponse200303OverrideSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *InlineResponse200303Override) GetSecurityOk() (*InlineResponse200303OverrideSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *InlineResponse200303Override) SetSecurity(v InlineResponse200303OverrideSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *InlineResponse200303Override) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


