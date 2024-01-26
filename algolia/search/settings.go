// Code generated by go generate. DO NOT EDIT.

package search

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/debug"
	"github.com/subchanyayanbachtiar/algoliasearch-client-go/v3/algolia/opt"
)

// Settings represents an index settings configuration.
type Settings struct {
	SearchableAttributes                    *opt.SearchableAttributesOption                    `json:"searchableAttributes,omitempty"`
	AttributesForFaceting                   *opt.AttributesForFacetingOption                   `json:"attributesForFaceting,omitempty"`
	UnretrievableAttributes                 *opt.UnretrievableAttributesOption                 `json:"unretrievableAttributes,omitempty"`
	AttributesToRetrieve                    *opt.AttributesToRetrieveOption                    `json:"attributesToRetrieve,omitempty"`
	Ranking                                 *opt.RankingOption                                 `json:"ranking,omitempty"`
	CustomRanking                           *opt.CustomRankingOption                           `json:"customRanking,omitempty"`
	RelevancyStrictness                     *opt.RelevancyStrictnessOption                     `json:"relevancyStrictness,omitempty"`
	Replicas                                *opt.ReplicasOption                                `json:"replicas,omitempty"`
	Primary                                 *opt.PrimaryOption                                 `json:"primary,omitempty"`
	MaxValuesPerFacet                       *opt.MaxValuesPerFacetOption                       `json:"maxValuesPerFacet,omitempty"`
	SortFacetValuesBy                       *opt.SortFacetValuesByOption                       `json:"sortFacetValuesBy,omitempty"`
	AttributesToHighlight                   *opt.AttributesToHighlightOption                   `json:"attributesToHighlight,omitempty"`
	AttributesToSnippet                     *opt.AttributesToSnippetOption                     `json:"attributesToSnippet,omitempty"`
	HighlightPreTag                         *opt.HighlightPreTagOption                         `json:"highlightPreTag,omitempty"`
	HighlightPostTag                        *opt.HighlightPostTagOption                        `json:"highlightPostTag,omitempty"`
	SnippetEllipsisText                     *opt.SnippetEllipsisTextOption                     `json:"snippetEllipsisText,omitempty"`
	RestrictHighlightAndSnippetArrays       *opt.RestrictHighlightAndSnippetArraysOption       `json:"restrictHighlightAndSnippetArrays,omitempty"`
	HitsPerPage                             *opt.HitsPerPageOption                             `json:"hitsPerPage,omitempty"`
	PaginationLimitedTo                     *opt.PaginationLimitedToOption                     `json:"paginationLimitedTo,omitempty"`
	MinWordSizefor1Typo                     *opt.MinWordSizefor1TypoOption                     `json:"minWordSizefor1Typo,omitempty"`
	MinWordSizefor2Typos                    *opt.MinWordSizefor2TyposOption                    `json:"minWordSizefor2Typos,omitempty"`
	TypoTolerance                           *opt.TypoToleranceOption                           `json:"typoTolerance,omitempty"`
	AllowTyposOnNumericTokens               *opt.AllowTyposOnNumericTokensOption               `json:"allowTyposOnNumericTokens,omitempty"`
	DisableTypoToleranceOnAttributes        *opt.DisableTypoToleranceOnAttributesOption        `json:"disableTypoToleranceOnAttributes,omitempty"`
	DisableTypoToleranceOnWords             *opt.DisableTypoToleranceOnWordsOption             `json:"disableTypoToleranceOnWords,omitempty"`
	SeparatorsToIndex                       *opt.SeparatorsToIndexOption                       `json:"separatorsToIndex,omitempty"`
	IgnorePlurals                           *opt.IgnorePluralsOption                           `json:"ignorePlurals,omitempty"`
	AttributesToTransliterate               *opt.AttributesToTransliterateOption               `json:"attributesToTransliterate,omitempty"`
	RemoveStopWords                         *opt.RemoveStopWordsOption                         `json:"removeStopWords,omitempty"`
	CamelCaseAttributes                     *opt.CamelCaseAttributesOption                     `json:"camelCaseAttributes,omitempty"`
	DecompoundedAttributes                  *opt.DecompoundedAttributesOption                  `json:"decompoundedAttributes,omitempty"`
	KeepDiacriticsOnCharacters              *opt.KeepDiacriticsOnCharactersOption              `json:"keepDiacriticsOnCharacters,omitempty"`
	CustomNormalization                     *opt.CustomNormalizationOption                     `json:"customNormalization,omitempty"`
	QueryLanguages                          *opt.QueryLanguagesOption                          `json:"queryLanguages,omitempty"`
	IndexLanguages                          *opt.IndexLanguagesOption                          `json:"indexLanguages,omitempty"`
	DecompoundQuery                         *opt.DecompoundQueryOption                         `json:"decompoundQuery,omitempty"`
	QueryType                               *opt.QueryTypeOption                               `json:"queryType,omitempty"`
	RemoveWordsIfNoResults                  *opt.RemoveWordsIfNoResultsOption                  `json:"removeWordsIfNoResults,omitempty"`
	AdvancedSyntax                          *opt.AdvancedSyntaxOption                          `json:"advancedSyntax,omitempty"`
	OptionalWords                           *opt.OptionalWordsOption                           `json:"optionalWords,omitempty"`
	DisablePrefixOnAttributes               *opt.DisablePrefixOnAttributesOption               `json:"disablePrefixOnAttributes,omitempty"`
	DisableExactOnAttributes                *opt.DisableExactOnAttributesOption                `json:"disableExactOnAttributes,omitempty"`
	ExactOnSingleWordQuery                  *opt.ExactOnSingleWordQueryOption                  `json:"exactOnSingleWordQuery,omitempty"`
	AlternativesAsExact                     *opt.AlternativesAsExactOption                     `json:"alternativesAsExact,omitempty"`
	AdvancedSyntaxFeatures                  *opt.AdvancedSyntaxFeaturesOption                  `json:"advancedSyntaxFeatures,omitempty"`
	EnableRules                             *opt.EnableRulesOption                             `json:"enableRules,omitempty"`
	EnablePersonalization                   *opt.EnablePersonalizationOption                   `json:"enablePersonalization,omitempty"`
	EnableReRanking                         *opt.EnableReRankingOption                         `json:"enableReRanking,omitempty"`
	ReRankingApplyFilter                    *opt.ReRankingApplyFilterOption                    `json:"reRankingApplyFilter,omitempty"`
	NumericAttributesForFiltering           *opt.NumericAttributesForFilteringOption           `json:"numericAttributesForFiltering,omitempty"`
	AllowCompressionOfIntegerArray          *opt.AllowCompressionOfIntegerArrayOption          `json:"allowCompressionOfIntegerArray,omitempty"`
	AttributeForDistinct                    *opt.AttributeForDistinctOption                    `json:"attributeForDistinct,omitempty"`
	Distinct                                *opt.DistinctOption                                `json:"distinct,omitempty"`
	ReplaceSynonymsInHighlight              *opt.ReplaceSynonymsInHighlightOption              `json:"replaceSynonymsInHighlight,omitempty"`
	MinProximity                            *opt.MinProximityOption                            `json:"minProximity,omitempty"`
	ResponseFields                          *opt.ResponseFieldsOption                          `json:"responseFields,omitempty"`
	MaxFacetHits                            *opt.MaxFacetHitsOption                            `json:"maxFacetHits,omitempty"`
	Advanced                                *opt.AdvancedOption                                `json:"advanced,omitempty"`
	AttributeCriteriaComputedByMinProximity *opt.AttributeCriteriaComputedByMinProximityOption `json:"attributeCriteriaComputedByMinProximity,omitempty"`
	UserData                                *opt.UserDataOption                                `json:"userData,omitempty"`
	RenderingContent                        *RenderingContent                                  `json:"renderingContent,omitempty"`
	Extensions                              *Extensions                                        `json:"extensions,omitempty"`
	CustomSettings                          map[string]interface{}                             `json:"-"`
}

