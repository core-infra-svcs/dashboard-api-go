# InlineResponse200243OptOutEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eligible** | Pointer to **bool** | Condition flag to opt out from the feature | [optional] 
**Reason** | Pointer to **string** | User friendly message regarding opt-out eligibility | [optional] 
**Help** | Pointer to [**InlineResponse200243OptOutEligibilityHelp**](InlineResponse200243OptOutEligibilityHelp.md) |  | [optional] 

## Methods

### NewInlineResponse200243OptOutEligibility

`func NewInlineResponse200243OptOutEligibility() *InlineResponse200243OptOutEligibility`

NewInlineResponse200243OptOutEligibility instantiates a new InlineResponse200243OptOutEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200243OptOutEligibilityWithDefaults

`func NewInlineResponse200243OptOutEligibilityWithDefaults() *InlineResponse200243OptOutEligibility`

NewInlineResponse200243OptOutEligibilityWithDefaults instantiates a new InlineResponse200243OptOutEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligible

`func (o *InlineResponse200243OptOutEligibility) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *InlineResponse200243OptOutEligibility) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *InlineResponse200243OptOutEligibility) SetEligible(v bool)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *InlineResponse200243OptOutEligibility) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetReason

`func (o *InlineResponse200243OptOutEligibility) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InlineResponse200243OptOutEligibility) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InlineResponse200243OptOutEligibility) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InlineResponse200243OptOutEligibility) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetHelp

`func (o *InlineResponse200243OptOutEligibility) GetHelp() InlineResponse200243OptOutEligibilityHelp`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *InlineResponse200243OptOutEligibility) GetHelpOk() (*InlineResponse200243OptOutEligibilityHelp, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *InlineResponse200243OptOutEligibility) SetHelp(v InlineResponse200243OptOutEligibilityHelp)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *InlineResponse200243OptOutEligibility) HasHelp() bool`

HasHelp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


