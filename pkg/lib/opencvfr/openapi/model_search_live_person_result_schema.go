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

// checks if the SearchLivePersonResultSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchLivePersonResultSchema{}

// SearchLivePersonResultSchema struct for SearchLivePersonResultSchema
type SearchLivePersonResultSchema struct {
	LivenessScore float32                    `json:"liveness_score"`
	Persons       []SearchPersonResultSchema `json:"persons"`
}

type _SearchLivePersonResultSchema SearchLivePersonResultSchema

// NewSearchLivePersonResultSchema instantiates a new SearchLivePersonResultSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchLivePersonResultSchema(livenessScore float32, persons []SearchPersonResultSchema) *SearchLivePersonResultSchema {
	this := SearchLivePersonResultSchema{}
	this.LivenessScore = livenessScore
	this.Persons = persons
	return &this
}

// NewSearchLivePersonResultSchemaWithDefaults instantiates a new SearchLivePersonResultSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchLivePersonResultSchemaWithDefaults() *SearchLivePersonResultSchema {
	this := SearchLivePersonResultSchema{}
	return &this
}

// GetLivenessScore returns the LivenessScore field value
func (o *SearchLivePersonResultSchema) GetLivenessScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LivenessScore
}

// GetLivenessScoreOk returns a tuple with the LivenessScore field value
// and a boolean to check if the value has been set.
func (o *SearchLivePersonResultSchema) GetLivenessScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LivenessScore, true
}

// SetLivenessScore sets field value
func (o *SearchLivePersonResultSchema) SetLivenessScore(v float32) {
	o.LivenessScore = v
}

// GetPersons returns the Persons field value
func (o *SearchLivePersonResultSchema) GetPersons() []SearchPersonResultSchema {
	if o == nil {
		var ret []SearchPersonResultSchema
		return ret
	}

	return o.Persons
}

// GetPersonsOk returns a tuple with the Persons field value
// and a boolean to check if the value has been set.
func (o *SearchLivePersonResultSchema) GetPersonsOk() ([]SearchPersonResultSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Persons, true
}

// SetPersons sets field value
func (o *SearchLivePersonResultSchema) SetPersons(v []SearchPersonResultSchema) {
	o.Persons = v
}

func (o SearchLivePersonResultSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchLivePersonResultSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["liveness_score"] = o.LivenessScore
	toSerialize["persons"] = o.Persons
	return toSerialize, nil
}

func (o *SearchLivePersonResultSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"liveness_score",
		"persons",
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

	varSearchLivePersonResultSchema := _SearchLivePersonResultSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchLivePersonResultSchema)

	if err != nil {
		return err
	}

	*o = SearchLivePersonResultSchema(varSearchLivePersonResultSchema)

	return err
}

type NullableSearchLivePersonResultSchema struct {
	value *SearchLivePersonResultSchema
	isSet bool
}

func (v NullableSearchLivePersonResultSchema) Get() *SearchLivePersonResultSchema {
	return v.value
}

func (v *NullableSearchLivePersonResultSchema) Set(val *SearchLivePersonResultSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchLivePersonResultSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchLivePersonResultSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchLivePersonResultSchema(val *SearchLivePersonResultSchema) *NullableSearchLivePersonResultSchema {
	return &NullableSearchLivePersonResultSchema{value: val, isSet: true}
}

func (v NullableSearchLivePersonResultSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchLivePersonResultSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}