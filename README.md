# Welcome, Paper Hoarder!

If you have a huge pile of paper records like receipts, etc..., you can use `packrat` to help you automatically sort and digitize them.

I started this project after I read [Brian](https://github.com/bketelsen)'s [blog post](https://www.brianketelsen.com/ai-ml-documents-one/) about how he used OCR to digitize a lot of his receipts. I fell in love with the idea because I have a pile of receipts and other paper records that keeps growing and it stresses me out. I also never, ever want to have to go and sort & file them.

So, I followed Brian's lead and built a rough version of the OCR system he talked about in the post. Using [the Microsoft OCR sample code written in C#](https://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/quickstarts/csharp-print-text) as a guide, I was able to throw something together very quickly.

## Usage

Make sure you have the following three environment variables set:

- `FILENAME` - the file that you want to run OCR on
- `AZURE_COG_SVCS_ENDPOINT` - your Azure cognitive services endpoint
- `AZURE_SUBSCRIPTION_ID` - your Azure subscription ID

I like to use [`direnv`](https://direnv.net/) to store these environment variables inside an `.envrc` file in my current working directory, and I highly recommend using it to set these environment variables.
