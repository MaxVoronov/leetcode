#!/usr/bin/env bash

# Read from the file words.txt and output the word frequency list to stdout.
cat words.txt | tr ' ' '\n' | sed '/^$/d' - | sort | uniq -c | sort -nr | sed -En "s/^\s+([0-9]+)\s([a-z]+)$/\2 \1/p"
