from PIL import Image
from torchvision import transforms
import torch
from pytorch_pretrained_vit import ViT
import torch.optim as optim
import argparse

parser = argparse.ArgumentParser()
parser.add_argument('--model', type=int, default=0, help="What kind of model you want to classified 0:food, 1:animal")
parser.add_argument('--img', type=str, help="Image file name to predict")
args = vars(parser.parse_args())

if args['model'] == 0:
    PATH = 'model_food.pth'
    NUM_CLASSES = 16
    CLASS_NAME = 'class_food.txt'
else:
    PATH = 'none'
    NUM_CLASSES = 0
    CLASS_NAME = ''

IMAGE_SIZE = 224
IMAGE_PATH = args['img']
device = ('cuda' if torch.cuda.is_available() else 'cpu')

model = ViT('B_32_imagenet1k',
    dim=128,
    image_size=224,
    num_classes=NUM_CLASSES
)

optimizer = optim.SGD(model.parameters(), lr=0.0001, momentum=0.9)

# load checkpoints
checkpoint = torch.load(PATH, map_location=device)
model.load_state_dict(checkpoint['model_state_dict'])

def predict(image):
    normalize = transforms.Normalize(
            mean=[0.485, 0.456, 0.406],
            std=[0.229, 0.224, 0.225]
            )
    transform = transforms.Compose([
        transforms.Resize((IMAGE_SIZE, IMAGE_SIZE)),
        transforms.ToTensor(),
        normalize
    ])

    # img = Image.open(image)
    batch = torch.unsqueeze(transform(image),0)

    # predict
    model.eval()
    out = model(batch)

    with open(CLASS_NAME) as f:
        classes = [line.strip() for line in f.readlines()]

    prob = torch.nn.functional.softmax(out, dim = 1)[0] * 100
    _, indices = torch.sort(out, descending = True)
    return [(classes[idx], prob[idx].item()) for idx in indices[0][:5]]

image = Image.open(IMAGE_PATH).convert('RGB')
labels = predict(image)

# predicted class
for i in labels:
    name = i[0].split()
    print("Prediction ", name[1], ", Score: ", i[1])