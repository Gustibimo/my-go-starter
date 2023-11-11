package pkg

import "my-go-starter/shared"

func MaskUser(user shared.UserType) shared.UserType {
	var userMasked shared.UserType

	// mask 3 first character
	userMasked.Name = "***" + user.Name[3:]
	userMasked.Email = user.Email

	return userMasked
}