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
[[{"class": "som-tam", "probability": 99.84957885742188}], [{"class": "nam-prik-oong", "probability": 0.04002401977777481}], [{"class": "khao-tom-mud", "probability": 0.017101913690567017}], [{"class": "kang-khee-lek", "probability": 0.014392317272722721}], [{"class": "nam-prik-nhum", "probability": 0.011167040094733238}]]
```