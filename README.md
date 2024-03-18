### Store CLI
  Store is a CLI tool developed using Cobra and the Go programming language. It perform  file operations using command line. This project is implemented with uploading file, deleting file, list the current files in the server.


## Commands
```
  store add: This command will upload files to the fileserver.  
 
  store ls: This command will list the files from the fileserver.
 
  store wc: This command will count total number of words in the fileserver.
  
  store rm: This command will delete the file from the fileserver.
  
  store update: This command will update the file in the fileserver.
 
  store freq-words: This command will retrieves the most frequent word used across all uploaded documents from the fileserver.
```

## Installation

   Set up environment variables:
 

# API URL for the Text Store project

```
   export STORE_URL="http://localhost:8080"
```

1. Clone the repository:
```
   https://github.com/mrajkumar21/backed-file-store.git

   cd backed-file-store/manifest/misc
   
   To run server ./store-project
   
   To run cli ./store
   
2. Install on kubernetes.
```
	kubectl apply -f manifest/deployment.yaml
```   

## Testing 
    
# To create files on the server:
```
   store add "filename"
   
   eg: store add file.txt
   
```
#To list the files from the fileserver.:
```
   store ls
````
# To delete the file from the fileserver.:
```
   store rm "filename"
   
   eg: store rm file.txt
````
# To count total number of words in the fileserver. 
```
   store wc
```
# To get the most frequently used words from the uploaded documents, RUN:
```
   store freq-words
```

##Testing on Kubernetes

# To create files on the server:
```
  curl -X POST -F "files=@test.txt" -F "files=@test.txt"  http://"svc-ip":8080/add
   
   Response:
     Files uploaded successfully
   
```
#To list the files from the fileserver.:
```
   curl -X GET  http://"svc-ip":8080/list
   
    Response:
   {"files":["test.txt"]}
````
# To delete the file from the fileserver.:
```
   curl -X DELETE -d  '{"files": ["test.txt"] }'  http://"svc-ip":8080/delete
   
   Response:
   File test.txt deleted successfully
````
# To count total number of words in the fileserver. 
```
  curl -X GET  http://"svc-ip":8080/wordcount
  
  Response:
 {"WordCount":1}
```



