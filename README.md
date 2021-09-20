# gtree

[![GitHub release](https://img.shields.io/github/release/ddddddO/gtree.svg)](https://github.com/ddddddO/gtree/releases) [![Go Reference](https://pkg.go.dev/badge/github.com/ddddddO/gtree/v6.svg)](https://pkg.go.dev/github.com/ddddddO/gtree/v6) [![ci](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml/badge.svg)](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml) [![codecov](https://codecov.io/gh/ddddddO/gtree/branch/master/graph/badge.svg?token=JLGSLF33RH)](https://codecov.io/gh/ddddddO/gtree) [![Go Report Card](https://goreportcard.com/badge/github.com/ddddddO/gtree)](https://goreportcard.com/report/github.com/ddddddO/gtree)

(markdown || program) to tree.

```
├── There are three ways to generate tree (CLI, Package(1), Package(2)). They are explained below.
├── CLI and Package(1)
│   ├── Given a Markdown file or format, the result of linux tree command is printed.
│   └── Create Markdown file by referring to the file in the `testdata/` directory.
│       ├── Hierarchy is represented by hyphen and indentation.
│       └── Indentation should be unified by one of the following.
│           ├── Tab (default)
│           ├── Two spaces (required: `-ts`)
│           └── Four spaces (required: `-fs`)
├── Package(1)
│   └── You can customize branch format.
└── Package(2)
    ├── You can also generate a tree programmatically.
    ├── Markdown is irrelevant.
    └── You can customize branch format.
```
generated by `cat testdata/sample0.md | gtree -fs`


---

## CLI
[read me!](https://github.com/ddddddO/gtree/blob/master/README_CLI.md)


## Package(1) / like CLI
[read me!](https://github.com/ddddddO/gtree/blob/master/README_Package_1.md)


## Package(2) / generate a tree programmatically
[read me!](https://github.com/ddddddO/gtree/blob/master/README_Package_2.md)

---

## Documents
- [Markdown形式の入力からtreeを出力するCLI](https://zenn.dev/ddddddo/articles/ad97623a004496)
- [Goでtreeを表現する](https://zenn.dev/ddddddo/articles/8cd85c68763f2e)
