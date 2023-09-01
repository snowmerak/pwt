package pwt

import "errors"

var errInvalidPayloadLength = errors.New("invalid payload length")

func IsInvalidPayloadLengthErr(err error) bool {
	return errors.Is(err, errInvalidPayloadLength)
}

var errInvalidPayloadType = errors.New("invalid payload type")

func IsInvalidPayloadTypeErr(err error) bool {
	return errors.Is(err, errInvalidPayloadType)
}

var errAlreadySetAnotherType = errors.New("already set another type")

func IsAlreadySetAnotherTypeErr(err error) bool {
	return errors.Is(err, errAlreadySetAnotherType)
}

var errInvalidSignature = errors.New("invalid signature")

func IsInvalidSignatureErr(err error) bool {
	return errors.Is(err, errInvalidSignature)
}

var errInvalidHeader = errors.New("invalid header")

func IsInvalidHeaderErr(err error) bool {
	return errors.Is(err, errInvalidHeader)
}

var errInvalidPayloadKey = errors.New("invalid payload key")

func IsInvalidPayloadKeyErr(err error) bool {
	return errors.Is(err, errInvalidPayloadKey)
}
