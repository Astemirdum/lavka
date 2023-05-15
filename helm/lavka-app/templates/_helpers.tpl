{{/*
Expand the name of the chart.
*/}}
{{- define "lavka-app.name" -}}
{{- default .Chart.Name .Values.app.name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "lavka-app.fullname" -}}
{{- if .Values.app.fullname }}
{{- .Values.app.fullname | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.app.name }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}


{{- define "db.name" -}}
{{- if .Values.db.name }}
{{- .Values.db.name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.app.name }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-"  }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "lavka-app.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}



{{/*
Common labels
*/}}
{{- define "lavka-app.labels" -}}
helm.sh/chart: {{ include "lavka-app.chart" . }}
{{ include "lavka-app.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{/*{{- if .Values.version }}*/}}
{{/*app.kubernetes.io/version: {{ .Values.version | quote }}*/}}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}


{{- define "db.labels" -}}
helm.sh/chart: {{ include "lavka-app.chart" . }}
{{ include "db.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Expand the namespace.
*/}}
{{- define "lavka-app.namespace" -}}
{{- default .Values.namespace .Release.Namespace | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "db.namespace" -}}
{{- default .Values.db.namespace .Release.Namespace | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "lavka-app.selectorLabels" -}}
app.kubernetes.io/name: {{ include "lavka-app.fullname" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}


{{- define "db.selectorLabels" -}}
app.kubernetes.io/name: {{ include "db.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "lavka-app.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "lavka-app.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*app strategy*/}}
{{- define "lavka-app.strategy" -}}
rollingUpdate:
  maxSurge: {{ .Values.app.strategy.rollingUpdate.maxSurge}}
  maxUnavailable: {{ .Values.app.strategy.rollingUpdate.maxUnavailable}}
type: {{ .Values.app.strategy.type}}
{{- end }}


{{/*health*/}}
{{- define "lavka-app.health" -}}
readinessProbe:
  httpGet: &health
    path: /health
    port: {{ .Values.configData.http.port }}
    scheme: HTTP
  initialDelaySeconds: 20
  failureThreshold: 3
  periodSeconds: 30
  timeoutSeconds: 5
livenessProbe:
  httpGet: *health
  failureThreshold: 5
  periodSeconds: 60
  timeoutSeconds: 5
  successThreshold: 1
  initialDelaySeconds: 10
startupProbe:
  failureThreshold: 10
  httpGet: *health
  periodSeconds: 10
  timeoutSeconds: 5
  successThreshold: 1
{{- end }}

{{- define "db.health" -}}
livenessProbe:
  exec:
    command:
      - bash
      - -ec
      - 'PGPASSWORD=$POSTGRES_PASSWORD psql -w -U "${POSTGRES_USER}" -d "${POSTGRES_DB}"  -h 127.0.0.1 -c "SELECT 1"'
  initialDelaySeconds: 30
  periodSeconds: 10
  timeoutSeconds: 5
  successThreshold: 1
  failureThreshold: 6
readinessProbe:
  exec:
    command:
      - bash
      - -ec
      - 'PGPASSWORD=$POSTGRES_PASSWORD psql -w -U "${POSTGRES_USER}" -d "${POSTGRES_DB}"  -h 127.0.0.1 -c "SELECT 1"'
  initialDelaySeconds: 5
  periodSeconds: 10
  timeoutSeconds: 5
  successThreshold: 1
  failureThreshold: 6
{{- end}}


{{/*
env pgHostPortDB
*/}}
{{- define "app.env.pgHostPortDB" -}}
- name: DB_HOST
  valueFrom:
    configMapKeyRef:
      name: {{ include "lavka-app.fullname" . }}-config
      key: DB_HOST
- name: DB_PORT
  valueFrom:
    configMapKeyRef:
      name: {{ include "lavka-app.fullname" . }}-config
      key: DB_PORT
{{- end }}


{{/*
Wait for DB to be ready
*/}}
{{- define "app.pgWait" -}}
['sh', '-c', "until nc -w 2 $(DB_HOST) $(DB_PORT); do echo Waiting for $(DB_HOST):$(DB_PORT) to be ready; sleep 5; done"]
{{- end }}

{{/*
Change how the image is assigned based on the skaffold flag.
*/}}
{{- define "k8-golang-demo.image" -}}
{{- if .Values.skaffold -}}
{{- .Values.skaffoldImage -}}
{{- else -}}
{{- printf "%s%s:%s" .Values.image.hostname .Values.image.repository .Values.image.tag -}}
{{- end -}}
{{- end -}}

{{/*
Change how the data migration image is assigned based on the skaffold flag.
*/}}
{{- define "k8-golang-demo.migration-image" -}}
{{- if .Values.skaffold -}}
{{- .Values.migration.skaffoldImage -}}
{{- else -}}
{{- printf "%s%s:%s" .Values.migration.image.hostname .Values.migration.image.repository .Values.migration.image.tag -}}
{{- end -}}
{{- end -}}