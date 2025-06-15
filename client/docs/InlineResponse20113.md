# InlineResponse20113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse20113Network**](InlineResponse20113Network.md) |  | [optional] 
**RuleId** | Pointer to **string** | Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default) | [optional] 
**Type** | Pointer to **string** | Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default) | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Updated at timestamp | [optional] 
**CreatedAt** | Pointer to **time.Time** | Created at timestamp | [optional] 
**Match** | Pointer to [**InlineResponse20113Match**](InlineResponse20113Match.md) |  | [optional] 

## Methods

### NewInlineResponse20113

`func NewInlineResponse20113() *InlineResponse20113`

NewInlineResponse20113 instantiates a new InlineResponse20113 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20113WithDefaults

`func NewInlineResponse20113WithDefaults() *InlineResponse20113`

NewInlineResponse20113WithDefaults instantiates a new InlineResponse20113 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse20113) GetNetwork() InlineResponse20113Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20113) GetNetworkOk() (*InlineResponse20113Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20113) SetNetwork(v InlineResponse20113Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse20113) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRuleId

`func (o *InlineResponse20113) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *InlineResponse20113) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *InlineResponse20113) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *InlineResponse20113) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20113) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20113) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20113) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20113) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse20113) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse20113) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse20113) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse20113) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse20113) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse20113) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse20113) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse20113) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMatch

`func (o *InlineResponse20113) GetMatch() InlineResponse20113Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *InlineResponse20113) GetMatchOk() (*InlineResponse20113Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *InlineResponse20113) SetMatch(v InlineResponse20113Match)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *InlineResponse20113) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


