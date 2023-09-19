# Microservice overview
1. This service is designed to handle author-user subscription model
2. An author is a premium author if certain conditions are achieved & upon becoming a premium author the content published is sent to premium users - by doing this, the earnings are split between the authors
3. This service keeps a sync of the author's isPremiumAuthor flag in database.

# Tech stack used:
1. Backend Framework : Golang with Gorilla/Mux
2. Database: Mongodb
3. Message queue :  RabbitMQ
4. Cloud platform used : Google Cloud Platform

# Deployed Here
https://car-app-62a04.uc.r.appspot.com

# High level design of the microservice

![image](https://github.com/Radhika877/pratilipi-author-handler-service/assets/56670369/dfa9e771-5c24-46e1-b41d-1a1225076960)
