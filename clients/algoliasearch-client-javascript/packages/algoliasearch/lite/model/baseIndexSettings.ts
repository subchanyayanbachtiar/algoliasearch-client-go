// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

export type BaseIndexSettings = {
  /**
   * Creates replicas, exact copies of an index.
   */
  replicas?: string[];

  /**
   * Set the maximum number of hits accessible via pagination.
   */
  paginationLimitedTo?: number;

  /**
   * List of attributes that can\'t be retrieved at query time.
   */
  unretrievableAttributes?: string[];

  /**
   * A list of words for which you want to turn off typo tolerance.
   */
  disableTypoToleranceOnWords?: string[];

  /**
   * Specify on which attributes in your index Algolia should apply Japanese transliteration to make words indexed in Katakana or Kanji searchable in Hiragana.
   */
  attributesToTransliterate?: string[];

  /**
   * List of attributes on which to do a decomposition of camel case words.
   */
  camelCaseAttributes?: string[];

  /**
   * Specify on which attributes in your index Algolia should apply word segmentation, also known as decompounding.
   */
  decompoundedAttributes?: Record<string, any>;

  /**
   * Sets the languages at the index level for language-specific processing such as tokenization and normalization.
   */
  indexLanguages?: string[];

  /**
   * List of attributes on which you want to disable prefix matching.
   */
  disablePrefixOnAttributes?: string[];

  /**
   * Enables compression of large integer arrays.
   */
  allowCompressionOfIntegerArray?: boolean;

  /**
   * List of numeric attributes that can be used as numerical filters.
   */
  numericAttributesForFiltering?: string[];

  /**
   * Control which separators are indexed.
   */
  separatorsToIndex?: string;

  /**
   * The complete list of attributes used for searching.
   */
  searchableAttributes?: string[];

  /**
   * Lets you store custom data in your indices.
   */
  userData?: Record<string, any>;

  /**
   * Overrides Algolia\'s default normalization.
   */
  customNormalization?: Record<string, Record<string, string>>;
};
