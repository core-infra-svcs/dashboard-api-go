# InlineResponse200248Profiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomApns** | Pointer to **[]string** | Available custom APNs for the profile | [optional] 
**Iccid** | Pointer to **string** | eSIM profile ID | [optional] 
**Status** | Pointer to **string** | eSIM profile status | [optional] 
**ServiceProvider** | Pointer to [**InlineResponse200248ServiceProvider**](InlineResponse200248ServiceProvider.md) |  | [optional] 

## Methods

### NewInlineResponse200248Profiles

`func NewInlineResponse200248Profiles() *InlineResponse200248Profiles`

NewInlineResponse200248Profiles instantiates a new InlineResponse200248Profiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200248ProfilesWithDefaults

`func NewInlineResponse200248ProfilesWithDefaults() *InlineResponse200248Profiles`

NewInlineResponse200248ProfilesWithDefaults instantiates a new InlineResponse200248Profiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomApns

`func (o *InlineResponse200248Profiles) GetCustomApns() []string`

GetCustomApns returns the CustomApns field if non-nil, zero value otherwise.

### GetCustomApnsOk

`func (o *InlineResponse200248Profiles) GetCustomApnsOk() (*[]string, bool)`

GetCustomApnsOk returns a tuple with the CustomApns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomApns

`func (o *InlineResponse200248Profiles) SetCustomApns(v []string)`

SetCustomApns sets CustomApns field to given value.

### HasCustomApns

`func (o *InlineResponse200248Profiles) HasCustomApns() bool`

HasCustomApns returns a boolean if a field has been set.

### GetIccid

`func (o *InlineResponse200248Profiles) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *InlineResponse200248Profiles) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *InlineResponse200248Profiles) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *InlineResponse200248Profiles) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200248Profiles) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200248Profiles) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200248Profiles) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200248Profiles) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetServiceProvider

`func (o *InlineResponse200248Profiles) GetServiceProvider() InlineResponse200248ServiceProvider`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *InlineResponse200248Profiles) GetServiceProviderOk() (*InlineResponse200248ServiceProvider, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *InlineResponse200248Profiles) SetServiceProvider(v InlineResponse200248ServiceProvider)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *InlineResponse200248Profiles) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