type settings Settings

type backwardCompatibleSettings struct {
	AttributesToIndex        *opt.SearchableAttributesOption          `json:"attributesToIndex,omitempty"`
	Slaves                   *opt.ReplicasOption                      `json:"slaves,omitempty"`
	NumericAttributesToIndex *opt.NumericAttributesForFilteringOption `json:"numericAttributesToIndex,omitempty"`
}

func (s *Settings) UnmarshalJSON(data []byte) error {
	var bcSettings backwardCompatibleSettings
	err := json.Unmarshal(data, &bcSettings)
	if err != nil {
		return fmt.Errorf("cannot unmarshal backward-compatible settings: %v", err)
	}

	err = json.Unmarshal(data, (*settings)(s))
	if err != nil {
		return fmt.Errorf("cannot unmarshal settings: %v", err)
	}

	if s.SearchableAttributes == nil {
		s.SearchableAttributes = bcSettings.AttributesToIndex
	}

	if s.Replicas == nil {
		s.Replicas = bcSettings.Slaves
	}

	if s.NumericAttributesForFiltering == nil {
		s.NumericAttributesForFiltering = bcSettings.NumericAttributesToIndex
	}

	err = json.Unmarshal(data, &s.CustomSettings)
	if err != nil {
		return fmt.Errorf("cannot unmarshal unknown settings: %v", err)
	}
	for _, knownSetting := range []string{
		"searchableAttributes",
		"attributesToIndex",
		"attributesForFaceting",
		"unretrievableAttributes",
		"attributesToRetrieve",
		"ranking",
		"customRanking",
		"relevancyStrictness",
		"replicas",
		"slaves",
		"primary",
		"maxValuesPerFacet",
		"sortFacetValuesBy",
		"attributesToHighlight",
		"attributesToSnippet",
		"highlightPreTag",
		"highlightPostTag",
		"snippetEllipsisText",
		"restrictHighlightAndSnippetArrays",
		"hitsPerPage",
		"paginationLimitedTo",
		"minWordSizefor1Typo",
		"minWordSizefor2Typos",
		"typoTolerance",
		"allowTyposOnNumericTokens",
		"disableTypoToleranceOnAttributes",
		"disableTypoToleranceOnWords",
		"separatorsToIndex",
		"ignorePlurals",
		"attributesToTransliterate",
		"removeStopWords",
		"camelCaseAttributes",
		"decompoundedAttributes",
		"keepDiacriticsOnCharacters",
		"customNormalization",
		"queryLanguages",
		"indexLanguages",
		"decompoundQuery",
		"queryType",
		"removeWordsIfNoResults",
		"advancedSyntax",
		"optionalWords",
		"disablePrefixOnAttributes",
		"disableExactOnAttributes",
		"exactOnSingleWordQuery",
		"alternativesAsExact",
		"advancedSyntaxFeatures",
		"enableRules",
		"enablePersonalization",
		"enableReRanking",
		"reRankingApplyFilter",
		"numericAttributesForFiltering",
		"numericAttributesToIndex",
		"allowCompressionOfIntegerArray",
		"attributeForDistinct",
		"distinct",
		"replaceSynonymsInHighlight",
		"minProximity",
		"responseFields",
		"maxFacetHits",
		"advanced",
		"attributeCriteriaComputedByMinProximity",
		"userData",
	} {
		delete(s.CustomSettings, knownSetting)
	}

	return nil
}

