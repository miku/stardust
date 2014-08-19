stardust
========

Stardust, strdist. String distance measures for the command line.

    $ go run cmd/stardust/stardust.go -h
    Usage of stardust:
      -f="ngram": distance measure
      -l=false: list available measures

    $ go run cmd/stardust/stardust.go -l
    hamming
    levenshtein
    ngram

    $ go run cmd/stardust/stardust.go -f levenshtein "München" "Munich"
    4

    $ go run cmd/stardust/stardust.go -f ngram "The quick brown fox" "The fox brown quick"
    0.6190476190476191

[![Build Status](http://img.shields.io/travis/miku/stardust.svg?style=flat)](https://travis-ci.org/miku/stardust)
