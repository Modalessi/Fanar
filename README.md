# Fanar
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
  
- Resrouese
  - id: uuid
  - course_id: uuid (refrence course.id)
  - title: string
  - description: string (optional)
  - url: string (the actual file)
  - tags: []string
    - tags are hardcoded (Notes, Homeworks, Quizes, Labs, Slides, Midterms, Finals, Exams, OldExams)


## TODO
Auth
- register [DONE]
- login [DONE]
- admin middlewre [DONE]
- restore passwords
- confirm email


- Course
- Create Course (admin) [DONE]
- Delete Course (admin) [DONE]
- Edit Course (admin) [DONE]
- Get Course [DONE]



- Resource
- download resource [DONEs]
- create resource (upload) [DONE]
- delete resource (deleteing) 
  - admin can delete any file
  - user can delete only the files he uploaded
  