// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	versioned "github.com/jenkins-x/jx/pkg/client/clientset/versioned"
)

func AnyVersionedInterface() versioned.Interface {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(versioned.Interface))(nil)).Elem()))
	var nullValue versioned.Interface
	return nullValue
}

func EqVersionedInterface(value versioned.Interface) versioned.Interface {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue versioned.Interface
	return nullValue
}