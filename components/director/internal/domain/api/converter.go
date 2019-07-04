package api

import (
	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/kyma-incubator/compass/components/director/pkg/graphql"
)

//go:generate mockery -name=AuthConverter -output=automock -outpkg=automock -case=underscore
type AuthConverter interface {
	ToGraphQL(in *model.Auth) *graphql.Auth
	InputFromGraphQL(in *graphql.AuthInput) *model.AuthInput
}

//go:generate mockery -name=FetchRequestConverter -output=automock -outpkg=automock -case=underscore
type FetchRequestConverter interface {
	ToGraphQL(in *model.FetchRequest) *graphql.FetchRequest
	InputFromGraphQL(in *graphql.FetchRequestInput) *model.FetchRequestInput
}

//go:generate mockery -name=VersionConverter -output=automock -outpkg=automock -case=underscore
type VersionConverter interface {
	ToGraphQL(in *model.Version) *graphql.Version
	InputFromGraphQL(in *graphql.VersionInput) *model.VersionInput
}

type converter struct {
	auth    AuthConverter
	fr      FetchRequestConverter
	version VersionConverter
}

func NewConverter(auth AuthConverter, fr FetchRequestConverter, version VersionConverter) *converter {
	return &converter{auth: auth, fr: fr, version: version}
}

func (c *converter) ToGraphQL(in *model.APIDefinition) *graphql.APIDefinition {
	if in == nil {
		return nil
	}

	return &graphql.APIDefinition{
		ID:            in.ID,
		ApplicationID: in.ApplicationID,
		Name:          in.Name,
		Description:   in.Description,
		Spec:          c.apiSpecToGraphQL(in.Spec),
		TargetURL:     in.TargetURL,
		Group:         in.Group,
		Auths:         c.runtimeAuthArrToGraphQL(in.Auths),
		DefaultAuth:   c.auth.ToGraphQL(in.DefaultAuth),
		Version:       c.version.ToGraphQL(in.Version),
	}
}

func (c *converter) MultipleToGraphQL(in []*model.APIDefinition) []*graphql.APIDefinition {
	var apis []*graphql.APIDefinition
	for _, a := range in {
		if a == nil {
			continue
		}
		apis = append(apis, c.ToGraphQL(a))
	}

	return apis
}

func (c *converter) MultipleInputFromGraphQL(in []*graphql.APIDefinitionInput) []*model.APIDefinitionInput {
	var arr []*model.APIDefinitionInput
	for _, item := range in {
		api := c.InputFromGraphQL(item)
		arr = append(arr, api)
	}

	return arr
}

func (c *converter) InputFromGraphQL(in *graphql.APIDefinitionInput) *model.APIDefinitionInput {
	if in == nil {
		return nil
	}

	return &model.APIDefinitionInput{
		Name:        in.Name,
		Description: in.Description,
		TargetURL:   in.TargetURL,
		Group:       in.Group,
		Spec:        c.apiSpecInputFromGraphQL(in.Spec),
		Version:     c.version.InputFromGraphQL(in.Version),
		DefaultAuth: c.auth.InputFromGraphQL(in.DefaultAuth),
	}
}

func (c *converter) apiSpecToGraphQL(in *model.APISpec) *graphql.APISpec {
	if in == nil {
		return nil
	}

	var data *graphql.CLOB
	if in.Data != nil {
		tmp := graphql.CLOB(*in.Data)
		data = &tmp
	}

	var format *graphql.SpecFormat
	if in.Format != nil {
		f := graphql.SpecFormat(*in.Format)
		format = &f
	}

	return &graphql.APISpec{
		Data:         data,
		Type:         graphql.APISpecType(in.Type),
		Format:       format,
		FetchRequest: c.fr.ToGraphQL(in.FetchRequest),
	}
}

func (c *converter) apiSpecInputFromGraphQL(in *graphql.APISpecInput) *model.APISpecInput {
	if in == nil {
		return nil
	}

	var format model.SpecFormat
	if in.Format != "" {
		format = model.SpecFormat(in.Format)
	}

	return &model.APISpecInput{
		Data:         (*string)(in.Data),
		Type:         model.APISpecType(in.Type),
		Format:       &format,
		FetchRequest: c.fr.InputFromGraphQL(in.FetchRequest),
	}
}

func (c *converter) runtimeAuthToGraphQL(in *model.RuntimeAuth) *graphql.RuntimeAuth {
	if in == nil {
		return nil
	}

	return &graphql.RuntimeAuth{
		RuntimeID: in.RuntimeID,
		Auth:      c.auth.ToGraphQL(in.Auth),
	}
}

func (c *converter) runtimeAuthArrToGraphQL(in []*model.RuntimeAuth) []*graphql.RuntimeAuth {
	var auths []*graphql.RuntimeAuth
	for _, item := range in {
		auths = append(auths, &graphql.RuntimeAuth{
			RuntimeID: item.RuntimeID,
			Auth:      c.auth.ToGraphQL(item.Auth),
		})
	}

	return auths
}
