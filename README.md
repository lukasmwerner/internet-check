# internet-check



Part of my small basic tools collection. Does what it says, tells you if you are
online. The main motivation for creating this program was the incredibly
inconsistent internet on my college campus.

## Config
`internet-check` loads a JSON config file from `~/.config/internet-check/hosts.json`
```json
[
    {
        "Host": "https://apple.com",
        "StatusCode": 200
    },
    {
        "Host": "https://google.com",
        "StatusCode": 200
    },
    {
        "Host": "https://microsoft.com",
        "StatusCode": 200
    },
    {
        "Host": "https://ard.de",
        "StatusCode": 200
    },
    {
        "Host": "https://nytimes.com",
        "StatusCode": 200
    },
    {
        "Host": "https://example.com",
        "StatusCode": 200
    },
    {
        "Host": "https://news.ycombinator.com",
        "StatusCode": 200
    }
]
```

![](/demo.gif)
