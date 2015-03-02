# htmlq

Query HTML from StdIn using css selectors.

## Example
```
curl https://api.slack.com/rtm | htmlq "a[class='bold block']"
```

## Build 
```
go build -o /usr/local/bin/htmlq main.go
```

## Why 
Sometimes it is necessary to get just some values from a webpage.

## Thanks 
Many thanks to PuerkitoBio (https://github.com/PuerkitoBio/goquery). Without his library this tool would have been many more than the current few lines.

## Contributions

Contributions are welcome.

## Creators

**Remco Verhoef**
- <https://twitter.com/remco_verhoef>
- <https://twitter.com/dutchcoders>
