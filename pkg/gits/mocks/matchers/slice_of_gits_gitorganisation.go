// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	gits "github.com/jenkins-x/jx/v2/pkg/gits"
	"github.com/petergtz/pegomock"
)

func AnySliceOfGitsGitOrganisation() []gits.GitOrganisation {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*([]gits.GitOrganisation))(nil)).Elem()))
	var nullValue []gits.GitOrganisation
	return nullValue
}

func EqSliceOfGitsGitOrganisation(value []gits.GitOrganisation) []gits.GitOrganisation {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue []gits.GitOrganisation
	return nullValue
}
