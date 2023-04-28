// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { BigQueryDataType } from './bigQueryDataType';

export type SourceBigQuery = {
  /**
   * Project ID of the BigQuery Source.
   */
  projectID: string;

  /**
   * Dataset ID of the BigQuery Source.
   */
  datasetID: string;

  dataType?: BigQueryDataType;

  /**
   * Table name (for default BQ).
   */
  table?: string;

  /**
   * Table prefix (for Google Analytics).
   */
  tablePrefix?: string;
};
