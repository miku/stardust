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

    COMMANDS:
       ngram    Ngram similarity
       hamming  Hamming distance
       levenshtein  Levenshtein distance
       jaro     Jaro distance
       jaro-winkler Jaro-Winkler distance
       plain    Plain passthrough (for IO benchmarks)
       help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       -f '1,2'     c1,c2 the two columns to use for the comparison
       --delimiter, -d '    '   column delimiter (defaults to tab)
       --version, -v    print the version
       --help, -h       show help

For starters
------------

    $ stardust hamming "Hallo" "Hello"
    Hallo   Hello   1

    $ stardust ngram "Hallo" "Hello"
    Hallo   Hello   0.2

    $ stardust ngram "Hallo Welt" "Hello World"
    Hallo Welt  Hello World 0.21428571428571427

Are the man pages of `cp` and `mv` more similar that those of `ls` and `mv`,
when measured with a trigram model?

    $ stardust ngram "$(echo $(man ls))" "$(echo $(man mv))" | cut -f3
    0.29057337220602525

    $ stardust ngram "$(echo $(man cp))" "$(echo $(man mv))" | cut -f3
    0.4792746113989637

They seem to. And according to Jaro similarity?

    $ stardust jaro "$(echo $(man ls))" "$(echo $(man mv))" | cut -f3
    0.5597612762544908

    $ stardust jaro "$(echo $(man cp))" "$(echo $(man mv))" | cut -f3
    0.6376732132890776

Still.

Specific options
----------------

Some measures come with additional options, e.g. ngram will take a size
option, which corresponds to the `n` in ngram.

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
