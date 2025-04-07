package shareerrs

// ValidationError はバリデーションエラーを表す構造体です。
// フィールド名とエラーを保持します。
type validationError struct {
	Field string
	Err   error
}

// ValidationErrors は複数のバリデーションエラーを保持するためのスライスです。
type ValidationErrors []validationError

func (errs ValidationErrors) Error() string {
	if len(errs) == 0 {
		return ""
	}

	var result string
	for _, err := range errs {
		result += err.Field + ": " + err.Err.Error() + "\n"
	}
	return result
}
func (errs ValidationErrors) IsEmpty() bool {
	return len(errs) == 0
}
func (errs ValidationErrors) Add(
	field string,
	err error,
) ValidationErrors {
	return append(errs, validationError{Field: field, Err: err})
}

// ドメインバリデーションの結果を表す構造体
type DomainValidationResult interface {
	ExternalDataRequests() []ExternalDataRequest
	ValidationErrors() ValidationErrors
	IsComplete() bool
	HasRequest() bool
	HasError() bool
	Merge(DomainValidationResult) DomainValidationResult
}

type domainValidationResult struct {
	externalDataRequests []ExternalDataRequest
	validationErrors     ValidationErrors
}

func NewDomainValidationResult(
	externalDataRequests []ExternalDataRequest,
	validationErrors ValidationErrors,
) DomainValidationResult {
	return &domainValidationResult{
		externalDataRequests: externalDataRequests,
		validationErrors:     validationErrors,
	}
}

func (v domainValidationResult) IsComplete() bool {

	return v.validationErrors.IsEmpty() && len(v.externalDataRequests) == 0
}

func (v domainValidationResult) HasRequest() bool {
	return len(v.externalDataRequests) > 0
}
func (v domainValidationResult) HasError() bool {
	return !v.validationErrors.IsEmpty()
}

func (v *domainValidationResult) Merge(
	other DomainValidationResult,
) DomainValidationResult {
	if v == nil {
		return other
	}
	return &domainValidationResult{
		externalDataRequests: append(v.externalDataRequests, other.ExternalDataRequests()...),
		validationErrors:     append(v.validationErrors, other.ValidationErrors()...),
	}
}

func (v domainValidationResult) ExternalDataRequests() []ExternalDataRequest {
	return v.externalDataRequests
}
func (v domainValidationResult) ValidationErrors() ValidationErrors {
	return v.validationErrors
}

// 外部データリクエストのインターフェース
// ワークフローをピュアに保つために、
// バリデーション時に依存データがある場合は外にリクエストをだすように設計しています
// 例) 重複チェック, 外部APIの呼び出しなど
type ExternalDataRequest interface {
	Key() string
	Description() string
}

type ExternalDataType int
