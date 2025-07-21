# HTML Comprehensive Summary

HTML (HyperText Markup Language) is the standard language for creating webpages. It provides the structure of a web page, allowing you to define elements such as text, images, links, forms, and more.

---

## 📚 Table of Contents

1. [Introduction to HTML](#1-introduction-to-html)
2. [HTML Document Structure](#2-html-document-structure)
3. [Common HTML Tags](#3-common-html-tags)
4. [Text Formatting](#4-text-formatting)
5. [Lists in HTML](#5-lists-in-html)
6. [Links and Images](#6-links-and-images)
7. [Tables](#7-tables)
8. [Forms](#8-forms)
9. [Semantic HTML](#9-semantic-html)
10. [HTML Entities](#10-html-entities)
11. [Media Elements](#11-media-elements)
12. [HTML5 Features](#12-html5-features)
13. [Cheatsheet](#13-cheatsheet)

---

## 1. Introduction to HTML

- **HTML** is not a programming language; it's a markup language.
- It structures content and elements on the web using tags.
- Tags are enclosed in angle brackets: `<tagname>...</tagname>`.

---

## 2. HTML Document Structure

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>My First Page</title>
</head>
<body>
  <h1>Hello World!</h1>
</body>
</html>
```

| Part       | Purpose                          |
|------------|----------------------------------|
| `<!DOCTYPE>` | Declares HTML5 version         |
| `<html>`     | Root element                   |
| `<head>`     | Metadata (not displayed)       |
| `<title>`    | Page title (in browser tab)    |
| `<body>`     | Visible page content           |

---

## 3. Common HTML Tags

| Tag         | Description                      |
|-------------|----------------------------------|
| `<h1>` to `<h6>` | Headings (h1 is largest)     |
| `<p>`       | Paragraph                        |
| `<br>`      | Line break                       |
| `<hr>`      | Horizontal rule                  |
| `<strong>`  | Bold (semantic)                  |
| `<em>`      | Italic (semantic)                |
| `<div>`     | Generic block container          |
| `<span>`    | Generic inline container         |

---

## 4. Text Formatting

| Tag         | Effect                           |
|-------------|----------------------------------|
| `<b>`       | Bold                             |
| `<i>`       | Italic                           |
| `<u>`       | Underline                        |
| `<mark>`    | Highlighted text                 |
| `<sub>`     | Subscript                        |
| `<sup>`     | Superscript                      |
| `<code>`    | Inline code                      |
| `<pre>`     | Preformatted text (respects spaces) |

---

## 5. Lists in HTML

### Unordered List
```html
<ul>
  <li>Item 1</li>
  <li>Item 2</li>
</ul>
```

### Ordered List
```html
<ol>
  <li>First</li>
  <li>Second</li>
</ol>
```

### Description List
```html
<dl>
  <dt>HTML</dt>
  <dd>HyperText Markup Language</dd>
</dl>
```

---

## 6. Links and Images

### Links
```html
<a href="https://example.com">Visit Site</a>
```

### Images
```html
<img src="image.jpg" alt="Description" />
```

| Attribute | Description                         |
|-----------|-------------------------------------|
| `href`    | URL to link                         |
| `target="_blank"` | Opens link in new tab       |
| `src`     | Path to image                       |
| `alt`     | Alternative text for accessibility  |

---

## 7. Tables

```html
<table>
  <tr>
    <th>Name</th>
    <th>Age</th>
  </tr>
  <tr>
    <td>Alice</td>
    <td>30</td>
  </tr>
</table>
```

| Tag     | Purpose               |
|---------|-----------------------|
| `<table>` | Table container     |
| `<tr>`    | Table row           |
| `<th>`    | Table header cell   |
| `<td>`    | Table data cell     |

---

## 8. Forms

```html
<form action="/submit" method="post">
  <label for="name">Name:</label>
  <input type="text" id="name" name="name" />
  <input type="submit" value="Submit" />
</form>
```

| Input Type     | Description                  |
|----------------|------------------------------|
| `text`         | Single-line input            |
| `email`        | Email field with validation  |
| `password`     | Password field               |
| `radio`        | Radio button                 |
| `checkbox`     | Checkbox                     |
| `submit`       | Submit button                |
| `textarea`     | Multi-line text              |
| `select`       | Dropdown selection           |

---

## 9. Semantic HTML

Semantic tags describe their meaning clearly:

| Tag         | Meaning                        |
|-------------|--------------------------------|
| `<header>`  | Page or section header         |
| `<nav>`     | Navigation links               |
| `<main>`    | Main content of document       |
| `<section>` | A section of related content   |
| `<article>` | Independent article or entry   |
| `<aside>`   | Sidebar or complementary info  |
| `<footer>`  | Footer of section/page         |

---

## 10. HTML Entities

Used to display reserved or special characters.

| Symbol | Code     |
|--------|----------|
| `<`    | `&lt;`   |
| `>`    | `&gt;`   |
| `&`    | `&amp;`  |
| `"`    | `&quot;` |
| `'`    | `&apos;` |
| `©`    | `&copy;` |
| `→`    | `&rarr;` |

---

## 11. Media Elements

### Audio
```html
<audio controls>
  <source src="sound.mp3" type="audio/mpeg" />
</audio>
```

### Video
```html
<video width="320" height="240" controls>
  <source src="movie.mp4" type="video/mp4" />
</video>
```

---

## 12. HTML5 Features

- `<video>` and `<audio>` support
- `<canvas>` for 2D graphics
- `<section>`, `<article>`, `<nav>`, `<footer>`
- Local storage via `localStorage` and `sessionStorage`
- Geolocation API

---

## 13. Cheatsheet

### Tags Summary

| Element      | Description             |
|--------------|-------------------------|
| `<html>`     | Root HTML element       |
| `<head>`     | Metadata section        |
| `<title>`    | Page title              |
| `<body>`     | Page content            |
| `<h1>`–`<h6>`| Headings                |
| `<p>`        | Paragraph               |
| `<a>`        | Anchor/Link             |
| `<img>`      | Image                   |
| `<ul>`, `<ol>`, `<li>` | Lists         |
| `<form>`     | Form                    |
| `<input>`    | Input field             |
| `<table>`    | Table                   |

### Useful Attributes

| Attribute  | Used In       | Description                          |
|------------|---------------|--------------------------------------|
| `href`     | `<a>`         | Link target                          |
| `src`      | `<img>`, `<video>`, `<audio>` | Source path   |
| `alt`      | `<img>`       | Alt text                             |
| `action`   | `<form>`      | URL to submit to                     |
| `method`   | `<form>`      | GET or POST                          |
| `type`     | `<input>`     | Text, password, radio, etc.          |
| `id`       | Any element   | Identifier                           |
| `class`    | Any element   | CSS class name                       |

---

🧠 **Tip**: Use a code editor like VS Code with Emmet for fast HTML scaffolding!

---

## 📎 Conclusion

HTML is the backbone of web development. With a good understanding of structure, elements, and attributes, you can build well-structured, accessible, and user-friendly websites. Once you're comfortable with HTML, move on to CSS and JavaScript to enhance your web development skills.


