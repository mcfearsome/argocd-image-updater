package common

// This file contains a list of constants required by other packages

const ImageUpdaterAnnotationPrefix = "argocd-image-updater.argoproj.io"

// The annotation on the application resources to indicate the list of images
// allowed for updates.
const ImageUpdaterAnnotation = ImageUpdaterAnnotationPrefix + "/image-list"

// Defaults for Helm parameter names
const (
	DefaultHelmImageName = "image.name"
	DefaultHelmImageTag  = "image.tag"
)

// Helm related annotations
const (
	HelmParamImageNameAnnotation = ImageUpdaterAnnotationPrefix + "/%s.helm.image-name"
	HelmParamImageTagAnnotation  = ImageUpdaterAnnotationPrefix + "/%s.helm.image-tag"
	HelmParamImageSpecAnnotation = ImageUpdaterAnnotationPrefix + "/%s.helm.image-spec"
)

// Kustomize related annotations
const (
	KustomizeApplicationNameAnnotation = ImageUpdaterAnnotationPrefix + "/%s.kustomize.image-name"
)

// Template related annotations
const ()

// Image specific configuration annotations
const (
	OldMatchOptionAnnotation    = ImageUpdaterAnnotationPrefix + "/%s.tag-match" // Deprecated and will be removed
	AllowTagsOptionAnnotation   = ImageUpdaterAnnotationPrefix + "/%s.allow-tags"
	IgnoreTagsOptionAnnotation  = ImageUpdaterAnnotationPrefix + "/%s.ignore-tags"
	ForceUpdateOptionAnnotation = ImageUpdaterAnnotationPrefix + "/%s.force-update"
	UpdateStrategyAnnotation    = ImageUpdaterAnnotationPrefix + "/%s.update-strategy"
	PullSecretAnnotation        = ImageUpdaterAnnotationPrefix + "/%s.pull-secret"
	PlatformsAnnotation         = ImageUpdaterAnnotationPrefix + "/%s.platforms"
)

// Application-wide update strategy related annotations
const (
	ApplicationWideAllowTagsOptionAnnotation   = ImageUpdaterAnnotationPrefix + "/allow-tags"
	ApplicationWideIgnoreTagsOptionAnnotation  = ImageUpdaterAnnotationPrefix + "/ignore-tags"
	ApplicationWideForceUpdateOptionAnnotation = ImageUpdaterAnnotationPrefix + "/force-update"
	ApplicationWideUpdateStrategyAnnotation    = ImageUpdaterAnnotationPrefix + "/update-strategy"
	ApplicationWidePullSecretAnnotation        = ImageUpdaterAnnotationPrefix + "/pull-secret"
)

// Application update configuration related annotations
const (
	WriteBackMethodAnnotation   = ImageUpdaterAnnotationPrefix + "/write-back-method"
	GitBranchAnnotation         = ImageUpdaterAnnotationPrefix + "/git-branch"
	WriteBackTargetAnnotation   = ImageUpdaterAnnotationPrefix + "/write-back-target"
	WriteBackTemplateAnnotation = ImageUpdaterAnnotationPrefix + "/write-back-template"
	// For Internal Use
	WriteBackTemplateBuildCacheAnnotation = ImageUpdaterAnnotationPrefix + "write-back-template-cache"
	KustomizationPrefix                   = "kustomization"
	HelmPrefix                            = "helm"
	TemplatePrefix                        = "template"
)

// DefaultTargetFilePattern configurations related to the write-back functionality
const DefaultTargetFilePattern = ".argocd-source-%s.yaml"

// The default Git commit message's template
const DefaultGitCommitMessage = `build: automatic update of {{ .AppName }}

{{ range .AppChanges -}}
updates image {{ .Image }} tag '{{ .OldTag }}' to '{{ .NewTag }}'
{{ end -}}
`

const DefaultTemplate = `
image:
{{- if ne .Index 0 }}
  {{ .Alias }}:
    repository: {{ .Repository }}
    tag: {{ .Tag }}
{{ else }}
  repository: {{ .Repository }}
  tag: {{ .Tag }}
{{ end }}
`
