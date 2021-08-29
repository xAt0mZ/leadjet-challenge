
export function extractResultsFromHtml({
  resultWrapperClass,
  urlWrapperClass,
  descriptionWrapperClass
}) {
  const res = []
  const results = document.getElementsByClassName(resultWrapperClass);
  for (let idx = 0; idx < results.length; idx++) {
    const elem = results[idx];
    const link = elem.getElementsByClassName(urlWrapperClass)[0].getElementsByTagName('a')[0];
    const url = link.href;
    const title = link.getElementsByTagName('h3')[0].textContent;
    const description = elem.getElementsByClassName(descriptionWrapperClass)[0].getElementsByTagName('span')[0].textContent;
    res.push({ url, title, description });
  }
  return res;
}