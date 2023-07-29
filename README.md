# TikTok Interview Test - Instant Messaging System Implementation

## Description

This project is my implementation of a TikTok interview test, where we have been tasked with building an instant messaging system. The requirements for this project were provided as part of the test and served as guidelines for developing the system.

## Main Features

The instant messaging system implementation includes the following features:

1. **Main**: Only develop the backend side of the system, focusing on core message features without the front-end part and the account/authentication part
2. **Instant Messaging**: Users can send and receive instant messages to and from their contacts.
3. **Data storage**: The system should store messages data. Receivers can access this data at any time. At least one database must be used. There is no limitation on data schema design.
4. **Message Delivery**: The system should be able to deliver messages to the intended recipients by PULL mode in a timely and consistent manner.
5. **Must** contain a runnable docker-compose.yml

## Extra features added by myself

> Extra features will be placed in a seperate branch withe name: `additional-features`

1. **User Registration**: Users can create new accounts by providing their desired username, email address, and password.
2. **User Authentication**: Registered users can log in to the system using their credentials.
3. **User Profiles**: Each user has a profile that contains information such as their username, display name, profile picture, and status.
4. **Contact List**: Users can manage their contact list, which consists of other registered users.
5. **Message Notifications**: Users receive real-time notifications when they receive new messages.
6. **Online/Offline Status**: Users can see the online/offline status of their contacts.

## Technologies Used

The implementation of this instant messaging system is built using the following technologies:

- **Backend Framework**: [Go](https://go.dev/)
- **Database**: [Postgres](https://www.postgresql.org/)
- **Cache**: [Redis](https://redis.io/)
- **Real-Time Communication**: [Gorilla Websocket](https://github.com/gorilla/websocket)

## Setup Instructions

To set up and run this instant messaging system on your local machine, follow these instructions:

TBC

## Limitations and Future Improvements

Although this implementation covers the basic requirements of the TikTok interview test, there are some limitations and areas for future improvements, including:

1. Group Messaging: Currently, the system only supports one-on-one messaging. Implementing group messaging functionality would be a valuable addition.
2. Message Encryption: To enhance security, implementing end-to-end encryption for messages would be beneficial.
3. User Blocking: Adding the ability to block and unblock users would provide users with better control over their messaging experience.

### Dependencies

install protoc-gen-go (Protobuf plugin specifically for Go)

- https://github.com/protocolbuffers/protobuf/releases
  ...Or manually download and unzip into a path where you system can find the binary (include it to your $PATH)

install protoc (`protoc` is the actual Protocol Buffers compiler binary. It is responsible for compiling .proto files into language-specific code. The protoc binary is provided by the Protocol Buffers project, and you need to install it separately on your system)

- brew install protoc
- Get the zip file from

### Setup

Make sure you have your $GOPATH, $GOROOT & $GOBIN set in your path. This is so that the `protoc-gen-go` and `protoc` binaries can be reached by your system
