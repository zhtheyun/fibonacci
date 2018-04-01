## Write RAML with api-designer

### Install
```
npm install -g request
npm install -g api-designer
```

### run the server
```
api-designer -p 4000
```

### start the browser and visit localhost:4000


## raml2html

A simple RAML to HTML documentation generator, written for Node.js, with theme support.

### RAML version support
raml2html 4 and higher only support RAML 1.0 files. Please stick with raml2html 3.x for RAML 0.8 support.


### Install
```
npm i -g raml2html
```


### Themes
raml2html ships with a default theme, but you can install more from NPM. For example, to render
RAML to Markdown, you can install the raml2html-markdown-theme theme:

```
npm i -g raml2html-markdown-theme
```

Search NPM for the "raml2html-theme" keyword (or use [this link](https://www.npmjs.com/browse/keyword/raml2html-theme))
to find more themes.

### Usage

#### As a command line script
```
raml2html --help
raml2html example.raml > example.html
raml2html --theme raml2html-markdown-theme example.raml > example.html
raml2html --template my-custom-template.nunjucks -i example.raml -o example.html
```
