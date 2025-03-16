package userdomain

import shareerrs "example-ch7_8/internal/share/domain/errs"

var ToValidateUserImpl ToValidateUser = func(ext ExternalUserData) func(UnvalidatedUser) (*ValidatedUser, shareerrs.DomainValidationResult) {
	validateEmail := toValidateEmailImpl(
		toFormattedEmailImpl,
		toUniqueEmailImpl,
		ext.ExternalEmailData,
	)
	return func(user UnvalidatedUser) (*ValidatedUser, shareerrs.DomainValidationResult) {
		var result shareerrs.DomainValidationResult
		email, res := validateEmail(user.Email)
		result = result.Merge(res)

		if (result != nil) && !result.IsComplete() {
			return nil, result
		}

		return &ValidatedUser{
			FirstName:      FirstName(user.FirstName),
			LastName:       LastName(user.LastName),
			Email:          email,
			Password:       Password(user.Password),
			Zipcode:        Zipcode(user.Zipcode),
			Prefecture:     Prefecture(user.Prefecture),
			Municipalities: Municipalities(user.Municipalities),
			Address:        Address(user.Address),
			Telephone:      Telephone(user.Telephone),
		}, result
	}

}