func (s Settings) MarshalJSON() ([]byte, error) {
	dataWithoutCustomSettings, err := json.Marshal(settings(s))
	if err != nil {
		return nil, fmt.Errorf("cannot marshal known settings: %v", err)
	}
	var m map[string]json.RawMessage
	err = json.Unmarshal(dataWithoutCustomSettings, &m)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal known settings into temporary map: %v", err)
	}
	for k, v := range s.CustomSettings {
		data, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("cannot marshal unknown settings %q: %v", k, err)
		}
		m[k] = data
	}
	return json.Marshal(m)
}

// Equal returns true if given settings are the same as the instance one. Empty
// settings are considered equal to their default value counterpart.
func (s Settings) Equal(s2 Settings) bool {
	if !opt.SearchableAttributesEqual(s.SearchableAttributes, s2.SearchableAttributes) {
		debug.Printf("Settings.SearchableAttributes are not equal: %#v != %#v\n", s.SearchableAttributes, s2.SearchableAttributes)
		return false
	}
	if !opt.AttributesForFacetingEqual(s.AttributesForFaceting, s2.AttributesForFaceting) {
		debug.Printf("Settings.AttributesForFaceting are not equal: %#v != %#v\n", s.AttributesForFaceting, s2.AttributesForFaceting)
		return false
	}
	if !opt.UnretrievableAttributesEqual(s.UnretrievableAttributes, s2.UnretrievableAttributes) {
		debug.Printf("Settings.UnretrievableAttributes are not equal: %#v != %#v\n", s.UnretrievableAttributes, s2.UnretrievableAttributes)
		return false
	}
	if !opt.AttributesToRetrieveEqual(s.AttributesToRetrieve, s2.AttributesToRetrieve) {
		debug.Printf("Settings.AttributesToRetrieve are not equal: %#v != %#v\n", s.AttributesToRetrieve, s2.AttributesToRetrieve)
		return false
	}
	if !opt.RankingEqual(s.Ranking, s2.Ranking) {
		debug.Printf("Settings.Ranking are not equal: %#v != %#v\n", s.Ranking, s2.Ranking)
		return false
	}
	if !opt.CustomRankingEqual(s.CustomRanking, s2.CustomRanking) {
		debug.Printf("Settings.CustomRanking are not equal: %#v != %#v\n", s.CustomRanking, s2.CustomRanking)
		return false
	}
	if !opt.RelevancyStrictnessEqual(s.RelevancyStrictness, s2.RelevancyStrictness) {
		debug.Printf("Settings.RelevancyStrictness are not equal: %#v != %#v\n", s.RelevancyStrictness, s2.RelevancyStrictness)
		return false
	}
	if !opt.ReplicasEqual(s.Replicas, s2.Replicas) {
		debug.Printf("Settings.Replicas are not equal: %#v != %#v\n", s.Replicas, s2.Replicas)
		return false
	}
	if !opt.PrimaryEqual(s.Primary, s2.Primary) {
		debug.Printf("Settings.Primary are not equal: %#v != %#v\n", s.Primary, s2.Primary)
		return false
	}
	if !opt.MaxValuesPerFacetEqual(s.MaxValuesPerFacet, s2.MaxValuesPerFacet) {
		debug.Printf("Settings.MaxValuesPerFacet are not equal: %#v != %#v\n", s.MaxValuesPerFacet, s2.MaxValuesPerFacet)
		return false
	}
	if !opt.SortFacetValuesByEqual(s.SortFacetValuesBy, s2.SortFacetValuesBy) {
		debug.Printf("Settings.SortFacetValuesBy are not equal: %#v != %#v\n", s.SortFacetValuesBy, s2.SortFacetValuesBy)
		return false
	}
	if !opt.AttributesToHighlightEqual(s.AttributesToHighlight, s2.AttributesToHighlight) {
		debug.Printf("Settings.AttributesToHighlight are not equal: %#v != %#v\n", s.AttributesToHighlight, s2.AttributesToHighlight)
		return false
	}
	if !opt.AttributesToSnippetEqual(s.AttributesToSnippet, s2.AttributesToSnippet) {
		debug.Printf("Settings.AttributesToSnippet are not equal: %#v != %#v\n", s.AttributesToSnippet, s2.AttributesToSnippet)
		return false
	}
	if !opt.HighlightPreTagEqual(s.HighlightPreTag, s2.HighlightPreTag) {
		debug.Printf("Settings.HighlightPreTag are not equal: %#v != %#v\n", s.HighlightPreTag, s2.HighlightPreTag)
		return false
	}
	if !opt.HighlightPostTagEqual(s.HighlightPostTag, s2.HighlightPostTag) {
		debug.Printf("Settings.HighlightPostTag are not equal: %#v != %#v\n", s.HighlightPostTag, s2.HighlightPostTag)
		return false
	}
	if !opt.SnippetEllipsisTextEqual(s.SnippetEllipsisText, s2.SnippetEllipsisText) {
		debug.Printf("Settings.SnippetEllipsisText are not equal: %#v != %#v\n", s.SnippetEllipsisText, s2.SnippetEllipsisText)
		return false
	}
	if !opt.RestrictHighlightAndSnippetArraysEqual(s.RestrictHighlightAndSnippetArrays, s2.RestrictHighlightAndSnippetArrays) {
		debug.Printf("Settings.RestrictHighlightAndSnippetArrays are not equal: %#v != %#v\n", s.RestrictHighlightAndSnippetArrays, s2.RestrictHighlightAndSnippetArrays)
		return false
	}
	if !opt.HitsPerPageEqual(s.HitsPerPage, s2.HitsPerPage) {
		debug.Printf("Settings.HitsPerPage are not equal: %#v != %#v\n", s.HitsPerPage, s2.HitsPerPage)
		return false
	}
	if !opt.PaginationLimitedToEqual(s.PaginationLimitedTo, s2.PaginationLimitedTo) {
		debug.Printf("Settings.PaginationLimitedTo are not equal: %#v != %#v\n", s.PaginationLimitedTo, s2.PaginationLimitedTo)
		return false
	}
	if !opt.MinWordSizefor1TypoEqual(s.MinWordSizefor1Typo, s2.MinWordSizefor1Typo) {
		debug.Printf("Settings.MinWordSizefor1Typo are not equal: %#v != %#v\n", s.MinWordSizefor1Typo, s2.MinWordSizefor1Typo)
		return false
	}
	if !opt.MinWordSizefor2TyposEqual(s.MinWordSizefor2Typos, s2.MinWordSizefor2Typos) {
		debug.Printf("Settings.MinWordSizefor2Typos are not equal: %#v != %#v\n", s.MinWordSizefor2Typos, s2.MinWordSizefor2Typos)
		return false
	}
	if !opt.TypoToleranceEqual(s.TypoTolerance, s2.TypoTolerance) {
		debug.Printf("Settings.TypoTolerance are not equal: %#v != %#v\n", s.TypoTolerance, s2.TypoTolerance)
		return false
	}
	if !opt.AllowTyposOnNumericTokensEqual(s.AllowTyposOnNumericTokens, s2.AllowTyposOnNumericTokens) {
		debug.Printf("Settings.AllowTyposOnNumericTokens are not equal: %#v != %#v\n", s.AllowTyposOnNumericTokens, s2.AllowTyposOnNumericTokens)
		return false
	}
	if !opt.DisableTypoToleranceOnAttributesEqual(s.DisableTypoToleranceOnAttributes, s2.DisableTypoToleranceOnAttributes) {
		debug.Printf("Settings.DisableTypoToleranceOnAttributes are not equal: %#v != %#v\n", s.DisableTypoToleranceOnAttributes, s2.DisableTypoToleranceOnAttributes)
		return false
	}
	if !opt.DisableTypoToleranceOnWordsEqual(s.DisableTypoToleranceOnWords, s2.DisableTypoToleranceOnWords) {
		debug.Printf("Settings.DisableTypoToleranceOnWords are not equal: %#v != %#v\n", s.DisableTypoToleranceOnWords, s2.DisableTypoToleranceOnWords)
		return false
	}
	if !opt.SeparatorsToIndexEqual(s.SeparatorsToIndex, s2.SeparatorsToIndex) {
		debug.Printf("Settings.SeparatorsToIndex are not equal: %#v != %#v\n", s.SeparatorsToIndex, s2.SeparatorsToIndex)
		return false
	}
	if !opt.IgnorePluralsEqual(s.IgnorePlurals, s2.IgnorePlurals) {
		debug.Printf("Settings.IgnorePlurals are not equal: %#v != %#v\n", s.IgnorePlurals, s2.IgnorePlurals)
		return false
	}
	if !opt.AttributesToTransliterateEqual(s.AttributesToTransliterate, s2.AttributesToTransliterate) {
		debug.Printf("Settings.AttributesToTransliterate are not equal: %#v != %#v\n", s.AttributesToTransliterate, s2.AttributesToTransliterate)
		return false
	}
	if !opt.RemoveStopWordsEqual(s.RemoveStopWords, s2.RemoveStopWords) {
		debug.Printf("Settings.RemoveStopWords are not equal: %#v != %#v\n", s.RemoveStopWords, s2.RemoveStopWords)
		return false
	}
	if !opt.CamelCaseAttributesEqual(s.CamelCaseAttributes, s2.CamelCaseAttributes) {
		debug.Printf("Settings.CamelCaseAttributes are not equal: %#v != %#v\n", s.CamelCaseAttributes, s2.CamelCaseAttributes)
		return false
	}
	if !opt.DecompoundedAttributesEqual(s.DecompoundedAttributes, s2.DecompoundedAttributes) {
		debug.Printf("Settings.DecompoundedAttributes are not equal: %#v != %#v\n", s.DecompoundedAttributes, s2.DecompoundedAttributes)
		return false
	}
	if !opt.KeepDiacriticsOnCharactersEqual(s.KeepDiacriticsOnCharacters, s2.KeepDiacriticsOnCharacters) {
		debug.Printf("Settings.KeepDiacriticsOnCharacters are not equal: %#v != %#v\n", s.KeepDiacriticsOnCharacters, s2.KeepDiacriticsOnCharacters)
		return false
	}
	if !opt.CustomNormalizationEqual(s.CustomNormalization, s2.CustomNormalization) {
		debug.Printf("Settings.CustomNormalization are not equal: %#v != %#v\n", s.CustomNormalization, s2.CustomNormalization)
		return false
	}
	if !opt.QueryLanguagesEqual(s.QueryLanguages, s2.QueryLanguages) {
		debug.Printf("Settings.QueryLanguages are not equal: %#v != %#v\n", s.QueryLanguages, s2.QueryLanguages)
		return false
	}
	if !opt.IndexLanguagesEqual(s.IndexLanguages, s2.IndexLanguages) {
		debug.Printf("Settings.IndexLanguages are not equal: %#v != %#v\n", s.IndexLanguages, s2.IndexLanguages)
		return false
	}
	if !opt.DecompoundQueryEqual(s.DecompoundQuery, s2.DecompoundQuery) {
		debug.Printf("Settings.DecompoundQuery are not equal: %#v != %#v\n", s.DecompoundQuery, s2.DecompoundQuery)
		return false
	}
	if !opt.QueryTypeEqual(s.QueryType, s2.QueryType) {
		debug.Printf("Settings.QueryType are not equal: %#v != %#v\n", s.QueryType, s2.QueryType)
		return false
	}
	if !opt.RemoveWordsIfNoResultsEqual(s.RemoveWordsIfNoResults, s2.RemoveWordsIfNoResults) {
		debug.Printf("Settings.RemoveWordsIfNoResults are not equal: %#v != %#v\n", s.RemoveWordsIfNoResults, s2.RemoveWordsIfNoResults)
		return false
	}
	if !opt.AdvancedSyntaxEqual(s.AdvancedSyntax, s2.AdvancedSyntax) {
		debug.Printf("Settings.AdvancedSyntax are not equal: %#v != %#v\n", s.AdvancedSyntax, s2.AdvancedSyntax)
		return false
	}
	if !opt.OptionalWordsEqual(s.OptionalWords, s2.OptionalWords) {
		debug.Printf("Settings.OptionalWords are not equal: %#v != %#v\n", s.OptionalWords, s2.OptionalWords)
		return false
	}
	if !opt.DisablePrefixOnAttributesEqual(s.DisablePrefixOnAttributes, s2.DisablePrefixOnAttributes) {
		debug.Printf("Settings.DisablePrefixOnAttributes are not equal: %#v != %#v\n", s.DisablePrefixOnAttributes, s2.DisablePrefixOnAttributes)
		return false
	}
	if !opt.DisableExactOnAttributesEqual(s.DisableExactOnAttributes, s2.DisableExactOnAttributes) {
		debug.Printf("Settings.DisableExactOnAttributes are not equal: %#v != %#v\n", s.DisableExactOnAttributes, s2.DisableExactOnAttributes)
		return false
	}
	if !opt.ExactOnSingleWordQueryEqual(s.ExactOnSingleWordQuery, s2.ExactOnSingleWordQuery) {
		debug.Printf("Settings.ExactOnSingleWordQuery are not equal: %#v != %#v\n", s.ExactOnSingleWordQuery, s2.ExactOnSingleWordQuery)
		return false
	}
	if !opt.AlternativesAsExactEqual(s.AlternativesAsExact, s2.AlternativesAsExact) {
		debug.Printf("Settings.AlternativesAsExact are not equal: %#v != %#v\n", s.AlternativesAsExact, s2.AlternativesAsExact)
		return false
	}
	if !opt.AdvancedSyntaxFeaturesEqual(s.AdvancedSyntaxFeatures, s2.AdvancedSyntaxFeatures) {
		debug.Printf("Settings.AdvancedSyntaxFeatures are not equal: %#v != %#v\n", s.AdvancedSyntaxFeatures, s2.AdvancedSyntaxFeatures)
		return false
	}
	if !opt.EnableRulesEqual(s.EnableRules, s2.EnableRules) {
		debug.Printf("Settings.EnableRules are not equal: %#v != %#v\n", s.EnableRules, s2.EnableRules)
		return false
	}
	if !opt.EnablePersonalizationEqual(s.EnablePersonalization, s2.EnablePersonalization) {
		debug.Printf("Settings.EnablePersonalization are not equal: %#v != %#v\n", s.EnablePersonalization, s2.EnablePersonalization)
		return false
	}
	if !opt.EnableReRankingEqual(s.EnableReRanking, s2.EnableReRanking) {
		debug.Printf("Settings.EnableReRanking are not equal: %#v != %#v\n", s.EnableReRanking, s2.EnableReRanking)
		return false
	}
	if !opt.ReRankingApplyFilterEqual(s.ReRankingApplyFilter, s2.ReRankingApplyFilter) {
		debug.Printf("Settings.ReRankingApplyFilter are not equal: %#v != %#v\n", s.ReRankingApplyFilter, s2.ReRankingApplyFilter)
		return false
	}
	if !opt.NumericAttributesForFilteringEqual(s.NumericAttributesForFiltering, s2.NumericAttributesForFiltering) {
		debug.Printf("Settings.NumericAttributesForFiltering are not equal: %#v != %#v\n", s.NumericAttributesForFiltering, s2.NumericAttributesForFiltering)
		return false
	}
	if !opt.AllowCompressionOfIntegerArrayEqual(s.AllowCompressionOfIntegerArray, s2.AllowCompressionOfIntegerArray) {
		debug.Printf("Settings.AllowCompressionOfIntegerArray are not equal: %#v != %#v\n", s.AllowCompressionOfIntegerArray, s2.AllowCompressionOfIntegerArray)
		return false
	}
	if !opt.AttributeForDistinctEqual(s.AttributeForDistinct, s2.AttributeForDistinct) {
		debug.Printf("Settings.AttributeForDistinct are not equal: %#v != %#v\n", s.AttributeForDistinct, s2.AttributeForDistinct)
		return false
	}
	if !opt.DistinctEqual(s.Distinct, s2.Distinct) {
		debug.Printf("Settings.Distinct are not equal: %#v != %#v\n", s.Distinct, s2.Distinct)
		return false
	}
	if !opt.ReplaceSynonymsInHighlightEqual(s.ReplaceSynonymsInHighlight, s2.ReplaceSynonymsInHighlight) {
		debug.Printf("Settings.ReplaceSynonymsInHighlight are not equal: %#v != %#v\n", s.ReplaceSynonymsInHighlight, s2.ReplaceSynonymsInHighlight)
		return false
	}
	if !opt.MinProximityEqual(s.MinProximity, s2.MinProximity) {
		debug.Printf("Settings.MinProximity are not equal: %#v != %#v\n", s.MinProximity, s2.MinProximity)
		return false
	}
	if !opt.ResponseFieldsEqual(s.ResponseFields, s2.ResponseFields) {
		debug.Printf("Settings.ResponseFields are not equal: %#v != %#v\n", s.ResponseFields, s2.ResponseFields)
		return false
	}
	if !opt.MaxFacetHitsEqual(s.MaxFacetHits, s2.MaxFacetHits) {
		debug.Printf("Settings.MaxFacetHits are not equal: %#v != %#v\n", s.MaxFacetHits, s2.MaxFacetHits)
		return false
	}
	if !opt.AdvancedEqual(s.Advanced, s2.Advanced) {
		debug.Printf("Settings.Advanced are not equal: %#v != %#v\n", s.Advanced, s2.Advanced)
		return false
	}
	if !opt.AttributeCriteriaComputedByMinProximityEqual(s.AttributeCriteriaComputedByMinProximity, s2.AttributeCriteriaComputedByMinProximity) {
		debug.Printf("Settings.AttributeCriteriaComputedByMinProximity are not equal: %#v != %#v\n", s.AttributeCriteriaComputedByMinProximity, s2.AttributeCriteriaComputedByMinProximity)
		return false
	}
	if !opt.UserDataEqual(s.UserData, s2.UserData) {
		debug.Printf("Settings.UserData are not equal: %#v != %#v\n", s.UserData, s2.UserData)
		return false
	}
	return true
}

