# what-the-food

**what-the-food** or **WTF** is a mobile application that allows users to ~~classify Thai food from their mobile devices.~~

### Project Structure

```
.
├── LICENSE
├── README.md
├── backend
│   ├── Dockerfile
│   ├── README.md
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── model
│   ├── README.md
│   ├── app.py
│   ├── class_food.txt
│   ├── output.json
│   └── test.jpeg
└── wtf_app
    ├── app
    │   ├── build.gradle
    │   ├── proguard-rules.pro
    │   ├── release
    │   │   ├── app-release.apk
    │   │   └── output-metadata.json
    │   └── src
    │       └── main
    .
    .
    └── settings.gradle

36 directories, 48 files
```

`model` : PyTorch inference script

`backend` : go backend

`android-app` : android application

> For more details, you can read through each README for each directory.

### To run `what-the-food` application


### Contributing

We wish to expand this project up and correct all of our faults in order to make our project great.

### License
[MIT](https://choosealicense.com/licenses/mit/)
