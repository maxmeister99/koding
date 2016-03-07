// A series of kite error type constants. Because the usage of the kite error
// types can be used anywhere, these types should never be changed.
// Incosistencies in formatting likely means that the type came from somewhere
// else. New types should adhere to camelcase.
//
// Leave existing cases as is!
//
package kiteerrortypes

const (
	//
	// Kite error types not generated from Klient:
	//

	AuthenticationError = "authenticationError"

	//
	// Kite error types generated from klient:
	//

	MachineNotFound           = "MachineNotFound"
	MachineUnreachable        = "MachineUnreachable"
	AuthErrTokenIsExpired     = "AuthErrTokenIsExpired"
	AuthErrTokenIsNotValidYet = "AuthErrTokenIsNotValidYet"
	MissingArgument           = "MissingArgument"

	// DialingFailed is the kite.Error.Type used for errors encountered when
	// dialing the remote.
	DialingFailed = "dialing failed"

	// MountNotFound is the kite.Error.Type used for errors encountered when
	// the mount name given cannot be found.
	MountNotFound = "mount not found"

	// SystemUnmountFailed is the Kite Error Type used when either the unmounter
	// fails, or generic (non-instanced) unmount fails.
	SystemUnmountFailed = "system-unmount-failed"
)