### Grader Sheets
Automatically generate grader sheets. Supports conflicts of interest and different assignment numbers for different graders.

Designed for use with the Canvas LMS. Canvas can download all submissions with a predictable pattern:
`lastfirst_submitted-filename.ext`. 

### Quickstart
- Create `graders.json`
  - "grade" - how many assignments should this program attempt to assign (default in `strutures/grader.go`)
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
- Point to submissions directory
  - Default is `./submissions`, change it in `functions/count_submissions.go`
- From the project root directory, `go run main.go`
- Run `./print.sh` to view the results

### Todo
- Support different directories for submissions, passed from command line
- Pass `graders.json` from command line
- Make the code less unreadable