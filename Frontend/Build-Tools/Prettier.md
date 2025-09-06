# Prettier

## Installation
- npm i -D --save-exact prettier
- Should be '--save-exact' so when there's new version of prettier then it will not override our app's formatting if there are changes in prettier's way of formatting.

## CLI usage
- npx prettier --write <filename>

## Vs Code Extension
- Local code prettier config will take precedence over vscode ext.
- 

## Watching for Changes
```
npx onchange "**/*" -- npx prettier --werite --ignore-unknown {{changed}}
```

<br>
- package.json
```
{
  "scripts": {
    "prettier-watch": "onchange \"**/*\" -- prettier --write --ignore-unknown {{changed}}"
  }
}

```

<br>
## Ignoring Files
- .prettierignore


<br>
## Configuration

1. in package.json add 'pretter' 
```
{

 "name": "project-name",
 "scripts" : ....,
 "prettier": { ... } 
} 
```


2. adding '.prettierrc' 

```
{
  "semi" : true,
  "overrides": [
    {
     "files": ".ts",
     "options" : {
 	"semi": false
      }
     }
   
  ] 
}

```

<br>

## Conflict with ESLint
- install 'eslint-config-prettier'
- .eslintrc extends 'prettier'
- npx eslint-config-prettier <file.js>  // check rules conflict


