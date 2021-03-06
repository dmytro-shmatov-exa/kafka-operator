// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package currentalert

import (
	emperror "emperror.dev/errors"
)

type resizePvcValidator struct {
	Alert *currentAlertStruct
}

func newResizePvcValidator(curerentAlert *currentAlertStruct) resizePvcValidator {
	return resizePvcValidator{
		Alert: curerentAlert,
	}
}

func (a resizePvcValidator) validateAlert() error {
	if !checkLabelExists(a.Alert.Labels, "persistentvolumeclaim") {
		return emperror.New("persistentvolumeclaim label doesn't exist")
	}
	if a.Alert.Annotations["command"] != ResizePvcCommand {
		return emperror.NewWithDetails("unsupported command", "command", a.Alert.Annotations["command"])
	}

	return nil
}
