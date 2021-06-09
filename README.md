Summary
-------
Generate a quiz using a CSV

Description
-----------
Generate a boolean-only quiz with CSV data.  After each answer has been submitted, display notes about the question when available.  When the User has completed the quiz then display their final grade.

A Climate Action Quiz will presented by default, but a CSV may be provided by the User instead.

Acceptance Criteria
-------------------
- Given a valid User supplied quiz CSV
- When the quiz has been completed by the User
- Then display the User grade

CSV Example
----------------------------------------
```
question,answer,notes
Are we experiencing the sixth mass extinction?,true,We have identified biodiversity loss that supports an extinction level event.
```

Usage
-----
```
go build cmd/csv_to_quiz/main.go
./main [-c quiz.csv]
```
