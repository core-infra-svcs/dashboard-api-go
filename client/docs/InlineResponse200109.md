# InlineResponse200109

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of client | [optional] 
**ClientId** | Pointer to **string** | ID of client | [optional] 
**Assigned** | Pointer to [**[]NetworksNetworkIdPoliciesByClientAssigned**](NetworksNetworkIdPoliciesByClientAssigned.md) | Assigned policies | [optional] 

## Methods

### NewInlineResponse200109

`func NewInlineResponse200109() *InlineResponse200109`

NewInlineResponse200109 instantiates a new InlineResponse200109 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200109WithDefaults

`func NewInlineResponse200109WithDefaults() *InlineResponse200109`

NewInlineResponse200109WithDefaults instantiates a new InlineResponse200109 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200109) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200109) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200109) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200109) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientId

`func (o *InlineResponse200109) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineResponse200109) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineResponse200109) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineResponse200109) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAssigned

`func (o *InlineResponse200109) GetAssigned() []NetworksNetworkIdPoliciesByClientAssigned`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *InlineResponse200109) GetAssignedOk() (*[]NetworksNetworkIdPoliciesByClientAssigned, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *InlineResponse200109) SetAssigned(v []NetworksNetworkIdPoliciesByClientAssigned)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *InlineResponse200109) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


