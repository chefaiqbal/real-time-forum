<<<<<<< HEAD
# 🌟 Real-Time Forum Project


## 🚀 Objectives

This project focuses on the following key points:

- **User Registration and Login** 🔐
- **Post Creation** 📝
- **Commenting on Posts** 💬
- **Private Messaging** 💌

The forum includes five distinct components:

- **SQLite**: Data storage 💾
- **Golang**: Data handling and Websockets (Backend) 🏗️
- **JavaScript**: Handling all Frontend events and client Websockets 🌐
- **HTML**: Structuring the page elements 🏷️
- **CSS**: Styling the page elements 🎨

This project uses a single HTML file, so all page changes are managed in JavaScript (Single Page Application).


## Tools & Skills 🧰

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-003B57?style=for-the-badge&logo=sqlite&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-000000?style=for-the-badge&logo=javascript&logoColor=F7DF1E)
![HTML5](https://img.shields.io/badge/HTML5-000000?style=for-the-badge&logo=html5&logoColor=E34F26)
![CSS3](https://img.shields.io/badge/CSS3-000000?style=for-the-badge&logo=css3&logoColor=1572B6)
![Git](https://img.shields.io/badge/Git-000000?style=for-the-badge&logo=git&logoColor=F05032)
![Gitea](https://img.shields.io/badge/Gitea-34495E?style=for-the-badge&logo=gitea&logoColor=5D9425)
![GitHub Copilot](https://img.shields.io/badge/githubcopilot-%23026AA7.svg?style=for-the-badge&logo=githubcopilot&logoColor=white)
![ChatGPT](https://img.shields.io/badge/chatGPT-74aa9c?style=for-the-badge&logo=openai&logoColor=white)
![YouTube](https://img.shields.io/badge/YouTube-%23FF0000.svg?style=for-the-badge&logo=YouTube&logoColor=white)
![Markdown](https://img.shields.io/badge/Markdown-000000?style=for-the-badge&logo=markdown&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)

## 📁 Project Architecture

This project employs a microservices architecture. Each service is responsible for a specific functionality of the application. Here is the overall project structure:



```
.
├── backend
│ ├── cmd
│ │ ├── app
│ │ │ └── main.go
│ │ └── bootservices
│ │ └── main.go
│ ├── orm
│ ├── server
│ │ ├── gateway
│ │ ├── microservices
│ │ ├── middleware
│ │ └── router
│ ├── services
│ │ ├── auth
│ │ │ ├── controllers
│ │ │ ├── database
│ │ │ │ └── migrates
│ │ │ └── models
│ │ ├── chat
│ │ │ ├── controllers
│ │ │ ├── database
│ │ │ │ └── migrates
│ │ │ └── models
│ │ ├── notification
│ │ │ ├── controllers
│ │ │ ├── database
│ │ │ │ └── migrates
│ │ │ └── models
│ │ └── posts
│ │ ├── controllers
│ │ ├── database
│ │ │ └── migrates
│ │ └── models
│ └── utils
│ ├── jwt
│ ├── key
│ └── validation
│ └── test
├── frontend
│ └── assets
│ ├── css
│ │ └── img
│ ├── images
│ └── js
│ ├── api
│ ├── components
│ ├── pages
│ ├── router
│ └── utils
```

### 🏗️ Architecture Details

1. **Microservices**
    - Each microservice is autonomous and responsible for a specific part of the application (authentication, chat, notifications, posts). 🔄
    - Defined in the `backend/services` directory.
    - Interaction via REST APIs. 🌐

2. **Gateway**
    - A single entry point for all requests managed by the gateway located in `backend/server/gateway`. 🚪
    - Routes requests to the appropriate microservices based on the URL. 🔀

3. **Websockets**
    - Handles real-time messaging in chat and notifications. 💬
    - Implemented in Golang (backend) and JavaScript (frontend). 🌐

4. **JWT Authentication**
    - Secures endpoints using JSON Web Tokens (JWT). 🔐
    - Custom JWT implementation in `backend/utils/jwt`.

### 🔧 Services

- **Auth**: Manages user registration, login, and logout. 🧑‍🤝‍🧑
- **Chat**: Manages private messages between users. 💬
- **Notification**: Manages real-time notifications. 📲
- **Posts**: Manages post creation and comments. 📝

## 🚀 Usage

**Start the services in following order**

```bash
cd backend/cmd/app
go run main.go

cd backend/cmd/bootservices
go run main.go
```

Open your browser and go to http://localhost:3000 🌐

## 🎉 Conclusion

This project uses a microservices architecture to separate responsibilities and facilitate maintenance and scalability. Each service is autonomous and communicates with others via REST APIs. Websockets enable real-time communication for chat and notifications. Authentication is secured using JWT.
=======
# real-time-forum
A real-time forum application built with Golang, SQLite, and JavaScript, featuring user registration, post creation, commenting, and private messaging.
>>>>>>> 80733a1a796ddee80876801792b213d191829fb6
