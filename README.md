# Bloom Filter
Bloom filter written in Go based on
[this](https://codeburst.io/lets-implement-a-bloom-filter-in-go-b2da8a4b849f) article.
Below is a definition of a bloom filter, paraphrased from [Wikipedia](https://en.wikipedia.org/wiki/Bloom_filter).

> A bloom filter is a space-efficient probabilistic data structure that is used to test whether an element is a
> member of a set. False positive matches are possible, but false negatives are not; i.e. a query returns either
> “possibly in set” or “definitely not in set”. Elements can be added to the set, but not removed.
