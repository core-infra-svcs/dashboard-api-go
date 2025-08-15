# InlineResponse200208SentryEnrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemsManagerNetwork** | Pointer to [**InlineResponse200208SentryEnrollmentSystemsManagerNetwork**](InlineResponse200208SentryEnrollmentSystemsManagerNetwork.md) |  | [optional] 
**Strength** | Pointer to **string** | The strength of the enforcement of selected system types. | [optional] 
**EnforcedSystems** | Pointer to **[]string** | The system types that the Sentry enforces. | [optional] 

## Methods

### NewInlineResponse200208SentryEnrollment

`func NewInlineResponse200208SentryEnrollment() *InlineResponse200208SentryEnrollment`

NewInlineResponse200208SentryEnrollment instantiates a new InlineResponse200208SentryEnrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200208SentryEnrollmentWithDefaults

`func NewInlineResponse200208SentryEnrollmentWithDefaults() *InlineResponse200208SentryEnrollment`

NewInlineResponse200208SentryEnrollmentWithDefaults instantiates a new InlineResponse200208SentryEnrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemsManagerNetwork

`func (o *InlineResponse200208SentryEnrollment) GetSystemsManagerNetwork() InlineResponse200208SentryEnrollmentSystemsManagerNetwork`

GetSystemsManagerNetwork returns the SystemsManagerNetwork field if non-nil, zero value otherwise.

### GetSystemsManagerNetworkOk

`func (o *InlineResponse200208SentryEnrollment) GetSystemsManagerNetworkOk() (*InlineResponse200208SentryEnrollmentSystemsManagerNetwork, bool)`

GetSystemsManagerNetworkOk returns a tuple with the SystemsManagerNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemsManagerNetwork

`func (o *InlineResponse200208SentryEnrollment) SetSystemsManagerNetwork(v InlineResponse200208SentryEnrollmentSystemsManagerNetwork)`

SetSystemsManagerNetwork sets SystemsManagerNetwork field to given value.

### HasSystemsManagerNetwork

`func (o *InlineResponse200208SentryEnrollment) HasSystemsManagerNetwork() bool`

HasSystemsManagerNetwork returns a boolean if a field has been set.

### GetStrength

`func (o *InlineResponse200208SentryEnrollment) GetStrength() string`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *InlineResponse200208SentryEnrollment) GetStrengthOk() (*string, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *InlineResponse200208SentryEnrollment) SetStrength(v string)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *InlineResponse200208SentryEnrollment) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetEnforcedSystems

`func (o *InlineResponse200208SentryEnrollment) GetEnforcedSystems() []string`

GetEnforcedSystems returns the EnforcedSystems field if non-nil, zero value otherwise.

### GetEnforcedSystemsOk

`func (o *InlineResponse200208SentryEnrollment) GetEnforcedSystemsOk() (*[]string, bool)`

GetEnforcedSystemsOk returns a tuple with the EnforcedSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcedSystems

`func (o *InlineResponse200208SentryEnrollment) SetEnforcedSystems(v []string)`

SetEnforcedSystems sets EnforcedSystems field to given value.

### HasEnforcedSystems

`func (o *InlineResponse200208SentryEnrollment) HasEnforcedSystems() bool`

HasEnforcedSystems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


