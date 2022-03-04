import { querySuggestionsApi } from '@experimental-api-clients-automation/client-query-suggestions';
import { ApiError } from '@experimental-api-clients-automation/client-common';
import dotenv from 'dotenv';

dotenv.config({ path: '../../.env' });

const appId = process.env.ALGOLIA_APPLICATION_ID || '**** APP_ID *****';
const apiKey =
  process.env.ALGOLIA_QUERY_SUGGESTIONS_KEY ||
  '**** QUERY_SUGGESTIONS_KEY *****';

// Init client with appId and apiKey
const client = querySuggestionsApi(appId, apiKey, 'us');

async function testQuerySuggestions() {
  try {
    const res = await client.getAllConfigs();

    console.log(`[OK]`, res);
  } catch (e) {
    if (e instanceof ApiError) {
      return console.log(`[${e.status}] ${e.message}`, e.stackTrace, e);
    }

    console.log('[ERROR]', e);
  }
}

testQuerySuggestions();
