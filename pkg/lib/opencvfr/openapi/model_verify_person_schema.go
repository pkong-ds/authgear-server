/*
OpenCV Face Recognition

 OpenCV Face Recognition allows you to either **manually** (through our [Developer Portal](https://developer.opencv.fr) UI) or **programmatically** (through our SDKs or REST API) detect, recognize, and verify faces in images. It is based on state-of-the-art (SOTA) algorithms and uses deep learning face recognition models. The API is designed to be easy to use and integrate into your applications. It is available as this hosted service or you can deploy it yourself on your own servers.  There are four ways that you can use the product: - Using the [Face Recognition Developer Portal](https://developer.opencv.fr) User Interface to manage your **Developer** teams, **Persons** who are registered for the search API to recognize, and **Collections** - groups of persons. - Using the [Python Face Recognition SDK](https://docs.opencv.fr/python) to integrate the API into your Python applications. - Using the [C++ Face Recognition SDK](https://docs.opencv.fr/cpp) to integrate the API into your C++ applications. - Using the REST API (below) to integrate the API functionality into your applications in other languages.  To use the REST API described below, you will need to create an account and obtain an API key. You can do this by signing up for a free account at [Face Recognition Developer Portal](https://developer.opencv.fr).  Once you have signed up, you will see an **API Developer Key** in the Dashboard. This is the key that you will use to authenticate your requests to the API. You can also create additional Developers (each with their own key) for your applications.  To use the API, you will need to send the API key in the `X-API-Key` header of each request. For example, using `curl`: ``` curl -X GET \"https://<region>.opencv.fr/persons\" -H \"accept: application/json\" -H \"X-API-Key: <your API key>\" ```  `<region>` is the data storage region that you selected when you created your account. It can be `us`, `eu`, or `sg`.  To help you try out the functionality quickly, the below live docs include a **Try it out** button for each endpoint. This will allow you to send a request to the API and see the response. Before you can use this, you will need to grab your API key from the Dashboard and enter it into the field that shows up when you click the green **Authorize** button (below this line on the right).

API version: 2024.07.05.1135
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const DefaultMinScore float32 = 0.7

// checks if the VerifyPersonSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyPersonSchema{}

// VerifyPersonSchema struct for VerifyPersonSchema
type VerifyPersonSchema struct {
	Id               string                 `json:"id"`
	Images           []string               `json:"images"`
	MinScore         NullableFloat32        `json:"min_score,omitempty"`
	VerificationMode NullableSearchModeEnum `json:"verification_mode,omitempty"`
}

type _VerifyPersonSchema VerifyPersonSchema

// NewVerifyPersonSchema instantiates a new VerifyPersonSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyPersonSchema(id string, images []string) *VerifyPersonSchema {
	this := VerifyPersonSchema{}
	this.Id = id
	this.Images = images
	this.MinScore = *NewNullableFloat32(PtrFloat32(DefaultMinScore))
	return &this
}

// NewVerifyPersonSchemaWithDefaults instantiates a new VerifyPersonSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyPersonSchemaWithDefaults() *VerifyPersonSchema {
	this := VerifyPersonSchema{}
	this.MinScore = *NewNullableFloat32(PtrFloat32(DefaultMinScore))
	return &this
}

// GetId returns the Id field value
func (o *VerifyPersonSchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VerifyPersonSchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VerifyPersonSchema) SetId(v string) {
	o.Id = v
}

// GetImages returns the Images field value
func (o *VerifyPersonSchema) GetImages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *VerifyPersonSchema) GetImagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *VerifyPersonSchema) SetImages(v []string) {
	o.Images = v
}

// GetMinScore returns the MinScore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyPersonSchema) GetMinScore() float32 {
	if o == nil || IsNil(o.MinScore.Get()) {
		var ret float32
		return ret
	}
	return *o.MinScore.Get()
}

// GetMinScoreOk returns a tuple with the MinScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyPersonSchema) GetMinScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinScore.Get(), o.MinScore.IsSet()
}

// HasMinScore returns a boolean if a field has been set.
func (o *VerifyPersonSchema) HasMinScore() bool {
	if o != nil && o.MinScore.IsSet() {
		return true
	}

	return false
}

// SetMinScore gets a reference to the given NullableMinScore and assigns it to the MinScore field.
func (o *VerifyPersonSchema) SetMinScore(v float32) {
	o.MinScore.Set(&v)
}

// SetMinScoreNil sets the value for MinScore to be an explicit nil
func (o *VerifyPersonSchema) SetMinScoreNil() {
	o.MinScore.Set(nil)
}

// UnsetMinScore ensures that no value is present for MinScore, not even an explicit nil
func (o *VerifyPersonSchema) UnsetMinScore() {
	o.MinScore.Unset()
}

// GetVerificationMode returns the VerificationMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyPersonSchema) GetVerificationMode() SearchModeEnum {
	if o == nil || IsNil(o.VerificationMode.Get()) {
		var ret SearchModeEnum
		return ret
	}
	return *o.VerificationMode.Get()
}

// GetVerificationModeOk returns a tuple with the VerificationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyPersonSchema) GetVerificationModeOk() (*SearchModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerificationMode.Get(), o.VerificationMode.IsSet()
}

// HasVerificationMode returns a boolean if a field has been set.
func (o *VerifyPersonSchema) HasVerificationMode() bool {
	if o != nil && o.VerificationMode.IsSet() {
		return true
	}

	return false
}

// SetVerificationMode gets a reference to the given NullableSearchModeEnum and assigns it to the VerificationMode field.
func (o *VerifyPersonSchema) SetVerificationMode(v SearchModeEnum) {
	o.VerificationMode.Set(&v)
}

// SetVerificationModeNil sets the value for VerificationMode to be an explicit nil
func (o *VerifyPersonSchema) SetVerificationModeNil() {
	o.VerificationMode.Set(nil)
}

// UnsetVerificationMode ensures that no value is present for VerificationMode, not even an explicit nil
func (o *VerifyPersonSchema) UnsetVerificationMode() {
	o.VerificationMode.Unset()
}

func (o VerifyPersonSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyPersonSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["images"] = o.Images
	if o.MinScore.IsSet() {
		toSerialize["min_score"] = o.MinScore.Get()
	}
	if o.VerificationMode.IsSet() {
		toSerialize["verification_mode"] = o.VerificationMode.Get()
	}
	return toSerialize, nil
}

func (o *VerifyPersonSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"images",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVerifyPersonSchema := _VerifyPersonSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerifyPersonSchema)

	if err != nil {
		return err
	}

	*o = VerifyPersonSchema(varVerifyPersonSchema)

	return err
}

type NullableVerifyPersonSchema struct {
	value *VerifyPersonSchema
	isSet bool
}

func (v NullableVerifyPersonSchema) Get() *VerifyPersonSchema {
	return v.value
}

func (v *NullableVerifyPersonSchema) Set(val *VerifyPersonSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyPersonSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyPersonSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyPersonSchema(val *VerifyPersonSchema) *NullableVerifyPersonSchema {
	return &NullableVerifyPersonSchema{value: val, isSet: true}
}

func (v NullableVerifyPersonSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyPersonSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}