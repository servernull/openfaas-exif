# openfaas-exif

An [openfaas](https://www.openfaas.com) function that takes a base64 encoded image, and returns EXIF information.

```bash

### deploy
faas-cli deploy -f stack.yml

### invoke
cat image.jpg | base64 | faas-cli invoke openfaas-exif | jq

[
  {
    "Make": "samsung"
  },
  {
    "Model": "SM-N920T"
  },
  {
    "Orientation": "1"
  },
  {
    "XResolution": "72/1"
  },
  {
    "YResolution": "72/1"
  },
  {
    "ResolutionUnit": "2"
  },
  {
    "Software": "GIMP 2.8.20"
  },
  {
    "DateTime": "2018:06:09 01:07:30"
  },
  {
    "YCbCrPositioning": "1"
  },
  {
    "ExifTag": "212"
  },
  {
    "ExposureTime": "1/13"
  },
  {
    "FNumber": "19/10"
  },
  {
    "ExposureProgram": "2"
  },
  {
    "ISOSpeedRatings": "200"
  },
  {
    "ExifVersion": "0220"
  },
  {
    "DateTimeOriginal": "2018:04:28 21:23:12"
  },
  {
    "DateTimeDigitized": "2018:04:28 21:23:12"
  },
  {
    "ShutterSpeedValue": "374/100"
  },
  {
    "ApertureValue": "185/100"
  },
  {
    "BrightnessValue": "-57/100"
  },
  {
    "ExposureBiasValue": "0/10"
  },
  {
    "MaxApertureValue": "185/100"
  },
  {
    "MeteringMode": "2"
  },
  {
    "Flash": "0"
  },
  {
    "FocalLength": "430/100"
  },
  {
    "UserComment": "UserComment<SIZE=(13) ENCODING=[ASCII] V=[0 0 0 73 73 67 83 65]... LEN=(13)>"
  },
  {
    "FlashpixVersion": "0100"
  },
  {
    "ColorSpace": "1"
  },
  {
    "PixelXDimension": "920"
  },
  {
    "PixelYDimension": "570"
  },
  {
    "InteroperabilityTag": "690"
  },
  {
    "InteroperabilityIndex": "R98"
  },
  {
    "InteroperabilityVersion": "0100"
  },
  {
    "ExposureMode": "0"
  },
  {
    "WhiteBalance": "0"
  },
  {
    "FocalLengthIn35mmFilm": "28"
  },
  {
    "SceneCaptureType": "0"
  },
  {
    "ImageUniqueID": "b2f7216f7a3e04a00000000000000000"
  },
  {
    "GPSTag": "720"
  },
  {
    "GPSVersionID": "0x02"
  },
  {
    "GPSLatitudeRef": "N"
  },
  {
    "GPSLatitude": "26/1"
  },
  {
    "GPSLongitudeRef": "W"
  },
  {
    "GPSLongitude": "80/1"
  },
  {
    "GPSAltitudeRef": "0x01"
  },
  {
    "GPSAltitude": "0/1"
  },
  {
    "GPSTimeStamp": "1/1"
  },
  {
    "GPSDateStamp": "2018:04:29"
  },
  {
    "Compression": "6"
  },
  {
    "XResolution": "72/1"
  },
  {
    "YResolution": "72/1"
  },
  {
    "ResolutionUnit": "2"
  }
]

```
