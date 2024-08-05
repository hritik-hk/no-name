## Go App to aggregate your RSS feeds 
##### App optimized using multithreading to continuosly scrape rss feeds stored in Db and then store new feed post to your database for different usecases.
![Screenshot from 2024-08-04 22-03-43](https://github.com/user-attachments/assets/4253be28-3c13-4d2c-a0a0-564da5d3be63)

### How to run?
Add the required .env config
 1. `using go:` 
```
go build -o main && ./main
```
 2. `using docker:` 
```
docker build -t your_docker_image_name .
docker run -p 7777:7777 your_docker_image_name
```


