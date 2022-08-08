# csvt - CSV Transform
A CLI tool for easy CSV file transformations such as

- Filtering by columns
- Flatten the CSV table

Further operations might be added int the future, but these are the ones I mainly need now.

# Usage
```bash
cat my-data.csv |\
  ./csvt filter col --keep label |\
  ./csvt flatten \
  > flattened.csv
```