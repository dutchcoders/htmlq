# htmlq

A querying tool for html using css selectors.

## Example
```
curl https://api.slack.com/rtm | htmlq --filter "a[class='bold block']" @href
curl https://api.slack.com/rtm | htmlq --filter img @src
curl https://api.slack.com/rtm | htmlq --filter a
```

## Build 
```
go build -o /usr/local/bin/htmlq main.go
```

## Thanks 
Many thanks to PuerkitoBio (https://github.com/PuerkitoBio/goquery). Without his library this tool would have been many more than the current few lines.

## Contributions

Contributions are welcome.

## Creators

**Remco Verhoef**
- <https://twitter.com/remco_verhoef>
- <https://twitter.com/dutchcoders>
