package util

import (
	"strings"

	corev1 "k8s.io/api/core/v1"
)

func StringSliceToEnvVarSlice(envs []string) []corev1.EnvVar {
	envVars := []corev1.EnvVar{}

	for _, e := range envs {
		d := strings.Split(e, "=")
		envVars = append(envVars, corev1.EnvVar{Name: d[0], Value: d[1]})
	}
	return envVars
}

func StringSliceToMap(kv []string) map[string]string {
	m := map[string]string{}

	for _, e := range kv {
		d := strings.Split(e, "=")
		m[d[0]] = d[1]
	}
	return m
}
