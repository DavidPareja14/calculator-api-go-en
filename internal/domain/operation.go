package domain

// Operation identifies a supported arithmetic operation.
type Operation string

const (
	OperationAdd      Operation = "add"
	OperationSubtract Operation = "subtract"
	OperationMultiply Operation = "multiply"
	OperationDivide   Operation = "divide"
)

// ParseOperation converts a raw string to a supported Operation.
func ParseOperation(raw string) (Operation, error) {
	switch Operation(raw) {
	case OperationAdd, OperationSubtract, OperationMultiply, OperationDivide:
		return Operation(raw), nil
	default:
		return "", ErrInvalidOperation
	}
}
