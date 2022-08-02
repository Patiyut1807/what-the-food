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

### Running
```bash
python app.py --model 0 --img test.jpeg
```
```
usage: app.py [-h] [--model MODEL] [--img IMG]

optional arguments:
  -h, --help     show this help message and exit
  --model MODEL  What kind of model you want to classified 0:food, 1:animal
  --img IMG      Image file name to predict
```
This will produce json file `output.json`
```json
[{"class": "som-tam", "probability": 99.8066635131836}, {"class": "nam-prik-oong", "probability": 0.05312320962548256}, {"class": "nam-prik-nhum", "probability": 0.01790415309369564}, {"class": "pad-thai", "probability": 0.01440698653459549}, {"class": "khao-tom-mud", "probability": 0.014378681778907776}]
```