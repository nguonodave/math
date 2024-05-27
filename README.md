# About

This is a simple go project for calculating statistical measures such as average, median, variance, and standard deviation from a given dataset.

## Usage

To use it, follow these steps:

1.  Clone the repository:

```
git clone <REPOSITORY-URL-OR-PATH>
```

Replace `<REPOSITORY-URL-OR-PATH>` with the convention of the url or path you want to use. If you use a personal access token, you can clone the following way:

```
git clone https://<PAT-TOKEN>@github.com/nguonodave/math.git
```

Replace `<PAT-TOKEN>` with your token.

2.  Navigate to the project directory:

```
cd math
```

3.  Run the program:

```
go run main.go data.txt
```
Replace data.txt with the path to your dataset file. Each line in the file should contain a single data point.

# Input Format

The input dataset should be provided in a text file (**data.txt** as per the above example), where each line contains a single numerical value.

# Output

The program will calculate and display the following statistical measures:

- Average
- Median
- Variance
- Standard Deviation

## Example

Suppose you have a dataset file named data.txt with the following contents:

```
534
2
92355
12
0
```

Running the program with this dataset will produce the following output:

```
Average: 18581
Median: 12
Variance: 1360707569
Standard Deviation: 36888
```

# Contributing

Contributions are welcome! If you have any suggestions, feature requests, or bug reports, please open an issue or create a pull request.

# License

This project is licensed under the [MIT License](LICENCE).
