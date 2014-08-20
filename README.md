stardust
========



Stardust, strdist. String distance measures for the command line.

[![Build Status](http://img.shields.io/travis/miku/stardust.svg?style=flat)](https://travis-ci.org/miku/stardust)

![Actual star dust](http://www.jpl.nasa.gov/images/herschel/20120110/pia15254-640.jpg)

Overview
--------

    $ stardust
    NAME:
       stardust - String similarity measures for tab separated values.

    USAGE:
       stardust [global options] command [command options] [arguments...]

    VERSION:
       0.1.0

    AUTHOR:
      Martin Czygan - <martin.czygan@gmail.com>

    COMMANDS:
       ngram    Ngram similarity
       hamming  Hamming distance
       jaro     Jaro similarity
       plain    Plain passthrough (for IO benchmarks)
       help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       -f '1,2'     c1,c2 the two columns to use for the comparison
       --help, -h       show help
       --version, -v    print the version

    $ stardust hamming "Hallo" "Hello"
    Hallo   Hello   1

    $ stardust ngram "Hallo" "Hello"
    Hallo   Hello   0.2

Measures can have their own options, too:

    $ stardust ngram --help
    NAME:
       ngram - Ngram similarity

    USAGE:
       command ngram [command options] [arguments...]

    DESCRIPTION:
       Compute Ngram similarity, which lies between 0 and 1.

    OPTIONS:
       --size, -s '3'   value of n


    $ stardust ngram --size 2 "Hello" "Hallo"
    Hello   Hallo   0.3333333333333333

    $ stardust ngram --size 1 "Hallo" "Hello"
    Hallo   Hello   0.6
