# InlineResponse200133

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The type of command sent to the device. | [optional] 
**Name** | Pointer to **string** | The name of the device to which the command is sent. | [optional] 
**Details** | Pointer to **string** | A JSON string object containing command details. | [optional] 
**DashboardUser** | Pointer to **string** | The Meraki dashboard user who initiated the command. | [optional] 
**Ts** | Pointer to **string** | The time the command was sent to the device. | [optional] 

## Methods

### NewInlineResponse200133

`func NewInlineResponse200133() *InlineResponse200133`

NewInlineResponse200133 instantiates a new InlineResponse200133 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200133WithDefaults

`func NewInlineResponse200133WithDefaults() *InlineResponse200133`

NewInlineResponse200133WithDefaults instantiates a new InlineResponse200133 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *InlineResponse200133) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InlineResponse200133) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InlineResponse200133) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InlineResponse200133) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200133) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200133) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200133) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200133) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200133) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200133) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200133) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200133) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDashboardUser

`func (o *InlineResponse200133) GetDashboardUser() string`

GetDashboardUser returns the DashboardUser field if non-nil, zero value otherwise.

### GetDashboardUserOk

`func (o *InlineResponse200133) GetDashboardUserOk() (*string, bool)`

GetDashboardUserOk returns a tuple with the DashboardUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUser

`func (o *InlineResponse200133) SetDashboardUser(v string)`

SetDashboardUser sets DashboardUser field to given value.

### HasDashboardUser

`func (o *InlineResponse200133) HasDashboardUser() bool`

HasDashboardUser returns a boolean if a field has been set.

### GetTs

`func (o *InlineResponse200133) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200133) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200133) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200133) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


