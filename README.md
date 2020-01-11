# rsrc - Tool for embedding binary resources in Go programs.for windows

this project fork form [rsrc](https://github.com/akavel/rsrc)

INSTALL:

```bash 
go get github.com/srfirouzi/rsrc
```

## USAGE:

rsrc [-i main.rc] [-o rsrc.syso] [-arch amd64]
  Generates a .syso file with specified resources embedded in .rsrc section.
  The .syso file can be linked by Go linker when building Win32 executables.
  Icon embedded this way will show up on application's .exe instead of empty icon.
  Manifest file embedded this way will be recognized and detected by Windows.

The generated *.syso files should get automatically recognized by 'go build'
command and linked into an executable/library, as long as there are any *.go
files in the same directory.

OPTIONS:
- -arch="386": architecture of output file - one of: 386, amd64
- -i="main.rc": path of input rc file
- -o="rsrc.syso": name of output COFF (.res or .syso) file

rc file format,like standard rc format but different
```
id format file
id format file
```
format supprt

- ICON 
- MANIFEST

sample

```
100 ICON "icon.ico"
200 MANIFEST "manifest.manifest"
```

LICENSE: MIT
