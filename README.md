# `dodl`

`dodl` is a command-line tool designed to save you time when creating structured plaintext documents.  Instead of spending time manually setting up your file system and formatting each document, `dodl` automates these steps, allowing you to focus on the content itself.  With just one command, `dodl` generates your document and places it in the right location, using templates you define.

### What Problem Does `dodl` Solve?

Creating and organising notes can be a time-consuming task, especially when you're trying to maintain consistency.  Imagine being in a meeting and needing to take notes quickly: before you know it, you've spent valuable time managing file names, directories, and document structure, all while the conversation moves forward.

`dodl` eliminates this overhead.  It automates the creation of structured documents, so you don’t have to worry about directory organisation or file formatting in the moment.  This lets you focus on what matters most: getting your ideas down fast and efficiently.

If you have any questions or suggestions, feel free to open an issue on [GitHub](https://github.com/matthewchivers/dodl).

## Get Started

Here's how to get started with `dodl` in just a few steps:

### Installation

You can install `dodl` easily using Homebrew:

```shell
brew tap matthewchivers/dodl
brew install dodl
```

### Initialise Your Workspace

To start using `dodl`, navigate to a directory where you'd like to store your documents, and initialize the workspace by running:

```shell
dodl init
```

This creates a `.dodl` directory in your current location, which stores all of the configuration files and templates that `dodl` uses to generate documents.

### Create a Document

Once your workspace is set up, you can create a document by specifying the document type you’ve defined in your `config.yaml` (see [configure your workspace](#configure-your-workspace)):

```shell
dodl create [document-type]
```

You can also override specific fields at runtime with custom values. For example:

```shell
dodl create note --date "21/03/2027" --topic "Milestone Birthday Planning"
```

For more on how the fields work and why they’re useful, see [Templating](#templating) below.

## Templating

`dodl` includes a powerful templating system that allows you to dynamically generate content in your documents, directories, and filenames. By using placeholders like `{{ .Now }}` or `{{ .Topic }}`, you can create consistent structures across your files. It’s simple to get started with, but also flexible enough to handle more complex needs.

#### Examples:

Fields can be configured (via `config.yaml` or overridden at runtime) and used as placeholders in your templates. For example, you can access the `Topic` field or use calculated fields like `Now`. Here are a few examples:

* `{{ .Topic }}` – If your `Topic` field in `config.yaml` is "Project X", this will output: `Project X`.
* `{{ .Topic | upper }}` – This will output the `Topic` in uppercase: `PROJECT X`.
* `{{ .Now | date "2006/01" }}` – Outputs the current date in `yyyy/mm` format (e.g., `2024/10`).
* `{{ addDays .Now 6 | date "02-01-06" }}` – If today is `28-Oct-24`, this will add six days and output `03-11-24`.

The templating system uses Go’s `text/template` library along with the [Sprig functions](http://masterminds.github.io/sprig/), which add extra functionality like `upper` and `date` for date formatting and other dynamic operations.

## Configure Your Workspace

Customise `dodl` by editing `.dodl/config.yaml` and adding templates to `.dodl/templates`. This allows you to tailor document types, fields, and storage patterns to match your workflow.

### Example `config.yaml`
``` yaml
default_document_type: journal
custom_fields:
  author: "Full Name"
document_types:
  journal:
    template_file: "journal.md"
    file_name_pattern: "journal-{{ .Now | date \"02-01-2006\" }}.md" # e.g. journal-28-10-24.md
    directory_pattern: # e.g. 2024/October/journal
      - "{{ .Now | date \"2006\" }}"
      - "{{ .Now | date \"January\" }}"
      - "journal"
    custom_fields:
      author: "First Name"
  meeting:
    template_file: "meeting.md"
    file_name_pattern: "{{ .Topic }}-{{ .Now | date \"02-01-2006\" }}.md"
    directory_pattern: [ "{{ .Now | date \"2006\"}}", "{{ .Now | date \"01\"}}" ]
```

The above config sets up a default document type of `journal`, which will create files such as `journal-28-10-24.md` in a directory such as `2024/10/journal.`

### Template Files

Document templates also take advantage of the same placeholder mechanism. For example, a `journal` template based on the example configuration above might look like:

``` md
#  Journal Entry - {{ .Now | date "02 January 2006" }}

**Author**: {{ .author }}

## Daily Thoughts
...
```

## Status

`dodl` is in active development, with core features like templating, configuration, and document creation already implemented. Natural language date parsing is coming soo.  Contributions and feedback are welcome!

## Contributing

Open-source is good for the world, and contributions are encouraged.  Fork the repository and submit a pull request!  For major changes, please open an issue first to discuss your ideas.

## License
`dodl` is released under the MIT license, making it free to use, modify, and share.  Requires attribution to Matthew Chivers (the original author of `dodl`).
