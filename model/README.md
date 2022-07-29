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
[
    [
        {
            "class": "\u0e2a\u0e49\u0e21\u0e15\u0e33",
            "probability": 99.84957885742188
        }
    ],
    [
        {
            "class": "\u0e19\u0e49\u0e33\u0e1e\u0e23\u0e34\u0e01\u0e2d\u0e48\u0e2d\u0e07",
            "probability": 0.04002401977777481
        }
    ],
    [
        {
            "class": "\u0e02\u0e49\u0e32\u0e27\u0e15\u0e49\u0e21\u0e21\u0e31\u0e14",
            "probability": 0.017101913690567017
        }
    ],
    [
        {
            "class": "\u0e41\u0e01\u0e07\u0e02\u0e35\u0e49\u0e40\u0e2b\u0e25\u0e47\u0e01",
            "probability": 0.014392317272722721
        }
    ],
    [
        {
            "class": "\u0e19\u0e49\u0e33\u0e1e\u0e23\u0e34\u0e01\u0e2b\u0e19\u0e38\u0e48\u0e21",
            "probability": 0.011167040094733238
        }
    ]
]
```