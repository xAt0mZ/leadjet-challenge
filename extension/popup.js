import * as Consts from './query/consts.js';
import { Result } from './query/types.js';

import { extractResultsFromHtml } from './query/extractor.js';
import QueryService from './query/rest.js';

let saveResultsButton = document.getElementById("saveResultsButton");

saveResultsButton.addEventListener("click", async () => {
  const [tab] = await chrome.tabs.query({ active: true, currentWindow: true });

  const query = tab.url.match(/(?:.*?)\.google\.(?:\w+?)\/search\?q=(.*?)&(?:.*)/);

  if (query) {
    chrome.scripting.executeScript({
      target: { tabId: tab.id },
      function: extractResultsFromHtml,
      args: [Consts]
    },
      (results) => onResultsExtracted(query, results));
  }
});

function onResultsExtracted(query, res) {
  const queryString = query[1];
  const results = res[0].result.map((r) => new Result(r));

  QueryService.post(queryString, results);
}
