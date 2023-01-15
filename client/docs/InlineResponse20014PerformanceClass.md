# InlineResponse20014PerformanceClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of this performance class. Must be one of: &#39;builtin&#39; or &#39;custom&#39; | 
**BuiltinPerformanceClassName** | Pointer to **string** | Name of builtin performance class. Must be present when performanceClass type is &#39;builtin&#39; and value must be one of: &#39;VoIP&#39; | [optional] 
**CustomPerformanceClassId** | Pointer to **string** | ID of created custom performance class, must be present when performanceClass type is \&quot;custom\&quot; | [optional] 

## Methods

### NewInlineResponse20014PerformanceClass

`func NewInlineResponse20014PerformanceClass(type_ string, ) *InlineResponse20014PerformanceClass`

NewInlineResponse20014PerformanceClass instantiates a new InlineResponse20014PerformanceClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20014PerformanceClassWithDefaults

`func NewInlineResponse20014PerformanceClassWithDefaults() *InlineResponse20014PerformanceClass`

NewInlineResponse20014PerformanceClassWithDefaults instantiates a new InlineResponse20014PerformanceClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse20014PerformanceClass) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20014PerformanceClass) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20014PerformanceClass) SetType(v string)`

SetType sets Type field to given value.


### GetBuiltinPerformanceClassName

`func (o *InlineResponse20014PerformanceClass) GetBuiltinPerformanceClassName() string`

GetBuiltinPerformanceClassName returns the BuiltinPerformanceClassName field if non-nil, zero value otherwise.

### GetBuiltinPerformanceClassNameOk

`func (o *InlineResponse20014PerformanceClass) GetBuiltinPerformanceClassNameOk() (*string, bool)`

GetBuiltinPerformanceClassNameOk returns a tuple with the BuiltinPerformanceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltinPerformanceClassName

`func (o *InlineResponse20014PerformanceClass) SetBuiltinPerformanceClassName(v string)`

SetBuiltinPerformanceClassName sets BuiltinPerformanceClassName field to given value.

### HasBuiltinPerformanceClassName

`func (o *InlineResponse20014PerformanceClass) HasBuiltinPerformanceClassName() bool`

HasBuiltinPerformanceClassName returns a boolean if a field has been set.

### GetCustomPerformanceClassId

`func (o *InlineResponse20014PerformanceClass) GetCustomPerformanceClassId() string`

GetCustomPerformanceClassId returns the CustomPerformanceClassId field if non-nil, zero value otherwise.

### GetCustomPerformanceClassIdOk

`func (o *InlineResponse20014PerformanceClass) GetCustomPerformanceClassIdOk() (*string, bool)`

GetCustomPerformanceClassIdOk returns a tuple with the CustomPerformanceClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPerformanceClassId

`func (o *InlineResponse20014PerformanceClass) SetCustomPerformanceClassId(v string)`

SetCustomPerformanceClassId sets CustomPerformanceClassId field to given value.

### HasCustomPerformanceClassId

`func (o *InlineResponse20014PerformanceClass) HasCustomPerformanceClassId() bool`

HasCustomPerformanceClassId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


