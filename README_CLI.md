## CLI

### Installation
```sh
go get github.com/ddddddO/gtree/v6/cmd/gtree
```

or, download from [here](https://github.com/ddddddO/gtree/releases).


### Usage

```sh
20:25:28 > gtree -ts << EOS
> - a
>   - vvv
>     - jjj
>   - kggg
>     - kkdd
>     - tggg
>   - edddd
>     - orrr
>   - gggg
> EOS
a
├── vvv
│   └── jjj
├── kggg
│   ├── kkdd
│   └── tggg
├── edddd
│   └── orrr
└── gggg
```


#### OR

When Markdown data is indented as a tab.

```
├── gtree -f testdata/sample1.md
├── cat testdata/sample1.md | gtree -f -
└── cat testdata/sample1.md | gtree
```

For 2 or 4 spaces instead of tabs, `-ts` or `-fs` is required.


<details>
<summary>More details</summary>

- Usage other than representing a directory.

```sh
16:31:42 > cat testdata/sample2.md | gtree
k8s_resources
├── (Tier3)
│   └── (Tier2)
│       └── (Tier1)
│           └── (Tier0)
├── Deployment
│   └── ReplicaSet
│       └── Pod
│           └── container(s)
├── CronJob
│   └── Job
│       └── Pod
│           └── container(s)
├── (empty)
│   └── DaemonSet
│       └── Pod
│           └── container(s)
└── (empty)
    └── StatefulSet
        └── Pod
            └── container(s)
```

---
- Two spaces indent

```sh
01:15:25 > cat testdata/sample4.md | gtree -ts
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
```

- Four spaces indent

```sh
01:16:46 > cat testdata/sample5.md | gtree -fs
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
```

- Multiple roots

```sh
13:06:26 > cat testdata/sample6.md | gtree
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
```


</details>