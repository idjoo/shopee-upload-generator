# shopee-upload-generator

A program that helps you generate [Mass Upload](https://seller.shopee.co.id/edu/article/112) file
from Mass Update files. Written in pure go.

## Features
- Very Fast
- Read file concurrently

## How it works
This program will read 4 file:  
- Mass Update basic info
- Mass Update sales
- Mass Update shipping
- Mass Update media

The program will then generate it to an existing empty Mass Upload file.
