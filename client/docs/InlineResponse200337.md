# InlineResponse200337

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** | The success or failure status of the API | [optional] 
**States** | Pointer to **[]string** | The set of different states on the spaces integration process | [optional] 
**Email** | Pointer to **string** | The meraki user who attempts the spaces integration | [optional] 
**AccountName** | Pointer to **string** | The spaces dashboard account name created in spaces integration | [optional] 
**AccountType** | Pointer to **string** | The spaces dashboard account type created in spaces integration | [optional] 

## Methods

### NewInlineResponse200337

`func NewInlineResponse200337() *InlineResponse200337`

NewInlineResponse200337 instantiates a new InlineResponse200337 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200337WithDefaults

`func NewInlineResponse200337WithDefaults() *InlineResponse200337`

NewInlineResponse200337WithDefaults instantiates a new InlineResponse200337 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InlineResponse200337) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200337) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200337) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200337) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStates

`func (o *InlineResponse200337) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *InlineResponse200337) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *InlineResponse200337) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *InlineResponse200337) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetEmail

`func (o *InlineResponse200337) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse200337) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse200337) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineResponse200337) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAccountName

`func (o *InlineResponse200337) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *InlineResponse200337) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *InlineResponse200337) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *InlineResponse200337) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountType

`func (o *InlineResponse200337) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *InlineResponse200337) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *InlineResponse200337) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *InlineResponse200337) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


