# InlineObject207

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confirmed** | Pointer to **bool** | Set to true for immediate execution. Set to false if the action should be previewed before executing. This property cannot be unset once it is true. Defaults to false. | [optional] 
**Synchronous** | Pointer to **bool** | Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch. Defaults to false. | [optional] 
**Actions** | [**[]OrganizationsOrganizationIdActionBatchesActions1**](OrganizationsOrganizationIdActionBatchesActions1.md) | A set of changes to make as part of this action (&lt;a href&#x3D;&#39;https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/&#39;&gt;more details&lt;/a&gt;) | 
**Callback** | Pointer to [**DevicesSerialLiveToolsArpTableCallback**](DevicesSerialLiveToolsArpTableCallback.md) |  | [optional] 

## Methods

### NewInlineObject207

`func NewInlineObject207(actions []OrganizationsOrganizationIdActionBatchesActions1, ) *InlineObject207`

NewInlineObject207 instantiates a new InlineObject207 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject207WithDefaults

`func NewInlineObject207WithDefaults() *InlineObject207`

NewInlineObject207WithDefaults instantiates a new InlineObject207 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmed

`func (o *InlineObject207) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *InlineObject207) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *InlineObject207) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *InlineObject207) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetSynchronous

`func (o *InlineObject207) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *InlineObject207) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *InlineObject207) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *InlineObject207) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.

### GetActions

`func (o *InlineObject207) GetActions() []OrganizationsOrganizationIdActionBatchesActions1`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *InlineObject207) GetActionsOk() (*[]OrganizationsOrganizationIdActionBatchesActions1, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *InlineObject207) SetActions(v []OrganizationsOrganizationIdActionBatchesActions1)`

SetActions sets Actions field to given value.


### GetCallback

`func (o *InlineObject207) GetCallback() DevicesSerialLiveToolsArpTableCallback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *InlineObject207) GetCallbackOk() (*DevicesSerialLiveToolsArpTableCallback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *InlineObject207) SetCallback(v DevicesSerialLiveToolsArpTableCallback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *InlineObject207) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


