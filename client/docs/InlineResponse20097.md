# InlineResponse20097

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Alert identifier. Value can be empty | [optional] 
**Category** | Pointer to **string** | Category of the alert | [optional] 
**Type** | Pointer to **string** | Alert type | [optional] 
**Severity** | Pointer to **string** | Severity of the alert | [optional] 
**Scope** | Pointer to [**NetworksNetworkIdHealthAlertsScope**](NetworksNetworkIdHealthAlertsScope.md) |  | [optional] 

## Methods

### NewInlineResponse20097

`func NewInlineResponse20097() *InlineResponse20097`

NewInlineResponse20097 instantiates a new InlineResponse20097 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20097WithDefaults

`func NewInlineResponse20097WithDefaults() *InlineResponse20097`

NewInlineResponse20097WithDefaults instantiates a new InlineResponse20097 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20097) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20097) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20097) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20097) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse20097) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse20097) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse20097) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse20097) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20097) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20097) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20097) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20097) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSeverity

`func (o *InlineResponse20097) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InlineResponse20097) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InlineResponse20097) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *InlineResponse20097) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetScope

`func (o *InlineResponse20097) GetScope() NetworksNetworkIdHealthAlertsScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineResponse20097) GetScopeOk() (*NetworksNetworkIdHealthAlertsScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineResponse20097) SetScope(v NetworksNetworkIdHealthAlertsScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineResponse20097) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


