package main

const docsHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>hnrss — Hacker News RSS Feeds</title>
<style>
  :root {
    --orange: #ff6600;
    --bg: #f6f6ef;
    --text: #222;
    --muted: #666;
    --code-bg: #eee;
    --border: #ddd;
    --link: #0066cc;
  }
  * { box-sizing: border-box; margin: 0; padding: 0; }
  body {
    font-family: Verdana, Geneva, sans-serif;
    font-size: 14px;
    line-height: 1.6;
    color: var(--text);
    background: var(--bg);
    padding: 1.5rem;
  }
  header {
    background: var(--orange);
    color: #fff;
    padding: 0.75rem 1rem;
    border-radius: 4px;
    margin-bottom: 1.5rem;
    display: flex;
    align-items: baseline;
    gap: 0.75rem;
  }
  header h1 { font-size: 1.2rem; font-weight: bold; }
  header p  { font-size: 0.9rem; opacity: 0.9; }
  a { color: var(--link); text-decoration: none; }
  a:hover { text-decoration: underline; }
  h2 {
    font-size: 1rem;
    font-weight: bold;
    border-bottom: 2px solid var(--orange);
    padding-bottom: 0.25rem;
    margin: 1.5rem 0 0.75rem;
  }
  h3 {
    font-size: 0.9rem;
    font-weight: bold;
    color: var(--muted);
    margin: 1rem 0 0.25rem;
    text-transform: uppercase;
    letter-spacing: 0.04em;
  }
  code {
    background: var(--code-bg);
    padding: 0.1em 0.35em;
    border-radius: 3px;
    font-family: monospace;
    font-size: 0.9em;
  }
  .endpoint-grid {
    display: grid;
    grid-template-columns: minmax(220px, auto) 1fr;
    gap: 0.1rem 1rem;
    margin-bottom: 0.5rem;
  }
  .ep-url  { font-family: monospace; font-size: 0.9em; }
  .ep-desc { color: var(--muted); font-size: 0.88em; align-self: center; }
  .badge {
    display: inline-block;
    font-size: 0.7em;
    font-weight: bold;
    padding: 0.1em 0.4em;
    border-radius: 3px;
    margin-left: 0.35em;
    vertical-align: middle;
  }
  .badge-new  { background: #d4edda; color: #1a5c2e; border: 1px solid #a3d3ae; }
  .badge-req  { background: #fff3cd; color: #7a5500; border: 1px solid #ffd78a; }
  table {
    border-collapse: collapse;
    width: 100%;
    margin-bottom: 1rem;
  }
  th, td {
    text-align: left;
    padding: 0.35rem 0.6rem;
    border: 1px solid var(--border);
    font-size: 0.88em;
  }
  th { background: #e8e8e0; font-weight: bold; }
  tr:nth-child(even) td { background: #f0f0ea; }
  .note {
    background: #fff8e1;
    border-left: 3px solid var(--orange);
    padding: 0.5rem 0.75rem;
    margin: 0.75rem 0;
    font-size: 0.88em;
  }
  .example {
    background: #fff;
    border: 1px solid var(--border);
    border-radius: 4px;
    padding: 0.6rem 0.8rem;
    margin: 0.4rem 0;
    font-family: monospace;
    font-size: 0.88em;
    word-break: break-all;
  }
  .example a { color: var(--link); }
  footer {
    margin-top: 2rem;
    font-size: 0.8rem;
    color: var(--muted);
    border-top: 1px solid var(--border);
    padding-top: 0.75rem;
  }
</style>
</head>
<body>

<header>
  <h1>hnrss</h1>
  <p>Custom, realtime RSS feeds for <a href="https://news.ycombinator.com/" style="color:#fff;text-decoration:underline">Hacker News</a></p>
</header>

<p>This service provides RSS, Atom, and JSON Feed endpoints for Hacker News content.
All feeds are generated on the fly from the <a href="https://hn.algolia.com/api">Algolia HN Search API</a> — no cache, no database.</p>

<h2>Feed Formats</h2>
<p>Append a suffix to any endpoint URL to select the feed format:</p>
<table>
  <tr><th>Suffix</th><th>Format</th><th>MIME type</th></tr>
  <tr><td><em>(none or <code>.rss</code>)</em></td><td>RSS 2.0</td><td>application/rss+xml</td></tr>
  <tr><td><code>.atom</code></td><td>Atom 1.0</td><td>application/atom+xml</td></tr>
  <tr><td><code>.jsonfeed</code></td><td>JSON Feed 1.0</td><td>application/json</td></tr>
</table>

<h2>Endpoints</h2>

<h3>Stories &amp; Polls</h3>
<div class="endpoint-grid">
  <span class="ep-url"><code>/newest</code></span>
  <span class="ep-desc">Newest stories and polls</span>

  <span class="ep-url"><code>/newest/ai</code> <span class="badge badge-new">NEW</span></span>
  <span class="ep-desc">Newest stories matching AI-related keywords (see below)</span>

  <span class="ep-url"><code>/newest/noai</code> <span class="badge badge-new">NEW</span></span>
  <span class="ep-desc">Newest stories with AI-related stories filtered out</span>

  <span class="ep-url"><code>/frontpage</code></span>
  <span class="ep-desc">Current front page stories</span>

  <span class="ep-url"><code>/ask</code></span>
  <span class="ep-desc">Ask HN threads</span>

  <span class="ep-url"><code>/show</code></span>
  <span class="ep-desc">Show HN threads</span>

  <span class="ep-url"><code>/polls</code></span>
  <span class="ep-desc">Polls</span>

  <span class="ep-url"><code>/jobs</code></span>
  <span class="ep-desc">Job postings</span>
</div>

<h3>Comments</h3>
<div class="endpoint-grid">
  <span class="ep-url"><code>/newcomments</code></span>
  <span class="ep-desc">Newest comments site-wide</span>

  <span class="ep-url"><code>/bestcomments</code></span>
  <span class="ep-desc">Best comments (scraped from HN best comments page)</span>

  <span class="ep-url"><code>/item</code> <span class="badge badge-req">id=</span></span>
  <span class="ep-desc">Comments on a specific story — requires <code>id=&lt;item_id&gt;</code></span>
</div>

<h3>User Feeds</h3>
<div class="endpoint-grid">
  <span class="ep-url"><code>/user</code> <span class="badge badge-req">id=</span></span>
  <span class="ep-desc">All submissions and comments by a user — requires <code>id=&lt;username&gt;</code></span>

  <span class="ep-url"><code>/threads</code> <span class="badge badge-req">id=</span></span>
  <span class="ep-desc">Comments only by a user — requires <code>id=&lt;username&gt;</code></span>

  <span class="ep-url"><code>/submitted</code> <span class="badge badge-req">id=</span></span>
  <span class="ep-desc">Submissions only by a user — requires <code>id=&lt;username&gt;</code></span>

  <span class="ep-url"><code>/replies</code> <span class="badge badge-req">id=</span></span>
  <span class="ep-desc">Replies to a user or item — requires <code>id=&lt;username&gt;</code></span>

  <span class="ep-url"><code>/favorites</code> <span class="badge badge-req">id=</span></span>
  <span class="ep-desc">User&#39;s public favorites — requires <code>id=&lt;username&gt;</code></span>
</div>

<h3>Special Pages</h3>
<div class="endpoint-grid">
  <span class="ep-url"><code>/classic</code></span>
  <span class="ep-desc">HN Classic page</span>

  <span class="ep-url"><code>/best</code></span>
  <span class="ep-desc">HN Best page</span>

  <span class="ep-url"><code>/active</code></span>
  <span class="ep-desc">HN Active page</span>

  <span class="ep-url"><code>/invited</code></span>
  <span class="ep-desc">HN Invited page</span>

  <span class="ep-url"><code>/pool</code></span>
  <span class="ep-desc">HN Pool page</span>

  <span class="ep-url"><code>/launches</code></span>
  <span class="ep-desc">HN Launches page</span>
</div>

<h3>Who Is Hiring</h3>
<div class="endpoint-grid">
  <span class="ep-url"><code>/whoishiring</code></span>
  <span class="ep-desc">All posts from the latest "Ask HN: Who is hiring?" thread</span>

  <span class="ep-url"><code>/whoishiring/jobs</code></span>
  <span class="ep-desc">Job offers only (seeking employees)</span>

  <span class="ep-url"><code>/whoishiring/hired</code></span>
  <span class="ep-desc">Candidates only (seeking employers)</span>

  <span class="ep-url"><code>/whoishiring/freelance</code></span>
  <span class="ep-desc">Freelance posts only</span>
</div>

<h2>Query Parameters</h2>
<table>
  <tr><th>Parameter</th><th>Description</th><th>Example</th></tr>
  <tr><td><code>q</code></td><td>Full-text search query</td><td><code>q=golang</code></td></tr>
  <tr><td><code>id</code></td><td>Username or item ID (required by some endpoints)</td><td><code>id=pg</code></td></tr>
  <tr><td><code>points</code></td><td>Minimum points filter</td><td><code>points=100</code></td></tr>
  <tr><td><code>comments</code></td><td>Minimum comment count filter</td><td><code>comments=10</code></td></tr>
  <tr><td><code>count</code></td><td>Number of items to return (max 100, default 20)</td><td><code>count=50</code></td></tr>
  <tr><td><code>description</code></td><td>Set to <code>0</code> to omit item descriptions from feed entries</td><td><code>description=0</code></td></tr>
  <tr><td><code>link</code></td><td>Set to <code>url</code> to link directly to the article instead of HN comments</td><td><code>link=url</code></td></tr>
  <tr><td><code>author</code></td><td>Filter by submitter username</td><td><code>author=dang</code></td></tr>
  <tr><td><code>search_attrs</code></td><td>Algolia search attributes (comma-separated: <code>title</code>, <code>story_text</code>, <code>comment_text</code>, <code>url</code>, <code>author</code>)</td><td><code>search_attrs=title</code></td></tr>
</table>

<h2>AI Filtering Endpoints</h2>

<p><code>/newest/ai</code> and <code>/newest/noai</code> filter the newest stories feed based on AI content detection.
Matching is case-insensitive and treats hyphens, slashes, and dots as word boundaries
(so <em>LLM-based</em> matches <em>llm</em>).</p>

<p>The following terms trigger a match:</p>
<table>
  <tr><th>Category</th><th>Terms</th></tr>
  <tr>
    <td>Concepts</td>
    <td>artificial intelligence, machine learning, deep learning, neural network, large language model,
    generative ai, natural language processing, foundation model, diffusion model, retrieval augmented,
    prompt engineering, fine tuning, finetuning, vector database, embedding model, agentic, ai agent,
    ai model, ai system, ai tool, vibe coding, attention mechanism</td>
  </tr>
  <tr>
    <td>Acronyms (whole-word)</td>
    <td>AI, LLM, LLMs, AGI, RLHF, NLP, GenAI</td>
  </tr>
  <tr>
    <td>Companies</td>
    <td>OpenAI, Anthropic, DeepMind, Mistral, Cohere, Hugging Face, Stability AI, Inflection AI, xAI</td>
  </tr>
  <tr>
    <td>Products / Models</td>
    <td>ChatGPT, GPT-4, GPT-3, Claude, Gemini, Llama, Copilot, Midjourney, DALL-E, Stable Diffusion,
    Whisper, Grok, Mixtral, Perplexity, Devin, Cursor, Phi-3, Phi-4, o1, o3</td>
  </tr>
</table>

<div class="note">
  <strong>Note:</strong> <code>/newest/ai</code> fetches up to 3× the requested count from Algolia and
  then filters client-side, so fewer items may be returned than <code>count</code> requests when AI
  content is sparse. The same applies to <code>/newest/noai</code> when AI content is dense.
</div>

<h2>Examples</h2>

<div class="example"><a href="/newest?q=golang&amp;points=50">/newest?q=golang&amp;points=50</a> — newest Go stories with ≥ 50 points (RSS)</div>
<div class="example"><a href="/newest/ai">/newest/ai</a> — newest AI stories (RSS)</div>
<div class="example"><a href="/newest/noai.atom">/newest/noai.atom</a> — newest non-AI stories (Atom)</div>
<div class="example"><a href="/frontpage.atom">/frontpage.atom</a> — front page as Atom feed</div>
<div class="example"><a href="/threads?id=pg">/threads?id=pg</a> — all comments by user <em>pg</em></div>
<div class="example"><a href="/item?id=17821181">/item?id=17821181</a> — comments on a specific story</div>
<div class="example"><a href="/ask?count=5&amp;points=100">/ask?count=5&amp;points=100</a> — top 5 Ask HN with ≥ 100 points</div>
<div class="example"><a href="/newest.jsonfeed?q=rust&amp;link=url">/newest.jsonfeed?q=rust&amp;link=url</a> — newest Rust stories as JSON Feed, linking to article URLs</div>
<div class="example"><a href="/whoishiring/jobs">/whoishiring/jobs</a> — job listings from current "Who is Hiring" thread</div>

<footer>
  Self-hosted instance at <a href="https://hnrss.kakwalab.ovh/">https://hnrss.kakwalab.ovh/</a> —
  source on <a href="https://github.com/kakwa/hnrss">GitHub</a> —
  data via <a href="https://hn.algolia.com/api">Algolia HN Search API</a>
</footer>

</body>
</html>
`
