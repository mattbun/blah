# blah

Parse a subtitles file, replace every line with some text, write the result into a new subtitles file.

## Usage

```shell
# Show usage info
blah --help

# Replace every line with "blah blah"
blah -i original.srt -o blahblah.srt -t "blah blah"

# Replace the lines in each subtitle with a beautiful diamond made of '+'
#   +
#  +++
# +++++
#  +++
#   +
blah -i original.srt -o diamonds.srt -1 "+" -2 "+++" -3 "+++++" -4 "+++" -5 "+"
```
