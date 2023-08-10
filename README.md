# mk

Simply **m**a**k**e files/folders

This command-line tool allows you to create directories and files using a straightforward syntax. Directories are created using the `+<folder>` syntax, where `<folder>` represents the directory name. The `-` syntax is used to remove the last added directory from the directory stack. You can create empty files by providing filenames, optionally nested within added directories. **mk** mantains a directory stack to facilitate creating nested directories and files. See [example](#Example).

> Tested on Linux (idk if it works on Windows and Mac)

## Usage

`mk [+<folder>] [-] [filename]`

- `+<folder>`: Create a new directory with the specified name and add it to the directory stack.
- `-`: Remove the last added directory from the directory stack.
- `filename`: Create an empty file with the specified name. If directories are added to the directory stack, the file will be created within them. Example: `myfile.txt`

## Example

Folder structure:

```txt
example/
├── tool.go
├── README.md
└── files/
    ├── documents/
    │   └── report.txt
    ├── projects/
    │   └── go/
    │       ├── main.go
    │       └── README.md
    └── output.txt
```

To create it we use the following command:

```sh
mk +example tool.go README.md +files +documents report.txt - +projects +go main.go README.md - - output.txt
```

```txt
mk 
    +example   | Stack: [example]
    tool.go    | Created: example/tool.go
    README.md  | Created: example/README.md
    +files     | Stack: [example, files]
    +documents | Stack: [example, files, documents]
    report.txt | Created: example/files/documents/report.txt
    -          | Stack: [example, files]
    +projects  | Stack: [example, files, projects]
    +go        | Stack: [example, files, projects, go]
    main.go    | Created: example/files/projects/go/main.go
    README.md  | Created: example/files/projects/go/README.md
    -          | Stack: [example, files, projects]
    -          | Stack: [example, files]
    output.txt | Created: example/files/output.txt
```

## License

i don't really care how you're using it :P as it is a simple cli tool
