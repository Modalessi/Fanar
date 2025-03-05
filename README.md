# IAU RESOURCES
where you find everything you need



the idea of the website :
a website where iau students can find everything they need

like what ?
- major plans, and their details
- courses, and their details
- resources for each course
- calender
- doctors evaluations (in the future)


what is a plan :
a set of years
each year is a set of terms
each term is a set of courses


what is a course :
- it has a title
- description
- credit hours
- course-code
- contact hours
- pre-requiues


- what is resours
- it is a file, linked to a course
- with tags of course (tags are powerfull)


## Data Strcture

- Course
  - id: uuid
  - title: string
  - course_page: md file
  - course-code: string
  - credit hours: int
  - contact-hours: int 

- Prerequest
  - course: uuid (refrence Course.id)
  - prerequest: uuid (refrence Course.id)
  
- Resroueses
  - id: uuid
  - title: string
  - description: string (optional)
  - url: string
  - tags: []string


## TODO
Auth
- register [DONE]
- login [DONE]
- admin middlewre [DONE]