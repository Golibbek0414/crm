# crm

         Features/Functionalities:
1.  teachers & students (connected in groups)
2.  Console for teachers (list of students, attendance, marks, homework)
3.  Schedule
4.  console for students (marks, schedule, attendance)
5.  Admin (set up the app - register teachers & students, etc.)
        Architecture:
api gateway,
teachers service (includes teachers & subjects tables)
students service (includes students & groups tables)
schedules service (schedule table)
performance service (journal table)
auth service (authenticates requests)