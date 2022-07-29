## wtf Model

### Project Structure


### Description

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
### Output format
```
[{99.84957885742188, 'ส้มตำ'}, {0.04002401977777481, 'น้ำพริกอ่อง'}, {0.017101913690567017, 'ข้าวต้มมัด'}, {0.014392317272722721, 'แกงขี้เหล็ก'}, {0.011167040094733238, 'น้ำพริกหนุ่ม'}]
```

### Training

This model uses pretrained ViT model.