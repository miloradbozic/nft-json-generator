# NFT Json Generator
## Metaplex Candy machine NFT JSON files generator

[![LCT Labs](https://lct-labs-public.s3.amazonaws.com/logo.png)](https://lctlabs.io)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

NFT Json generator allows you to generate arbitrary number of JSON files containing metadata for your NFTs
NFT Json Generator is generator JSON files representing metadata for NFTs on Solana blockchain following Metaplex metadata standard.
It is written in Go thus it is very fast in comparison to some other solutions written in Javascript etc.

- ✨Magic ✨
## Features, Usage

- Provide CSV file containing description of your NFTs
- Add wallet addresses of creators as parameters
- Run program and find your JSOn files in output folder

This text you see here is *actually- written in Markdown! To get a feel
for Markdown's syntax, type some text into the left window and
watch the results in the right.

## CSV Format

Input file for the generator is a comma separated csv file which should have trait names as the first column. Example of input csv file:

| Color | Age | Rarity |
| ------ | ------ | ------ |
| Red | Young | Common |
| Blue | Young | Common |
| Golden | Young | Medium |
| Cameleon | Old | Rare |


## Installation

Dillinger requires [Go](https://golang.org/) v? to run.

Install the dependencies and devDependencies and start the server.

```sh
cd nft-json-generator
go run main.go <path_to_csv> <wallet1> <wallet2> ...
```

## Development

Want to contribute? Great!


## Docker

Nft-Json-Generator is very easy to install and deploy in a Docker container.

```sh
docker run -d -p 8000:8080 --restart=always --cap-add=SYS_ADMIN --name=dillinger <youruser>/dillinger:${package.json.version}
```




## License

MIT

**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [dill]: <https://github.com/joemccann/dillinger>
   [git-repo-url]: <https://github.com/joemccann/dillinger.git>
   [john gruber]: <http://daringfireball.net>
   [df1]: <http://daringfireball.net/projects/markdown/>
   [markdown-it]: <https://github.com/markdown-it/markdown-it>
   [Ace Editor]: <http://ace.ajax.org>
   [node.js]: <http://nodejs.org>
   [Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
   [jQuery]: <http://jquery.com>
   [@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
   [express]: <http://expressjs.com>
   [AngularJS]: <http://angularjs.org>
   [Gulp]: <http://gulpjs.com>

   [PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
   [PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
   [PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
   [PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
   [PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
   [PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>
