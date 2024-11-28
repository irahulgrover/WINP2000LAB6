# WINP2000LAB6
# PLEASE PREVIEW THIS README FILE TO SEE THE SCREENSHOTS ATTACHED
# NAME - RAHUL
# STUDENT ID - 500225972

Building a Go API for Current Toronto Time with MySQL Database Logging

Overview

In this assignment, you will develop an API using the Go programming language. This API will provide the current time in Toronto in JSON format. Additionally, each request to the API will log the current time to a MySQL database.

Objectives
1.	Create a Go API Endpoint:
Develop an endpoint that returns the current Toronto time in JSON format.
2.	MySQL Database Integration:
Connect to a MySQL database and store the time data for each API request.
3.	Time Zone Handling:
Ensure that the time returned is accurately adjusted to Toronto's time zone.

Install MySQL using command “brew install MySQL” -

<img width="504" alt="image" src="https://github.com/user-attachments/assets/e7c4815c-7356-4b31-8771-f453e2db740d">

Login to MySQL and Create a Database and Table Inside the MySQL shell-

<img width="440" alt="image" src="https://github.com/user-attachments/assets/a5c61980-196a-4cf3-878c-776734588bc1">

Create a directory “go__api_project” for project-
Install dependencies using “go get -u github.com/go-sql-driver/mysql” -

<img width="468" alt="image" src="https://github.com/user-attachments/assets/804dfd45-b43b-42ac-b231-1c419b12e847">

Install dependencies “ go get -u github.com/gin-gonic/gin”. -

<img width="468" alt="image" src="https://github.com/user-attachments/assets/1e15405d-bf8d-4ce4-8e27-c344cc667338">

Create a file named main.go  in project directory and write the Go Application – 

<img width="474" alt="image" src="https://github.com/user-attachments/assets/76c21ab9-10bf-4d16-9af9-5d47877f7316">

Started the Server using “go run main.go”-

<img width="468" alt="image" src="https://github.com/user-attachments/assets/c517c923-c749-472e-8be5-3d869e144354">

Test the API using  http://localhost:8080/current-time , http://localhost:8080/logged-times OR curl http://localhost:8080/current-time , curl http://localhost:8080/logged-times . I have tested using both and provided the screenshots below.

<img width="520" alt="image" src="https://github.com/user-attachments/assets/3d87123c-5ae6-441e-bd70-8afdb954f39e">

<img width="582" alt="image" src="https://github.com/user-attachments/assets/2f61d4b3-5cb5-42c3-bd41-eb3a27e20997">

<img width="532" alt="image" src="https://github.com/user-attachments/assets/025ba039-b363-4e79-bb21-ac7412d05713">

Hence all the objectives of the asssignment were completed-
Created a Go API Endpoint
MySQL Database Integration
Time Zone Handling

Bonus Challenges
•	Implement logging in your Go application to log events and errors.
•	Create an additional endpoint to retrieve all logged times from the database.
•	Dockerize your Go application and the MySQL database for easy deployment.


Build and Run Containers using command “docker-compose up—build”-

<img width="468" alt="image" src="https://github.com/user-attachments/assets/b13d0c6b-9719-44a0-b758-8d14c1081e3f">



<img width="468" alt="image" src="https://github.com/user-attachments/assets/c3e723d3-c87e-454c-9346-60eb3468e111">

Implement logging in your Go application to log events and errors-

<img width="468" alt="image" src="https://github.com/user-attachments/assets/6d357ddf-98b1-46bf-b050-6633fd36d0aa">

Create an additional endpoint to retrieve all logged times from the database-

<img width="669" alt="image" src="https://github.com/user-attachments/assets/986acb4a-daad-4300-9028-ddc69d74d2ee">






















