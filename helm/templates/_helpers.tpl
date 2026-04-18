{{- define "hnrss.middlewareAuthName" -}}
{{ printf "%s-basic-auth" .Release.Name }}
{{- end }}

{{- define "hnrss.serversTransportName" -}}
{{ printf "%s-transport" .Release.Name }}
{{- end }}

{{- define "hnrss.middlewareAuthRef" -}}
{{ printf "%s-%s@kubernetescrd" .Release.Namespace (include "hnrss.middlewareAuthName" .) }}
{{- end }}

{{- define "hnrss.serversTransportRef" -}}
{{ printf "%s-%s@kubernetescrd" .Release.Namespace (include "hnrss.serversTransportName" .) }}
{{- end }}

{{- define "hnrss.dockerconfigjson" -}}
{{- $server := .Values.registryAuth.server -}}
{{- $user := .Values.registryAuth.username -}}
{{- $pass := .Values.registryAuth.password -}}
{{- $auth := printf "%s:%s" $user $pass | b64enc -}}
{{- $entry := dict "username" $user "password" $pass "auth" $auth -}}
{{- dict "auths" (dict $server $entry) | toJson -}}
{{- end }}
