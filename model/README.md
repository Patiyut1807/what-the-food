## wtf Model

### Project Structure
```
.
├── app.py
├── class_food.txt
├── model_food.pth
├── output.json
├── README.md
└── test.jpeg

0 directories, 6 files
```

> For `model_food.pth` please follow this link and download it from google drive [model_food.pth](https://drive.google.com/drive/folders/113kF1b1pMHqLYTcLYiL1uxTc2VVzi0N1?usp=sharing) 

### Running
```bash
python app.py --model 0 --img test.jpeg
```
```
usage: app.py [-h] [--model MODEL] [--img IMG]

optional arguments:
  -h, --help     show this help message and exit
  --model MODEL  What kind of model you want to classified 0:food
  --img IMG      Image file name to predict
```

In the future, we hope that we can provide much more model type than just Thai-Food classification.

This will produce json file `output.json`
```json
[{"class": "som-tam", "probability": 99.8066635131836}, {"class": "nam-prik-oong", "probability": 0.05312320962548256}, {"class": "nam-prik-nhum", "probability": 0.01790415309369564}, {"class": "pad-thai", "probability": 0.01440698653459549}, {"class": "khao-tom-mud", "probability": 0.014378681778907776}]
```