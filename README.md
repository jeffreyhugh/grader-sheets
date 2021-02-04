### Grader Sheets
Automatically generate grader sheets. Supports conflicts of interest and different assignment counts for different graders.

Designed for use with the Canvas LMS. Canvas can download all submissions with a predictable pattern:
`lastfirst_submitted-filename.ext`. 

### Quickstart
- Create `config.json`
  - "file_type" - file type of submissions (recommended: `.tar.gz`)
  - "grade" - how many assignments should this program attempt to assign by default (I don't know if this works)
```
{
  "file_type" : ".tar.gz",
  "grade" : 4
}
```
- Create `graders.json`
  - "grade" - how many assignments should this program attempt to assign (default in `config.json`)
  - "conflicts" - list of \[lastfirst] that the grader cannot grade
 ```
{
  "graders": [
    {
      "first": "john",
      "last": "doe",
      "email": "john.doe@my.school.edu",
      "grade": 4,
      "conflicts": [smithjane]
    }
  ]
}
```
- Run the executable
  - `-c` specifies the location of `config.json`
  - `-g` specifies the location of `graders.json`
  - `-s` specifies the location of `submissions.zip`
- Run `./print.sh` to view the results

### Scripts
- If you need to modify submission filenames or whatever before the sheet gets created, put the
script in `scripts/`. These scripts will be run after the `.zip` is extracted but before the
sheet gets created.
  - Since the name of the directory of submissions is based on the system's time, it will be 
  passed as the first argument to the script

### Todo
- Make the code less unreadable