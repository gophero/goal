package errorx

import "fmt"

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w", err)
}

// Wrapf returns an error that wraps err with given format and args.
func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	if len(args) == 0 {
		return fmt.Errorf("%s: %w", format, err)
	}
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
}
