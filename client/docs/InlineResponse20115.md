# InlineResponse20115

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId} | [optional] 
**OrganizationId** | Pointer to **string** | ID of the organization this action batch belongs to | [optional] 
**Confirmed** | Pointer to **bool** | Flag describing whether the action should be previewed before executing or not | [optional] 
**Synchronous** | Pointer to **bool** | Flag describing whether actions should run synchronously or asynchronously | [optional] 
**Status** | Pointer to [**OrganizationsOrganizationIdActionBatchesStatus**](OrganizationsOrganizationIdActionBatchesStatus.md) |  | [optional] 
**Actions** | [**[]OrganizationsOrganizationIdActionBatchesActions**](OrganizationsOrganizationIdActionBatchesActions.md) | A set of changes made as part of this action (&lt;a href&#x3D;&#39;https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/&#39;&gt;more details&lt;/a&gt;) | 
**Callback** | Pointer to [**InlineResponse2011Callback**](InlineResponse2011Callback.md) |  | [optional] 

## Methods

### NewInlineResponse20115

`func NewInlineResponse20115(actions []OrganizationsOrganizationIdActionBatchesActions, ) *InlineResponse20115`

NewInlineResponse20115 instantiates a new InlineResponse20115 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20115WithDefaults

`func NewInlineResponse20115WithDefaults() *InlineResponse20115`

NewInlineResponse20115WithDefaults instantiates a new InlineResponse20115 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20115) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20115) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20115) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20115) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InlineResponse20115) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InlineResponse20115) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InlineResponse20115) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InlineResponse20115) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetConfirmed

`func (o *InlineResponse20115) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *InlineResponse20115) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *InlineResponse20115) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *InlineResponse20115) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetSynchronous

`func (o *InlineResponse20115) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *InlineResponse20115) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *InlineResponse20115) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *InlineResponse20115) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20115) GetStatus() OrganizationsOrganizationIdActionBatchesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20115) GetStatusOk() (*OrganizationsOrganizationIdActionBatchesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20115) SetStatus(v OrganizationsOrganizationIdActionBatchesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20115) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActions

`func (o *InlineResponse20115) GetActions() []OrganizationsOrganizationIdActionBatchesActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *InlineResponse20115) GetActionsOk() (*[]OrganizationsOrganizationIdActionBatchesActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *InlineResponse20115) SetActions(v []OrganizationsOrganizationIdActionBatchesActions)`

SetActions sets Actions field to given value.


### GetCallback

`func (o *InlineResponse20115) GetCallback() InlineResponse2011Callback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineResponse20115) GetCallbackOk() (*InlineResponse2011Callback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineResponse20115) SetCallback(v InlineResponse2011Callback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineResponse20115) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


