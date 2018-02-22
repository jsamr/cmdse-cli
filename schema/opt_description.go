package schema

type OptDescription struct {
	Description string
	MatchModels MatchModels
}

// This function returns the semantic tokens type associated with the provided argument
// if it matched at least one, nil otherwise
func (optDescription *OptDescription) MatchArgument(arg string) ([]*SemanticTokenType, bool) {
	var matches = make([]*SemanticTokenType, 0, 10)
	for _, matchModel := range optDescription.MatchModels {
		if matchModel.leftSideRegex.MatchString(arg) {
			matches = append(matches, matchModel.variant.flagTokenType)
		}
	}
	return matches, len(matches) > 0
}

// Initialize an OptDescription ; assign description address to matchModels' description field.
func NewOptDescription(description string, matchModels ...*MatchModel) *OptDescription {
	for _, desc := range matchModels {
		desc.description = description
	}
	return &OptDescription{
		description,
		matchModels,
	}
}
