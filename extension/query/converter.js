import { QueryPayload, ResultPayload} from './payloads.js';

export function createPayload(queryString, results) {
  return new QueryPayload({
    QueryString: queryString,
    Results: results.map((r) => new ResultPayload({
      Title: r.title,
      URL: r.url,
      Description: r.description,
    })),
  });
}