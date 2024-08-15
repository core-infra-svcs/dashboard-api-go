# InlineResponse200194SentryEnrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemsManagerNetwork** | Pointer to [**InlineResponse200194SentryEnrollmentSystemsManagerNetwork**](InlineResponse200194SentryEnrollmentSystemsManagerNetwork.md) |  | [optional] 
**Strength** | Pointer to **string** | The strength of the enforcement of selected system types. | [optional] 
**EnforcedSystems** | Pointer to **[]string** | The system types that the Sentry enforces. | [optional] 

## Methods

### NewInlineResponse200194SentryEnrollment

`func NewInlineResponse200194SentryEnrollment() *InlineResponse200194SentryEnrollment`

NewInlineResponse200194SentryEnrollment instantiates a new InlineResponse200194SentryEnrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200194SentryEnrollmentWithDefaults

`func NewInlineResponse200194SentryEnrollmentWithDefaults() *InlineResponse200194SentryEnrollment`

NewInlineResponse200194SentryEnrollmentWithDefaults instantiates a new InlineResponse200194SentryEnrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemsManagerNetwork

`func (o *InlineResponse200194SentryEnrollment) GetSystemsManagerNetwork() InlineResponse200194SentryEnrollmentSystemsManagerNetwork`

GetSystemsManagerNetwork returns the SystemsManagerNetwork field if non-nil, zero value otherwise.

### GetSystemsManagerNetworkOk

`func (o *InlineResponse200194SentryEnrollment) GetSystemsManagerNetworkOk() (*InlineResponse200194SentryEnrollmentSystemsManagerNetwork, bool)`

GetSystemsManagerNetworkOk returns a tuple with the SystemsManagerNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemsManagerNetwork

`func (o *InlineResponse200194SentryEnrollment) SetSystemsManagerNetwork(v InlineResponse200194SentryEnrollmentSystemsManagerNetwork)`

SetSystemsManagerNetwork sets SystemsManagerNetwork field to given value.

### HasSystemsManagerNetwork

`func (o *InlineResponse200194SentryEnrollment) HasSystemsManagerNetwork() bool`

HasSystemsManagerNetwork returns a boolean if a field has been set.

### GetStrength

`func (o *InlineResponse200194SentryEnrollment) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *InlineResponse200194SentryEnrollment) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *InlineResponse200194SentryEnrollment) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *InlineResponse200194SentryEnrollment) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetEnforcedSystems

`func (o *InlineResponse200194SentryEnrollment) GetEnforcedSystems() []string`

GetEnforcedSystems returns the EnforcedSystems field if non-nil, zero value otherwise.

### GetEnforcedSystemsOk

`func (o *InlineResponse200194SentryEnrollment) GetEnforcedSystemsOk() (*[]string, bool)`

GetEnforcedSystemsOk returns a tuple with the EnforcedSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcedSystems

`func (o *InlineResponse200194SentryEnrollment) SetEnforcedSystems(v []string)`

SetEnforcedSystems sets EnforcedSystems field to given value.

### HasEnforcedSystems

`func (o *InlineResponse200194SentryEnrollment) HasEnforcedSystems() bool`

HasEnforcedSystems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


