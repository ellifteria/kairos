# kairos

A Discord timestamp generator

## Usage

```
kairos [--format=mode] [time]
```

### Possible modes

- `short-time`
- `long-time`
- `short-date`
- `long-date`
- `short-date-time`*
- `long-date-time`
- `relative`

For more, see [https://discord.com/developers/docs/reference#message-formatting-timestamp-styles for more](https://discord.com/developers/docs/reference#message-formatting-timestamp-styles) for more.

\* default

### Time format

kairos will pattern matche a time in the format of `YYYY-MM-DD HH:mm UTC±O` where:

- `YYYY` is a four-digit year
- `MM` is a two-digit month
- `DD` is a two-digit day
- `HH` is a two-digit hour
- `mm` is a two-digit minute
- `UTC±O` is the string `UTC` followed by `+` or `-` and then a float offset `O`

If anything is left out, the default of the current date, time, and timezone will be used.

### Examples

```
kairos
```

This will print out the Discord timestamp in short date time format using local time.

```
kairos --format relative 1970-01-01 00:00 UTC+0
```

This will return `<t:0:R>` which will print out the amount of time that has elapsed since the Unix epoch.

