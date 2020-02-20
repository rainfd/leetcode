#!/bin/bash

awk '{ for(i=1;i<=NF;i++) counter[$i]++ } END{ for(word in counter) print word, counter[word]}' words.txt | sort -r -n -k 2
