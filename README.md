# `dodl` 

`dodl` is a command-line tool that makes it super easy to create and manage your documents. Whether you're jotting down meeting minutes, planning your next big project, keeping track of to-do lists, or just taking notes, `dodl` has got you covered.

**Note:** This tool is a work-in-progress and a lot of the functionality isn't implemented yet. Contributions and feedback are welcome!

## What is `dodl`?

`dodl` is all about simplifying your document management. It helps you quickly create documents in the right place with the right content using customisable templates and directory patterns. No more manual file creation and organisation—let `dodl` do the heavy lifting for you!

## Who is `dodl` for?

`dodl` is perfect for:
- **Students** who need to organise lecture notes and study plans.
- **Professionals** who want to keep their meeting notes and project plans neat and tidy.
- **DIY Enthusiasts** who love planning and tracking their projects.
- **Anyone** who wants an easy way to manage their notes, to-do lists, and more.

## Why `dodl`?

- **Quick and Easy**: Create documents in seconds with predefined templates.
- **Stay Organised**: Keep all your notes and plans in the right place.
- **Customisable**: Set up your own templates and directory structures to fit your needs.

## How will `dodl` be used?

### Initialising a New `dodl` Register

Kick things off by setting up the necessary directory structure and configuration files:

```sh
dodl init
```

### Creating a New Document

Generate a new document effortlessly with a simple command:

```sh
dodl new doc todo-list
```

This command will create a new document in the right directory with the right placeholder content for meeting notes.
(`meeting-notes` directory and document content is generated from patterns specified in the config)

### Retrieve a Document

**Get all documents**
```sh
dodl get docs
```

**Get documents with a type**
```sh
dodl get docs --type "meeting-note"
```

Document creation and retrieval will be expanded to support topics, tags, and other metadata.

## Summary

With `dodl`, creating and organising your documents is a breeze. Spend less time managing files and more time being productive!

## Contributing

`dodl` is just getting started, and we’d love your help to make it awesome! If you have ideas, feedback, or want to contribute, jump right in!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

**Remember, this tool is still in development! Your feedback and contributions will help shape its future.**
