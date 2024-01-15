# OrganizationsOrganizationIdActionBatchesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **bool** | Flag describing whether all actions in the action batch have completed | [optional] 
**Failed** | Pointer to **bool** | Flag describing whether any actions in the action batch failed | [optional] 
**Errors** | Pointer to **[]string** | List of errors encountered when running actions in the action batch | [optional] 
**CreatedResources** | [**[]OrganizationsOrganizationIdActionBatchesStatusCreatedResources**](OrganizationsOrganizationIdActionBatchesStatusCreatedResources.md) | Resources created as a result of this action batch | 

## Methods

### NewOrganizationsOrganizationIdActionBatchesStatus

`func NewOrganizationsOrganizationIdActionBatchesStatus(createdResources []OrganizationsOrganizationIdActionBatchesStatusCreatedResources, ) *OrganizationsOrganizationIdActionBatchesStatus`

NewOrganizationsOrganizationIdActionBatchesStatus instantiates a new OrganizationsOrganizationIdActionBatchesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdActionBatchesStatusWithDefaults

`func NewOrganizationsOrganizationIdActionBatchesStatusWithDefaults() *OrganizationsOrganizationIdActionBatchesStatus`

NewOrganizationsOrganizationIdActionBatchesStatusWithDefaults instantiates a new OrganizationsOrganizationIdActionBatchesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *OrganizationsOrganizationIdActionBatchesStatus) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *OrganizationsOrganizationIdActionBatchesStatus) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *OrganizationsOrganizationIdActionBatchesStatus) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *OrganizationsOrganizationIdActionBatchesStatus) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetFailed

`func (o *OrganizationsOrganizationIdActionBatchesStatus) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *OrganizationsOrganizationIdActionBatchesStatus) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *OrganizationsOrganizationIdActionBatchesStatus) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *OrganizationsOrganizationIdActionBatchesStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetErrors

`func (o *OrganizationsOrganizationIdActionBatchesStatus) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *OrganizationsOrganizationIdActionBatchesStatus) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *OrganizationsOrganizationIdActionBatchesStatus) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *OrganizationsOrganizationIdActionBatchesStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetCreatedResources

`func (o *OrganizationsOrganizationIdActionBatchesStatus) GetCreatedResources() []OrganizationsOrganizationIdActionBatchesStatusCreatedResources`

GetCreatedResources returns the CreatedResources field if non-nil, zero value otherwise.

### GetCreatedResourcesOk

`func (o *OrganizationsOrganizationIdActionBatchesStatus) GetCreatedResourcesOk() (*[]OrganizationsOrganizationIdActionBatchesStatusCreatedResources, bool)`

GetCreatedResourcesOk returns a tuple with the CreatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedResources

`func (o *OrganizationsOrganizationIdActionBatchesStatus) SetCreatedResources(v []OrganizationsOrganizationIdActionBatchesStatusCreatedResources)`

SetCreatedResources sets CreatedResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


