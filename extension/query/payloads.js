export function ResultPayload({ Title, URL, Description }) {
  Object.assign(this, { Title, URL, Description })
}

export function QueryPayload({ QueryString, Results }) {
  Object.assign(this, { QueryString, Results });
}