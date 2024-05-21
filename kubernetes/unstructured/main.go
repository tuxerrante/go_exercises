package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func main() {

	myObject := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"status": map[string]interface{}{
				"binding": map[string]interface{}{
					"name": "K8SCustomResource",
				},
				"conditions": []map[string]interface{}{
					{
						"reason":  "SecretSynced",
						"message": "Secret was synced",
					},
					{
						"reason":  "SecretNotSynced",
						"message": "Secret was not synced",
					},
				},
			},
		},
	}

	conditions, found, err := unstructured.NestedSlice(myObject.Object, "status", "conditions")
	checkUnstructuredSearchResult("myObject", found, "conditions", conditions, err)

	conditionMap, ok := conditions[0].(map[string]interface{})
	if !ok {
		panic("cannot convert condition field to map[string]interface{}")
	}
	reason, found, err := unstructured.NestedString(conditionMap, "reason")
	checkUnstructuredSearchResult("reason", found, "reason", reason, err)

	fmt.Println(reason)
}

func checkUnstructuredSearchResult(resourceName string, found bool, keyword string, foundObject any, err error) {
	if !found {
		panic(fmt.Sprintf("'%s' field not found in %s: Got:{{%v}} -- %v.", keyword, resourceName, foundObject, err))
	}
}
