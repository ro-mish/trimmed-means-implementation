# Set random seed for reproducibility
set.seed(1234)

# Generate same test data as Go imlementation
int_data <- sample(0:999, 100, replace=TRUE)
float_data <- runif(100, 0, 1000)

# Calculate the 5% trimmed means
int_trimmed_mean <- mean(int_data, trim=0.05)
float_trimmed_mean <- mean(float_data, trim=0.05)

# Print first 10 values of each dataset
cat("Integer Data (first 10):\n")
print(head(int_data, 10))
cat("\nInteger 5% Trimmed Mean:", sprintf("%.2f", int_trimmed_mean), "\n")


# For verification, let's also show how many values were trimmed
n_trimmed <- floor(100 * 0.05)
cat("\nNumber of values trimmed from each end:", n_trimmed, "\n")
cat("Total number of values used in mean:", 100 - 2*n_trimmed, "\n") 