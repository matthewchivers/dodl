# dodl

`dodl` is a command-line tool designed to streamline the creation of structured documents quickly and efficiently. Whether you're in a meeting, attending a lecture, or brainstorming new ideas, `dodl` helps you generate well-organized documents with the correct templates, stored in the appropriate folders, and pre-filled with key information—all in an instant.

## Features

- **Dynamic Templating**: Build documents using placeholders like `{{ .Today }}`, `{{ .Topic }}`, and more, creating structured notes in seconds.
- **Dynamic Directory Structuring**: Automatically organize documents into the right folders based on context, ensuring everything is always where it belongs.
- **Flexible Date Input**: Need a document templated for a specific date? Specify dates in standard formats like `DD/MM/YYYY`. *Natural language date parsing (e.g., "next Monday", "three weeks from now") is coming soon!*
- **Seamless Experience**: Focus on simplicity—no unnecessary prompts or inputs. You get exactly what you need, fast.
- **Customizable Workflows**: Configure templates and file structures to suit personal and project-specific needs.

## Getting Started

### Installation

1. Clone the repo
1. Run "Make install"

> Prerequisites: `golang` and `make`

### Set Up Your Workspace

Initialize your workspace by running `dodl init` in your desired directory. This command creates a `.dodl` directory, marking it as the root of your workspace:

```bash
dodl init
```

If you prefer to initialize a workspace in a different location, specify the path:

```bash
dodl init /path/to/your/workspace
```

### Create a Document

Need meeting notes, a project report, or any other type of document? Use the `create` command with the document type you've configured:

```bash
dodl create [document_type]
```

For example, if you have a document type called `projectX` defined in your `config.yaml`, you can create it with:

```bash
dodl create projectX
```

You can also add details like date and topic to customize it further:

```bash
dodl create projectX -d "29/10/2024" -t "Weekly Update"
```

### Check Workspace Status

To view your workspace setup, run:

```bash
dodl status
```

This command displays the active configuration, available templates, and workspace root path.

## The `dodl` Workspace

When you run `dodl init`, a `.dodl` directory is created in your specified location. This directory serves as the root of your workspace and includes:

- `config.yaml`: Defines document types, templates, and behavior.
- `templates` directory: Holds your document templates.

All `dodl` commands (`create`, `status`, etc.) executed within this workspace will use the configurations and templates from this root, ensuring consistent document creation and organization.

## Configuration and Templates

Customize `dodl` to fit your workflow by editing the `config.yaml` file and adding templates in the `templates` directory.

### `config.yaml` Example

```yaml
default_document_type: journal
document_types:
  journal:
    template_file: "journal.md"
    file_name_pattern: "{{ .Today | date \"2006-01-02\" }}.md"
    directory_pattern: "{{ .Today | date \"2006\" }/{{ .Today | date \"January\" }}"
    custom_values:
      author: "Your Name"
```

### Template Files

Templates use Go's `text/template` syntax, allowing you to include dynamic content based on the context. Place your template files in the `templates` directory.

Example `journal.md` template:

```
# Journal Entry - {{ .Today | date "02 January 2006" }}

**Author**: {{ .author }}

## Thoughts

...
```

## The Problem `dodl` Solves

In spontaneous situations where you need to take notes or document ideas, setting up a structured document can be a hassle. `dodl` eliminates this friction by providing you with a ready-to-use document tailored to your needs, so you can focus on capturing information without worrying about formatting or file organization.

## Status

`dodl` is currently in active development. The core features—including dynamic templating, configuration, and document creation—have been implemented. Natural language date parsing is planned for a future release. For now, you can specify custom dates (at runtime) using standard formats like `DD/MM/YYYY`. We welcome feedback and contributions to continue improving the tool.

## Contributing

Contributions are welcome! If you'd like to contribute, please fork the repository and submit a pull request. For major changes, please open an issue first to discuss what you'd like to change.

## License

`dodl` is released under the MIT License, ensuring it's free to use, modify, and share (with attribution).

---

Feel free to explore and customize `dodl` to suit your workflow. If you have any questions or suggestions, please open an issue on the [GitHub repository](https://github.com/matthewchivers/dodl).
