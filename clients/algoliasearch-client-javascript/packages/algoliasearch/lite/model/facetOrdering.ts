// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { Facets } from './facets';
import type { Value } from './value';

/**
 * Defining how facets should be ordered.
 */
export type FacetOrdering = {
  facets?: Facets;

  /**
   * The ordering of facet values, within an individual list.
   */
  values?: Record<string, Value>;
};
