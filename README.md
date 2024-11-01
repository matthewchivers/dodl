# dodl

`dodl` is a command-line tool for creating structured documents quickly and effortlessly.  With a few simple comamnds, `dodl` generates documents and automatically organises them in appropriate directories, using templates you define.

`dodl` supports standard plaintext documents such as `.txt` and `.md` (markdown).

### What Problem Does `dodl` Solve?

Creating structured documents can be time-consuming.  Setting up folders, creating filenames, and ensuring consistent formatting takes effort. `dodl` automates these repetitive tasks, so you can focus directly on writing and getting work done.

Explore `dodl` today and simplify your workflow.  For questions or suggestions, please open an issue on [GitHub](https://github.com/matthewchivers/dodl).

## Get Started

Here's how to get started with `dodl` in just a few steps:

### Installation

Make sure you've got [golang](https://go.dev/doc/install) installed (preferably `1.23`) and then run:

```
go install github.com/matthewchivers/dodl@latest
```

### Initialise Your Workspace

Navigate to a directory that is to be your `dodl` workspace and run:
```
dodl init
```

This command creates a `.dodl` directory in your current location.  This directory holds all configurations and templates that `dodl` uses to operate.

### Create a document

```
dodl create [document-type]
```

Replace `[document-type]` with any document type you've defined in `config.yaml` (see [configure your workspace](#configure-your-workspace)).

Override fields at runtime for custom values, example:

```
dodl create journal -d "20/03/2027" -t "Milestone Birthday Planning"
```

The document will be saved based on your templates, e.g. `workspace/2027/March/20/journal-20-03-27.md`.

## Templating

Use placeholders like `{{ .Now }}` or `{{ .Topic }}` to dynamically generate content in document templates, directory structures, or filenames.  `dodl` templating makes it easy to have consistent file naming, directory structuring, and document formatting.  The templating is simple to start with, but can be extended in powerful ways (see examples below).

#### Examples:
* `{{ .Topic }}` where the Topic field is "Project X" -> `Project X`
* `{{ .Topic | upper }}` where the Topic field is "Project X" -> `PROJECT X`
* `{{ .Now | date "2006/01" }}` uses golang date formatting -> `2024/10`
* `{{ addDays .Now 6 | date "02-01-06 }}` (if today is 28-Oct-24) -> `03-11-24`

Templating uses golang's `text/template` alongside http://masterminds.github.io/sprig/ - which provides functions such as `upper` and `date` seen above.

## Configure Your Workspace

Customise `dodl` by editing `.dodl/config.yaml`, and adding templates to `.dodl/templates`.

### Example `config.yaml`
``` yaml
default_document_type: journal
custom_fields:
  author: "Full Name"
document_types:
  journal:
    template_file: "journal.md"
    file_name_pattern: "journal-{{ .Now | date \"02-01-2006\" }}.md" # e.g. journal-
    directory_pattern: # e.g. 2024/October
      - "{{ .Now | date \"2006\" }}"
      - "{{ .Now | date \"January\" }}"
    custom_fields:
      author: "First Name"
  meeting:
    template_file: "meeting.md"
    file_name_pattern: "{{ .Topic }}-{{ .Now | date \"02-01-2006\" }}.md"
    directory_pattern: [ "{{ .Now | date \"2006\"}}", "{{ .Now | date \"01\"}}" ]
```

### Template Files

Example template using the custom field `author` (as in example config above).

``` md
#  Journal Entry - {{ .Now | date "02 January 2006" }}

**Author**: {{ .author }}

## Thoughts

...
```

## Status

`dodl` is in active development, with core features like templating, configuration, and document creation already implemented. Natural language date parsing is coming soo.  Contributions and feedback are welcome!

## Contributing

Open-source is good for the world, and contributions are encouraged.  Fork the repository and submit a pull request!  For major changes, please open an issue first to discuss your ideas.

## License
`dodl` is released under the MIT license, making it free to use, modify, and share.  Requires attribution to Matthew Chivers (the original author of `dodl`).
