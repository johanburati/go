# go-face-demo1
Testing faces recognition in Go using [go-face](https://github.com/Kagami/go-face) library.

[go-face](https://github.com/Kagami/go-face) implements face recognition for Go using [dlib](http://dlib.net/)'s state-of-the-art face recognition built with deep learning. The model has an accuracy of 99.38% on the [Labeled Faces in the Wild](http://vis-www.cs.umass.edu/lfw/) benchmark.

# Demo

First the program will detect the faces from left to right in this group photo of the [Seinfeld](https://www.imdb.com/title/tt0098904/) cast, and we will name the faces accordingly.

| ![](seinfeld.jpg) |
| --- |
| The *Seinfeld* cast: Kramer, Jerry, Elaine and George |

Then program will scan the face in this recent picture of the each actor and determine if the face appears in the group photo.

| ![](kramer.jpg) | ![](jerry.jpg) | ![](elaine.jpg) | ![](george.jpg) |
| --- | --- | --- | --- |
| Kramer | Jerry | Elaine | George |

Here are the results:

```
$ go-face-demo1 --group seinfeld.jpg --names Kramer,Jerry,Elaine,George --target kramer.jpg
# GO-FACE DEMO #1
## Detect faces on seinfeld.jpg
- Detected: Kramer
- Detected: Jerry
- Detected: Elaine
- Detected: George
## Recognize face on kramer.jpg
- Found: Kramer

$ go-face-demo1 --group seinfeld.jpg --names Kramer,Jerry,Elaine,George --target jerry.jpg
# GO-FACE DEMO #1
## Detect faces on seinfeld.jpg
- Detected: Kramer
- Detected: Jerry
- Detected: Elaine
- Detected: George
## Recognize face on jerry.jpg
- Found: Jerry

$ go-face-demo1 --group seinfeld.jpg --names Kramer,Jerry,Elaine,George --target elaine.jpg
# GO-FACE DEMO #1
## Detect faces on seinfeld.jpg
- Detected: Kramer
- Detected: Jerry
- Detected: Elaine
- Detected: George
## Recognize face on elaine.jpg
- Found: Elaine

$ go-face-demo1 --group seinfeld.jpg --names Kramer,Jerry,Elaine,George --target george.jpg
# GO-FACE DEMO #1
## Detect faces on seinfeld.jpg
- Detected: Kramer
- Detected: Jerry
- Detected: Elaine
- Detected: George
## Recognize face on george.jpg
- Found: George

```

Success, quite amazing !

The program could recognize each actor even though they are a bit older now, wear glasses or change hairstyle.
