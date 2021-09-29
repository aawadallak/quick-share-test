
[![Lint Go Code](https://github.com/aawadallak/quick-share-test/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/aawadallak/quick-share-test/actions/workflows/go.yml)

[![Build & Test](https://github.com/aawadallak/quick-share-test/actions/workflows/test.yml/badge.svg)](https://github.com/aawadallak/quick-share-test/actions/workflows/test.yml)


## About the project
The service aims to serve as an app for quick and hassle-free file sharing. The idea was to generate something that didn't require registration to share a file. Because we often want to send files to other people and this process can be a bit complicated and painful. The proposal has only two limiter

 - **Max size** of sharing **100 MB**
 - **Max time** the archive will be stored on server: **24 HOURS**

## Technologies
The project was built with **Golang**, the chosen database was **MongoDB** due to the project needs and the data structure, the same was also added to **Docker** so  the ***continuous implementation*** process can be generated using the concept of ***container***. It was also added a route to check if the application is accepting requests, basically a ping route. Architecture is an extremely important process in the decision to create a project and thinking about the scalability of the project, we chose to use the concept of **Domain Driven Design (DDD)**. The next step of the project will be to add **Unit and Automated Tests** to the application, to guarantee the security of the development and to certify that the implementations occurred correctly.

## Future implementations

 - [ ] Unit and automated tests
 - [ ] Database migrations trough **Go**
 - [ ] Structure different log levels

## Routes

```
Check the server status and return a message as response
Method: GET
Path: http://localhost:9004/api/v1/ping
```
```
Download files from an id.
Method: GET
Path: http://localhost:9004/api/v1/download?id={id}
```
```
Responsible for uploading the files
Method: POST
Path: http://localhost:9004/api/v1/upload
The requisition body must be adequate.
```
