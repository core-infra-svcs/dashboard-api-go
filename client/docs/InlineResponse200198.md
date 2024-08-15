# InlineResponse200198

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId} | [optional] 
**OrganizationId** | Pointer to **string** | ID of the organization this action batch belongs to | [optional] 
**Confirmed** | Pointer to **bool** | Flag describing whether the action should be previewed before executing or not | [optional] 
**Synchronous** | Pointer to **bool** | Flag describing whether actions should run synchronously or asynchronously | [optional] 
**Status** | Pointer to [**OrganizationsOrganizationIdActionBatchesStatus**](OrganizationsOrganizationIdActionBatchesStatus.md) |  | [optional] 
**Actions** | [**[]OrganizationsOrganizationIdActionBatchesActions**](OrganizationsOrganizationIdActionBatchesActions.md) | A set of changes made as part of this action (&lt;a href&#x3D;&#39;https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/&#39;&gt;more details&lt;/a&gt;) | 

## Methods

### NewInlineResponse200198

`func NewInlineResponse200198(actions []OrganizationsOrganizationIdActionBatchesActions, ) *InlineResponse200198`

NewInlineResponse200198 instantiates a new InlineResponse200198 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200198WithDefaults

`func NewInlineResponse200198WithDefaults() *InlineResponse200198`

NewInlineResponse200198WithDefaults instantiates a new InlineResponse200198 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200198) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200198) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200198) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200198) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InlineResponse200198) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InlineResponse200198) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InlineResponse200198) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InlineResponse200198) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetConfirmed

`func (o *InlineResponse200198) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *InlineResponse200198) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *InlineResponse200198) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *InlineResponse200198) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetSynchronous

`func (o *InlineResponse200198) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *InlineResponse200198) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *InlineResponse200198) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *InlineResponse200198) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200198) GetStatus() OrganizationsOrganizationIdActionBatchesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200198) GetStatusOk() (*OrganizationsOrganizationIdActionBatchesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200198) SetStatus(v OrganizationsOrganizationIdActionBatchesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200198) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActions

`func (o *InlineResponse200198) GetActions() []OrganizationsOrganizationIdActionBatchesActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *InlineResponse200198) GetActionsOk() (*[]OrganizationsOrganizationIdActionBatchesActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *InlineResponse200198) SetActions(v []OrganizationsOrganizationIdActionBatchesActions)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


