// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package runtime_env

import (
	"fmt"

	"net/url"

	"openpitrix.io/openpitrix/pkg/models"
)

func LabelStringToMap(labelString string) (map[string]string, error) {
	mapLabel := make(map[string]string)
	m, err := url.ParseQuery(labelString)
	if err != nil {
		return nil, err
	}
	for mKey, mValue := range m {
		if len(mValue) != 1 {
			return nil, fmt.Errorf("bad label format %v", labelString)
		}
		mapLabel[mKey] = mValue[0]
	}
	return mapLabel, nil
}

func SelectorStringToMap(selectorString string) (map[string][]string, error) {
	selectorMap, err := url.ParseQuery(selectorString)
	if err != nil {
		return nil, err
	}
	return selectorMap, nil
}

func LabelMapDiff(oldLabelMap, newLabelMap map[string]string) (additions, deletions map[string]string) {
	additions = make(map[string]string)
	deletions = make(map[string]string)
	for i := 0; i < 2; i++ {
		for oldLabelKey, oldLabelValue := range oldLabelMap {
			found := false
			if newLabelValue, ok := newLabelMap[oldLabelKey]; ok {
				if oldLabelValue == newLabelValue {
					found = true
				}
			}
			if !found {
				if i == 0 {
					deletions[oldLabelKey] = oldLabelValue
				} else {
					additions[oldLabelKey] = oldLabelValue
				}
			}
		}
		if i == 0 {
			oldLabelMap, newLabelMap = newLabelMap, oldLabelMap
		}
	}
	return additions, deletions
}

func LabelStructToMap(labelStructs []*models.RuntimeEnvLabel) map[string]string {
	mapLabel := make(map[string]string)
	for _, labelStruct := range labelStructs {
		mapLabel[labelStruct.LabelKey] = labelStruct.LabelValue
	}
	return mapLabel
}
