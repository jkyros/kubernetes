/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v2beta1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"k8s.io/api/autoscaling/v2beta1"
	autoscalingv2beta1 "k8s.io/api/autoscaling/v2beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	autoscalingapiv2 "k8s.io/kubernetes/pkg/apis/autoscaling/v2"
)

// Testing nil pointer panic uncovered by #70806
// TODO(yue9944882): Test nil/empty conversion across all resource types
func TestNilOrEmptyConversion(t *testing.T) {
	scheme := runtime.NewScheme()
	assert.NoError(t, RegisterConversions(scheme))

	testCases := []struct {
		obj1 interface{}
		obj2 interface{}
	}{
		{
			obj1: &autoscaling.ExternalMetricSource{},
			obj2: &v2beta1.ExternalMetricSource{},
		},
		{
			obj1: &autoscaling.ExternalMetricStatus{},
			obj2: &v2beta1.ExternalMetricStatus{},
		},
		{
			obj1: &autoscaling.PodsMetricSource{},
			obj2: &v2beta1.PodsMetricSource{},
		},
		{
			obj1: &autoscaling.PodsMetricStatus{},
			obj2: &v2beta1.PodsMetricStatus{},
		},
		{
			obj1: &autoscaling.ObjectMetricSource{},
			obj2: &v2beta1.ObjectMetricSource{},
		},
		{
			obj1: &autoscaling.ObjectMetricStatus{},
			obj2: &v2beta1.ObjectMetricStatus{},
		},
		{
			obj1: &autoscaling.ResourceMetricSource{},
			obj2: &v2beta1.ResourceMetricSource{},
		},
		{
			obj1: &autoscaling.ResourceMetricStatus{},
			obj2: &v2beta1.ResourceMetricStatus{},
		},
		{
			obj1: &autoscaling.HorizontalPodAutoscaler{},
			obj2: &v2beta1.HorizontalPodAutoscaler{},
		},
		{
			obj1: &autoscaling.MetricTarget{},
			obj2: &v2beta1.CrossVersionObjectReference{},
		},
	}
	for _, testCase := range testCases {
		assert.NoError(t, scheme.Convert(testCase.obj1, testCase.obj2, nil))
		assert.NoError(t, scheme.Convert(testCase.obj2, testCase.obj1, nil))
	}
}

func TestConvert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(t *testing.T) {

	desiredUp := &autoscaling.HPAScalingRules{}
	desiredDown := &autoscaling.HPAScalingRules{}
	// we have to deal in terms of autoscaling but all our defaults are in autoscalingv2, so we have to convert them from there
	assert.NoError(t, autoscalingapiv2.Convert_v2_HPAScalingRules_To_autoscaling_HPAScalingRules(autoscalingapiv2.GenerateHPAScaleUpRules(nil), desiredUp, nil))
	assert.NoError(t, autoscalingapiv2.Convert_v2_HPAScalingRules_To_autoscaling_HPAScalingRules(autoscalingapiv2.GenerateHPAScaleDownRules(nil), desiredDown, nil))

	type args struct {
		out       *autoscaling.HorizontalPodAutoscaler
		in        *autoscalingv2beta1.HorizontalPodAutoscaler
		expectOut *autoscaling.HorizontalPodAutoscaler
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"TestConversionWithNilBehaviors",
			args{
				in: &autoscalingv2beta1.HorizontalPodAutoscaler{

					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{autoscaling.BehaviorSpecsAnnotation: "{}"},
					},
					Spec:   autoscalingv2beta1.HorizontalPodAutoscalerSpec{},
					Status: autoscalingv2beta1.HorizontalPodAutoscalerStatus{},
				},
				out: &autoscaling.HorizontalPodAutoscaler{},
				expectOut: &autoscaling.HorizontalPodAutoscaler{
					Spec: autoscaling.HorizontalPodAutoscalerSpec{
						Behavior: &autoscaling.HorizontalPodAutoscalerBehavior{
							ScaleUp:   desiredUp,
							ScaleDown: desiredDown,
						},
					},
					Status: autoscaling.HorizontalPodAutoscalerStatus{},
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.NotPanics(t, func() {
				if err := Convert_v2beta1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(tt.args.in, tt.args.out, nil); (err != nil) != tt.wantErr {
					t.Errorf("Convert_v2beta1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler() error = %v, wantErr %v", err, tt.wantErr)
				}
			})

			assert.Equal(t, tt.args.expectOut.Spec.Behavior, tt.args.out.Spec.Behavior)
		})
	}
}
