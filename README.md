# csvt - CSV Transform
A CLI tool for easy CSV file transformations such as

- Filtering by column names
- Flatten the CSV into a single string

Further operations might be added int the future, but these are the ones I mainly need now.

# Usage
```bash
cat my-data.csv |\
  ./csvt filter col --keep column-a,column-b |\
  ./csvt flatten \
  > flattened.csv
```