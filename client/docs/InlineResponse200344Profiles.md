# InlineResponse200344Profiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomApns** | Pointer to **[]string** | Available custom APNs for the profile | [optional] 
**Iccid** | Pointer to **string** | eSIM profile ID | [optional] 
**Status** | Pointer to **string** | eSIM profile status | [optional] 
**ServiceProvider** | Pointer to [**InlineResponse200344ServiceProvider**](InlineResponse200344ServiceProvider.md) |  | [optional] 

## Methods

### NewInlineResponse200344Profiles

`func NewInlineResponse200344Profiles() *InlineResponse200344Profiles`

NewInlineResponse200344Profiles instantiates a new InlineResponse200344Profiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200344ProfilesWithDefaults

`func NewInlineResponse200344ProfilesWithDefaults() *InlineResponse200344Profiles`

NewInlineResponse200344ProfilesWithDefaults instantiates a new InlineResponse200344Profiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomApns

`func (o *InlineResponse200344Profiles) GetCustomApns() []string`

GetCustomApns returns the CustomApns field if non-nil, zero value otherwise.

### GetCustomApnsOk

`func (o *InlineResponse200344Profiles) GetCustomApnsOk() (*[]string, bool)`

GetCustomApnsOk returns a tuple with the CustomApns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomApns

`func (o *InlineResponse200344Profiles) SetCustomApns(v []string)`

SetCustomApns sets CustomApns field to given value.

### HasCustomApns

`func (o *InlineResponse200344Profiles) HasCustomApns() bool`

HasCustomApns returns a boolean if a field has been set.

### GetIccid

`func (o *InlineResponse200344Profiles) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *InlineResponse200344Profiles) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *InlineResponse200344Profiles) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *InlineResponse200344Profiles) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200344Profiles) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200344Profiles) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200344Profiles) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200344Profiles) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetServiceProvider

`func (o *InlineResponse200344Profiles) GetServiceProvider() InlineResponse200344ServiceProvider`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *InlineResponse200344Profiles) GetServiceProviderOk() (*InlineResponse200344ServiceProvider, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *InlineResponse200344Profiles) SetServiceProvider(v InlineResponse200344ServiceProvider)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *InlineResponse200344Profiles) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


