# Go-JJC
A boilerplate for **Go** projects, which covers the basics of setting up a server with **Golang**. This repo covers server routing in a more professional way as well as how to serve/render static web pages and database setup/connections for **MongoDB, MySQL, MariaDB, GraphQL, FireBase** etc. and **WebSocket** using **Golang**. This project is free to use and maintained by PolarSoft Technologies. Support from like-minded Go developers and correctional help is highly appreciated.

**NOTE:**
This project uses **<a href="https://docs.gofiber.io/">Fiber</a>** which is an **Express** inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.

Basically, if you want to learn how to use **Go** as your backend and you are coming from a **NodeJs** background using **Express.js**, then this repo is for you! It handles and represent to you the usage of Express-like **NodeJs** setup in a **Go** way, while being rigorious but yet simple to read and understand.

Also, for those who don't have any knowledge of Express in NodeJs, you will definitely love **Go-JJC** for its simplicity and super fast ability to power your APIs/server/backend effectively!


# Starting A Server
Starting a **Go** server is very simple to do; simply navigate to the project folder you want to use (say **cd onlyrouting**) and then run **go run main.go** to start up your **Go** server. This is simple but has a downside, that every time you make a sigle change you need re-run **go run main.go** command again in order to test your app effectively with the latest changes! 
It would be really cool if our server could automatically restart itself whenever it detects one or more saved changes of our project; just like using **Nodemon** with **Express** in a **NodeJs** app.

To handle this issue, every Go-JJC project/module/package uses a package called **<a href="https://github.com/cespare/reflex">Reflex</a>** which would automatically restart our server whenever we save our code(s), thus leaving us to focus more on the things that matter the most than manually and painstakingly restarting a server over and over again.
Thus to start a server, we would do this **reflex -r '\.go' -s -- sh -c "go run main.go"** instead of **go run main.go** in the current working directory.
Doing this would automatically restart our server whenever we save changes in our code! 
Do well to read about this package here **<a href="https://github.com/cespare/reflex">Reflex</a>**

# Available Modules
Go-JJC comes with different directories/folders called modules/packages. Each module/package does or accomplishes a predefined task. The different modules are singly explained below with each file in each module fully documented for better understanding of what's been done:

## 1. Only Routing (the onlyrouting folder) Module
This module explains the basics of api routing using **Golang** in a more professional way. The entry point of this module is the main.go file; it's in this file we setup our **Fiber** instance/app, configure middlewares, and app/project-wide routes.

There are two controller (product_controller.go & user_controller.go) files in this module which handle user and product related operations.

There are three routes in the product_controller.go file which are
**localhost:3330/api/v1/products/add , localhost:3330/api/v1/products/update , and localhost:3330/api/v1/products/fetchAll**. You can test out these routes in Postman or web browser to see their individual response.

The user_controller.go file also has three routes. In this file, more light is being thrown on how to retrieve params, query strings and body data sent by the client. The following are the routes to watch out for in this controller:
**localhost:3330/api/v1/user/signIn?email=test@test.com&password=123456** for sending data from the client as query strings/parameters,

**localhost:3330/api/v1/user/signIn/:email/:password** for sending data from the client as parameters, and

**localhost:3330/api/v1/user/signIn** for sending data from the client using the body object. 

Refer to the user_controller.go file for more explanation on these!

## 2. With Web Pages (the with_web_pages folder) Module
This module explains how to serve/render web pages with **Go**. This module is very straightforward and easy to understand.


# Need More Features?
If you think this repo (or you) needs more feature(s) please kindly create an issue on this repo; we'll go through it and revert to you once it's been implemented. Thanks.

**Enjoy!**
