{{- define "k8s-operator.fullname" -}}
{{- if eq .Release.Name .Chart.Name -}}
{{ .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else -}}
{{ printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" }}
{{- end -}}
{{- end }}

{{- define "ukc.metro" -}}
{{- required "UKC metro is required." .Values.ukc.metro -}}
{{- end }}

{{- define "ukc.metroSlug" -}}
{{- (urlParse (include "ukc.metro" . | trim)).host | replace "." "-" -}}
{{- end }}

{{- define "k8s-operator.configmap-name" -}}
{{ include "k8s-operator.fullname" . }}
{{- end }}

{{- define "k8s-operator.secret-name" -}}
{{ include "k8s-operator.fullname" . }}
{{- end }}

{{- define "k8s-operator.leader-election-role-name" -}}
{{ printf "%s-leader-election" (include "k8s-operator.fullname" .) }}
{{- end }}

{{- define "k8s-operator.image" -}}
{{- printf "%s:%s" .Values.image.name (default "latest" .Values.image.tag) -}}
{{- end }}
