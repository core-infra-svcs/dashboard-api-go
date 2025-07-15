# InlineResponse200383Redundancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | Wireless LAN controller role(Active, Active recovery, Standby hot, Standby recovery and Offline) | [optional] 
**Id** | Pointer to **string** | Wireless LAN controller redundancy ID | [optional] 
**ChassisName** | Pointer to **string** | Wireless LAN controller chassis name | [optional] 
**RedundantSerial** | Pointer to **string** | Wireless LAN controller redundant device serial | [optional] 
**Management** | Pointer to [**InlineResponse200383RedundancyManagement**](InlineResponse200383RedundancyManagement.md) |  | [optional] 

## Methods

### NewInlineResponse200383Redundancy

`func NewInlineResponse200383Redundancy() *InlineResponse200383Redundancy`

NewInlineResponse200383Redundancy instantiates a new InlineResponse200383Redundancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200383RedundancyWithDefaults

`func NewInlineResponse200383RedundancyWithDefaults() *InlineResponse200383Redundancy`

NewInlineResponse200383RedundancyWithDefaults instantiates a new InlineResponse200383Redundancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *InlineResponse200383Redundancy) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InlineResponse200383Redundancy) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InlineResponse200383Redundancy) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *InlineResponse200383Redundancy) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse200383Redundancy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200383Redundancy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200383Redundancy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200383Redundancy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChassisName

`func (o *InlineResponse200383Redundancy) GetChassisName() string`

GetChassisName returns the ChassisName field if non-nil, zero value otherwise.

### GetChassisNameOk

`func (o *InlineResponse200383Redundancy) GetChassisNameOk() (*string, bool)`

GetChassisNameOk returns a tuple with the ChassisName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisName

`func (o *InlineResponse200383Redundancy) SetChassisName(v string)`

SetChassisName sets ChassisName field to given value.

### HasChassisName

`func (o *InlineResponse200383Redundancy) HasChassisName() bool`

HasChassisName returns a boolean if a field has been set.

### GetRedundantSerial

`func (o *InlineResponse200383Redundancy) GetRedundantSerial() string`

GetRedundantSerial returns the RedundantSerial field if non-nil, zero value otherwise.

### GetRedundantSerialOk

`func (o *InlineResponse200383Redundancy) GetRedundantSerialOk() (*string, bool)`

GetRedundantSerialOk returns a tuple with the RedundantSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantSerial

`func (o *InlineResponse200383Redundancy) SetRedundantSerial(v string)`

SetRedundantSerial sets RedundantSerial field to given value.

### HasRedundantSerial

`func (o *InlineResponse200383Redundancy) HasRedundantSerial() bool`

HasRedundantSerial returns a boolean if a field has been set.

### GetManagement

`func (o *InlineResponse200383Redundancy) GetManagement() InlineResponse200383RedundancyManagement`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *InlineResponse200383Redundancy) GetManagementOk() (*InlineResponse200383RedundancyManagement, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *InlineResponse200383Redundancy) SetManagement(v InlineResponse200383RedundancyManagement)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *InlineResponse200383Redundancy) HasManagement() bool`

HasManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


