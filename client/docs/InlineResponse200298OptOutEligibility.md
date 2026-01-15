# InlineResponse200298OptOutEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eligible** | Pointer to **bool** | Condition flag to opt out from the feature | [optional] 
**Reason** | Pointer to **string** | User friendly message regarding opt-out eligibility | [optional] 
**Help** | Pointer to [**InlineResponse200298OptOutEligibilityHelp**](InlineResponse200298OptOutEligibilityHelp.md) |  | [optional] 

## Methods

### NewInlineResponse200298OptOutEligibility

`func NewInlineResponse200298OptOutEligibility() *InlineResponse200298OptOutEligibility`

NewInlineResponse200298OptOutEligibility instantiates a new InlineResponse200298OptOutEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200298OptOutEligibilityWithDefaults

`func NewInlineResponse200298OptOutEligibilityWithDefaults() *InlineResponse200298OptOutEligibility`

NewInlineResponse200298OptOutEligibilityWithDefaults instantiates a new InlineResponse200298OptOutEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligible

`func (o *InlineResponse200298OptOutEligibility) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *InlineResponse200298OptOutEligibility) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *InlineResponse200298OptOutEligibility) SetEligible(v bool)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *InlineResponse200298OptOutEligibility) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetReason

`func (o *InlineResponse200298OptOutEligibility) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InlineResponse200298OptOutEligibility) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InlineResponse200298OptOutEligibility) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InlineResponse200298OptOutEligibility) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetHelp

`func (o *InlineResponse200298OptOutEligibility) GetHelp() InlineResponse200298OptOutEligibilityHelp`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *InlineResponse200298OptOutEligibility) GetHelpOk() (*InlineResponse200298OptOutEligibilityHelp, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *InlineResponse200298OptOutEligibility) SetHelp(v InlineResponse200298OptOutEligibilityHelp)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *InlineResponse200298OptOutEligibility) HasHelp() bool`

HasHelp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


