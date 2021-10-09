# Task Submission for Appointy, made in Go.
The task is to develop a basic version of Instagram. We are only required to develop the API for the system.

## Getting Started

Installation prerequisites include only an Installation of Go on your system.

After installing Go, Run the following commands to run the API locally: -

# download the starter kit
git clone https://github.com/shreshth02/Instagram_Backend_API.git

# and then
go run main.go

The following endpoints are available:

* `POST /users`: To Create an User.
* `GET /users/{user_id}`: To Get an user using id. 
* `POST /posts`: To Create a Post.
* `GET /posts/{post_id} `: To Get a post using id.
* `GET /posts/user/{user_id} `: To List all posts of a user.