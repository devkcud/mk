# mk

Simply **m**a**k**e files/folders

This command-line tool allows you to create directories and files using a
straightforward syntax. Directories are created using the `<folder>/` syntax,
where `<folder>` represents the directory name. The `..` syntax is used to remove
the last added directory from the directory stack. You can create empty files by
providing filenames, optionally nested within added directories. **mk** mantains
a directory stack to facilitate creating nested directories and files.
[View example](#examples).

> Tested on Linux (idk if it works on Windows and Mac)

## Building

View [BUILD.md](BUILD.md)

## Usage

`mk [-flags] [..] [folder/] [files...]`

- `-flags`: View mk -help for info
- `..`: Returns the a directory in the stack.
- `folder/`: Creates a directory and add it to the directory stack.
- `files...`: Creates files inside the directory stack.

## Example

### Example 1

Folder structure:

```txt
example/
└── docs/
    └── myfile.txt
```

To create it we use the following command:

```sh
mk example/docs/myfile.txt
```

### Example 2

Folder structure:

```txt
example/
├── tool.go
├── README.md
└── files/
    ├── documents/
    │   └── +page.svelte
    ├── projects/
    │   └── go/
    │       ├── main.go
    │       └── README.md
    └── output.txt
```

To create it we use the following command:

```sh
mk example/ tool.go README.md files/ documents/ +page.svelte .. projects/ go/ main.go README.md ... output.txt
```

```txt
mk
    example/      | create a new folder called example/
    tool.go       | create a new file called tool.go inside example/
    README.md     | create a new file called README.md inside example/
    files/        | create a new folder called files/
    documents/    | create a new folder called documents/ inside files/
    +page.svelte  | create a new file called +page.svelte inside documents/
    ..            | return one directory in the stack
    projects/     | create a new folder called projects/ inside documents/
    go/           | create a new folder called go/ inside projects/
    main.go       | create a new file called main.go inside go/
    README.md     | create a new file called README.md inside go/
    ...           | return two directories in the stack
    output.txt    | create a new file called output.txt inside go/
```

## License

[This project is under the The Unlicense license](LICENSE)
