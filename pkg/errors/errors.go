package errors

import (
	"github.com/pkg/errors"
)

func New(message string) error {
	return errors.New(message)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args)
}

func WithStack(err error) error {
	return errors.WithStack(err)
}

// type fundamental struct {
// 	Message string `json:"message"`
// 	// Code     int    `json:"code"`
// 	// Internal error  `json:"-"`
// }

// // Error .
// func (f *fundamental) Error() string {
// 	return f.Message
// }

// func New(text string) error {
// 	return &errorString{text}
// }

// // errorString is a trivial implementation of error.
// type errorString struct {
// 	s string
// }

// func (e *errorString) Error() string {
// 	return e.s
// }

// // Unwrap returns the result of calling the Unwrap method on err, if err's
// // type contains an Unwrap method returning error.
// // Otherwise, Unwrap returns nil.
// func Unwrap(err error) error {
// 	u, ok := err.(interface {
// 		Unwrap() error
// 	})
// 	if !ok {
// 		return nil
// 	}
// 	return u.Unwrap()
// }

// // Is reports whether any error in err's chain matches target.
// //
// // The chain consists of err itself followed by the sequence of errors obtained by
// // repeatedly calling Unwrap.
// //
// // An error is considered to match a target if it is equal to that target or if
// // it implements a method Is(error) bool such that Is(target) returns true.
// //
// // An error type might provide an Is method so it can be treated as equivalent
// // to an existing error. For example, if MyError defines
// //
// //	func (m MyError) Is(target error) bool { return target == fs.ErrExist }
// //
// // then Is(MyError{}, fs.ErrExist) returns true. See syscall.Errno.Is for
// // an example in the standard library.
// func Is(err, target error) bool {
// 	if target == nil {
// 		return err == target
// 	}

// 	isComparable := reflectlite.TypeOf(target).Comparable()
// 	for {
// 		if isComparable && err == target {
// 			return true
// 		}
// 		if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
// 			return true
// 		}
// 		// TODO: consider supporting target.Is(err). This would allow
// 		// user-definable predicates, but also may allow for coping with sloppy
// 		// APIs, thereby making it easier to get away with them.
// 		if err = Unwrap(err); err == nil {
// 			return false
// 		}
// 	}
// }

// // As finds the first error in err's chain that matches target, and if so, sets
// // target to that error value and returns true. Otherwise, it returns false.
// //
// // The chain consists of err itself followed by the sequence of errors obtained by
// // repeatedly calling Unwrap.
// //
// // An error matches target if the error's concrete value is assignable to the value
// // pointed to by target, or if the error has a method As(interface{}) bool such that
// // As(target) returns true. In the latter case, the As method is responsible for
// // setting target.
// //
// // An error type might provide an As method so it can be treated as if it were a
// // different error type.
// //
// // As panics if target is not a non-nil pointer to either a type that implements
// // error, or to any interface type.
// func As(err error, target interface{}) bool {
// 	if target == nil {
// 		panic("errors: target cannot be nil")
// 	}
// 	val := reflectlite.ValueOf(target)
// 	typ := val.Type()
// 	if typ.Kind() != reflectlite.Ptr || val.IsNil() {
// 		panic("errors: target must be a non-nil pointer")
// 	}
// 	if e := typ.Elem(); e.Kind() != reflectlite.Interface && !e.Implements(errorType) {
// 		panic("errors: *target must be interface or implement error")
// 	}
// 	targetType := typ.Elem()
// 	for err != nil {
// 		if reflectlite.TypeOf(err).AssignableTo(targetType) {
// 			val.Elem().Set(reflectlite.ValueOf(err))
// 			return true
// 		}
// 		if x, ok := err.(interface{ As(interface{}) bool }); ok && x.As(target) {
// 			return true
// 		}
// 		err = Unwrap(err)
// 	}
// 	return false
// }

// var errorType = reflectlite.TypeOf((*error)(nil)).Elem()
