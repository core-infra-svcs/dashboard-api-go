# InlineResponse200253Profiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomApns** | Pointer to **[]string** | Available custom APNs for the profile | [optional] 
**Iccid** | Pointer to **string** | eSIM profile ID | [optional] 
**Status** | Pointer to **string** | eSIM profile status | [optional] 
**ServiceProvider** | Pointer to [**InlineResponse200253ServiceProvider**](InlineResponse200253ServiceProvider.md) |  | [optional] 

## Methods

### NewInlineResponse200253Profiles

`func NewInlineResponse200253Profiles() *InlineResponse200253Profiles`

NewInlineResponse200253Profiles instantiates a new InlineResponse200253Profiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200253ProfilesWithDefaults

`func NewInlineResponse200253ProfilesWithDefaults() *InlineResponse200253Profiles`

NewInlineResponse200253ProfilesWithDefaults instantiates a new InlineResponse200253Profiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomApns

`func (o *InlineResponse200253Profiles) GetCustomApns() []string`

GetCustomApns returns the CustomApns field if non-nil, zero value otherwise.

### GetCustomApnsOk

`func (o *InlineResponse200253Profiles) GetCustomApnsOk() (*[]string, bool)`

GetCustomApnsOk returns a tuple with the CustomApns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomApns

`func (o *InlineResponse200253Profiles) SetCustomApns(v []string)`

SetCustomApns sets CustomApns field to given value.

### HasCustomApns

`func (o *InlineResponse200253Profiles) HasCustomApns() bool`

HasCustomApns returns a boolean if a field has been set.

### GetIccid

`func (o *InlineResponse200253Profiles) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *InlineResponse200253Profiles) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *InlineResponse200253Profiles) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *InlineResponse200253Profiles) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200253Profiles) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200253Profiles) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200253Profiles) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200253Profiles) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetServiceProvider

`func (o *InlineResponse200253Profiles) GetServiceProvider() InlineResponse200253ServiceProvider`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *InlineResponse200253Profiles) GetServiceProviderOk() (*InlineResponse200253ServiceProvider, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *InlineResponse200253Profiles) SetServiceProvider(v InlineResponse200253ServiceProvider)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *InlineResponse200253Profiles) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


