![license](https://img.shields.io/github/license/SKKUGoon/goServer)

# Kowl - Korean Stock Market Event Driven Strategy Watchmen

## New projects Notice

This project will be implemented into new korean stock exchange API project. Which will be coded after the kimchi premium project - python based. (golang version kimchi project is deprecated)


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

#### Parsing

RSS feed gives us direct access to report container. However, the values inside the table is not directly retrievable.
The table is loaded onto an separate `iframe` which serves the static html file mentioned in the src variable.
The url for raw material can be recreated by obtaining dcmNo value and rcpNo value from the container.


```go
const (
    // dcmNo parsing info
    dcmId     = "'dcmNo'"
    dcmLength = 7

    // rcpNo parsing info
    rcpId     = "'rcpNo'"
    rcpLength = 14

    dcmParseIndSrt = len(dcmId) + len("] = '")
    rcpParseIndSrt = len(rcpId) + len("] = '")
)

func getDcmNo(ctx string) (string, string) {
    i := strings.Index(ctx, dcmId)
    dcmNo := ctx[i+dcmParseIndSrt : i+dcmParseIndSrt+dcmLength]
    return "dcmNo", dcmNo
}

func getRcpNo(ctx string) (string, string) {
    i := strings.Index(ctx, rcpId)
    rcpNo := ctx[i+rcpParseIndSrt : i+rcpParseIndSrt+rcpLength]
    return "rcpNo", rcpNo
}
```

Using `dcmNo` and `rcpNo` and other parameters such as `eleId`, `offset`, `length`, and `dtd` to default,
we can finally access raw data table.


#### Parsing the raw data

