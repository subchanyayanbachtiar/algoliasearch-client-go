// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

/**
 * Response from the API when the Destination is successfully updated.
 */
export type DestinationUpdateResponse = {
  /**
   * The destination UUID.
   */
  destinationID: string;

  /**
   * An human readable name describing the object.
   */
  name: string;

  /**
   * Date of last update (RFC3339 format).
   */
  updatedAt: string;
};
