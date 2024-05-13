## A "Bloom Filter" implementation for educational purposes.
Space efficient probabilistic data structure.

| 0 | 1 | 0 | 1 | 0 | 0 | 1 | 0 | 0 | 0 |
|---|---|---|---|---|---|---|---|---|---|

### Query -> Contains ?
- NO
- Probably YES

### FALSE POSITIVE; 

- __Despite element doesn't exist but bloom filter return element exist__

### FALSE NEGATIVE;

- __Element doesn't exist__

### Main Functions
- ADD (O(n_HashFunc))
- QUERY (O(n_HashFunc))
Memory fixed

### Hash Functions
should be like;
- Fast
- Randomly distributed
- Collisions are okay
- Collisions should be rare

Some known hashing functions;
  - Murmur3
  - StringSum
  - FNV-1a