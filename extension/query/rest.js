import { createPayload } from './converter.js';

/**
 * Post queryString and Results to the API
 * @param {string} queryString
 * @param {Result[]} results 
 */
async function post(queryString, results) {
  chrome.storage.sync.get(async ({ apiURL }) => {
    try {
      const url = `${apiURL}/query`;

      const payload = createPayload(queryString, results);

      await fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        headers: {
          'Content-Type': 'application/json'
        },
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(payload)
      });
    } catch (err) {
      console.log('Post failed', err)
    }
  });
}

export default {
  post,
}
