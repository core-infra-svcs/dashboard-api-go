# InlineResponse20112

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse20112Network**](InlineResponse20112Network.md) |  | [optional] 
**RuleId** | Pointer to **string** | Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default) | [optional] 
**Type** | Pointer to **string** | Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default) | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Updated at timestamp | [optional] 
**CreatedAt** | Pointer to **time.Time** | Created at timestamp | [optional] 
**Match** | Pointer to [**InlineResponse20112Match**](InlineResponse20112Match.md) |  | [optional] 

## Methods

### NewInlineResponse20112

`func NewInlineResponse20112() *InlineResponse20112`

NewInlineResponse20112 instantiates a new InlineResponse20112 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20112WithDefaults

`func NewInlineResponse20112WithDefaults() *InlineResponse20112`

NewInlineResponse20112WithDefaults instantiates a new InlineResponse20112 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse20112) GetNetwork() InlineResponse20112Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20112) GetNetworkOk() (*InlineResponse20112Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20112) SetNetwork(v InlineResponse20112Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse20112) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRuleId

`func (o *InlineResponse20112) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse20112) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse20112) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse20112) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20112) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20112) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20112) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20112) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse20112) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse20112) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse20112) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse20112) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse20112) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse20112) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse20112) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse20112) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMatch

`func (o *InlineResponse20112) GetMatch() InlineResponse20112Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *InlineResponse20112) GetMatchOk() (*InlineResponse20112Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *InlineResponse20112) SetMatch(v InlineResponse20112Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *InlineResponse20112) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


