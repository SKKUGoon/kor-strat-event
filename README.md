![license](https://img.shields.io/github/license/SKKUGoon/goServer)

# Kowl - Korean Stock Market Event Driven Strategy Watchmen

Kowl provides near-real-time event driven strategy for korean stocks in
KOSPI and KOSDAQ by scraping DART(Data Analysis, Retrieval and Transfer System)
provided by Korean FINANCIAL SUPERVISORY SERVICE. By parsing native RSS feed,
Kowl creates signal that is tradable in any other stock trading APIs.

## Specifics

#### RSS Feed

Innately, DART RSS Feed has 500 requests per minute limit.
Therefore, request by seconds is more than enough.

#### Signals

KOWL focuses on 2 major events;

1. Bonus Issue
2. Rights Issue

Detailed Strategy will not be disclosed, but the scheme is fairly obvious.