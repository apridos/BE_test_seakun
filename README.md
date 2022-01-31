### Seakun Back End Test

#### Here's the API for problem no.5

```url : seakun.aprido.my.id```

Please firstly log in here:

| POST | ```/api/v1/admin/login```|
|------|------------------------------|
|Body   |```username : string, password : string``` |
|data | ```{"username" : "seakun", "password" : "superaman"}```|


*All body is JSON formatted.*

#### Creating new teacher

| POST | ```/api/v1/teacher/create ```|
|------|------------------------------|
|Body   |```name : string, birth_date : string datetime``` |
|Example | ```{"name" : "Haryanto", "birth_date" : "2006-01-02T15:04:05.000Z"}```|

#### Edit a teacher

| POST | ```/api/v1/teacher/edit ```|
|------|------------------------------|
|Body   |```id : int, name : string, birth_date : string datetime``` |
|Example | ```{"id" : 11, "name" : "Haryantos", "birth_date" : "2006-01-02T15:04:05.000Z"}```|

#### Delete a teacher

| POST | ```/api/v1/teacher/delete ```|
|------|------------------------------|
|Body   |```id : int``` |
|Example | ```{"id" : 11}```|

#### Read all teacher data

| GET | ```/api/v1/teacher/all ```|
|------|------------------------------|

#### Read a teacher data

| GET | ```/api/v1/teacher/data/:id ```|
|------|------------------------------|
|Params   |```id : int``` |
|Example | ```/api/v1/teacher/data/9```|

#### Read a teacher data

| GET | ```/api/v1/teacher/search/:name/:birth_date ```|
|------|------------------------------|
|Params   |```name : string, birth_date : string datetime``` |
|Example | ```/api/v1/teacher/search/aku/2006-01-02```|


