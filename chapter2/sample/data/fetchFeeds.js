import Parser from 'rss-parser';
import fs from 'fs';

const parser = new Parser();

// Soma feeds.json
const feeds = JSON.parse(fs.readFileSync('feeds.json', 'utf-8'));

async function fetchAllFeeds() {
  for (const feed of feeds) {
    try {
      const rss = await parser.parseURL(feed.link);
      console.log(`\nLatest posts from ${feed.site}:`);
      rss.items.forEach(item => {
        console.log(`- ${item.title} (${item.link})`);
      });
    } catch (err) {
      console.log(`Error fetching ${feed.site}:`, err.message);
    }
  }
}

fetchAllFeeds();