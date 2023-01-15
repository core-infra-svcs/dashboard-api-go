# InlineResponse20074SentryEnrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemsManagerNetwork** | Pointer to [**InlineResponse20074SentryEnrollmentSystemsManagerNetwork**](InlineResponse20074SentryEnrollmentSystemsManagerNetwork.md) |  | [optional] 
**Strength** | Pointer to **string** | The strength of the enforcement of selected system types. | [optional] 
**EnforcedSystems** | Pointer to **[]string** | The system types that the Sentry enforces. | [optional] 

## Methods

### NewInlineResponse20074SentryEnrollment

`func NewInlineResponse20074SentryEnrollment() *InlineResponse20074SentryEnrollment`

NewInlineResponse20074SentryEnrollment instantiates a new InlineResponse20074SentryEnrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20074SentryEnrollmentWithDefaults

`func NewInlineResponse20074SentryEnrollmentWithDefaults() *InlineResponse20074SentryEnrollment`

NewInlineResponse20074SentryEnrollmentWithDefaults instantiates a new InlineResponse20074SentryEnrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemsManagerNetwork

`func (o *InlineResponse20074SentryEnrollment) GetSystemsManagerNetwork() InlineResponse20074SentryEnrollmentSystemsManagerNetwork`

GetSystemsManagerNetwork returns the SystemsManagerNetwork field if non-nil, zero value otherwise.

### GetSystemsManagerNetworkOk

`func (o *InlineResponse20074SentryEnrollment) GetSystemsManagerNetworkOk() (*InlineResponse20074SentryEnrollmentSystemsManagerNetwork, bool)`

GetSystemsManagerNetworkOk returns a tuple with the SystemsManagerNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemsManagerNetwork

`func (o *InlineResponse20074SentryEnrollment) SetSystemsManagerNetwork(v InlineResponse20074SentryEnrollmentSystemsManagerNetwork)`

SetSystemsManagerNetwork sets SystemsManagerNetwork field to given value.

### HasSystemsManagerNetwork

`func (o *InlineResponse20074SentryEnrollment) HasSystemsManagerNetwork() bool`

HasSystemsManagerNetwork returns a boolean if a field has been set.

### GetStrength

`func (o *InlineResponse20074SentryEnrollment) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *InlineResponse20074SentryEnrollment) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *InlineResponse20074SentryEnrollment) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *InlineResponse20074SentryEnrollment) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetEnforcedSystems

`func (o *InlineResponse20074SentryEnrollment) GetEnforcedSystems() []string`

GetEnforcedSystems returns the EnforcedSystems field if non-nil, zero value otherwise.

### GetEnforcedSystemsOk

`func (o *InlineResponse20074SentryEnrollment) GetEnforcedSystemsOk() (*[]string, bool)`

GetEnforcedSystemsOk returns a tuple with the EnforcedSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcedSystems

`func (o *InlineResponse20074SentryEnrollment) SetEnforcedSystems(v []string)`

SetEnforcedSystems sets EnforcedSystems field to given value.

### HasEnforcedSystems

`func (o *InlineResponse20074SentryEnrollment) HasEnforcedSystems() bool`

HasEnforcedSystems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


