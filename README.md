# dodl

`dodl` is an upcoming command-line tool designed to solve a familiar problem: you're in a meeting, a lecture, or someone mentions a new idea, and you suddenly need to take notes. You need a structured document, stored in the right place, **right now**.

Our vision is to make this effortless: instantly generate a well-organized document, with the correct template, stored in the appropriate folder, with key information pre-filled. No manual setup—just a ready-to-use document, so you can focus on content, not formatting or navigating folder structures.

## What dodl Will Do

- **Dynamic Templating**: Build documents using placeholders like `{{date}}`, `{{topic}}`, and more, creating structured notes in seconds.
- **Dynamic Directory Structuring**: Automatically organize documents into the right folders based on context, so everything is always where it belongs.
- **Flexible Date Input**: Need a document templated for a meeting in the future? Specify dates in regular formats (`DD-MM-YY`) or natural language expressions like "next Monday" or "three weeks from now."
- **Seamless Experience**: Focus on simplicity — no unnecessary prompts or input. You get exactly what you need, fast.
- **Flexible Workflows**: Configure templates and file structures to suit personal and project-specific needs.

## Examples

### Set Up Your Workspace

To get started, run `dodl init`, which creates a `.dodl` directory in your current location:

```bash
dodl init
```

Want it somewhere else? Just specify the path:

```bash
dodl init /path/to/your/workspace
```

### Create a Document

Need meeting notes or a project report? Just tell dodl what type of document you want by running:

```bash
dodl create [document_type]
```

This should be a document type configured in your workspace's `config.yaml`. For example, if you have a configured document type `projectX`, you can create it with:

```bash
dodl create projectX
```

You can also add details like date and topic to customize it further:

```bash
dodl create projectX -d "2023-10-28" -t "Weekly Update"
```

### Check Workspace Status

Curious about your workspace setup? Run `dodl status` to see the active configuration, templates, and workspace root path:

```bash
dodl status
```

## The `dodl` Workspace

Running `dodl init` creates a `.dodl` directory in your chosen location, marking that location (the directory containing `.dodl`) as the root of your workspace. This `.dodl` folder includes:

- `config.yaml`: Defines document types and behavior.
- `templates` directory: Holds document templates.

The workspace root controls all subfolders: any `dodl` commands run within the workspace (`create`, `status`, etc.) will use the config and templates from the `.dodl` directory at the root, ensuring consistent document creation and user information.

## The Problem `dodl` Solves

In impromptu meetings or when jotting down ideas, dodl removes the friction from note-taking. It provides the document template you need, pre-filled with relevant data and stored in the right folder, so you can start typing immediately without worrying about setup. The goal is to streamline workflows, letting you focus on capturing ideas, not managing files.

## Status

`dodl` is currently in active development. The basic command system is in place using Cobra commands, and users can now use the `init` command to create a new dodl workspace. A configuration system has been implemented, paving the way for specifying documents that the program can create. Follow along as this vision is brought to life!

## License

`dodl` will be released under the MIT License, ensuring it's free to use, modify, and share (with attribution).
