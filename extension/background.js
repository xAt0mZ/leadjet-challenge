let apiURL = 'http://localhost:10000/api';

chrome.runtime.onInstalled.addListener(() => {
  chrome.storage.sync.set({ apiURL });
});