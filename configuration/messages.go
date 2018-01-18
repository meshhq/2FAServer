package configuration

// Success string literal.
const Success = "Success."

// KeyIDMissing string literal.
const KeyIDMissing = "Key identifier cannot be emtpy."

// KeyIDMustBeEmtpy string literal
const KeyIDMustBeEmtpy = "Key identifier must be emtpy for insertion."

// KeySecretMissing string literal
const KeySecretMissing = "Key secret cannot be emtpy."

// ProviderMissing string literal
const ProviderMissing = "Provider cannot be emtpy."

// UserIDMissing string literal.
const UserIDMissing = "Property 'user_id' is missing."

// InvalidRequestPayload string literal.
const InvalidRequestPayload = "Invalid request payload."

// ElementMissing string literal.
const ElementMissing = "Element does not exist."

// KeysFetchError string literal.
const KeysFetchError = "Could not retrieve Keys."

// DeleteError string literal.
const DeleteError = "Could not remove Key."

// UpdateKeyError string literal.
const UpdateKeyError = "Could not update Key."

// CreateKeyError string literal.
const CreateKeyError = "Could not create Key in storage."

// CreateOTPError string literal.
const CreateOTPError = "Could not create TOTP token."

// OTPValidationFailed string literal.
const OTPValidationFailed = "Token validation failed."

// CreateQRCodeError string literal.
const CreateQRCodeError = "Could not generate QR Code for enrollment."

// NoResutlsForUserIDAndProvider string literal.
const NoResutlsForUserIDAndProvider = "No key exist for given Provider and UserID association."
