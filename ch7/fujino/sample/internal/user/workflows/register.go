package userworkflows

import (
	shareerrs "example-ch7_8/internal/share/domain/errs"
	userdomain "example-ch7_8/internal/user/domain"
)

var NewRegisterUserWorkflow RegisterUser = func(
	validateUser ValidateUser,
	registUser RegistUser,
) RegisterUserWorkflow {
	return func(cmd RegisterUserCommand) ([]RegisterUserEvent, shareerrs.DomainValidationResult, error) {
		validated, res := validateUser(cmd.Data.UnvalidatedUser, cmd.Data.ExternalUserData)
		if !res.IsComplete() {
			return nil, res, nil
		}
		registered, err := registUser(*validated)
		if err != nil {
			return nil, nil, nil
		}
		return []RegisterUserEvent{&UserRegistered{RegisteredUser: *registered}}, nil, nil
	}
}

var ValidateUserImpl ValidateUser = func(uu userdomain.UnvalidatedUser, ext userdomain.ExternalUserData) (*userdomain.ValidatedUser, shareerrs.DomainValidationResult) {
	var result shareerrs.DomainValidationResult
	toValidateUser := userdomain.ToValidateUserImpl(ext)
	validated, res := toValidateUser(uu)
	result = result.Merge(res)
	return validated, result
}

var RegistUserImpl RegistUser = func(user userdomain.ValidatedUser) (*userdomain.RegisteredUser, error) {

	return &userdomain.RegisteredUser{}, nil
}
