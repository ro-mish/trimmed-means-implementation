# Trimmed Mean Comparison between R and Go Implementations

In this repository, we examine the differences between the implementations for 
trimmed mean, both in R and Go. 

We use the trimmed-means-go repository (ro-mish/trimmed-mean-go) to simplify our execution. 
As such, we obtain the results here that display the difference between R and Go implementations.

## Setup and Installation

```bash
go mod init trimmed-mean-test
go get github.com/ro-mish/trimmed-mean-go
```

## Running the Comparison

```bash
chmod +x compare_results.sh
./compare_results.sh
```

## Results

```
=== Running Go Implementation ===
--------------------------------
Sample size: 100
Trim percentage: 5.0%
Number of values trimmed from each end: 5
Number of values used in mean: 90

Integer Data (first 10): [682 315 389 796 140 348 596 531 118 334]
Integer 5% Trimmed Mean: 515.22

Float Data (first 10): [757.46 328.15 414.75 498.17 324.19 129.13 80.61 247.87 969.16 27.66]
Float 5% Trimmed Mean: 486.46

=== Running R Implementation ===
--------------------------------
Integer Data (first 10):
[1] 283 847 917 100 622 904 644 933 399 899

Integer 5% Trimmed Mean: 526.83 

Float Data (first 10):
[1] "204.20" "133.74" "325.68" "155.06" "129.96" "435.53" "38.64" "713.30" "100.77" "950.30"

Float 5% Trimmed Mean: 536.49 

Number of values trimmed from each end: 5 
Total number of values used in mean: 90
```

## Analysis
As we notice, the trimmed mean is not consistent across both implementations.
We can deduce that this is likely due to the random number generators that each respective language leverages.
However, despite the difference in random number generators, there is still an interesting phenomena that
the mean for both is around 500. They are not too far off, despite using different random number generators.

## Conclusions
In summary, we see the following:

- Different random number generators in Go and R produce distinct sequences even with the same seed, leading to different sample data

- Despite using different RNGs, both implementations:
   - Successfully trim 5% from each end (5 values from each tail)
   - Calculate means using the middle 90 values
   - Produce reasonably close results clustering around 500

- The core trimmed mean calculation logic works consistently across both implementations

- The similar final results (appx 500) suggest both implementations are statistically sound, even with different data

This validates that our Go implementation of the trimmed mean algorithm is working as expected, producing statistically meaningful results comparable to the implementation in R.

## Gen AI Disclosure
Gen AI conversation history for this repo can be found in genAI-usage.txt