func (s Settings) String() string {
	settingsStr := "Settings{\n"
	settingsStr += fmt.Sprintf("\tSearchableAttributes: %v,\n", stringifyReturnValues(s.SearchableAttributes.Get()))
	settingsStr += fmt.Sprintf("\tAttributesForFaceting: %v,\n", stringifyReturnValues(s.AttributesForFaceting.Get()))
	settingsStr += fmt.Sprintf("\tUnretrievableAttributes: %v,\n", stringifyReturnValues(s.UnretrievableAttributes.Get()))
	settingsStr += fmt.Sprintf("\tAttributesToRetrieve: %v,\n", stringifyReturnValues(s.AttributesToRetrieve.Get()))
	settingsStr += fmt.Sprintf("\tRanking: %v,\n", stringifyReturnValues(s.Ranking.Get()))
	settingsStr += fmt.Sprintf("\tCustomRanking: %v,\n", stringifyReturnValues(s.CustomRanking.Get()))
	settingsStr += fmt.Sprintf("\tRelevancyStrictness: %v,\n", stringifyReturnValues(s.RelevancyStrictness.Get()))
	settingsStr += fmt.Sprintf("\tReplicas: %v,\n", stringifyReturnValues(s.Replicas.Get()))
	settingsStr += fmt.Sprintf("\tPrimary: %v,\n", stringifyReturnValues(s.Primary.Get()))
	settingsStr += fmt.Sprintf("\tMaxValuesPerFacet: %v,\n", stringifyReturnValues(s.MaxValuesPerFacet.Get()))
	settingsStr += fmt.Sprintf("\tSortFacetValuesBy: %v,\n", stringifyReturnValues(s.SortFacetValuesBy.Get()))
	settingsStr += fmt.Sprintf("\tAttributesToHighlight: %v,\n", stringifyReturnValues(s.AttributesToHighlight.Get()))
	settingsStr += fmt.Sprintf("\tAttributesToSnippet: %v,\n", stringifyReturnValues(s.AttributesToSnippet.Get()))
	settingsStr += fmt.Sprintf("\tHighlightPreTag: %v,\n", stringifyReturnValues(s.HighlightPreTag.Get()))
	settingsStr += fmt.Sprintf("\tHighlightPostTag: %v,\n", stringifyReturnValues(s.HighlightPostTag.Get()))
	settingsStr += fmt.Sprintf("\tSnippetEllipsisText: %v,\n", stringifyReturnValues(s.SnippetEllipsisText.Get()))
	settingsStr += fmt.Sprintf("\tRestrictHighlightAndSnippetArrays: %v,\n", stringifyReturnValues(s.RestrictHighlightAndSnippetArrays.Get()))
	settingsStr += fmt.Sprintf("\tHitsPerPage: %v,\n", stringifyReturnValues(s.HitsPerPage.Get()))
	settingsStr += fmt.Sprintf("\tPaginationLimitedTo: %v,\n", stringifyReturnValues(s.PaginationLimitedTo.Get()))
	settingsStr += fmt.Sprintf("\tMinWordSizefor1Typo: %v,\n", stringifyReturnValues(s.MinWordSizefor1Typo.Get()))
	settingsStr += fmt.Sprintf("\tMinWordSizefor2Typos: %v,\n", stringifyReturnValues(s.MinWordSizefor2Typos.Get()))
	settingsStr += fmt.Sprintf("\tTypoTolerance: %v,\n", stringifyReturnValues(s.TypoTolerance.Get()))
	settingsStr += fmt.Sprintf("\tAllowTyposOnNumericTokens: %v,\n", stringifyReturnValues(s.AllowTyposOnNumericTokens.Get()))
	settingsStr += fmt.Sprintf("\tDisableTypoToleranceOnAttributes: %v,\n", stringifyReturnValues(s.DisableTypoToleranceOnAttributes.Get()))
	settingsStr += fmt.Sprintf("\tDisableTypoToleranceOnWords: %v,\n", stringifyReturnValues(s.DisableTypoToleranceOnWords.Get()))
	settingsStr += fmt.Sprintf("\tSeparatorsToIndex: %v,\n", stringifyReturnValues(s.SeparatorsToIndex.Get()))
	settingsStr += fmt.Sprintf("\tIgnorePlurals: %v,\n", stringifyReturnValues(s.IgnorePlurals.Get()))
	settingsStr += fmt.Sprintf("\tAttributesToTransliterate: %v,\n", stringifyReturnValues(s.AttributesToTransliterate.Get()))
	settingsStr += fmt.Sprintf("\tRemoveStopWords: %v,\n", stringifyReturnValues(s.RemoveStopWords.Get()))
	settingsStr += fmt.Sprintf("\tCamelCaseAttributes: %v,\n", stringifyReturnValues(s.CamelCaseAttributes.Get()))
	settingsStr += fmt.Sprintf("\tDecompoundedAttributes: %v,\n", stringifyReturnValues(s.DecompoundedAttributes.Get()))
	settingsStr += fmt.Sprintf("\tKeepDiacriticsOnCharacters: %v,\n", stringifyReturnValues(s.KeepDiacriticsOnCharacters.Get()))
	settingsStr += fmt.Sprintf("\tCustomNormalization: %v,\n", stringifyReturnValues(s.CustomNormalization.Get()))
	settingsStr += fmt.Sprintf("\tQueryLanguages: %v,\n", stringifyReturnValues(s.QueryLanguages.Get()))
	settingsStr += fmt.Sprintf("\tIndexLanguages: %v,\n", stringifyReturnValues(s.IndexLanguages.Get()))
	settingsStr += fmt.Sprintf("\tDecompoundQuery: %v,\n", stringifyReturnValues(s.DecompoundQuery.Get()))
	settingsStr += fmt.Sprintf("\tQueryType: %v,\n", stringifyReturnValues(s.QueryType.Get()))
	settingsStr += fmt.Sprintf("\tRemoveWordsIfNoResults: %v,\n", stringifyReturnValues(s.RemoveWordsIfNoResults.Get()))
	settingsStr += fmt.Sprintf("\tAdvancedSyntax: %v,\n", stringifyReturnValues(s.AdvancedSyntax.Get()))
	settingsStr += fmt.Sprintf("\tOptionalWords: %v,\n", stringifyReturnValues(s.OptionalWords.Get()))
	settingsStr += fmt.Sprintf("\tDisablePrefixOnAttributes: %v,\n", stringifyReturnValues(s.DisablePrefixOnAttributes.Get()))
	settingsStr += fmt.Sprintf("\tDisableExactOnAttributes: %v,\n", stringifyReturnValues(s.DisableExactOnAttributes.Get()))
	settingsStr += fmt.Sprintf("\tExactOnSingleWordQuery: %v,\n", stringifyReturnValues(s.ExactOnSingleWordQuery.Get()))
	settingsStr += fmt.Sprintf("\tAlternativesAsExact: %v,\n", stringifyReturnValues(s.AlternativesAsExact.Get()))
	settingsStr += fmt.Sprintf("\tAdvancedSyntaxFeatures: %v,\n", stringifyReturnValues(s.AdvancedSyntaxFeatures.Get()))
	settingsStr += fmt.Sprintf("\tEnableRules: %v,\n", stringifyReturnValues(s.EnableRules.Get()))
	settingsStr += fmt.Sprintf("\tEnablePersonalization: %v,\n", stringifyReturnValues(s.EnablePersonalization.Get()))
	settingsStr += fmt.Sprintf("\tEnableReRanking: %v,\n", stringifyReturnValues(s.EnableReRanking.Get()))
	settingsStr += fmt.Sprintf("\tReRankingApplyFilter: %v,\n", stringifyReturnValues(s.ReRankingApplyFilter.Get()))
	settingsStr += fmt.Sprintf("\tNumericAttributesForFiltering: %v,\n", stringifyReturnValues(s.NumericAttributesForFiltering.Get()))
	settingsStr += fmt.Sprintf("\tAllowCompressionOfIntegerArray: %v,\n", stringifyReturnValues(s.AllowCompressionOfIntegerArray.Get()))
	settingsStr += fmt.Sprintf("\tAttributeForDistinct: %v,\n", stringifyReturnValues(s.AttributeForDistinct.Get()))
	settingsStr += fmt.Sprintf("\tDistinct: %v,\n", stringifyReturnValues(s.Distinct.Get()))
	settingsStr += fmt.Sprintf("\tReplaceSynonymsInHighlight: %v,\n", stringifyReturnValues(s.ReplaceSynonymsInHighlight.Get()))
	settingsStr += fmt.Sprintf("\tMinProximity: %v,\n", stringifyReturnValues(s.MinProximity.Get()))
	settingsStr += fmt.Sprintf("\tResponseFields: %v,\n", stringifyReturnValues(s.ResponseFields.Get()))
	settingsStr += fmt.Sprintf("\tMaxFacetHits: %v,\n", stringifyReturnValues(s.MaxFacetHits.Get()))
	settingsStr += fmt.Sprintf("\tAdvanced: %v,\n", stringifyReturnValues(s.Advanced.Get()))
	settingsStr += fmt.Sprintf("\tAttributeCriteriaComputedByMinProximity: %v,\n", stringifyReturnValues(s.AttributeCriteriaComputedByMinProximity.Get()))
	settingsStr += fmt.Sprintf("\tUserData: %v,\n", stringifyReturnValues(s.UserData.Get()))

	settingsStr += fmt.Sprintf("\tRenderingContent: %+v\n", s.RenderingContent)

	settingsStr += "\tCustomSettings{\n"
	for k, v := range s.CustomSettings {
		settingsStr += fmt.Sprintf("\t\t%s: %v,\n", k, v)
	}
	settingsStr += "\t},\n}"

	return settingsStr
}

func stringifyReturnValues(args ...interface{}) string {
	var s []string
	for _, arg := range args {
		s = append(s, fmt.Sprintf("%v", arg))
	}
	return "(" + strings.Join(s, ",") + ")"
}
