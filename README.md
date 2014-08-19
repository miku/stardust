stardust
========



Stardust, strdist. String distance measures for the command line.

[![Build Status](http://img.shields.io/travis/miku/stardust.svg?style=flat)](https://travis-ci.org/miku/stardust)

![Actual star dust](http://www.jpl.nasa.gov/images/herschel/20120110/pia15254-640.jpg)

Overview
--------

    $ go run cmd/stardust/stardust.go -h
    Usage of stardust:
      -m="ngram": distance measure
      -l=false: list available measures

    $ go run cmd/stardust/stardust.go -l
    hamming
    jaro
    levenshtein
    ngram

    $ go run cmd/stardust/stardust.go -m levenshtein "München" "Munich"
    4

    $ go run cmd/stardust/stardust.go -m ngram "The quick brown fox" "The fox brown quick"
    0.6190476190476191

    $ go run cmd/stardust/stardust.go -m jaro "The quick brown fox" "The fox brown quick"
    0.6432748538011696
